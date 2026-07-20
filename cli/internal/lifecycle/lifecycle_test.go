package lifecycle

import "testing"

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

	withoutOutput := make(map[string]string, len(base)-1)
	for key, value := range base {
		if key != "output" {
			withoutOutput[key] = value
		}
	}
	if _, err := evaluationRunRequest(withoutOutput); err == nil {
		t.Fatal("evaluation run without output was accepted")
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

func TestEvaluationCleanupAndSignalRequireOutputOwnedState(t *testing.T) {
	service := New()
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
