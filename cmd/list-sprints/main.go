package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

func main() {
	projectName := flag.String("project", "", "Project to show the sprints for")
	sprintState := flag.String("state", "", "Sprint state filter. Can be combined in comma separated values. Valid values: active,future,closed")
	last := flag.Bool("last", false, "Show only the last created sprint.")
	next := flag.Bool("next", false, "Show next sprint name.")
	flag.Parse()

	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(*projectName)
	if err != nil {
		log.Fatalln(err)
	}

	var sprints *jira.SprintList
	if *next {
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
				jira.Sprint{
					Name: nextSprintName,
					State: "-",
				},
			},
		}
	} else {
		if *last {
			lastSprint, err := project.GetLastCreatedSprint()
			if err != nil {
				log.Fatalln(err)
			}

			sprints = &jira.SprintList{
				Values: []jira.Sprint{*lastSprint},
			}
		} else {
			sprints, err = project.GetSprints(*sprintState)
			if err != nil {
				log.Fatalln(err)
			}
		}
	}

	fmt.Printf("Sprints for project %v:\n", project.Name)
	fmt.Printf("%v", sprints)
}
