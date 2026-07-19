package replay

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"
	"sync"
)

type piAdapter struct {
	path      string
	directory string
	environ   []string
	model     string
}

func newPiAdapter(options Options) (Adapter, error) {
	path, err := resolveExecutable("pi", options.Executable)
	if err != nil {
		return nil, fmt.Errorf("pi executable: %w", err)
	}
	return &piAdapter{path: path, directory: options.Directory, environ: options.Environment, model: options.Model}, nil
}

func (adapter *piAdapter) HarnessID() HarnessID { return HarnessPi }

func (adapter *piAdapter) Start(ctx context.Context) (Session, error) {
	args := []string{"--mode", "rpc"}
	if adapter.model != "" {
		args = append(args, "--model", adapter.model)
	}
	command := exec.CommandContext(ctx, adapter.path, args...)
	command.Dir = adapter.directory
	command.Env = mergedEnvironment(adapter.environ)
	stdin, err := command.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := command.StdoutPipe()
	if err != nil {
		return nil, err
	}
	var stderr lockedBuffer
	command.Stderr = &stderr
	if err := command.Start(); err != nil {
		return nil, err
	}
	return &piSession{command: command, stdin: stdin, stdout: bufio.NewReader(stdout), stderr: &stderr}, nil
}

type piSession struct {
	command *exec.Cmd
	stdin   io.WriteCloser
	stdout  *bufio.Reader
	stderr  *lockedBuffer
	pending bool
	closed  bool
}

type lockedBuffer struct {
	mutex sync.Mutex
	data  bytes.Buffer
}

func (buffer *lockedBuffer) Write(data []byte) (int, error) {
	buffer.mutex.Lock()
	defer buffer.mutex.Unlock()
	return buffer.data.Write(data)
}

func (buffer *lockedBuffer) String() string {
	buffer.mutex.Lock()
	defer buffer.mutex.Unlock()
	return buffer.data.String()
}

func (session *piSession) SendPrompt(_ context.Context, prompt string) error {
	if session.closed {
		return errors.New("session is closed")
	}
	if session.pending {
		return errors.New("previous prompt is still running")
	}
	message, err := json.Marshal(struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}{Type: "prompt", Message: prompt})
	if err != nil {
		return err
	}
	if _, err := session.stdin.Write(append(message, '\n')); err != nil {
		return fmt.Errorf("write Pi RPC prompt: %w", err)
	}
	session.pending = true
	return nil
}

func (session *piSession) Wait(ctx context.Context) (Capture, error) {
	if !session.pending {
		return Capture{}, errors.New("no prompt is running")
	}
	session.pending = false
	type readResult struct {
		line []byte
		err  error
	}
	events := make([]json.RawMessage, 0, 8)
	var transcript bytes.Buffer
	for {
		resultChannel := make(chan readResult, 1)
		go func() {
			line, err := session.stdout.ReadBytes('\n')
			resultChannel <- readResult{line: line, err: err}
		}()
		var result readResult
		select {
		case <-ctx.Done():
			return Capture{}, ctx.Err()
		case result = <-resultChannel:
		}
		if result.err != nil {
			return Capture{}, fmt.Errorf("Pi RPC ended before terminal event: %w: %s", result.err, strings.TrimSpace(session.stderr.String()))
		}
		line := bytes.TrimSpace(result.line)
		var event map[string]any
		if len(line) == 0 || json.Unmarshal(line, &event) != nil {
			return Capture{}, fmt.Errorf("%w: invalid Pi RPC event", ErrProtocol)
		}
		transcript.Write(result.line)
		events = append(events, append(json.RawMessage(nil), line...))
		if event["type"] == "agent_end" {
			return Capture{Transcript: transcript.String(), Stderr: session.stderr.String(), Events: events}, nil
		}
	}
}

func (session *piSession) Close() error {
	if session.closed {
		return nil
	}
	session.closed = true
	_ = session.stdin.Close()
	err := session.command.Wait()
	if err == nil {
		return nil
	}
	if session.command.ProcessState != nil && session.command.ProcessState.Exited() {
		return fmt.Errorf("Pi RPC exited unsuccessfully: %w: %s", err, strings.TrimSpace(session.stderr.String()))
	}
	return session.command.Process.Kill()
}
