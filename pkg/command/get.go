package command

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

var getUsage = `Gets stats for active and future sprints in specified project.

Usage: mecha get [options]

Options:
  --project Project to show the sprints for
  --state   Sprint state filter. Can be combined in comma separated values. Valid values: active,future,closed
`

// NewGetCommand returns a command for the Get operation.
func NewGetCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("get", flag.ExitOnError),
		Execute: getFunc,
	}

	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.StringVar(&sprintState, "state", "", "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, getUsage)
	}

	return cmd
}

var (
	projectName = ""
	sprintState = ""
)

var getFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(context.TODO(), os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"))

	project, err := c.FindProject(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	sprints, err := project.GetSprints(sprintState)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Sprint Stats for project %v:\n", project.Name)
	missingSP := 0
	for _, s := range sprints.Values {
		fmt.Printf("- %v (%v)\n", s.Name, s.State)
		ss, err := s.GetStats()
		if err != nil {
			log.Fatalln(err)
		}

		for status, counts := range ss {
			if counts["Missing SP"] > 0 {
				missingSP += counts["Missing SP"]
			}
			fmt.Printf("  %v/%v stories/SP %v\n", counts["Stories"], counts["SP"], status)
		}
	}
	if missingSP > 0 {
		fmt.Println("There are", missingSP, "stories without SP")
	}

	os.Exit(0)
}
