package evaluation

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ericwimp8/skill-issue/cli/internal/replay"
	"github.com/ericwimp8/skill-issue/cli/internal/runstate"
)

func TestTranscriptSanitizerCleansCompleteReplayResult(t *testing.T) {
	temporaryRoot := t.TempDir()
	workspace := filepath.Join(temporaryRoot, "workspace", "project")
	outputRoot := filepath.Join(temporaryRoot, "output")
	stateRoot := filepath.Join(outputRoot, ".skill-issue")
	runtimeRoot := filepath.Join(temporaryRoot, "runtime", "run")
	cliPath := filepath.Join(workspace, "cli", "skill-issue")
	windowsWorkspace := strings.Join([]string{"C:", "Users", "example", "project"}, `\`)
	sanitizer, err := newTranscriptSanitizer(transcriptSanitizerConfig{
		Workspace:   workspace,
		OutputRoot:  outputRoot,
		StateRoot:   stateRoot,
		RuntimeRoot: runtimeRoot,
		CLIPath:     cliPath,
	})
	if err != nil {
		t.Fatal(err)
	}
	sanitizer.addPath(windowsWorkspace, "[windows-workspace]")
	sanitizer.addIdentity("Developer-Mac", "[host]")

	event, err := json.Marshal(map[string]any{
		workspace + "/event-key": "message",
		"nested": map[string]any{
			"path":   filepath.Join(workspace, "notes.md"),
			"number": json.Number("9007199254740993"),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	result := replay.Result{
		HarnessID: replay.HarnessClaude,
		Scenario: replay.Scenario{Turns: []replay.Turn{{
			ID:     "turn-1",
			Prompt: "Inspect " + filepath.Join(workspace, "prompt.md"),
		}}},
		Turns: []replay.TurnResult{{
			TurnID: "turn-1",
			Capture: replay.Capture{
				SessionID:  "session-1",
				Transcript: `Run "` + cliPath + `" signal "token" "` + stateRoot + `" from ` + runtimeRoot + ` and ` + windowsWorkspace + ` with ` + strings.ReplaceAll(workspace, "/", "\\/") + ` on Developer-Mac.`,
				Stderr:     "workspace: " + workspace,
				Events:     []json.RawMessage{event},
			},
		}},
	}

	if err := sanitizer.sanitize(&result); err != nil {
		t.Fatal(err)
	}
	encoded, err := json.Marshal(result)
	if err != nil {
		t.Fatal(err)
	}
	text := string(encoded)
	for _, private := range []string{workspace, outputRoot, stateRoot, runtimeRoot, cliPath, windowsWorkspace, strings.ReplaceAll(workspace, "/", `\/`), "Developer-Mac"} {
		if strings.Contains(text, private) {
			t.Fatalf("sanitized transcript retained %q: %s", private, text)
		}
	}
	for _, placeholder := range []string{"[workspace]", "[skill-issue-cli]", "[evaluation-state]", "[runtime]", "[windows-workspace]", "[host]"} {
		if !strings.Contains(text, placeholder) {
			t.Fatalf("sanitized transcript is missing %q: %s", placeholder, text)
		}
	}
	if result.Turns[0].Capture.SessionID != "session-1" || !strings.Contains(text, "9007199254740993") {
		t.Fatalf("sanitization changed non-private capture data: %s", text)
	}
}

func TestTranscriptSanitizerRespectsPathBoundaries(t *testing.T) {
	sanitizer := transcriptSanitizer{}
	workspace := filepath.Join(t.TempDir(), "workspace")
	workspaceFile := filepath.Join(workspace, "file.txt")
	backupFile := filepath.Join(workspace+"-backup", "file.txt")
	sanitizer.addPath(workspace, "[workspace]")

	value := sanitizer.sanitizeText(workspaceFile + " " + backupFile)
	if value != filepath.Join("[workspace]", "file.txt")+" "+backupFile {
		t.Fatalf("unexpected boundary sanitization: %s", value)
	}
}

func TestTranscriptSanitizerRejectsMalformedStructuredEvent(t *testing.T) {
	result := replay.Result{Turns: []replay.TurnResult{{
		TurnID:  "turn-1",
		Capture: replay.Capture{Events: []json.RawMessage{json.RawMessage(`{"broken"`)}},
	}}}

	if err := (transcriptSanitizer{}).sanitize(&result); err == nil {
		t.Fatal("malformed transcript event was accepted")
	}
}

func TestTranscriptSanitizationFollowsCodexSignalRecording(t *testing.T) {
	stateRoot := t.TempDir()
	service := New(stateRoot)
	runID, err := runstate.NewRunID()
	if err != nil {
		t.Fatal(err)
	}
	token, err := runstate.NewToken()
	if err != nil {
		t.Fatal(err)
	}
	if err := service.runs.Create(runstate.Run{
		SchemaVersion: 1,
		ID:            runID,
		ActiveTurn:    "turn-1",
		Tokens:        map[string]string{token: "prompt-writing"},
	}); err != nil {
		t.Fatal(err)
	}
	command := `echo "` + token + `"`
	event, err := json.Marshal(map[string]any{
		"type": "item.completed",
		"item": map[string]any{"type": "command_execution", "command": command},
	})
	if err != nil {
		t.Fatal(err)
	}
	capture := replay.Capture{Events: []json.RawMessage{event}, Transcript: command}
	if err := service.recordCodexSignals(runID, "turn-1", capture, map[string]string{token: "prompt-writing"}); err != nil {
		t.Fatal(err)
	}

	sanitizer := transcriptSanitizer{}
	sanitizer.addPath(stateRoot, "[evaluation-state]")
	replayResult := replay.Result{Turns: []replay.TurnResult{{TurnID: "turn-1", Capture: capture}}}
	if err := sanitizer.sanitize(&replayResult); err != nil {
		t.Fatal(err)
	}
	events, err := service.runs.Events(runID)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 || events[0].Skill != "prompt-writing" || events[0].TurnID != "turn-1" {
		t.Fatalf("unexpected recorded signals: %#v", events)
	}
	if strings.Contains(replayResult.Turns[0].Capture.Transcript, stateRoot) {
		t.Fatal("stored replay capture retained scoring paths")
	}
}

func TestCursorSignalAttemptWithoutMarkerIsToolingFailure(t *testing.T) {
	stateRoot := t.TempDir()
	service := New(stateRoot)
	runID, token := createCursorSignalRun(t, service, "turn-1")
	cliPath := filepath.Join(stateRoot, "bin", "skill-issue")
	capture := cursorSignalCapture(t, cliPath+` signal `+token+` `+stateRoot)

	if err := service.validateCursorSignals(runID, "turn-1", capture, map[string]string{token: "prompt-writing"}, cliPath); err == nil {
		t.Fatal("Cursor signal attempt without a recorded marker was accepted")
	}
}

func TestCursorSignalAttemptWithMarkerIsAccepted(t *testing.T) {
	stateRoot := t.TempDir()
	service := New(stateRoot)
	runID, token := createCursorSignalRun(t, service, "turn-1")
	if err := service.runs.Mark(token); err != nil {
		t.Fatal(err)
	}
	cliPath := filepath.Join(stateRoot, "bin", "skill-issue")
	capture := cursorSignalCapture(t, cliPath+` signal `+token+` `+stateRoot)

	if err := service.validateCursorSignals(runID, "turn-1", capture, map[string]string{token: "prompt-writing"}, cliPath); err != nil {
		t.Fatal(err)
	}
}

func TestCursorWithoutSignalAttemptRemainsModelMiss(t *testing.T) {
	stateRoot := t.TempDir()
	service := New(stateRoot)
	runID, token := createCursorSignalRun(t, service, "turn-1")
	if err := service.validateCursorSignals(runID, "turn-1", replay.Capture{}, map[string]string{token: "prompt-writing"}, filepath.Join(stateRoot, "bin", "skill-issue")); err != nil {
		t.Fatal(err)
	}
}

func createCursorSignalRun(t *testing.T, service Service, turnID string) (string, string) {
	t.Helper()
	runID, err := runstate.NewRunID()
	if err != nil {
		t.Fatal(err)
	}
	token, err := runstate.NewToken()
	if err != nil {
		t.Fatal(err)
	}
	if err := service.runs.Create(runstate.Run{SchemaVersion: 1, ID: runID, ActiveTurn: turnID, Tokens: map[string]string{token: "prompt-writing"}}); err != nil {
		t.Fatal(err)
	}
	return runID, token
}

func cursorSignalCapture(t *testing.T, command string) replay.Capture {
	t.Helper()
	event, err := json.Marshal(map[string]any{
		"type":    "tool_call",
		"subtype": "started",
		"tool_call": map[string]any{
			"shellToolCall": map[string]any{"args": map[string]any{"command": command}},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	return replay.Capture{Events: []json.RawMessage{event}}
}
