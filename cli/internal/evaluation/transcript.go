package evaluation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/user"
	"path/filepath"
	"slices"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/ericwimp8/skill-issue/cli/internal/replay"
)

type transcriptSanitizerConfig struct {
	Workspace   string
	OutputRoot  string
	StateRoot   string
	RuntimeRoot string
	CLIPath     string
}

type transcriptReplacement struct {
	value       string
	replacement string
	identity    bool
}

type transcriptSanitizer struct {
	replacements []transcriptReplacement
}

func newTranscriptSanitizer(config transcriptSanitizerConfig) (transcriptSanitizer, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return transcriptSanitizer{}, fmt.Errorf("resolve transcript privacy home: %w", err)
	}
	hostname, err := os.Hostname()
	if err != nil {
		return transcriptSanitizer{}, fmt.Errorf("resolve transcript privacy hostname: %w", err)
	}

	sanitizer := transcriptSanitizer{}
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

	slices.SortStableFunc(sanitizer.replacements, func(left, right transcriptReplacement) int {
		return len(right.value) - len(left.value)
	})
	return sanitizer, nil
}

func (sanitizer *transcriptSanitizer) addPath(value, replacement string) {
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

func (sanitizer *transcriptSanitizer) addIdentity(value, replacement string) {
	value = strings.TrimSpace(value)
	if value == "" || genericIdentity(value) {
		return
	}
	sanitizer.addReplacement(value, replacement, true)
}

func (sanitizer *transcriptSanitizer) addReplacement(value, replacement string, identity bool) {
	if value == "" {
		return
	}
	if slices.ContainsFunc(sanitizer.replacements, func(existing transcriptReplacement) bool { return existing.value == value }) {
		return
	}
	sanitizer.replacements = append(sanitizer.replacements, transcriptReplacement{value: value, replacement: replacement, identity: identity})
	escaped, err := json.Marshal(value)
	if err == nil {
		encoded := string(escaped[1 : len(escaped)-1])
		if encoded != value {
			sanitizer.replacements = append(sanitizer.replacements, transcriptReplacement{value: encoded, replacement: replacement, identity: identity})
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

func (sanitizer transcriptSanitizer) sanitize(result *replay.Result) error {
	for index := range result.Scenario.Turns {
		result.Scenario.Turns[index].Prompt = sanitizer.sanitizeText(result.Scenario.Turns[index].Prompt)
	}
	for turnIndex := range result.Turns {
		capture := &result.Turns[turnIndex].Capture
		capture.Transcript = sanitizer.sanitizeText(capture.Transcript)
		capture.Stderr = sanitizer.sanitizeText(capture.Stderr)
		for eventIndex, event := range capture.Events {
			sanitized, err := sanitizer.sanitizeEvent(event)
			if err != nil {
				return fmt.Errorf("sanitize transcript event for turn %q: %w", result.Turns[turnIndex].TurnID, err)
			}
			capture.Events[eventIndex] = sanitized
		}
	}
	return nil
}

func (sanitizer transcriptSanitizer) sanitizeEvent(event json.RawMessage) (json.RawMessage, error) {
	decoder := json.NewDecoder(bytes.NewReader(event))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return nil, err
	}
	if err := decoder.Decode(&struct{}{}); err != io.EOF {
		return nil, errors.New("structured event contains trailing data")
	}
	value = sanitizer.sanitizeJSON(value)
	data, err := json.Marshal(value)
	if err != nil {
		return nil, err
	}
	return json.RawMessage(data), nil
}

func (sanitizer transcriptSanitizer) sanitizeJSON(value any) any {
	switch typed := value.(type) {
	case string:
		return sanitizer.sanitizeText(typed)
	case []any:
		for index := range typed {
			typed[index] = sanitizer.sanitizeJSON(typed[index])
		}
	case map[string]any:
		sanitized := make(map[string]any, len(typed))
		for key, child := range typed {
			sanitized[sanitizer.sanitizeText(key)] = sanitizer.sanitizeJSON(child)
		}
		return sanitized
	}
	return value
}

func (sanitizer transcriptSanitizer) sanitizeText(value string) string {
	for _, replacement := range sanitizer.replacements {
		value = replaceTranscriptValue(value, replacement)
	}
	return value
}

func replaceTranscriptValue(value string, replacement transcriptReplacement) string {
	start := 0
	for {
		index := strings.Index(value[start:], replacement.value)
		if index < 0 {
			return value
		}
		index += start
		end := index + len(replacement.value)
		if transcriptBoundary(value, index, end, replacement.identity) {
			value = value[:index] + replacement.replacement + value[end:]
			start = index + len(replacement.replacement)
			continue
		}
		start = end
	}
}

func transcriptBoundary(value string, start, end int, identity bool) bool {
	if start > 0 {
		previous, _ := utf8.DecodeLastRuneInString(value[:start])
		if transcriptWordRune(previous, identity) {
			return false
		}
	}
	if end < len(value) {
		next, _ := utf8.DecodeRuneInString(value[end:])
		if transcriptWordRune(next, identity) {
			return false
		}
	}
	return true
}

func transcriptWordRune(value rune, identity bool) bool {
	if unicode.IsLetter(value) || unicode.IsDigit(value) || value == '_' || value == '-' {
		return true
	}
	return !identity && value == '.'
}
