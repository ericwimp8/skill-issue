package evaluation

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/replay"
	"github.com/ericwimp8/skill-issue/cli/internal/runstate"
)

func TestCustomInputsUseExistingScenarioAndAnswerShapes(t *testing.T) {
	directory := t.TempDir()
	workspace := filepath.Join(directory, "workspace")
	if err := os.Mkdir(workspace, 0o700); err != nil {
		t.Fatal(err)
	}
	scenarioPath := filepath.Join(directory, "scenario.json")
	answerPath := filepath.Join(directory, "answer.json")
	skillsPath := filepath.Join(directory, "skills")
	writeSkillFixture(t, skillsPath, "prompt-writing")
	scenario := map[string]any{
		"schema_version": 1,
		"scenario_id":    "custom",
		"turns": []map[string]string{{
			"turn_id": "turn-1",
			"prompt":  "Write a concise prompt for another agent.",
		}},
	}
	answer := map[string]any{
		"schema_version": 1,
		"scenario_id":    "custom",
		"expected": []map[string]string{{
			"turn_id": "turn-1",
			"skill":   "prompt-writing",
		}},
	}
	writeJSONFixture(t, scenarioPath, scenario)
	writeJSONFixture(t, answerPath, answer)

	loadedScenario, loadedAnswer, err := loadInputs(RunRequest{Workspace: workspace, SkillsPath: skillsPath, ScenarioPath: scenarioPath, AnswerSheet: answerPath})
	if err != nil {
		t.Fatal(err)
	}
	if loadedScenario.ID != "custom" || len(loadedScenario.Turns) != 1 || len(loadedAnswer.Expected) != 1 {
		t.Fatalf("unexpected custom inputs: %#v %#v", loadedScenario, loadedAnswer)
	}

	insideAnswer := filepath.Join(workspace, "answer.json")
	writeJSONFixture(t, insideAnswer, answer)
	if _, _, err := loadInputs(RunRequest{Workspace: workspace, SkillsPath: skillsPath, ScenarioPath: scenarioPath, AnswerSheet: insideAnswer}); err == nil {
		t.Fatal("answer sheet inside evaluated workspace was accepted")
	}
}

func TestBuiltInIdentifierLoadsScenarioAndAnswerTogether(t *testing.T) {
	shapes := map[string]struct {
		turns    int
		expected int
	}{
		"gardening-web-application":                   {turns: 30, expected: 46},
		"community-archive-desktop-application":       {turns: 30, expected: 46},
		"neighborhood-emergency-preparedness-program": {turns: 30, expected: 45},
	}
	for identifier, shape := range shapes {
		t.Run(identifier, func(t *testing.T) {
			scenario, answer, err := loadInputs(RunRequest{EvaluationID: identifier})
			if err != nil {
				t.Fatal(err)
			}
			if scenario.ID != identifier || answer.ScenarioID != scenario.ID {
				t.Fatalf("built-in parts do not match: %#v %#v", scenario, answer)
			}
			if len(scenario.Turns) != shape.turns || len(answer.Expected) != shape.expected {
				t.Fatalf("unexpected built-in shape: %d turns, %d expected calls", len(scenario.Turns), len(answer.Expected))
			}
		})
	}
}

func TestWebsiteResultPreservesOrderAndDoesNotInflateRepeatedCalls(t *testing.T) {
	scenario := replay.Scenario{
		SchemaVersion: 1,
		ID:            "custom",
		Turns: []replay.Turn{
			{ID: "opening", Prompt: "one"},
			{ID: "middle", Prompt: "two"},
			{ID: "review-step", Prompt: "three"},
			{ID: "finish", Prompt: "four"},
		},
	}
	result := Result{
		RunID:      "run",
		ScenarioID: "custom",
		Harness:    "codex",
		Model:      "model",
		Expected: []SkillCall{
			{TurnID: "opening", Skill: "plan-maintenance"},
			{TurnID: "opening", Skill: "document-update-discipline"},
			{TurnID: "opening", Skill: "plan-maintenance"},
			{TurnID: "review-step", Skill: "prompt-writing"},
		},
		Observed: []SkillCall{
			{TurnID: "opening", Skill: "plan-maintenance"},
			{TurnID: "opening", Skill: "plan-maintenance"},
			{TurnID: "opening", Skill: "document-update-discipline"},
			{TurnID: "opening", Skill: "prompt-writing"},
			{TurnID: "opening", Skill: "prompt-writing"},
			{TurnID: "finish", Skill: "prompt-writing"},
			{TurnID: "finish", Skill: "prompt-writing"},
		},
		Additional: []SkillCall{
			{TurnID: "opening", Skill: "prompt-writing"},
			{TurnID: "opening", Skill: "prompt-writing"},
			{TurnID: "finish", Skill: "prompt-writing"},
			{TurnID: "finish", Skill: "prompt-writing"},
		},
	}

	website := deriveWebsiteResult(result, scenario)
	if website.SchemaVersion != 2 || website.TotalTurns != 4 || len(website.Points) != 3 {
		t.Fatalf("unexpected website shape: %#v", website)
	}
	if website.Points[0] != (WebsitePoint{Turn: 1, TurnID: "opening", Called: 2, Missed: 0, Unexpected: 1}) {
		t.Fatalf("unexpected opening point: %#v", website.Points[0])
	}
	if website.Points[1] != (WebsitePoint{Turn: 3, TurnID: "review-step", Called: 0, Missed: 1}) {
		t.Fatalf("unexpected review point: %#v", website.Points[1])
	}
	if website.Points[2] != (WebsitePoint{Turn: 4, TurnID: "finish", Unexpected: 1}) {
		t.Fatalf("unexpected finish point: %#v", website.Points[2])
	}
}

func TestWebsiteResultProjectsAdditionalCallsAsUnexpected(t *testing.T) {
	scenario := replay.Scenario{
		SchemaVersion: 1,
		ID:            "custom",
		Turns: []replay.Turn{
			{ID: "expected", Prompt: "one"},
			{ID: "unexpected", Prompt: "two"},
		},
	}
	result := deriveResult(
		"run",
		RunRequest{Harness: harness.Codex, Model: "model", Reasoning: "medium"},
		"evaluation",
		scenario.ID,
		time.Now(),
		[]SkillCall{{TurnID: "expected", Skill: "prompt-writing"}},
		[]runstate.Event{
			{TurnID: "unexpected", Skill: "prompt-writing", Attributed: true},
			{TurnID: "unexpected", Skill: "prompt-writing", Attributed: true},
			{TurnID: "unexpected", Skill: "document-update-discipline"},
		},
	)

	website := deriveWebsiteResult(result, scenario)
	if len(result.Additional) != 2 || len(result.Unattributed) != 1 {
		t.Fatalf("unexpected detailed classification: %#v", result)
	}
	if len(website.Points) != 2 || website.Points[1] != (WebsitePoint{Turn: 2, TurnID: "unexpected", Unexpected: 1}) {
		t.Fatalf("unexpected website projection: %#v", website)
	}
}

func TestGardeningWebsitePointsUseEveryExpectedTurn(t *testing.T) {
	scenario, answer, err := loadInputs(RunRequest{EvaluationID: "gardening-web-application"})
	if err != nil {
		t.Fatal(err)
	}
	website := deriveWebsiteResult(Result{
		RunID:      "run",
		ScenarioID: scenario.ID,
		Harness:    "codex",
		Model:      "model",
		Expected:   answer.Expected,
	}, scenario)
	if website.TotalTurns != 30 || len(website.Points) != 27 {
		t.Fatalf("unexpected gardening website shape: %#v", website)
	}
	expectedTurns := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 14, 15, 16, 17, 19, 20, 21, 22, 23, 25, 26, 27, 28, 29, 30}
	for index, turn := range expectedTurns {
		if website.Points[index].Turn != turn {
			t.Fatalf("point %d uses turn %d", index, website.Points[index].Turn)
		}
	}
	if website.Points[0].Called != 0 || website.Points[0].Missed != 2 {
		t.Fatalf("turn 1 did not retain its expected calls: %#v", website.Points[0])
	}
}

func TestToolingCompleteRunWritesWebsiteArtifact(t *testing.T) {
	directory := t.TempDir()
	workspace := filepath.Join(directory, "workspace")
	inputs := filepath.Join(directory, "inputs")
	output := filepath.Join(directory, "output")
	state := filepath.Join(output, ".skill-issue")
	if err := os.Mkdir(workspace, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.Mkdir(inputs, 0o700); err != nil {
		t.Fatal(err)
	}
	scenarioPath := filepath.Join(inputs, "scenario.json")
	answerPath := filepath.Join(inputs, "answer.json")
	skillsPath := filepath.Join(inputs, "skills")
	writeSkillFixture(t, skillsPath, "prompt-writing")
	writeJSONFixture(t, scenarioPath, map[string]any{
		"schema_version": 1,
		"scenario_id":    "file-output",
		"turns": []map[string]string{{
			"turn_id": "alpha",
			"prompt":  "Write a concise prompt.",
		}},
	})
	writeJSONFixture(t, answerPath, map[string]any{
		"schema_version": 1,
		"scenario_id":    "file-output",
		"expected": []map[string]string{{
			"turn_id": "alpha",
			"skill":   "prompt-writing",
		}},
	})
	cliPath, err := os.Executable()
	if err != nil {
		t.Fatal(err)
	}
	service := New(state)
	var runtimeRoot string
	service.adapterFactory = func(id replay.HarnessID, options replay.Options) (replay.Adapter, error) {
		runtimeRoot = filepath.Dir(options.ClaudeSkillsRoot)
		return staticAdapter{id: id}, nil
	}
	result, err := service.Run(context.Background(), RunRequest{
		Workspace:         workspace,
		OutputRoot:        output,
		Harness:           harness.ClaudeCode,
		Model:             "model",
		ScenarioPath:      scenarioPath,
		AnswerSheet:       answerPath,
		SkillsPath:        skillsPath,
		CLIPath:           cliPath,
		IncludeEvents:     true,
		IncludeTranscript: true,
	})
	if err != nil {
		t.Fatal(err)
	}
	outputDirectories, err := filepath.Glob(filepath.Join(output, "claude-code-*-"+result.RunID[:8]))
	if err != nil || len(outputDirectories) != 1 {
		t.Fatalf("unexpected output directories: %v %v", outputDirectories, err)
	}
	outputDirectory := outputDirectories[0]
	for _, name := range []string{"events.jsonl", "transcript.json", "result.json", "website.json"} {
		if _, err := os.Stat(filepath.Join(outputDirectory, name)); err != nil {
			t.Fatalf("missing output artifact %s: %v", name, err)
		}
	}
	data, err := os.ReadFile(filepath.Join(outputDirectory, "website.json"))
	if err != nil {
		t.Fatal(err)
	}
	var website WebsiteResult
	if err := json.Unmarshal(data, &website); err != nil {
		t.Fatal(err)
	}
	if website.SchemaVersion != 2 || website.RunID != result.RunID || website.TotalTurns != 1 || len(website.Points) != 1 {
		t.Fatalf("unexpected website file: %#v", website)
	}
	if _, err := os.Stat(filepath.Join(output, ".skill-issue", "runs", result.RunID, "run.json")); !os.IsNotExist(err) {
		t.Fatalf("private run state was retained after cleanup: %v", err)
	}
	if _, err := os.Stat(runtimeRoot); !os.IsNotExist(err) {
		t.Fatalf("private harness runtime was not removed: %v", err)
	}
}

func TestEvaluationOutputMustRemainOutsideWorkspace(t *testing.T) {
	workspace := t.TempDir()
	if _, err := prepareOutputRoot(workspace, filepath.Join(workspace, "output")); err == nil {
		t.Fatal("evaluation output inside workspace was accepted")
	}
}

type staticAdapter struct {
	id replay.HarnessID
}

func (adapter staticAdapter) HarnessID() replay.HarnessID { return adapter.id }
func (adapter staticAdapter) Start(context.Context) (replay.Session, error) {
	return staticSession{}, nil
}

type staticSession struct{}

func (staticSession) SendPrompt(context.Context, string) error { return nil }
func (staticSession) Wait(context.Context) (replay.Capture, error) {
	return replay.Capture{SessionID: "session"}, nil
}
func (staticSession) Close() error { return nil }

func writeJSONFixture(t *testing.T, path string, value any) {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		t.Fatal(err)
	}
}

func writeSkillFixture(t *testing.T, root, name string) {
	t.Helper()
	directory := filepath.Join(root, name)
	if err := os.MkdirAll(directory, 0o700); err != nil {
		t.Fatal(err)
	}
	content := "---\nname: " + name + "\ndescription: Test skill.\n---\n\n# Test Skill\n"
	if err := os.WriteFile(filepath.Join(directory, "SKILL.md"), []byte(content), 0o600); err != nil {
		t.Fatal(err)
	}
}

func TestPrepareRequestCachesInputsForTheRun(t *testing.T) {
	directory := t.TempDir()
	workspace := filepath.Join(directory, "workspace")
	if err := os.Mkdir(workspace, 0o700); err != nil {
		t.Fatal(err)
	}
	scenarioPath := filepath.Join(directory, "scenario.json")
	answerPath := filepath.Join(directory, "answer.json")
	skillsPath := filepath.Join(directory, "skills")
	writeSkillFixture(t, skillsPath, "prompt-writing")
	writeJSONFixture(t, scenarioPath, map[string]any{
		"schema_version": 1,
		"scenario_id":    "custom",
		"turns":          []map[string]string{{"turn_id": "turn-1", "prompt": "Write a concise prompt."}},
	})
	writeJSONFixture(t, answerPath, map[string]any{
		"schema_version": 1,
		"scenario_id":    "custom",
		"expected":       []map[string]string{{"turn_id": "turn-1", "skill": "prompt-writing"}},
	})

	prepared, err := PrepareRequest(RunRequest{Harness: harness.Codex, Workspace: workspace, SkillsPath: skillsPath, ScenarioPath: scenarioPath, AnswerSheet: answerPath})
	if err != nil {
		t.Fatal(err)
	}
	if prepared.inputs == nil || prepared.AvailableTurns != 1 {
		t.Fatalf("prepared request did not cache inputs: %#v", prepared)
	}

	// Once the user has reviewed the prepared request, changing or deleting
	// the input files must not affect what runs.
	if err := os.Remove(scenarioPath); err != nil {
		t.Fatal(err)
	}
	if err := os.Remove(answerPath); err != nil {
		t.Fatal(err)
	}
	again, err := PrepareRequest(prepared)
	if err != nil {
		t.Fatalf("prepared request re-read deleted inputs: %v", err)
	}
	if again.inputs != prepared.inputs || again.AvailableTurns != 1 {
		t.Fatalf("re-prepared request lost its cached inputs: %#v", again)
	}
}

func TestCustomInputPathsMayBeRelative(t *testing.T) {
	directory := t.TempDir()
	workspace := filepath.Join(directory, "workspace")
	if err := os.Mkdir(workspace, 0o700); err != nil {
		t.Fatal(err)
	}
	writeSkillFixture(t, filepath.Join(directory, "skills"), "prompt-writing")
	writeJSONFixture(t, filepath.Join(directory, "scenario.json"), map[string]any{
		"schema_version": 1,
		"scenario_id":    "custom",
		"turns":          []map[string]string{{"turn_id": "turn-1", "prompt": "Write a concise prompt."}},
	})
	writeJSONFixture(t, filepath.Join(directory, "answer.json"), map[string]any{
		"schema_version": 1,
		"scenario_id":    "custom",
		"expected":       []map[string]string{{"turn_id": "turn-1", "skill": "prompt-writing"}},
	})
	t.Chdir(directory)
	if _, _, err := loadInputs(RunRequest{Workspace: workspace, SkillsPath: "skills", ScenarioPath: "scenario.json", AnswerSheet: "answer.json"}); err != nil {
		t.Fatalf("relative custom input paths were rejected: %v", err)
	}
}
