package main

import (
	"fmt"
	"os"

	"github.com/nobbmaestro/beam/cmd"
)

var (
	version = "dev"
	commit  = "unknown"
	date    = "unknown"
)

func main() {
	cmd.SetVersion(version, commit, date)

	err := cmd.Execute()
	if err != nil {
		fmt.Errorf("%w", err)
		os.Exit(1)
	}
}
