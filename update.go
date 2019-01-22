package main

import (
	"book/ch4/upgit/github"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

//Checks which field the user asked to update and if it is valid
func validateUpdateType (userUpdateType string) (validatedUpdateType string, err error){
	userUpdateType = strings.ToLower(userUpdateType)
	switch userUpdateType {
	case "title", "body", "state":
		return userUpdateType, nil
	default:
		return "", fmt.Errorf("invalid update field: %s", userUpdateType)
	}
}

//Attempts to update the passed issue id, will open an editor if the description is being edited
func updateIssue (issueNumber string, updateType string){

	res, err := github.GetIssueDetails(issueNumber)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	selectedIssue := *res
	switch updateType {
	case	"state":
	    fmt.Printf("Issue %s current status: %s enter open | closed to set state.\n", issueNumber, selectedIssue.State)
        reader := bufio.NewReader(os.Stdin)
        status, _ := reader.ReadString('\n')
        fmt.Println(status)
	default:
		//Should not be possible but crash out
		log.Fatalf("Invalid update type reached update command.")
		os.Exit(1)
	}

}