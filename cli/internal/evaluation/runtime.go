package evaluation

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
)

type runtimePreparation struct {
	environment           []string
	workingDirectory      string
	evaluationSkillRoot   string
	codexConfiguration    []string
	cursorPluginDir       string
	claudeSettings        string
	claudeSkillsRoot      string
	claudeWorkspacePrompt string
	piSkillsRoot          string
	signalExecutable      string
}

func (service Service) prepareRuntime(harnessID harness.ID, model, reasoning, runID, workspace, executable, cliPath string, skillNames []string) (runtimePreparation, error) {
	if harnessID == harness.Codex {
		skills, err := codexSkillsToDisable()
		if err != nil {
			return runtimePreparation{}, err
		}
		configuration := []string{
			`approvals_reviewer="auto_review"`,
			fmt.Sprintf("model_reasoning_effort=%s", strconv.Quote(reasoning)),
			`project_doc_max_bytes=0`,
			`apps._default.enabled=false`,
		}
		if len(skills) > 0 {
			configuration = append(configuration, codexSkillConfiguration(skills))
		}
		return runtimePreparation{workingDirectory: workspace, evaluationSkillRoot: filepath.Join(workspace, ".agents", "skills"), codexConfiguration: configuration}, nil
	}

	root := privateRuntimeRunRoot(runID)
	if err := os.MkdirAll(root, 0o700); err != nil {
		return runtimePreparation{}, fmt.Errorf("create private harness runtime: %w", err)
	}
	if err := os.MkdirAll(filepath.Join(root, "tmp"), 0o700); err != nil {
		return runtimePreparation{}, fmt.Errorf("create private harness temporary directory: %w", err)
	}

	switch harnessID {
	case harness.Cursor:
		return prepareCursorRuntime(root, workspace, service.stateRoot, executable, cliPath)
	case harness.ClaudeCode:
		return prepareClaudeRuntime(root, workspace)
	case harness.OpenCode:
		return prepareOpenCodeRuntime(root, workspace, executable, model, reasoning, cliPath, skillNames)
	case harness.KiloCode:
		return prepareKiloRuntime(root, workspace, executable, model, reasoning, cliPath, skillNames)
	case harness.Pi:
		return preparePiRuntime(root, workspace, executable)
	default:
		return runtimePreparation{}, fmt.Errorf("unsupported evaluation harness %q", harnessID)
	}
}

// structuredRuntimeSpec parameterizes the OpenCode-style harnesses (OpenCode
// and its Kilo fork) that share a configuration schema, permission model, and
// XDG-based environment isolation. Fields hold only what genuinely differs.
type structuredRuntimeSpec struct {
	harnessID        harness.ID
	name             string            // harness name used in error messages
	configDir        string            // configuration directory name under <root>/config
	configFile       string            // configuration file name
	schema           string            // configuration $schema URL
	defaultAgent     string            // agent that receives the reasoning variant
	agentConfig      map[string]any    // extra settings for the default agent
	configExtras     map[string]any    // harness-specific top-level configuration
	permissionExtras map[string]string // harness-specific permission denials
	environment      []string          // harness-specific environment variables
}

func kiloRuntimeSpec(skillRoot string) structuredRuntimeSpec {
	return structuredRuntimeSpec{
		harnessID:    harness.KiloCode,
		name:         "Kilo",
		configDir:    "kilo",
		configFile:   "kilo.json",
		schema:       "https://app.kilo.ai/config.json",
		defaultAgent: "code",
		agentConfig: map[string]any{
			"permission": map[string]string{
				"codebase_search": "deny",
				"semantic_search": "deny",
			},
		},
		configExtras: map[string]any{
			"remote_control": false,
			"indexing":       map[string]any{"enabled": false},
			"skills":         map[string]any{"paths": []string{skillRoot}, "urls": []string{}},
		},
		permissionExtras: map[string]string{
			"semantic_search":    "deny",
			"codebase_search":    "deny",
			"kilo_memory_save":   "deny",
			"kilo_memory_recall": "deny",
		},
		environment: []string{
			"KILO_AUTO_SHARE=false",
			"KILO_DISABLE_AUTOUPDATE=true",
			"KILO_DISABLE_CLAUDE_CODE=true",
			"KILO_DISABLE_CLAUDE_CODE_PROMPT=true",
			"KILO_DISABLE_CLAUDE_CODE_SKILLS=true",
			"KILO_DISABLE_CODEBASE_INDEXING=true",
			"KILO_DISABLE_EXTERNAL_SKILLS=true",
			"KILO_DISABLE_LSP_DOWNLOAD=true",
			"KILO_DISABLE_PRESENCE=true",
			"KILO_DISABLE_PROJECT_CONFIG=true",
			"KILO_DISABLE_SESSION_INGEST=true",
			"KILO_DISABLE_SHARE=true",
			"KILO_EXPERIMENTAL_DISABLE_FILEWATCHER=true",
			"KILO_NO_DAEMON=true",
			"KILO_PURE=true",
			"KILO_REMOTE=false",
		},
	}
}

func openCodeRuntimeSpec() structuredRuntimeSpec {
	return structuredRuntimeSpec{
		harnessID:    harness.OpenCode,
		name:         "OpenCode",
		configDir:    "opencode",
		configFile:   "opencode.json",
		schema:       "https://opencode.ai/config.json",
		defaultAgent: "build",
		configExtras: map[string]any{
			"plugin":       []string{},
			"instructions": []string{},
			"mcp":          map[string]any{},
		},
		environment: []string{
			"OPENCODE_AUTO_SHARE=false",
			"OPENCODE_DISABLE_AUTOUPDATE=true",
			"OPENCODE_DISABLE_CLAUDE_CODE=true",
			"OPENCODE_DISABLE_CLAUDE_CODE_PROMPT=true",
			"OPENCODE_DISABLE_CLAUDE_CODE_SKILLS=true",
			"OPENCODE_DISABLE_EXTERNAL_SKILLS=true",
			"OPENCODE_DISABLE_LSP_DOWNLOAD=true",
			"OPENCODE_DISABLE_PROJECT_CONFIG=true",
			"OPENCODE_DISABLE_SHARE=true",
			"OPENCODE_EXPERIMENTAL_DISABLE_FILEWATCHER=true",
		},
	}
}

func prepareKiloRuntime(root, workspace, executable, model, reasoning, cliPath string, skillNames []string) (runtimePreparation, error) {
	skillRoot := filepath.Join(root, "passed-skills")
	return prepareStructuredRuntime(kiloRuntimeSpec(skillRoot), root, workspace, executable, model, reasoning, cliPath, skillNames, skillRoot, []string{skillRoot})
}

func prepareOpenCodeRuntime(root, workspace, executable, model, reasoning, cliPath string, skillNames []string) (runtimePreparation, error) {
	skillRoot := filepath.Join(root, "config", "opencode", "skills")
	return prepareStructuredRuntime(openCodeRuntimeSpec(), root, workspace, executable, model, reasoning, cliPath, skillNames, skillRoot, nil)
}

func prepareStructuredRuntime(spec structuredRuntimeSpec, root, workspace, executable, model, reasoning, cliPath string, skillNames []string, skillRoot string, extraPaths []string) (runtimePreparation, error) {
	configRoot := filepath.Join(root, "config", spec.configDir)
	paths := append([]string{configRoot, filepath.Join(root, "state"), filepath.Join(root, "cache")}, extraPaths...)
	for _, path := range paths {
		if err := os.MkdirAll(path, 0o700); err != nil {
			return runtimePreparation{}, fmt.Errorf("create %s runtime path: %w", spec.name, err)
		}
	}
	provider, err := structuredModelProvider(spec.name, model)
	if err != nil {
		return runtimePreparation{}, err
	}
	config := structuredRuntimeConfig(spec, model, provider, reasoning, cliPath, skillNames)
	if err := writeRuntimeJSON(filepath.Join(configRoot, spec.configFile), config); err != nil {
		return runtimePreparation{}, err
	}
	environment, err := structuredEnvironment(spec, root, executable)
	if err != nil {
		return runtimePreparation{}, err
	}
	return runtimePreparation{
		environment:         environment,
		workingDirectory:    workspace,
		evaluationSkillRoot: skillRoot,
	}, nil
}

func structuredModelProvider(name, model string) (string, error) {
	provider, _, found := strings.Cut(model, "/")
	if !found || provider == "" {
		return "", fmt.Errorf("%s model must use provider/model format", name)
	}
	return provider, nil
}

func structuredRuntimeConfig(spec structuredRuntimeSpec, model, provider, reasoning, cliPath string, skillNames []string) map[string]any {
	skillPermissions := map[string]string{"*": "deny"}
	for _, name := range skillNames {
		skillPermissions[name] = "allow"
	}
	permission := map[string]any{
		"*": "deny",
		"read": map[string]string{
			"*":             "allow",
			"*.env":         "deny",
			"*.env.*":       "deny",
			"*.env.example": "allow",
		},
		"edit":  "allow",
		"glob":  "allow",
		"grep":  "allow",
		"list":  "allow",
		"skill": skillPermissions,
		"bash": map[string]string{
			"*":                          "deny",
			"*" + cliPath + `" signal *`: "allow",
			cliPath + " signal *":        "allow",
		},
		"external_directory": "deny",
		"question":           "deny",
		"task":               "deny",
		"webfetch":           "deny",
		"websearch":          "deny",
	}
	for key, value := range spec.permissionExtras {
		permission[key] = value
	}
	agent := map[string]any{"variant": reasoning}
	for key, value := range spec.agentConfig {
		agent[key] = value
	}
	config := map[string]any{
		"$schema":           spec.schema,
		"model":             model,
		"small_model":       model,
		"default_agent":     spec.defaultAgent,
		"enabled_providers": []string{provider},
		"share":             "disabled",
		"autoupdate":        false,
		"snapshot":          false,
		"formatter":         false,
		"lsp":               false,
		"agent":             map[string]any{spec.defaultAgent: agent},
		"permission":        permission,
	}
	for key, value := range spec.configExtras {
		config[key] = value
	}
	return config
}

func structuredEnvironment(spec structuredRuntimeSpec, root, executable string) ([]string, error) {
	path, err := runtimeExecutable(spec.harnessID, executable)
	if err != nil {
		return nil, err
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("resolve %s home: %w", spec.name, err)
	}
	dataHome := os.Getenv("XDG_DATA_HOME")
	if dataHome == "" {
		dataHome = filepath.Join(home, ".local", "share")
	}
	environment := controlledEnvironment(home, filepath.Join(root, "tmp"), path, false)
	environment = append(environment,
		"XDG_CONFIG_HOME="+filepath.Join(root, "config"),
		"XDG_DATA_HOME="+dataHome,
		"XDG_STATE_HOME="+filepath.Join(root, "state"),
		"XDG_CACHE_HOME="+filepath.Join(root, "cache"),
	)
	return append(environment, spec.environment...), nil
}

func kiloEnvironment(root, executable string) ([]string, error) {
	return structuredEnvironment(kiloRuntimeSpec(filepath.Join(root, "passed-skills")), root, executable)
}

func openCodeEnvironment(root, executable string) ([]string, error) {
	return structuredEnvironment(openCodeRuntimeSpec(), root, executable)
}

func prepareCursorRuntime(root, workspace, stateRoot, executable, cliPath string) (runtimePreparation, error) {
	home := filepath.Join(root, "home")
	plugin := filepath.Join(root, "plugin")
	bin := filepath.Join(root, "bin")
	for _, path := range []string{filepath.Join(home, ".cursor"), filepath.Join(home, "Library"), filepath.Join(root, "store"), filepath.Join(plugin, ".cursor-plugin"), filepath.Join(plugin, "skills"), bin} {
		if err := os.MkdirAll(path, 0o700); err != nil {
			return runtimePreparation{}, fmt.Errorf("create Cursor runtime path: %w", err)
		}
	}
	userHome, err := os.UserHomeDir()
	if err != nil {
		return runtimePreparation{}, fmt.Errorf("resolve user home: %w", err)
	}
	keychain := filepath.Join(home, "Library", "Keychains")
	if err := os.Symlink(filepath.Join(userHome, "Library", "Keychains"), keychain); err != nil && !os.IsExist(err) {
		return runtimePreparation{}, fmt.Errorf("link Cursor keychain: %w", err)
	}
	signalExecutable := filepath.Join(bin, "skill-issue")
	if err := copyExecutable(cliPath, signalExecutable); err != nil {
		return runtimePreparation{}, fmt.Errorf("copy Cursor signal executable: %w", err)
	}
	config := map[string]any{
		"version": 1,
		"permissions": map[string]any{
			"allow": []string{"Read(**)", "Write(**)", "Shell(" + signalExecutable + ")"},
			"deny":  []string{"Shell(rm)", "Shell(git)", "Read(.env*)", "Write(**/*.key)"},
		},
		"approvalMode":        "allowlist",
		"notifications":       false,
		"hints":               false,
		"modelSlashCommands":  false,
		"autoAcceptWebSearch": false,
		"sandbox":             map[string]any{"mode": "enabled", "networkAccess": "user_config_with_defaults"},
	}
	if err := writeRuntimeJSON(filepath.Join(home, ".cursor", "cli-config.json"), config); err != nil {
		return runtimePreparation{}, err
	}
	sandbox := map[string]any{
		"type":                     "workspace_readwrite",
		"additionalReadwritePaths": []string{stateRoot},
	}
	if err := writeRuntimeJSON(filepath.Join(home, ".cursor", "sandbox.json"), sandbox); err != nil {
		return runtimePreparation{}, err
	}
	manifest := map[string]any{
		"$schema":     "https://cursor.com/schemas/cursor-plugin/plugin.json",
		"name":        "skill-issue-evaluation-runtime",
		"version":     "1.0.0",
		"description": "Generated Skill Issue evaluation skills.",
		"skills":      "skills",
	}
	if err := writeRuntimeJSON(filepath.Join(plugin, ".cursor-plugin", "plugin.json"), manifest); err != nil {
		return runtimePreparation{}, err
	}
	path, err := runtimeExecutable(harness.Cursor, executable)
	if err != nil {
		return runtimePreparation{}, err
	}
	environment := controlledEnvironment(home, filepath.Join(root, "tmp"), path, false)
	environment = append(environment,
		"CURSOR_CONFIG_DIR="+filepath.Join(home, ".cursor"),
		"CURSOR_DATA_DIR="+filepath.Join(home, ".cursor"),
		"CURSOR_AGENT_STORE_DIR="+filepath.Join(root, "store"),
	)
	return runtimePreparation{
		environment:         environment,
		workingDirectory:    workspace,
		evaluationSkillRoot: filepath.Join(plugin, "skills"),
		cursorPluginDir:     plugin,
		signalExecutable:    signalExecutable,
	}, nil
}

func copyExecutable(source, destination string) error {
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()
	output, err := os.OpenFile(destination, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0o700)
	if err != nil {
		return err
	}
	if _, err := io.Copy(output, input); err != nil {
		output.Close()
		return err
	}
	return output.Close()
}

func prepareClaudeRuntime(root, workspace string) (runtimePreparation, error) {
	launch := filepath.Join(root, "launch")
	skills := filepath.Join(root, "passed-skills", ".claude", "skills")
	if err := os.MkdirAll(launch, 0o700); err != nil {
		return runtimePreparation{}, fmt.Errorf("create Claude launch directory: %w", err)
	}
	if err := os.MkdirAll(skills, 0o700); err != nil {
		return runtimePreparation{}, fmt.Errorf("create Claude skills directory: %w", err)
	}
	settings := map[string]any{
		"autoMemoryEnabled": false,
		"env": map[string]string{
			"CLAUDE_CODE_DISABLE_AUTO_MEMORY":                      "1",
			"CLAUDE_CODE_DISABLE_CLAUDE_MDS":                       "1",
			"CLAUDE_CODE_DISABLE_NONESSENTIAL_TRAFFIC":             "1",
			"CLAUDE_CODE_DISABLE_OFFICIAL_MARKETPLACE_AUTOINSTALL": "1",
			"CLAUDE_CODE_DISABLE_BACKGROUND_TASKS":                 "1",
			"CLAUDE_CODE_DISABLE_CRON":                             "1",
		},
		"permissions": map[string]any{"additionalDirectories": []string{workspace}},
	}
	data, err := json.Marshal(settings)
	if err != nil {
		return runtimePreparation{}, fmt.Errorf("encode Claude settings: %w", err)
	}
	return runtimePreparation{
		workingDirectory:      launch,
		evaluationSkillRoot:   skills,
		claudeSettings:        string(data),
		claudeSkillsRoot:      filepath.Join(root, "passed-skills"),
		claudeWorkspacePrompt: fmt.Sprintf("The evaluation workspace is %s. Treat that absolute directory as the project root. Resolve every relative project path against it. Do not write project files in the process working directory.", workspace),
	}, nil
}

func preparePiRuntime(root, workspace, executable string) (runtimePreparation, error) {
	for _, path := range []string{
		filepath.Join(root, "home"), filepath.Join(root, "sessions"),
		filepath.Join(root, "passed-skills"),
	} {
		if err := os.MkdirAll(path, 0o700); err != nil {
			return runtimePreparation{}, fmt.Errorf("create Pi runtime path: %w", err)
		}
	}
	path, err := runtimeExecutable(harness.Pi, executable)
	if err != nil {
		return runtimePreparation{}, err
	}
	environment := controlledEnvironment(filepath.Join(root, "home"), filepath.Join(root, "tmp"), path, true)
	if voltaHome := os.Getenv("VOLTA_HOME"); voltaHome != "" {
		environment = appendEnvironmentPath(environment, filepath.Join(voltaHome, "bin"))
		environment = append(environment, "VOLTA_HOME="+voltaHome)
	}
	piAgentDirectory := os.Getenv("PI_CODING_AGENT_DIR")
	if piAgentDirectory == "" {
		userHome, err := os.UserHomeDir()
		if err != nil {
			return runtimePreparation{}, fmt.Errorf("resolve Pi agent directory: %w", err)
		}
		piAgentDirectory = filepath.Join(userHome, ".pi", "agent")
	}
	environment = append(environment,
		"PI_CODING_AGENT_DIR="+piAgentDirectory,
		"PI_CODING_AGENT_SESSION_DIR="+filepath.Join(root, "sessions"),
		"PI_OFFLINE=1",
	)
	return runtimePreparation{
		environment:         environment,
		workingDirectory:    workspace,
		evaluationSkillRoot: filepath.Join(root, "passed-skills"),
		piSkillsRoot:        filepath.Join(root, "passed-skills"),
	}, nil
}

func appendEnvironmentPath(environment []string, entry string) []string {
	for index, value := range environment {
		if strings.HasPrefix(value, "PATH=") {
			environment[index] = value + ":" + entry
			return environment
		}
	}
	return environment
}

func writeRuntimeJSON(path string, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("encode runtime configuration: %w", err)
	}
	if err := os.WriteFile(path, append(data, '\n'), 0o600); err != nil {
		return fmt.Errorf("write runtime configuration: %w", err)
	}
	return nil
}

func runtimeExecutable(harnessID harness.ID, override string) (string, error) {
	names := []string{override}
	if override == "" {
		switch harnessID {
		case harness.Cursor:
			names = []string{"agent", "cursor-agent"}
		case harness.ClaudeCode:
			names = []string{"claude"}
		case harness.OpenCode:
			names = []string{"opencode"}
		case harness.KiloCode:
			names = []string{"kilo"}
		case harness.Pi:
			names = []string{"pi"}
		}
	}
	for _, name := range names {
		if name == "" {
			continue
		}
		path, err := exec.LookPath(name)
		if err == nil {
			return path, nil
		}
	}
	return "", fmt.Errorf("%s executable: command not found", harnessID)
}

func controlledEnvironment(home, temporary, executable string, includeCredentials bool) []string {
	user := os.Getenv("USER")
	if user == "" {
		user = "unknown"
	}
	logname := os.Getenv("LOGNAME")
	if logname == "" {
		logname = user
	}
	term := os.Getenv("TERM")
	if term == "" {
		term = "dumb"
	}
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = "en_US.UTF-8"
	}
	environment := []string{
		"HOME=" + home,
		"TMPDIR=" + temporary,
		"PATH=/usr/bin:/bin:/usr/sbin:/sbin:" + filepath.Dir(executable),
		"SHELL=/bin/zsh",
		"TERM=" + term,
		"LANG=" + lang,
		"USER=" + user,
		"LOGNAME=" + logname,
	}
	if !includeCredentials {
		return environment
	}
	for _, entry := range os.Environ() {
		key, _, found := strings.Cut(entry, "=")
		if !found || !credentialEnvironmentKey(key) {
			continue
		}
		environment = append(environment, entry)
	}
	return environment
}

func credentialEnvironmentKey(key string) bool {
	if strings.HasSuffix(key, "_API_KEY") || strings.HasSuffix(key, "_TOKEN") {
		return true
	}
	for _, prefix := range []string{"AWS_", "AZURE_", "CLOUDFLARE_", "OPENAI_", "GOOGLE_", "GEMINI_", "XAI_"} {
		if strings.HasPrefix(key, prefix) {
			return true
		}
	}
	return false
}

func privateRuntimeRunRoot(runID string) string {
	return filepath.Join(os.TempDir(), "skill-issue", runID)
}

func codexSkillsToDisable() ([]string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("resolve user home: %w", err)
	}
	codexHome := os.Getenv("CODEX_HOME")
	if codexHome == "" {
		codexHome = filepath.Join(home, ".codex")
	}
	roots := []string{
		filepath.Join(codexHome, "skills"),
		filepath.Join(home, ".agents", "skills"),
		filepath.FromSlash("/etc/codex/skills"),
	}
	unique := map[string]bool{}
	for _, root := range roots {
		files, err := skillFilesUnder(root)
		if err != nil {
			return nil, err
		}
		for _, file := range files {
			unique[file] = true
		}
	}
	skills := make([]string, 0, len(unique))
	for file := range unique {
		skills = append(skills, file)
	}
	sort.Strings(skills)
	return skills, nil
}

func skillFilesUnder(root string) ([]string, error) {
	resolved, err := filepath.EvalSymlinks(root)
	if os.IsNotExist(err) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("resolve Codex skill root %s: %w", root, err)
	}
	files := []string{}
	err = filepath.WalkDir(resolved, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || entry.Name() != "SKILL.md" {
			return nil
		}
		relative, err := filepath.Rel(resolved, path)
		if err != nil {
			return err
		}
		files = append(files, filepath.Join(root, relative))
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("scan Codex skill root %s: %w", root, err)
	}
	return files, nil
}

func codexSkillConfiguration(skills []string) string {
	entries := make([]string, 0, len(skills))
	for _, skill := range skills {
		entries = append(entries, fmt.Sprintf("{path=%s,enabled=false}", strconv.Quote(skill)))
	}
	return "skills.config=[" + strings.Join(entries, ",") + "]"
}
