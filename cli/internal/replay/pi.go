package replay

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"
)

// piRequestTimeout bounds control-plane RPC exchanges (get_state, get_commands)
// so a hung pi process fails the run instead of blocking it indefinitely.
const piRequestTimeout = 30 * time.Second

type piAdapter struct {
	path             string
	directory        string
	environ          []string
	cleanEnvironment bool
	model            string
	reasoning        string
	skillsRoot       string
	browserPolicy    BrowserPolicy
}

func newPiAdapter(options Options) (Adapter, error) {
	if err := options.BrowserPolicy.Validate(); err != nil {
		return nil, err
	}
	path, err := resolveExecutable("pi", options.Executable)
	if err != nil {
		return nil, fmt.Errorf("pi executable: %w", err)
	}
	return &piAdapter{
		path:             path,
		directory:        options.Directory,
		environ:          options.Environment,
		cleanEnvironment: options.CleanEnvironment,
		model:            options.Model,
		reasoning:        options.Reasoning,
		skillsRoot:       options.PiSkillsRoot,
		browserPolicy:    options.BrowserPolicy,
	}, nil
}

func (adapter *piAdapter) HarnessID() HarnessID { return HarnessPi }

func (adapter *piAdapter) Start(ctx context.Context) (Session, error) {
	provider, model := piModel(adapter.model)
	sessionID, err := generatedSessionID()
	if err != nil {
		return nil, err
	}
	args := []string{"--mode", "rpc", "--provider", provider, "--model", model, "--thinking", adapter.reasoning, "--no-session", "--session-id", sessionID, "--no-approve", "--no-extensions", "--no-skills"}
	skills, err := explicitPiSkills(adapter.skillsRoot)
	if err != nil {
		return nil, err
	}
	for _, skill := range skills {
		args = append(args, "--skill", skill)
	}
	args = append(args, "--no-prompt-templates", "--no-themes", "--no-context-files", "--tools", "read,bash,edit,write,grep,find,ls", "--offline")
	command := evaluationCommand(ctx, adapter.path, args, adapter.browserPolicy)
	configureOwnedProcess(command)
	command.Dir = adapter.directory
	command.Env = environment(adapter.environ, adapter.cleanEnvironment)
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
		return nil, fmt.Errorf("start Pi executable: %w: %s", err, strings.TrimSpace(stderr.String()))
	}
	session := &piSession{
		command:        command,
		commandLine:    strings.Join(append([]string{adapter.path}, args...), " "),
		stdin:          stdin,
		stdout:         bufio.NewReader(stdout),
		stderr:         &stderr,
		sessionID:      sessionID,
		nextID:         1,
		expectedSkills: skills,
		provider:       provider,
		model:          model,
		reasoning:      adapter.reasoning,
	}
	if err := session.preflight(ctx); err != nil {
		_ = session.Close()
		return nil, session.failure(err, "")
	}
	return session, nil
}

type piSession struct {
	command        *exec.Cmd
	commandLine    string
	stdin          io.WriteCloser
	stdout         *bufio.Reader
	stderr         *lockedBuffer
	sessionID      string
	nextID         int
	expectedSkills []string
	provider       string
	model          string
	reasoning      string
	pending        bool
	pendingID      string
	closed         bool
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

// failure wraps a Pi error with the launch command, the transcript captured
// so far, and Pi's stderr for post-mortem diagnostics.
func (session *piSession) failure(err error, transcript string) error {
	return &DiagnosticError{
		Diagnostic: Diagnostic{Command: session.commandLine, Stdout: transcript, Stderr: session.stderr.String()},
		Err:        err,
	}
}

func (session *piSession) preflight(ctx context.Context) error {
	state, err := session.request(ctx, "get_state")
	if err != nil {
		return fmt.Errorf("Pi get_state preflight failed: %w", err)
	}
	if err := validatePiState(state, session.sessionID, session.provider, session.model, session.reasoning); err != nil {
		return err
	}
	commands, err := session.request(ctx, "get_commands")
	if err != nil {
		return fmt.Errorf("Pi get_commands preflight failed: %w", err)
	}
	for _, skill := range session.expectedSkills {
		if !containsString(commands, skill) {
			return fmt.Errorf("%w: Pi did not load supplied skill %s", ErrProtocol, filepath.Base(skill))
		}
	}
	return nil
}

func (session *piSession) SendPrompt(_ context.Context, prompt string) error {
	if session.closed {
		return errors.New("session is closed")
	}
	if session.pending {
		return errors.New("previous prompt is still running")
	}
	id := session.commandID("turn")
	message, err := json.Marshal(map[string]any{"id": id, "type": "prompt", "message": prompt})
	if err != nil {
		return err
	}
	if _, err := session.stdin.Write(append(message, '\n')); err != nil {
		return fmt.Errorf("write Pi RPC prompt: %w", err)
	}
	session.pending = true
	session.pendingID = id
	return nil
}

func (session *piSession) Wait(ctx context.Context) (Capture, error) {
	if !session.pending {
		return Capture{}, errors.New("no prompt is running")
	}
	session.pending = false
	events := make([]json.RawMessage, 0, 16)
	var transcript bytes.Buffer
	accepted := false
	settled := false
	var agentError string
	for !(accepted && settled) {
		event, err := session.readEvent(ctx)
		if err != nil {
			if ctx.Err() != nil {
				session.abort()
			}
			return Capture{}, session.failure(err, transcript.String())
		}
		transcript.Write(event)
		transcript.WriteByte('\n')
		events = append(events, append(json.RawMessage(nil), bytes.TrimSpace(event)...))
		var value map[string]any
		if json.Unmarshal(event, &value) != nil {
			return Capture{}, session.failure(fmt.Errorf("%w: invalid Pi RPC event", ErrProtocol), transcript.String())
		}
		if value["type"] == "response" {
			if success, ok := value["success"].(bool); ok && !success {
				session.abort()
				return Capture{}, session.failure(fmt.Errorf("Pi RPC command failed: %s", strings.TrimSpace(string(event))), transcript.String())
			}
			if value["id"] == session.pendingID && value["command"] == "prompt" {
				accepted = true
			}
		}
		if value["type"] == "agent_settled" {
			settled = true
		}
		if value["type"] == "agent_end" {
			if stop, _ := value["stopReason"].(string); stop == "error" || stop == "aborted" {
				agentError = fmt.Sprintf("Pi agent ended with %s: %s", stop, strings.TrimSpace(string(event)))
			}
		}
		if value["type"] == "error" {
			session.abort()
			return Capture{}, session.failure(fmt.Errorf("Pi RPC reported an error: %s", strings.TrimSpace(string(event))), transcript.String())
		}
	}
	if agentError != "" {
		return Capture{}, session.failure(errors.New(agentError), transcript.String())
	}
	state, err := session.request(ctx, "get_state")
	if err != nil {
		return Capture{}, session.failure(fmt.Errorf("Pi get_state after turn failed: %w", err), transcript.String())
	}
	if err := validatePiState(state, session.sessionID, session.provider, session.model, session.reasoning); err != nil {
		return Capture{}, session.failure(err, transcript.String())
	}
	transcript.Write(state)
	transcript.WriteByte('\n')
	events = append(events, append(json.RawMessage(nil), bytes.TrimSpace(state)...))
	return Capture{SessionID: session.sessionID, Transcript: transcript.String(), Stderr: session.stderr.String(), Events: events}, nil
}

func (session *piSession) request(ctx context.Context, command string) ([]byte, error) {
	ctx, cancel := context.WithTimeout(ctx, piRequestTimeout)
	defer cancel()
	id := session.commandID(command)
	message, err := json.Marshal(map[string]any{"id": id, "type": command})
	if err != nil {
		return nil, err
	}
	if _, err := session.stdin.Write(append(message, '\n')); err != nil {
		return nil, err
	}
	for {
		event, err := session.readEvent(ctx)
		if err != nil {
			return nil, err
		}
		var value map[string]any
		if json.Unmarshal(event, &value) != nil {
			return nil, fmt.Errorf("%w: invalid Pi RPC event", ErrProtocol)
		}
		if value["type"] != "response" || value["id"] != id {
			continue
		}
		if success, ok := value["success"].(bool); ok && !success {
			return nil, fmt.Errorf("native response: %s", strings.TrimSpace(string(event)))
		}
		return event, nil
	}
}

func (session *piSession) readEvent(ctx context.Context) ([]byte, error) {
	type result struct {
		data []byte
		err  error
	}
	channel := make(chan result, 1)
	go func() {
		line, err := session.stdout.ReadBytes('\n')
		channel <- result{data: line, err: err}
	}()
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case result := <-channel:
		if result.err != nil {
			return nil, fmt.Errorf("Pi RPC ended before terminal event: %w: %s", result.err, strings.TrimSpace(session.stderr.String()))
		}
		line := bytes.TrimSpace(result.data)
		if len(line) == 0 || !isJSONObject(line) {
			return nil, fmt.Errorf("%w: invalid Pi RPC event", ErrProtocol)
		}
		return append([]byte(nil), line...), nil
	}
}

func (session *piSession) abort() {
	message, err := json.Marshal(map[string]any{"id": session.commandID("abort"), "type": "abort"})
	if err == nil {
		_, _ = session.stdin.Write(append(message, '\n'))
	}
}

func (session *piSession) commandID(prefix string) string {
	id := fmt.Sprintf("%s-%d", prefix, session.nextID)
	session.nextID++
	return id
}

func (session *piSession) Close() error {
	if session.closed {
		return nil
	}
	session.closed = true
	if session.pending {
		session.abort()
	}
	_ = session.stdin.Close()
	done := make(chan error, 1)
	go func() { done <- session.command.Wait() }()
	select {
	case err := <-done:
		if err == nil {
			return nil
		}
		return fmt.Errorf("Pi RPC exited unsuccessfully: %w: %s", err, strings.TrimSpace(session.stderr.String()))
	case <-time.After(2 * time.Second):
		stopOwnedProcessGroup(session.command)
		return nil
	}
}

func piModel(value string) (string, string) {
	provider, model, found := strings.Cut(value, "/")
	if !found {
		return "openai-codex", value
	}
	return provider, model
}

func explicitPiSkills(root string) ([]string, error) {
	if root == "" {
		return nil, errors.New("Pi evaluation skill root is required")
	}
	entries, err := os.ReadDir(root)
	if err != nil {
		return nil, fmt.Errorf("read Pi evaluation skills: %w", err)
	}
	paths := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		paths = append(paths, filepath.Join(root, entry.Name()))
	}
	slices.Sort(paths)
	if len(paths) == 0 {
		return nil, errors.New("Pi evaluation requires at least one supplied skill")
	}
	return paths, nil
}

func validatePiState(event []byte, sessionID, provider, model, reasoning string) error {
	var value map[string]any
	if json.Unmarshal(event, &value) != nil {
		return fmt.Errorf("%w: malformed Pi state response", ErrProtocol)
	}
	state, ok := value["data"].(map[string]any)
	if !ok {
		return fmt.Errorf("%w: Pi state response has no data", ErrProtocol)
	}
	if id, _ := state["sessionId"].(string); id != sessionID {
		return fmt.Errorf("%w: Pi session ID mismatch: %s", ErrProtocol, id)
	}
	level, ok := state["thinkingLevel"].(string)
	if !ok || level != reasoning {
		return fmt.Errorf("%w: Pi thinking level is %s instead of %s", ErrProtocol, level, reasoning)
	}
	streaming, ok := state["isStreaming"].(bool)
	if !ok || streaming {
		return fmt.Errorf("%w: Pi state is streaming during preflight", ErrProtocol)
	}
	compacting, ok := state["isCompacting"].(bool)
	if !ok || compacting {
		return fmt.Errorf("%w: Pi state is compacting during preflight", ErrProtocol)
	}
	pending, ok := state["pendingMessageCount"].(float64)
	if !ok || pending != 0 {
		return fmt.Errorf("%w: Pi has %d pending messages", ErrProtocol, int(pending))
	}
	if sessionFile, _ := state["sessionFile"].(string); sessionFile != "" {
		return fmt.Errorf("%w: Pi ephemeral run has session file %s", ErrProtocol, sessionFile)
	}
	if provider != "" || model != "" {
		resolvedProvider, resolvedModel := piStateModel(state)
		if resolvedProvider != provider || resolvedModel != model {
			return fmt.Errorf("%w: Pi resolved model %s/%s instead of %s/%s", ErrProtocol, resolvedProvider, resolvedModel, provider, model)
		}
	}
	return nil
}

func piStateModel(state map[string]any) (string, string) {
	modelValue, ok := state["model"].(map[string]any)
	if !ok {
		return "", ""
	}
	provider, _ := modelValue["provider"].(string)
	model, _ := modelValue["id"].(string)
	return provider, model
}

func containsString(value []byte, wanted string) bool {
	return strings.Contains(string(value), wanted)
}
