package payload

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"path"
	"regexp"
	"sort"
	"strings"

	skillissue "github.com/ericwimp8/skill-issue"
)

//go:embed assets/manifest.json
var manifestData []byte

type Component struct {
	ID             string `json:"id"`
	EvaluationOnly bool   `json:"evaluation_only,omitempty"`
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

var localReference = regexp.MustCompile("`((?:references|scripts|assets)/[^`]+)`")

func ReadManifest() (Manifest, error) {
	var manifest Manifest
	if err := json.Unmarshal(manifestData, &manifest); err != nil {
		return Manifest{}, fmt.Errorf("decode embedded payload manifest: %w", err)
	}
	if manifest.SchemaVersion != 1 || manifest.Product == "" {
		return Manifest{}, fmt.Errorf("embedded payload manifest is invalid")
	}
	return manifest, nil
}

func Skills() ([]Skill, error) {
	return readSkills(false)
}

func EvaluationSkills() ([]Skill, error) {
	return readSkills(true)
}

func readSkills(includeEvaluationOnly bool) ([]Skill, error) {
	manifest, err := ReadManifest()
	if err != nil {
		return nil, err
	}

	components := make(map[string]Component, len(manifest.Components))
	for _, component := range manifest.Components {
		components[component.ID] = component
	}

	var skills []Skill
	for _, root := range []string{"skills", "supporting-skills", "evaluation-skills"} {
		entries, err := fs.ReadDir(skillissue.CanonicalSkills, root)
		if err != nil {
			return nil, fmt.Errorf("read canonical %s: %w", root, err)
		}
		for _, entry := range entries {
			if !entry.IsDir() {
				continue
			}
			name := entry.Name()
			component, ok := components[name]
			if !ok {
				return nil, fmt.Errorf("canonical skill %q is absent from payload manifest", name)
			}
			if root == "evaluation-skills" && !component.EvaluationOnly {
				return nil, fmt.Errorf("evaluation skill %q is not marked evaluation-only", name)
			}
			if root != "evaluation-skills" && component.EvaluationOnly {
				return nil, fmt.Errorf("ordinary skill %q is marked evaluation-only", name)
			}
			files, err := readSkill(root, name)
			if err != nil {
				return nil, err
			}
			if includeEvaluationOnly || !component.EvaluationOnly {
				skills = append(skills, Skill{Name: name, Files: files})
			}
			delete(components, name)
		}
	}
	if len(components) != 0 {
		return nil, fmt.Errorf("payload manifest references absent canonical skills")
	}
	sort.Slice(skills, func(i, j int) bool { return skills[i].Name < skills[j].Name })
	return skills, nil
}

func readSkill(root, name string) (map[string][]byte, error) {
	base := path.Join(root, name)
	files := map[string][]byte{}
	err := fs.WalkDir(skillissue.CanonicalSkills, base, func(filePath string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
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
	if err := validateReferenceClosure(name, entrypoint, files); err != nil {
		return nil, err
	}
	return files, nil
}

func validateReferenceClosure(name string, entrypoint []byte, files map[string][]byte) error {
	inFence := false
	for _, line := range strings.Split(string(entrypoint), "\n") {
		if strings.HasPrefix(strings.TrimSpace(line), "```") {
			inFence = !inFence
			continue
		}
		if inFence {
			continue
		}
		for _, match := range localReference.FindAllStringSubmatch(line, -1) {
			reference := path.Clean(match[1])
			if strings.HasPrefix(reference, "../") {
				return fmt.Errorf("canonical skill %q has escaping reference %q", name, reference)
			}
			if _, ok := files[reference]; !ok {
				return fmt.Errorf("canonical skill %q references absent file %q", name, reference)
			}
		}
	}
	return nil
}

func validateFrontmatter(name string, data []byte) error {
	text := string(data)
	if !strings.HasPrefix(text, "---\n") {
		return fmt.Errorf("canonical skill %q has invalid frontmatter", name)
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		return fmt.Errorf("canonical skill %q has unterminated frontmatter", name)
	}
	frontmatter := text[4 : 4+end]
	if !strings.Contains(frontmatter, "name: "+name+"\n") && !strings.HasSuffix(frontmatter, "name: "+name) {
		return fmt.Errorf("canonical skill %q frontmatter name does not match its directory", name)
	}
	if !strings.Contains(frontmatter, "description:") {
		return fmt.Errorf("canonical skill %q has no description", name)
	}
	return nil
}
