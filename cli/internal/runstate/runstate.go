package runstate

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type Run struct {
	SchemaVersion     int               `json:"schema_version"`
	ID                string            `json:"id"`
	Workspace         string            `json:"workspace"`
	Harness           string            `json:"harness"`
	Model             string            `json:"model"`
	Reasoning         string            `json:"reasoning"`
	EvaluationID      string            `json:"evaluation_id"`
	Scenario          string            `json:"scenario"`
	Scope             string            `json:"scope"`
	Status            string            `json:"status"`
	ActiveTurn        string            `json:"active_turn,omitempty"`
	HarnessSession    string            `json:"harness_session,omitempty"`
	HarnessExecutable string            `json:"harness_executable,omitempty"`
	Tokens            map[string]string `json:"tokens"`
	InstallationState string            `json:"installation_state,omitempty"`
	EvidencePath      string            `json:"evidence_path,omitempty"`
	TranscriptPath    string            `json:"transcript_path,omitempty"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}

type Event struct {
	SchemaVersion int       `json:"schema_version"`
	Event         string    `json:"event"`
	RunID         string    `json:"run_id"`
	TurnID        string    `json:"turn_id,omitempty"`
	Attributed    bool      `json:"attributed"`
	Harness       string    `json:"harness"`
	Model         string    `json:"model"`
	Reasoning     string    `json:"reasoning"`
	EvaluationID  string    `json:"evaluation_id"`
	Skill         string    `json:"skill"`
	RecordedAt    time.Time `json:"recorded_at"`
}

type Store struct {
	root string
}

func NewStore(root string) Store {
	return Store{root: root}
}

func NewToken() (string, error) {
	data := make([]byte, 32)
	if _, err := rand.Read(data); err != nil {
		return "", fmt.Errorf("generate opaque token: %w", err)
	}
	return hex.EncodeToString(data), nil
}

func NewRunID() (string, error) {
	data := make([]byte, 16)
	if _, err := rand.Read(data); err != nil {
		return "", fmt.Errorf("generate run ID: %w", err)
	}
	return hex.EncodeToString(data), nil
}

func (store Store) Create(run Run) error {
	if run.SchemaVersion != 1 || run.ID == "" || len(run.Tokens) == 0 {
		return errors.New("run state is incomplete")
	}
	if err := os.MkdirAll(store.runDir(run.ID), 0o700); err != nil {
		return fmt.Errorf("create run directory: %w", err)
	}
	if err := os.MkdirAll(store.tokenDir(), 0o700); err != nil {
		return fmt.Errorf("create token directory: %w", err)
	}
	for token := range run.Tokens {
		if err := validateToken(token); err != nil {
			return err
		}
		if err := writeAtomic(store.tokenPath(token), []byte(run.ID+"\n"), 0o600); err != nil {
			return err
		}
	}
	now := time.Now().UTC()
	run.CreatedAt = now
	run.UpdatedAt = now
	return store.Save(run)
}

func (store Store) Save(run Run) error {
	if run.ID == "" || strings.ContainsAny(run.ID, `/\\`) {
		return errors.New("run ID is invalid")
	}
	run.UpdatedAt = time.Now().UTC()
	data, err := json.MarshalIndent(run, "", "  ")
	if err != nil {
		return fmt.Errorf("encode run state: %w", err)
	}
	return writeAtomic(store.runPath(run.ID), append(data, '\n'), 0o600)
}

func (store Store) Load(runID string) (Run, error) {
	data, err := os.ReadFile(store.runPath(runID))
	if err != nil {
		return Run{}, fmt.Errorf("read run state: %w", err)
	}
	var run Run
	if err := json.Unmarshal(data, &run); err != nil {
		return Run{}, fmt.Errorf("decode run state: %w", err)
	}
	if run.SchemaVersion != 1 || run.ID != runID {
		return Run{}, errors.New("stored run state is invalid")
	}
	return run, nil
}

func (store Store) SetActiveTurn(runID, turnID string) error {
	return store.withLock(runID, func() error {
		run, err := store.Load(runID)
		if err != nil {
			return err
		}
		run.ActiveTurn = turnID
		return store.Save(run)
	})
}

func (store Store) SetHarnessSession(runID, sessionID string) error {
	return store.withLock(runID, func() error {
		run, err := store.Load(runID)
		if err != nil {
			return err
		}
		run.HarnessSession = sessionID
		return store.Save(run)
	})
}

func (store Store) Mark(token string) error {
	if err := validateToken(token); err != nil {
		return err
	}
	data, err := os.ReadFile(store.tokenPath(token))
	if err != nil {
		return fmt.Errorf("resolve opaque token: %w", err)
	}
	runID := strings.TrimSpace(string(data))
	return store.withLock(runID, func() error {
		run, err := store.Load(runID)
		if err != nil {
			return err
		}
		skill, ok := run.Tokens[token]
		if !ok {
			return errors.New("opaque token is not owned by its run")
		}
		event := Event{
			SchemaVersion: 1,
			Event:         "skill_invocation",
			RunID:         run.ID,
			TurnID:        run.ActiveTurn,
			Attributed:    run.ActiveTurn != "",
			Harness:       run.Harness,
			Model:         run.Model,
			Reasoning:     run.Reasoning,
			EvaluationID:  run.EvaluationID,
			Skill:         skill,
			RecordedAt:    time.Now().UTC(),
		}
		encoded, err := json.Marshal(event)
		if err != nil {
			return fmt.Errorf("encode invocation event: %w", err)
		}
		file, err := os.OpenFile(store.eventsPath(runID), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0o600)
		if err != nil {
			return fmt.Errorf("open invocation event log: %w", err)
		}
		defer file.Close()
		if _, err := file.Write(append(encoded, '\n')); err != nil {
			return fmt.Errorf("append invocation event: %w", err)
		}
		return file.Sync()
	})
}

func (store Store) Events(runID string) ([]Event, error) {
	data, err := os.ReadFile(store.eventsPath(runID))
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read invocation events: %w", err)
	}
	var events []Event
	for _, line := range strings.Split(strings.TrimSpace(string(data)), "\n") {
		if line == "" {
			continue
		}
		var event Event
		if err := json.Unmarshal([]byte(line), &event); err != nil {
			return nil, fmt.Errorf("decode invocation event: %w", err)
		}
		events = append(events, event)
	}
	return events, nil
}

func (store Store) DeletePrivateMappings(runID string) error {
	run, err := store.Load(runID)
	if err != nil {
		return err
	}
	for token := range run.Tokens {
		if err := os.Remove(store.tokenPath(token)); err != nil && !errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("remove token mapping: %w", err)
		}
	}
	run.Tokens = nil
	return store.Save(run)
}

func (store Store) RunDir(runID string) string {
	return store.runDir(runID)
}

func (store Store) DeleteRun(runID string) error {
	if runID == "" || strings.ContainsAny(runID, `/\\`) {
		return errors.New("run ID is invalid")
	}
	if err := os.RemoveAll(store.runDir(runID)); err != nil {
		return fmt.Errorf("remove private run state: %w", err)
	}
	for _, path := range []string{store.tokenDir(), filepath.Join(store.root, "runs"), store.root} {
		if err := removeEmptyDirectory(path); err != nil {
			return err
		}
	}
	return nil
}

func (store Store) withLock(runID string, action func() error) error {
	lockPath := filepath.Join(store.runDir(runID), ".lock")
	for attempt := 0; attempt < 100; attempt++ {
		file, err := os.OpenFile(lockPath, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o600)
		if err == nil {
			file.Close()
			defer os.Remove(lockPath)
			return action()
		}
		if !errors.Is(err, os.ErrExist) {
			return fmt.Errorf("acquire run lock: %w", err)
		}
		time.Sleep(10 * time.Millisecond)
	}
	return errors.New("timed out waiting for run lock")
}

func (store Store) runDir(runID string) string {
	return filepath.Join(store.root, "runs", runID)
}

func (store Store) runPath(runID string) string {
	return filepath.Join(store.runDir(runID), "run.json")
}

func (store Store) eventsPath(runID string) string {
	return filepath.Join(store.runDir(runID), "events.jsonl")
}

func (store Store) tokenDir() string {
	return filepath.Join(store.root, "tokens")
}

func (store Store) tokenPath(token string) string {
	return filepath.Join(store.tokenDir(), token)
}

func validateToken(token string) error {
	if len(token) != 64 {
		return errors.New("opaque token is invalid")
	}
	_, err := hex.DecodeString(token)
	if err != nil {
		return errors.New("opaque token is invalid")
	}
	return nil
}

func removeEmptyDirectory(path string) error {
	err := os.Remove(path)
	if err == nil || errors.Is(err, os.ErrNotExist) {
		return nil
	}
	entries, readErr := os.ReadDir(path)
	if readErr == nil && len(entries) > 0 {
		return nil
	}
	return fmt.Errorf("remove empty private state directory: %w", err)
}

func writeAtomic(path string, data []byte, mode os.FileMode) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o700); err != nil {
		return fmt.Errorf("create private state directory: %w", err)
	}
	temporary, err := os.CreateTemp(filepath.Dir(path), ".state-*")
	if err != nil {
		return fmt.Errorf("create temporary state: %w", err)
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)
	if err := temporary.Chmod(mode); err != nil {
		temporary.Close()
		return fmt.Errorf("secure temporary state: %w", err)
	}
	if _, err := temporary.Write(data); err != nil {
		temporary.Close()
		return fmt.Errorf("write temporary state: %w", err)
	}
	if err := temporary.Sync(); err != nil {
		temporary.Close()
		return fmt.Errorf("sync temporary state: %w", err)
	}
	if err := temporary.Close(); err != nil {
		return fmt.Errorf("close temporary state: %w", err)
	}
	if err := os.Rename(temporaryPath, path); err != nil {
		return fmt.Errorf("commit private state: %w", err)
	}
	return nil
}
