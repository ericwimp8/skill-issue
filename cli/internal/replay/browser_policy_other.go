//go:build !darwin

package replay

import (
	"context"
	"os/exec"
)

func evaluationCommand(ctx context.Context, path string, args []string, _ BrowserPolicy) *exec.Cmd {
	return exec.CommandContext(ctx, path, args...)
}
