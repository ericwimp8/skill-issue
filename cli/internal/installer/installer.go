package installer

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/payload"
)

type Request struct {
	Harness         harness.ID
	Scope           harness.Scope
	Workspace       string
	EvaluationRoot  string
	Home            string
	CLIPath         string
	SignalStateRoot string
	Tokens          map[string]string
	Skills          []payload.Skill
}

type Installation struct {
	InstallRoot string   `json:"install_root"`
	Paths       []string `json:"paths"`
}

type EvaluationInstallation struct {
	Root        string   `json:"root,omitempty"`
	Preexisting []string `json:"preexisting"`
	Skills      []string `json:"skills"`
}

type Service struct{}

func New() Service {
	return Service{}
}

func (Service) Install(request Request) (Installation, error) {
	skills, err := payload.Skills()
	if err != nil {
		return Installation{}, err
	}
	root, err := ensureSkillRoot(request)
	if err != nil {
		return Installation{}, err
	}
	return materializeSkills(root, skills)
}

func (Service) Uninstall(request Request) (Installation, error) {
	skills, err := payload.Skills()
	if err != nil {
		return Installation{}, err
	}
	root, err := resolveSkillRoot(request)
	if err != nil {
		return Installation{}, err
	}
	removed := make([]string, 0, len(skills))
	for _, skill := range skills {
		target := filepath.Join(root, skill.Name)
		if err := os.RemoveAll(target); err != nil {
			return Installation{}, fmt.Errorf("remove Skill Issue skill %s: %w", target, err)
		}
		removed = append(removed, target)
	}
	return Installation{InstallRoot: root, Paths: removed}, nil
}

func (Service) PrepareEvaluation(request Request) (EvaluationInstallation, Installation, error) {
	if request.Scope != harness.ScopeProject {
		return EvaluationInstallation{}, Installation{}, errors.New("evaluation installations are always project-local")
	}
	skills, err := selectedEvaluationSkills(request)
	if err != nil {
		return EvaluationInstallation{}, Installation{}, err
	}
	skills, err = instrument(skills, request.CLIPath, request.SignalStateRoot, request.Tokens)
	if err != nil {
		return EvaluationInstallation{}, Installation{}, err
	}
	root, err := ensureSkillRoot(request)
	if err != nil {
		return EvaluationInstallation{}, Installation{}, err
	}
	ordinary, err := payload.Skills()
	if err != nil {
		return EvaluationInstallation{}, Installation{}, err
	}
	ordinaryNames := skillNames(ordinary)
	state := EvaluationInstallation{Root: root}
	state.Skills = skillNamesSorted(skills)
	for _, skill := range skills {
		target := filepath.Join(root, skill.Name)
		if _, statErr := os.Stat(target); statErr == nil {
			if !ordinaryNames[skill.Name] {
				return EvaluationInstallation{}, Installation{}, fmt.Errorf("evaluation-only skill collision at %s", target)
			}
			state.Preexisting = append(state.Preexisting, skill.Name)
		} else if !errors.Is(statErr, os.ErrNotExist) {
			return EvaluationInstallation{}, Installation{}, fmt.Errorf("inspect evaluation skill destination: %w", statErr)
		}
	}
	sort.Strings(state.Preexisting)
	installed, err := materializeSkills(root, skills)
	if err != nil {
		cleanupErr := Service{}.CleanupEvaluation(request, state)
		return EvaluationInstallation{}, Installation{}, errors.Join(err, cleanupErr)
	}
	return state, installed, nil
}

func (Service) CleanupEvaluation(request Request, state EvaluationInstallation) error {
	if request.Scope != harness.ScopeProject {
		return errors.New("evaluation installations are always project-local")
	}
	if state.Root != "" {
		request.EvaluationRoot = state.Root
	}
	root, err := resolveSkillRoot(request)
	if err != nil {
		return err
	}
	evaluationSkillNames, err := evaluationSkillsForState(state)
	if err != nil {
		return err
	}
	ordinarySkills, err := payload.Skills()
	if err != nil {
		return err
	}
	ordinaryByName := make(map[string]payload.Skill, len(ordinarySkills))
	for _, skill := range ordinarySkills {
		ordinaryByName[skill.Name] = skill
	}
	preexisting := make(map[string]bool, len(state.Preexisting))
	for _, name := range state.Preexisting {
		if _, ok := ordinaryByName[name]; !ok {
			return fmt.Errorf("invalid preexisting evaluation skill %q", name)
		}
		preexisting[name] = true
	}
	if len(preexisting) > 0 {
		if err := os.MkdirAll(root, 0o755); err != nil {
			return fmt.Errorf("create skill root: %w", err)
		}
	}
	for _, name := range evaluationSkillNames {
		target := filepath.Join(root, name)
		if preexisting[name] {
			if err := materialize(root, target, ordinaryByName[name]); err != nil {
				return err
			}
			continue
		}
		if err := os.RemoveAll(target); err != nil {
			return fmt.Errorf("remove temporary evaluation skill %s: %w", target, err)
		}
	}
	return nil
}

func EncodeEvaluationInstallation(state EvaluationInstallation) ([]byte, error) {
	return json.MarshalIndent(state, "", "  ")
}

func DecodeEvaluationInstallation(data []byte) (EvaluationInstallation, error) {
	var state EvaluationInstallation
	if err := json.Unmarshal(data, &state); err != nil {
		return EvaluationInstallation{}, err
	}
	if state.Root != "" && !filepath.IsAbs(state.Root) {
		return EvaluationInstallation{}, errors.New("evaluation installation state is invalid")
	}
	for _, name := range state.Preexisting {
		if name == "" || filepath.Base(name) != name {
			return EvaluationInstallation{}, errors.New("evaluation installation state is invalid")
		}
	}
	for _, name := range state.Skills {
		if name == "" || filepath.Base(name) != name {
			return EvaluationInstallation{}, errors.New("evaluation installation state is invalid")
		}
	}
	return state, nil
}

func selectedEvaluationSkills(request Request) ([]payload.Skill, error) {
	if len(request.Skills) > 0 {
		return request.Skills, nil
	}
	return payload.EvaluationSkills()
}

func evaluationSkillsForState(state EvaluationInstallation) ([]string, error) {
	if len(state.Skills) > 0 {
		return state.Skills, nil
	}
	skills, err := payload.EvaluationSkills()
	if err != nil {
		return nil, err
	}
	return skillNamesSorted(skills), nil
}

func skillNamesSorted(skills []payload.Skill) []string {
	names := make([]string, 0, len(skills))
	for _, skill := range skills {
		names = append(names, skill.Name)
	}
	sort.Strings(names)
	return names
}

func resolveSkillRoot(request Request) (string, error) {
	if request.EvaluationRoot != "" {
		if !filepath.IsAbs(request.EvaluationRoot) {
			return "", errors.New("evaluation skill root must be absolute")
		}
		return request.EvaluationRoot, nil
	}
	root, err := harness.SkillRoot(request.Harness, request.Scope, request.Workspace, request.Home)
	if err != nil {
		return "", err
	}
	return root, nil
}

func ensureSkillRoot(request Request) (string, error) {
	root, err := resolveSkillRoot(request)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return "", fmt.Errorf("create skill root: %w", err)
	}
	return root, nil
}

func materializeSkills(root string, skills []payload.Skill) (Installation, error) {
	paths := make([]string, 0, len(skills))
	for _, skill := range skills {
		target := filepath.Join(root, skill.Name)
		if err := materialize(root, target, skill); err != nil {
			return Installation{}, err
		}
		paths = append(paths, target)
	}
	return Installation{InstallRoot: root, Paths: paths}, nil
}

func skillNames(skills []payload.Skill) map[string]bool {
	names := make(map[string]bool, len(skills))
	for _, skill := range skills {
		names[skill.Name] = true
	}
	return names
}

func instrument(skills []payload.Skill, cliPath, signalStateRoot string, tokens map[string]string) ([]payload.Skill, error) {
	if !filepath.IsAbs(cliPath) {
		return nil, errors.New("evaluation CLI path must be absolute")
	}
	if !filepath.IsAbs(signalStateRoot) {
		return nil, errors.New("evaluation signal state root must be absolute")
	}
	result := make([]payload.Skill, 0, len(skills))
	for _, skill := range skills {
		token := ""
		for candidate, name := range tokens {
			if name == skill.Name {
				token = candidate
				break
			}
		}
		if token == "" {
			return nil, fmt.Errorf("skill %q has no opaque token", skill.Name)
		}
		files := make(map[string][]byte, len(skill.Files))
		for name, data := range skill.Files {
			files[name] = append([]byte(nil), data...)
		}
		entrypoint, err := inject(files["SKILL.md"], cliPath, signalStateRoot, token)
		if err != nil {
			return nil, fmt.Errorf("instrument %s: %w", skill.Name, err)
		}
		files["SKILL.md"] = entrypoint
		result = append(result, payload.Skill{Name: skill.Name, Files: files})
	}
	return result, nil
}

func inject(data []byte, cliPath, signalStateRoot, token string) ([]byte, error) {
	text := string(data)
	if !strings.HasPrefix(text, "---\n") {
		return nil, errors.New("SKILL.md has no opening frontmatter delimiter")
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		return nil, errors.New("SKILL.md has no closing frontmatter delimiter")
	}
	insertAt := 4 + end + len("\n---\n")
	instruction := fmt.Sprintf("\nRun %q signal %q %q, then continue normally.\n", cliPath, token, signalStateRoot)
	return []byte(text[:insertAt] + instruction + text[insertAt:]), nil
}

func materialize(root, target string, skill payload.Skill) error {
	staging, err := os.MkdirTemp(root, ".skill-issue-stage-*")
	if err != nil {
		return fmt.Errorf("create skill staging directory: %w", err)
	}
	defer os.RemoveAll(staging)
	for relative, data := range skill.Files {
		path := filepath.Join(staging, filepath.FromSlash(relative))
		if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
			return fmt.Errorf("create staged skill directory: %w", err)
		}
		if err := os.WriteFile(path, data, 0o644); err != nil {
			return fmt.Errorf("write staged skill file: %w", err)
		}
	}
	if err := os.RemoveAll(target); err != nil {
		return fmt.Errorf("remove prior Skill Issue skill: %w", err)
	}
	if err := os.Rename(staging, target); err != nil {
		return fmt.Errorf("commit skill directory: %w", err)
	}
	return nil
}
