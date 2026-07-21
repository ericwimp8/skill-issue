package evaluation

import (
	"encoding/json"
	"fmt"
	"net"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/ericwimp8/skill-issue/cli/internal/replay"
)

type artifactSanitizerConfig struct {
	Workspace    string
	OutputRoot   string
	StateRoot    string
	RuntimeRoot  string
	CLIPath      string
	SignalTokens []string
}

type artifactReplacement struct {
	value       string
	replacement string
	identity    bool
}

type artifactPattern struct {
	pattern     *regexp.Regexp
	replacement string
}

type artifactSanitizer struct {
	replacements []artifactReplacement
	patterns     []artifactPattern
}

type TranscriptArtifact struct {
	SchemaVersion int              `json:"schema_version"`
	Turns         []TranscriptTurn `json:"turns"`
}

type TranscriptTurn struct {
	TurnID    string `json:"turn_id"`
	User      string `json:"user"`
	Assistant string `json:"assistant"`
}

func newArtifactSanitizer(config artifactSanitizerConfig) (artifactSanitizer, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return artifactSanitizer{}, fmt.Errorf("resolve artifact privacy home: %w", err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		return artifactSanitizer{}, fmt.Errorf("resolve artifact privacy hostname: %w", err)
	}

	sanitizer := artifactSanitizer{patterns: defaultArtifactPatterns()}
	sanitizer.addPath(config.CLIPath, "[skill-issue-cli]")
	sanitizer.addPath(config.StateRoot, "[evaluation-state]")
	sanitizer.addPath(config.OutputRoot, "[evaluation-output]")
	sanitizer.addPath(config.RuntimeRoot, "[runtime]")
	sanitizer.addPath(config.Workspace, "[workspace]")
	sanitizer.addPath(home, "[home]")
	sanitizer.addPath(os.TempDir(), "[temporary-directory]")

	for _, key := range []string{"USER", "LOGNAME", "USERNAME"} {
		sanitizer.addIdentity(os.Getenv(key), "[user]")
	}
	for _, value := range []string{hostname, os.Getenv("HOSTNAME"), os.Getenv("COMPUTERNAME")} {
		sanitizer.addIdentity(value, "[host]")
	}
	if current, currentErr := user.Current(); currentErr == nil {
		sanitizer.addIdentity(current.Username, "[user]")
		sanitizer.addIdentity(current.Name, "[user]")
		sanitizer.addPath(current.HomeDir, "[home]")
	}
	for _, token := range config.SignalTokens {
		sanitizer.addIdentity(token, "[signal-token]")
	}

	slices.SortStableFunc(sanitizer.replacements, func(left, right artifactReplacement) int {
		return len(right.value) - len(left.value)
	})
	return sanitizer, nil
}

func defaultArtifactPatterns() []artifactPattern {
	return []artifactPattern{
		{regexp.MustCompile(`(?s)-----BEGIN (?:RSA |EC |OPENSSH )?PRIVATE KEY-----.*?-----END (?:RSA |EC |OPENSSH )?PRIVATE KEY-----`), "[private-key]"},
		{regexp.MustCompile(`(?i)(\bauthorization\s*[:=]\s*["']?(?:bearer|basic)\s+)[A-Za-z0-9._~+/=-]{8,}`), `${1}[authorization]`},
		{regexp.MustCompile(`(?i)(\b(?:api[_-]?key|access[_-]?token|auth[_-]?token|password|passwd|secret|client[_-]?secret)\b["']?\s*[:=]\s*)(?:"[^"\r\n]*"|'[^'\r\n]*'|[^\s,;#]+)`), `${1}[secret]`},
		{regexp.MustCompile(`(?i)(\bhttps?://)[^\s/@:]+:[^\s/@]+@`), `${1}[url-credentials]@`},
		{regexp.MustCompile(`\b(?:sk-[A-Za-z0-9_-]{16,}|ghp_[A-Za-z0-9]{16,}|xox[baprs]-[A-Za-z0-9-]{16,})\b`), "[token]"},
		{regexp.MustCompile(`\beyJ[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\.[A-Za-z0-9_-]{8,}\b`), "[jwt]"},
		{regexp.MustCompile(`(?i)\b[A-Z0-9._%+-]+@[A-Z0-9.-]+\.[A-Z]{2,}\b`), "[email]"},
		{regexp.MustCompile(`(/Users/|/home/)[^/\\\s]+`), `${1}[user]`},
		{regexp.MustCompile(`(?i)(C:\\Users\\)[^\\/\s]+`), `${1}[user]`},
	}
}

func (sanitizer *artifactSanitizer) addPath(value, replacement string) {
	if strings.TrimSpace(value) == "" {
		return
	}
	value = filepath.Clean(value)
	if value == "." || value == string(filepath.Separator) {
		return
	}
	sanitizer.addReplacement(value, replacement, false)
	if slash := filepath.ToSlash(value); slash != value {
		sanitizer.addReplacement(slash, replacement, false)
	}
	if escapedSlash := strings.ReplaceAll(value, "/", `\/`); escapedSlash != value {
		sanitizer.addReplacement(escapedSlash, replacement, false)
	}
}

func (sanitizer *artifactSanitizer) addIdentity(value, replacement string) {
	value = strings.TrimSpace(value)
	if value == "" || genericIdentity(value) {
		return
	}
	sanitizer.addReplacement(value, replacement, true)
}

func (sanitizer *artifactSanitizer) addReplacement(value, replacement string, identity bool) {
	if value == "" {
		return
	}
	if slices.ContainsFunc(sanitizer.replacements, func(existing artifactReplacement) bool { return existing.value == value }) {
		return
	}
	sanitizer.replacements = append(sanitizer.replacements, artifactReplacement{value: value, replacement: replacement, identity: identity})
	escaped, err := json.Marshal(value)
	if err == nil {
		encoded := string(escaped[1 : len(escaped)-1])
		if encoded != value {
			sanitizer.replacements = append(sanitizer.replacements, artifactReplacement{value: encoded, replacement: replacement, identity: identity})
		}
	}
}

func genericIdentity(value string) bool {
	switch strings.ToLower(value) {
	case "root", "user", "admin", "administrator", "unknown", "localhost":
		return true
	default:
		return false
	}
}

func newTranscriptArtifact(result replay.Result) (TranscriptArtifact, error) {
	captures := make(map[string]replay.Capture, len(result.Turns))
	for _, turn := range result.Turns {
		captures[turn.TurnID] = turn.Capture
	}
	artifact := TranscriptArtifact{SchemaVersion: 1, Turns: make([]TranscriptTurn, 0, len(result.Scenario.Turns))}
	for _, turn := range result.Scenario.Turns {
		capture, ok := captures[turn.ID]
		if !ok {
			return TranscriptArtifact{}, fmt.Errorf("transcript capture missing turn %q", turn.ID)
		}
		assistant, err := extractAssistantResponse(result.HarnessID, capture.Events)
		if err != nil {
			return TranscriptArtifact{}, fmt.Errorf("extract assistant response for turn %q: %w", turn.ID, err)
		}
		artifact.Turns = append(artifact.Turns, TranscriptTurn{TurnID: turn.ID, User: turn.Prompt, Assistant: assistant})
	}
	return artifact, nil
}

func extractAssistantResponse(harnessID replay.HarnessID, events []json.RawMessage) (string, error) {
	responses := make([]string, 0, 4)
	for _, event := range events {
		var value struct {
			Type string `json:"type"`
			Item struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"item"`
			Part struct {
				Type string `json:"type"`
				Text string `json:"text"`
			} `json:"part"`
			Message struct {
				Role    string `json:"role"`
				Content []struct {
					Type string `json:"type"`
					Text string `json:"text"`
				} `json:"content"`
			} `json:"message"`
		}
		if err := json.Unmarshal(event, &value); err != nil {
			return "", err
		}
		switch harnessID {
		case replay.HarnessCodex:
			if value.Type == "item.completed" && value.Item.Type == "agent_message" {
				responses = appendConversationText(responses, value.Item.Text)
			}
		case replay.HarnessCursor, replay.HarnessClaude:
			if value.Type == "assistant" && value.Message.Role == "assistant" {
				responses = appendMessageContent(responses, value.Message.Content)
			}
		case replay.HarnessOpenCode:
			if value.Type == "text" && value.Part.Type == "text" {
				responses = appendConversationText(responses, value.Part.Text)
			}
		case replay.HarnessPi:
			if value.Type == "message_end" && value.Message.Role == "assistant" {
				responses = appendMessageContent(responses, value.Message.Content)
			}
		default:
			return "", fmt.Errorf("unsupported transcript harness %q", harnessID)
		}
	}
	return strings.Join(responses, "\n\n"), nil
}

func appendMessageContent(responses []string, content []struct {
	Type string `json:"type"`
	Text string `json:"text"`
}) []string {
	for _, item := range content {
		if item.Type == "text" || item.Type == "output_text" {
			responses = appendConversationText(responses, item.Text)
		}
	}
	return responses
}

func appendConversationText(responses []string, value string) []string {
	value = strings.TrimSpace(value)
	if value == "" {
		return responses
	}
	return append(responses, value)
}

func (sanitizer artifactSanitizer) sanitizeTranscript(artifact *TranscriptArtifact) {
	for index := range artifact.Turns {
		artifact.Turns[index].User = sanitizer.sanitizeText(artifact.Turns[index].User)
		artifact.Turns[index].Assistant = sanitizer.sanitizeText(artifact.Turns[index].Assistant)
	}
}

func (sanitizer artifactSanitizer) sanitizeFailure(record *FailureRecord) {
	record.Error = sanitizer.sanitizeText(record.Error)
}

func (sanitizer artifactSanitizer) sanitizeText(value string) string {
	for _, replacement := range sanitizer.replacements {
		value = replaceArtifactValue(value, replacement)
	}
	for _, pattern := range sanitizer.patterns {
		value = pattern.pattern.ReplaceAllString(value, pattern.replacement)
	}
	value = artifactIPv4Pattern.ReplaceAllStringFunc(value, func(candidate string) string {
		if net.ParseIP(candidate) == nil {
			return candidate
		}
		return "[ip-address]"
	})
	return value
}

var artifactIPv4Pattern = regexp.MustCompile(`(?:\d{1,3}\.){3}\d{1,3}`)

func replaceArtifactValue(value string, replacement artifactReplacement) string {
	start := 0
	for {
		index := strings.Index(value[start:], replacement.value)
		if index < 0 {
			return value
		}
		index += start
		end := index + len(replacement.value)
		if artifactBoundary(value, index, end, replacement.identity) {
			value = value[:index] + replacement.replacement + value[end:]
			start = index + len(replacement.replacement)
			continue
		}
		start = end
	}
}

func artifactBoundary(value string, start, end int, identity bool) bool {
	if start > 0 {
		previous, _ := utf8.DecodeLastRuneInString(value[:start])
		if artifactWordRune(previous, identity) {
			return false
		}
	}
	if end < len(value) {
		next, _ := utf8.DecodeRuneInString(value[end:])
		if artifactWordRune(next, identity) {
			return false
		}
	}
	return true
}

func artifactWordRune(value rune, identity bool) bool {
	if unicode.IsLetter(value) || unicode.IsDigit(value) || value == '_' || value == '-' {
		return true
	}
	return !identity && value == '.'
}
