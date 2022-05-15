package command

import (
	"flag"
	"fmt"
	"os"
)

var listUsage = `Lists projects, sprints, and issues.

Usage: mecha list object [options]

Commands:
  projects Lists all projects in the current Jira server
  sprints  Lists sprints for the specified project
  issues   Lists issues in sprints for the specified project
`

// NewListCommand returns the command for all List operations.
func NewListCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("list", flag.ExitOnError),
		Execute: listFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, listUsage)
	}

	return cmd
}

var listFunc = func(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.usageAndExit("mecha: You need to provide a list command.\n")
	}

	switch args[0] {
	case "projects":
		cmd = NewListProjectsCommand()
	case "sprints":
		cmd = NewListSprintsCommand()
	case "issues":
		cmd = NewListIssuesCommand()
	default:
		cmd.usageAndExit(fmt.Sprintf("mecha: '%s' is not a mecha list comand.\n", args[1]))
	}

	cmd.Init(args[1:])
	cmd.Run()
}
