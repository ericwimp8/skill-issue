package installer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/payload"
)

type Request struct {
	Harness              harness.ID
	Scope                harness.Scope
	Workspace            string
	EvaluationRoot       string
	Home                 string
	CLIPath              string
	SignalStateRoot      string
	BackupRoot           string
	Tokens               map[string]string
	Skills               []payload.Skill
	ApplyHarnessMetadata bool
	ConfirmReplace       func(differing []string) (bool, error)
}

type Installation struct {
	InstallRoot string   `json:"install_root"`
	Paths       []string `json:"paths"`
}

type EvaluationInstallation struct {
	Root        string   `json:"root,omitempty"`
	BackupRoot  string   `json:"backup_root,omitempty"`
	Preexisting []string `json:"preexisting"`
	Skills      []string `json:"skills"`
}

var ErrReplacementDeclined = errors.New("declined replacement of preexisting skills")

type Service struct{}

func New() Service {
	return Service{}
}

func (Service) Install(request Request) (Installation, error) {
	skills, err := payload.Skills()
	if err != nil {
		return Installation{}, err
	}
	skills, err = applyHarnessMetadata(request.Harness, skills)
	if err != nil {
		return Installation{}, err
	}
	root, err := ensureSkillRoot(request)
	if err != nil {
		return Installation{}, err
	}
	return materializeSkills(root, request.Harness, skills)
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
	if request.ApplyHarnessMetadata {
		skills, err = applyHarnessMetadata(request.Harness, skills)
		if err != nil {
			return EvaluationInstallation{}, Installation{}, err
		}
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
	canonical, err := applyHarnessMetadata(request.Harness, ordinary)
	if err != nil {
		return EvaluationInstallation{}, Installation{}, err
	}
	canonicalByName := make(map[string]payload.Skill, len(canonical))
	for _, skill := range canonical {
		canonicalByName[skill.Name] = skill
	}
	state := EvaluationInstallation{Root: root}
	state.Skills = skillNamesSorted(skills)
	var differing []string
	for _, skill := range skills {
		target := filepath.Join(root, skill.Name)
		if _, statErr := os.Stat(target); statErr == nil {
			if !ordinaryNames[skill.Name] {
				return EvaluationInstallation{}, Installation{}, fmt.Errorf("temporary evaluation skill collision at %s", target)
			}
			state.Preexisting = append(state.Preexisting, skill.Name)
			matches, matchErr := directoryMatchesSkill(target, request.Harness, canonicalByName[skill.Name])
			if matchErr != nil {
				return EvaluationInstallation{}, Installation{}, fmt.Errorf("inspect preexisting skill %s: %w", target, matchErr)
			}
			if !matches {
				differing = append(differing, skill.Name)
			}
		} else if !errors.Is(statErr, os.ErrNotExist) {
			return EvaluationInstallation{}, Installation{}, fmt.Errorf("inspect evaluation skill destination: %w", statErr)
		}
	}
	sort.Strings(state.Preexisting)
	sort.Strings(differing)
	if len(differing) > 0 {
		if request.ConfirmReplace == nil {
			return EvaluationInstallation{}, Installation{}, fmt.Errorf("preexisting skills differ from their canonical versions: %s", strings.Join(differing, ", "))
		}
		confirmed, confirmErr := request.ConfirmReplace(append([]string(nil), differing...))
		if confirmErr != nil {
			return EvaluationInstallation{}, Installation{}, confirmErr
		}
		if !confirmed {
			return EvaluationInstallation{}, Installation{}, ErrReplacementDeclined
		}
	}
	if len(state.Preexisting) > 0 && request.BackupRoot != "" {
		if !filepath.IsAbs(request.BackupRoot) {
			return EvaluationInstallation{}, Installation{}, errors.New("evaluation skill backup root must be absolute")
		}
		state.BackupRoot = request.BackupRoot
		for _, name := range state.Preexisting {
			if err := copyDirectory(filepath.Join(root, name), filepath.Join(request.BackupRoot, name)); err != nil {
				return EvaluationInstallation{}, Installation{}, fmt.Errorf("back up preexisting skill %s: %w", name, err)
			}
		}
	}
	installed, err := materializeSkills(root, request.Harness, skills)
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
	ordinarySkills, err = applyHarnessMetadata(request.Harness, ordinarySkills)
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
			restored, err := restoreFromBackup(state.BackupRoot, name, target)
			if err != nil {
				return err
			}
			if restored {
				continue
			}
			if err := materialize(root, target, request.Harness, ordinaryByName[name]); err != nil {
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

func restoreFromBackup(backupRoot, name, target string) (bool, error) {
	if backupRoot == "" {
		return false, nil
	}
	backup := filepath.Join(backupRoot, name)
	info, err := os.Stat(backup)
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("inspect preexisting skill backup %s: %w", backup, err)
	}
	if !info.IsDir() {
		return false, fmt.Errorf("preexisting skill backup %s is not a directory", backup)
	}
	if err := os.RemoveAll(target); err != nil {
		return false, fmt.Errorf("remove temporary evaluation skill %s: %w", target, err)
	}
	if err := copyDirectory(backup, target); err != nil {
		return false, fmt.Errorf("restore preexisting skill %s: %w", name, err)
	}
	return true, nil
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
	if state.BackupRoot != "" && !filepath.IsAbs(state.BackupRoot) {
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
	return payload.Skills()
}

func evaluationSkillsForState(state EvaluationInstallation) ([]string, error) {
	if len(state.Skills) > 0 {
		return state.Skills, nil
	}
	skills, err := payload.Skills()
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

func materializeSkills(root string, harnessID harness.ID, skills []payload.Skill) (Installation, error) {
	paths := make([]string, 0, len(skills))
	for _, skill := range skills {
		target := filepath.Join(root, skill.Name)
		if err := materialize(root, target, harnessID, skill); err != nil {
			return Installation{}, err
		}
		paths = append(paths, target)
	}
	if err := verifyMaterializedSkills(root, harnessID, skills); err != nil {
		return Installation{}, err
	}
	return Installation{InstallRoot: root, Paths: paths}, nil
}

func verifyMaterializedSkills(root string, harnessID harness.ID, skills []payload.Skill) error {
	for _, skill := range skills {
		for relative := range skill.Files {
			include, err := harness.IncludeSkillFile(harnessID, relative)
			if err != nil {
				return err
			}
			if !include {
				continue
			}
			installedPath := filepath.Join(root, skill.Name, filepath.FromSlash(relative))
			info, err := os.Stat(installedPath)
			if err != nil {
				return fmt.Errorf("verify installed skill file %s: %w", installedPath, err)
			}
			if !info.Mode().IsRegular() {
				return fmt.Errorf("verify installed skill file %s: not a regular file", installedPath)
			}
		}
	}
	return nil
}

func directoryMatchesSkill(target string, harnessID harness.ID, skill payload.Skill) (bool, error) {
	expected := make(map[string][]byte, len(skill.Files))
	for relative, data := range skill.Files {
		include, err := harness.IncludeSkillFile(harnessID, relative)
		if err != nil {
			return false, err
		}
		if include {
			expected[relative] = data
		}
	}
	actual, err := readDirectoryFiles(target)
	if err != nil {
		return false, err
	}
	if len(actual) != len(expected) {
		return false, nil
	}
	for relative, data := range expected {
		if !bytes.Equal(actual[relative], data) {
			return false, nil
		}
	}
	return true, nil
}

func readDirectoryFiles(root string) (map[string][]byte, error) {
	files := map[string][]byte{}
	err := filepath.WalkDir(root, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.IsDir() || entry.Name() == ".DS_Store" {
			return nil
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("unsupported file type at %s", filePath)
		}
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(root, filePath)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(relative)] = data
		return nil
	})
	if err != nil {
		return nil, err
	}
	return files, nil
}

func copyDirectory(source, destination string) error {
	return filepath.WalkDir(source, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(source, filePath)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, relative)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		if !entry.Type().IsRegular() {
			return fmt.Errorf("unsupported file type at %s", filePath)
		}
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
}

func skillNames(skills []payload.Skill) map[string]bool {
	names := make(map[string]bool, len(skills))
	for _, skill := range skills {
		names[skill.Name] = true
	}
	return names
}

func applyHarnessMetadata(harnessID harness.ID, skills []payload.Skill) ([]payload.Skill, error) {
	supported, err := harness.SupportsDisableModelInvocation(harnessID)
	if err != nil {
		return nil, err
	}
	if !supported {
		return skills, nil
	}

	result := make([]payload.Skill, 0, len(skills))
	for _, skill := range skills {
		if skill.Name != "skill-intake" {
			result = append(result, skill)
			continue
		}
		files := make(map[string][]byte, len(skill.Files))
		for relative, data := range skill.Files {
			files[relative] = append([]byte(nil), data...)
		}
		entrypoint, err := addDisableModelInvocation(files["SKILL.md"])
		if err != nil {
			return nil, fmt.Errorf("apply %s metadata to %s: %w", harnessID, skill.Name, err)
		}
		files["SKILL.md"] = entrypoint
		result = append(result, payload.Skill{Name: skill.Name, Files: files})
	}
	return result, nil
}

func addDisableModelInvocation(data []byte) ([]byte, error) {
	document, err := payload.ParseFrontmatter(data)
	if err != nil {
		return nil, err
	}
	for _, line := range document.Lines() {
		line = strings.TrimSpace(line)
		if !strings.HasPrefix(line, "disable-model-invocation:") {
			continue
		}
		if strings.TrimSpace(strings.TrimPrefix(line, "disable-model-invocation:")) == "true" {
			return append([]byte(nil), data...), nil
		}
		return nil, errors.New("SKILL.md has conflicting disable-model-invocation metadata")
	}
	text := string(data)
	return []byte(text[:document.CloseStart] + document.Newline + "disable-model-invocation: true" + text[document.CloseStart:]), nil
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
	document, err := payload.ParseFrontmatter(data)
	if err != nil {
		return nil, err
	}
	text := string(data)
	instruction := fmt.Sprintf("%sRun %q signal %q %q, then continue normally.%s", document.Newline, cliPath, token, signalStateRoot, document.Newline)
	return []byte(text[:document.BodyStart] + instruction + text[document.BodyStart:]), nil
}

func materialize(root, target string, harnessID harness.ID, skill payload.Skill) error {
	staging, err := os.MkdirTemp(root, ".skill-issue-stage-*")
	if err != nil {
		return fmt.Errorf("create skill staging directory: %w", err)
	}
	defer os.RemoveAll(staging)
	// MkdirTemp creates the directory 0700; widen it so the installed skill
	// directory matches the 0755 used for every other created directory.
	if err := os.Chmod(staging, 0o755); err != nil {
		return fmt.Errorf("open skill staging directory permissions: %w", err)
	}
	for relative, data := range skill.Files {
		include, err := harness.IncludeSkillFile(harnessID, relative)
		if err != nil {
			return err
		}
		if !include {
			continue
		}
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
