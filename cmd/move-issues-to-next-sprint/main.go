package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

func main() {
	projectName := flag.String("project", "", "Project on which move the issues from the active sprint to the new one")
	issueStatusesToMove := flag.String("to-move", "", "Issue statuses to move")
	flag.Parse()

	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(*projectName)
	if err != nil {
		log.Fatalln(err)
	}

	active, err := project.GetCurrentSprint()
	if err != nil {
		log.Fatalln(err)
	}

	future, err := project.GetNextSprint()
	if err != nil {
		log.Fatalln(err)
	}

	var issuesToMove *jira.IssueList
	if *issueStatusesToMove == "" {
		fmt.Printf(
			"Moving all issues from %v to %v\n",
			active.Name,
			future.Name,
		)

		issuesToMove, err = active.GetIssues()
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		fmt.Printf(
			"Moving %v issues from %v to %v\n",
			*issueStatusesToMove,
			active.Name,
			future.Name,
		)

		issuesToMove, err = active.GetIssuesByStatus(*issueStatusesToMove)
		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Printf("%s", issuesToMove)

	err = active.MoveIssuesToNextSprint(issuesToMove)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Done")
}
