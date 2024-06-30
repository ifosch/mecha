package command

import (
	"flag"
	"fmt"
	"os"
)

var getUsage = `Get stats or filters.

Usage: mecha get object [options]

Commands:
  stats    Gets stats for sprints (active or future) on a project
  filter   Gets filter
`

// NewGetCommand returns the command for all Get operations.
func NewGetCommand() *Command {
	cmd := &Command{
		flags:   flag.NewFlagSet("get", flag.ExitOnError),
		Execute: getFunc,
	}

	cmd.flags.Usage = func() {
		fmt.Fprintln(os.Stderr, getUsage)
	}

	return cmd
}

var getFunc = func(cmd *Command, args []string) {
	if len(args) == 0 {
		cmd.usageAndExit("mecha: You need to provide a get command.\n")
	}

	switch args[0] {
	case "stats":
		cmd = NewGetStatsCommand()
	case "filter":
		cmd = NewGetFilterCommand()
	default:
		cmd.usageAndExit(fmt.Sprintf("mecha: '%s' is not a mecha get comand.\n", args[0]))
	}

	cmd.Init(args[1:])
	cmd.Run()
}
