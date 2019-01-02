package github

import "time"

const IssuesURL = "https://api.github.com/repos/jjsdiner/upgit-simple/issues"

type IssuesListResult []*IssueSummary
type IssueDetailsResult *IssueDetails

//Used to store a summary of an issue on GitHub
type IssueSummary struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
}

//Used to store information about a user on a GitHub issue
type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

//Used to store detailed information about a GitHub issue
type IssueDetails struct {
	Number      int
	HTMLURL     string `json:"html_url"`
	Title       string
	State       string
	User        *User
	CreatedAt   time.Time `json:"created_at"`
	Body        string
	UpdatedAt   *time.Time   `json:"updated_at,omitempty"` //Pointer as can be nil but Time can not be
	ClosedAt    *time.Time   `json:"closed_at,omitempty"`
	PullRequest *PullRequest `json:"pull_request,omitempty"`
}

//Used to store information about a pull request on a GitHub issue.
type PullRequest struct {
	HTMLURL string `json:"html_url"`
}

//The data needed to create a new issue at GitHub.
type NewIssue struct {
	Title       string `json:"title"`
	Description string `json:"body,omitempty"`
}