package main

import (
	"book/ch4/upgit/github"
	"fmt"
	"log"
	"os"
	"strings"
)

func validateListArgument(userArgument string) (validatedArgument string){
	userText := strings.ToLower(userArgument)
	switch userText {
	case "all", "open","closed":
		return userText
	default:
		return "all"
	}
}

func displayIssues(issueStatus string){
	result, err := github.ListIssues(issueStatus)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	for _, issue := range *result{
		fmt.Printf("%d\t%-50s\t%s\t%s\n", issue.Number, issue.Title, issue.User.Login, issue.HTMLURL)
	}
}
