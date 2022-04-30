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
	flag.Parse()

	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(*projectName)
	if err != nil {
		log.Fatalln(err)
	}

	sprints, err := project.GetSprints(*sprintState)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("Sprints for project %v:\n", project.Name)
	fmt.Printf("%v", sprints)
}
