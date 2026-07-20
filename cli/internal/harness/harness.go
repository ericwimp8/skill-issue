package harness

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

type ID string

const (
	ClaudeCode ID = "claude-code"
	Codex      ID = "codex"
	Cursor     ID = "cursor"
	OpenCode   ID = "opencode"
	KiloCode   ID = "kilo-code"
	Pi         ID = "pi"
)

type Scope string

const (
	ScopeProject Scope = "project"
	ScopeUser    Scope = "user"
)

type Spec struct {
	ID                     ID
	Executable             string
	ProjectSkillDir        string
	UserSkillDir           func(home string) string
	HarnessSkillFiles      []string
	DisableModelInvocation bool
}

type EvaluationDefaults struct {
	Model     string
	Reasoning string
}

var specs = map[ID]Spec{
	ClaudeCode: {ID: ClaudeCode, Executable: "claude", ProjectSkillDir: ".claude/skills", UserSkillDir: claudeUserPath, DisableModelInvocation: true},
	Codex:      {ID: Codex, Executable: "codex", ProjectSkillDir: ".agents/skills", UserSkillDir: homePath(".agents", "skills"), HarnessSkillFiles: []string{"agents/openai.yaml"}},
	Cursor:     {ID: Cursor, Executable: "cursor-agent", ProjectSkillDir: ".cursor/skills", UserSkillDir: homePath(".cursor", "skills"), DisableModelInvocation: true},
	OpenCode:   {ID: OpenCode, Executable: "opencode", ProjectSkillDir: ".opencode/skills", UserSkillDir: homePath(".config", "opencode", "skills")},
	KiloCode:   {ID: KiloCode, Executable: "kilo", ProjectSkillDir: ".kilo/skills", UserSkillDir: homePath(".kilo", "skills")},
	Pi:         {ID: Pi, Executable: "pi", ProjectSkillDir: ".pi/skills", UserSkillDir: homePath(".pi", "agent", "skills"), DisableModelInvocation: true},
}

var evaluationDefaults = map[ID]EvaluationDefaults{
	ClaudeCode: {Model: "opus", Reasoning: "medium"},
	Codex:      {Model: "gpt-5.6-sol", Reasoning: "medium"},
	Cursor:     {Model: "auto", Reasoning: "medium"},
	OpenCode:   {Model: "openai/gpt-5.6-sol", Reasoning: "medium"},
	KiloCode:   {Model: "openai/gpt-5.6-sol", Reasoning: "medium"},
	Pi:         {Model: "openai-codex/gpt-5.6-sol", Reasoning: "medium"},
}

func SupportedIDs() []ID {
	return []ID{
		ClaudeCode,
		Codex,
		Cursor,
		OpenCode,
		KiloCode,
		Pi,
	}
}

func InstallationAvailable(id ID) bool {
	switch id {
	case ClaudeCode, Codex, Cursor, OpenCode, KiloCode, Pi:
		return true
	default:
		return false
	}
}

func ParseID(value string) (ID, error) {
	id := ID(value)
	if _, ok := specs[id]; !ok {
		return "", fmt.Errorf("unsupported harness %q", value)
	}
	return id, nil
}

func ParseEvaluationID(value string) (ID, error) {
	id, err := ParseID(value)
	if err != nil {
		return "", err
	}
	if _, ok := evaluationDefaults[id]; !ok {
		return "", fmt.Errorf("unsupported evaluation harness %q", value)
	}
	return id, nil
}

func EvaluationDefaultsFor(id ID) (EvaluationDefaults, error) {
	defaults, ok := evaluationDefaults[id]
	if !ok {
		return EvaluationDefaults{}, fmt.Errorf("unsupported evaluation harness %q", id)
	}
	return defaults, nil
}

func ParseScope(value string) (Scope, error) {
	scope := Scope(value)
	if scope != ScopeProject && scope != ScopeUser {
		return "", fmt.Errorf("unsupported scope %q", value)
	}
	return scope, nil
}

func Lookup(id ID) (Spec, error) {
	spec, ok := specs[id]
	if !ok {
		return Spec{}, fmt.Errorf("unsupported harness %q", id)
	}
	return spec, nil
}

func SkillRoot(id ID, scope Scope, workspace, home string) (string, error) {
	spec, err := Lookup(id)
	if err != nil {
		return "", err
	}
	if scope == ScopeProject {
		if workspace == "" {
			return "", errors.New("workspace is required for project scope")
		}
		absolute, err := filepath.Abs(workspace)
		if err != nil {
			return "", fmt.Errorf("resolve workspace: %w", err)
		}
		return filepath.Join(absolute, filepath.FromSlash(spec.ProjectSkillDir)), nil
	}
	if scope != ScopeUser {
		return "", fmt.Errorf("unsupported scope %q", scope)
	}
	if home == "" {
		return "", errors.New("home directory is required for user scope")
	}
	return spec.UserSkillDir(home), nil
}

func IncludeSkillFile(id ID, relative string) (bool, error) {
	spec, err := Lookup(id)
	if err != nil {
		return false, err
	}
	relative = path.Clean(relative)
	if contains(spec.HarnessSkillFiles, relative) {
		return true, nil
	}
	for _, candidate := range specs {
		if contains(candidate.HarnessSkillFiles, relative) {
			return false, nil
		}
	}
	return true, nil
}

func SupportsDisableModelInvocation(id ID) (bool, error) {
	spec, err := Lookup(id)
	if err != nil {
		return false, err
	}
	return spec.DisableModelInvocation, nil
}

func contains(values []string, wanted string) bool {
	for _, value := range values {
		if value == wanted {
			return true
		}
	}
	return false
}

func homePath(parts ...string) func(string) string {
	return func(home string) string {
		items := append([]string{home}, parts...)
		return filepath.Join(items...)
	}
}

func claudeUserPath(home string) string {
	if configured := os.Getenv("CLAUDE_CONFIG_DIR"); configured != "" {
		return filepath.Join(configured, "skills")
	}
	return filepath.Join(home, ".claude", "skills")
}
