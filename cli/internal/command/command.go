package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"github.com/ericwimp8/skill-issue/cli/internal/lifecycle"
	"github.com/ericwimp8/skill-issue/cli/internal/payload"
	"github.com/ericwimp8/skill-issue/cli/internal/platform"
)

type BuildInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"build_date"`
}

type App struct {
	stdout    io.Writer
	stderr    io.Writer
	buildInfo BuildInfo
	lifecycle lifecycle.Manager
}

func New(stdout, stderr io.Writer, buildInfo BuildInfo) App {
	return NewWithLifecycle(stdout, stderr, buildInfo, lifecycle.UnavailableManager{})
}

func NewWithLifecycle(stdout, stderr io.Writer, buildInfo BuildInfo, lifecycleManager lifecycle.Manager) App {
	return App{
		stdout:    stdout,
		stderr:    stderr,
		buildInfo: buildInfo,
		lifecycle: lifecycleManager,
	}
}

func (app App) Run(args []string) int {
	if len(args) == 0 {
		app.printHelp()
		return 0
	}

	switch args[0] {
	case "help", "-h", "--help":
		app.printHelp()
		return 0
	case "version", "--version":
		return app.writeJSON(app.buildInfo)
	case "system":
		return app.writeJSON(platform.Current())
	case "payload":
		manifest, err := payload.ReadManifest()
		if err != nil {
			return app.fail(err)
		}
		return app.writeJSON(manifest)
	case "diagnose":
		return app.diagnose()
	case "install", "verify", "repair", "update", "uninstall", "evaluate":
		return app.runLifecycle(lifecycle.Action(args[0]), args[1:])
	case "signal":
		return app.runMarker(args[1:])
	default:
		fmt.Fprintf(app.stderr, "unknown command %q\n", args[0])
		app.printHelpTo(app.stderr)
		return 2
	}
}

func (app App) runMarker(args []string) int {
	_, err := app.lifecycle.Execute(lifecycle.ActionMark, args)
	if err != nil {
		return app.fail(err)
	}
	return 0
}

func (app App) diagnose() int {
	manifest, err := payload.ReadManifest()
	if err != nil {
		return app.fail(err)
	}
	report := struct {
		Build         BuildInfo        `json:"build"`
		Platform      platform.Info    `json:"platform"`
		Payload       payload.Manifest `json:"payload"`
		AdaptersReady bool             `json:"adapters_ready"`
	}{
		Build:         app.buildInfo,
		Platform:      platform.Current(),
		Payload:       manifest,
		AdaptersReady: app.lifecycle.Ready(),
	}
	return app.writeJSON(report)
}

func (app App) runLifecycle(action lifecycle.Action, args []string) int {
	result, err := app.lifecycle.Execute(action, args)
	if errors.Is(err, lifecycle.ErrUnavailable) {
		fmt.Fprintln(app.stderr, err)
		return 2
	}
	if err != nil {
		return app.fail(err)
	}
	return app.writeJSON(result)
}

func (app App) writeJSON(value any) int {
	encoder := json.NewEncoder(app.stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(value); err != nil {
		return app.fail(fmt.Errorf("write output: %w", err))
	}
	return 0
}

func (app App) fail(err error) int {
	fmt.Fprintln(app.stderr, err)
	return 1
}

func (app App) printHelp() {
	app.printHelpTo(app.stdout)
}

func (app App) printHelpTo(writer io.Writer) {
	fmt.Fprintln(writer, `Skill Issue CLI

Usage:
  skill-issue <command>

Foundation commands:
  help        Show this help
  version     Show build metadata
  system      Show operating system and architecture
  payload     Show embedded payload metadata
  diagnose    Show foundation diagnostic information

Lifecycle commands:
  install     Install canonical skills into a supported harness
  verify      Verify an owned installation
  repair      Reinstall an owned canonical payload
  update      Replace an owned canonical payload
  uninstall   Remove an owned installation
  evaluate    Run or clean up a governed scenario
Run "skill-issue help" and see cli/README.md for command arguments.`)
}
