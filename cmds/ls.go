package cmds

import (
	"context"
	"fmt"

	"github.com/jmhobbs/littlefs-cli/lfs"
	"github.com/jmhobbs/littlefs-cli/path"

	"github.com/peterbourgon/ff/v3/ffcli"
	"tinygo.org/x/tinyfs/littlefs"
)

var List *ffcli.Command = &ffcli.Command{
	Name:       "ls",
	ShortUsage: "littlefs ls <path>",
	ShortHelp:  "List files on a littlefs filesystem.",
	Exec: func(ctx context.Context, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("required 1 argument, got %d", len(args))
		}

		file := path.Parse(args[0])

		return lfs.WithReadOnly(file, func(fs *littlefs.LFS) error {
			lfsPath := file.VolumePath
			if lfsPath == "" {
				lfsPath = "/"
			}

			stat, err := fs.Stat(lfsPath)
			if err != nil {
				return err
			}

			if stat.IsDir() {
				root, err := fs.Open(lfsPath)
				if err != nil {
					return err
				}
				if root.IsDir() {
					files, err := root.Readdir(0)
					if err != nil {
						return err
					}
					for _, f := range files {
						if f.IsDir() {
							fmt.Print(f.Mode(), f.Size(), f.ModTime(), " ", f.Name()+"/", "\n")
						} else {
							fmt.Print(f.Mode(), f.Size(), f.ModTime(), " ", f.Name(), "\n")
						}
					}
				}
			} else {
				fmt.Print(stat.Mode(), stat.Size(), stat.ModTime(), " ", stat.Name(), "\n")
			}

			return nil
		})
	},
}
