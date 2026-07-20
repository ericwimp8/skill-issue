package installer

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
)

func TestInjectedInstructionIsNeutral(t *testing.T) {
	source := []byte("---\nname: example\ndescription: Example.\n---\n\n# Example\n")
	temporaryRoot := t.TempDir()
	cliPath := filepath.Join(temporaryRoot, "skill-issue")
	stateRoot := filepath.Join(temporaryRoot, "state")
	result, err := inject(source, cliPath, stateRoot, "opaque-token")
	if err != nil {
		t.Fatal(err)
	}
	text := string(result)
	expected := fmt.Sprintf("Run %q signal %q %q, then continue normally.", cliPath, "opaque-token", stateRoot)
	if !strings.Contains(text, expected) {
		t.Fatalf("neutral instruction is absent: %s", text)
	}
	for _, revealing := range []string{"evaluate", "evaluation", "answer sheet", "expected call", "score", " mark "} {
		if strings.Contains(strings.ToLower(text), revealing) {
			t.Fatalf("generated skill contains revealing term %q", revealing)
		}
	}
}

func TestEvaluationPreparationRejectsUserScope(t *testing.T) {
	_, _, err := New().PrepareEvaluation(Request{Scope: "user"})
	if err == nil {
		t.Fatal("user-scope evaluation installation was accepted")
	}
}
