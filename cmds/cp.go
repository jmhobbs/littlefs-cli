package cmds

import (
	"context"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var Cp *ffcli.Command = &ffcli.Command{
	Name:       "cp",
	ShortUsage: "littlefs cp <path> <path>",
	ShortHelp:  "Copy files to, from, or on a littlefs filesystem.",
	Exec: func(ctx context.Context, args []string) error {
		println("TODO")
		return nil
	},
}
