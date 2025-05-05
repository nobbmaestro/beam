package main

import (
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
		os.Exit(1)
	}
}
