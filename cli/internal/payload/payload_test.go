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
		"code-implementation-discipline",
		"code-testing-discipline",
		"dictate-plan",
		"document-update-discipline",
		"prompt-writing",
		"skill-authoring-discipline",
		"skill-evaluation-and-refinement",
		"skill-generation",
		"skill-intake",
		"system-change-ownership",
		"systematic-debugging",
	}
	if !slices.Equal(names, expected) {
		t.Fatalf("unexpected canonical skills: %v", names)
	}
}

func TestCanonicalManifestSources(t *testing.T) {
	manifest, err := ReadManifest()
	if err != nil {
		t.Fatal(err)
	}
	expected := map[string]string{
		"code-implementation-discipline":  "evaluations/scenario-skill-refinement/code-implementation-discipline/skill",
		"code-testing-discipline":         "evaluations/scenario-skill-refinement/code-testing-discipline/skill",
		"dictate-plan":                    "supporting-skills/dictate-plan",
		"document-update-discipline":      "evaluations/scenario-skill-refinement/document-update-discipline/skill",
		"prompt-writing":                  "evaluations/scenario-skill-refinement/prompt-writing/skill",
		"skill-authoring-discipline":      "evaluations/scenario-skill-refinement/skill-authoring-discipline/skill",
		"skill-evaluation-and-refinement": "skills/skill-evaluation-and-refinement",
		"skill-generation":                "skills/skill-generation",
		"skill-intake":                    "skills/skill-intake",
		"system-change-ownership":         "evaluations/scenario-skill-refinement/system-change-ownership/skill",
		"systematic-debugging":            "evaluations/scenario-skill-refinement/systematic-debugging/skill",
	}
	if len(manifest.Components) != len(expected) {
		t.Fatalf("unexpected manifest component count: %d", len(manifest.Components))
	}
	for _, component := range manifest.Components {
		if expected[component.ID] != component.Source {
			t.Fatalf("unexpected source for %q: %q", component.ID, component.Source)
		}
	}
}

func TestBuiltInEvaluationsAreComplete(t *testing.T) {
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
			if len(unit.Scenario.Turns) != shape.turns || len(unit.AnswerSheet.Expected) != shape.expected {
				t.Fatalf("unexpected governed evaluation shape: %d turns, %d expected calls", len(unit.Scenario.Turns), len(unit.AnswerSheet.Expected))
			}
			for turnIndex, turn := range unit.Scenario.Turns {
				if turn.TurnID != fmt.Sprintf("turn-%d", turnIndex+1) || turn.Prompt == "" {
					t.Fatalf("embedded turn is incomplete or unordered: %#v", turn)
				}
			}
			scoredTurns := make(map[string]bool, len(unit.Scenario.Turns))
			for _, expected := range unit.AnswerSheet.Expected {
				scoredTurns[expected.TurnID] = true
			}
			if len(scoredTurns) != 27 {
				t.Fatalf("unexpected scored-turn count: %d", len(scoredTurns))
			}
			for _, turnID := range []string{"turn-13", "turn-18", "turn-24"} {
				if scoredTurns[turnID] {
					t.Fatalf("factual reminder turn %s is scored", turnID)
				}
			}
		})
	}
}

func TestValidateReferenceClosureTracksFenceMarkers(t *testing.T) {
	entrypoint := "---\nname: demo\ndescription: d\n---\n" +
		"see `references/present.md`\n\n" +
		"~~~\n`references/absent.md`\n~~~\n\n" +
		"````\n`references/absent.md`\n```\n`references/absent.md`\n````\n"
	files := map[string][]byte{
		"SKILL.md":              []byte(entrypoint),
		"references/present.md": []byte("content"),
	}
	if err := validateReferenceClosure("demo", files); err != nil {
		t.Fatalf("fenced references were validated: %v", err)
	}
	broken := "---\nname: demo\ndescription: d\n---\nsee `references/absent.md`\n"
	if err := validateReferenceClosure("demo", map[string][]byte{"SKILL.md": []byte(broken)}); err == nil {
		t.Fatal("absent unfenced reference was accepted")
	}
}
