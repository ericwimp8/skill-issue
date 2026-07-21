//go:build darwin

package replay

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestForbiddenBrowserPolicyBlocksDescendantExecutable(t *testing.T) {
	directory := t.TempDir()
	browser := filepath.Join(directory, "chrome")
	harness := filepath.Join(directory, "harness")
	if err := os.WriteFile(browser, []byte("#!/bin/sh\nprintf launched\n"), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(harness, []byte("#!/bin/sh\nexec \"$BROWSER\"\n"), 0o700); err != nil {
		t.Fatal(err)
	}

	command := evaluationCommand(context.Background(), harness, nil, BrowserPolicyForbidden)
	command.Env = append(os.Environ(), "BROWSER="+browser)
	if output, err := command.CombinedOutput(); err == nil {
		t.Fatalf("browser descendant ran under forbidden policy: %s", output)
	}

	command = evaluationCommand(context.Background(), harness, nil, BrowserPolicyAllowed)
	command.Env = append(os.Environ(), "BROWSER="+browser)
	output, err := command.CombinedOutput()
	if err != nil {
		t.Fatalf("allowed browser descendant failed: %v: %s", err, output)
	}
	if string(output) != "launched" {
		t.Fatalf("unexpected allowed output %q", output)
	}
}
