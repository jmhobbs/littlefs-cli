package cmds

import (
	"context"
	"fmt"

	"github.com/jmhobbs/littlefs-cli/lfs"
	"github.com/jmhobbs/littlefs-cli/path"
	"github.com/peterbourgon/ff/v3/ffcli"
	"tinygo.org/x/tinyfs/littlefs"
)

var Remove *ffcli.Command = &ffcli.Command{
	Name:       "rm",
	ShortUsage: "littlefs rm <path>",
	ShortHelp:  "Remove files from a littlefs filesystem.",
	FlagSet:    commonFlagSet,
	Exec: func(ctx context.Context, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("required 1 argument, got %d", len(args))
		}
		file := path.Parse(args[0])

		if file.VolumePath == "" {
			return fmt.Errorf("littefs path required")
		}

		return lfs.WithReadWrite(file, blockSize, blocks, func(fs *littlefs.LFS) error {
			return fs.Remove(file.VolumePath)
		})
	},
}
