package cmds

import (
	"context"
	"fmt"

	"github.com/jmhobbs/littlefs-cli/path"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var Move *ffcli.Command = &ffcli.Command{
	Name:       "mv",
	ShortUsage: "littlefs mv <path> <path>",
	ShortHelp:  "Move files to, from, or in a littlefs filesystem.",
	FlagSet:    commonFlagSet,
	Exec: func(ctx context.Context, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("required 2 argument, got %d", len(args))
		}

		sourceFile := path.Parse(args[0])
		targetFile := path.Parse(args[1])

		err := cp(sourceFile, targetFile)
		if err != nil {
			return err
		}

		return rm(sourceFile)
	},
}
