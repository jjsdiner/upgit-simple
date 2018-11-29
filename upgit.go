//Program to allow the user to create, read, update and close Github issues.
package main

import (
	"book/ch4/upgit/github"
	"flag"
	"fmt"
	"log"
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
			issueStat4 := "all"
			if len(os.Args) >2{
				issueStat4 = os.Args[2]
			}
			result, err := github.ListIssues(issueStat4)
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			for _, issue := range *result{
				fmt.Printf("%d\t%-50s\t%s\t%s\n", issue.Number, issue.Title, issue.User.Login, issue.HTMLURL)
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
	fmt.Println("   list [state]      List issues on the repo, [all|open|closed]"   )
	os.Exit(0)
}

func showVersion() {
	fmt.Println("upgit version 0.1")
	os.Exit(0)
}
