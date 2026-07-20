package evaluation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestPrepareOpenCodeRuntimeOwnsConfigurationAndSkillPermissions(t *testing.T) {
	root := t.TempDir()
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_DATA_HOME", filepath.Join(home, "opencode-data"))

	runtime, err := prepareOpenCodeRuntime(
		root,
		filepath.Join(root, "workspace"),
		"/bin/sh",
		"openai/gpt-5.6-sol",
		"medium",
		"/tmp/skill-issue",
		[]string{"prompt-writing", "document-update-discipline"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if runtime.evaluationSkillRoot != filepath.Join(root, "config", "opencode", "skills") {
		t.Fatalf("unexpected OpenCode skill root: %s", runtime.evaluationSkillRoot)
	}

	data, err := os.ReadFile(filepath.Join(root, "config", "opencode", "opencode.json"))
	if err != nil {
		t.Fatal(err)
	}
	var config map[string]any
	if err := json.Unmarshal(data, &config); err != nil {
		t.Fatal(err)
	}
	permission := config["permission"].(map[string]any)
	skills := permission["skill"].(map[string]any)
	bash := permission["bash"].(map[string]any)
	if skills["*"] != "deny" || skills["prompt-writing"] != "allow" || skills["document-update-discipline"] != "allow" {
		t.Fatalf("unexpected OpenCode skill permissions: %#v", skills)
	}
	if bash["*"] != "deny" || bash["/tmp/skill-issue signal *"] != "allow" || bash[`*/tmp/skill-issue" signal *`] != "allow" {
		t.Fatalf("unexpected OpenCode Bash permissions: %#v", bash)
	}

	joined := strings.Join(runtime.environment, "\n")
	for _, expected := range []string{
		"HOME=" + home,
		"XDG_CONFIG_HOME=" + filepath.Join(root, "config"),
		"XDG_DATA_HOME=" + filepath.Join(home, "opencode-data"),
		"OPENCODE_DISABLE_PROJECT_CONFIG=true",
		"OPENCODE_DISABLE_EXTERNAL_SKILLS=true",
		"OPENCODE_EXPERIMENTAL_DISABLE_FILEWATCHER=true",
	} {
		if !strings.Contains(joined, expected) {
			t.Fatalf("OpenCode environment lacks %q: %v", expected, runtime.environment)
		}
	}
}

func TestPrepareKiloRuntimeOwnsConfigurationAndSkillPermissions(t *testing.T) {
	root := t.TempDir()
	home := t.TempDir()
	dataHome := filepath.Join(home, "kilo-data")
	t.Setenv("HOME", home)
	t.Setenv("XDG_DATA_HOME", dataHome)

	runtime, err := prepareKiloRuntime(
		root,
		filepath.Join(root, "workspace"),
		"/bin/sh",
		"openai/gpt-5.6-sol",
		"medium",
		"/tmp/skill-issue",
		[]string{"prompt-writing", "document-update-discipline"},
	)
	if err != nil {
		t.Fatal(err)
	}
	if runtime.evaluationSkillRoot != filepath.Join(root, "passed-skills") {
		t.Fatalf("unexpected Kilo skill root: %s", runtime.evaluationSkillRoot)
	}

	data, err := os.ReadFile(filepath.Join(root, "config", "kilo", "kilo.json"))
	if err != nil {
		t.Fatal(err)
	}
	var config map[string]any
	if err := json.Unmarshal(data, &config); err != nil {
		t.Fatal(err)
	}
	if config["default_agent"] != "code" || config["remote_control"] != false {
		t.Fatalf("unexpected Kilo runtime settings: %#v", config)
	}
	permission := config["permission"].(map[string]any)
	skills := permission["skill"].(map[string]any)
	bash := permission["bash"].(map[string]any)
	if skills["*"] != "deny" || skills["prompt-writing"] != "allow" || skills["document-update-discipline"] != "allow" {
		t.Fatalf("unexpected Kilo skill permissions: %#v", skills)
	}
	if permission["semantic_search"] != "deny" || permission["kilo_memory_save"] != "deny" {
		t.Fatalf("unexpected Kilo extended permissions: %#v", permission)
	}
	if bash["*"] != "deny" || bash["/tmp/skill-issue signal *"] != "allow" {
		t.Fatalf("unexpected Kilo Bash permissions: %#v", bash)
	}
	agent := config["agent"].(map[string]any)["code"].(map[string]any)
	agentPermission := agent["permission"].(map[string]any)
	if agentPermission["semantic_search"] != "deny" || agentPermission["codebase_search"] != "deny" {
		t.Fatalf("unexpected Kilo agent permissions: %#v", agentPermission)
	}

	joined := strings.Join(runtime.environment, "\n")
	for _, expected := range []string{
		"HOME=" + home,
		"XDG_CONFIG_HOME=" + filepath.Join(root, "config"),
		"XDG_DATA_HOME=" + dataHome,
		"KILO_DISABLE_PROJECT_CONFIG=true",
		"KILO_DISABLE_EXTERNAL_SKILLS=true",
		"KILO_DISABLE_CODEBASE_INDEXING=true",
		"KILO_NO_DAEMON=true",
		"KILO_PURE=true",
	} {
		if !strings.Contains(joined, expected) {
			t.Fatalf("Kilo environment lacks %q: %v", expected, runtime.environment)
		}
	}
}
