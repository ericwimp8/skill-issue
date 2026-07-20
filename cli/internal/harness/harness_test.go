package harness

import (
	"path/filepath"
	"testing"
)

func TestEveryHarnessRetainsItsNativeProjectSkillRoot(t *testing.T) {
	workspace := t.TempDir()
	expected := map[ID]string{
		ClaudeCode: ".claude/skills",
		Codex:      ".agents/skills",
		Cursor:     ".cursor/skills",
		OpenCode:   ".opencode/skills",
		KiloCode:   ".kilo/skills",
		Pi:         ".pi/skills",
	}
	for id, relative := range expected {
		root, err := SkillRoot(id, ScopeProject, workspace, "")
		if err != nil {
			t.Fatalf("%s: %v", id, err)
		}
		if root != filepath.Join(workspace, filepath.FromSlash(relative)) {
			t.Fatalf("%s project root = %s", id, root)
		}
	}
}

func TestOpenCodeIsAvailableForInstallationAndEvaluation(t *testing.T) {
	if !InstallationAvailable(OpenCode) {
		t.Fatal("OpenCode installation is unavailable")
	}
	defaults, err := EvaluationDefaultsFor(OpenCode)
	if err != nil {
		t.Fatal(err)
	}
	if defaults.Model != "openai/gpt-5.6-sol" || defaults.Reasoning != "medium" {
		t.Fatalf("unexpected OpenCode defaults: %#v", defaults)
	}
}

func TestKiloIsAvailableForInstallationAndEvaluation(t *testing.T) {
	if !InstallationAvailable(KiloCode) {
		t.Fatal("Kilo installation is unavailable")
	}
	defaults, err := EvaluationDefaultsFor(KiloCode)
	if err != nil {
		t.Fatal(err)
	}
	if defaults.Model != "openai/gpt-5.6-sol" || defaults.Reasoning != "medium" {
		t.Fatalf("unexpected Kilo defaults: %#v", defaults)
	}
}
