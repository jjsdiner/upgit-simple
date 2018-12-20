package main

import (
	"book/ch4/upgit/github"
	"fmt"
	"log"
	"os"
	"strings"
)

//Validates the passed issue state is one of the three supported by GitHub
func validateListArgument(userArgument string) (validatedArgument string) {
	userText := strings.ToLower(userArgument)
	switch userText {
	case "all", "open", "closed":
		return userText
	default:
		return "all"
	}
}

//Writes out the summary information about GitHub issues in the passed in state
func displayIssues(issueStatus string) {
	result, err := github.ListIssueSummaries(issueStatus)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for _, issue := range *result {
		fmt.Printf("%d\t%-50s\t%s\t%s\t", issue.Number, issue.Title, issue.User.Login, issue.HTMLURL)
		fmt.Printf("%s\n", issue.CreatedAt.Format("2006-01-02T15:04:05"))
	}
}
