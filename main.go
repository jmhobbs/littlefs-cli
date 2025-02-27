package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/jmhobbs/littlefs-cli/cmds"
	"github.com/peterbourgon/ff/v3/ffcli"
)

func main() {
	if err := cmds.Root.ParseAndRun(context.Background(), os.Args[1:]); err != nil {
		if errors.As(err, &ffcli.NoExecError{}) {
			fmt.Println(ffcli.DefaultUsageFunc(cmds.Root))
			os.Exit(0)
		}
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
