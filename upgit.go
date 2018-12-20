//Program to allow the user to create, read, update and close Github issues.
package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	helpPtr := flag.Bool("help", false, "User requests usage information")
	verPtr := flag.Bool("version", false, "User request version information")
	flag.Parse()
	helpCheck := *helpPtr
	verCheck := *verPtr
	if helpCheck {
		showHelp()
	}
	if verCheck {
		showVersion()
	}
	if len(os.Args) < 2 {
		showHelp()
	} else {
		switch os.Args[1] {
		case "list":
			issueStatus := "all"
			if len(os.Args) > 2 {
				issueStatus = os.Args[2]
				issueStatus = validateListArgument(issueStatus)
			}
			displayIssues(issueStatus)
		case "read":
			issueID := ""
			if len(os.Args) > 2 {
				issueID = os.Args[2]
				displayIssueDetails(issueID)
			} else {
				showHelp()
			}
		case "create":
			issueTitle := ""
			switch len(os.Args){
			case 3: {
				issueTitle = os.Args[2]
			}
			default:
				showHelp()
			}
			createNewIssue(issueTitle)
		default:
			showHelp()
		}
	}

}

//Writes out the help information for the application
func showHelp() {
	fmt.Println("usage: upgit [--version] [--help]")
	fmt.Println("<command> [<args>]")
	fmt.Println("Commands:")
	fmt.Println("   list [state]      List issues on the repo, [all|open|closed]")
	fmt.Println("   read [issueNumber]   Displays the details of an issue by number")
	fmt.Println("   create ['Issue Title'] Creates a new issue, opens vi to add description")
	os.Exit(0)
}

//Returns the current version number
func showVersion() {
	fmt.Println("upgit version 0.1.1")
	os.Exit(0)
}
