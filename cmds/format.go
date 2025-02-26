package cmds

import (
	"context"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var Format *ffcli.Command = &ffcli.Command{
	Name:       "fmt",
	ShortUsage: "littlefs fmt <device|image>",
	ShortHelp:  "Format a block device, or create a new image file.",
	Exec: func(ctx context.Context, args []string) error {
		println("TODO")
		return nil
	},
}
