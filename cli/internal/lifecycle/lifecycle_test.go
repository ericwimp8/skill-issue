package lifecycle

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
	"github.com/ericwimp8/skill-issue/cli/internal/replay"
)

func TestEvaluationRunIsProjectLocalAndHasOneInputMode(t *testing.T) {
	base := map[string]string{
		"workspace":  t.TempDir(),
		"output":     t.TempDir(),
		"harness":    "codex",
		"model":      "model",
		"evaluation": "gardening-web-application",
	}
	request, err := evaluationRunRequest(base)
	if err != nil {
		t.Fatal(err)
	}
	if request.Workspace == "" || request.OutputRoot == "" || request.EvaluationID != "gardening-web-application" {
		t.Fatalf("unexpected evaluation request: %#v", request)
	}

	withCustom := make(map[string]string, len(base)+2)
	for key, value := range base {
		withCustom[key] = value
	}
	withCustom["scenario"] = "scenario.json"
	withCustom["answer-sheet"] = "answer.json"
	if _, err := evaluationRunRequest(withCustom); err == nil {
		t.Fatal("built-in and custom inputs were accepted together")
	}
}

func TestEvaluationRunDefaultsWorkspaceOutputAndEvaluation(t *testing.T) {
	directory := t.TempDir()
	t.Chdir(directory)
	request, err := evaluationRunRequest(map[string]string{"harness": "codex"})
	if err != nil {
		t.Fatal(err)
	}
	if !request.WorkspaceCreated || !strings.Contains(filepath.Base(request.Workspace), "skill-issue-workspace-") {
		t.Fatalf("workspace was not created with the default name: %#v", request)
	}
	if info, err := os.Stat(request.Workspace); err != nil || !info.IsDir() {
		t.Fatalf("default workspace does not exist: %v", err)
	}
	if filepath.Base(request.OutputRoot) != "skill-issue-output" || !filepath.IsAbs(request.OutputRoot) {
		t.Fatalf("output root was not defaulted: %q", request.OutputRoot)
	}
	if request.EvaluationID != defaultEvaluationID || request.AvailableTurns != 30 {
		t.Fatalf("built-in evaluation was not defaulted: %#v", request)
	}

	// A second defaulted run must get its own fresh workspace.
	second, err := evaluationRunRequest(map[string]string{"harness": "codex"})
	if err != nil {
		t.Fatal(err)
	}
	if second.Workspace == request.Workspace {
		t.Fatal("defaulted workspaces were reused across runs")
	}
}

func TestEvaluationRunCreatesANamedMissingWorkspace(t *testing.T) {
	directory := t.TempDir()
	t.Chdir(directory)
	request, err := evaluationRunRequest(map[string]string{"harness": "codex", "workspace": "fresh-workspace"})
	if err != nil {
		t.Fatal(err)
	}
	if !request.WorkspaceCreated {
		t.Fatalf("named missing workspace was not marked created: %#v", request)
	}
	if info, err := os.Stat(filepath.Join(directory, "fresh-workspace")); err != nil || !info.IsDir() {
		t.Fatalf("named workspace was not created: %v", err)
	}

	// An existing workspace is used as-is and never marked created.
	again, err := evaluationRunRequest(map[string]string{"harness": "codex", "workspace": "fresh-workspace"})
	if err != nil {
		t.Fatal(err)
	}
	if again.WorkspaceCreated {
		t.Fatalf("existing workspace was marked created: %#v", again)
	}
}

func TestCancelledEvaluationRemovesItsCreatedWorkspace(t *testing.T) {
	directory := t.TempDir()
	t.Chdir(directory)
	service := New(nil)
	decline := func(evaluation.RunRequest) (bool, error) { return false, nil }
	result, err := service.ExecuteEvaluationRun([]string{"--harness", "codex"}, decline, nil)
	if err != nil || result.Status != "cancelled" {
		t.Fatalf("cancelled run result: %#v %v", result, err)
	}
	leftovers, err := filepath.Glob(filepath.Join(directory, "skill-issue-workspace-*"))
	if err != nil || len(leftovers) != 0 {
		t.Fatalf("cancelled run left its created workspace behind: %v %v", leftovers, err)
	}
	if _, err := os.Stat(filepath.Join(directory, "skill-issue-output")); !os.IsNotExist(err) {
		t.Fatalf("cancelled run created the output root: %v", err)
	}
}

func TestEvaluationCleanupAndSignalRequireOutputOwnedState(t *testing.T) {
	service := New(nil)
	if _, err := service.evaluate([]string{"cleanup", "--run", "run-id"}); err == nil {
		t.Fatal("evaluation cleanup without output was accepted")
	}
	if _, err := service.mark([]string{"opaque-token"}); err == nil {
		t.Fatal("signal without state root was accepted")
	}
	if _, err := service.mark([]string{"opaque-token", "relative-state"}); err == nil {
		t.Fatal("signal with relative state root was accepted")
	}
}

func TestTurnProgressRendererReportsStartAndFinishWithStats(t *testing.T) {
	var output bytes.Buffer
	renderer := newTurnProgressRenderer(&output)
	defer renderer.stop()
	renderer.handle(evaluation.TurnProgress{TurnID: "turn-2", Index: 2, Total: 3, Phase: replay.BoundaryBefore})
	renderer.handle(evaluation.TurnProgress{TurnID: "turn-2", Index: 2, Total: 3, Phase: replay.BoundaryAfter, Duration: 12 * time.Second, HarnessEvents: 34, SkillCalls: 1})
	wanted := "Starting turn 2 of 3: turn-2\nFinished turn 2 of 3: turn-2 (12s, 34 harness events, 1 skill call)\n"
	if output.String() != wanted {
		t.Fatalf("progress output = %q", output.String())
	}
	if renderer.interactive {
		t.Fatal("plain writer was treated as an interactive terminal")
	}
}

func TestParseOptionsSupportsInlineValues(t *testing.T) {
	options, err := parseOptions([]string{"--workspace=/tmp/w", "--events", "--transcript=false", "--scenario=--odd-name"}, "events", "transcript")
	if err != nil {
		t.Fatal(err)
	}
	if options["workspace"] != "/tmp/w" || options["events"] != "true" || options["transcript"] != "false" || options["scenario"] != "--odd-name" {
		t.Fatalf("unexpected options: %v", options)
	}
	if _, err := parseOptions([]string{"--events=maybe"}, "events"); err == nil {
		t.Fatal("non-boolean inline value for a boolean flag was accepted")
	}
	if _, err := parseOptions([]string{"--workspace=/a", "--workspace", "/b"}); err == nil {
		t.Fatal("duplicate option was accepted")
	}
	if _, err := parseOptions([]string{"--workspace"}); err == nil {
		t.Fatal("value option without a value was accepted")
	}
}

func TestRequiredDistinguishesMissingFromEmpty(t *testing.T) {
	if _, err := required(map[string]string{}, "output"); err == nil || !strings.Contains(err.Error(), "required") {
		t.Fatalf("missing option error = %v", err)
	}
	if _, err := required(map[string]string{"output": ""}, "output"); err == nil || !strings.Contains(err.Error(), "empty") {
		t.Fatalf("explicitly empty option error = %v", err)
	}
}

func TestTurnProgressRendererSpinnerRendersAndClearsWhenInteractive(t *testing.T) {
	var output bytes.Buffer
	renderer := newTurnProgressRenderer(&output)
	renderer.interactive = true
	renderer.handle(evaluation.TurnProgress{TurnID: "turn-1", Index: 1, Total: 2, Phase: replay.BoundaryBefore})
	time.Sleep(300 * time.Millisecond)
	renderer.handle(evaluation.TurnProgress{TurnID: "turn-1", Index: 1, Total: 2, Phase: replay.BoundaryAfter})
	text := output.String()
	if !strings.Contains(text, "turn 1 of 2 running") {
		t.Fatalf("spinner label missing from output: %q", text)
	}
	if !strings.Contains(text, "\r\x1b[2K") {
		t.Fatalf("spinner did not redraw and clear its line: %q", text)
	}
	if !strings.HasSuffix(text, "Finished turn 1 of 2: turn-1 (0s, 0 harness events, 0 skill calls)\n") {
		t.Fatalf("finished line missing after spinner: %q", text)
	}
}
