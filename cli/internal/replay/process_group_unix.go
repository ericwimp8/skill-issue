//go:build !windows

package replay

import (
	"errors"
	"os/exec"
	"syscall"
)

func configureOwnedProcess(command *exec.Cmd) {
	command.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	command.Cancel = func() error {
		return killOwnedProcessGroup(command)
	}
}

func stopOwnedProcessGroup(command *exec.Cmd) {
	_ = killOwnedProcessGroup(command)
}

func killOwnedProcessGroup(command *exec.Cmd) error {
	if command.Process == nil {
		return nil
	}
	err := syscall.Kill(-command.Process.Pid, syscall.SIGKILL)
	if errors.Is(err, syscall.ESRCH) {
		return nil
	}
	return err
}
