package command

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strings"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
	"github.com/ericwimp8/skill-issue/cli/internal/lifecycle"
)

type BuildInfo struct {
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	BuildDate string `json:"build_date"`
}

type App struct {
	stdin     io.Reader
	stdout    io.Writer
	stderr    io.Writer
	buildInfo BuildInfo
	lifecycle lifecycle.Service
}

func New(stdin io.Reader, stdout, stderr io.Writer, buildInfo BuildInfo) App {
	return App{
		stdin:     stdin,
		stdout:    stdout,
		stderr:    stderr,
		buildInfo: buildInfo,
		lifecycle: lifecycle.New(),
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
	case "install", "uninstall", "evaluate":
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

func (app App) runLifecycle(action lifecycle.Action, args []string) int {
	if action == lifecycle.ActionEvaluate {
		app.warnEvaluation(args)
	}
	if action == lifecycle.ActionEvaluate && len(args) > 0 && args[0] == "run" {
		return app.runEvaluation(args)
	}
	result, err := app.lifecycle.Execute(action, args)
	if err != nil {
		return app.fail(err)
	}
	return app.writeJSON(result)
}

func (app App) runEvaluation(args []string) int {
	result, err := app.lifecycle.ExecuteEvaluationRun(args[1:], app.reviewEvaluation)
	if err != nil {
		return app.fail(err)
	}
	return app.writeJSON(result)
}

func (app App) reviewEvaluation(request evaluation.RunRequest) (bool, error) {
	fmt.Fprintln(app.stderr, "evaluation ready:")
	if request.EvaluationID != "" {
		fmt.Fprintf(app.stderr, "  evaluation: %s (built-in)\n", request.EvaluationID)
	} else {
		fmt.Fprintf(app.stderr, "  evaluation: %s (custom)\n", filepath.Base(request.ScenarioPath))
	}
	if request.TurnLimit > request.AvailableTurns {
		fmt.Fprintf(app.stderr, "  turns: %d of %d available (requested %d)\n", request.EffectiveTurns, request.AvailableTurns, request.TurnLimit)
	} else {
		fmt.Fprintf(app.stderr, "  turns: %d of %d available\n", request.EffectiveTurns, request.AvailableTurns)
	}
	fmt.Fprintf(app.stderr, "  harness: %s\n", request.Harness)
	fmt.Fprintf(app.stderr, "  model: %s\n", request.Model)
	fmt.Fprintf(app.stderr, "  reasoning: %s\n", request.Reasoning)
	fmt.Fprintf(app.stderr, "  workspace: %s\n", request.Workspace)
	fmt.Fprintf(app.stderr, "  output: %s\n", request.OutputRoot)
	if request.Executable != "" {
		fmt.Fprintf(app.stderr, "  executable: %s\n", request.Executable)
	}
	if request.EvaluationID == "" {
		fmt.Fprintf(app.stderr, "  skills: %s\n", request.SkillsPath)
		fmt.Fprintf(app.stderr, "  scenario: %s\n", request.ScenarioPath)
		fmt.Fprintf(app.stderr, "  answer-sheet: %s\n", request.AnswerSheet)
		fmt.Fprintln(app.stderr, "warning: answer-sheet correctness is caller-owned; an incorrect key can make the result invalid or misleading")
	}
	fmt.Fprint(app.stderr, "start evaluation? [y/N]: ")
	reader := bufio.NewReader(app.stdin)
	answer, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return false, fmt.Errorf("read evaluation confirmation: %w", err)
	}
	switch strings.ToLower(strings.TrimSpace(answer)) {
	case "y", "yes":
		fmt.Fprintln(app.stderr)
		return true, nil
	default:
		fmt.Fprintln(app.stderr, "evaluation cancelled")
		return false, nil
	}
}

func (app App) warnEvaluation(args []string) {
	if len(args) == 0 || args[0] != "run" {
		return
	}
	if hasOption(args[1:], "scenario") || hasOption(args[1:], "answer-sheet") {
		fmt.Fprintln(app.stderr, "warning: custom evaluation inputs must not contain personal, confidential, or sensitive information; the harness conversation may retain their full contents")
	}
	if hasOption(args[1:], "transcript") {
		fmt.Fprintln(app.stderr, "warning: --transcript writes the evaluation conversation, including prompts, responses, commands, command output, errors, and harness events; known local paths and machine identifiers are sanitized, but personal or confidential conversation content may remain, so review transcript.json before sharing")
	}
}

func hasOption(args []string, name string) bool {
	wanted := "--" + name
	for _, argument := range args {
		if argument == wanted {
			return true
		}
	}
	return false
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

Lifecycle commands:
  install     Install canonical skills into a supported harness
  uninstall   Remove canonical Skill Issue skills
  evaluate    Run or clean up a governed scenario
Run "skill-issue help" and see cli/README.md for command arguments.`)
}
