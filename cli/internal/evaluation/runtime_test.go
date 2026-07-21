package evaluation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
)

func TestPrepareCodexRuntimeOwnsStateAndDisablesAgents(t *testing.T) {
	defer os.RemoveAll(privateRuntimeRunRoot("run-id"))
	root := t.TempDir()
	workspace := filepath.Join(root, "workspace")
	userCodexHome := filepath.Join(root, "user-codex-home")
	if err := os.MkdirAll(userCodexHome, 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(userCodexHome, "auth.json"), []byte("authentication"), 0o600); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CODEX_HOME", userCodexHome)

	runtime, err := (Service{}).prepareRuntime(harness.Codex, "gpt-5.6-sol", "medium", "run-id", workspace, "/bin/sh", "/tmp/skill-issue", nil)
	if err != nil {
		t.Fatal(err)
	}
	codexHome := filepath.Join(privateRuntimeRunRoot("run-id"), "codex-home")
	if !slices.Contains(runtime.environment, "CODEX_HOME="+codexHome) {
		t.Fatalf("Codex runtime does not own CODEX_HOME: %v", runtime.environment)
	}
	authentication, err := os.ReadFile(filepath.Join(codexHome, "auth.json"))
	if err != nil {
		t.Fatal(err)
	}
	if string(authentication) != "authentication" {
		t.Fatalf("unexpected copied authentication: %q", authentication)
	}
	if !slices.Contains(runtime.codexConfiguration, `features.multi_agent=false`) || !slices.Contains(runtime.codexConfiguration, `features.multi_agent_v2=false`) {
		t.Fatalf("Codex runtime leaves agents enabled: %v", runtime.codexConfiguration)
	}
	if runtime.evaluationSkillRoot != filepath.Join(workspace, ".agents", "skills") {
		t.Fatalf("unexpected Codex skill root: %s", runtime.evaluationSkillRoot)
	}
}

func TestPrepareCursorRuntimeOwnsSignalExecutable(t *testing.T) {
	root := t.TempDir()
	workspace := filepath.Join(root, "workspace")
	stateRoot := filepath.Join(root, "output", ".skill-issue")
	home := t.TempDir()
	t.Setenv("HOME", home)
	source := filepath.Join(root, "source-skill-issue")
	if err := os.WriteFile(source, []byte("binary"), 0o700); err != nil {
		t.Fatal(err)
	}

	runtime, err := prepareCursorRuntime(root, workspace, stateRoot, "/bin/sh", source)
	if err != nil {
		t.Fatal(err)
	}
	if runtime.signalExecutable != filepath.Join(root, "bin", "skill-issue") {
		t.Fatalf("unexpected Cursor signal executable: %s", runtime.signalExecutable)
	}
	data, err := os.ReadFile(runtime.signalExecutable)
	if err != nil {
		t.Fatal(err)
	}
	if string(data) != "binary" {
		t.Fatalf("unexpected copied executable: %q", data)
	}
	configData, err := os.ReadFile(filepath.Join(root, "home", ".cursor", "cli-config.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(configData), "Shell("+runtime.signalExecutable+")") {
		t.Fatalf("Cursor config does not allow the run-owned signal executable: %s", configData)
	}
	sandboxData, err := os.ReadFile(filepath.Join(root, "home", ".cursor", "sandbox.json"))
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(sandboxData), stateRoot) {
		t.Fatalf("Cursor sandbox does not allow the run-owned marker state: %s", sandboxData)
	}
}

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

func TestStructuredEnvironmentLeavesUnselectedDataRootToExecutable(t *testing.T) {
	home := t.TempDir()
	t.Setenv("HOME", home)
	t.Setenv("XDG_DATA_HOME", "")

	environment, err := structuredEnvironment(openCodeRuntimeSpec(), t.TempDir(), "/bin/sh")
	if err != nil {
		t.Fatal(err)
	}
	for _, entry := range environment {
		if strings.HasPrefix(entry, "XDG_DATA_HOME=") {
			t.Fatalf("runtime replaced the executable's native data-root default: %v", environment)
		}
	}
}

func TestControlledEnvironmentForwardsOnlyAllowlistedCredentials(t *testing.T) {
	t.Setenv("OPENAI_API_KEY", "forwarded")
	t.Setenv("MY_PRIVATE_TOKEN", "secret")
	t.Setenv("AWS_VAULT_BACKEND", "secret")
	environment := controlledEnvironment("/home", "/tmpdir", "/usr/bin/pi", true)
	joined := strings.Join(environment, "\n")
	if !strings.Contains(joined, "OPENAI_API_KEY=forwarded") {
		t.Fatalf("allowlisted credential was not forwarded: %v", environment)
	}
	if strings.Contains(joined, "MY_PRIVATE_TOKEN") {
		t.Fatalf("arbitrary *_TOKEN variable was forwarded: %v", environment)
	}
	if strings.Contains(joined, "AWS_VAULT_BACKEND") {
		t.Fatalf("non-allowlisted AWS_-prefixed variable was forwarded: %v", environment)
	}
}
