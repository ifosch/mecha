package command

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

var moveUsage = `Moves issues from the currently active sprint to the next one for the specified project.

Usage: mecha move [options]

Options:
  --project Project on which move the issues from the active sprint to the new one
  --to-move Issue statuses to move
`

// NewMoveCommand returns the command for the Move operation.
func NewMoveCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("move", flag.ExitOnError),
		Execute: moveFunc,
	}

	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.StringVar(&issueStatusesToMove, "to-move", "", "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, moveUsage)
	}

	return cmd
}

var (
	issueStatusesToMove = ""
)

var moveFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(context.TODO(), os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"))

	project, err := c.FindProject(projectName)
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
	if issueStatusesToMove == "" {
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
			issueStatusesToMove,
			active.Name,
			future.Name,
		)

		issuesToMove, err = active.GetIssuesByStatus(issueStatusesToMove)
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
