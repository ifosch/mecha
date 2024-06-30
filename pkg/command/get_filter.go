package command

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/ifosch/mecha/pkg/jira"
)

var getFilterUsage = `Gets filter

Usage: mecha get filter ID [options]
`

// NewGetFilterCommand returns a command for the Get Filter operation.
func NewGetFilterCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("get-filter", flag.ExitOnError),
		Execute: getFilterFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, getFilterUsage)
	}

	return cmd
}

var getFilterFunc = func(cmd *Command, args []string) {
	c := jira.NewClient(context.TODO(), os.Getenv("JIRA_URL"), os.Getenv("JIRA_USERNAME"), os.Getenv("JIRA_API_TOKEN"))

	ID, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, err := c.GetFilter(ID)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(f)

	os.Exit(0)
}
