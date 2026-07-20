package replay

import (
	"os"
	"strings"
	"testing"
)

func TestMergedEnvironmentReplacesInheritedValues(t *testing.T) {
	t.Setenv("CODEX_HOME", "inherited-home")
	environment := mergedEnvironment([]string{"CODEX_HOME=replacement-home", "CODEX_SQLITE_HOME=replacement-home"})

	var homes []string
	for _, entry := range environment {
		if strings.HasPrefix(entry, "CODEX_HOME=") {
			homes = append(homes, entry)
		}
	}
	if len(homes) != 1 || homes[0] != "CODEX_HOME=replacement-home" {
		t.Fatalf("unexpected CODEX_HOME entries: %v", homes)
	}
	if len(environment) < len(os.Environ()) {
		t.Fatal("environment overlay removed unrelated inherited values")
	}
}

func TestCodexCommandsUseBoundedNonInteractivePermissions(t *testing.T) {
	for _, args := range [][]string{codexInitial("prompt"), codexResume("session", "prompt")} {
		joined := strings.Join(args, " ")
		if !strings.Contains(joined, "--ask-for-approval never") || !strings.Contains(joined, "--sandbox workspace-write") {
			t.Fatalf("Codex command lacks bounded permissions: %v", args)
		}
		if strings.Contains(joined, "danger-full-access") || strings.Contains(joined, "dangerously-bypass") {
			t.Fatalf("Codex command widens permissions: %v", args)
		}
	}
}
