package cmds

import (
	"github.com/peterbourgon/ff/v3/ffcli"
)

var Root *ffcli.Command = &ffcli.Command{
	ShortUsage: "littlefs <command>",
	Subcommands: []*ffcli.Command{
		Cat,
		Copy,
		Format,
		List,
		Move,
		Remove,
		Tree,
	},
}
