package replay

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

type Options struct {
	Executable  string
	Directory   string
	Environment []string
	Model       string
}

type commandSpec struct {
	executable string
	initial    func(string) []string
	resume     func(string, string) []string
}

type processAdapter struct {
	harnessID HarnessID
	path      string
	directory string
	environ   []string
	spec      commandSpec
	model     string
}

func NewAdapter(harnessID HarnessID, options Options) (Adapter, error) {
	if harnessID == HarnessPi {
		return newPiAdapter(options)
	}
	spec, ok := commandSpecs()[harnessID]
	if !ok {
		return nil, fmt.Errorf("unsupported harness %q", harnessID)
	}
	path, err := resolveExecutable(spec.executable, options.Executable)
	if err != nil {
		return nil, fmt.Errorf("%s executable: %w", harnessID, err)
	}
	return &processAdapter{harnessID: harnessID, path: path, directory: options.Directory, environ: options.Environment, spec: spec, model: options.Model}, nil
}

func (adapter *processAdapter) HarnessID() HarnessID { return adapter.harnessID }

func (adapter *processAdapter) Start(context.Context) (Session, error) {
	return &processSession{adapter: adapter}, nil
}

type processSession struct {
	adapter   *processAdapter
	sessionID string
	pending   *exec.Cmd
	stdout    bytes.Buffer
	stderr    bytes.Buffer
	closed    bool
}

func (session *processSession) SendPrompt(ctx context.Context, prompt string) error {
	if session.closed {
		return errors.New("session is closed")
	}
	if session.pending != nil {
		return errors.New("previous prompt is still running")
	}
	args := session.adapter.spec.initial(prompt)
	if session.sessionID != "" {
		args = session.adapter.spec.resume(session.sessionID, prompt)
	}
	if session.adapter.model != "" {
		args = append(args, "--model", session.adapter.model)
	}
	session.stdout.Reset()
	session.stderr.Reset()
	command := exec.CommandContext(ctx, session.adapter.path, args...)
	command.Dir = session.adapter.directory
	command.Env = mergedEnvironment(session.adapter.environ)
	command.Stdout = &session.stdout
	command.Stderr = &session.stderr
	if err := command.Start(); err != nil {
		return fmt.Errorf("start executable: %w", err)
	}
	session.pending = command
	return nil
}

func (session *processSession) Wait(context.Context) (Capture, error) {
	if session.pending == nil {
		return Capture{}, errors.New("no prompt is running")
	}
	command := session.pending
	session.pending = nil
	if err := command.Wait(); err != nil {
		return Capture{}, fmt.Errorf("harness exited unsuccessfully: %w: %s", err, strings.TrimSpace(session.stderr.String()))
	}
	events, err := parseEvents(session.stdout.Bytes())
	if err != nil {
		if session.adapter.harnessID != HarnessCopilot {
			return Capture{}, err
		}
		events = []json.RawMessage{json.RawMessage(`{"type":"result"}`)}
	}
	if session.sessionID == "" {
		session.sessionID = findSessionID(events)
		if session.sessionID == "" && session.adapter.harnessID == HarnessCopilot {
			session.sessionID = findCopilotSessionID(session.stdout.String() + "\n" + session.stderr.String())
		}
		if session.sessionID == "" {
			return Capture{}, fmt.Errorf("%w: missing session ID", ErrProtocol)
		}
	}
	return Capture{SessionID: session.sessionID, Transcript: session.stdout.String(), Stderr: session.stderr.String(), Events: events}, nil
}

func findCopilotSessionID(output string) string {
	match := regexp.MustCompile(`--resume(?:=|\s+)([A-Za-z0-9-]+)`).FindStringSubmatch(output)
	if len(match) == 2 {
		return match[1]
	}
	return ""
}

func (session *processSession) Close() error {
	session.closed = true
	if session.pending == nil || session.pending.Process == nil {
		return nil
	}
	return session.pending.Process.Kill()
}

func resolveExecutable(defaultName, override string) (string, error) {
	name := defaultName
	if override != "" {
		name = override
	}
	path, err := exec.LookPath(name)
	if err != nil {
		return "", err
	}
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	if info.IsDir() || info.Mode()&0o111 == 0 {
		return "", fmt.Errorf("%q is not executable", path)
	}
	return path, nil
}

func mergedEnvironment(extra []string) []string {
	if extra == nil {
		return os.Environ()
	}
	return append(append([]string{}, os.Environ()...), extra...)
}

func parseEvents(output []byte) ([]json.RawMessage, error) {
	trimmed := bytes.TrimSpace(output)
	if len(trimmed) == 0 {
		return nil, fmt.Errorf("%w: empty output", ErrProtocol)
	}
	var array []json.RawMessage
	if trimmed[0] == '[' {
		if err := json.Unmarshal(trimmed, &array); err != nil || len(array) == 0 {
			return nil, fmt.Errorf("%w: invalid JSON event array", ErrProtocol)
		}
		for _, event := range array {
			if !isJSONObject(event) {
				return nil, fmt.Errorf("%w: event is not a JSON object", ErrProtocol)
			}
		}
		return array, nil
	}
	lines := bytes.Split(trimmed, []byte{'\n'})
	events := make([]json.RawMessage, 0, len(lines))
	for _, line := range lines {
		line = bytes.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if !isJSONObject(line) {
			return nil, fmt.Errorf("%w: invalid JSON object event", ErrProtocol)
		}
		events = append(events, append(json.RawMessage(nil), line...))
	}
	if len(events) == 0 {
		return nil, fmt.Errorf("%w: empty event stream", ErrProtocol)
	}
	return events, nil
}

func isJSONObject(event []byte) bool {
	var value map[string]any
	return json.Unmarshal(event, &value) == nil && value != nil
}

func findSessionID(events []json.RawMessage) string {
	keys := []string{"session_id", "sessionId", "thread_id", "threadId"}
	for _, event := range events {
		var value any
		if json.Unmarshal(event, &value) == nil {
			if id := findStringField(value, keys); id != "" {
				return id
			}
		}
	}
	return ""
}

func findStringField(value any, keys []string) string {
	switch typed := value.(type) {
	case map[string]any:
		for _, key := range keys {
			if found, ok := typed[key].(string); ok && strings.TrimSpace(found) != "" {
				return found
			}
		}
		for _, child := range typed {
			if found := findStringField(child, keys); found != "" {
				return found
			}
		}
	case []any:
		for _, child := range typed {
			if found := findStringField(child, keys); found != "" {
				return found
			}
		}
	}
	return ""
}

func commandSpecs() map[HarnessID]commandSpec {
	return map[HarnessID]commandSpec{
		HarnessCopilot:  {executable: "copilot", initial: promptArgs("-p"), resume: resumeArgs("--resume", "-p")},
		HarnessClaude:   {executable: "claude", initial: promptArgs("-p", "--output-format", "stream-json", "--verbose"), resume: resumeArgs("--resume", "-p", "--output-format", "stream-json", "--verbose")},
		HarnessCodex:    {executable: "codex", initial: codexInitial, resume: codexResume},
		HarnessCursor:   {executable: "cursor-agent", initial: promptArgs("-p", "--output-format", "stream-json"), resume: resumeArgs("--resume", "-p", "--output-format", "stream-json")},
		HarnessGemini:   {executable: "gemini", initial: promptArgs("-p", "--output-format", "stream-json"), resume: resumeArgs("--resume", "-p", "--output-format", "stream-json")},
		HarnessGrok:     {executable: "grok", initial: promptArgs("-p", "--output-format", "json"), resume: resumeArgs("--resume", "-p", "--output-format", "json")},
		HarnessOpenCode: {executable: "opencode", initial: runArgs, resume: runResumeArgs},
		HarnessKilo:     {executable: "kilo", initial: runArgs, resume: runResumeArgs},
	}
}

func promptArgs(prefix string, suffix ...string) func(string) []string {
	return func(prompt string) []string { return append([]string{prefix, prompt}, suffix...) }
}

func resumeArgs(resumeFlag, promptFlag string, suffix ...string) func(string, string) []string {
	return func(sessionID, prompt string) []string {
		return append([]string{resumeFlag, sessionID, promptFlag, prompt}, suffix...)
	}
}

func codexInitial(prompt string) []string { return []string{"exec", "--json", prompt} }
func codexResume(sessionID, prompt string) []string {
	return []string{"exec", "--json", "resume", sessionID, prompt}
}
func runArgs(prompt string) []string { return []string{"run", "--format", "json", prompt} }
func runResumeArgs(sessionID, prompt string) []string {
	return []string{"run", "--session", sessionID, "--format", "json", prompt}
}
