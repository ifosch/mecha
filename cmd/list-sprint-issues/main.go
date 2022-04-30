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
