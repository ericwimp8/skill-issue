package evaluation

import (
	"encoding/json"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ericwimp8/skill-issue/cli/internal/replay"
	"github.com/ericwimp8/skill-issue/cli/internal/runstate"
)

func TestTranscriptSanitizerCleansConversationArtifact(t *testing.T) {
	temporaryRoot := t.TempDir()
	workspace := filepath.Join(temporaryRoot, "workspace", "project")
	outputRoot := filepath.Join(temporaryRoot, "output")
	stateRoot := filepath.Join(outputRoot, ".skill-issue")
	runtimeRoot := filepath.Join(temporaryRoot, "runtime", "run")
	cliPath := filepath.Join(workspace, "cli", "skill-issue")
	windowsWorkspace := strings.Join([]string{"C:", "Users", "example", "project"}, `\`)
	sanitizer, err := newArtifactSanitizer(artifactSanitizerConfig{
		Workspace: workspace, OutputRoot: outputRoot, StateRoot: stateRoot, RuntimeRoot: runtimeRoot, CLIPath: cliPath,
	})
	if err != nil {
		t.Fatal(err)
	}
	sanitizer.addPath(windowsWorkspace, "[windows-workspace]")
	sanitizer.addIdentity("Developer-Mac", "[host]")
	artifact := TranscriptArtifact{SchemaVersion: 1, Turns: []TranscriptTurn{{
		TurnID: "turn-1",
		User:   "Inspect " + filepath.Join(workspace, "prompt.md"),
		Assistant: `Run "` + cliPath + `" from ` + stateRoot + ` and ` + runtimeRoot + ` and ` + windowsWorkspace +
			` with ` + strings.ReplaceAll(workspace, "/", "\\/") + ` on Developer-Mac.`,
	}}}

	sanitizer.sanitizeTranscript(&artifact)
	encoded, err := json.Marshal(artifact)
	if err != nil {
		t.Fatal(err)
	}
	text := string(encoded)
	for _, private := range []string{workspace, outputRoot, stateRoot, runtimeRoot, cliPath, windowsWorkspace, "Developer-Mac"} {
		if strings.Contains(text, private) {
			t.Fatalf("sanitized transcript retained %q: %s", private, text)
		}
	}
	for _, placeholder := range []string{"[workspace]", "[skill-issue-cli]", "[evaluation-state]", "[runtime]", "[windows-workspace]", "[host]"} {
		if !strings.Contains(text, placeholder) {
			t.Fatalf("sanitized transcript is missing %q: %s", placeholder, text)
		}
	}
}

func TestTranscriptSanitizerRespectsPathBoundaries(t *testing.T) {
	sanitizer := artifactSanitizer{}
	workspace := filepath.Join(t.TempDir(), "workspace")
	workspaceFile := filepath.Join(workspace, "file.txt")
	backupFile := filepath.Join(workspace+"-backup", "file.txt")
	sanitizer.addPath(workspace, "[workspace]")

	value := sanitizer.sanitizeText(workspaceFile + " " + backupFile)
	if value != filepath.Join("[workspace]", "file.txt")+" "+backupFile {
		t.Fatalf("unexpected boundary sanitization: %s", value)
	}
}

func TestArtifactSanitizerCleansSensitivePatterns(t *testing.T) {
	sanitizer := artifactSanitizer{patterns: defaultArtifactPatterns()}
	cases := map[string]struct {
		input    string
		expected string
	}{
		"authorization":   {"authorization: Bearer private-credential-value", "authorization: Bearer [authorization]"},
		"email":           {"private@example.com", "[email]"},
		"ip":              {"203.0.113.42", "[ip-address]"},
		"jwt":             {"eyJabcdefghijk.abcdefghijkl.abcdefghijkl", "[jwt]"},
		"private-key":     {"-----BEGIN PRIVATE KEY-----\nprivate-material\n-----END PRIVATE KEY-----", "[private-key]"},
		"provider-token":  {"ghp_1234567890abcdefghijklmnop", "[token]"},
		"secret":          {`client_secret="private-value"`, "client_secret=[secret]"},
		"unix-home":       {"/home/private-person/project", "/home/[user]/project"},
		"url-credentials": {"https://private-user:private-password@example.com/path", "https://[url-credentials]@example.com/path"},
		"windows-home":    {`C:\Users\private-person\project`, `C:\Users\[user]\project`},
	}
	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			if sanitized := sanitizer.sanitizeText(testCase.input); sanitized != testCase.expected {
				t.Fatalf("sanitized value = %q, want %q", sanitized, testCase.expected)
			}
		})
	}
	if sanitized := sanitizer.sanitizeText("version 999.999.999.999"); sanitized != "version 999.999.999.999" {
		t.Fatalf("invalid IP-like value was changed: %q", sanitized)
	}
}

func TestTranscriptArtifactKeepsOnlyConversationForEveryHarness(t *testing.T) {
	cases := map[replay.HarnessID][]json.RawMessage{
		replay.HarnessCodex:    {json.RawMessage(`{"type":"item.completed","item":{"type":"command_execution","command":"private command"}}`), json.RawMessage(`{"type":"item.completed","item":{"type":"agent_message","text":"assistant reply"}}`)},
		replay.HarnessCursor:   {json.RawMessage(`{"type":"tool_call","subtype":"started","tool_call":{"command":"private command"}}`), json.RawMessage(`{"type":"assistant","message":{"role":"assistant","content":[{"type":"text","text":"assistant reply"}]}}`)},
		replay.HarnessClaude:   {json.RawMessage(`{"type":"assistant","message":{"role":"assistant","content":[{"type":"thinking","thinking":"private reasoning"},{"type":"text","text":"assistant reply"}]}}`)},
		replay.HarnessOpenCode: {json.RawMessage(`{"type":"tool_use","part":{"type":"tool","state":{"input":{"command":"private command"}}}}`), json.RawMessage(`{"type":"text","part":{"type":"text","text":"assistant reply"}}`)},
		replay.HarnessPi:       {json.RawMessage(`{"type":"tool_execution_start","args":{"command":"private command"}}`), json.RawMessage(`{"type":"message_end","message":{"role":"assistant","content":[{"type":"thinking","thinking":"private reasoning"},{"type":"text","text":"assistant reply"}]}}`)},
	}
	for harnessID, events := range cases {
		t.Run(string(harnessID), func(t *testing.T) {
			artifact, err := newTranscriptArtifact(replay.Result{
				HarnessID: harnessID,
				Scenario:  replay.Scenario{Turns: []replay.Turn{{ID: "turn-1", Prompt: "user prompt"}}},
				Turns:     []replay.TurnResult{{TurnID: "turn-1", Capture: replay.Capture{SessionID: "private-session", Transcript: "private transport", Stderr: "private stderr", Events: events}}},
			})
			if err != nil {
				t.Fatal(err)
			}
			encoded, err := json.Marshal(artifact)
			if err != nil {
				t.Fatal(err)
			}
			if len(artifact.Turns) != 1 || artifact.Turns[0].User != "user prompt" || artifact.Turns[0].Assistant != "assistant reply" {
				t.Fatalf("unexpected conversation artifact: %#v", artifact)
			}
			for _, excluded := range []string{"capture", "events", "session", "stderr", "tool", "command", "transport", "reasoning"} {
				if strings.Contains(string(encoded), excluded) {
					t.Fatalf("conversation artifact retained %q: %s", excluded, encoded)
				}
			}
		})
	}
}

func TestTranscriptExtractionFollowsSignalRecording(t *testing.T) {
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
	if err := service.runs.Create(runstate.Run{SchemaVersion: 1, ID: runID, ActiveTurn: "turn-1", Tokens: map[string]string{token: "prompt-writing"}}); err != nil {
		t.Fatal(err)
	}
	event, err := json.Marshal(map[string]any{"type": "item.completed", "item": map[string]any{"type": "command_execution", "command": `sed -n '1,240p' .agents/skills/prompt-writing/SKILL.md`}})
	if err != nil {
		t.Fatal(err)
	}
	capture := replay.Capture{Events: []json.RawMessage{event}}
	if err := service.recordCodexSignals(runID, "turn-1", capture, map[string]string{token: "prompt-writing"}); err != nil {
		t.Fatal(err)
	}
	events, err := service.runs.Events(runID)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 || events[0].Skill != "prompt-writing" || events[0].TurnID != "turn-1" {
		t.Fatalf("unexpected recorded signals: %#v", events)
	}
}

func TestCursorSignalAttemptWithoutMarkerIsToolingFailure(t *testing.T) {
	stateRoot := t.TempDir()
	service := New(stateRoot)
	runID, token := createCursorSignalRun(t, service, "turn-1")
	cliPath := filepath.Join(stateRoot, "bin", "skill-issue")
	capture := cursorSignalCapture(t, cliPath+` signal `+token+` `+stateRoot)
	if err := service.validateCursorSignals(runID, "turn-1", capture, map[string]string{token: "prompt-writing"}, cliPath); err == nil {
		t.Fatal("Cursor attempted signal without a recorded marker was accepted")
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
		"type": "tool_call", "subtype": "started",
		"tool_call": map[string]any{"shellToolCall": map[string]any{"args": map[string]any{"command": command}}},
	})
	if err != nil {
		t.Fatal(err)
	}
	return replay.Capture{Events: []json.RawMessage{event}}
}
