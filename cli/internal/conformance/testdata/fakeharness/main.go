// Command fakeharness emulates the native CLI protocol of every supported
// harness for the conformance suite. The behavior is selected by a
// fake-mode.json file beside the invoked executable, because several harness
// runtimes launch with a fully controlled environment that forwards no test
// variables. In happy mode the fake reads the instrumented skills the
// evaluator installed and executes (or, for Codex, reports) their real signal
// commands, so a conformance run exercises the complete attribution pipeline
// without any vendor credentials.
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
)

type settings struct {
	Harness string `json:"harness"`
	Mode    string `json:"mode"`
	Version string `json:"version,omitempty"`
}

const (
	modeHappy             = "happy"
	modeDieOnResume       = "die-on-resume"
	modeSessionChange     = "session-change"
	modeMissingCompletion = "missing-completion"
	modeConfigReject      = "config-reject"
	modeAgentError        = "agent-error"
	modeMarkerFailure     = "marker-failure"
	modeMarkerRecovered   = "marker-recovered"
	modeHideSkills        = "hide-skills"
)

var defaultVersions = map[string]string{
	"claude-code": "2.1.205 (Claude Code)",
	"codex":       "codex-cli 0.144.6",
	"cursor":      "2026.07.16-899851b",
	"opencode":    "1.18.4",
	"pi":          "0.80.10",
}

func main() {
	directory, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		fail("resolve fake harness directory: %v", err)
	}
	data, err := os.ReadFile(filepath.Join(directory, "fake-mode.json"))
	if err != nil {
		fail("read fake-mode.json: %v", err)
	}
	var configuration settings
	if err := json.Unmarshal(data, &configuration); err != nil {
		fail("decode fake-mode.json: %v", err)
	}
	if slices.Contains(os.Args[1:], "--version") {
		version := configuration.Version
		if version == "" {
			version = defaultVersions[configuration.Harness]
		}
		fmt.Println(version)
		return
	}
	switch configuration.Harness {
	case "codex":
		runCodex(configuration)
	case "claude-code":
		runClaude(configuration)
	case "cursor":
		runCursor(configuration)
	case "opencode":
		runStructured(configuration, directory)
	case "pi":
		runPi(configuration)
	default:
		fail("unsupported fake harness %q", configuration.Harness)
	}
}

func runCodex(configuration settings) {
	arguments := os.Args[1:]
	if slices.Contains(arguments, "login") && slices.Contains(arguments, "status") {
		os.Exit(0)
	}
	if configuration.Mode == modeConfigReject {
		for index, argument := range arguments {
			if argument == "--config" && index+1 < len(arguments) && strings.HasPrefix(arguments[index+1], "features.multi_agent=") {
				fail("Error loading config.toml: invalid type: boolean `false`, expected struct FakeFeatureToml in `features`")
			}
		}
	}
	workspace := argumentValue("--cd")
	if workspace == "" {
		fail("codex fake requires --cd")
	}
	resumed := slices.Contains(arguments, "resume")
	if resumed {
		switch configuration.Mode {
		case modeDieOnResume:
			fmt.Println(`{"type":"item.started"}`)
			fail("fake codex died mid-turn")
		case modeSessionChange:
			emit(map[string]any{"type": "thread.started", "thread_id": "thread-changed"})
		}
	} else {
		emit(map[string]any{"type": "thread.started", "thread_id": "thread-0001"})
	}
	for _, token := range captureTokens(filepath.Join(workspace, ".agents", "skills")) {
		emit(map[string]any{"type": "item.completed", "item": map[string]any{"type": "command_execution", "command": fmt.Sprintf("echo %q", token)}})
	}
	if configuration.Mode == modeMissingCompletion {
		return
	}
	emit(map[string]any{"type": "turn.completed"})
}

func runClaude(configuration settings) {
	if len(os.Args) > 1 && (os.Args[1] == "project" || os.Args[1] == "auth") {
		os.Exit(0)
	}
	skillsBase := argumentValue("--add-dir")
	if skillsBase == "" {
		fail("claude fake requires --add-dir")
	}
	sessionID := argumentValue("--session-id")
	resumed := false
	if resume := argumentValue("--resume"); resume != "" {
		sessionID = resume
		resumed = true
	}
	if sessionID == "" {
		fail("claude fake requires a session ID")
	}
	if resumed && configuration.Mode == modeDieOnResume {
		fmt.Println(`{"type":"assistant"}`)
		fail("fake claude died mid-turn")
	}
	if resumed && configuration.Mode == modeSessionChange {
		sessionID = "claude-session-changed"
	}
	skillRoot := filepath.Join(skillsBase, ".claude", "skills")
	visible := []string{}
	if configuration.Mode == modeHideSkills {
		visible = append(visible, "unrelated-skill")
	} else {
		for _, entrypoint := range skillEntrypoints(skillRoot) {
			visible = append(visible, filepath.Base(filepath.Dir(entrypoint)))
		}
		if configuration.Mode == modeMarkerFailure {
			toolUseID := "claude-marker-failure"
			command := firstSignalCommand(skillRoot)
			emit(map[string]any{"type": "assistant", "message": map[string]any{"content": []map[string]any{{"type": "tool_use", "name": "Bash", "id": toolUseID, "input": map[string]any{"command": command}}}}})
			emit(map[string]any{"type": "user", "message": map[string]any{"content": []map[string]any{{"type": "tool_result", "tool_use_id": toolUseID, "content": "permission denied by rule", "is_error": true}}}, "tool_use_result": map[string]any{"stderr": "permission denied by rule"}})
		} else {
			executeSignals(skillRoot)
		}
	}
	emit(map[string]any{"type": "system", "subtype": "init", "session_id": sessionID, "skills": visible})
	if configuration.Mode == modeMissingCompletion {
		return
	}
	emit(map[string]any{"type": "result", "session_id": sessionID})
}

func runCursor(configuration settings) {
	if len(os.Args) > 1 && os.Args[1] == "status" {
		os.Exit(0)
	}
	plugin := argumentValue("--plugin-dir")
	if plugin == "" {
		fail("cursor fake requires --plugin-dir")
	}
	sessionID := "cursor-session-0001"
	resumed := false
	if resume := argumentValue("--resume"); resume != "" {
		sessionID = resume
		resumed = true
	}
	if resumed && configuration.Mode == modeDieOnResume {
		fmt.Println(`{"type":"assistant"}`)
		fail("fake cursor died mid-turn")
	}
	if resumed && configuration.Mode == modeSessionChange {
		sessionID = "cursor-session-changed"
	}
	executeSignals(filepath.Join(plugin, "skills"))
	emit(map[string]any{"type": "system", "subtype": "init", "session_id": sessionID})
	if configuration.Mode == modeMissingCompletion {
		return
	}
	emit(map[string]any{"type": "result", "subtype": "success", "session_id": sessionID})
}

func runStructured(configuration settings, directory string) {
	arguments := os.Args[1:]
	if len(arguments) == 0 {
		fail("%s fake requires a subcommand", configuration.Harness)
	}
	version := configuration.Version
	if version == "" {
		version = "1.18.4"
	}
	switch arguments[0] {
	case "--version":
		fmt.Println(version)
	case "auth":
		fmt.Println("openai")
	case "models":
		fmt.Println("openai/gpt-5.6-sol")
	case "debug":
		if configuration.Mode == modeHideSkills {
			fmt.Println("[]")
			return
		}
		listSkills(structuredSkillRoot())
	case "session":
		handleStructuredSession(directory, arguments)
	case "run":
		runStructuredTurn(configuration, directory)
	default:
		fail("unsupported %s subcommand: %v", configuration.Harness, arguments)
	}
}

func listSkills(root string) {
	skills := make([]map[string]any, 0, 8)
	for _, entrypoint := range skillEntrypoints(root) {
		skills = append(skills, map[string]any{"name": filepath.Base(filepath.Dir(entrypoint))})
	}
	data, err := json.Marshal(skills)
	if err != nil {
		fail("encode skill listing: %v", err)
	}
	fmt.Println(string(data))
}

func handleStructuredSession(directory string, arguments []string) {
	if len(arguments) < 2 {
		fail("session subcommand requires an action")
	}
	sessionsPath := filepath.Join(directory, "sessions.json")
	switch arguments[1] {
	case "list":
		sessions := readSessions(sessionsPath)
		listed := make([]map[string]any, 0, len(sessions))
		for _, id := range sessions {
			listed = append(listed, map[string]any{"id": id})
		}
		data, err := json.Marshal(listed)
		if err != nil {
			fail("encode session listing: %v", err)
		}
		fmt.Println(string(data))
	case "delete":
		if len(arguments) < 3 {
			fail("session delete requires an ID")
		}
		id := arguments[2]
		remaining := slices.DeleteFunc(readSessions(sessionsPath), func(existing string) bool { return existing == id })
		writeSessions(sessionsPath, remaining)
	default:
		fail("unsupported session action %q", arguments[1])
	}
}

func runStructuredTurn(configuration settings, directory string) {
	name := configuration.Harness
	sessionID := argumentValue("--session")
	resumed := sessionID != ""
	if !resumed {
		sessionID = name + "-session-0001"
		sessionsPath := filepath.Join(directory, "sessions.json")
		writeSessions(sessionsPath, append(readSessions(sessionsPath), sessionID))
	}
	if resumed && configuration.Mode == modeDieOnResume {
		fmt.Println(`{"type":"step_start"}`)
		fail("fake %s died mid-turn", name)
	}
	if resumed && configuration.Mode == modeSessionChange {
		sessionID = name + "-session-changed"
	}
	if configuration.Mode != modeMarkerFailure {
		executeSignals(structuredSkillRoot())
	}
	emit(map[string]any{"type": "step_start", "sessionID": sessionID})
	if configuration.Mode == modeMarkerFailure || configuration.Mode == modeMarkerRecovered {
		// A compound command chaining the real signal with another action,
		// denied by the deny-first policy — the live failure shape.
		command := firstSignalCommand(structuredSkillRoot()) + ` && mkdir -p "plans/fake"`
		emit(map[string]any{"type": "tool_use", "sessionID": sessionID, "part": map[string]any{"tool": "bash", "state": map[string]any{"status": "error", "error": "permission denied by rule", "input": map[string]any{"command": command}}}})
	}
	reason := "stop"
	if configuration.Mode == modeMissingCompletion {
		reason = "tool-calls"
	}
	emit(map[string]any{"type": "step_finish", "sessionID": sessionID, "part": map[string]any{"reason": reason}})
}

func structuredSkillRoot() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		fail("XDG_CONFIG_HOME is not set")
	}
	return filepath.Join(configHome, "opencode", "skills")
}

func runPi(configuration settings) {
	sessionID := argumentValue("--session-id")
	thinking := argumentValue("--thinking")
	provider := argumentValue("--provider")
	model := argumentValue("--model")
	var skills []string
	for index, argument := range os.Args {
		if argument == "--skill" && index+1 < len(os.Args) {
			skills = append(skills, os.Args[index+1])
		}
	}
	prompts := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var message map[string]any
		if json.Unmarshal([]byte(line), &message) != nil {
			continue
		}
		id, _ := message["id"].(string)
		switch message["type"] {
		case "get_state":
			reported := sessionID
			if configuration.Mode == modeSessionChange && prompts >= 2 {
				reported = "pi-session-changed"
			}
			emit(map[string]any{"id": id, "type": "response", "command": "get_state", "success": true, "data": map[string]any{
				"sessionId":           reported,
				"thinkingLevel":       thinking,
				"isStreaming":         false,
				"isCompacting":        false,
				"pendingMessageCount": 0,
				"model":               map[string]any{"provider": provider, "id": model},
			}})
		case "get_commands":
			emit(map[string]any{"id": id, "type": "response", "command": "get_commands", "success": true, "data": map[string]any{"commands": skills}})
		case "prompt":
			prompts++
			if configuration.Mode == modeDieOnResume && prompts == 2 {
				fail("fake pi died mid-turn")
			}
			if configuration.Mode == modeMarkerFailure {
				toolCallID := "pi-marker-failure"
				command := firstSignalCommand(filepath.Dir(skills[0]))
				emit(map[string]any{"type": "tool_execution_start", "toolName": "bash", "toolCallId": toolCallID, "args": map[string]any{"command": command}})
				emit(map[string]any{"type": "tool_execution_end", "toolName": "bash", "toolCallId": toolCallID, "isError": true, "result": map[string]any{"content": []map[string]any{{"type": "text", "text": "permission denied by rule"}}}})
			} else {
				for _, skill := range skills {
					executeSignalEntrypoint(filepath.Join(skill, "SKILL.md"))
				}
			}
			emit(map[string]any{"id": id, "type": "response", "command": "prompt", "success": true})
			if configuration.Mode == modeAgentError {
				emit(map[string]any{"type": "agent_end", "stopReason": "error"})
			}
			emit(map[string]any{"type": "agent_settled"})
		case "abort":
			emit(map[string]any{"id": id, "type": "response", "command": "abort", "success": true})
		}
	}
}

var signalInstruction = regexp.MustCompile(`Run "([^"]+)" signal "([^"]+)" "([^"]+)", then continue normally\.`)
var captureInstruction = regexp.MustCompile(`Run echo "([^"]+)", then continue normally\.`)

func skillEntrypoints(root string) []string {
	entries, err := os.ReadDir(root)
	if err != nil {
		fail("read skill root %s: %v", root, err)
	}
	entrypoints := make([]string, 0, len(entries))
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		entrypoint := filepath.Join(root, entry.Name(), "SKILL.md")
		if _, err := os.Stat(entrypoint); err == nil {
			entrypoints = append(entrypoints, entrypoint)
		}
	}
	slices.Sort(entrypoints)
	return entrypoints
}

func executeSignals(root string) {
	for _, entrypoint := range skillEntrypoints(root) {
		executeSignalEntrypoint(entrypoint)
	}
}

func executeSignalEntrypoint(entrypoint string) {
	data, err := os.ReadFile(entrypoint)
	if err != nil {
		fail("read instrumented skill %s: %v", entrypoint, err)
	}
	match := signalInstruction.FindSubmatch(data)
	if match == nil {
		fail("no signal instruction in %s", entrypoint)
	}
	command := exec.Command(string(match[1]), "signal", string(match[2]), string(match[3]))
	output, err := command.CombinedOutput()
	if err != nil {
		fail("signal command failed: %v: %s", err, output)
	}
}

func firstSignalCommand(root string) string {
	entrypoints := skillEntrypoints(root)
	if len(entrypoints) == 0 {
		fail("no skills available for a signal command")
	}
	data, err := os.ReadFile(entrypoints[0])
	if err != nil {
		fail("read instrumented skill %s: %v", entrypoints[0], err)
	}
	match := signalInstruction.FindSubmatch(data)
	if match == nil {
		fail("no signal instruction in %s", entrypoints[0])
	}
	return fmt.Sprintf("%q signal %q %q", match[1], match[2], match[3])
}

func captureTokens(root string) []string {
	tokens := make([]string, 0, 8)
	for _, entrypoint := range skillEntrypoints(root) {
		data, err := os.ReadFile(entrypoint)
		if err != nil {
			fail("read instrumented skill %s: %v", entrypoint, err)
		}
		match := captureInstruction.FindSubmatch(data)
		if match == nil {
			fail("no capture instruction in %s", entrypoint)
		}
		tokens = append(tokens, string(match[1]))
	}
	return tokens
}

func readSessions(path string) []string {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil
	}
	var sessions []string
	if json.Unmarshal(data, &sessions) != nil {
		return nil
	}
	return sessions
}

func writeSessions(path string, sessions []string) {
	data, err := json.Marshal(sessions)
	if err != nil {
		fail("encode sessions: %v", err)
	}
	if err := os.WriteFile(path, data, 0o600); err != nil {
		fail("write sessions: %v", err)
	}
}

func argumentValue(name string) string {
	for index, argument := range os.Args {
		if argument == name && index+1 < len(os.Args) {
			return os.Args[index+1]
		}
	}
	return ""
}

func emit(value map[string]any) {
	data, err := json.Marshal(value)
	if err != nil {
		fail("encode event: %v", err)
	}
	fmt.Println(string(data))
}

func fail(format string, arguments ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", arguments...)
	os.Exit(1)
}
