# CLI Portability Hardening — Zero-Context Handoff

You are picking up an active effort with no prior conversation. This document is self-contained: read it, then the authoritative sources it names, and you can continue the work. The mission: make the Skill Issue CLI as portable as possible — a stranger's machine with normally installed harnesses must go from download to a passing evaluation using only the documentation — with no shortcuts, meaning portability gets proven on clean machines and defended with automation, not asserted.

## What This System Is

The Skill Issue CLI (`cli/`, Go, module root at the repository root) installs canonical skills into six coding-agent harnesses (`claude-code`, `codex`, `cursor`, `opencode`, `kilo-code`, `pi`) and runs blind, turn-attributed skill-calling evaluations against them: it installs instrumented temporary skills into an isolated per-run runtime, replays scenario turns through the harness's native CLI, records skill invocations via an opaque-token `signal` subcommand, and derives expected/observed/missing results. Read `cli/README.md` in full — it is the product contract.

## Authoritative Sources (read before acting)

- `AGENTS.md` — CLI channel discipline, evaluation routing, output placement, planning-document policy.
- `.repository-privacy.md` — never record machine-specific absolute paths in tracked files; build distributables with `-trimpath`; commit hooks must be installed (`scripts/setup-repository-hooks.sh`).
- `cli/README.md` — commands, per-harness defaults, qualified versions, artifacts.
- `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md` — the campaign's executable routes, environment requirements, model-identifier resolution, failure classification, and fail-fast rules. (It replaced the retired hardening and evaluation orchestration prompts; the hardening record survives in `hardening-report.md`.)

## State at Handoff

Recently completed and verified on this machine (some committed, some still in the working tree — check `git status` and `git log` first):

- A full code audit with four fix batches: safety (preexisting-skill backup/restore + confirmation, PID-stamped run locks, run-ID validation), correctness (ESC-key handling, Pi process-group ownership and RPC timeouts, CRLF frontmatter, strict shell-word signal matching), structure (harness spec registry as single source of truth, per-harness `buildArgs` dispatch, OpenCode/Kilo runtime dedup, typed run statuses), polish (`slices`/`errors.Is` idioms, `--key=value` option parsing, macOS/Linux fail-fast, `SKILL_ISSUE_ALLOW_UNQUALIFIED_HARNESS=1` version-pin escape hatch).
- A six-harness one-turn smoke: all harnesses verified tooling-clean. Fixes that came out of it: relative custom-input paths, Codex `--skip-git-repo-check` on both `exec` forms (non-git workspaces now work end-to-end, verified live).
- Failure diagnostics: failed runs write `failure.json` into the run's output directory with the exact harness command line, active turn, full error chain, and complete native stdout/stderr (`replay.DiagnosticError` → `evaluation.toolingFailure`). Verified live with a fake failing harness.
- Turn progress output: spinner on interactive terminals, per-turn `Starting`/`Finished` lines with duration, harness events, and skill-call counts.
- Measured performance: CLI overhead is under one second per run; all remaining latency is model inference. Do not add timeouts or retries around normal turns.
- Phase 3 hardening campaign (four-turn matrix, all six harnesses) completed tooling-clean; report at `plans/skill-calling-evaluation-campaign/hardening-report.md`. It fixed the codex config defect noted below. That fix is uncommitted.

### Known risks and open conflicts at handoff

- **Another agent session has been working the same tree in parallel** (Codex runtime home isolation, payload manifest sources, website data). Reconcile with `git status`/`git log` before assuming anything below is still true.
- **Codex config (resolved)**: an earlier working tree carried `agents.enabled=false`, which codex-cli rejects fatally (`invalid type: boolean, expected struct AgentRoleToml`) because `agents` now holds agent-role config. The Phase 3 hardening campaign fixed this in `cli/internal/evaluation/runtime.go` to `features.multi_agent=false` (plus the retained `features.multi_agent_v2=false`), verified against codex `0.144.6`; the runtime test and `cli/README.md` were updated to match. This fix is **not yet committed or promoted to known-good**. The cheap verification technique still applies to any future codex config change: `codex -c '<key>=<value>' login status` parses configuration in under a second with no model cost.
- `npm run format:check` fails on `src/data/evaluationData.ts` (website work, not CLI); `cli/dist/` cross-platform binaries are stale and predate all of the above; failed runs may leave an empty timestamped output directory plus `failure.json` (expected); Codex workspaces keep a cosmetic empty `.agents/` directory after cleanup.

## The Roadmap (execute in order)

### Phase 0 — Freeze the baseline (~half day)

Commit the working tree (coordinate if the parallel session is still active; privacy hooks must pass), rebuild both CLI channels (`./cli/scripts/local-cli.sh build-development` and `build-known-good`), run `gofmt`/`go vet`/`go test ./cli/...`, and run one six-harness one-turn smoke against that exact commit per the hardening orchestration document. Every later failure must be attributable to environment versus change; that needs a fixed reference point.

### Phase 1 — Self-diagnosing CLI (~1–2 days)

Build `skill-issue doctor [--harness <id>]`: resolve the executable, check the version against pinned/tested, verify authentication, validate the generated harness configuration against the real binary (the cheap parse check above), confirm skill discovery where the harness supports it, and check platform/workspace/output preconditions — seconds, zero model cost, actionable output per finding. In the same pass: add skill-visibility verification to real runs for every harness (`CheckOpenCodeSkills` exists for OpenCode; add the Kilo equivalent; Claude Code exposes its visible skills in the first `system/init` event; Pi already validates via `get_commands`) so a run can never again "succeed" with the governed skills silently unloaded — this exact failure happened and was only caught by an operator noticing an empty `events.jsonl`. Add a `--yes` flag for scripted confirmation (current workaround is piping `echo y`).

Also close a robustness gap the hardening campaign surfaced: a **relative `--executable` (or `--cli-path`) fails late and cryptically** — the CLI switches to a run-owned working directory before spawning the harness, so a caller-relative path yields `fork/exec <path>: no such file or directory` at turn 1 instead of a clear early error. Resolve `--executable` to an absolute path against the invocation cwd at parse time (as workspace/output already are), or reject a non-resolvable one in preflight/`doctor`. This is exactly the papercut a stranger passing `./bin/claude` would hit.

### Phase 2 — Fake-harness conformance suite (~2–3 days)

Scripted fake binaries speaking each harness's protocol (the pattern already exists in `cli/internal/replay/process_test.go` and was used live to verify `failure.json`): happy path, protocol violations, mid-turn death, session-ID change, config rejection. Drive full `evaluate run` invocations against them inside `go test`. This makes the entire adapter layer testable on any machine with no vendor credentials. Then CI: GitHub Actions on macOS **and** Linux running vet, tests, and this suite on every push.

### Phase 3 — Real-harness hardening campaign (DONE — verify, don't redo)

Completed on this machine; see `plans/skill-calling-evaluation-campaign/hardening-report.md` and the `hardening-progress.md` ledger. All six harnesses passed a four-turn `gardening-web-application` run tooling-clean over ten attempts. One production defect (the codex config key above) and three invocation/environment corrections were found; the four fixed environment requirements (nested-codex outer escalation, writable Claude session state, absolute project-local routes, Cursor's `cursor-grok-4.5-medium` identifier) are lab-facing, not product defects. **Do not re-run this as new work.** Re-run the matrix only as a regression gate after a shared-code change (replay/evaluation/installer) or after committing the codex fix, since the passing runs were on an uncommitted development build. Evidence lives under `output/harness-hardening/` (ignored, unsanitized — keep it out of commits).

### Phase 4 — Clean-machine validation (the non-negotiable core, ~2–3 days)

Two environments that have never seen this repository: a clean macOS machine or VM, and a Linux machine or container. Install each harness the normal public way (brew/npm/official installers), authenticate normally, run `doctor`, then a one-turn smoke per harness. Expect Linux to surface real defects — it has never been exercised, and the controlled environment carries macOS assumptions (the Cursor keychain symlink under `Library/Keychains`, `SHELL=/bin/zsh`, PATH construction in `cli/internal/evaluation/runtime.go` `controlledEnvironment`). Fix every finding in product code; never with machine-specific workarounds. Exit criterion: a stranger's laptop goes from clone/download to a passing evaluation using only the README.

### Phase 5 — Version policy for unpinned harnesses (~half day)

OpenCode (`1.18.4`) and Kilo (`7.4.11`) are pinned exactly with the escape-hatch env var. Extend a recorded-and-warn model to claude-code, codex, cursor, and pi: record the tested version, warn on drift, never hard-block. Turns future vendor breakage into a labeled event.

### Phase 6 — Distribution and stranger-facing docs (~1 day)

Rebuild `cli/dist/` via `cli/scripts/build-cross-platform.sh` (`-trimpath`; verify binaries contain no personal paths per the privacy policy). Write a per-harness QUICKSTART: install normally, authenticate, `doctor`, `evaluate run`; troubleshooting section that leads with `failure.json`. Keep Windows as-is: install works, evaluation fails fast with a clear message.

### Explicitly out of scope

Do not move this machine's qualification-lab specifics (the `.skill-issue/` harness environments, `XDG_DATA_HOME` exports, wrapper knowledge) into the product — a normal install needs none of it. Do not weaken harness sandboxes or version pins to make errors disappear. Do not pursue Windows evaluation support.

## Operational Knowledge You Will Need

- **Channels**: `./cli/scripts/local-cli.sh` — `build-development` builds from the working tree; `development <args>` runs it; known-good builds from committed `HEAD` via git archive and is the default channel. Rebuild development after every source change; promote known-good only from a committed baseline.
- **On this machine only**, harness routes are non-standard (qualified lab installs). All routes, required `XDG_DATA_HOME` exports, wrapper traps, and model defaults are in the hardening orchestration document — follow it verbatim here. On clean machines (Phase 4), use none of it: normal installs, PATH resolution, real data homes.
- **Reading results**: `result.json` (expected/observed/missing/additional/unattributed), `events.jsonl` (`--events`), `transcript.json` (`--transcript`, sanitized), `failure.json` (unsanitized post-mortem — never commit or share). **Missing skill calls are model behavior — and a hardening PASS — only once corroborated: the run recorded at least one attributed signal, and each missing skill is proven visible to the harness (its own visibility evidence, or an observed invocation of that skill on the same harness and build). An uncorroborated miss is unclassified until a `--transcript` visibility check settles it; "evaluation encountered a tooling error" is always a defect.**
- **Recovery**: if an error message names an `evaluate cleanup --run <id> --output <root>` command, run exactly that, then verify workspace skills are cleaned/restored and no private state remains under the system temp directory's `skill-issue/` path.
- **Validation gates for every change**: `gofmt -w cli/`, `go build ./cli/...`, `go vet ./cli/...`, `go test ./cli/...`; prettier for Markdown you touch; keep test fixtures identity-neutral (no machine-style personal paths — see existing tests for the convention).
- **Outputs**: always under repository-root `output/` in descriptive subdirectories, workspaces always outside the repository, never force-add `output/` or `.skill-issue/`.

## Definition of Done

Portability is done when: `doctor` exists and passes on this machine for all six harnesses; the fake-harness conformance suite runs in CI on macOS and Linux; the hardening campaign matrix is tooling-clean; both clean machines (macOS and Linux, normal installs) complete `doctor` plus a one-turn smoke on every harness with zero lab-specific setup; version drift warns instead of mystifying; and fresh `-trimpath` distributables plus a QUICKSTART exist. At that point, recommend removal or archival of this handoff and the hardening orchestration document per the repository's working-document policy.
