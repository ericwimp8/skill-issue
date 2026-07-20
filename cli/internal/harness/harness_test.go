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
