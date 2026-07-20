package installer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/payload"
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

func differingSkillRequest(t *testing.T) (Request, string, string) {
	t.Helper()
	ordinary, err := payload.Skills()
	if err != nil {
		t.Fatal(err)
	}
	name := ordinary[0].Name
	directory := t.TempDir()
	root := filepath.Join(directory, "skills-root")
	if err := os.MkdirAll(filepath.Join(root, name), 0o755); err != nil {
		t.Fatal(err)
	}
	userContent := "---\nname: " + name + "\ndescription: Locally modified.\n---\n\n# Local edits\n"
	if err := os.WriteFile(filepath.Join(root, name, "SKILL.md"), []byte(userContent), 0o644); err != nil {
		t.Fatal(err)
	}
	request := Request{
		Harness:         harness.Codex,
		Scope:           harness.ScopeProject,
		Workspace:       directory,
		EvaluationRoot:  root,
		CLIPath:         filepath.Join(directory, "skill-issue"),
		SignalStateRoot: filepath.Join(directory, "state"),
		BackupRoot:      filepath.Join(directory, "backup"),
		Tokens:          map[string]string{"opaque-token": name},
		Skills: []payload.Skill{{
			Name:  name,
			Files: map[string][]byte{"SKILL.md": []byte("---\nname: " + name + "\ndescription: Canonical.\n---\n\n# Canonical\n")},
		}},
	}
	return request, name, userContent
}

func TestPreexistingDifferingSkillRequiresConfirmation(t *testing.T) {
	request, _, _ := differingSkillRequest(t)
	if _, _, err := New().PrepareEvaluation(request); err == nil || !strings.Contains(err.Error(), "differ from their canonical versions") {
		t.Fatalf("differing preexisting skill was replaced without confirmation: %v", err)
	}

	request.ConfirmReplace = func([]string) (bool, error) { return false, nil }
	if _, _, err := New().PrepareEvaluation(request); !errors.Is(err, ErrReplacementDeclined) {
		t.Fatalf("declined replacement did not stop the evaluation: %v", err)
	}
}

func TestPreexistingSkillIsBackedUpAndRestoredExactly(t *testing.T) {
	request, name, userContent := differingSkillRequest(t)
	var asked []string
	request.ConfirmReplace = func(differing []string) (bool, error) {
		asked = differing
		return true, nil
	}
	state, _, err := New().PrepareEvaluation(request)
	if err != nil {
		t.Fatal(err)
	}
	if len(asked) != 1 || asked[0] != name {
		t.Fatalf("confirmation did not list the differing skill: %v", asked)
	}
	installedPath := filepath.Join(request.EvaluationRoot, name, "SKILL.md")
	installed, err := os.ReadFile(installedPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(installed) == userContent {
		t.Fatal("evaluation skill was not replaced")
	}
	if state.BackupRoot != request.BackupRoot {
		t.Fatalf("backup root was not recorded: %#v", state)
	}
	if err := New().CleanupEvaluation(request, state); err != nil {
		t.Fatal(err)
	}
	restored, err := os.ReadFile(installedPath)
	if err != nil {
		t.Fatal(err)
	}
	if string(restored) != userContent {
		t.Fatalf("preexisting skill was not restored byte-for-byte: %q", restored)
	}
}

func TestPreexistingCanonicalSkillProceedsWithoutConfirmation(t *testing.T) {
	directory := t.TempDir()
	workspace := filepath.Join(directory, "workspace")
	if err := os.MkdirAll(workspace, 0o755); err != nil {
		t.Fatal(err)
	}
	install := Request{Harness: harness.Codex, Scope: harness.ScopeProject, Workspace: workspace}
	if _, err := New().Install(install); err != nil {
		t.Fatal(err)
	}
	ordinary, err := payload.Skills()
	if err != nil {
		t.Fatal(err)
	}
	tokens := make(map[string]string, len(ordinary))
	for index, skill := range ordinary {
		tokens[fmt.Sprintf("opaque-token-%d", index)] = skill.Name
	}
	request := Request{
		Harness:         harness.Codex,
		Scope:           harness.ScopeProject,
		Workspace:       workspace,
		CLIPath:         filepath.Join(directory, "skill-issue"),
		SignalStateRoot: filepath.Join(directory, "state"),
		BackupRoot:      filepath.Join(directory, "backup"),
		Tokens:          tokens,
		ConfirmReplace: func(differing []string) (bool, error) {
			t.Fatalf("identical preexisting skills triggered confirmation: %v", differing)
			return false, nil
		},
	}
	state, _, err := New().PrepareEvaluation(request)
	if err != nil {
		t.Fatal(err)
	}
	if len(state.Preexisting) != len(ordinary) {
		t.Fatalf("expected every canonical skill to be preexisting: %#v", state)
	}
	if err := New().CleanupEvaluation(request, state); err != nil {
		t.Fatal(err)
	}
}
