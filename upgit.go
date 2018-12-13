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

		default:
			showHelp()
		}
	}

}

func showHelp() {
	fmt.Println("usage: upgit [--version] [--help]")
	fmt.Println("<command> [<args>]")
	fmt.Println("Commands:")
	fmt.Println("   list [state]      List issues on the repo, [all|open|closed] ")
	fmt.Println("   read [issueID]   Displays the details of an issue by ID")
	os.Exit(0)
}

func showVersion() {
	fmt.Println("upgit version 0.1")
	os.Exit(0)
}
