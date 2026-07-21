// Package doctor diagnoses whether the current machine can run Skill Issue
// evaluations. Every check completes in seconds with zero model cost: it
// resolves each harness executable, compares the installed version against
// the tested or pinned version, verifies authentication through each
// harness's native status surface, and validates the generated Codex
// configuration against the installed binary's own parser.
package doctor

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	goruntime "runtime"
	"strings"
	"time"

	"github.com/ericwimp8/skill-issue/cli/internal/evaluation"
	"github.com/ericwimp8/skill-issue/cli/internal/harness"
)

type Status string

const (
	StatusOK   Status = "ok"
	StatusWarn Status = "warn"
	StatusFail Status = "fail"
)

type Check struct {
	Harness string `json:"harness,omitempty"`
	Name    string `json:"check"`
	Status  Status `json:"status"`
	Detail  string `json:"detail"`
}

type Report struct {
	SchemaVersion int     `json:"schema_version"`
	Healthy       bool    `json:"healthy"`
	Checks        []Check `json:"checks"`
}

const probeTimeout = 15 * time.Second

// allowUnqualifiedHarnessEnv mirrors the evaluator's version-pin escape
// hatch so doctor reports the same decision a run would make.
const allowUnqualifiedHarnessEnv = "SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS"

// Run diagnoses the selected harness, or every evaluation harness when the
// filter is empty, and streams one human-readable line per finding.
func Run(ctx context.Context, harnessFilter harness.ID, executableOverride string, progress io.Writer) (Report, error) {
	if progress == nil {
		progress = io.Discard
	}
	targets := harness.SupportedIDs()
	if harnessFilter != "" {
		if _, err := harness.ParseEvaluationID(string(harnessFilter)); err != nil {
			return Report{}, err
		}
		targets = []harness.ID{harnessFilter}
	} else if executableOverride != "" {
		return Report{}, errors.New("--executable requires --harness to name the harness it overrides")
	}
	report := Report{SchemaVersion: 1, Healthy: true}
	record := func(check Check) {
		report.Checks = append(report.Checks, check)
		if check.Status == StatusFail {
			report.Healthy = false
		}
		subject := check.Name
		if check.Harness != "" {
			subject = check.Harness + " " + check.Name
		}
		fmt.Fprintf(progress, "%4s  %s: %s\n", check.Status, subject, check.Detail)
	}
	for _, check := range platformChecks() {
		record(check)
	}
	for _, id := range targets {
		for _, check := range harnessChecks(ctx, id, executableOverride) {
			record(check)
		}
	}
	return report, nil
}

func platformChecks() []Check {
	checks := []Check{}
	if goruntime.GOOS == "darwin" || goruntime.GOOS == "linux" {
		checks = append(checks, Check{Name: "platform", Status: StatusOK, Detail: goruntime.GOOS + " supports evaluation"})
	} else {
		checks = append(checks, Check{Name: "platform", Status: StatusFail, Detail: fmt.Sprintf("evaluation requires macOS or Linux; %s is not supported", goruntime.GOOS)})
	}
	temporary := os.TempDir()
	if probe, err := os.CreateTemp(temporary, "skill-issue-doctor-"); err != nil {
		checks = append(checks, Check{Name: "temporary-directory", Status: StatusFail, Detail: fmt.Sprintf("cannot write private runtime state under %s: %v", temporary, err)})
	} else {
		probe.Close()
		os.Remove(probe.Name())
		checks = append(checks, Check{Name: "temporary-directory", Status: StatusOK, Detail: temporary + " is writable for private runtime state"})
	}
	if home, err := os.UserHomeDir(); err != nil {
		checks = append(checks, Check{Name: "home-directory", Status: StatusFail, Detail: fmt.Sprintf("cannot resolve the home directory: %v", err)})
	} else {
		checks = append(checks, Check{Name: "home-directory", Status: StatusOK, Detail: home})
	}
	return checks
}

func harnessChecks(ctx context.Context, id harness.ID, executableOverride string) []Check {
	name := string(id)
	path, err := resolveHarnessExecutable(id, executableOverride)
	if err != nil {
		return []Check{{Harness: name, Name: "executable", Status: StatusFail, Detail: fmt.Sprintf("%v; install %s normally or pass --executable", err, name)}}
	}
	checks := []Check{{Harness: name, Name: "executable", Status: StatusOK, Detail: path}}
	checks = append(checks, versionCheck(ctx, id, path))
	checks = append(checks, authenticationChecks(ctx, id, path)...)
	if id == harness.Codex {
		checks = append(checks, codexConfigurationCheck(ctx, path))
	}
	return checks
}

// resolveHarnessExecutable mirrors the evaluator's resolution: an override
// with a path separator is absolutized against the invocation directory,
// bare names resolve through PATH, and Cursor tries its two native names.
func resolveHarnessExecutable(id harness.ID, override string) (string, error) {
	names := []string{override}
	if override == "" {
		switch id {
		case harness.Cursor:
			names = []string{"agent", "cursor-agent"}
		default:
			spec, err := harness.Lookup(id)
			if err != nil {
				return "", err
			}
			names = []string{spec.Executable}
		}
	} else if strings.ContainsRune(override, os.PathSeparator) {
		absolute, err := filepath.Abs(override)
		if err != nil {
			return "", fmt.Errorf("resolve --executable: %w", err)
		}
		names = []string{absolute}
	}
	var firstErr error
	for _, candidate := range names {
		path, err := exec.LookPath(candidate)
		if err == nil {
			return path, nil
		}
		if firstErr == nil {
			firstErr = err
		}
	}
	return "", firstErr
}

func versionCheck(ctx context.Context, id harness.ID, path string) Check {
	name := string(id)
	output, err := probe(ctx, path, "--version")
	if err != nil {
		return Check{Harness: name, Name: "version", Status: StatusFail, Detail: fmt.Sprintf("--version failed: %v", err)}
	}
	installed := extractVersion(output)
	if installed == "" {
		return Check{Harness: name, Name: "version", Status: StatusWarn, Detail: fmt.Sprintf("cannot parse version from %q", strings.TrimSpace(output))}
	}
	tested, pinned, err := harness.TestedVersion(id)
	if err != nil || tested == "" {
		return Check{Harness: name, Name: "version", Status: StatusWarn, Detail: installed + " (no tested version recorded)"}
	}
	if installed == tested {
		return Check{Harness: name, Name: "version", Status: StatusOK, Detail: fmt.Sprintf("%s matches the tested version", installed)}
	}
	if pinned {
		if os.Getenv(allowUnqualifiedHarnessEnv) == "1" {
			return Check{Harness: name, Name: "version", Status: StatusWarn, Detail: fmt.Sprintf("%s is not the qualified version %s; runs continue because %s=1", installed, tested, allowUnqualifiedHarnessEnv)}
		}
		return Check{Harness: name, Name: "version", Status: StatusFail, Detail: fmt.Sprintf("%s is not the qualified version %s; install the qualified version or set %s=1", installed, tested, allowUnqualifiedHarnessEnv)}
	}
	return Check{Harness: name, Name: "version", Status: StatusWarn, Detail: fmt.Sprintf("%s differs from the tested version %s; behavior may have drifted", installed, tested)}
}

func authenticationChecks(ctx context.Context, id harness.ID, path string) []Check {
	name := string(id)
	switch id {
	case harness.ClaudeCode:
		if output, err := probe(ctx, path, "auth", "status"); err != nil {
			return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("claude auth status failed: %v: %s", err, firstLine(output))}}
		}
		return []Check{{Harness: name, Name: "authentication", Status: StatusOK, Detail: "claude auth status succeeded"}}
	case harness.Codex:
		if output, err := probe(ctx, path, "login", "status"); err != nil {
			return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("codex login status failed: %v: %s; run codex login", err, firstLine(output))}}
		}
		return []Check{{Harness: name, Name: "authentication", Status: StatusOK, Detail: "codex login status succeeded"}}
	case harness.Cursor:
		if output, err := probe(ctx, path, "status"); err != nil {
			return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("cursor status failed: %v: %s; run the agent login flow", err, firstLine(output))}}
		}
		return []Check{{Harness: name, Name: "authentication", Status: StatusOK, Detail: "cursor agent status succeeded"}}
	case harness.OpenCode:
		return structuredAuthenticationChecks(ctx, id, path)
	case harness.Pi:
		directory := os.Getenv("PI_CODING_AGENT_DIR")
		if directory == "" {
			home, err := os.UserHomeDir()
			if err != nil {
				return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("cannot resolve the Pi agent directory: %v", err)}}
			}
			directory = filepath.Join(home, ".pi", "agent")
		}
		if info, err := os.Stat(directory); err != nil || !info.IsDir() {
			return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("Pi agent directory %s is missing; run pi once to initialize it or set PI_CODING_AGENT_DIR", directory)}}
		}
		return []Check{{Harness: name, Name: "authentication", Status: StatusOK, Detail: "Pi agent directory present: " + directory}}
	}
	return nil
}

func structuredAuthenticationChecks(ctx context.Context, id harness.ID, path string) []Check {
	name := string(id)
	defaults, err := harness.EvaluationDefaultsFor(id)
	if err != nil {
		return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: err.Error()}}
	}
	provider, _, found := strings.Cut(defaults.Model, "/")
	if !found {
		return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("default model %q is not provider/model", defaults.Model)}}
	}
	authArgs := []string{"auth", "list"}
	modelArgs := []string{"models", provider}
	if id == harness.OpenCode {
		authArgs = append(authArgs, "--pure")
		modelArgs = append(modelArgs, "--pure")
	}
	output, err := probe(ctx, path, authArgs...)
	if err != nil {
		return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("auth list failed: %v: %s", err, firstLine(output))}}
	}
	if !strings.Contains(strings.ToLower(output), strings.ToLower(provider)) {
		return []Check{{Harness: name, Name: "authentication", Status: StatusFail, Detail: fmt.Sprintf("provider %q is not authenticated; run %s auth login", provider, name)}}
	}
	checks := []Check{{Harness: name, Name: "authentication", Status: StatusOK, Detail: fmt.Sprintf("provider %q is authenticated", provider)}}
	models, err := probe(ctx, path, modelArgs...)
	if err != nil {
		checks = append(checks, Check{Harness: name, Name: "model", Status: StatusFail, Detail: fmt.Sprintf("models %s failed: %v: %s", provider, err, firstLine(models))})
		return checks
	}
	if !containsLine(models, defaults.Model) {
		checks = append(checks, Check{Harness: name, Name: "model", Status: StatusFail, Detail: fmt.Sprintf("default model %q is unavailable for provider %q", defaults.Model, provider)})
		return checks
	}
	checks = append(checks, Check{Harness: name, Name: "model", Status: StatusOK, Detail: fmt.Sprintf("default model %q is available", defaults.Model)})
	return checks
}

// codexConfigurationCheck parses the exact generated evaluation configuration
// through the installed binary. codex -c <key>=<value> login status parses
// configuration in under a second with no model cost; a rejected key is the
// class of defect that once broke every Codex evaluation.
func codexConfigurationCheck(ctx context.Context, path string) Check {
	defaults, err := harness.EvaluationDefaultsFor(harness.Codex)
	if err != nil {
		return Check{Harness: "codex", Name: "configuration", Status: StatusFail, Detail: err.Error()}
	}
	args := []string{}
	for _, value := range evaluation.CodexBaseConfiguration(defaults.Reasoning) {
		args = append(args, "-c", value)
	}
	args = append(args, "login", "status")
	if output, err := probe(ctx, path, args...); err != nil {
		return Check{Harness: "codex", Name: "configuration", Status: StatusFail, Detail: fmt.Sprintf("the installed codex rejected the generated evaluation configuration: %v: %s", err, firstLine(output))}
	}
	return Check{Harness: "codex", Name: "configuration", Status: StatusOK, Detail: "the installed codex parses the generated evaluation configuration"}
}

func probe(ctx context.Context, path string, args ...string) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, probeTimeout)
	defer cancel()
	command := exec.CommandContext(ctx, path, args...)
	var output bytes.Buffer
	command.Stdout = &output
	command.Stderr = &output
	err := command.Run()
	return output.String(), err
}

// extractVersion returns the first whitespace-separated token that starts
// with a digit, tolerating banners like "codex-cli 0.144.6" and suffixes
// like "2.1.205 (Claude Code)".
func extractVersion(output string) string {
	for _, field := range strings.Fields(output) {
		if field != "" && field[0] >= '0' && field[0] <= '9' {
			return field
		}
	}
	return ""
}

func firstLine(output string) string {
	line, _, _ := strings.Cut(strings.TrimSpace(output), "\n")
	return line
}

func containsLine(output, wanted string) bool {
	for _, line := range strings.Split(output, "\n") {
		if strings.TrimSpace(line) == wanted {
			return true
		}
	}
	return false
}
