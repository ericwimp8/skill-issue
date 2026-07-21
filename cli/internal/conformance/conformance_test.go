//go:build darwin || linux

package conformance

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
)

var binaries struct {
	cli  string
	fake string
}

func TestMain(m *testing.M) {
	directory, err := os.MkdirTemp("", "skill-issue-conformance-")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	binaries.cli = filepath.Join(directory, "skill-issue")
	binaries.fake = filepath.Join(directory, "fake-harness")
	builds := [][2]string{
		{binaries.cli, "github.com/ericwimp8/skill-issue/cli/cmd/skill-issue"},
		{binaries.fake, "./testdata/fakeharness"},
	}
	for _, build := range builds {
		command := exec.Command("go", "build", "-o", build[0], build[1])
		command.Stderr = os.Stderr
		if err := command.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "build %s: %v\n", build[1], err)
			os.RemoveAll(directory)
			os.Exit(1)
		}
	}
	code := m.Run()
	os.RemoveAll(directory)
	os.Exit(code)
}

var evaluationHarnesses = []string{"claude-code", "codex", "cursor", "opencode", "kilo-code", "pi"}

var harnessExecutables = map[string]string{
	"claude-code": "claude",
	"codex":       "codex",
	"cursor":      "agent",
	"opencode":    "opencode",
	"kilo-code":   "kilo",
	"pi":          "pi",
}

type conformanceRun struct {
	t          *testing.T
	harness    string
	root       string
	workspace  string
	output     string
	tmp        string
	home       string
	fakeDir    string
	executable string
	inputs     string
	extraEnv   []string
}

func newConformanceRun(t *testing.T, harnessID, mode string) *conformanceRun {
	t.Helper()
	run := &conformanceRun{t: t, harness: harnessID, root: t.TempDir()}
	run.workspace = filepath.Join(run.root, "workspace")
	run.output = filepath.Join(run.root, "output")
	run.tmp = filepath.Join(run.root, "tmp")
	run.home = filepath.Join(run.root, "home")
	run.fakeDir = filepath.Join(run.root, "harness-bin")
	run.inputs = filepath.Join(run.root, "inputs")
	for _, directory := range []string{run.workspace, run.output, run.tmp, run.home, run.fakeDir} {
		if err := os.MkdirAll(directory, 0o700); err != nil {
			t.Fatal(err)
		}
	}
	run.executable = filepath.Join(run.fakeDir, harnessExecutables[harnessID])
	if err := os.Symlink(binaries.fake, run.executable); err != nil {
		t.Fatal(err)
	}
	run.setMode(mode, "")
	run.writeInputs()
	if harnessID == "codex" {
		codexHome := filepath.Join(run.root, "codex-home")
		if err := os.MkdirAll(codexHome, 0o700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(filepath.Join(codexHome, "auth.json"), []byte(`{"token":"conformance"}`+"\n"), 0o600); err != nil {
			t.Fatal(err)
		}
		run.extraEnv = append(run.extraEnv, "CODEX_HOME="+codexHome)
	}
	return run
}

func (run *conformanceRun) setMode(mode, version string) {
	run.t.Helper()
	value := map[string]string{"harness": run.harness, "mode": mode}
	if version != "" {
		value["version"] = version
	}
	data, err := json.Marshal(value)
	if err != nil {
		run.t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(run.fakeDir, "fake-mode.json"), append(data, '\n'), 0o600); err != nil {
		run.t.Fatal(err)
	}
}

func (run *conformanceRun) writeInputs() {
	run.t.Helper()
	skill := filepath.Join(run.inputs, "skills", "conformance-skill")
	if err := os.MkdirAll(skill, 0o700); err != nil {
		run.t.Fatal(err)
	}
	entrypoint := "---\nname: conformance-skill\ndescription: Conformance test skill.\n---\n\n# Conformance Skill\n"
	if err := os.WriteFile(filepath.Join(skill, "SKILL.md"), []byte(entrypoint), 0o600); err != nil {
		run.t.Fatal(err)
	}
	run.writeJSON(filepath.Join(run.inputs, "scenario.json"), map[string]any{
		"schema_version": 1,
		"scenario_id":    "conformance",
		"turns": []map[string]string{
			{"turn_id": "turn-1", "prompt": "Complete the first conformance step."},
			{"turn_id": "turn-2", "prompt": "Complete the second conformance step."},
		},
	})
	run.writeJSON(filepath.Join(run.inputs, "answer.json"), map[string]any{
		"schema_version": 1,
		"scenario_id":    "conformance",
		"expected": []map[string]string{
			{"turn_id": "turn-1", "skill": "conformance-skill"},
			{"turn_id": "turn-2", "skill": "conformance-skill"},
		},
	})
}

func (run *conformanceRun) writeJSON(path string, value any) {
	run.t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		run.t.Fatal(err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0o600); err != nil {
		run.t.Fatal(err)
	}
}

// execute drives one full evaluate run through the built CLI binary against
// the fake harness, with a hermetic identity-neutral environment.
func (run *conformanceRun) execute(extraArgs, extraEnv []string) (int, string, string) {
	run.t.Helper()
	args := []string{
		"evaluate", "run",
		"--workspace", run.workspace,
		"--output", run.output,
		"--harness", run.harness,
		"--executable", run.executable,
		"--skills", filepath.Join(run.inputs, "skills"),
		"--scenario", filepath.Join(run.inputs, "scenario.json"),
		"--answer-sheet", filepath.Join(run.inputs, "answer.json"),
		"--events",
	}
	args = append(args, extraArgs...)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()
	command := exec.CommandContext(ctx, binaries.cli, args...)
	command.Dir = run.root
	command.Stdin = strings.NewReader("y\n")
	command.Env = append([]string{
		"HOME=" + run.home,
		"TMPDIR=" + run.tmp,
		"PATH=/usr/bin:/bin",
		"USER=conformance",
		"LOGNAME=conformance",
		"TERM=dumb",
		"LANG=en_US.UTF-8",
	}, append(append([]string{}, run.extraEnv...), extraEnv...)...)
	var stdout, stderr bytes.Buffer
	command.Stdout = &stdout
	command.Stderr = &stderr
	err := command.Run()
	if ctx.Err() != nil {
		run.t.Fatalf("evaluate run timed out: %s", stderr.String())
	}
	code := 0
	if err != nil {
		exit, ok := err.(*exec.ExitError)
		if !ok {
			run.t.Fatalf("run CLI: %v: %s", err, stderr.String())
		}
		code = exit.ExitCode()
	}
	return code, stdout.String(), stderr.String()
}

func (run *conformanceRun) completedResult(stdout string) evaluation.Result {
	run.t.Helper()
	var output struct {
		Action string            `json:"action"`
		Status string            `json:"status"`
		Data   evaluation.Result `json:"data"`
	}
	if err := json.Unmarshal([]byte(stdout), &output); err != nil {
		run.t.Fatalf("decode CLI output: %v: %s", err, stdout)
	}
	if output.Action != "evaluate" || output.Status != "complete" {
		run.t.Fatalf("run did not complete: %#v", output)
	}
	return output.Data
}

func (run *conformanceRun) runDirectory() string {
	run.t.Helper()
	directories, err := filepath.Glob(filepath.Join(run.output, run.harness+"-*"))
	if err != nil || len(directories) != 1 {
		run.t.Fatalf("expected one run output directory: %v %v", directories, err)
	}
	return directories[0]
}

func (run *conformanceRun) failureRecord() evaluation.FailureRecord {
	run.t.Helper()
	data, err := os.ReadFile(filepath.Join(run.runDirectory(), "failure.json"))
	if err != nil {
		run.t.Fatalf("read failure.json: %v", err)
	}
	var record evaluation.FailureRecord
	if err := json.Unmarshal(data, &record); err != nil {
		run.t.Fatal(err)
	}
	return record
}

// assertCleanup verifies the run left no temporary skills in the workspace,
// no private harness runtime under the run's TMPDIR, and no private run state
// under the output root.
func (run *conformanceRun) assertCleanup() {
	run.t.Helper()
	err := filepath.WalkDir(run.workspace, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !entry.IsDir() && entry.Name() == "SKILL.md" {
			run.t.Errorf("temporary skill remains in workspace: %s", path)
		}
		return nil
	})
	if err != nil {
		run.t.Fatal(err)
	}
	if entries, err := os.ReadDir(filepath.Join(run.tmp, "skill-issue")); err == nil && len(entries) > 0 {
		run.t.Errorf("private harness runtime remains: %v", entries)
	}
	states, err := filepath.Glob(filepath.Join(run.output, ".skill-issue", "runs", "*", "run.json"))
	if err != nil || len(states) > 0 {
		run.t.Errorf("private run state remains: %v %v", states, err)
	}
}

func expectedConformanceCalls(result evaluation.Result) error {
	observed := map[string]int{}
	for _, call := range result.Observed {
		observed[call.TurnID+" "+call.Skill]++
	}
	for _, turn := range []string{"turn-1", "turn-2"} {
		if observed[turn+" conformance-skill"] == 0 {
			return fmt.Errorf("no observed conformance-skill call on %s: %#v", turn, result.Observed)
		}
	}
	if len(result.Missing) != 0 || len(result.Additional) != 0 || len(result.Unattributed) != 0 {
		return fmt.Errorf("unexpected classification: missing %v additional %v unattributed %v", result.Missing, result.Additional, result.Unattributed)
	}
	return nil
}

func TestHappyPathAttributesSignalsAndCleansUp(t *testing.T) {
	for _, harnessID := range evaluationHarnesses {
		t.Run(harnessID, func(t *testing.T) {
			t.Parallel()
			run := newConformanceRun(t, harnessID, "happy")
			var extraArgs []string
			if harnessID == "claude-code" {
				extraArgs = append(extraArgs, "--transcript")
			}
			code, stdout, stderr := run.execute(extraArgs, nil)
			if code != 0 {
				t.Fatalf("run failed: %s", stderr)
			}
			result := run.completedResult(stdout)
			if err := expectedConformanceCalls(result); err != nil {
				t.Fatal(err)
			}
			directory := run.runDirectory()
			artifacts := []string{"result.json", "website.json", "events.jsonl"}
			if harnessID == "claude-code" {
				artifacts = append(artifacts, "transcript.json")
			}
			for _, name := range artifacts {
				if _, err := os.Stat(filepath.Join(directory, name)); err != nil {
					t.Errorf("missing artifact %s: %v", name, err)
				}
			}
			if _, err := os.Stat(filepath.Join(directory, "failure.json")); !os.IsNotExist(err) {
				t.Errorf("passing run wrote failure.json: %v", err)
			}
			run.assertCleanup()
			switch harnessID {
			case "opencode":
				data, err := os.ReadFile(filepath.Join(run.fakeDir, "sessions.json"))
				if err != nil || strings.TrimSpace(string(data)) != "[]" {
					t.Errorf("OpenCode session was not deleted: %q %v", data, err)
				}
			case "kilo-code":
				data, err := os.ReadFile(filepath.Join(run.fakeDir, "deleted-sessions"))
				if err != nil || !strings.Contains(string(data), "kilo-code-session-0001") {
					t.Errorf("Kilo session was not deleted: %q %v", data, err)
				}
			}
		})
	}
}

func TestMidTurnDeathWritesTurnAttributedDiagnostics(t *testing.T) {
	for _, harnessID := range evaluationHarnesses {
		t.Run(harnessID, func(t *testing.T) {
			t.Parallel()
			run := newConformanceRun(t, harnessID, "die-on-resume")
			code, _, stderr := run.execute(nil, nil)
			if code == 0 {
				t.Fatal("mid-turn death was reported as success")
			}
			if !strings.Contains(stderr, "evaluation encountered a tooling error") {
				t.Fatalf("stderr does not report a tooling error: %s", stderr)
			}
			record := run.failureRecord()
			if record.TurnID != "turn-2" {
				t.Errorf("failure not attributed to turn-2: %#v", record)
			}
			if record.Command == "" {
				t.Errorf("failure record lacks the harness command: %#v", record)
			}
			if !strings.Contains(record.Stderr, "died mid-turn") {
				t.Errorf("failure record lacks native stderr: %#v", record)
			}
			run.assertCleanup()
		})
	}
}

func TestSessionIDChangeIsProtocolViolation(t *testing.T) {
	for _, harnessID := range evaluationHarnesses {
		t.Run(harnessID, func(t *testing.T) {
			t.Parallel()
			run := newConformanceRun(t, harnessID, "session-change")
			code, _, stderr := run.execute(nil, nil)
			if code == 0 {
				t.Fatal("session ID change was reported as success")
			}
			if !strings.Contains(stderr, "evaluation encountered a tooling error") {
				t.Fatalf("stderr does not report a tooling error: %s", stderr)
			}
			record := run.failureRecord()
			if !strings.Contains(record.Error, "session ID") {
				t.Errorf("failure does not name the session ID change: %#v", record)
			}
			if record.TurnID != "turn-2" {
				t.Errorf("failure not attributed to turn-2: %#v", record)
			}
			run.assertCleanup()
		})
	}
}

func TestIncompleteProtocolIsViolation(t *testing.T) {
	cases := map[string]struct {
		mode    string
		message string
	}{
		"claude-code": {mode: "missing-completion", message: "system/init and result"},
		"codex":       {mode: "missing-completion", message: "turn.completed"},
		"cursor":      {mode: "missing-completion", message: "system/init and result"},
		"opencode":    {mode: "missing-completion", message: "step_finish"},
		"kilo-code":   {mode: "missing-completion", message: "step_finish"},
		"pi":          {mode: "agent-error", message: "Pi agent ended with error"},
	}
	for _, harnessID := range evaluationHarnesses {
		t.Run(harnessID, func(t *testing.T) {
			t.Parallel()
			expectation := cases[harnessID]
			run := newConformanceRun(t, harnessID, expectation.mode)
			code, _, stderr := run.execute(nil, nil)
			if code == 0 {
				t.Fatal("incomplete protocol was reported as success")
			}
			record := run.failureRecord()
			if !strings.Contains(record.Error, expectation.message) {
				t.Errorf("failure %q does not name %q: %s", record.Error, expectation.message, stderr)
			}
			if record.TurnID != "turn-1" {
				t.Errorf("failure not attributed to turn-1: %#v", record)
			}
			run.assertCleanup()
		})
	}
}

func TestCodexConfigurationRejectionIsDiagnosed(t *testing.T) {
	t.Parallel()
	run := newConformanceRun(t, "codex", "config-reject")
	code, _, stderr := run.execute(nil, nil)
	if code == 0 {
		t.Fatal("rejected configuration was reported as success")
	}
	if !strings.Contains(stderr, "evaluation encountered a tooling error") {
		t.Fatalf("stderr does not report a tooling error: %s", stderr)
	}
	record := run.failureRecord()
	if !strings.Contains(record.Stderr, "Error loading config.toml") {
		t.Errorf("failure record lacks the native configuration error: %#v", record)
	}
	if record.TurnID != "turn-1" {
		t.Errorf("failure not attributed to turn-1: %#v", record)
	}
	run.assertCleanup()
}

func TestOpenCodeFailedMarkerIsToolingFailure(t *testing.T) {
	t.Parallel()
	run := newConformanceRun(t, "opencode", "marker-failure")
	code, _, _ := run.execute(nil, nil)
	if code == 0 {
		t.Fatal("failed marker command was reported as success")
	}
	record := run.failureRecord()
	if !strings.Contains(record.Error, "marker command failed") {
		t.Errorf("failure does not name the marker command: %#v", record)
	}
	run.assertCleanup()
}

func TestVersionPinBlocksUnqualifiedHarness(t *testing.T) {
	for _, harnessID := range []string{"opencode", "kilo-code"} {
		t.Run(harnessID, func(t *testing.T) {
			t.Parallel()
			run := newConformanceRun(t, harnessID, "happy")
			run.setMode("happy", "9.9.9")
			code, _, stderr := run.execute(nil, nil)
			if code == 0 {
				t.Fatal("unqualified version was accepted")
			}
			record := run.failureRecord()
			if !strings.Contains(record.Error, "qualified version") {
				t.Errorf("failure does not name the version pin: %s", stderr)
			}
			run.assertCleanup()
		})
	}
}

func TestVersionPinEscapeHatchWarnsAndProceeds(t *testing.T) {
	for _, harnessID := range []string{"opencode", "kilo-code"} {
		t.Run(harnessID, func(t *testing.T) {
			t.Parallel()
			run := newConformanceRun(t, harnessID, "happy")
			run.setMode("happy", "9.9.9")
			code, stdout, stderr := run.execute(nil, []string{"SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS=1"})
			if code != 0 {
				t.Fatalf("escape hatch did not allow the run: %s", stderr)
			}
			if !strings.Contains(stderr, "not the qualified version") {
				t.Errorf("escape hatch did not warn about the drift: %s", stderr)
			}
			result := run.completedResult(stdout)
			if err := expectedConformanceCalls(result); err != nil {
				t.Fatal(err)
			}
			run.assertCleanup()
		})
	}
}
