package main

import (
	"book/ch4/upgit/github"
	"fmt"
	"log"
	"os"
)

//Writes out detailed information about the GitHub issue number passed in.
func displayIssueDetails(issueNumber string) {
	result, err := github.GetIssueDetails(issueNumber)
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
