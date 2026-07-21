package payload

import (
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"slices"
	"strings"

	skillissue "github.com/ericwimp8/skill-issue"
)

//go:embed assets/manifest.json
var manifestData []byte

type Component struct {
	ID     string `json:"id"`
	Source string `json:"source"`
}

type Manifest struct {
	SchemaVersion  int         `json:"schema_version"`
	Product        string      `json:"product"`
	PayloadVersion string      `json:"payload_version"`
	Components     []Component `json:"components"`
}

type Skill struct {
	Name  string
	Files map[string][]byte
}

var localReference = regexp.MustCompile("`((?:(?:references|scripts|assets)/|\\.\\.?/)[^`]+)`")

const builtInEvaluationRoot = "evaluations/skill-calling/built-ins"

func ReadManifest() (Manifest, error) {
	var manifest Manifest
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return Manifest{}, fmt.Errorf("decode embedded payload manifest: %w", err)
	}
	if manifest.SchemaVersion != 1 || manifest.Product == "" {
		return Manifest{}, fmt.Errorf("embedded payload manifest is invalid")
	}
	identifiers := make(map[string]bool, len(manifest.Components))
	sources := make(map[string]bool, len(manifest.Components))
	for _, component := range manifest.Components {
		cleanSource := path.Clean(component.Source)
		if component.ID == "" || path.Base(component.ID) != component.ID {
			return Manifest{}, fmt.Errorf("embedded payload manifest has invalid component ID %q", component.ID)
		}
		if component.Source == "" || cleanSource != component.Source || path.IsAbs(cleanSource) || cleanSource == ".." || strings.HasPrefix(cleanSource, "../") {
			return Manifest{}, fmt.Errorf("embedded payload manifest has invalid source %q", component.Source)
		}
		if identifiers[component.ID] {
			return Manifest{}, fmt.Errorf("embedded payload manifest repeats component %q", component.ID)
		}
		if sources[component.Source] {
			return Manifest{}, fmt.Errorf("embedded payload manifest repeats source %q", component.Source)
		}
		identifiers[component.ID] = true
		sources[component.Source] = true
	}
	return manifest, nil
}

func Skills() ([]Skill, error) {
	return readSkills()
}

func LoadSkills(root string) ([]Skill, error) {
	if strings.TrimSpace(root) == "" {
		return nil, fmt.Errorf("custom skill directory is required")
	}
	abs, err := filepath.Abs(root)
	if err != nil {
		return nil, fmt.Errorf("resolve custom skill directory: %w", err)
	}
	info, err := os.Stat(abs)
	if err != nil {
		return nil, fmt.Errorf("inspect custom skill directory: %w", err)
	}
	if !info.IsDir() {
		return nil, fmt.Errorf("custom skill path must be a directory")
	}
	entries, err := os.ReadDir(abs)
	if err != nil {
		return nil, fmt.Errorf("read custom skill directory: %w", err)
	}
	if len(entries) == 0 {
		return nil, fmt.Errorf("custom skill directory must contain at least one skill directory")
	}

	skills := make([]Skill, 0, len(entries))
	for _, entry := range entries {
		if entry.Type()&os.ModeSymlink != 0 {
			return nil, fmt.Errorf("custom skill directory contains unsupported symlink %q", entry.Name())
		}
		if !entry.IsDir() {
			return nil, fmt.Errorf("custom skill directory entry %q is not a skill directory", entry.Name())
		}
		files, err := readExternalSkill(abs, entry.Name())
		if err != nil {
			return nil, err
		}
		skills = append(skills, Skill{Name: entry.Name(), Files: files})
	}
	slices.SortFunc(skills, func(left, right Skill) int { return strings.Compare(left.Name, right.Name) })
	return skills, nil
}

func BuiltInEvaluation(id string) ([]byte, error) {
	if strings.TrimSpace(id) == "" || path.Base(id) != id {
		return nil, fmt.Errorf("invalid built-in evaluation %q", id)
	}
	data, err := fs.ReadFile(skillissue.CanonicalSkills, path.Join(builtInEvaluationRoot, id+".json"))
	if err != nil {
		return nil, fmt.Errorf("read built-in evaluation %q: %w", id, err)
	}
	return append([]byte(nil), data...), nil
}

func readSkills() ([]Skill, error) {
	manifest, err := ReadManifest()
	if err != nil {
		return nil, err
	}

	skills := make([]Skill, 0, len(manifest.Components))
	for _, component := range manifest.Components {
		files, err := readSkill(component.Source, component.ID)
		if err != nil {
			return nil, err
		}
		skills = append(skills, Skill{Name: component.ID, Files: files})
	}
	slices.SortFunc(skills, func(left, right Skill) int { return strings.Compare(left.Name, right.Name) })
	return skills, nil
}

func readSkill(source, name string) (map[string][]byte, error) {
	base := source
	files := map[string][]byte{}
	err := fs.WalkDir(skillissue.CanonicalSkills, base, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Name() == ".DS_Store" {
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		data, err := fs.ReadFile(skillissue.CanonicalSkills, filePath)
		if err != nil {
			return err
		}
		relative := strings.TrimPrefix(filePath, base+"/")
		files[relative] = data
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("read canonical skill %q: %w", name, err)
	}
	entrypoint, ok := files["SKILL.md"]
	if !ok {
		return nil, fmt.Errorf("canonical skill %q has no SKILL.md", name)
	}
	if err := validateFrontmatter(name, entrypoint); err != nil {
		return nil, err
	}
	if err := validateReferenceClosure(name, files); err != nil {
		return nil, err
	}
	return files, nil
}

func readExternalSkill(root, name string) (map[string][]byte, error) {
	base := filepath.Join(root, name)
	files := map[string][]byte{}
	err := filepath.WalkDir(base, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if entry.Type()&os.ModeSymlink != 0 {
			return fmt.Errorf("custom skill %q contains unsupported symlink %q", name, filePath)
		}
		if entry.Name() == ".DS_Store" {
			return nil
		}
		if entry.IsDir() {
			return nil
		}
		data, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		relative, err := filepath.Rel(base, filePath)
		if err != nil {
			return err
		}
		files[filepath.ToSlash(relative)] = data
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("read custom skill %q: %w", name, err)
	}
	entrypoint, ok := files["SKILL.md"]
	if !ok {
		return nil, fmt.Errorf("custom skill %q has no SKILL.md", name)
	}
	if err := validateFrontmatter(name, entrypoint); err != nil {
		return nil, err
	}
	if err := validateReferenceClosure(name, files); err != nil {
		return nil, err
	}
	return files, nil
}

func validateReferenceClosure(name string, files map[string][]byte) error {
	for source, data := range files {
		if path.Ext(source) != ".md" {
			continue
		}
		inFence := false
		var fenceMarker byte
		var fenceLength int
		for _, line := range strings.Split(string(data), "\n") {
			trimmed := strings.TrimSpace(line)
			if inFence {
				// Only a fence of the same character at least as long as the
				// opener closes the block.
				if fenceClosing(trimmed, fenceMarker, fenceLength) {
					inFence = false
				}
				continue
			}
			if marker, length, ok := fenceOpening(trimmed); ok {
				inFence, fenceMarker, fenceLength = true, marker, length
				continue
			}
			for _, match := range localReference.FindAllStringSubmatch(line, -1) {
				reference := path.Clean(match[1])
				if strings.HasPrefix(match[1], ".") {
					reference = path.Clean(path.Join(path.Dir(source), match[1]))
				}
				if reference == ".." || strings.HasPrefix(reference, "../") {
					return fmt.Errorf("canonical skill %q has escaping reference %q in %q", name, match[1], source)
				}
				if _, ok := files[reference]; !ok {
					return fmt.Errorf("canonical skill %q references absent file %q in %q", name, reference, source)
				}
			}
		}
	}
	return nil
}

// fenceOpening reports whether a line opens a Markdown code fence: at least
// three backticks or tildes, where a backtick fence's info string may not
// contain further backticks (per CommonMark).
func fenceOpening(line string) (byte, int, bool) {
	if line == "" || (line[0] != '`' && line[0] != '~') {
		return 0, 0, false
	}
	marker := line[0]
	length := fenceRunLength(line, marker)
	if length < 3 {
		return 0, 0, false
	}
	if marker == '`' && strings.ContainsRune(line[length:], '`') {
		return 0, 0, false
	}
	return marker, length, true
}

func fenceClosing(line string, marker byte, minimum int) bool {
	length := fenceRunLength(line, marker)
	return length >= minimum && strings.TrimSpace(line[length:]) == ""
}

func fenceRunLength(line string, marker byte) int {
	length := 0
	for length < len(line) && line[length] == marker {
		length++
	}
	return length
}

func validateFrontmatter(name string, data []byte) error {
	document, err := ParseFrontmatter(data)
	if errors.Is(err, ErrMissingOpeningFrontmatter) {
		return fmt.Errorf("canonical skill %q has invalid frontmatter", name)
	}
	if err != nil {
		return fmt.Errorf("canonical skill %q has unterminated frontmatter", name)
	}
	hasName := false
	hasDescription := false
	for _, line := range document.Lines() {
		if strings.TrimRight(line, " \t") == "name: "+name {
			hasName = true
		}
		if strings.HasPrefix(line, "description:") {
			hasDescription = true
		}
	}
	if !hasName {
		return fmt.Errorf("canonical skill %q frontmatter name does not match its directory", name)
	}
	if !hasDescription {
		return fmt.Errorf("canonical skill %q has no description", name)
	}
	return nil
}
