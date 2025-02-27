package cmds

import (
	"flag"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var rootFlagSet *flag.FlagSet = flag.NewFlagSet("littlefs", flag.ExitOnError)

var (
	blocks    = rootFlagSet.Int64("block-count", 0, "littlefs block count, 0 is auto-detect")
	blockSize = rootFlagSet.Int64("block-size", 0, "littlefs block size, 0 is auto-detect for devices, or 512 for images")
)

var Root *ffcli.Command = &ffcli.Command{
	FlagSet:    rootFlagSet,
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
