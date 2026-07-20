package command

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/lifecycle"
	"github.com/ericwimp8/skill-issue/cli/internal/payload"
	"golang.org/x/term"
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
	case "install":
		return app.runInstall(args[1:])
	case "uninstall", "evaluate":
		return app.runLifecycle(lifecycle.Action(args[0]), args[1:])
	case "signal":
		return app.runMarker(args[1:])
	default:
		fmt.Fprintf(app.stderr, "unknown command %q\n", args[0])
		app.printHelpTo(app.stderr)
		return 2
	}
}

type menuOption struct {
	label    string
	value    string
	disabled bool
}

func (app App) runInstall(args []string) int {
	if len(args) > 0 {
		return app.runLifecycle(lifecycle.ActionInstall, args)
	}
	return app.runGuidedInstall()
}

func (app App) runGuidedInstall() int {
	input, ok := app.stdin.(*os.File)
	if !ok || !term.IsTerminal(int(input.Fd())) {
		return app.fail(fmt.Errorf("guided installation requires an interactive terminal; supply --harness, --scope, and --workspace for scripted installation"))
	}

	harnessOptions := make([]menuOption, 0, len(harness.SupportedIDs()))
	for _, id := range harness.SupportedIDs() {
		available := harness.InstallationAvailable(id)
		label := string(id)
		if !available {
			label += " (in progress)"
		}
		harnessOptions = append(harnessOptions, menuOption{label: label, value: string(id), disabled: !available})
	}
	harnessValue, err := selectMenu(input, app.stderr, "Choose a harness", harnessOptions)
	if err != nil {
		return app.fail(err)
	}
	scopeValue, err := selectMenu(input, app.stderr, "Choose an installation scope", []menuOption{
		{label: "project", value: string(harness.ScopeProject)},
		{label: "user", value: string(harness.ScopeUser)},
	})
	if err != nil {
		return app.fail(err)
	}

	harnessID, err := harness.ParseID(harnessValue)
	if err != nil {
		return app.fail(err)
	}
	scope, err := harness.ParseScope(scopeValue)
	if err != nil {
		return app.fail(err)
	}
	workspace, err := os.Getwd()
	if err != nil {
		return app.fail(fmt.Errorf("resolve current workspace: %w", err))
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return app.fail(fmt.Errorf("resolve home directory: %w", err))
	}
	destination, err := harness.SkillRoot(harnessID, scope, workspace, home)
	if err != nil {
		return app.fail(err)
	}
	skills, err := payload.Skills()
	if err != nil {
		return app.fail(err)
	}

	fmt.Fprintln(app.stderr, "Installation ready:")
	fmt.Fprintf(app.stderr, "  harness: %s\n", harnessID)
	fmt.Fprintf(app.stderr, "  scope: %s\n", scope)
	if scope == harness.ScopeProject {
		fmt.Fprintf(app.stderr, "  workspace: %s\n", workspace)
	}
	fmt.Fprintf(app.stderr, "  destination: %s\n", destination)
	fmt.Fprintln(app.stderr, "  skills:")
	for _, skill := range skills {
		fmt.Fprintf(app.stderr, "    - %s\n", skill.Name)
	}
	confirmed, err := readConfirmation(app.stdin, app.stderr, "Install these skills? [y/N]: ")
	if err != nil {
		return app.fail(err)
	}
	if !confirmed {
		fmt.Fprintln(app.stderr, "installation cancelled")
		return 0
	}

	installArgs := []string{"--harness", string(harnessID), "--scope", string(scope)}
	if scope == harness.ScopeProject {
		installArgs = append(installArgs, "--workspace", workspace)
	}
	result, err := app.lifecycle.Execute(lifecycle.ActionInstall, installArgs)
	if err != nil {
		return app.fail(err)
	}
	if code := app.writeJSON(result); code != 0 {
		return code
	}
	fmt.Fprintf(app.stderr, "Installed and verified %d skills.\n", len(skills))
	fmt.Fprintf(app.stderr, "Open %s and explicitly invoke skill-intake to begin.\n", harnessID)
	return 0
}

func selectMenu(input *os.File, output io.Writer, title string, options []menuOption) (string, error) {
	if len(options) == 0 {
		return "", errors.New("selection menu has no options")
	}
	selected := firstEnabledOption(options)
	if selected < 0 {
		return "", errors.New("selection menu has no available options")
	}
	state, err := term.MakeRaw(int(input.Fd()))
	if err != nil {
		return "", fmt.Errorf("open interactive selection: %w", err)
	}
	defer term.Restore(int(input.Fd()), state)

	fmt.Fprintf(output, "%s\r\nUse the up and down arrows, then press Enter.\r\n", title)
	renderMenu(output, options, selected, false)
	for {
		key, err := readMenuKey(input)
		if err != nil {
			return "", err
		}
		switch key {
		case "up":
			selected = moveSelection(options, selected, -1)
			renderMenu(output, options, selected, true)
		case "down":
			selected = moveSelection(options, selected, 1)
			renderMenu(output, options, selected, true)
		case "select":
			fmt.Fprint(output, "\r\n")
			return options[selected].value, nil
		case "cancel":
			return "", errors.New("installation selection cancelled")
		}
	}
}

func renderMenu(output io.Writer, options []menuOption, selected int, redraw bool) {
	if redraw {
		fmt.Fprintf(output, "\x1b[%dA", len(options))
	}
	for index, option := range options {
		indicator := "  "
		if index == selected {
			indicator = "> "
		} else if option.disabled {
			indicator = "- "
		}
		fmt.Fprintf(output, "\r\x1b[2K%s%s\r\n", indicator, option.label)
	}
}

func firstEnabledOption(options []menuOption) int {
	for index, option := range options {
		if !option.disabled {
			return index
		}
	}
	return -1
}

func moveSelection(options []menuOption, selected, direction int) int {
	for offset := 1; offset <= len(options); offset++ {
		candidate := (selected + direction*offset + len(options)) % len(options)
		if !options[candidate].disabled {
			return candidate
		}
	}
	return selected
}

func readMenuKey(input io.Reader) (string, error) {
	var current [1]byte
	if _, err := io.ReadFull(input, current[:]); err != nil {
		return "", fmt.Errorf("read interactive selection: %w", err)
	}
	switch current[0] {
	case '\r', '\n':
		return "select", nil
	case 3:
		return "cancel", nil
	case 0, 0xe0:
		if _, err := io.ReadFull(input, current[:]); err != nil {
			return "", fmt.Errorf("read interactive selection: %w", err)
		}
		if current[0] == 72 {
			return "up", nil
		}
		if current[0] == 80 {
			return "down", nil
		}
	case 0x1b:
		var sequence [2]byte
		if _, err := io.ReadFull(input, sequence[:]); err != nil {
			return "", fmt.Errorf("read interactive selection: %w", err)
		}
		if (sequence[0] == '[' || sequence[0] == 'O') && sequence[1] == 'A' {
			return "up", nil
		}
		if (sequence[0] == '[' || sequence[0] == 'O') && sequence[1] == 'B' {
			return "down", nil
		}
	}
	return "", nil
}

func readConfirmation(input io.Reader, output io.Writer, prompt string) (bool, error) {
	fmt.Fprint(output, prompt)
	reader := bufio.NewReader(input)
	answer, err := reader.ReadString('\n')
	if err != nil && err != io.EOF {
		return false, fmt.Errorf("read installation confirmation: %w", err)
	}
	switch strings.ToLower(strings.TrimSpace(answer)) {
	case "y", "yes":
		fmt.Fprintln(output)
		return true, nil
	default:
		return false, nil
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
