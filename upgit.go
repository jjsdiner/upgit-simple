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
			result, err := github.ListIssues()
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}
			fmt.Println((*result)[0])

		default:
			showHelp()
		}
	}

}

func showHelp() {
	fmt.Println("usage: upgit [--version] [--help]")
	fmt.Println("<command> [<args>]")
	os.Exit(0)
}

func showVersion() {
	fmt.Println("upgit version 0.1")
	os.Exit(0)
}
