package cmds

import (
	"context"
	"fmt"
	"os"

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
		return rm(path.Parse(args[0]))
	},
}

func rm(file path.Path) error {
	if file.VolumePath == "" {
		return os.Remove(file.Path)
	}
	return lfs.WithReadWrite(file, blockSize, blocks, func(fs *littlefs.LFS) error {
		return fs.Remove(file.VolumePath)
	})
}
