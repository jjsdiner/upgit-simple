package github

import "time"

const IssuesListURL = "https://api.github.com/repos/jjsdiner/upgit-simple/issues"

type IssuesListResult []*Issue

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string    // in Markdown format
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
