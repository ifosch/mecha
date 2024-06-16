package command

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/mecha/pkg/jira"
)

var listProjectsUsage = `Lists projects.

Usage: mecha list projects
`

// NewListProjectsCommand returns the command for the ListProjects operation.
func NewListProjectsCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("list-projects", flag.ExitOnError),
		Execute: listProjectsFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listProjectsUsage)
	}

	return cmd
}

var listProjectsFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(context.TODO(), os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"))

	pl, err := c.GetProjects()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v", pl)
}
