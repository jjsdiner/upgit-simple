package github

import (
	"book/ch4/github"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//Posts new issue from user to GitHub
func CreateIssue (issue *NewIssue) (*IssueDetailsResult, error){
	issueJson, err := json.Marshal(issue)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	req, err := http.NewRequest("POST", github.IssuesURL, bytes.NewBuffer(issueJson))
	req.Header.Set("Authorization", "token "+os.Getenv("UPGITUSER"))
	req.Header.Set( "Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusCreated {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		body := string(bodyBytes)
		resp.Body.Close()
		return nil, fmt.Errorf("create issue failed:%s", resp.Status + "\ntext: " + body)
	}
	var result IssueDetailsResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil

}
