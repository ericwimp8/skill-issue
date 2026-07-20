package payload

import (
	"encoding/json"
	"fmt"
	"slices"
	"testing"
)

func TestSkillsMatchCanonicalManifest(t *testing.T) {
	skills, err := Skills()
	if err != nil {
		t.Fatal(err)
	}
	names := make([]string, 0, len(skills))
	for _, skill := range skills {
		names = append(names, skill.Name)
	}
	expected := []string{
		"document-update-discipline",
		"prompt-writing",
		"skill-authoring-discipline",
		"skill-evaluation-and-refinement",
		"skill-generation",
		"skill-intake",
		"system-change-ownership",
	}
	if !slices.Equal(names, expected) {
		t.Fatalf("unexpected canonical skills: %v", names)
	}
}

func TestBuiltInEvaluationsAreComplete(t *testing.T) {
	identifiers := []string{
		"gardening-web-application",
		"community-archive-desktop-application",
		"neighborhood-emergency-preparedness-program",
	}
	for _, identifier := range identifiers {
		t.Run(identifier, func(t *testing.T) {
			data, err := BuiltInEvaluation(identifier)
			if err != nil {
				t.Fatal(err)
			}
			var unit struct {
				SchemaVersion int    `json:"schema_version"`
				EvaluationID  string `json:"evaluation_id"`
				Scenario      struct {
					ScenarioID string `json:"scenario_id"`
					Turns      []struct {
						TurnID string `json:"turn_id"`
						Prompt string `json:"prompt"`
					} `json:"turns"`
				} `json:"scenario"`
				AnswerSheet struct {
					ScenarioID string `json:"scenario_id"`
					Expected   []struct {
						TurnID string `json:"turn_id"`
						Skill  string `json:"skill"`
					} `json:"expected"`
				} `json:"answer_sheet"`
			}
			if err := json.Unmarshal(data, &unit); err != nil {
				t.Fatal(err)
			}
			if unit.SchemaVersion != 1 || unit.EvaluationID != identifier {
				t.Fatalf("unexpected embedded identity: %#v", unit)
			}
			if unit.Scenario.ScenarioID != identifier || unit.AnswerSheet.ScenarioID != identifier {
				t.Fatalf("scenario and answer sheet do not match identifier")
			}
			if len(unit.Scenario.Turns) != 30 || len(unit.AnswerSheet.Expected) != 4 {
				t.Fatalf("unexpected governed evaluation shape: %d turns, %d expected calls", len(unit.Scenario.Turns), len(unit.AnswerSheet.Expected))
			}
			for turnIndex, turn := range unit.Scenario.Turns {
				if turn.TurnID != fmt.Sprintf("turn-%d", turnIndex+1) || turn.Prompt == "" {
					t.Fatalf("embedded turn is incomplete or unordered: %#v", turn)
				}
			}
		})
	}
}
