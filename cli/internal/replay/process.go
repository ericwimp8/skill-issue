package replay

import (
	"bytes"
	"context"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type Options struct {
	Executable            string
	Directory             string
	Environment           []string
	CleanEnvironment      bool
	Model                 string
	ModelOverride         bool
	Reasoning             string
	ReasoningOverride     bool
	CodexConfiguration    []string
	CursorPluginDir       string
	ClaudeSettings        string
	ClaudeSkillsRoot      string
	ClaudeWorkspacePrompt string
	PiSkillsRoot          string
	SkillIssueExecutable  string
}

type commandSpec struct {
	executable string
	initial    func(string) []string
	resume     func(string, string) []string
}

type processAdapter struct {
	harnessID             HarnessID
	path                  string
	directory             string
	environ               []string
	spec                  commandSpec
	model                 string
	modelOverride         bool
	reasoning             string
	config                []string
	cleanEnvironment      bool
	cursorPluginDir       string
	claudeSettings        string
	claudeSkillsRoot      string
	claudeWorkspacePrompt string
	piSkillsRoot          string
	skillIssueExecutable  string
	cleanup               func() error
}

const qualifiedOpenCodeVersion = "1.18.4"
const qualifiedKiloVersion = "7.4.11"

func NewAdapter(harnessID HarnessID, options Options) (Adapter, error) {
	if harnessID == HarnessPi {
		return newPiAdapter(options)
	}
	spec, ok := commandSpecs()[harnessID]
	if !ok {
		return nil, fmt.Errorf("unsupported harness %q", harnessID)
	}
	path, err := resolveExecutable(spec.executable, options.Executable)
	if harnessID == HarnessCursor && options.Executable == "" {
		path, err = resolveCursorExecutable()
	}
	if err != nil {
		return nil, fmt.Errorf("%s executable: %w", harnessID, err)
	}
	adapter := &processAdapter{
		harnessID:             harnessID,
		path:                  path,
		directory:             options.Directory,
		environ:               options.Environment,
		spec:                  spec,
		model:                 options.Model,
		modelOverride:         options.ModelOverride,
		reasoning:             options.Reasoning,
		config:                options.CodexConfiguration,
		cleanEnvironment:      options.CleanEnvironment,
		cursorPluginDir:       options.CursorPluginDir,
		claudeSettings:        options.ClaudeSettings,
		claudeSkillsRoot:      options.ClaudeSkillsRoot,
		claudeWorkspacePrompt: options.ClaudeWorkspacePrompt,
		piSkillsRoot:          options.PiSkillsRoot,
		skillIssueExecutable:  options.SkillIssueExecutable,
	}
	if harnessID == HarnessClaude {
		adapter.cleanup = func() error {
			command := exec.CommandContext(context.Background(), adapter.path, "project", "purge", "--yes", adapter.directory)
			configureOwnedProcess(command)
			command.Env = environment(adapter.environ, adapter.cleanEnvironment)
			var output bytes.Buffer
			command.Stdout = &output
			command.Stderr = &output
			err := command.Run()
			stopOwnedProcessGroup(command)
			if err != nil {
				return fmt.Errorf("Claude Code project purge failed: %w: %s", err, strings.TrimSpace(output.String()))
			}
			return nil
		}
	}
	return adapter, nil
}

func (adapter *processAdapter) HarnessID() HarnessID { return adapter.harnessID }

func (adapter *processAdapter) Start(context.Context) (Session, error) {
	session := &processSession{adapter: adapter}
	if adapter.harnessID == HarnessClaude {
		id, err := generatedSessionID()
		if err != nil {
			return nil, err
		}
		session.sessionID = id
	}
	return session, nil
}

type processSession struct {
	adapter   *processAdapter
	sessionID string
	pending   *exec.Cmd
	stdout    bytes.Buffer
	stderr    bytes.Buffer
	closed    bool
	started   bool
}

func (session *processSession) SendPrompt(ctx context.Context, prompt string) error {
	if session.closed {
		return errors.New("session is closed")
	}
	if session.pending != nil {
		return errors.New("previous prompt is still running")
	}
	args := session.adapter.spec.initial(prompt)
	if session.adapter.harnessID == HarnessCursor {
		args = cursorArgs(session.adapter, session.sessionID, prompt)
	} else if session.adapter.harnessID == HarnessClaude {
		if session.started {
			args = claudeArgs(session.adapter, session.sessionID, prompt, true)
		} else {
			args = claudeArgs(session.adapter, session.sessionID, prompt, false)
		}
	} else if session.adapter.harnessID == HarnessOpenCode || session.adapter.harnessID == HarnessKilo {
		if session.adapter.harnessID == HarnessKilo {
			args = kiloArgs(session.adapter, session.sessionID, prompt)
		} else {
			args = openCodeArgs(session.adapter, session.sessionID, prompt)
		}
	} else if session.sessionID != "" {
		args = session.adapter.spec.resume(session.sessionID, prompt)
	}
	if session.adapter.harnessID == HarnessCodex {
		prefix := []string{
			"--cd", session.adapter.directory,
			"--ask-for-approval", "on-request",
			"--sandbox", "workspace-write",
			"--disable", "plugins",
		}
		if session.adapter.model != "" {
			prefix = append(prefix, "--model", session.adapter.model)
		}
		for _, value := range session.adapter.config {
			prefix = append(prefix, "--config", value)
		}
		args = append(prefix, args...)
	}
	session.stdout.Reset()
	session.stderr.Reset()
	command := exec.CommandContext(ctx, session.adapter.path, args...)
	configureOwnedProcess(command)
	command.Dir = session.adapter.directory
	command.Env = environment(session.adapter.environ, session.adapter.cleanEnvironment)
	command.Stdout = &session.stdout
	command.Stderr = &session.stderr
	if err := command.Start(); err != nil {
		return fmt.Errorf("start executable: %w", err)
	}
	session.pending = command
	session.started = true
	return nil
}

func (session *processSession) Wait(context.Context) (Capture, error) {
	if session.pending == nil {
		return Capture{}, errors.New("no prompt is running")
	}
	command := session.pending
	session.pending = nil
	err := command.Wait()
	stopOwnedProcessGroup(command)
	if err != nil {
		return Capture{}, fmt.Errorf("harness exited unsuccessfully: %w: %s", err, strings.TrimSpace(session.stderr.String()))
	}
	events, err := parseEvents(session.stdout.Bytes())
	if err != nil {
		return Capture{}, fmt.Errorf("%s harness produced invalid structured output: %w: %s", session.adapter.harnessID, err, strings.TrimSpace(session.stderr.String()))
	}
	if session.adapter.harnessID == HarnessKilo {
		events = collapseAdjacentExactDuplicateEvents(events)
	}
	requireSessionStart := session.sessionID == ""
	if err := validateHarnessOutput(session.adapter.harnessID, events, session.stderr.String(), requireSessionStart); err != nil {
		return Capture{}, err
	}
	if session.sessionID == "" {
		session.sessionID = findSessionID(events)
		if session.sessionID == "" {
			return Capture{}, fmt.Errorf("%w: missing session ID: %s", ErrProtocol, strings.TrimSpace(session.stderr.String()))
		}
	}
	if err := validateSessionID(session.adapter.harnessID, session.sessionID, events); err != nil {
		return Capture{}, err
	}
	return Capture{SessionID: session.sessionID, Transcript: session.stdout.String(), Stderr: session.stderr.String(), Events: events}, nil
}

func validateHarnessOutput(harnessID HarnessID, events []json.RawMessage, stderr string, requireSessionStart bool) error {
	types := map[string]bool{}
	for _, event := range events {
		var value struct {
			Type    string `json:"type"`
			Subtype string `json:"subtype"`
			Error   any    `json:"error"`
			IsError bool   `json:"is_error"`
		}
		if json.Unmarshal(event, &value) != nil {
			return fmt.Errorf("%w: invalid %s event: %s", ErrProtocol, harnessID, strings.TrimSpace(string(event)))
		}
		types[value.Type+"\x00"+value.Subtype] = true
		types[value.Type+"\x00"] = true
		if value.Type == "turn.failed" || value.Type == "error" || value.Subtype == "error" || (value.Type == "result" && value.IsError) {
			return fmt.Errorf("%s harness reported an error: %s", harnessID, strings.TrimSpace(string(event)))
		}
	}
	switch harnessID {
	case HarnessCodex:
		if requireSessionStart && !types["thread.started\x00"] {
			return fmt.Errorf("%w: Codex output missing thread.started: %s", ErrProtocol, strings.TrimSpace(stderr))
		}
		if !types["turn.completed\x00"] {
			return fmt.Errorf("%w: Codex output missing turn.completed: %s", ErrProtocol, strings.TrimSpace(stderr))
		}
	case HarnessCursor:
		if !types["system\x00init"] || !types["result\x00success"] {
			return fmt.Errorf("%w: Cursor output missing successful system/init and result events: %s", ErrProtocol, strings.TrimSpace(stderr))
		}
	case HarnessClaude:
		if !types["system\x00init"] || !types["result\x00"] {
			return fmt.Errorf("%w: Claude Code output missing successful system/init and result events: %s", ErrProtocol, strings.TrimSpace(stderr))
		}
	case HarnessOpenCode, HarnessKilo:
		if markerFailure, failed := failedSignalMarker(events); failed {
			return fmt.Errorf("%s marker command failed: %s", harnessID, markerFailure)
		}
		if !structuredRunStopped(events) {
			return fmt.Errorf("%w: %s output missing terminal step_finish with reason stop: %s", ErrProtocol, harnessID, strings.TrimSpace(stderr))
		}
	}
	return nil
}

func failedSignalMarker(events []json.RawMessage) (string, bool) {
	for _, event := range events {
		var value struct {
			Type string `json:"type"`
			Part struct {
				Tool  string `json:"tool"`
				State struct {
					Status string `json:"status"`
					Error  string `json:"error"`
					Input  struct {
						Command string `json:"command"`
					} `json:"input"`
				} `json:"state"`
			} `json:"part"`
		}
		if json.Unmarshal(event, &value) == nil && value.Type == "tool_use" && value.Part.Tool == "bash" && value.Part.State.Status == "error" && strings.Contains(value.Part.State.Input.Command, " signal ") {
			return value.Part.State.Input.Command + ": " + value.Part.State.Error, true
		}
	}
	return "", false
}

func validateSessionID(harnessID HarnessID, sessionID string, events []json.RawMessage) error {
	if sessionID == "" {
		return fmt.Errorf("%w: missing %s session ID", ErrProtocol, harnessID)
	}
	found := findSessionID(events)
	if found != "" && found != sessionID {
		return fmt.Errorf("%w: %s session ID changed from %s to %s", ErrProtocol, harnessID, sessionID, found)
	}
	if harnessID == HarnessOpenCode || harnessID == HarnessKilo {
		for _, event := range events {
			var value struct {
				SessionID string `json:"sessionID"`
			}
			if json.Unmarshal(event, &value) == nil && value.SessionID != "" && value.SessionID != sessionID {
				return fmt.Errorf("%w: %s session ID changed from %s to %s", ErrProtocol, harnessID, sessionID, value.SessionID)
			}
		}
	}
	return nil
}

func structuredRunStopped(events []json.RawMessage) bool {
	for _, event := range events {
		var value struct {
			Type string `json:"type"`
			Part struct {
				Reason string `json:"reason"`
			} `json:"part"`
		}
		if json.Unmarshal(event, &value) == nil && value.Type == "step_finish" && value.Part.Reason == "stop" {
			return true
		}
	}
	return false
}

func (session *processSession) Close() error {
	session.closed = true
	var closeErr error
	if session.pending != nil && session.pending.Process != nil {
		stopOwnedProcessGroup(session.pending)
		_ = session.pending.Wait()
		session.pending = nil
	}
	if session.adapter.cleanup != nil {
		cleanupErr := session.adapter.cleanup()
		if closeErr == nil {
			closeErr = cleanupErr
		}
	}
	if session.adapter.harnessID == HarnessOpenCode {
		if session.sessionID == "" {
			session.sessionID = findSessionIDFromPartialOutput(session.stdout.Bytes())
		}
		if session.sessionID != "" {
			cleanupErr := DeleteOpenCodeSession(context.Background(), session.adapter.path, session.adapter.directory, session.adapter.environ, session.adapter.cleanEnvironment, session.sessionID)
			if closeErr == nil {
				closeErr = cleanupErr
			}
		}
	}
	if session.adapter.harnessID == HarnessKilo {
		if session.sessionID == "" {
			session.sessionID = findSessionIDFromPartialOutput(session.stdout.Bytes())
		}
		if session.sessionID != "" {
			cleanupErr := DeleteKiloSession(context.Background(), session.adapter.path, session.adapter.directory, session.adapter.environ, session.adapter.cleanEnvironment, session.sessionID)
			if closeErr == nil {
				closeErr = cleanupErr
			}
		}
	}
	return closeErr
}

func openCodeArgs(adapter *processAdapter, sessionID, prompt string) []string {
	args := []string{"run", "--pure", "--format", "json", "--model", adapter.model, "--variant", adapter.reasoning}
	if sessionID != "" {
		args = append(args, "--session", sessionID)
	}
	return append(args, prompt)
}

func kiloArgs(adapter *processAdapter, sessionID, prompt string) []string {
	args := []string{"run", "--pure", "--format", "json", "--model", adapter.model, "--variant", adapter.reasoning, "--agent", "code", "--dir", adapter.directory}
	if sessionID != "" {
		args = append(args, "--session", sessionID)
	}
	return append(args, prompt)
}

func findSessionIDFromPartialOutput(output []byte) string {
	for _, line := range bytes.Split(output, []byte{'\n'}) {
		var value any
		if json.Unmarshal(bytes.TrimSpace(line), &value) == nil {
			if id := findStringField(value, []string{"sessionID"}); id != "" {
				return id
			}
		}
	}
	return ""
}

func collapseAdjacentExactDuplicateEvents(events []json.RawMessage) []json.RawMessage {
	result := make([]json.RawMessage, 0, len(events))
	for _, event := range events {
		if len(result) > 0 && bytes.Equal(result[len(result)-1], event) {
			continue
		}
		result = append(result, event)
	}
	return result
}

func DeleteKiloSession(ctx context.Context, executable, directory string, env []string, clean bool, sessionID string) error {
	path, err := resolveExecutable("kilo", executable)
	if err != nil {
		return fmt.Errorf("Kilo session executable: %w", err)
	}
	command := exec.CommandContext(ctx, path, "session", "delete", sessionID)
	configureOwnedProcess(command)
	command.Dir = directory
	command.Env = environment(env, clean)
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	err = command.Run()
	stopOwnedProcessGroup(command)
	if err != nil && !strings.Contains(strings.ToLower(output.String()), "not found") {
		return fmt.Errorf("Kilo session deletion failed: %w: %s", err, strings.TrimSpace(output.String()))
	}
	return nil
}

func DeleteOpenCodeSession(ctx context.Context, executable, directory string, env []string, clean bool, sessionID string) error {
	path, err := resolveExecutable("opencode", executable)
	if err != nil {
		return fmt.Errorf("OpenCode session executable: %w", err)
	}
	listed, err := runStatusCommandAt(ctx, path, directory, env, clean, "session", "list", "--pure", "--format", "json")
	if err != nil {
		return fmt.Errorf("OpenCode session listing failed: %w", err)
	}
	if !jsonContainsString([]byte(listed), sessionID) {
		return nil
	}
	command := exec.CommandContext(ctx, path, "session", "delete", sessionID, "--pure")
	configureOwnedProcess(command)
	command.Dir = directory
	command.Env = environment(env, clean)
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	err = command.Run()
	stopOwnedProcessGroup(command)
	if err != nil {
		return fmt.Errorf("OpenCode session deletion failed: %w: %s", err, strings.TrimSpace(output.String()))
	}
	listed, err = runStatusCommandAt(ctx, path, directory, env, clean, "session", "list", "--pure", "--format", "json")
	if err != nil {
		return fmt.Errorf("OpenCode session deletion verification failed: %w", err)
	}
	if jsonContainsString([]byte(listed), sessionID) {
		return errors.New("OpenCode session deletion verification failed: session still exists")
	}
	return nil
}

func CheckOpenCodeSkills(ctx context.Context, executable, directory string, env []string, clean bool, expected []string) error {
	path, err := resolveExecutable("opencode", executable)
	if err != nil {
		return fmt.Errorf("OpenCode skill discovery executable: %w", err)
	}
	output, err := runStatusCommandAt(ctx, path, directory, env, clean, "debug", "skill", "--pure")
	if err != nil {
		return fmt.Errorf("OpenCode skill discovery failed: %w", err)
	}
	var discovered []struct {
		Name string `json:"name"`
	}
	if err := json.Unmarshal([]byte(output), &discovered); err != nil {
		return fmt.Errorf("decode OpenCode skill discovery: %w", err)
	}
	found := make(map[string]bool, len(discovered))
	for _, skill := range discovered {
		found[skill.Name] = true
	}
	for _, name := range expected {
		if !found[name] {
			return fmt.Errorf("OpenCode did not discover installed evaluation skill %q", name)
		}
	}
	return nil
}

func jsonContainsString(data []byte, wanted string) bool {
	var value any
	if json.Unmarshal(data, &value) != nil {
		return false
	}
	return valueContainsString(value, wanted)
}

func valueContainsString(value any, wanted string) bool {
	switch typed := value.(type) {
	case string:
		return typed == wanted
	case []any:
		for _, child := range typed {
			if valueContainsString(child, wanted) {
				return true
			}
		}
	case map[string]any:
		for _, child := range typed {
			if valueContainsString(child, wanted) {
				return true
			}
		}
	}
	return false
}

func cursorArgs(adapter *processAdapter, sessionID, prompt string) []string {
	args := []string{"--disable-auto-update", "--disable-project-configs", "--workspace", adapter.directory, "--plugin-dir", adapter.cursorPluginDir}
	if adapter.modelOverride || adapter.model != "auto" {
		args = append(args, "--model", adapter.model)
	}
	args = append(args, "--trust", "--sandbox", "enabled", "--auto-review")
	if sessionID != "" {
		args = append(args, "--resume", sessionID)
	}
	return append(args, "-p", "--output-format", "stream-json", prompt)
}

func claudeArgs(adapter *processAdapter, sessionID, prompt string, resume bool) []string {
	allowedTools := "Read,Write,Edit,Glob,Grep,Bash(" + adapter.skillIssueExecutable + " signal *)"
	args := []string{"-p", "--setting-sources", "project", "--settings", adapter.claudeSettings, "--strict-mcp-config", "--no-chrome", "--add-dir", adapter.claudeSkillsRoot, "--tools", "Read,Write,Edit,Glob,Grep,Bash", "--allowedTools", allowedTools, "--permission-mode", "dontAsk", "--append-system-prompt", adapter.claudeWorkspacePrompt, "--model", adapter.model, "--effort", adapter.reasoning}
	if resume {
		args = append(args, "--resume", sessionID)
	} else if sessionID != "" {
		args = append(args, "--session-id", sessionID)
	}
	return append(args, "--output-format", "stream-json", "--verbose", prompt)
}

func generatedSessionID() (string, error) {
	data := make([]byte, 16)
	if _, err := rand.Read(data); err != nil {
		return "", fmt.Errorf("generate harness session ID: %w", err)
	}
	return fmt.Sprintf("%x-%x-%x-%x-%x", data[0:4], data[4:6], data[6:8], data[8:10], data[10:]), nil
}

func resolveCursorExecutable() (string, error) {
	for _, name := range []string{"agent", "cursor-agent"} {
		if path, err := resolveExecutable(name, ""); err == nil {
			return path, nil
		}
	}
	return "", errors.New("Cursor executable: neither agent nor cursor-agent was found")
}

func CheckAuthentication(ctx context.Context, harnessID HarnessID, override, model string, env []string, clean bool) error {
	if harnessID != HarnessCodex && harnessID != HarnessCursor && harnessID != HarnessOpenCode && harnessID != HarnessKilo {
		return nil
	}
	var path string
	var err error
	if harnessID == HarnessCursor && override == "" {
		path, err = resolveCursorExecutable()
	} else {
		name := "codex"
		if harnessID == HarnessCursor {
			name = "cursor-agent"
		} else if harnessID == HarnessOpenCode {
			name = "opencode"
		} else if harnessID == HarnessKilo {
			name = "kilo"
		}
		path, err = resolveExecutable(name, override)
	}
	if err != nil {
		return fmt.Errorf("%s authentication executable: %w", harnessID, err)
	}
	if harnessID == HarnessOpenCode {
		return checkOpenCodeAuthentication(ctx, path, model, env, clean)
	}
	if harnessID == HarnessKilo {
		return checkKiloAuthentication(ctx, path, model, env, clean)
	}
	command := exec.CommandContext(ctx, path, "login", "status")
	if harnessID == HarnessCursor {
		command = exec.CommandContext(ctx, path, "status")
	}
	command.Env = environment(env, clean)
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	if err := command.Run(); err != nil {
		return fmt.Errorf("%s authentication status failed: %w: %s", harnessID, err, strings.TrimSpace(output.String()))
	}
	return nil
}

func checkKiloAuthentication(ctx context.Context, path, model string, env []string, clean bool) error {
	version, err := runStatusCommand(ctx, path, env, clean, "--version")
	if err != nil {
		return fmt.Errorf("Kilo version check failed: %w", err)
	}
	if strings.TrimSpace(version) != qualifiedKiloVersion {
		return fmt.Errorf("Kilo version %q is unsupported; install qualified version %s", strings.TrimSpace(version), qualifiedKiloVersion)
	}
	provider, _, found := strings.Cut(model, "/")
	if !found || provider == "" {
		return errors.New("Kilo model must use provider/model format")
	}
	auth, err := runStatusCommand(ctx, path, env, clean, "auth", "list")
	if err != nil {
		return fmt.Errorf("Kilo authentication status failed: %w", err)
	}
	if !strings.Contains(strings.ToLower(auth), strings.ToLower(provider)) {
		return fmt.Errorf("Kilo provider %q is not authenticated; run kilo auth login", provider)
	}
	models, err := runStatusCommand(ctx, path, env, clean, "models", provider)
	if err != nil {
		return fmt.Errorf("Kilo model availability failed: %w", err)
	}
	if !containsLine(models, model) {
		return fmt.Errorf("Kilo model %q is unavailable for provider %q", model, provider)
	}
	return nil
}

func checkOpenCodeAuthentication(ctx context.Context, path, model string, env []string, clean bool) error {
	version, err := runStatusCommand(ctx, path, env, clean, "--version")
	if err != nil {
		return fmt.Errorf("OpenCode version check failed: %w", err)
	}
	if strings.TrimSpace(version) != qualifiedOpenCodeVersion {
		return fmt.Errorf("OpenCode version %q is unsupported; install qualified version %s", strings.TrimSpace(version), qualifiedOpenCodeVersion)
	}
	provider, _, found := strings.Cut(model, "/")
	if !found || provider == "" {
		return errors.New("OpenCode model must use provider/model format")
	}
	auth, err := runStatusCommand(ctx, path, env, clean, "auth", "list", "--pure")
	if err != nil {
		return fmt.Errorf("OpenCode authentication status failed: %w", err)
	}
	if !strings.Contains(strings.ToLower(auth), strings.ToLower(provider)) {
		return fmt.Errorf("OpenCode provider %q is not authenticated; run opencode auth login", provider)
	}
	models, err := runStatusCommand(ctx, path, env, clean, "models", provider, "--pure")
	if err != nil {
		return fmt.Errorf("OpenCode model availability failed: %w", err)
	}
	if !containsLine(models, model) {
		return fmt.Errorf("OpenCode model %q is unavailable for provider %q", model, provider)
	}
	return nil
}

func runStatusCommand(ctx context.Context, path string, env []string, clean bool, args ...string) (string, error) {
	return runStatusCommandAt(ctx, path, "", env, clean, args...)
}

func runStatusCommandAt(ctx context.Context, path, directory string, env []string, clean bool, args ...string) (string, error) {
	command := exec.CommandContext(ctx, path, args...)
	command.Dir = directory
	command.Env = environment(env, clean)
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	if err := command.Run(); err != nil {
		return "", fmt.Errorf("%w: %s", err, strings.TrimSpace(output.String()))
	}
	return output.String(), nil
}

func containsLine(output, wanted string) bool {
	for _, line := range strings.Split(output, "\n") {
		if strings.TrimSpace(line) == wanted {
			return true
		}
	}
	return false
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
	overrides := make(map[string]string, len(extra))
	for _, entry := range extra {
		key, _, found := strings.Cut(entry, "=")
		if found {
			overrides[key] = entry
		}
	}
	environment := make([]string, 0, len(os.Environ())+len(extra))
	for _, entry := range os.Environ() {
		key, _, found := strings.Cut(entry, "=")
		if found {
			if _, replaced := overrides[key]; replaced {
				continue
			}
		}
		environment = append(environment, entry)
	}
	return append(environment, extra...)
}

func environment(extra []string, clean bool) []string {
	if !clean {
		return mergedEnvironment(extra)
	}
	return append([]string(nil), extra...)
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
	keys := []string{"session_id", "sessionId", "sessionID", "thread_id", "threadId"}
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
		HarnessClaude:   {executable: "claude", initial: promptArgs("-p", "--output-format", "stream-json", "--verbose"), resume: resumeArgs("--resume", "-p", "--output-format", "stream-json", "--verbose")},
		HarnessCodex:    {executable: "codex", initial: codexInitial, resume: codexResume},
		HarnessCursor:   {executable: "cursor-agent", initial: promptArgs("-p", "--output-format", "stream-json"), resume: resumeArgs("--resume", "-p", "--output-format", "stream-json")},
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

func codexInitial(prompt string) []string {
	return []string{"exec", "--ignore-user-config", "--ignore-rules", "--json", prompt}
}
func codexResume(sessionID, prompt string) []string {
	return []string{"exec", "resume", "--ignore-user-config", "--ignore-rules", "--json", sessionID, prompt}
}
func runArgs(prompt string) []string { return []string{"run", "--format", "json", prompt} }
func runResumeArgs(sessionID, prompt string) []string {
	return []string{"run", "--session", sessionID, "--format", "json", prompt}
}
