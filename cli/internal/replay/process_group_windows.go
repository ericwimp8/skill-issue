//go:build windows

package replay

import "os/exec"

func configureOwnedProcess(command *exec.Cmd) {}

func stopOwnedProcessGroup(command *exec.Cmd) {
	if command.Process != nil {
		_ = command.Process.Kill()
	}
}
