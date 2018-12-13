package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//ListIssueSummaries queries the GitHub API and returns a list of all issues in the passed state.
func ListIssueSummaries(state string) (*IssuesListResult, error) {
	resp, err := http.Get(IssuesURL + "?state=" + state)
	if err != nil {
		return nil, err
	}

	//We must close resp.Body on all execution paths.
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result IssuesListResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}

//GetIssueDetails queries the GitHub API and gets detailed information on specific issue.
func GetIssueDetails(issueNumber string) (*IssueDetailsResult, error){
	resp, err := http.Get(IssuesURL + "/" + issueNumber)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed:%s", resp.Status)
	}

	var result IssueDetailsResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &result, nil
}