package harness

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type ID string

const (
	Copilot    ID = "copilot"
	ClaudeCode ID = "claude-code"
	Codex      ID = "codex"
	Cursor     ID = "cursor"
	GeminiCLI  ID = "gemini-cli"
	GrokBuild  ID = "grok-build"
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
	ID              ID
	Executable      string
	ProjectSkillDir string
	UserSkillDir    func(home string) string
}

var specs = map[ID]Spec{
	Copilot:    {ID: Copilot, Executable: "copilot", ProjectSkillDir: ".github/skills", UserSkillDir: homePath(".copilot", "skills")},
	ClaudeCode: {ID: ClaudeCode, Executable: "claude", ProjectSkillDir: ".claude/skills", UserSkillDir: claudeUserPath},
	Codex:      {ID: Codex, Executable: "codex", ProjectSkillDir: ".agents/skills", UserSkillDir: homePath(".agents", "skills")},
	Cursor:     {ID: Cursor, Executable: "cursor-agent", ProjectSkillDir: ".cursor/skills", UserSkillDir: homePath(".cursor", "skills")},
	GeminiCLI:  {ID: GeminiCLI, Executable: "gemini", ProjectSkillDir: ".gemini/skills", UserSkillDir: homePath(".gemini", "skills")},
	GrokBuild:  {ID: GrokBuild, Executable: "grok", ProjectSkillDir: ".grok/skills", UserSkillDir: homePath(".grok", "skills")},
	OpenCode:   {ID: OpenCode, Executable: "opencode", ProjectSkillDir: ".opencode/skills", UserSkillDir: openCodeUserPath},
	KiloCode:   {ID: KiloCode, Executable: "kilo", ProjectSkillDir: ".kilo/skills", UserSkillDir: homePath(".kilo", "skills")},
	Pi:         {ID: Pi, Executable: "pi", ProjectSkillDir: ".pi/skills", UserSkillDir: homePath(".pi", "agent", "skills")},
}

func ParseID(value string) (ID, error) {
	id := ID(value)
	if _, ok := specs[id]; !ok {
		return "", fmt.Errorf("unsupported harness %q", value)
	}
	return id, nil
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

func IDs() []ID {
	ids := make([]ID, 0, len(specs))
	for id := range specs {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool { return ids[i] < ids[j] })
	return ids
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

func openCodeUserPath(home string) string {
	if configured := os.Getenv("XDG_CONFIG_HOME"); configured != "" {
		return filepath.Join(configured, "opencode", "skills")
	}
	return filepath.Join(home, ".config", "opencode", "skills")
}
