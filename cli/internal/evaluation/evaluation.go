package evaluation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/installer"
	"github.com/ericwimp8/skill-issue/cli/internal/payload"
	"github.com/ericwimp8/skill-issue/cli/internal/replay"
	"github.com/ericwimp8/skill-issue/cli/internal/runstate"
)

type SkillCall struct {
	TurnID string `json:"turn_id"`
	Skill  string `json:"skill"`
}

type AnswerSheet struct {
	SchemaVersion int         `json:"schema_version"`
	ScenarioID    string      `json:"scenario_id"`
	Expected      []SkillCall `json:"expected"`
}

type RunRequest struct {
	Workspace         string
	OutputRoot        string
	Harness           harness.ID
	Model             string
	ModelOverride     bool
	Reasoning         string
	ReasoningOverride bool
	EvaluationID      string
	SkillsPath        string
	ScenarioPath      string
	AnswerSheet       string
	Executable        string
	CLIPath           string
	IncludeEvents     bool
	IncludeTranscript bool
	TurnLimit         int
	AvailableTurns    int
	EffectiveTurns    int
	Progress          func(TurnProgress)
}

type TurnProgress struct {
	TurnID string
	Index  int
	Total  int
	Phase  replay.BoundaryPhase
}

func ResolveRequest(request RunRequest) (RunRequest, error) {
	defaults, err := harness.EvaluationDefaultsFor(request.Harness)
	if err != nil {
		return RunRequest{}, err
	}
	if request.Harness == harness.Cursor && request.ReasoningOverride {
		return RunRequest{}, errors.New("cursor does not support an independent --reasoning override; omit --reasoning to use model-native reasoning")
	}
	if request.Model == "" {
		request.Model = defaults.Model
	}
	if request.Reasoning == "" {
		request.Reasoning = defaults.Reasoning
	}
	return request, nil
}

func PrepareRequest(request RunRequest) (RunRequest, error) {
	request, err := ResolveRequest(request)
	if err != nil {
		return RunRequest{}, err
	}
	if request.TurnLimit < 0 {
		return RunRequest{}, errors.New("--turns must be a positive integer")
	}
	inputs, err := loadEvaluationInputs(request)
	if err != nil {
		return RunRequest{}, err
	}
	request.AvailableTurns = len(inputs.scenario.Turns)
	request.EffectiveTurns = request.AvailableTurns
	if request.TurnLimit > 0 && request.TurnLimit < request.AvailableTurns {
		request.EffectiveTurns = request.TurnLimit
	}
	return request, nil
}

type BuiltInEvaluation struct {
	SchemaVersion int             `json:"schema_version"`
	EvaluationID  string          `json:"evaluation_id"`
	Scenario      replay.Scenario `json:"scenario"`
	AnswerSheet   AnswerSheet     `json:"answer_sheet"`
}

type loadedInputs struct {
	scenario replay.Scenario
	answer   AnswerSheet
	skills   []payload.Skill
}

type Result struct {
	SchemaVersion  int         `json:"schema_version"`
	RunID          string      `json:"run_id"`
	Harness        string      `json:"harness"`
	Model          string      `json:"model"`
	Reasoning      string      `json:"reasoning"`
	EvaluationID   string      `json:"evaluation_id"`
	ScenarioID     string      `json:"scenario_id"`
	Scope          string      `json:"scope"`
	StartedAt      time.Time   `json:"started_at"`
	CompletedAt    time.Time   `json:"completed_at"`
	Expected       []SkillCall `json:"expected"`
	Observed       []SkillCall `json:"observed"`
	Missing        []SkillCall `json:"missing"`
	Additional     []SkillCall `json:"additional"`
	Unattributed   []SkillCall `json:"unattributed"`
	TranscriptPath string      `json:"transcript_path,omitempty"`
}

type WebsitePoint struct {
	Turn   int    `json:"turn"`
	TurnID string `json:"turn_id"`
	Called int    `json:"called"`
	Missed int    `json:"missed"`
}

type WebsiteResult struct {
	SchemaVersion int            `json:"schema_version"`
	RunID         string         `json:"run_id"`
	ScenarioID    string         `json:"scenario_id"`
	Harness       string         `json:"harness"`
	Model         string         `json:"model"`
	TotalTurns    int            `json:"total_turns"`
	Points        []WebsitePoint `json:"points"`
}

type Service struct {
	stateRoot      string
	runs           runstate.Store
	installer      installer.Service
	adapterFactory func(replay.HarnessID, replay.Options) (replay.Adapter, error)
}

func New(stateRoot string) Service {
	return Service{
		stateRoot:      stateRoot,
		runs:           runstate.NewStore(stateRoot),
		installer:      installer.New(),
		adapterFactory: replay.NewAdapter,
	}
}

func (service Service) Run(ctx context.Context, request RunRequest) (result Result, err error) {
	request, err = PrepareRequest(request)
	if err != nil {
		return Result{}, err
	}
	workspace, err := filepath.Abs(request.Workspace)
	if err != nil {
		return Result{}, fmt.Errorf("resolve evaluation workspace: %w", err)
	}
	workspace, err = filepath.EvalSymlinks(workspace)
	if err != nil {
		return Result{}, fmt.Errorf("canonicalize evaluation workspace: %w", err)
	}
	if info, statErr := os.Stat(workspace); statErr != nil || !info.IsDir() {
		return Result{}, errors.New("evaluation workspace must be an existing directory")
	}
	request.Workspace = workspace
	outputRoot, err := prepareOutputRoot(workspace, request.OutputRoot)
	if err != nil {
		return Result{}, err
	}
	inputs, err := loadEvaluationInputs(request)
	if err != nil {
		return Result{}, err
	}
	inputs = limitEvaluationInputs(inputs, request.EffectiveTurns)
	scenario, answer := inputs.scenario, inputs.answer
	cliPath, err := executablePath(request.CLIPath)
	if err != nil {
		return Result{}, err
	}
	if request.Harness == harness.OpenCode || request.Harness == harness.KiloCode {
		request.Executable, err = runtimeExecutable(request.Harness, request.Executable)
		if err != nil {
			return Result{}, err
		}
	}
	tokens, err := skillTokens(inputs.skills)
	if err != nil {
		return Result{}, err
	}
	runID, err := runstate.NewRunID()
	if err != nil {
		return Result{}, err
	}
	startedAt := time.Now().UTC()
	outputDirectory := filepath.Join(outputRoot, fmt.Sprintf("%s-%s-%s", request.Harness, startedAt.Format("20060102T150405Z"), runID[:8]))
	if err := os.Mkdir(outputDirectory, 0o700); err != nil {
		return Result{}, fmt.Errorf("create evaluation output directory: %w", err)
	}
	run := runstate.Run{
		SchemaVersion:     1,
		ID:                runID,
		Workspace:         workspace,
		Harness:           string(request.Harness),
		Model:             request.Model,
		Reasoning:         request.Reasoning,
		EvaluationID:      evaluationIdentity(request, scenario),
		Scenario:          scenario.ID,
		Scope:             string(harness.ScopeProject),
		Status:            "preparing",
		HarnessExecutable: request.Executable,
		Tokens:            tokens,
	}
	if err := service.runs.Create(run); err != nil {
		return Result{}, err
	}
	run, err = service.runs.Load(runID)
	if err != nil {
		return Result{}, err
	}
	defer func() {
		if err != nil {
			_ = service.runs.DeletePrivateMappings(runID)
			_ = service.runs.DeleteRun(runID)
			err = fmt.Errorf("evaluation run %s: %w", runID, err)
		}
	}()
	skillNames := make([]string, 0, len(inputs.skills))
	for _, skill := range inputs.skills {
		skillNames = append(skillNames, skill.Name)
	}
	runtime, err := service.prepareRuntime(request.Harness, request.Model, request.Reasoning, runID, workspace, request.Executable, cliPath, skillNames)
	if err != nil {
		return Result{}, err
	}
	if runtime.signalExecutable != "" {
		cliPath = runtime.signalExecutable
	}
	defer os.RemoveAll(privateRuntimeRunRoot(runID))
	if err := replay.CheckAuthentication(ctx, replay.HarnessID(request.Harness), request.Executable, request.Model, runtime.environment, request.Harness == harness.Cursor || request.Harness == harness.OpenCode || request.Harness == harness.KiloCode || request.Harness == harness.Pi); err != nil {
		return Result{}, fmt.Errorf("evaluation encountered a tooling error: %w", err)
	}
	installationState, _, err := service.installer.PrepareEvaluation(installer.Request{
		Harness:              request.Harness,
		Scope:                harness.ScopeProject,
		Workspace:            workspace,
		EvaluationRoot:       runtime.evaluationSkillRoot,
		CLIPath:              cliPath,
		SignalStateRoot:      service.stateRoot,
		Tokens:               tokens,
		Skills:               inputs.skills,
		ApplyHarnessMetadata: request.EvaluationID != "",
	})
	if err != nil {
		return Result{}, err
	}
	cleaned := false
	defer func() {
		if !cleaned {
			cleanupErr := service.cleanupWithInstallation(runID, installationState)
			if err == nil && cleanupErr != nil {
				err = cleanupErr
			}
		}
	}()
	installationStatePath := filepath.Join(service.runs.RunDir(runID), "installation-state.json")
	installationStateData, err := installer.EncodeEvaluationInstallation(installationState)
	if err != nil {
		return Result{}, err
	}
	if err := os.WriteFile(installationStatePath, append(installationStateData, '\n'), 0o600); err != nil {
		return Result{}, fmt.Errorf("write evaluation installation state: %w", err)
	}
	runtime.environment = append(runtime.environment, "PWD="+runtime.workingDirectory)
	if request.Harness == harness.OpenCode {
		if err := replay.CheckOpenCodeSkills(ctx, request.Executable, runtime.workingDirectory, runtime.environment, false, skillNames); err != nil {
			return Result{}, fmt.Errorf("evaluation encountered a tooling error: %w", err)
		}
	}
	run.InstallationState = installationStatePath
	run.Status = "running"
	if err := service.runs.Save(run); err != nil {
		return Result{}, err
	}
	adapter, err := service.adapterFactory(replay.HarnessID(request.Harness), replay.Options{
		Executable:            request.Executable,
		Directory:             runtime.workingDirectory,
		Environment:           runtime.environment,
		CleanEnvironment:      request.Harness == harness.Cursor || request.Harness == harness.KiloCode || request.Harness == harness.Pi,
		Model:                 request.Model,
		ModelOverride:         request.ModelOverride,
		Reasoning:             request.Reasoning,
		ReasoningOverride:     request.ReasoningOverride,
		CodexConfiguration:    runtime.codexConfiguration,
		CursorPluginDir:       runtime.cursorPluginDir,
		ClaudeSettings:        runtime.claudeSettings,
		ClaudeSkillsRoot:      runtime.claudeSkillsRoot,
		ClaudeWorkspacePrompt: runtime.claudeWorkspacePrompt,
		PiSkillsRoot:          runtime.piSkillsRoot,
		SkillIssueExecutable:  cliPath,
	})
	if err != nil {
		service.setStatus(runID, "tooling-failed")
		return Result{}, fmt.Errorf("evaluation encountered a tooling error: %w", err)
	}
	runner := replay.Runner{
		Adapter: adapter,
		OnBoundary: func(_ context.Context, boundary replay.Boundary) error {
			if boundary.Phase == replay.BoundaryBefore {
				if request.Progress != nil {
					request.Progress(TurnProgress{TurnID: boundary.TurnID, Index: boundary.TurnIndex, Total: boundary.TurnTotal, Phase: boundary.Phase})
				}
				return service.runs.SetActiveTurn(runID, boundary.TurnID)
			}
			if boundary.Capture != nil && boundary.Capture.SessionID != "" {
				if err := service.runs.SetHarnessSession(runID, boundary.Capture.SessionID); err != nil {
					return err
				}
			}
			if request.Harness == harness.Codex && boundary.Capture != nil {
				if err := service.recordCodexSignals(runID, boundary.TurnID, *boundary.Capture, tokens, cliPath); err != nil {
					return err
				}
			}
			if request.Harness == harness.Cursor && boundary.Capture != nil {
				if err := service.validateCursorSignals(runID, boundary.TurnID, *boundary.Capture, tokens, cliPath); err != nil {
					return err
				}
			}
			if err := service.runs.SetActiveTurn(runID, ""); err != nil {
				return err
			}
			if request.Progress != nil {
				request.Progress(TurnProgress{TurnID: boundary.TurnID, Index: boundary.TurnIndex, Total: boundary.TurnTotal, Phase: boundary.Phase})
			}
			return nil
		},
	}
	replayResult, err := runner.Run(ctx, scenario)
	if err != nil {
		service.setStatus(runID, "tooling-failed")
		return Result{}, fmt.Errorf("evaluation encountered a tooling error: %w", err)
	}
	if request.IncludeTranscript {
		sanitizer, err := newTranscriptSanitizer(transcriptSanitizerConfig{
			Workspace:   workspace,
			OutputRoot:  outputRoot,
			StateRoot:   service.stateRoot,
			RuntimeRoot: privateRuntimeRunRoot(runID),
			CLIPath:     cliPath,
		})
		if err != nil {
			return Result{}, err
		}
		if err := sanitizer.sanitize(&replayResult); err != nil {
			return Result{}, err
		}
	}
	events, err := service.runs.Events(runID)
	if err != nil {
		return Result{}, err
	}
	result = deriveResult(runID, request, evaluationIdentity(request, scenario), scenario.ID, startedAt, answer.Expected, events)
	if request.IncludeEvents {
		eventsPath := filepath.Join(outputDirectory, "events.jsonl")
		if err := writeEventsJSONL(eventsPath, events); err != nil {
			return Result{}, err
		}
	}
	if request.IncludeTranscript {
		transcriptPath := filepath.Join(outputDirectory, "transcript.json")
		if err := writeJSON(transcriptPath, replayResult); err != nil {
			return Result{}, err
		}
		result.TranscriptPath = filepath.Base(transcriptPath)
	}
	evidencePath := filepath.Join(outputDirectory, "result.json")
	if err := writeJSON(evidencePath, result); err != nil {
		return Result{}, err
	}
	websitePath := filepath.Join(outputDirectory, "website.json")
	if err := writeJSON(websitePath, deriveWebsiteResult(result, scenario)); err != nil {
		return Result{}, err
	}
	run, err = service.runs.Load(runID)
	if err != nil {
		return Result{}, err
	}
	run.EvidencePath = evidencePath
	if request.IncludeTranscript {
		run.TranscriptPath = filepath.Join(outputDirectory, result.TranscriptPath)
	}
	run.Status = "complete"
	if err := service.runs.Save(run); err != nil {
		return Result{}, err
	}
	if err := service.cleanupWithInstallation(runID, installationState); err != nil {
		return Result{}, err
	}
	cleaned = true
	return result, nil
}

func (service Service) recordCodexSignals(runID, turnID string, capture replay.Capture, tokens map[string]string, cliPath string) error {
	existing, err := service.runs.Events(runID)
	if err != nil {
		return err
	}
	observed := map[string]bool{}
	for _, event := range existing {
		if event.TurnID == turnID {
			observed[event.Skill] = true
		}
	}
	recorded := map[string]bool{}
	for _, event := range capture.Events {
		var value struct {
			Type string `json:"type"`
			Item struct {
				Type    string `json:"type"`
				Command string `json:"command"`
			} `json:"item"`
		}
		if json.Unmarshal(event, &value) != nil || value.Item.Type != "command_execution" {
			continue
		}
		for token := range tokens {
			skill := tokens[token]
			if recorded[token] || observed[skill] || !strings.Contains(value.Item.Command, cliPath) || !strings.Contains(value.Item.Command, " signal ") || !strings.Contains(value.Item.Command, token) || !strings.Contains(value.Item.Command, service.stateRoot) {
				continue
			}
			if err := service.runs.Mark(token); err != nil {
				return err
			}
			recorded[token] = true
		}
	}
	return nil
}

func (service Service) validateCursorSignals(runID, turnID string, capture replay.Capture, tokens map[string]string, cliPath string) error {
	existing, err := service.runs.Events(runID)
	if err != nil {
		return err
	}
	observed := map[string]bool{}
	for _, event := range existing {
		if event.TurnID == turnID {
			observed[event.Skill] = true
		}
	}
	for _, event := range capture.Events {
		var value struct {
			Type     string `json:"type"`
			Subtype  string `json:"subtype"`
			ToolCall struct {
				ShellToolCall struct {
					Args struct {
						Command string `json:"command"`
					} `json:"args"`
				} `json:"shellToolCall"`
			} `json:"tool_call"`
		}
		if json.Unmarshal(event, &value) != nil || value.Type != "tool_call" || value.Subtype != "started" {
			continue
		}
		command := value.ToolCall.ShellToolCall.Args.Command
		for token, skill := range tokens {
			if observed[skill] || !isSignalCommand(command, cliPath, token, service.stateRoot) {
				continue
			}
			detail := cursorSignalCompletionDetail(capture.Events, command)
			return fmt.Errorf("Cursor attempted the instrumented signal for skill %q but no marker was recorded: %s", skill, detail)
		}
	}
	return nil
}

func cursorSignalCompletionDetail(events []json.RawMessage, command string) string {
	for _, event := range events {
		var value struct {
			Type     string `json:"type"`
			Subtype  string `json:"subtype"`
			ToolCall struct {
				ShellToolCall struct {
					Args struct {
						Command string `json:"command"`
					} `json:"args"`
					Result struct {
						Failure struct {
							Stderr string `json:"stderr"`
						} `json:"failure"`
						Rejected struct {
							Reason string `json:"reason"`
						} `json:"rejected"`
					} `json:"result"`
				} `json:"shellToolCall"`
			} `json:"tool_call"`
		}
		if json.Unmarshal(event, &value) != nil || value.Type != "tool_call" || value.Subtype != "completed" || value.ToolCall.ShellToolCall.Args.Command != command {
			continue
		}
		if stderr := strings.TrimSpace(value.ToolCall.ShellToolCall.Result.Failure.Stderr); stderr != "" {
			return stderr
		}
		if reason := strings.TrimSpace(value.ToolCall.ShellToolCall.Result.Rejected.Reason); reason != "" {
			return reason
		}
		return "the command completed without recording its marker"
	}
	return "the command did not complete"
}

func isSignalCommand(command, cliPath, token, stateRoot string) bool {
	return strings.Contains(command, cliPath) && strings.Contains(command, " signal ") && strings.Contains(command, token) && strings.Contains(command, stateRoot)
}

func (service Service) Cleanup(runID string) error {
	run, err := service.runs.Load(runID)
	if err != nil {
		return err
	}
	if run.InstallationState == "" {
		if err := service.cleanupNativeSession(run); err != nil {
			return err
		}
		return service.finishCleanup(run)
	}
	data, err := os.ReadFile(run.InstallationState)
	if err != nil {
		return fmt.Errorf("read evaluation installation state: %w", err)
	}
	installationState, err := installer.DecodeEvaluationInstallation(data)
	if err != nil {
		return fmt.Errorf("decode evaluation installation state: %w", err)
	}
	if err := service.cleanupNativeSession(run); err != nil {
		return err
	}
	return service.cleanupWithInstallation(runID, installationState)
}

func (service Service) cleanupNativeSession(run runstate.Run) error {
	if run.HarnessSession == "" {
		return nil
	}
	switch run.Harness {
	case string(harness.KiloCode):
		environment, err := kiloEnvironment(privateRuntimeRunRoot(run.ID), run.HarnessExecutable)
		if err != nil {
			return err
		}
		if err := replay.DeleteKiloSession(context.Background(), run.HarnessExecutable, run.Workspace, environment, true, run.HarnessSession); err != nil {
			return fmt.Errorf("delete recovered Kilo session: %w", err)
		}
		return nil
	case string(harness.OpenCode):
		environment, err := openCodeEnvironment(privateRuntimeRunRoot(run.ID), run.HarnessExecutable)
		if err != nil {
			return err
		}
		if err := replay.DeleteOpenCodeSession(context.Background(), run.HarnessExecutable, run.Workspace, environment, true, run.HarnessSession); err != nil {
			return fmt.Errorf("delete recovered OpenCode session: %w", err)
		}
	}
	return nil
}

func (service Service) cleanupWithInstallation(runID string, installationState installer.EvaluationInstallation) error {
	run, err := service.runs.Load(runID)
	if err != nil {
		return err
	}
	harnessID, err := harness.ParseID(run.Harness)
	if err != nil {
		return err
	}
	if err := service.installer.CleanupEvaluation(installer.Request{
		Harness:        harnessID,
		Scope:          harness.ScopeProject,
		Workspace:      run.Workspace,
		EvaluationRoot: installationState.Root,
	}, installationState); err != nil {
		return err
	}
	if run.InstallationState != "" {
		run.InstallationState = ""
		if err := service.runs.Save(run); err != nil {
			return err
		}
	}
	return service.finishCleanup(run)
}

func (service Service) finishCleanup(run runstate.Run) error {
	if err := os.RemoveAll(privateRuntimeRunRoot(run.ID)); err != nil {
		return fmt.Errorf("remove private harness runtime: %w", err)
	}
	if err := service.runs.DeletePrivateMappings(run.ID); err != nil {
		return err
	}
	run, err := service.runs.Load(run.ID)
	if err != nil {
		return err
	}
	run.ActiveTurn = ""
	if run.Status == "complete" || run.Status == "complete-cleaned" {
		run.Status = "complete-cleaned"
	} else if run.Status != "cleaned" {
		run.Status = "cleaned"
	}
	if err := service.runs.Save(run); err != nil {
		return err
	}
	return service.runs.DeleteRun(run.ID)
}

func (service Service) Mark(token string) error {
	return service.runs.Mark(token)
}

func (service Service) setStatus(runID, status string) {
	run, err := service.runs.Load(runID)
	if err != nil {
		return
	}
	run.Status = status
	service.runs.Save(run)
}

func readScenario(path string) (replay.Scenario, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return replay.Scenario{}, fmt.Errorf("read scenario: %w", err)
	}
	var scenario replay.Scenario
	if err := json.Unmarshal(data, &scenario); err != nil {
		return replay.Scenario{}, fmt.Errorf("decode scenario: %w", err)
	}
	if err := scenario.Validate(); err != nil {
		return replay.Scenario{}, err
	}
	return scenario, nil
}

func loadInputs(request RunRequest) (replay.Scenario, AnswerSheet, error) {
	inputs, err := loadEvaluationInputs(request)
	if err != nil {
		return replay.Scenario{}, AnswerSheet{}, err
	}
	return inputs.scenario, inputs.answer, nil
}

func loadEvaluationInputs(request RunRequest) (loadedInputs, error) {
	if request.EvaluationID != "" {
		if request.SkillsPath != "" || request.ScenarioPath != "" || request.AnswerSheet != "" {
			return loadedInputs{}, errors.New("built-in evaluation cannot be combined with custom skill or input files")
		}
		data, err := payload.BuiltInEvaluation(request.EvaluationID)
		if err != nil {
			return loadedInputs{}, err
		}
		var builtIn BuiltInEvaluation
		if err := json.Unmarshal(data, &builtIn); err != nil {
			return loadedInputs{}, fmt.Errorf("decode built-in evaluation: %w", err)
		}
		if builtIn.SchemaVersion != 1 || builtIn.EvaluationID != request.EvaluationID {
			return loadedInputs{}, errors.New("built-in evaluation identity is invalid")
		}
		if err := builtIn.Scenario.Validate(); err != nil {
			return loadedInputs{}, err
		}
		if err := validateAnswerSheet(builtIn.AnswerSheet, builtIn.Scenario); err != nil {
			return loadedInputs{}, err
		}
		skills, err := payload.Skills()
		if err != nil {
			return loadedInputs{}, err
		}
		return loadedInputs{scenario: builtIn.Scenario, answer: builtIn.AnswerSheet, skills: skills}, nil
	}
	if request.SkillsPath == "" || request.ScenarioPath == "" || request.AnswerSheet == "" {
		return loadedInputs{}, errors.New("custom evaluation requires skills, scenario, and answer-sheet inputs")
	}
	inside, err := pathInsideWorkspace(request.Workspace, request.AnswerSheet)
	if err != nil {
		return loadedInputs{}, err
	}
	if inside {
		return loadedInputs{}, errors.New("custom answer sheet must remain outside the evaluated workspace")
	}
	scenario, err := readScenario(request.ScenarioPath)
	if err != nil {
		return loadedInputs{}, err
	}
	answer, err := readAnswerSheet(request.AnswerSheet)
	if err != nil {
		return loadedInputs{}, err
	}
	skills, err := payload.LoadSkills(request.SkillsPath)
	if err != nil {
		return loadedInputs{}, err
	}
	if err := validateAnswerSheetForSkills(answer, scenario, skills); err != nil {
		return loadedInputs{}, err
	}
	return loadedInputs{scenario: scenario, answer: answer, skills: skills}, nil
}

func limitEvaluationInputs(inputs loadedInputs, turns int) loadedInputs {
	if turns <= 0 || turns >= len(inputs.scenario.Turns) {
		return inputs
	}
	inputs.scenario.Turns = inputs.scenario.Turns[:turns]
	includedTurns := make(map[string]bool, turns)
	for _, turn := range inputs.scenario.Turns {
		includedTurns[turn.ID] = true
	}
	expected := make([]SkillCall, 0, len(inputs.answer.Expected))
	for _, call := range inputs.answer.Expected {
		if includedTurns[call.TurnID] {
			expected = append(expected, call)
		}
	}
	inputs.answer.Expected = expected
	return inputs
}

func pathInsideWorkspace(workspace, candidate string) (bool, error) {
	if workspace == "" {
		return false, nil
	}
	resolvedWorkspace, err := filepath.EvalSymlinks(workspace)
	if err != nil {
		return false, fmt.Errorf("resolve evaluation workspace: %w", err)
	}
	resolvedCandidate, err := filepath.EvalSymlinks(candidate)
	if err != nil {
		return false, fmt.Errorf("resolve custom answer sheet: %w", err)
	}
	relative, err := filepath.Rel(resolvedWorkspace, resolvedCandidate)
	if err != nil {
		return false, fmt.Errorf("compare custom answer sheet location: %w", err)
	}
	return relative == "." || (relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator))), nil
}

func prepareOutputRoot(workspace, configured string) (string, error) {
	if strings.TrimSpace(configured) == "" {
		return "", errors.New("evaluation output directory is required")
	}
	outputRoot, err := filepath.Abs(configured)
	if err != nil {
		return "", fmt.Errorf("resolve evaluation output directory: %w", err)
	}
	inside, err := pathWithin(workspace, outputRoot)
	if err != nil {
		return "", err
	}
	if inside {
		return "", errors.New("evaluation output directory must remain outside the evaluated workspace")
	}
	if err := os.MkdirAll(outputRoot, 0o700); err != nil {
		return "", fmt.Errorf("create evaluation output root: %w", err)
	}
	info, err := os.Stat(outputRoot)
	if err != nil {
		return "", fmt.Errorf("inspect evaluation output root: %w", err)
	}
	if !info.IsDir() {
		return "", errors.New("evaluation output root must be a directory")
	}
	inside, err = pathInsideWorkspace(workspace, outputRoot)
	if err != nil {
		return "", err
	}
	if inside {
		return "", errors.New("evaluation output directory must remain outside the evaluated workspace")
	}
	return outputRoot, nil
}

func pathWithin(root, candidate string) (bool, error) {
	relative, err := filepath.Rel(root, candidate)
	if err != nil {
		return false, fmt.Errorf("compare evaluation paths: %w", err)
	}
	return relative == "." || (relative != ".." && !strings.HasPrefix(relative, ".."+string(filepath.Separator))), nil
}

func readAnswerSheet(path string) (AnswerSheet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return AnswerSheet{}, fmt.Errorf("read answer sheet: %w", err)
	}
	var answer AnswerSheet
	if err := json.Unmarshal(data, &answer); err != nil {
		return AnswerSheet{}, fmt.Errorf("decode answer sheet: %w", err)
	}
	return answer, nil
}

func validateAnswerSheet(answer AnswerSheet, scenario replay.Scenario) error {
	skills, err := payload.Skills()
	if err != nil {
		return err
	}
	return validateAnswerSheetForSkills(answer, scenario, skills)
}

func validateAnswerSheetForSkills(answer AnswerSheet, scenario replay.Scenario, skills []payload.Skill) error {
	if answer.SchemaVersion != 1 || answer.ScenarioID != scenario.ID || len(answer.Expected) == 0 {
		return errors.New("answer sheet does not match the scenario")
	}
	turns := map[string]bool{}
	for _, turn := range scenario.Turns {
		turns[turn.ID] = true
	}
	available := map[string]bool{}
	for _, skill := range skills {
		available[skill.Name] = true
	}
	for _, expected := range answer.Expected {
		if !turns[expected.TurnID] || !available[expected.Skill] {
			return fmt.Errorf("invalid expected call for %s on turn %s", expected.Skill, expected.TurnID)
		}
	}
	return nil
}

func skillTokens(skills []payload.Skill) (map[string]string, error) {
	tokens := make(map[string]string, len(skills))
	for _, skill := range skills {
		token, err := runstate.NewToken()
		if err != nil {
			return nil, err
		}
		tokens[token] = skill.Name
	}
	return tokens, nil
}

func deriveResult(runID string, request RunRequest, evaluationID, scenarioID string, startedAt time.Time, expected []SkillCall, events []runstate.Event) Result {
	expectedSet := map[string]bool{}
	for _, item := range expected {
		expectedSet[item.TurnID+"\x00"+item.Skill] = true
	}
	observedSet := map[string]bool{}
	var observed []SkillCall
	var additional []SkillCall
	var unattributed []SkillCall
	for _, event := range events {
		call := SkillCall{TurnID: event.TurnID, Skill: event.Skill}
		if !event.Attributed {
			unattributed = append(unattributed, call)
			continue
		}
		observed = append(observed, call)
		key := event.TurnID + "\x00" + event.Skill
		observedSet[key] = true
		if !expectedSet[key] {
			additional = append(additional, call)
		}
	}
	var missing []SkillCall
	for _, item := range expected {
		if !observedSet[item.TurnID+"\x00"+item.Skill] {
			missing = append(missing, item)
		}
	}
	return Result{
		SchemaVersion: 1,
		RunID:         runID,
		Harness:       string(request.Harness),
		Model:         request.Model,
		Reasoning:     request.Reasoning,
		EvaluationID:  evaluationID,
		ScenarioID:    scenarioID,
		Scope:         string(harness.ScopeProject),
		StartedAt:     startedAt,
		CompletedAt:   time.Now().UTC(),
		Expected:      expected,
		Observed:      observed,
		Missing:       missing,
		Additional:    additional,
		Unattributed:  unattributed,
	}
}

func evaluationIdentity(request RunRequest, scenario replay.Scenario) string {
	if request.EvaluationID != "" {
		return request.EvaluationID
	}
	return scenario.ID
}

func deriveWebsiteResult(result Result, scenario replay.Scenario) WebsiteResult {
	expectedByTurn := make(map[string]map[string]struct{})
	for _, call := range result.Expected {
		skills := expectedByTurn[call.TurnID]
		if skills == nil {
			skills = make(map[string]struct{})
			expectedByTurn[call.TurnID] = skills
		}
		skills[call.Skill] = struct{}{}
	}
	observed := make(map[string]struct{})
	for _, call := range result.Observed {
		observed[call.TurnID+"\x00"+call.Skill] = struct{}{}
	}
	points := make([]WebsitePoint, 0, len(expectedByTurn))
	for index, turn := range scenario.Turns {
		skills := expectedByTurn[turn.ID]
		if len(skills) == 0 {
			continue
		}
		called := 0
		for skill := range skills {
			if _, ok := observed[turn.ID+"\x00"+skill]; ok {
				called++
			}
		}
		points = append(points, WebsitePoint{
			Turn:   index + 1,
			TurnID: turn.ID,
			Called: called,
			Missed: len(skills) - called,
		})
	}
	return WebsiteResult{
		SchemaVersion: 1,
		RunID:         result.RunID,
		ScenarioID:    result.ScenarioID,
		Harness:       result.Harness,
		Model:         result.Model,
		TotalTurns:    len(scenario.Turns),
		Points:        points,
	}
}

func executablePath(configured string) (string, error) {
	path := configured
	if path == "" {
		current, err := os.Executable()
		if err != nil {
			return "", fmt.Errorf("resolve CLI executable: %w", err)
		}
		path = current
	}
	absolute, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("resolve CLI path: %w", err)
	}
	info, err := os.Stat(absolute)
	if err != nil {
		return "", fmt.Errorf("inspect CLI executable: %w", err)
	}
	if info.IsDir() || info.Mode()&0o111 == 0 {
		return "", errors.New("CLI path is not executable")
	}
	return absolute, nil
}

func writeJSON(path string, value any) error {
	data, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("encode evaluation result: %w", err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0o600); err != nil {
		return fmt.Errorf("write evaluation result: %w", err)
	}
	return nil
}

func writeEventsJSONL(path string, events []runstate.Event) error {
	data := make([]byte, 0)
	for _, event := range events {
		encoded, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("encode evaluation event: %w", err)
		}
		data = append(data, encoded...)
		data = append(data, '\n')
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		return fmt.Errorf("write evaluation events: %w", err)
	}
	return nil
}
