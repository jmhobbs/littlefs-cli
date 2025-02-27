package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jmhobbs/littlefs-cli/cmds"

	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	var root *ffcli.Command
	root = &ffcli.Command{
		ShortUsage: "littlefs <command>",
		Exec: func(ctx context.Context, args []string) error {
			fmt.Println(ffcli.DefaultUsageFunc(root))
			return nil
		},
		Subcommands: []*ffcli.Command{
			cmds.Cat,
			cmds.Copy,
			cmds.Format,
			cmds.List,
			cmds.Move,
			cmds.Remove,
			cmds.Tree,
		},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
