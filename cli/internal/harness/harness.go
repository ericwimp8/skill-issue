package harness

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"slices"
)

type ID string

const (
	ClaudeCode ID = "claude-code"
	Codex      ID = "codex"
	Cursor     ID = "cursor"
	OpenCode   ID = "opencode"
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
	// InstallationInProgress marks a harness that is defined but whose
	// installation support is not ready for use yet.
	InstallationInProgress bool
	// CleanEvaluationEnvironment runs evaluation harness processes with only
	// the controlled environment instead of overlaying the caller's.
	CleanEvaluationEnvironment bool
	// CleanAuthenticationEnvironment runs the pre-run authentication check
	// with only the controlled environment.
	CleanAuthenticationEnvironment bool
	// TestedVersion is the harness version this CLI was qualified against.
	// When VersionPinned is set the evaluator requires that exact version
	// before side effects; otherwise drift is reported as a warning.
	TestedVersion string
	// VersionPinned marks TestedVersion as an exact requirement rather than
	// a recorded reference.
	VersionPinned bool
	// Evaluation holds the harness's evaluation defaults; a nil value means
	// the harness cannot run evaluations.
	Evaluation *EvaluationDefaults
}

type EvaluationDefaults struct {
	Model     string
	Reasoning string
}

// orderedSpecs is the single registry of supported harnesses; every lookup,
// listing, and capability check derives from it.
var orderedSpecs = []Spec{
	{ID: ClaudeCode, Executable: "claude", ProjectSkillDir: ".claude/skills", UserSkillDir: claudeUserPath, DisableModelInvocation: true, TestedVersion: "2.1.205", Evaluation: &EvaluationDefaults{Model: "opus", Reasoning: "medium"}},
	{ID: Codex, Executable: "codex", ProjectSkillDir: ".agents/skills", UserSkillDir: homePath(".agents", "skills"), HarnessSkillFiles: []string{"agents/openai.yaml"}, TestedVersion: "0.144.6", Evaluation: &EvaluationDefaults{Model: "gpt-5.6-sol", Reasoning: "medium"}},
	{ID: Cursor, Executable: "cursor-agent", ProjectSkillDir: ".cursor/skills", UserSkillDir: homePath(".cursor", "skills"), DisableModelInvocation: true, CleanEvaluationEnvironment: true, CleanAuthenticationEnvironment: true, TestedVersion: "2026.07.16-899851b", Evaluation: &EvaluationDefaults{Model: "auto", Reasoning: "medium"}},
	{ID: OpenCode, Executable: "opencode", ProjectSkillDir: ".opencode/skills", UserSkillDir: homePath(".config", "opencode", "skills"), CleanAuthenticationEnvironment: true, TestedVersion: "1.18.4", VersionPinned: true, Evaluation: &EvaluationDefaults{Model: "openai/gpt-5.6-sol", Reasoning: "medium"}},
	{ID: Pi, Executable: "pi", ProjectSkillDir: ".pi/skills", UserSkillDir: homePath(".pi", "agent", "skills"), DisableModelInvocation: true, CleanEvaluationEnvironment: true, CleanAuthenticationEnvironment: true, TestedVersion: "0.80.10", Evaluation: &EvaluationDefaults{Model: "openai-codex/gpt-5.6-sol", Reasoning: "medium"}},
}

var specs = func() map[ID]Spec {
	byID := make(map[ID]Spec, len(orderedSpecs))
	for _, spec := range orderedSpecs {
		byID[spec.ID] = spec
	}
	return byID
}()

func SupportedIDs() []ID {
	ids := make([]ID, 0, len(orderedSpecs))
	for _, spec := range orderedSpecs {
		ids = append(ids, spec.ID)
	}
	return ids
}

func InstallationAvailable(id ID) bool {
	spec, ok := specs[id]
	return ok && !spec.InstallationInProgress
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
	if specs[id].Evaluation == nil {
		return "", fmt.Errorf("unsupported evaluation harness %q", value)
	}
	return id, nil
}

func EvaluationDefaultsFor(id ID) (EvaluationDefaults, error) {
	spec, ok := specs[id]
	if !ok || spec.Evaluation == nil {
		return EvaluationDefaults{}, fmt.Errorf("unsupported evaluation harness %q", id)
	}
	return *spec.Evaluation, nil
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
	if slices.Contains(spec.HarnessSkillFiles, relative) {
		return true, nil
	}
	for _, candidate := range specs {
		if slices.Contains(candidate.HarnessSkillFiles, relative) {
			return false, nil
		}
	}
	return true, nil
}

// TestedVersion reports the harness version this CLI was qualified against
// and whether that version is an exact requirement.
func TestedVersion(id ID) (string, bool, error) {
	spec, err := Lookup(id)
	if err != nil {
		return "", false, err
	}
	return spec.TestedVersion, spec.VersionPinned, nil
}

func SupportsDisableModelInvocation(id ID) (bool, error) {
	spec, err := Lookup(id)
	if err != nil {
		return false, err
	}
	return spec.DisableModelInvocation, nil
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
