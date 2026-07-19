package receipt

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/harness"
)

const (
	ScopeProject = harness.ScopeProject
	ScopeUser    = harness.ScopeUser
)

type OwnedPath struct {
	Path   string `json:"path"`
	SHA256 string `json:"sha256,omitempty"`
}

type Receipt struct {
	SchemaVersion  int           `json:"schema_version"`
	InstallationID string        `json:"installation_id"`
	ProductVersion string        `json:"product_version"`
	Harness        string        `json:"harness"`
	Scope          harness.Scope `json:"scope"`
	Mode           string        `json:"mode"`
	RunID          string        `json:"run_id,omitempty"`
	InstallRoot    string        `json:"install_root"`
	PayloadSHA256  string        `json:"payload_sha256"`
	OwnedPaths     []OwnedPath   `json:"owned_paths"`
	CreatedAt      time.Time     `json:"created_at"`
	UpdatedAt      time.Time     `json:"updated_at"`
}

type Store struct {
	root string
}

func NewStore(root string) Store {
	return Store{root: root}
}

func (store Store) Save(receipt Receipt) error {
	if err := validateReceipt(receipt); err != nil {
		return err
	}
	if err := os.MkdirAll(store.root, 0o700); err != nil {
		return fmt.Errorf("create receipt directory: %w", err)
	}

	data, err := json.MarshalIndent(receipt, "", "  ")
	if err != nil {
		return fmt.Errorf("encode receipt: %w", err)
	}
	data = append(data, '\n')

	temporary, err := os.CreateTemp(store.root, ".receipt-*")
	if err != nil {
		return fmt.Errorf("create temporary receipt: %w", err)
	}
	temporaryPath := temporary.Name()
	defer os.Remove(temporaryPath)

	if err := writeTemporaryReceipt(temporary, data); err != nil {
		return err
	}
	if err := os.Rename(temporaryPath, store.path(receipt.InstallationID)); err != nil {
		return fmt.Errorf("commit receipt: %w", err)
	}
	return nil
}

func (store Store) Load(installationID string) (Receipt, error) {
	if err := validateID(installationID); err != nil {
		return Receipt{}, err
	}
	data, err := os.ReadFile(store.path(installationID))
	if err != nil {
		return Receipt{}, fmt.Errorf("read receipt: %w", err)
	}

	var receipt Receipt
	if err := json.Unmarshal(data, &receipt); err != nil {
		return Receipt{}, fmt.Errorf("decode receipt: %w", err)
	}
	if err := validateReceipt(receipt); err != nil {
		return Receipt{}, fmt.Errorf("stored receipt is invalid: %w", err)
	}
	return receipt, nil
}

func (store Store) Delete(installationID string) error {
	if err := validateID(installationID); err != nil {
		return err
	}
	if err := os.Remove(store.path(installationID)); err != nil {
		return fmt.Errorf("delete receipt: %w", err)
	}
	return nil
}

func (store Store) List() ([]Receipt, error) {
	entries, err := os.ReadDir(store.root)
	if errors.Is(err, os.ErrNotExist) {
		return nil, nil
	}
	if err != nil {
		return nil, fmt.Errorf("read receipt directory: %w", err)
	}
	var receipts []Receipt
	for _, entry := range entries {
		if entry.IsDir() || filepath.Ext(entry.Name()) != ".json" {
			continue
		}
		installationID := strings.TrimSuffix(entry.Name(), ".json")
		receipt, err := store.Load(installationID)
		if err != nil {
			return nil, err
		}
		receipts = append(receipts, receipt)
	}
	return receipts, nil
}

func (store Store) path(installationID string) string {
	return filepath.Join(store.root, installationID+".json")
}

func writeTemporaryReceipt(file *os.File, data []byte) error {
	if err := file.Chmod(0o600); err != nil {
		file.Close()
		return fmt.Errorf("secure temporary receipt: %w", err)
	}
	if _, err := file.Write(data); err != nil {
		file.Close()
		return fmt.Errorf("write temporary receipt: %w", err)
	}
	if err := file.Sync(); err != nil {
		file.Close()
		return fmt.Errorf("sync temporary receipt: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close temporary receipt: %w", err)
	}
	return nil
}

func validateReceipt(receipt Receipt) error {
	if receipt.SchemaVersion != 1 {
		return errors.New("receipt schema version must be 1")
	}
	if err := validateID(receipt.InstallationID); err != nil {
		return err
	}
	if receipt.Harness == "" || receipt.InstallRoot == "" || receipt.Mode == "" {
		return errors.New("receipt harness and install root are required")
	}
	if receipt.Scope != harness.ScopeProject && receipt.Scope != harness.ScopeUser {
		return errors.New("receipt scope must be project or user")
	}
	return nil
}

func validateID(installationID string) error {
	if installationID == "" || strings.ContainsAny(installationID, `/\\`) || installationID == "." || installationID == ".." {
		return errors.New("installation ID is invalid")
	}
	return nil
}
