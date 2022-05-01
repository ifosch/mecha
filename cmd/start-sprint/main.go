package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

func main() {
	projectName := flag.String("project", "", "Project for which first future sprint must start")
	flag.Parse()

	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	project, err := c.FindProject(*projectName)
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
