package command

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

var startUsage = `Starts the currently active sprint for the specified project.

Usage: mecha start [options]

Options:
  --project Project for which start the active sprint
`

// NewStartCommand returns a new command for a Start operation.
func NewStartCommand() *Command {
	cmd := &Command{
		flags: flag.NewFlagSet("start", flag.ExitOnError),
		Execute: startFunc,
	}

	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, startUsage)
	}

	return cmd
}

var startFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	future, err := project.GetNextSprint()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Starting", future.Name)

	err = future.Start()
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Done")
}
