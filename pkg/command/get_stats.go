package command

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/mecha/pkg/jira"
)

var getStatsUsage = `Gets stats for active and future sprints in specified project.

Usage: mecha get stats [options]

Options:
  --project Project to show the sprints for
  --state   Sprint state filter. Can be combined in comma separated values. Valid values: active,future,closed
`

// NewGetStatsCommand returns a command for the Get Stats operation.
func NewGetStatsCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("get-stats", flag.ExitOnError),
		Execute: getStatsFunc,
	}

	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.StringVar(&sprintState, "state", "", "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, getStatsUsage)
	}

	return cmd
}

var (
	projectName = ""
	sprintState = ""
)

var getStatsFunc = func(cmd *Command, args []string) {
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
