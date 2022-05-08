package command

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

var addUsage = `Adds the currently active sprint for the specified project.

Usage: mecha add [options]

Options:
  --project Project for which add the active sprint
`

// NewAddCommand returns the command for Add operation.
func NewAddCommand() *Command {
	cmd := &Command{
		flags: flag.NewFlagSet("add", flag.ExitOnError),
		Execute: addFunc,
	}

	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, addUsage)
	}

	return cmd
}

var addFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	newSprint, err := project.CreateSprint()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(newSprint, "created")
}
