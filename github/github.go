package github

import "time"

const IssuesURL = "https://api.github.com/repos/jjsdiner/upgit-simple/issues"

type IssuesListResult []*IssueSummary
type IssueDetailsResult *IssueDetails

type IssueSummary struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type IssueDetails struct {
	Number        int
	HTMLURL       string `json:"html_url"`
	Title         string
	State         string
	User          *User
	CreatedAt     time.Time `json:"created_at"`
	Body          string
	UpdatedAt     *time.Time `json:"updated_at,omitempty"` //Pointer as can be nil but Time can not be
	ClosedAt      *time.Time `json:"closed_at,omitempty"`
	PullRequest   *PullRequest `json:"pull_request,omitempty"`
}

type PullRequest struct {
	HTMLURL   string `json:"html_url"`
}
