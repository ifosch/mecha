package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ifosch/jira/pkg/jira"
)

func main() {
	c := jira.NewClient(os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"), nil)

	pl, err := c.GetProjects()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v", pl)
}
