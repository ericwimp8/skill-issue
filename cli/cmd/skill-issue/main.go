package main

import (
	"fmt"
	"os"

	"github.com/ericwimp8/skill-issue/cli/internal/command"
	"github.com/ericwimp8/skill-issue/cli/internal/lifecycle"
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
	stateRoot, err := lifecycle.DefaultStateRoot()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	manager := lifecycle.New(stateRoot, version)
	app := command.NewWithLifecycle(os.Stdout, os.Stderr, buildInfo, manager)
	os.Exit(app.Run(os.Args[1:]))
}
