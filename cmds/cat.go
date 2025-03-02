package cmds

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jmhobbs/littlefs-cli/lfs"
	"github.com/jmhobbs/littlefs-cli/path"

	"github.com/peterbourgon/ff/v3/ffcli"
	"tinygo.org/x/tinyfs/littlefs"
)

var Cat *ffcli.Command = &ffcli.Command{
	Name:       "cat",
	ShortUsage: "littlefs cat <path>",
	ShortHelp:  "Output contents of a file on a littlefs filesystem.",
	FlagSet:    commonFlagSet,
	Exec: func(ctx context.Context, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("required 1 argument, got %d", len(args))
		}
		file := path.Parse(args[0])

		if file.VolumePath == "" {
			return fmt.Errorf("littefs path required")
		}

		return lfs.WithReadOnly(file, blockSize, blocks, func(fs *littlefs.LFS) error {
			stat, err := fs.Stat(file.VolumePath)
			if err != nil {
				return err
			}

			if stat.IsDir() {
				return fmt.Errorf("cannot cat a directory")
			}

			f, err := fs.OpenFile(file.VolumePath, os.O_RDONLY)
			if err != nil {
				return err
			}
			defer f.Close()

			_, err = io.Copy(os.Stdout, f)
			return err
		})
	},
}
