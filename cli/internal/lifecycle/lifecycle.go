package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/signal"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/installer"
)

type Action string

const (
	ActionInstall   Action = "install"
	ActionUninstall Action = "uninstall"
	ActionEvaluate  Action = "evaluate"
	ActionMark      Action = "mark"
)

type Result struct {
	Action Action `json:"action"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}

type Service struct {
	installer installer.Service
	progress  io.Writer
}

type EvaluationReviewer func(evaluation.RunRequest) (bool, error)

type SkillReplacementConfirmer func(differing []string) (bool, error)

// New returns a lifecycle service that reports evaluation progress to the
// given writer; pass nil to discard progress output.
func New(progress io.Writer) Service {
	if progress == nil {
		progress = io.Discard
	}
	return Service{
		installer: installer.New(),
		progress:  progress,
	}
}

func (service Service) Execute(action Action, args []string) (Result, error) {
	switch action {
	case ActionInstall:
		return service.install(action, args)
	case ActionUninstall:
		return service.uninstall(action, args)
	case ActionEvaluate:
		return service.evaluate(args)
	case ActionMark:
		return service.mark(args)
	default:
		return Result{}, fmt.Errorf("unsupported lifecycle action %q", action)
	}
}

func (service Service) ExecuteEvaluationRun(args []string, reviewer EvaluationReviewer, confirmReplacement SkillReplacementConfirmer) (Result, error) {
	return service.evaluationRun(args, reviewer, confirmReplacement)
}

func (service Service) install(action Action, args []string) (Result, error) {
	options, err := parseOptions(args)
	if err != nil {
		return Result{}, err
	}
	request, err := installRequest(options)
	if err != nil {
		return Result{}, err
	}
	if !harness.InstallationAvailable(request.Harness) {
		return Result{}, fmt.Errorf("%s installation is still in progress", request.Harness)
	}
	installed, err := service.installer.Install(request)
	if err != nil {
		return Result{}, err
	}
	return Result{Action: action, Status: "installed", Data: installed}, nil
}

func (service Service) uninstall(action Action, args []string) (Result, error) {
	options, err := parseOptions(args)
	if err != nil {
		return Result{}, err
	}
	request, err := installRequest(options)
	if err != nil {
		return Result{}, err
	}
	removed, err := service.installer.Uninstall(request)
	if err != nil {
		return Result{}, err
	}
	return Result{Action: action, Status: "uninstalled", Data: removed}, nil
}

func (service Service) evaluate(args []string) (Result, error) {
	if len(args) == 0 {
		return Result{}, errors.New("evaluate requires run or cleanup")
	}
	switch args[0] {
	case "run":
		return service.evaluationRun(args[1:], nil, nil)
	case "cleanup":
		options, err := parseOptions(args[1:])
		if err != nil {
			return Result{}, err
		}
		runID, err := required(options, "run")
		if err != nil {
			return Result{}, err
		}
		output, err := requiredOutput(options)
		if err != nil {
			return Result{}, err
		}
		if err := evaluation.New(evaluationStateRoot(output)).Cleanup(runID); err != nil {
			return Result{}, err
		}
		return Result{Action: ActionEvaluate, Status: "cleaned"}, nil
	default:
		return Result{}, fmt.Errorf("unsupported evaluate command %q", args[0])
	}
}

func (service Service) evaluationRun(args []string, reviewer EvaluationReviewer, confirmReplacement SkillReplacementConfirmer) (Result, error) {
	options, err := parseOptions(args, "events", "transcript", "replace-preexisting-skills", "yes")
	if err != nil {
		return Result{}, err
	}
	request, err := evaluationRunRequest(options)
	if err != nil {
		return Result{}, err
	}
	request.ConfirmPreexisting = confirmReplacement
	if reviewer != nil {
		confirmed, err := reviewer(request)
		if err != nil || !confirmed {
			if request.WorkspaceCreated {
				// The created workspace is still empty; leave no trace behind
				// a cancelled or failed confirmation.
				_ = os.Remove(request.Workspace)
			}
			if err != nil {
				return Result{}, err
			}
			return Result{Action: ActionEvaluate, Status: "cancelled"}, nil
		}
	}
	return service.runEvaluation(request)
}

func (service Service) runEvaluation(request evaluation.RunRequest) (Result, error) {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()
	renderer := newTurnProgressRenderer(service.progress)
	defer renderer.stop()
	request.Progress = renderer.handle
	result, err := evaluation.New(evaluationStateRoot(request.OutputRoot)).Run(ctx, request)
	if err != nil {
		return Result{}, err
	}
	return Result{Action: ActionEvaluate, Status: "complete", Data: result}, nil
}

func (service Service) mark(args []string) (Result, error) {
	if len(args) != 2 {
		return Result{}, errors.New("signal requires one opaque token and state root")
	}
	if !filepath.IsAbs(args[1]) {
		return Result{}, errors.New("signal state root must be absolute")
	}
	if err := evaluation.New(args[1]).Mark(args[0]); err != nil {
		return Result{}, err
	}
	return Result{Action: ActionMark, Status: "recorded"}, nil
}

// defaultEvaluationID is the built-in scenario an evaluation runs when the
// caller selects none.
const defaultEvaluationID = "gardening-web-application"

func evaluationRunRequest(options map[string]string) (evaluation.RunRequest, error) {
	harnessValue, err := required(options, "harness")
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	id, err := harness.ParseEvaluationID(harnessValue)
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	output := options["output"]
	if output == "" {
		output = "skill-issue-output"
	}
	output, err = filepath.Abs(output)
	if err != nil {
		return evaluation.RunRequest{}, fmt.Errorf("resolve output directory: %w", err)
	}
	evaluationID := options["evaluation"]
	skills, err := optionalAbsolutePath(options, "skills")
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	scenario, err := optionalAbsolutePath(options, "scenario")
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	answer, err := optionalAbsolutePath(options, "answer-sheet")
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	turnLimit, err := optionalPositiveInteger(options, "turns")
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	executable, err := resolvedExecutableOption(options["executable"])
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	if evaluationID != "" && (skills != "" || scenario != "" || answer != "") {
		return evaluation.RunRequest{}, errors.New("--evaluation cannot be combined with --skills, --scenario, or --answer-sheet")
	}
	if evaluationID == "" {
		if skills == "" && scenario == "" && answer == "" {
			evaluationID = defaultEvaluationID
		} else if skills == "" || scenario == "" || answer == "" {
			return evaluation.RunRequest{}, errors.New("use --evaluation or supply --skills, --scenario, and --answer-sheet")
		}
	}
	workspace, workspaceCreated, err := preparedWorkspace(options["workspace"])
	if err != nil {
		return evaluation.RunRequest{}, err
	}
	request := evaluation.RunRequest{
		Workspace:          workspace,
		WorkspaceCreated:   workspaceCreated,
		OutputRoot:         output,
		Harness:            id,
		Model:              options["model"],
		ModelOverride:      options["model"] != "",
		Reasoning:          options["reasoning"],
		ReasoningOverride:  options["reasoning"] != "",
		EvaluationID:       evaluationID,
		SkillsPath:         skills,
		ScenarioPath:       scenario,
		AnswerSheet:        answer,
		Executable:         executable,
		CLIPath:            options["cli-path"],
		IncludeEvents:      options["events"] == "true",
		IncludeTranscript:  options["transcript"] == "true",
		ReplacePreexisting: options["replace-preexisting-skills"] == "true",
		AssumeYes:          options["yes"] == "true",
		TurnLimit:          turnLimit,
	}
	prepared, err := evaluation.PrepareRequest(request)
	if err != nil {
		if workspaceCreated {
			_ = os.Remove(workspace)
		}
		return evaluation.RunRequest{}, err
	}
	return prepared, nil
}

// preparedWorkspace resolves the evaluation workspace. Without a --workspace
// the CLI creates a fresh uniquely named directory adjacent to the invocation
// directory, so every defaulted run starts clean; a named workspace that does
// not exist yet is created for the caller.
func preparedWorkspace(configured string) (string, bool, error) {
	if configured == "" {
		created, err := os.MkdirTemp(".", "skill-issue-workspace-")
		if err != nil {
			return "", false, fmt.Errorf("create default workspace: %w", err)
		}
		absolute, err := filepath.Abs(created)
		if err != nil {
			return "", false, fmt.Errorf("resolve default workspace: %w", err)
		}
		return absolute, true, nil
	}
	absolute, err := filepath.Abs(configured)
	if err != nil {
		return "", false, fmt.Errorf("resolve workspace: %w", err)
	}
	if _, statErr := os.Stat(absolute); statErr == nil {
		return absolute, false, nil
	} else if !errors.Is(statErr, fs.ErrNotExist) {
		return "", false, fmt.Errorf("inspect workspace: %w", statErr)
	}
	if err := os.MkdirAll(absolute, 0o700); err != nil {
		return "", false, fmt.Errorf("create workspace: %w", err)
	}
	return absolute, true, nil
}

// resolvedExecutableOption absolutizes a path-form --executable against the
// invocation working directory at parse time. The evaluator later launches
// the harness from a run-owned directory, where a caller-relative path would
// fail at turn 1 with a cryptic fork/exec error. Bare command names keep
// resolving through PATH.
func resolvedExecutableOption(value string) (string, error) {
	if value == "" || !strings.ContainsRune(value, os.PathSeparator) {
		return value, nil
	}
	absolute, err := filepath.Abs(value)
	if err != nil {
		return "", fmt.Errorf("resolve --executable: %w", err)
	}
	return absolute, nil
}

func optionalAbsolutePath(options map[string]string, key string) (string, error) {
	value := options[key]
	if value == "" {
		return "", nil
	}
	absolute, err := filepath.Abs(value)
	if err != nil {
		return "", fmt.Errorf("resolve --%s: %w", key, err)
	}
	return absolute, nil
}

func optionalPositiveInteger(options map[string]string, key string) (int, error) {
	value := options[key]
	if value == "" {
		return 0, nil
	}
	number, err := strconv.Atoi(value)
	if err != nil || number < 1 {
		return 0, fmt.Errorf("--%s must be a positive integer", key)
	}
	return number, nil
}

func requiredOutput(options map[string]string) (string, error) {
	output, err := required(options, "output")
	if err != nil {
		return "", err
	}
	output, err = filepath.Abs(output)
	if err != nil {
		return "", fmt.Errorf("resolve output directory: %w", err)
	}
	return output, nil
}

func evaluationStateRoot(outputRoot string) string {
	return filepath.Join(outputRoot, ".skill-issue")
}

func installRequest(options map[string]string) (installer.Request, error) {
	id, scope, workspace, home, err := installationTarget(options)
	if err != nil {
		return installer.Request{}, err
	}
	return installer.Request{
		Harness:   id,
		Scope:     scope,
		Workspace: workspace,
		Home:      home,
	}, nil
}

func installationTarget(options map[string]string) (harness.ID, harness.Scope, string, string, error) {
	harnessValue, err := required(options, "harness")
	if err != nil {
		return "", "", "", "", err
	}
	id, err := harness.ParseID(harnessValue)
	if err != nil {
		return "", "", "", "", err
	}
	scopeValue, err := required(options, "scope")
	if err != nil {
		return "", "", "", "", err
	}
	scope, err := harness.ParseScope(scopeValue)
	if err != nil {
		return "", "", "", "", err
	}
	workspace := options["workspace"]
	if scope == harness.ScopeProject && workspace == "" {
		return "", "", "", "", errors.New("--workspace is required for project scope")
	}
	if workspace != "" {
		workspace, err = filepath.Abs(workspace)
		if err != nil {
			return "", "", "", "", fmt.Errorf("resolve workspace: %w", err)
		}
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", "", "", "", fmt.Errorf("resolve home directory: %w", err)
	}
	return id, scope, workspace, home, nil
}

func parseOptions(args []string, flags ...string) (map[string]string, error) {
	options := map[string]string{}
	boolean := map[string]bool{}
	for _, flag := range flags {
		boolean[flag] = true
	}
	for index := 0; index < len(args); index++ {
		argument := args[index]
		if !strings.HasPrefix(argument, "--") {
			return nil, fmt.Errorf("unexpected argument %q", argument)
		}
		key, inlineValue, hasInline := strings.Cut(strings.TrimPrefix(argument, "--"), "=")
		if key == "" {
			return nil, fmt.Errorf("invalid option %q", argument)
		}
		if _, exists := options[key]; exists {
			return nil, fmt.Errorf("duplicate option --%s", key)
		}
		if boolean[key] {
			if hasInline && inlineValue != "true" && inlineValue != "false" {
				return nil, fmt.Errorf("--%s accepts only true or false", key)
			}
			if hasInline {
				options[key] = inlineValue
			} else {
				options[key] = "true"
			}
			continue
		}
		if hasInline {
			options[key] = inlineValue
			continue
		}
		if index+1 >= len(args) || strings.HasPrefix(args[index+1], "--") {
			return nil, fmt.Errorf("%s requires a value", argument)
		}
		options[key] = args[index+1]
		index++
	}
	return options, nil
}

func required(options map[string]string, key string) (string, error) {
	value, ok := options[key]
	if !ok {
		return "", fmt.Errorf("--%s is required", key)
	}
	if value == "" {
		return "", fmt.Errorf("--%s must not be empty", key)
	}
	return value, nil
}
