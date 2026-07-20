package main

import (
	"os"

	"github.com/ericwimp8/skill-issue/cli/internal/command"
)

var (
	version   = "dev"
	commit    = "unknown"
	buildDate = "unknown"
)

func main() {
	buildInfo := command.BuildInfo{
		Version:   version,
		Commit:    commit,
		BuildDate: buildDate,
	}
	app := command.New(os.Stdin, os.Stdout, os.Stderr, buildInfo)
	os.Exit(app.Run(os.Args[1:]))
}
