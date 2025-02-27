package cmds

import (
	"context"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var Move *ffcli.Command = &ffcli.Command{
	Name:       "mv",
	ShortUsage: "littlefs mv <path> <path>",
	ShortHelp:  "Move files to, from, or in a littlefs filesystem.",
	Exec: func(ctx context.Context, args []string) error {
		println("TODO")
		return nil
	},
}
