package lifecycle

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/installer"
	"github.com/ericwimp8/skill-issue/cli/internal/runstate"
)

type Action string

const (
	ActionInstall   Action = "install"
	ActionVerify    Action = "verify"
	ActionRepair    Action = "repair"
	ActionUpdate    Action = "update"
	ActionUninstall Action = "uninstall"
	ActionEvaluate  Action = "evaluate"
	ActionMark      Action = "mark"
)

var ErrUnavailable = errors.New("verified harness adapters are not included")

type Result struct {
	Action Action `json:"action"`
	Status string `json:"status"`
	Data   any    `json:"data,omitempty"`
}

type Manager interface {
	Execute(action Action, args []string) (Result, error)
	Ready() bool
}

type Service struct {
	version    string
	installer  installer.Service
	evaluation evaluation.Service
}

func New(stateRoot, version string) Service {
	return Service{
		version:    version,
		installer:  installer.New(stateRoot),
		evaluation: evaluation.New(stateRoot),
	}
}

func (service Service) Ready() bool {
	return true
}

func (service Service) Execute(action Action, args []string) (Result, error) {
	switch action {
	case ActionInstall:
		return service.install(action, args)
	case ActionVerify:
		return service.verify(action, args)
	case ActionRepair, ActionUpdate:
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

func (service Service) install(action Action, args []string) (Result, error) {
	options, err := parseOptions(args)
	if err != nil {
		return Result{}, err
	}
	request, err := installRequest(options, service.version)
	if err != nil {
		return Result{}, err
	}
	installed, err := service.installer.Install(request)
	if err != nil {
		return Result{}, err
	}
	return Result{Action: action, Status: "installed", Data: installed}, nil
}

func (service Service) verify(action Action, args []string) (Result, error) {
	options, err := parseOptions(args)
	if err != nil {
		return Result{}, err
	}
	id, scope, workspace, home, err := installationTarget(options)
	if err != nil {
		return Result{}, err
	}
	receiptID, err := service.installer.ReceiptID(id, scope, workspace, home)
	if err != nil {
		return Result{}, err
	}
	installed, err := service.installer.Verify(receiptID)
	if err != nil {
		return Result{}, err
	}
	return Result{Action: action, Status: "verified", Data: installed}, nil
}

func (service Service) uninstall(action Action, args []string) (Result, error) {
	options, err := parseOptions(args)
	if err != nil {
		return Result{}, err
	}
	id, scope, workspace, home, err := installationTarget(options)
	if err != nil {
		return Result{}, err
	}
	receiptID, err := service.installer.ReceiptID(id, scope, workspace, home)
	if err != nil {
		return Result{}, err
	}
	if err := service.installer.Uninstall(receiptID); err != nil {
		return Result{}, err
	}
	return Result{Action: action, Status: "uninstalled"}, nil
}

func (service Service) evaluate(args []string) (Result, error) {
	if len(args) == 0 {
		return Result{}, errors.New("evaluate requires run or cleanup")
	}
	switch args[0] {
	case "run":
		options, err := parseOptions(args[1:])
		if err != nil {
			return Result{}, err
		}
		id, scope, workspace, _, err := installationTarget(options)
		if err != nil {
			return Result{}, err
		}
		model, err := required(options, "model")
		if err != nil {
			return Result{}, err
		}
		scenario, err := required(options, "scenario")
		if err != nil {
			return Result{}, err
		}
		answer, err := required(options, "answer-sheet")
		if err != nil {
			return Result{}, err
		}
		ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
		defer stop()
		result, err := service.evaluation.Run(ctx, evaluation.RunRequest{
			Workspace:      workspace,
			Harness:        id,
			Model:          model,
			ScenarioPath:   scenario,
			AnswerSheet:    answer,
			Scope:          scope,
			Executable:     options["executable"],
			CLIPath:        options["cli-path"],
			ProductVersion: service.version,
		})
		if err != nil {
			return Result{}, err
		}
		return Result{Action: ActionEvaluate, Status: "complete", Data: result}, nil
	case "cleanup":
		options, err := parseOptions(args[1:])
		if err != nil {
			return Result{}, err
		}
		runID, err := required(options, "run")
		if err != nil {
			return Result{}, err
		}
		if err := service.evaluation.Cleanup(runID); err != nil {
			return Result{}, err
		}
		return Result{Action: ActionEvaluate, Status: "cleaned"}, nil
	default:
		return Result{}, fmt.Errorf("unsupported evaluate command %q", args[0])
	}
}

func (service Service) mark(args []string) (Result, error) {
	if len(args) != 1 {
		return Result{}, errors.New("mark requires one opaque token")
	}
	if err := service.evaluation.Mark(args[0]); err != nil {
		return Result{}, err
	}
	return Result{Action: ActionMark, Status: "recorded"}, nil
}

func installRequest(options map[string]string, version string) (installer.Request, error) {
	id, scope, workspace, home, err := installationTarget(options)
	if err != nil {
		return installer.Request{}, err
	}
	return installer.Request{
		Harness:        id,
		Scope:          scope,
		Workspace:      workspace,
		Home:           home,
		ProductVersion: version,
		Mode:           installer.ModeOrdinary,
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

func parseOptions(args []string) (map[string]string, error) {
	options := map[string]string{}
	for index := 0; index < len(args); index++ {
		argument := args[index]
		if !strings.HasPrefix(argument, "--") {
			return nil, fmt.Errorf("unexpected argument %q", argument)
		}
		key := strings.TrimPrefix(argument, "--")
		if key == "" || index+1 >= len(args) || strings.HasPrefix(args[index+1], "--") {
			return nil, fmt.Errorf("%s requires a value", argument)
		}
		if _, exists := options[key]; exists {
			return nil, fmt.Errorf("duplicate option %s", argument)
		}
		options[key] = args[index+1]
		index++
	}
	return options, nil
}

func required(options map[string]string, key string) (string, error) {
	value := options[key]
	if value == "" {
		return "", fmt.Errorf("--%s is required", key)
	}
	return value, nil
}

type UnavailableManager struct{}

func (UnavailableManager) Execute(action Action, _ []string) (Result, error) {
	return Result{}, fmt.Errorf("%s: %w", action, ErrUnavailable)
}

func (UnavailableManager) Ready() bool {
	return false
}

func DefaultStateRoot() (string, error) {
	return runstate.DefaultRoot()
}
