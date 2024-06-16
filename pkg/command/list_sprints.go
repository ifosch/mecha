package command

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/mecha/pkg/jira"
)

var listSprintsUsage = `Lists sprints.

Usage: mecha list sprints

Options:
  project Project to show the sprints for
  state   Sprint state filter. Can be combined in comma separated values. Valid values: active,future,closed
  last    Shows only the last created sprint. Defaults to false
  next    Shows next sprint name. Defaults to false
`

// NewListSprintsCommand returns the command for the ListSprints operation.
func NewListSprintsCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("list-sprints", flag.ExitOnError),
		Execute: listSprintsFunc,
	}

	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.StringVar(&sprintState, "state", "", "")
	cmd.flags.BoolVar(&last, "last", false, "")
	cmd.flags.BoolVar(&next, "next", false, "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listSprintsUsage)
	}

	return cmd
}

var (
	last = false
	next = false
)

var listSprintsFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(context.TODO(), os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"))

	project, err := c.FindProject(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	var sprints *jira.SprintList
	if next {
		lastSprint, err := project.GetLastCreatedSprint()
		if err != nil {
			log.Fatalln(err)
		}

		nextSprintName, err := lastSprint.NextSprintName()
		if err != nil {
			log.Fatalln(err)
		}
		sprints = &jira.SprintList{
			Values: []jira.Sprint{
				{
					Name:  nextSprintName,
					State: "-",
				},
			},
		}
	} else {
		if last {
			lastSprint, err := project.GetLastCreatedSprint()
			if err != nil {
				log.Fatalln(err)
			}

			sprints = &jira.SprintList{
				Values: []jira.Sprint{*lastSprint},
			}
		} else {
			sprints, err = project.GetSprints(sprintState)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	fmt.Printf("Sprints for project %v:\n", project.Name)
	fmt.Printf("%v", sprints)
}
