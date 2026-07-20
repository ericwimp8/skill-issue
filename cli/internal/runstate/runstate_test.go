package runstate

import (
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
)

func testRun(id, token string) Run {
	return Run{
		SchemaVersion: 1,
		ID:            id,
		Harness:       "codex",
		Status:        "preparing",
		Tokens:        map[string]string{token: "prompt-writing"},
	}
}

func TestRunIDValidationRejectsPathTraversal(t *testing.T) {
	store := NewStore(t.TempDir())
	for _, runID := range []string{"", ".", "..", "../escape", `..\escape`, "nested/run"} {
		if _, err := store.Load(runID); err == nil || !strings.Contains(err.Error(), "run ID is invalid") {
			t.Fatalf("run ID %q was accepted by Load: %v", runID, err)
		}
		if _, err := store.Events(runID); err == nil || !strings.Contains(err.Error(), "run ID is invalid") {
			t.Fatalf("run ID %q was accepted by Events: %v", runID, err)
		}
		if err := store.DeleteRun(runID); err == nil {
			t.Fatalf("run ID %q was accepted by DeleteRun", runID)
		}
	}
}

func TestMarkRecoversFromStaleLock(t *testing.T) {
	token := strings.Repeat("ab", 32)
	store := NewStore(t.TempDir())
	runID := "stale-lock-run"
	if err := store.Create(testRun(runID, token)); err != nil {
		t.Fatal(err)
	}
	command := exec.Command("true")
	if err := command.Run(); err != nil {
		t.Fatal(err)
	}
	deadPID := command.Process.Pid
	lockPath := filepath.Join(store.runDir(runID), ".lock")
	if err := os.WriteFile(lockPath, []byte(strconv.Itoa(deadPID)+"\n"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := store.Mark(token); err != nil {
		t.Fatalf("stale lock was not recovered: %v", err)
	}
	if _, err := os.Stat(lockPath); !os.IsNotExist(err) {
		t.Fatalf("lock was not released: %v", err)
	}
	events, err := store.Events(runID)
	if err != nil {
		t.Fatal(err)
	}
	if len(events) != 1 || events[0].Skill != "prompt-writing" {
		t.Fatalf("unexpected recorded events: %#v", events)
	}
}

func TestLockIsHeldAndReleasedAroundActions(t *testing.T) {
	token := strings.Repeat("cd", 32)
	store := NewStore(t.TempDir())
	runID := "lock-roundtrip-run"
	if err := store.Create(testRun(runID, token)); err != nil {
		t.Fatal(err)
	}
	lockPath := filepath.Join(store.runDir(runID), ".lock")
	err := store.withLock(runID, func() error {
		data, readErr := os.ReadFile(lockPath)
		if readErr != nil {
			return readErr
		}
		pid, convErr := strconv.Atoi(strings.TrimSpace(string(data)))
		if convErr != nil {
			return convErr
		}
		if pid != os.Getpid() {
			t.Fatalf("lock holder PID = %d, want %d", pid, os.Getpid())
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	if _, statErr := os.Stat(lockPath); !os.IsNotExist(statErr) {
		t.Fatalf("lock was not released after action: %v", statErr)
	}
}
