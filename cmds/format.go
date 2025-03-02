package cmds

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/jmhobbs/littlefs-cli/lfs"
	"github.com/jmhobbs/littlefs-cli/path"
	"github.com/peterbourgon/ff/v3/ffcli"
)

var formatFlagSet *flag.FlagSet = flag.NewFlagSet("littlefs format", flag.ExitOnError)

var force = formatFlagSet.Bool("force", false, "do not ask for confirmation")

func init() {
	sharedFlags(formatFlagSet)
}

var Format *ffcli.Command = &ffcli.Command{
	Name:       "fmt",
	ShortUsage: "littlefs fmt [flags] <device|image>",
	ShortHelp:  "Format a block device, or create a new image file.",
	FlagSet:    formatFlagSet,
	Exec: func(ctx context.Context, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("required 1 argument, got %d", len(args))
		}

		file := path.Parse(args[0])
		if file.VolumePath != "" {
			return fmt.Errorf("cannot format a path on an existing littlefs filesystem")
		}

		fmt.Printf("formatting %q as littlefs, this will DESTROY anything which exists there. Continue? [y/N]", file.Path)
		reader := bufio.NewReader(os.Stdin)
		line, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		if line != "Y\n" && line != "y\n" {
			fmt.Println("aborting")
			return nil
		}

		err = lfs.Create(file.Path, blockSize, blocks)
		if err != nil {
			return err
		}
		fmt.Println("filesystem created")

		return nil
	},
}
