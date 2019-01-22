package main

import (
	"upgit-simple/github"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

//Creates a new issue at GitHub and returns writes out the URL of the new issue.
func createNewIssue (issueTitle string) {

	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "upgit" + time.Now().Format("2006-01-02T15:04:05"))
	if tmpFileErr != nil {
		log.Fatal(tmpFileErr)
		os.Exit(1)
	}

	path, err := exec.LookPath("vim")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//Detailed input for the issue will need an editor
	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()

	descriptionBytes, _ := ioutil.ReadFile(tmpFile.Name())
	description := string(descriptionBytes)

	//Post to git hub
    var issue github.NewIssue
	issue.Title = issueTitle
	issue.Description =  description

	result, err := github.CreateIssue(&issue)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	details := *result
	fmt.Printf("Issue Number: %d\n", details.Number)
	fmt.Printf("Title:        %s\n", details.Title)
	fmt.Printf("URL:          %s\n", details.HTMLURL)
	fmt.Printf("State:        %s\n", details.State)
	fmt.Printf("Created By:   %s\n", details.User.Login)
	fmt.Println("Created On:   " + details.CreatedAt.Format("2006-01-02T15:04:05"))
	if details.UpdatedAt != nil {
		fmt.Println("Last Updated: " + details.UpdatedAt.Format("2006-01-02T15:04:05"))
	}
	if details.ClosedAt != nil {
		fmt.Println("Closed:       " + details.ClosedAt.Format("2006-01-02T15:04:05"))
	}
	if details.PullRequest != nil {
		fmt.Println("Pull Request: " + details.PullRequest.HTMLURL)
	}
	fmt.Println("Description:  " + details.Body)
}
