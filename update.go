package main

import (
	"fmt"
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

//Attempts to update the passed issue id, will open an editor to 
func updateIssue (issueNumber int, updateType string){

}