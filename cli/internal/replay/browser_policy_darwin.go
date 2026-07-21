//go:build darwin

package replay

import (
	"context"
	"os/exec"
)

const browserForbiddenProfile = `(version 1)
(allow default)
(deny process-exec
  (regex #"/(Google Chrome([^/]*)|Chromium|chrome|chrome-headless-shell|chromium|chromium-browser|google-chrome|msedge|Microsoft Edge|Brave Browser|Firefox|firefox|MiniBrowser|Safari|open)$"))`

func evaluationCommand(ctx context.Context, path string, args []string, policy BrowserPolicy) *exec.Cmd {
	if policy != BrowserPolicyForbidden {
		return exec.CommandContext(ctx, path, args...)
	}
	sandboxArgs := append([]string{"-p", browserForbiddenProfile, path}, args...)
	return exec.CommandContext(ctx, "/usr/bin/sandbox-exec", sandboxArgs...)
}
