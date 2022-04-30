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
}
