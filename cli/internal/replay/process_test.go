package replay

import (
	"context"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMergedEnvironmentReplacesInheritedValues(t *testing.T) {
	t.Setenv("CODEX_HOME", "inherited-home")
	environment := mergedEnvironment([]string{"CODEX_HOME=replacement-home", "CODEX_SQLITE_HOME=replacement-home"})

	var homes []string
	for _, entry := range environment {
		if strings.HasPrefix(entry, "CODEX_HOME=") {
			homes = append(homes, entry)
		}
	}
	if len(homes) != 1 || homes[0] != "CODEX_HOME=replacement-home" {
		t.Fatalf("unexpected CODEX_HOME entries: %v", homes)
	}
	if len(environment) < len(os.Environ()) {
		t.Fatal("environment overlay removed unrelated inherited values")
	}
}

func TestCodexCommandsUseBoundedNonInteractivePermissions(t *testing.T) {
	for _, args := range [][]string{codexInitial("prompt"), codexResume("session", "prompt")} {
		joined := strings.Join(args, " ")
		if !strings.Contains(joined, "--skip-git-repo-check") || !strings.Contains(joined, "--ignore-user-config") || !strings.Contains(joined, "--ignore-rules") || !strings.Contains(joined, "--json") {
			t.Fatalf("Codex command lacks isolated structured execution: %v", args)
		}
		if strings.Contains(joined, "danger-full-access") || strings.Contains(joined, "dangerously-bypass") {
			t.Fatalf("Codex command widens permissions: %v", args)
		}
	}
}

func TestOpenCodeCommandsUseQualifiedStructuredSessionRoute(t *testing.T) {
	adapter := &processAdapter{model: "openai/gpt-5.6-sol", reasoning: "medium"}
	initial := openCodeArgs(adapter, "", "first prompt")
	resumed := openCodeArgs(adapter, "session-1", "second prompt")

	for _, args := range [][]string{initial, resumed} {
		joined := strings.Join(args, " ")
		for _, required := range []string{"run", "--pure", "--format json", "--model openai/gpt-5.6-sol", "--variant medium"} {
			if !strings.Contains(joined, required) {
				t.Fatalf("OpenCode command lacks %q: %v", required, args)
			}
		}
	}
	if strings.Contains(strings.Join(initial, " "), "--session") {
		t.Fatalf("initial OpenCode command resumes a session: %v", initial)
	}
	if !strings.Contains(strings.Join(resumed, " "), "--session session-1") {
		t.Fatalf("resumed OpenCode command lacks its session: %v", resumed)
	}
}

func TestOpenCodeProtocolRequiresStableStoppedSession(t *testing.T) {
	events := []json.RawMessage{
		json.RawMessage(`{"type":"step_start","sessionID":"session-1"}`),
		json.RawMessage(`{"type":"step_finish","sessionID":"session-1","part":{"reason":"stop"}}`),
	}
	if err := validateHarnessOutput(HarnessOpenCode, events, "", true); err != nil {
		t.Fatal(err)
	}
	if id := findSessionID(events); id != "session-1" {
		t.Fatalf("OpenCode session ID = %q", id)
	}
	if err := validateSessionID(HarnessOpenCode, "session-1", events); err != nil {
		t.Fatal(err)
	}

	changed := append(events, json.RawMessage(`{"type":"text","sessionID":"session-2"}`))
	if err := validateSessionID(HarnessOpenCode, "session-1", changed); err == nil {
		t.Fatal("OpenCode session ID change was accepted")
	}
	unfinished := []json.RawMessage{json.RawMessage(`{"type":"step_finish","sessionID":"session-1","part":{"reason":"tool-calls"}}`)}
	if err := validateHarnessOutput(HarnessOpenCode, unfinished, "", true); err == nil {
		t.Fatal("unfinished OpenCode turn was accepted")
	}
}

func TestOpenCodeProtocolToleratesErroredToolEventsAtReplayLayer(t *testing.T) {
	// Errored signal-bearing commands are classified at the evaluation layer,
	// which knows whether the marker was recorded; the replay layer must not
	// fail a stopped turn for them.
	events := []json.RawMessage{
		json.RawMessage(`{"type":"tool_use","sessionID":"session-1","part":{"tool":"bash","state":{"status":"error","input":{"command":"/tmp/skill-issue signal token state && mkdir plans"}}}}`),
		json.RawMessage(`{"type":"step_finish","sessionID":"session-1","part":{"reason":"stop"}}`),
	}
	if err := validateHarnessOutput(HarnessOpenCode, events, "", true); err != nil {
		t.Fatalf("errored tool event failed the replay layer: %v", err)
	}
}

func TestDeleteOpenCodeSessionDeletesOnlyAnExistingSession(t *testing.T) {
	directory := t.TempDir()
	executable := filepath.Join(directory, "opencode")
	logPath := filepath.Join(directory, "deleted")
	if err := os.WriteFile(filepath.Join(directory, ".workspace-marker"), nil, 0o600); err != nil {
		t.Fatal(err)
	}
	script := `#!/bin/sh
if [ "$1 $2" = "session list" ]; then
  [ -f .workspace-marker ] || exit 2
  if [ ! -f "$DELETE_LOG" ]; then
    printf '[{"id":"session-1"}]\n'
  else
    printf '[]\n'
  fi
  exit 0
fi
if [ "$1 $2 $3" = "session delete session-1" ]; then
  printf '%s\n' "$3" >> "$DELETE_LOG"
  exit 0
fi
exit 1
`
	if err := os.WriteFile(executable, []byte(script), 0o700); err != nil {
		t.Fatal(err)
	}
	environment := []string{"PATH=/usr/bin:/bin", "DELETE_LOG=" + logPath}
	if err := DeleteOpenCodeSession(context.Background(), executable, directory, environment, true, "session-1"); err != nil {
		t.Fatal(err)
	}
	data, err := os.ReadFile(logPath)
	if err != nil || strings.TrimSpace(string(data)) != "session-1" {
		t.Fatalf("OpenCode session was not deleted: %q %v", data, err)
	}
	if err := DeleteOpenCodeSession(context.Background(), executable, directory, environment, true, "missing"); err != nil {
		t.Fatal(err)
	}
	data, err = os.ReadFile(logPath)
	if err != nil || strings.Count(strings.TrimSpace(string(data)), "\n") != 0 {
		t.Fatalf("missing OpenCode session triggered deletion: %q %v", data, err)
	}
}

func TestCheckOpenCodeSkillsRequiresEveryExpectedSkill(t *testing.T) {
	directory := t.TempDir()
	executable := filepath.Join(directory, "opencode")
	if err := os.WriteFile(filepath.Join(directory, ".workspace-marker"), nil, 0o600); err != nil {
		t.Fatal(err)
	}
	script := `#!/bin/sh
[ -f .workspace-marker ] || exit 2
[ "$1 $2 $3" = "debug skill --pure" ] || exit 3
printf '[{"name":"alpha"},{"name":"beta"}]\n'
`
	if err := os.WriteFile(executable, []byte(script), 0o700); err != nil {
		t.Fatal(err)
	}
	environment := []string{"PATH=/usr/bin:/bin"}
	if err := CheckOpenCodeSkills(context.Background(), executable, directory, environment, true, []string{"alpha", "beta"}); err != nil {
		t.Fatal(err)
	}
	if err := CheckOpenCodeSkills(context.Background(), executable, directory, environment, true, []string{"missing"}); err == nil {
		t.Fatal("missing OpenCode evaluation skill was accepted")
	}
}

func TestClaudeVisibleSkillValidationRequiresEveryInstalledSkill(t *testing.T) {
	expected := []string{"alpha", "beta"}
	visible := []json.RawMessage{json.RawMessage(`{"type":"system","subtype":"init","session_id":"s","skills":["alpha","beta","operator-skill"]}`)}
	if err := validateClaudeVisibleSkills(visible, expected); err != nil {
		t.Fatal(err)
	}
	hidden := []json.RawMessage{json.RawMessage(`{"type":"system","subtype":"init","session_id":"s","skills":["operator-skill"]}`)}
	if err := validateClaudeVisibleSkills(hidden, expected); err == nil || !strings.Contains(err.Error(), "alpha") {
		t.Fatalf("invisible skill was accepted: %v", err)
	}
	unverifiable := []json.RawMessage{json.RawMessage(`{"type":"system","subtype":"init","session_id":"s"}`)}
	if err := validateClaudeVisibleSkills(unverifiable, expected); err == nil {
		t.Fatal("init event without skill evidence was accepted")
	}
	if err := validateClaudeVisibleSkills([]json.RawMessage{json.RawMessage(`{"type":"result"}`)}, expected); err == nil {
		t.Fatal("output without an init event was accepted")
	}
}

func TestWaitFailureCarriesDiagnostics(t *testing.T) {
	directory := t.TempDir()
	executable := filepath.Join(directory, "fake-harness")
	script := "#!/bin/sh\necho harness-stdout-line\necho harness-stderr-line >&2\nexit 1\n"
	if err := os.WriteFile(executable, []byte(script), 0o700); err != nil {
		t.Fatal(err)
	}
	adapter := &processAdapter{
		harnessID: HarnessClaude,
		path:      executable,
		directory: directory,
		spec: commandSpec{buildArgs: func(_ *processAdapter, _ string, _ bool, prompt string) []string {
			return []string{"-p", prompt}
		}},
	}
	session := &processSession{adapter: adapter}
	if err := session.SendPrompt(context.Background(), "prompt"); err != nil {
		t.Fatal(err)
	}
	_, err := session.Wait(context.Background())
	if err == nil {
		t.Fatal("failing harness was accepted")
	}
	var diagnostic *DiagnosticError
	if !errors.As(err, &diagnostic) {
		t.Fatalf("failure carries no diagnostics: %v", err)
	}
	if diagnostic.Diagnostic.Command != executable+" -p prompt" {
		t.Fatalf("unexpected failed command: %q", diagnostic.Diagnostic.Command)
	}
	if !strings.Contains(diagnostic.Diagnostic.Stdout, "harness-stdout-line") || !strings.Contains(diagnostic.Diagnostic.Stderr, "harness-stderr-line") {
		t.Fatalf("diagnostics lack native output: %#v", diagnostic.Diagnostic)
	}
}
