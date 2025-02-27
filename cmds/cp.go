package cmds

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/jmhobbs/littlefs-cli/lfs"
	"github.com/jmhobbs/littlefs-cli/path"

	"github.com/peterbourgon/ff/v3/ffcli"
)

var Copy *ffcli.Command = &ffcli.Command{
	Name:       "cp",
	ShortUsage: "littlefs cp <path> <path>",
	ShortHelp:  "Copy files to, from, or on a littlefs filesystem.",
	Exec: func(ctx context.Context, args []string) error {
		if len(args) != 2 {
			return fmt.Errorf("required 2 argument, got %d", len(args))
		}

		sourceFile := path.Parse(args[0])
		targetFile := path.Parse(args[1])

		var (
			source io.Reader
			sink   io.Writer
		)

		if sourceFile.VolumePath == "" {
			f, err := os.Open(sourceFile.Path)
			if err != nil {
				return err
			}
			defer f.Close()

			source = f
		} else {
			volume, lFS, f, err := lfs.OpenPath(sourceFile, os.O_RDONLY, *blockSize, *blocks)
			if err != nil {
				return err
			}
			defer volume.Close()
			defer lFS.Unmount()
			defer f.Close()

			source = f
		}

		if targetFile.VolumePath == "" {
			f, err := os.Create(targetFile.Path)
			if err != nil {
				return err
			}
			defer f.Close()

			sink = f
		} else {
			volume, lFS, f, err := lfs.OpenPath(targetFile, os.O_RDWR|os.O_CREATE, *blockSize, *blocks)
			if err != nil {
				return err
			}
			defer volume.Close()
			defer lFS.Unmount()
			defer f.Close()

			sink = f
		}

		_, err := io.Copy(sink, source)
		return err
	},
}
