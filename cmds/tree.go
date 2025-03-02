package cmds

import (
	"context"
	"fmt"
	sysfs "io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/jmhobbs/littlefs-cli/lfs"
	"github.com/jmhobbs/littlefs-cli/path"
	"github.com/jmhobbs/littlefs-cli/tree"

	"github.com/peterbourgon/ff/v3/ffcli"
	"tinygo.org/x/tinyfs/littlefs"
)

var Tree *ffcli.Command = &ffcli.Command{
	Name:       "tree",
	ShortUsage: "littlefs tree <path>",
	ShortHelp:  "Walk a littlefs filesystem and print a tree view.",
	FlagSet:    commonFlagSet,
	Exec: func(ctx context.Context, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("required 1 argument, got %d", len(args))
		}

		file := path.Parse(args[0])

		return lfs.WithReadOnly(file, blockSize, blocks, func(fs *littlefs.LFS) error {
			root := file.VolumePath

			if !strings.HasSuffix(root, string(filepath.Separator)) {
				root = root + string(filepath.Separator)
			}

			var (
				fileCount int        = 0
				dirCount  int        = 1
				files     *tree.Node = tree.NewNode(root)
			)

			err := sysfs.WalkDir(lfs.NewLFS_Sysfs(fs), root, func(path string, info sysfs.DirEntry, err error) error {
				if err != nil {
					return err
				}

				if info.IsDir() {
					dirCount++
				} else {
					fileCount++
				}

				if path != root {
					files.Insert(strings.TrimPrefix(path, root))
				}

				return nil
			})

			if err == nil {
				files.Print(os.Stdout)
				fmt.Printf("\n%d directories, %d files\n", dirCount, fileCount)
			}

			return err
		})
	},
}
