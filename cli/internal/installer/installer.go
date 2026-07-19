package installer

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
	"github.com/ericwimp8/skill-issue/cli/internal/payload"
	"github.com/ericwimp8/skill-issue/cli/internal/receipt"
)

const (
	ModeOrdinary   = "ordinary"
	ModeEvaluation = "evaluation"
)

type Request struct {
	Harness        harness.ID
	Scope          harness.Scope
	Workspace      string
	Home           string
	ProductVersion string
	Mode           string
	RunID          string
	CLIPath        string
	Tokens         map[string]string
}

type EvaluationBackup struct {
	ReceiptID    string
	PriorReceipt *receipt.Receipt
	BackupRoot   string
}

type Service struct {
	stateRoot string
	receipts  receipt.Store
}

func New(stateRoot string) Service {
	return Service{
		stateRoot: stateRoot,
		receipts:  receipt.NewStore(filepath.Join(stateRoot, "receipts")),
	}
}

func (service Service) Install(request Request) (receipt.Receipt, error) {
	skills, err := payload.Skills()
	if err != nil {
		return receipt.Receipt{}, err
	}
	if request.Mode == ModeEvaluation {
		skills, err = payload.EvaluationSkills()
		if err != nil {
			return receipt.Receipt{}, err
		}
		skills, err = instrument(skills, request.CLIPath, request.Tokens)
		if err != nil {
			return receipt.Receipt{}, err
		}
	}
	root, err := harness.SkillRoot(request.Harness, request.Scope, request.Workspace, request.Home)
	if err != nil {
		return receipt.Receipt{}, err
	}
	receiptID := installationID(request.Harness, request.Scope, root)
	existing, existingErr := service.receipts.Load(receiptID)
	if existingErr != nil && !errors.Is(existingErr, os.ErrNotExist) && !strings.Contains(existingErr.Error(), "no such file") {
		return receipt.Receipt{}, existingErr
	}
	if existingErr == nil && existing.Mode != request.Mode {
		return receipt.Receipt{}, fmt.Errorf("%s installation already owns %s", existing.Mode, root)
	}
	owned := map[string]bool{}
	if existingErr == nil {
		for _, item := range existing.OwnedPaths {
			owned[item.Path] = true
		}
	}
	for _, skill := range skills {
		target := filepath.Join(root, skill.Name)
		if _, err := os.Stat(target); err == nil && !owned[target] {
			return receipt.Receipt{}, fmt.Errorf("foreign skill collision at %s", target)
		} else if err != nil && !errors.Is(err, os.ErrNotExist) {
			return receipt.Receipt{}, fmt.Errorf("inspect skill destination: %w", err)
		}
	}
	if err := os.MkdirAll(root, 0o755); err != nil {
		return receipt.Receipt{}, fmt.Errorf("create skill root: %w", err)
	}
	backupRoot, err := os.MkdirTemp(service.stateRoot, ".install-backup-*")
	if err != nil {
		return receipt.Receipt{}, fmt.Errorf("create installation backup: %w", err)
	}
	defer os.RemoveAll(backupRoot)
	committed := []string{}
	for _, skill := range skills {
		target := filepath.Join(root, skill.Name)
		if _, err := os.Stat(target); err == nil {
			if err := copyTree(target, filepath.Join(backupRoot, skill.Name)); err != nil {
				service.rollback(root, backupRoot, committed)
				return receipt.Receipt{}, err
			}
		}
		committed = append(committed, skill.Name)
		if err := materialize(root, target, skill); err != nil {
			service.rollback(root, backupRoot, committed)
			return receipt.Receipt{}, err
		}
	}
	ownedPaths := make([]receipt.OwnedPath, 0, len(skills))
	for _, skill := range skills {
		digest, err := directoryDigest(filepath.Join(root, skill.Name))
		if err != nil {
			return receipt.Receipt{}, err
		}
		ownedPaths = append(ownedPaths, receipt.OwnedPath{Path: filepath.Join(root, skill.Name), SHA256: digest})
	}
	now := time.Now().UTC()
	createdAt := now
	if existingErr == nil {
		createdAt = existing.CreatedAt
	}
	installed := receipt.Receipt{
		SchemaVersion:  1,
		InstallationID: receiptID,
		ProductVersion: request.ProductVersion,
		Harness:        string(request.Harness),
		Scope:          request.Scope,
		Mode:           request.Mode,
		RunID:          request.RunID,
		InstallRoot:    root,
		PayloadSHA256:  skillsDigest(skills),
		OwnedPaths:     ownedPaths,
		CreatedAt:      createdAt,
		UpdatedAt:      now,
	}
	if err := service.receipts.Save(installed); err != nil {
		service.rollback(root, backupRoot, committed)
		return receipt.Receipt{}, err
	}
	return installed, nil
}

func (service Service) PrepareEvaluation(request Request, privateRunDir string) (EvaluationBackup, receipt.Receipt, error) {
	request.Mode = ModeEvaluation
	root, err := harness.SkillRoot(request.Harness, request.Scope, request.Workspace, request.Home)
	if err != nil {
		return EvaluationBackup{}, receipt.Receipt{}, err
	}
	receiptID := installationID(request.Harness, request.Scope, root)
	backup := EvaluationBackup{ReceiptID: receiptID, BackupRoot: filepath.Join(privateRunDir, "prior-installation")}
	prior, err := service.receipts.Load(receiptID)
	if err == nil {
		if prior.Mode != ModeOrdinary {
			return EvaluationBackup{}, receipt.Receipt{}, fmt.Errorf("installation at %s is already in %s mode", root, prior.Mode)
		}
		backup.PriorReceipt = &prior
		for _, item := range prior.OwnedPaths {
			if err := copyTree(item.Path, filepath.Join(backup.BackupRoot, filepath.Base(item.Path))); err != nil {
				return EvaluationBackup{}, receipt.Receipt{}, err
			}
		}
		if err := service.removeOwned(prior); err != nil {
			return EvaluationBackup{}, receipt.Receipt{}, err
		}
		if err := service.receipts.Delete(receiptID); err != nil {
			service.restorePrior(backup)
			return EvaluationBackup{}, receipt.Receipt{}, err
		}
	} else if !errors.Is(err, os.ErrNotExist) {
		return EvaluationBackup{}, receipt.Receipt{}, err
	}
	installed, err := service.Install(request)
	if err != nil {
		if backup.PriorReceipt != nil {
			if restoreErr := service.restorePrior(backup); restoreErr != nil {
				return EvaluationBackup{}, receipt.Receipt{}, errors.Join(err, restoreErr)
			}
		}
		return EvaluationBackup{}, receipt.Receipt{}, err
	}
	return backup, installed, nil
}

func (service Service) CleanupEvaluation(backup EvaluationBackup, runID string) error {
	current, err := service.receipts.Load(backup.ReceiptID)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}
	if err == nil {
		if current.Mode != ModeEvaluation {
			if backup.PriorReceipt != nil && current.InstallationID == backup.PriorReceipt.InstallationID && current.PayloadSHA256 == backup.PriorReceipt.PayloadSHA256 {
				return nil
			}
			return fmt.Errorf("refusing to clean non-evaluation installation %s", current.InstallationID)
		}
		if current.RunID != runID {
			return fmt.Errorf("evaluation installation belongs to run %s", current.RunID)
		}
		if err := service.removeOwned(current); err != nil {
			return err
		}
		if err := service.receipts.Delete(current.InstallationID); err != nil {
			return err
		}
	}
	if backup.PriorReceipt == nil {
		return nil
	}
	return service.restorePrior(backup)
}

func (service Service) restorePrior(backup EvaluationBackup) error {
	if backup.PriorReceipt == nil {
		return nil
	}
	for _, item := range backup.PriorReceipt.OwnedPaths {
		source := filepath.Join(backup.BackupRoot, filepath.Base(item.Path))
		if err := copyTree(source, item.Path); err != nil {
			return err
		}
	}
	return service.receipts.Save(*backup.PriorReceipt)
}

func (service Service) Verify(receiptID string) (receipt.Receipt, error) {
	installed, err := service.receipts.Load(receiptID)
	if err != nil {
		return receipt.Receipt{}, err
	}
	for _, item := range installed.OwnedPaths {
		digest, err := directoryDigest(item.Path)
		if err != nil {
			return receipt.Receipt{}, err
		}
		if digest != item.SHA256 {
			return receipt.Receipt{}, fmt.Errorf("owned path drift at %s", item.Path)
		}
	}
	return installed, nil
}

func (service Service) Uninstall(receiptID string) error {
	installed, err := service.Verify(receiptID)
	if err != nil {
		return err
	}
	if err := service.removeOwned(installed); err != nil {
		return err
	}
	return service.receipts.Delete(receiptID)
}

func (service Service) ReceiptID(id harness.ID, scope harness.Scope, workspace, home string) (string, error) {
	root, err := harness.SkillRoot(id, scope, workspace, home)
	if err != nil {
		return "", err
	}
	return installationID(id, scope, root), nil
}

func (service Service) ReceiptStore() receipt.Store {
	return service.receipts
}

func (service Service) removeOwned(installed receipt.Receipt) error {
	for _, item := range installed.OwnedPaths {
		if err := os.RemoveAll(item.Path); err != nil {
			return fmt.Errorf("remove owned path %s: %w", item.Path, err)
		}
	}
	return nil
}

func (service Service) rollback(root, backupRoot string, committed []string) {
	for _, name := range committed {
		target := filepath.Join(root, name)
		os.RemoveAll(target)
		backup := filepath.Join(backupRoot, name)
		if _, err := os.Stat(backup); err == nil {
			copyTree(backup, target)
		}
	}
}

func instrument(skills []payload.Skill, cliPath string, tokens map[string]string) ([]payload.Skill, error) {
	if !filepath.IsAbs(cliPath) {
		return nil, errors.New("evaluation CLI path must be absolute")
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
		entrypoint, err := inject(files["SKILL.md"], cliPath, token)
		if err != nil {
			return nil, fmt.Errorf("instrument %s: %w", skill.Name, err)
		}
		files["SKILL.md"] = entrypoint
		result = append(result, payload.Skill{Name: skill.Name, Files: files})
	}
	return result, nil
}

func inject(data []byte, cliPath, token string) ([]byte, error) {
	text := string(data)
	if !strings.HasPrefix(text, "---\n") {
		return nil, errors.New("SKILL.md has no opening frontmatter delimiter")
	}
	end := strings.Index(text[4:], "\n---\n")
	if end < 0 {
		return nil, errors.New("SKILL.md has no closing frontmatter delimiter")
	}
	insertAt := 4 + end + len("\n---\n")
	instruction := fmt.Sprintf("\nRun %q mark %q, then continue normally.\n", cliPath, token)
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
		return fmt.Errorf("remove prior owned skill: %w", err)
	}
	if err := os.Rename(staging, target); err != nil {
		return fmt.Errorf("commit skill directory: %w", err)
	}
	return nil
}

func copyTree(source, destination string) error {
	info, err := os.Stat(source)
	if err != nil {
		return fmt.Errorf("inspect copy source: %w", err)
	}
	if !info.IsDir() {
		return errors.New("copy source must be a directory")
	}
	if err := os.RemoveAll(destination); err != nil {
		return err
	}
	return filepath.Walk(source, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, relative)
		if info.IsDir() {
			return os.MkdirAll(target, info.Mode().Perm())
		}
		input, err := os.Open(path)
		if err != nil {
			return err
		}
		defer input.Close()
		output, err := os.OpenFile(target, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode().Perm())
		if err != nil {
			return err
		}
		if _, err := io.Copy(output, input); err != nil {
			output.Close()
			return err
		}
		return output.Close()
	})
}

func directoryDigest(root string) (string, error) {
	hash := sha256.New()
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if !info.IsDir() {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		return "", fmt.Errorf("enumerate owned directory: %w", err)
	}
	sort.Strings(files)
	for _, file := range files {
		relative, _ := filepath.Rel(root, file)
		hash.Write([]byte(filepath.ToSlash(relative)))
		data, err := os.ReadFile(file)
		if err != nil {
			return "", fmt.Errorf("hash owned file: %w", err)
		}
		hash.Write(data)
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func skillsDigest(skills []payload.Skill) string {
	hash := sha256.New()
	for _, skill := range skills {
		hash.Write([]byte(skill.Name))
		var names []string
		for name := range skill.Files {
			names = append(names, name)
		}
		sort.Strings(names)
		for _, name := range names {
			hash.Write([]byte(name))
			hash.Write(skill.Files[name])
		}
	}
	return hex.EncodeToString(hash.Sum(nil))
}

func installationID(id harness.ID, scope harness.Scope, root string) string {
	digest := sha256.Sum256([]byte(string(id) + "\x00" + string(scope) + "\x00" + root))
	return hex.EncodeToString(digest[:16])
}

func EncodeBackup(backup EvaluationBackup) ([]byte, error) {
	return json.MarshalIndent(backup, "", "  ")
}

func DecodeBackup(data []byte) (EvaluationBackup, error) {
	var backup EvaluationBackup
	if err := json.Unmarshal(data, &backup); err != nil {
		return EvaluationBackup{}, err
	}
	if backup.ReceiptID == "" {
		return EvaluationBackup{}, errors.New("evaluation backup is invalid")
	}
	return backup, nil
}
