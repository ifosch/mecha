package command

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

var listIssuesUsage = `Lists issues.

Usage: mecha list issues

Options:
  project Project to show the issues for
  state   Sprint state filter. Can be combined in comma separated values. Valid values: active,future,closed
`

// NewListIssuesCommand returns the command for the ListIssues operation.
func NewListIssuesCommand() *Command {
	cmd := &Command{
		flags: flag.NewFlagSet("list-issues", flag.ExitOnError),
		Execute: listIssuesFunc,
	}

 	cmd.flags.StringVar(&projectName, "project", "", "")
	cmd.flags.StringVar(&sprintState, "state", "", "")
	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listIssuesUsage)
	}

	return cmd
}

var listIssuesFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(projectName)
	if err != nil {
		log.Fatalln(err)
	}

	sprints, err := project.GetSprints(sprintState)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Sprints for project %v:\n", project.Name)
	for _, s := range sprints.Values {
		fmt.Printf("- %v (%v)\n", s.Name, s.State)
		il, err := s.GetIssues()
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println(il)
	}
}
