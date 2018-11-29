package github

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//ListIssues queries the GitHub issue tracker and returns all issue on our repo.
func ListIssues(stat4 string) (*IssuesListResult, error) {
	resp, err := http.Get(IssuesListURL + "?state=" + stat4)
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
