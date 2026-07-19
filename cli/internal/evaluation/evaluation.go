package evaluation

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
	Workspace      string
	Harness        harness.ID
	Model          string
	ScenarioPath   string
	AnswerSheet    string
	Scope          harness.Scope
	Executable     string
	CLIPath        string
	ProductVersion string
}

type Result struct {
	SchemaVersion  int         `json:"schema_version"`
	RunID          string      `json:"run_id"`
	Harness        string      `json:"harness"`
	Model          string      `json:"model"`
	ScenarioID     string      `json:"scenario_id"`
	Scope          string      `json:"scope"`
	StartedAt      time.Time   `json:"started_at"`
	CompletedAt    time.Time   `json:"completed_at"`
	Expected       []SkillCall `json:"expected"`
	Observed       []SkillCall `json:"observed"`
	Missing        []SkillCall `json:"missing"`
	Additional     []SkillCall `json:"additional"`
	Unattributed   []SkillCall `json:"unattributed"`
	TranscriptPath string      `json:"transcript_path"`
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
		installer:      installer.New(stateRoot),
		adapterFactory: replay.NewAdapter,
	}
}

func (service Service) Run(ctx context.Context, request RunRequest) (result Result, err error) {
	workspace, err := filepath.Abs(request.Workspace)
	if err != nil {
		return Result{}, fmt.Errorf("resolve evaluation workspace: %w", err)
	}
	if info, statErr := os.Stat(workspace); statErr != nil || !info.IsDir() {
		return Result{}, errors.New("evaluation workspace must be an existing directory")
	}
	scenario, err := readScenario(request.ScenarioPath)
	if err != nil {
		return Result{}, err
	}
	answer, err := readAnswerSheet(request.AnswerSheet, scenario)
	if err != nil {
		return Result{}, err
	}
	cliPath, err := executablePath(request.CLIPath)
	if err != nil {
		return Result{}, err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return Result{}, fmt.Errorf("resolve home directory: %w", err)
	}
	tokens, err := skillTokens(answer)
	if err != nil {
		return Result{}, err
	}
	runID, err := runstate.NewRunID()
	if err != nil {
		return Result{}, err
	}
	startedAt := time.Now().UTC()
	run := runstate.Run{
		SchemaVersion: 1,
		ID:            runID,
		Workspace:     workspace,
		Harness:       string(request.Harness),
		Model:         request.Model,
		Scenario:      scenario.ID,
		Scope:         string(request.Scope),
		Status:        "preparing",
		Tokens:        tokens,
	}
	if err := service.runs.Create(run); err != nil {
		return Result{}, err
	}
	defer func() {
		if err != nil {
			_ = service.runs.DeletePrivateMappings(runID)
			err = fmt.Errorf("evaluation run %s: %w", runID, err)
		}
	}()
	backup, _, err := service.installer.PrepareEvaluation(installer.Request{
		Harness:        request.Harness,
		Scope:          request.Scope,
		Workspace:      workspace,
		Home:           home,
		ProductVersion: request.ProductVersion,
		RunID:          runID,
		CLIPath:        cliPath,
		Tokens:         tokens,
	}, service.runs.RunDir(runID))
	if err != nil {
		return Result{}, err
	}
	cleaned := false
	defer func() {
		if !cleaned {
			cleanupErr := service.cleanupWithBackup(runID, backup)
			if err == nil && cleanupErr != nil {
				err = cleanupErr
			}
		}
	}()
	backupPath := filepath.Join(service.runs.RunDir(runID), "installation-backup.json")
	backupData, err := installer.EncodeBackup(backup)
	if err != nil {
		return Result{}, err
	}
	if err := os.WriteFile(backupPath, append(backupData, '\n'), 0o600); err != nil {
		return Result{}, fmt.Errorf("write evaluation restoration state: %w", err)
	}
	run.PriorReceipt = backupPath
	run.Status = "running"
	if err := service.runs.Save(run); err != nil {
		return Result{}, err
	}
	adapter, err := service.adapterFactory(replay.HarnessID(request.Harness), replay.Options{
		Executable: request.Executable,
		Directory:  workspace,
		Model:      request.Model,
	})
	if err != nil {
		service.setStatus(runID, "tooling-failed")
		return Result{}, err
	}
	runner := replay.Runner{
		Adapter: adapter,
		OnBoundary: func(_ context.Context, boundary replay.Boundary) error {
			if boundary.Phase == replay.BoundaryBefore {
				return service.runs.SetActiveTurn(runID, boundary.TurnID)
			}
			if boundary.Capture != nil && boundary.Capture.SessionID != "" {
				if err := service.runs.SetHarnessSession(runID, boundary.Capture.SessionID); err != nil {
					return err
				}
			}
			return service.runs.SetActiveTurn(runID, "")
		},
	}
	replayResult, err := runner.Run(ctx, scenario)
	if err != nil {
		service.setStatus(runID, "tooling-failed")
		return Result{}, err
	}
	events, err := service.runs.Events(runID)
	if err != nil {
		return Result{}, err
	}
	result = deriveResult(runID, request, scenario.ID, startedAt, answer.Expected, events)
	transcriptPath := filepath.Join(service.runs.RunDir(runID), "transcript.json")
	if err := writeJSON(transcriptPath, replayResult); err != nil {
		return Result{}, err
	}
	result.TranscriptPath = filepath.Base(transcriptPath)
	evidencePath := filepath.Join(service.runs.RunDir(runID), "result.json")
	if err := writeJSON(evidencePath, result); err != nil {
		return Result{}, err
	}
	run, err = service.runs.Load(runID)
	if err != nil {
		return Result{}, err
	}
	run.EvidencePath = evidencePath
	run.TranscriptPath = transcriptPath
	run.Status = "complete"
	if err := service.runs.Save(run); err != nil {
		return Result{}, err
	}
	if err := service.cleanupWithBackup(runID, backup); err != nil {
		return Result{}, err
	}
	cleaned = true
	return result, nil
}

func (service Service) Cleanup(runID string) error {
	run, err := service.runs.Load(runID)
	if err != nil {
		return err
	}
	if run.PriorReceipt == "" {
		return service.finishCleanup(run)
	}
	data, err := os.ReadFile(run.PriorReceipt)
	if err != nil {
		return fmt.Errorf("read evaluation restoration state: %w", err)
	}
	backup, err := installer.DecodeBackup(data)
	if err != nil {
		return fmt.Errorf("decode evaluation restoration state: %w", err)
	}
	return service.cleanupWithBackup(runID, backup)
}

func (service Service) cleanupWithBackup(runID string, backup installer.EvaluationBackup) error {
	run, err := service.runs.Load(runID)
	if err != nil {
		return err
	}
	if run.PriorReceipt != "" {
		if err := service.installer.CleanupEvaluation(backup, runID); err != nil {
			return err
		}
		run.PriorReceipt = ""
		if err := service.runs.Save(run); err != nil {
			return err
		}
	}
	return service.finishCleanup(run)
}

func (service Service) finishCleanup(run runstate.Run) error {
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
	return service.runs.Save(run)
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

func readAnswerSheet(path string, scenario replay.Scenario) (AnswerSheet, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return AnswerSheet{}, fmt.Errorf("read answer sheet: %w", err)
	}
	var answer AnswerSheet
	if err := json.Unmarshal(data, &answer); err != nil {
		return AnswerSheet{}, fmt.Errorf("decode answer sheet: %w", err)
	}
	if answer.SchemaVersion != 1 || answer.ScenarioID != scenario.ID || len(answer.Expected) == 0 {
		return AnswerSheet{}, errors.New("answer sheet does not match the scenario")
	}
	turns := map[string]bool{}
	for _, turn := range scenario.Turns {
		turns[turn.ID] = true
	}
	available := map[string]bool{}
	skills, err := payload.EvaluationSkills()
	if err != nil {
		return AnswerSheet{}, err
	}
	for _, skill := range skills {
		available[skill.Name] = true
	}
	for _, expected := range answer.Expected {
		if !turns[expected.TurnID] || !available[expected.Skill] {
			return AnswerSheet{}, fmt.Errorf("invalid expected call for %s on turn %s", expected.Skill, expected.TurnID)
		}
	}
	return answer, nil
}

func skillTokens(answer AnswerSheet) (map[string]string, error) {
	skills, err := payload.EvaluationSkills()
	if err != nil {
		return nil, err
	}
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

func deriveResult(runID string, request RunRequest, scenarioID string, startedAt time.Time, expected []SkillCall, events []runstate.Event) Result {
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
		ScenarioID:    scenarioID,
		Scope:         string(request.Scope),
		StartedAt:     startedAt,
		CompletedAt:   time.Now().UTC(),
		Expected:      expected,
		Observed:      observed,
		Missing:       missing,
		Additional:    additional,
		Unattributed:  unattributed,
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
