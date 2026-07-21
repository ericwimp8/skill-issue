package evaluation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	goruntime "runtime"
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
	Workspace          string
	OutputRoot         string
	Harness            harness.ID
	Model              string
	ModelOverride      bool
	Reasoning          string
	ReasoningOverride  bool
	EvaluationID       string
	SkillsPath         string
	ScenarioPath       string
	AnswerSheet        string
	Executable         string
	CLIPath            string
	IncludeEvents      bool
	IncludeTranscript  bool
	ReplacePreexisting bool
	// WorkspaceCreated marks a workspace directory the CLI itself created
	// while preparing the request, so the confirmation summary can say so and
	// cancellation can remove the empty directory again.
	WorkspaceCreated bool
	// AssumeYes accepts the pre-run confirmation for scripted callers; the
	// summary is still printed for the record.
	AssumeYes          bool
	TurnLimit          int
	AvailableTurns     int
	EffectiveTurns     int
	Progress           func(TurnProgress)
	ConfirmPreexisting func(differing []string) (bool, error)
	// inputs caches the parsed scenario, answer sheet, and skills from
	// PrepareRequest so the run executes exactly what the user confirmed,
	// even if the files change on disk in between.
	inputs *loadedInputs
}

type TurnProgress struct {
	TurnID string
	Index  int
	Total  int
	Phase  replay.BoundaryPhase
	// The following fields are populated only on the "after" phase.
	Duration      time.Duration
	HarnessEvents int
	SkillCalls    int
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
	// Evaluation runtimes assume a Unix environment (path layout, process
	// groups, controlled PATH); fail up front instead of deep in a run.
	if goruntime.GOOS != "darwin" && goruntime.GOOS != "linux" {
		return RunRequest{}, fmt.Errorf("evaluation requires macOS or Linux; %s is not supported", goruntime.GOOS)
	}
	request, err := ResolveRequest(request)
	if err != nil {
		return RunRequest{}, err
	}
	if request.TurnLimit < 0 {
		return RunRequest{}, errors.New("--turns must be a positive integer")
	}
	if request.inputs == nil {
		inputs, err := loadEvaluationInputs(request)
		if err != nil {
			return RunRequest{}, err
		}
		request.inputs = &inputs
	}
	request.AvailableTurns = len(request.inputs.scenario.Turns)
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
	Turn       int    `json:"turn"`
	TurnID     string `json:"turn_id"`
	Called     int    `json:"called"`
	Missed     int    `json:"missed"`
	Unexpected int    `json:"unexpected"`
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
	harnessSpec, err := harness.Lookup(request.Harness)
	if err != nil {
		return Result{}, err
	}
	outputRoot, err := prepareOutputRoot(workspace, request.OutputRoot)
	if err != nil {
		return Result{}, err
	}
	// PrepareRequest cached the parsed inputs; the run must use exactly what
	// the user reviewed rather than re-reading files from disk.
	inputs := limitEvaluationInputs(*request.inputs, request.EffectiveTurns)
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
		Status:            runstate.StatusPreparing,
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
	cleanupFailed := false
	defer func() {
		if err == nil {
			return
		}
		if cleanupFailed {
			err = fmt.Errorf("evaluation run %s: %w; run \"skill-issue evaluate cleanup --run %s --output %s\" to restore the workspace", runID, err, runID, request.OutputRoot)
			return
		}
		_ = service.runs.DeletePrivateMappings(runID)
		_ = service.runs.DeleteRun(runID)
		err = fmt.Errorf("evaluation run %s: %w", runID, err)
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
	if err := replay.CheckAuthentication(ctx, replay.HarnessID(request.Harness), request.Executable, request.Model, runtime.environment, harnessSpec.CleanAuthenticationEnvironment); err != nil {
		return Result{}, service.toolingFailure(outputDirectory, runID, request, err)
	}
	installationState, _, err := service.installer.PrepareEvaluation(installer.Request{
		Harness:              request.Harness,
		Scope:                harness.ScopeProject,
		Workspace:            workspace,
		EvaluationRoot:       runtime.evaluationSkillRoot,
		CLIPath:              cliPath,
		SignalStateRoot:      service.stateRoot,
		BackupRoot:           filepath.Join(service.runs.RunDir(runID), "preexisting-skills"),
		Tokens:               tokens,
		Skills:               inputs.skills,
		CaptureSignals:       request.Harness == harness.Codex,
		ApplyHarnessMetadata: request.EvaluationID != "",
		ConfirmReplace:       confirmPreexistingReplacement(request),
	})
	if err != nil {
		return Result{}, err
	}
	cleaned := false
	defer func() {
		if cleaned {
			return
		}
		if cleanupErr := service.cleanupWithInstallation(runID, installationState); cleanupErr != nil {
			cleanupFailed = true
			if err == nil {
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
			return Result{}, service.toolingFailure(outputDirectory, runID, request, err)
		}
	}
	if request.Harness == harness.KiloCode {
		if err := replay.CheckKiloSkills(ctx, request.Executable, runtime.workingDirectory, runtime.environment, true, skillNames); err != nil {
			return Result{}, service.toolingFailure(outputDirectory, runID, request, err)
		}
	}
	run.InstallationState = installationStatePath
	run.Status = runstate.StatusRunning
	if err := service.runs.Save(run); err != nil {
		return Result{}, err
	}
	adapter, err := service.adapterFactory(replay.HarnessID(request.Harness), replay.Options{
		Executable:            request.Executable,
		Directory:             runtime.workingDirectory,
		Environment:           runtime.environment,
		CleanEnvironment:      harnessSpec.CleanEvaluationEnvironment,
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
		ExpectedSkills:        skillNames,
	})
	if err != nil {
		return Result{}, service.toolingFailure(outputDirectory, runID, request, err)
	}
	var turnStartedAt time.Time
	runner := replay.Runner{
		Adapter: adapter,
		OnBoundary: func(_ context.Context, boundary replay.Boundary) error {
			if boundary.Phase == replay.BoundaryBefore {
				turnStartedAt = time.Now()
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
				if err := service.recordCodexSignals(runID, boundary.TurnID, *boundary.Capture, tokens); err != nil {
					return err
				}
			}
			if request.Harness == harness.Cursor && boundary.Capture != nil {
				if err := service.validateCursorSignals(runID, boundary.TurnID, *boundary.Capture, tokens, cliPath); err != nil {
					return err
				}
			}
			if (request.Harness == harness.OpenCode || request.Harness == harness.KiloCode) && boundary.Capture != nil {
				if err := service.validateStructuredSignals(string(request.Harness), runID, boundary.TurnID, *boundary.Capture, tokens, cliPath); err != nil {
					return err
				}
			}
			if err := service.runs.SetActiveTurn(runID, ""); err != nil {
				return err
			}
			if request.Progress != nil {
				progress := TurnProgress{TurnID: boundary.TurnID, Index: boundary.TurnIndex, Total: boundary.TurnTotal, Phase: boundary.Phase, Duration: time.Since(turnStartedAt)}
				if boundary.Capture != nil {
					progress.HarnessEvents = len(boundary.Capture.Events)
				}
				recorded, err := service.runs.Events(runID)
				if err != nil {
					return err
				}
				for _, event := range recorded {
					if event.TurnID == boundary.TurnID {
						progress.SkillCalls++
					}
				}
				request.Progress(progress)
			}
			return nil
		},
	}
	replayResult, err := runner.Run(ctx, scenario)
	if err != nil {
		return Result{}, service.toolingFailure(outputDirectory, runID, request, err)
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
	run.Status = runstate.StatusComplete
	if err := service.runs.Save(run); err != nil {
		return Result{}, err
	}
	cleaned = true
	if err := service.cleanupWithInstallation(runID, installationState); err != nil {
		cleanupFailed = true
		return Result{}, err
	}
	return result, nil
}

func confirmPreexistingReplacement(request RunRequest) func([]string) (bool, error) {
	return func(differing []string) (bool, error) {
		if request.ReplacePreexisting {
			return true, nil
		}
		if request.ConfirmPreexisting != nil {
			return request.ConfirmPreexisting(differing)
		}
		return false, fmt.Errorf("installed skills differ from their canonical versions: %s; rerun with --replace-preexisting-skills to temporarily replace them (local versions are restored after the run)", strings.Join(differing, ", "))
	}
}

func (service Service) recordCodexSignals(runID, turnID string, capture replay.Capture, tokens map[string]string) error {
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
			if recorded[token] || observed[skill] || !isCodexSignalCommand(value.Item.Command, token) {
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

func isCodexSignalCommand(command, token string) bool {
	return containsShellWord(command, "echo") && containsShellWord(command, token)
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

// validateStructuredSignals classifies errored OpenCode and Kilo bash events
// that carry the instrumented signal. A denied compound command whose signal
// the model retried — so the marker was recorded for the turn — is model
// behavior; an attempted signal whose marker was never recorded is a tooling
// failure, because the instrumentation cannot be trusted from that point.
func (service Service) validateStructuredSignals(harnessName, runID, turnID string, capture replay.Capture, tokens map[string]string, cliPath string) error {
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
			Type string `json:"type"`
			Part struct {
				Tool  string `json:"tool"`
				State struct {
					Status string `json:"status"`
					Error  string `json:"error"`
					Input  struct {
						Command string `json:"command"`
					} `json:"input"`
				} `json:"state"`
			} `json:"part"`
		}
		if json.Unmarshal(event, &value) != nil || value.Type != "tool_use" || value.Part.Tool != "bash" || value.Part.State.Status != "error" {
			continue
		}
		command := value.Part.State.Input.Command
		for token, skill := range tokens {
			if observed[skill] || !isSignalCommand(command, cliPath, token, service.stateRoot) {
				continue
			}
			detail := strings.TrimSpace(value.Part.State.Error)
			if detail == "" {
				detail = "the command did not record its marker"
			}
			return fmt.Errorf("%s attempted the instrumented signal for skill %q but no marker was recorded: %s", harnessName, skill, detail)
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

// isSignalCommand reports whether a captured shell command actually invokes
// the instrumented marker. Each component must appear as a standalone shell
// word (optionally quoted) so that commands merely mentioning a token — for
// example a model echoing its instructions — do not count as signals.
func isSignalCommand(command, cliPath, token, stateRoot string) bool {
	return containsShellWord(command, cliPath) &&
		containsShellWord(command, "signal") &&
		containsShellWord(command, token) &&
		containsShellWord(command, stateRoot)
}

func containsShellWord(command, wanted string) bool {
	if wanted == "" {
		return false
	}
	for start := 0; ; {
		index := strings.Index(command[start:], wanted)
		if index < 0 {
			return false
		}
		index += start
		end := index + len(wanted)
		if shellWordBoundary(command, index-1) && shellWordBoundary(command, end) {
			return true
		}
		start = index + 1
	}
}

// shellWordBoundary reports whether the byte at index (or the string edge)
// can delimit a shell word.
func shellWordBoundary(command string, index int) bool {
	if index < 0 || index >= len(command) {
		return true
	}
	switch command[index] {
	case ' ', '\t', '\n', '"', '\'', ';', '&', '|', '(', ')':
		return true
	default:
		return false
	}
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
	if run.Status == runstate.StatusComplete || run.Status == runstate.StatusCompleteCleaned {
		run.Status = runstate.StatusCompleteCleaned
	} else if run.Status != runstate.StatusCleaned {
		run.Status = runstate.StatusCleaned
	}
	if err := service.runs.Save(run); err != nil {
		return err
	}
	return service.runs.DeleteRun(run.ID)
}

func (service Service) Mark(token string) error {
	return service.runs.Mark(token)
}

func (service Service) setStatus(runID string, status runstate.Status) {
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
	candidate, err := filepath.Abs(candidate)
	if err != nil {
		return false, fmt.Errorf("resolve custom answer sheet: %w", err)
	}
	workspace, err = filepath.Abs(workspace)
	if err != nil {
		return false, fmt.Errorf("resolve evaluation workspace: %w", err)
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

// FailureRecord is the post-mortem artifact written to a failed run's output
// directory. It preserves the failed harness command and its full native
// output, which the error string alone cannot carry.
type FailureRecord struct {
	SchemaVersion int       `json:"schema_version"`
	RunID         string    `json:"run_id"`
	Harness       string    `json:"harness"`
	Model         string    `json:"model"`
	Reasoning     string    `json:"reasoning"`
	TurnID        string    `json:"turn_id,omitempty"`
	FailedAt      time.Time `json:"failed_at"`
	Error         string    `json:"error"`
	Command       string    `json:"command,omitempty"`
	Stdout        string    `json:"stdout,omitempty"`
	Stderr        string    `json:"stderr,omitempty"`
}

// toolingFailure marks the run failed and persists diagnostics best-effort;
// when the record is written, the returned error names its location.
func (service Service) toolingFailure(outputDirectory, runID string, request RunRequest, failure error) error {
	service.setStatus(runID, runstate.StatusToolingFailed)
	wrapped := fmt.Errorf("evaluation encountered a tooling error: %w", failure)
	record := FailureRecord{
		SchemaVersion: 1,
		RunID:         runID,
		Harness:       string(request.Harness),
		Model:         request.Model,
		Reasoning:     request.Reasoning,
		FailedAt:      time.Now().UTC(),
		Error:         failure.Error(),
	}
	if run, err := service.runs.Load(runID); err == nil {
		record.TurnID = run.ActiveTurn
	}
	var diagnostic *replay.DiagnosticError
	if errors.As(failure, &diagnostic) {
		record.Command = diagnostic.Diagnostic.Command
		record.Stdout = diagnostic.Diagnostic.Stdout
		record.Stderr = diagnostic.Diagnostic.Stderr
	}
	failurePath := filepath.Join(outputDirectory, "failure.json")
	if err := writeJSON(failurePath, record); err != nil {
		return wrapped
	}
	return fmt.Errorf("%w; failure diagnostics: %s", wrapped, failurePath)
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
	unexpectedByTurn := make(map[string]map[string]struct{})
	for _, call := range result.Additional {
		skills := unexpectedByTurn[call.TurnID]
		if skills == nil {
			skills = make(map[string]struct{})
			unexpectedByTurn[call.TurnID] = skills
		}
		skills[call.Skill] = struct{}{}
	}
	points := make([]WebsitePoint, 0, len(expectedByTurn)+len(unexpectedByTurn))
	for index, turn := range scenario.Turns {
		expectedSkills := expectedByTurn[turn.ID]
		unexpectedSkills := unexpectedByTurn[turn.ID]
		if len(expectedSkills) == 0 && len(unexpectedSkills) == 0 {
			continue
		}
		called := 0
		for skill := range expectedSkills {
			if _, ok := observed[turn.ID+"\x00"+skill]; ok {
				called++
			}
		}
		points = append(points, WebsitePoint{
			Turn:       index + 1,
			TurnID:     turn.ID,
			Called:     called,
			Missed:     len(expectedSkills) - called,
			Unexpected: len(unexpectedSkills),
		})
	}
	return WebsiteResult{
		SchemaVersion: 2,
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
	var data []byte
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
