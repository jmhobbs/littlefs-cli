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
		Subcommands: []*ffcli.Command{cmds.Format, cmds.Ls, cmds.Tree, cmds.Cat, cmds.Cp, cmds.Mv, cmds.Rm},
	}

	if err := root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
