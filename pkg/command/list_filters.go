package command

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/ifosch/mecha/pkg/jira"
)

var listFiltersUsage = `Lists filters.

Usage: mecha list filters
`

// NewListFiltersCommand returns the command for the ListFilters operation.
func NewListFiltersCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("list-filters", flag.ExitOnError),
		Execute: listFiltersFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listFiltersUsage)
	}

	return cmd
}

var listFiltersFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(context.TODO(), os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"))

	fl, err := c.GetFilters()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("%v", fl)
}
