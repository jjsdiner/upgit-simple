package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

//Creates a new issue at GitHub and returns writes out the URL of the new issue.
func createNewIssue (issueTitle string) {

	tmpDir := os.TempDir()
	tmpFile, tmpFileErr := ioutil.TempFile(tmpDir, "upgit" + time.Now().Format("2006-01-02T15:04:05"))
	if tmpFileErr != nil {
		log.Fatal(tmpFileErr)
		os.Exit(1)
	}

	path, err := exec.LookPath("vim")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//Detailed input for the issue will need an editor
	cmd := exec.Command(path, tmpFile.Name())
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Run()

	descriptionBytes, _ := ioutil.ReadFile(tmpFile.Name())
	description := string(descriptionBytes)


	fmt.Println("Issue Title:" + issueTitle)
	fmt.Println(description)
}
