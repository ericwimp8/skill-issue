package installer

import (
	"strings"
	"testing"
)

func TestInjectedInstructionIsNeutral(t *testing.T) {
	source := []byte("---\nname: example\ndescription: Example.\n---\n\n# Example\n")
	result, err := inject(source, "/opt/skill-issue", "opaque-token")
	if err != nil {
		t.Fatal(err)
	}
	text := string(result)
	if !strings.Contains(text, "Run \"/opt/skill-issue\" signal \"opaque-token\", then continue normally.") {
		t.Fatalf("neutral instruction is absent: %s", text)
	}
	for _, revealing := range []string{"evaluate", "evaluation", "answer sheet", "expected call", "score", " mark "} {
		if strings.Contains(strings.ToLower(text), revealing) {
			t.Fatalf("generated skill contains revealing term %q", revealing)
		}
	}
}

func TestEvaluationPreparationRejectsUserScope(t *testing.T) {
	_, _, err := New(t.TempDir()).PrepareEvaluation(Request{Scope: "user"}, t.TempDir())
	if err == nil {
		t.Fatal("user-scope evaluation installation was accepted")
	}
}
