package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

func main() {
	projectName := flag.String("project", "", "Project for which active sprint must complete")
	flag.Parse()

	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(*projectName)
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
