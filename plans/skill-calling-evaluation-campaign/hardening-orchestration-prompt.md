# Harness Evaluation Path Hardening Orchestration Prompt

## Goal

Harden the harness evaluation path by running small four-turn built-in evaluations on every supported harness, diagnosing every tooling failure to its owning cause, fixing what is broken while maintaining portability, and rerunning until each harness completes cleanly. These runs exist to exercise the tooling, not to collect campaign results: a model that skips an expected skill call is data; a run that reports `evaluation encountered a tooling error` is a defect to fix.

## Authoritative Sources

Read these before acting:

- `AGENTS.md` — CLI channel discipline, evaluation routing, output placement, repository privacy.
- `cli/README.md` — command interface, per-harness defaults, qualified versions, diagnostics artifacts.
- `plans/skill-calling-evaluation-campaign/evaluation-orchestration-prompt.md` — the campaign conventions this document inherits (isolation, process ownership, the Codex outer-sandbox escalation rule when the orchestrating agent is itself Codex).

## Run Matrix

Six runs, one per harness, four turns each, built-in scenario, single model per harness. Reasoning is `medium` everywhere, which is the default for every harness — do not pass `--reasoning` for any run (Cursor rejects the override outright). Models are the defaults everywhere except Cursor, so pass `--model` only there:

| Harness       | Model                                | `--model` flag            |
| ------------- | ------------------------------------ | ------------------------- |
| `codex`       | `gpt-5.6-sol` (default)              | omit                      |
| `claude-code` | `opus` (default)                     | omit                      |
| `cursor`      | Grok                                 | `--model grok` (required) |
| `opencode`    | `openai/gpt-5.6-sol` (default)       | omit                      |
| `kilo-code`   | `openai/gpt-5.6-sol` (default)       | omit                      |
| `pi`          | `openai-codex/gpt-5.6-sol` (default) | omit                      |

For Cursor, confirm the exact Grok model identifier the installed agent accepts before the first run (list the agent's available models if `grok` is rejected) and record the identifier you used. Always verify the pre-run summary shows the intended harness, executable, model, reasoning, workspace, and output root before confirming with `y`.

## CLI Channel

Hardening tests current source. Build and use the development channel:

```sh
./cli/scripts/local-cli.sh build-development
./cli/scripts/local-cli.sh development <arguments>
```

Rebuild `build-development` after every source fix and before every rerun. Do not rebuild known-good from uncommitted work; promote to known-good only after the hardening fixes are committed as the next baseline.

## Isolation Structure

Follow the campaign's isolation rules exactly:

- One unique external temporary git workspace per run, outside this repository. Create it fresh (`mktemp -d`), `git init` it, and commit an initial file. Never reuse a workspace across runs — leftover files from a prior run (for example a previously written `PLAN.md`) change model behavior and contaminate the comparison.
- One distinct output root per run under the repository's ignored output tree: `output/harness-hardening/<harness>-<attempt>/`. Never point two runs at the same workspace or output location.
- Never run two evaluations in the same workspace or output root; reuse a slot only after the prior command has ended and its cleanup is verified.
- Run harnesses sequentially, or at most two concurrently. Model inference dominates wall time (roughly 30 seconds to 5 minutes per turn); concurrency inflates every run's latency and muddies failure diagnosis.

## Command Template

```sh
echo y | ./cli/scripts/local-cli.sh development evaluate run \
  --workspace <fresh-external-git-workspace> \
  --output output/harness-hardening/<harness>-<attempt> \
  --harness <id> \
  --executable <explicit-route> \
  --turns 4 --events \
  --evaluation gardening-web-application
```

Add `--model grok` for Cursor only. Add `--transcript` on diagnostic reruns when you need the full sanitized harness event stream. The `echo y` answers the pre-run confirmation in non-interactive shells; in an interactive terminal, review the summary and confirm manually.

## Explicit Executable Routes

Shell aliases are invisible to the CLI (`exec.LookPath`); pass `--executable` explicitly for every run:

- `claude-code`: the operator's normal Claude Code executable — `"$(whence -p claude)"`. Do not use the `claude-codex` alias or the Codex-backed proxy launcher for these hardening runs.
- `codex`: `"$(whence -p codex)"`.
- `cursor`: the project-local isolated Cursor agent binary at `.skill-issue/cursor/home/.local/bin/agent`. Do not use the `cursorx` wrapper — it forces its own environment and workspace and would fight the CLI's run-owned runtime.
- `opencode`: `.skill-issue/opencode/bin/opencode` (qualified `1.18.4`). This wrapper only defaults unset variables, so the CLI's run-owned configuration still applies.
- `kilo-code`: the real binary at `.skill-issue/kilo/node_modules/@kilocode/cli-darwin-arm64/bin/kilo` (qualified `7.4.11`). Never use `.skill-issue/kilo/bin/kilo` — that wrapper hard-pins `XDG_CONFIG_HOME`, silently discarding the CLI's generated evaluation configuration, and the run then "succeeds" with zero skills loaded.
- `pi`: the installed Pi runtime entrypoint (typically `~/.volta/bin/pi`).

### Required environment per invocation

- `opencode` and `kilo-code`: export `XDG_DATA_HOME` to the qualified environment's data home for that single invocation — `.skill-issue/opencode/home/.local/share` or `.skill-issue/kilo/home/.local/share` respectively. The CLI passes the caller's data home through to the harness; without it the auth check fails with `provider "openai" is not authenticated` even though the qualified environment is logged in.
- `pi`: authentication resolves through `PI_CODING_AGENT_DIR` (default `~/.pi/agent`), which the CLI preserves. Providers keyed purely by environment variable must use a variable on the CLI's forwarding allowlist (`cli/internal/evaluation/runtime.go`, `forwardedCredentialKeys`).
- Version pins: OpenCode and Kilo must match their qualified versions exactly or the run stops before side effects. `SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS=1` downgrades the mismatch to a warning; use it only deliberately and record that you did.

## Understanding Outcomes

Read artifacts from the run's output directory `output/harness-hardening/<harness>-<attempt>/<harness>-<timestamp>-<prefix>/`:

- `result.json` — `expected`, `observed`, `missing`, `additional`, `unattributed` skill calls plus run metadata and timestamps. **Missing or additional calls are model behavior, not tooling failures.** A tooling-complete run with missing calls is a hardening pass.
- `events.jsonl` (with `--events`) — each recorded signal with turn attribution and time. An empty file plus a completed run means the model made no instrumented skill calls.
- `transcript.json` (with `--transcript`) — the full sanitized harness event stream per turn, including the harness's own stderr. Use it to see what the model actually saw and did (for Claude Code, the first `system/init` event lists the visible skills and tools).
- `failure.json` — written automatically when a run fails with a tooling error. It records the run ID, harness, model, reasoning, **active turn**, the full error chain, the **exact harness command line the CLI executed**, and the complete native stdout and stderr of the failed interaction. This is your primary post-mortem artifact; the terminal error message names its path. It is not sanitized — it contains local paths and the turn prompt — so keep it out of anything shared or committed.
- Progress stream — `Starting turn i of 4` / `Finished turn i of 4: <id> (<duration>, <n> harness events, <n> skill calls)` per turn. Zero harness events on a "finished" turn is suspicious; zero skill calls is not.

### Failure classification

Work every failure to one of these causes before changing anything:

1. **Route/environment** — wrong executable, missing `XDG_DATA_HOME`, version pin mismatch, missing authentication. Fix the invocation, not the CLI.
2. **Harness configuration rejected** — the harness exits refusing its own config (`failure.json` stdout/stderr shows a config parse error naming a key). The generated configuration lives in `cli/internal/evaluation/runtime.go`; validate a candidate key cheaply against the real binary before changing it (for example `codex -c '<key>=<value>' login status` parses configuration without running a turn).
3. **Protocol violation** — `malformed harness protocol` errors: missing init/completion events, session ID changes, unparseable output. Compare `failure.json` stdout against the validator expectations in `cli/internal/replay/process.go` (`validateHarnessOutput`); a harness version change is the usual suspect.
4. **Model behavior** — completed run, unexpected `missing`/`additional` sets. Record it; do not "fix" it.

### Recovery

- If the error message ends with a `skill-issue evaluate cleanup --run <id> --output <root>` instruction, cleanup failed mid-run: execute that exact command, then verify the workspace's temporary skills are gone and any preexisting skills were restored before rerunning.
- Otherwise failed runs self-clean; the surviving `failure.json` (and possibly an empty run directory) is expected residue.
- Verify after every run, pass or fail: no run-owned harness process is still alive, no temporary skill directories remain in the workspace, and no private runtime remains under the system temporary directory's `skill-issue/<run-id>` path. A leftover empty `.agents/` directory in Codex workspaces is known cosmetic residue, not a defect.

## Known Tips and Traps

- **Non-git workspaces**: Codex evaluations pass `--skip-git-repo-check`, so a bare directory works — but keep the git-init convention anyway for uniformity and easy diffing of workspace effects.
- **Claude Code shows more than seven skills**: its built-in skills (verify, code-review, and others) appear alongside the seven governed skills on any default install. Expected; not contamination.
- **Claude Code may complete tasks without invoking any skill** on subtle prompts. That is a valid result. To prove the signal path itself, the campaign's two-turn explicit custom smoke (`evaluations/skill-calling/smoke/`) forces an invocation.
- **Preexisting skill collision**: if a workspace already contains a skill directory with a canonical name, the CLI asks for confirmation (interactive) or requires `--replace-preexisting-skills` (scripted), backs the directory up, and restores it byte-for-byte at cleanup. Fresh workspaces avoid the prompt entirely.
- **Custom input paths may be relative** (fixed); workspace and output are absolutized by the CLI. The output root must remain outside the evaluated workspace and the workspace outside this repository.
- **Slow runs are model time.** The CLI's own overhead is under a second per run; per-turn cold starts and session resumption add a few seconds per turn for the Node/Bun-based harnesses. Do not add timeouts or retries around normal turns.
- **Interrupting a run** (Ctrl-C / SIGTERM) is handled: the process group is killed and cleanup runs. Still verify workspace and private-state cleanup afterward.
- **Do not use earlier passing runs as proof that a current failure is invalid.** Diagnose each failure from its own `failure.json`, native output, and the production source path that owns the behavior.

## Fix Discipline

- Fix route/environment problems in the invocation and record the requirement here or in the progress notes.
- Fix CLI defects in production source, keeping portability: no machine-specific paths, no platform assumptions beyond the documented macOS/Linux support, no weakening of harness sandboxes or version pins without an explicit recorded decision.
- After any source fix: `gofmt`, `go vet`, `go test ./cli/...`, rebuild the development channel, and rerun the affected harness from a fresh workspace. A fix that changes shared code (replay, evaluation, installer) invalidates prior hardening passes — rerun the full six-harness matrix after such a fix.
- Keep every diagnostic conclusion tied to retained evidence under `output/harness-hardening/`.

## Reporting

Track progress in a small table (harness, attempt, workspace disposition, outcome, failure classification, fix applied, evidence path) either in this document's companion notes or a sibling progress file in this directory. Record the exact executable path, effective model, and reasoning for every run — the pre-run summary and `result.json` both state them. When hardening ends, summarize defects found and fixed, surviving environment requirements, and recommend this document's own retention or removal per the repository's working-document policy. Never commit anything under `output/`, and never record machine-specific absolute paths in tracked files.

## Completion Criteria

Hardening is complete when every one of the six runs has, from a fresh workspace on the current development build: completed all four turns tooling-clean with one stable session; written `result.json`, `website.json`, and `events.jsonl`; left the workspace free of temporary skills and the machine free of run-owned processes and private state; and every earlier tooling failure has a diagnosed cause, an applied and validated fix, and a passing rerun recorded.
