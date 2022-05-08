package command

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

var completeUsage = `Completes the currently active sprint for the specified project.

Usage: mecha complete [options]

Options:
  --project Project for which complete the active sprint
`

// NewCompleteCommand returns a command for the Complete operation.
func NewCompleteCommand() *Command {
	cmd := &Command{
		flags: flag.NewFlagSet("complete", flag.ExitOnError),
		Execute: completeFunc,
	}

	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, completeUsage)
	}

	return cmd
}

var completeFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	active, err := project.GetCurrentSprint()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Setting", active.Name, "complete")

	err = active.Complete()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Done")
}
