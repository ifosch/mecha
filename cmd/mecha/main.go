package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ifosch/jira/pkg/command"
)

var usage = `Usage: mecha command [options]

A Jira CLI focused in project management tasks.

Commands:
  list     Lists projects, sprints, and issues
  add      Adds a new sprint to specified project
  get      Gets stats for active and future sprints in specified project
  move     Moves all issues in currently active sprint to next one for the specified project
  start    Starts the next sprint for the specified project
  complete Completes the active sprint in the specified project
`

func usageAndExit(msg string) {
	if msg != "" {
		fmt.Fprint(os.Stderr, msg)
		fmt.Fprintf(os.Stderr, "\n")
	}

	flag.Usage()
	os.Exit(1)
}

func main() {
	flag.Usage = func() {
		fmt.Fprint(os.Stderr, usage)
	}

	var cmd *command.Command

	if len(os.Args) == 1 {
		usageAndExit("mecha: You need to provide a command.\n")
	}

	switch os.Args[1] {
	case "get":
		cmd = command.NewGetCommand()
	case "list":
		cmd = command.NewListCommand()
	case "move":
		cmd = command.NewMoveCommand()
	case "add":
		cmd = command.NewAddCommand()
	case "complete":
		cmd = command.NewCompleteCommand()
	case "start":
		cmd = command.NewStartCommand()
	default:
		usageAndExit(fmt.Sprintf("mecha: '%s' is not a mecha comand.\n", os.Args[1]))
	}

	cmd.Init(os.Args[2:])
	cmd.Run()
}
