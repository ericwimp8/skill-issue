# Local Candidate Research: Grok Build, OpenCode, and Kilo Code

## Assignment

**Goal:** Map the repository's existing local evidence relevant to technical qualification of Grok Build, OpenCode, and Kilo Code, including production implementations, contracts, prior research, and safe local executable probes.

**Scope:** Repository production source, controlling plans and CLI documentation, retained direct-installation/manual-invocation/sub-agent research, retained smoke evidence, and read-only local executable discovery. The map covers isolation, temporary skills, noninteractive replay, session/resume behavior, protocol output, permissions, cancellation, cleanup, executable discovery, model/reasoning, authentication, and Codex-subscription claims.

**Exclusions:** Internet research, changing product source or support status, installing or authenticating any candidate harness, running a model session, and deciding final candidate support.

## Sources

- Production source: `cli/internal/harness/harness.go` — the nine harness specifications, direct-install roots, and the four-entry evaluation-defaults allowlist.
- Production source: `cli/internal/evaluation/runtime.go` and `cli/internal/evaluation/evaluation.go` — runtime preparation, authentication-preflight routing, temporary-skill preparation, evaluation cleanup, and the point at which unsupported evaluation harnesses fail.
- Production source: `cli/internal/replay/process.go`, `cli/internal/replay/replay.go`, and `cli/internal/replay/process_group_unix.go` — candidate command stubs, JSON-event parsing, session extraction, process-group ownership, cancellation, and close behavior.
- Production source: `cli/internal/installer/installer.go` — generic instrumented temporary-skill materialization and restoration/removal semantics.
- Local product contract: `cli/README.md` — installation/evaluation boundary, current qualified smoke boundary, unqualified-adapter caveats, and Grok's missing fail-closed discovery check.
- Local product contract: `README.md`, `plans/harness-setup.md`, `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`, and `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md` — minimum qualification tier, default-environment expectation, and complete proof requirements.
- Prior research: `research/deep-research/harness-direct-installation-architecture/assignments/06-grok-direct-installation.md`, `07-opencode-direct-installation.md`, and `08-kilo-code-direct-installation.md`; their synthesis at `research/deep-research/harness-direct-installation-architecture/harness-direct-installation-architecture-deep-research.md`.
- Prior research: `research/deep-research/manual-skill-invocation-nine-harnesses/assignments/06-grok-build.md`, `07-opencode.md`, and `08-kilo-code.md`; their synthesis at `research/deep-research/manual-skill-invocation-nine-harnesses/manual-skill-invocation-nine-harnesses-deep-research.md`.
- Prior research: `research/harness-subagent-launch/assignments/06-grok-build.md`, `07-opencode.md`, and `08-kilo-code.md`; their index at `research/harness-subagent-launch/harness-subagent-launch-reference.md`.
- Validation evidence: read-only local shell probes on 2026-07-21. `grok` and `kilo` were absent from `PATH`. `opencode` was present at version `1.14.39`; `opencode --help`, `opencode run --help`, and `opencode debug --help` exposed `run`, `--session`, `--format json`, `--model`, `--variant`, `--pure`, `--dangerously-skip-permissions`, and `debug skill`.
- Validation evidence boundary: `evaluations/skill-calling/smoke/real-harness-smoke-report.md` records only Codex, Cursor, Claude Code, and Pi smoke runs; it contains no candidate result for Grok Build, OpenCode, or Kilo Code.

## Findings

### 1. Candidate registration is production support for ordinary installation, not current evaluator support

The production harness registry gives all three candidates a native executable and project/user skill root: Grok Build uses `grok`, `.grok/skills`, and `~/.grok/skills`; OpenCode uses `opencode`, `.opencode/skills`, and `~/.config/opencode/skills`; Kilo Code uses `kilo`, `.kilo/skills`, and `~/.kilo/skills`. The installer can stage, replace, instrument, restore, or remove Skill Issue-owned directories for any registered harness when given a root. In contrast, `evaluationDefaults` contains only Codex, Cursor, Claude Code, and Pi; `ParseEvaluationID` rejects all three candidates before an evaluation request can proceed.

**Evidence:** Production source in `cli/internal/harness/harness.go` defines the three `Spec` entries and only four `evaluationDefaults` entries. `cli/internal/installer/installer.go` is generic over `harness.ID`, while `cli/README.md` explicitly says installation supports all nine and evaluation currently supports four.

**Implication:** Classify native roots and ordinary payload copying as **production source**, and candidate evaluation availability as **unresolved/unsupported in current production**. Any later qualification must first add a source-owned evaluation contract; direct installation alone supplies no evidence for temporary-skills, replay, or cleanup on these candidates.

### 2. Runtime preparation stops before temporary materialization for all three candidates

`Service.Run` creates output/private state and then calls `prepareRuntime` before authentication, temporary-skill installation, adapter construction, replay, or cleanup. `prepareRuntime` has dedicated branches only for Codex, Cursor, Claude Code, and Pi; its default returns `unsupported evaluation harness`. Therefore Grok Build, OpenCode, and Kilo Code cannot reach `PrepareEvaluation`, `replay.NewAdapter`, or an agent prompt through the CLI as currently implemented.

**Evidence:** Production source in `cli/internal/evaluation/evaluation.go` orders `prepareRuntime` before `CheckAuthentication`, `installer.PrepareEvaluation`, and `adapterFactory`. `cli/internal/evaluation/runtime.go` switches only on Cursor, Claude Code, and Pi after its Codex branch, then returns an unsupported-harness error. The checked-in smoke report contains no candidate run.

**Implication:** The following prerequisites are absent as **production implementation** for every candidate: a private/clean runtime home, ambient skill/config/plugin exclusion, a candidate-owned evaluation skill root, candidate authentication preflight, and a demonstrated cleanup path. The generic installer is a reusable component, but candidate end-to-end instrumentation remains **unresolved** rather than merely untested.

### 3. Grok Build has an unqualified replay stub and a fail-closed installation requirement that production does not enforce

The generic replay map proposes `grok -p <prompt> --output-format json` for an initial turn and `grok --resume <session> -p <prompt> --output-format json` for later turns. The generic adapter resolves the `grok` executable with `exec.LookPath`, parses JSON output, extracts a session-like ID, and owns the launched process group. These are production stubs, not a reachable evaluation route, because the runtime gate rejects Grok Build first. The local shell probe found no `grok` executable.

**Evidence:** Production source in `cli/internal/replay/process.go` contains the `HarnessGrok` command spec and generic `resolveExecutable`, `parseEvents`, `findSessionID`, and process close logic. Production source in `cli/internal/evaluation/runtime.go` blocks the path. The direct-installation report states that `grok inspect --json` should prove discovery before accepting installation; `cli/README.md` says the ordinary installer has not implemented that check.

**Implication:** Treat the command/resume/protocol shape as **production source but unqualified**, executable absence as **validation evidence**, and `grok inspect --json` as a **prior-research requirement not yet implemented**. A technical deep dive must verify the current Grok CLI's headless flags, session identifier/event schema, completion signal, cancellation semantics, and whether a clean process environment can retain supported authentication without importing ambient skills or settings.

### 4. OpenCode has the strongest local command evidence, but its present CLI path is still unreachable and unisolated

The installed OpenCode `1.14.39` exposes `opencode run [message..]` with `--session`, `--format json`, `--model`, `--variant`, and `--pure`; it also exposes `opencode debug skill`. This corroborates parts of the generic production command stub: `opencode run --format json <prompt>` followed by `opencode run --session <id> --format json <prompt>`. It also shows an available native plugin-exclusion control (`--pure`) and a dangerous permission-bypass flag that must not be used for qualification. The production CLI does not currently pass either flag, create a clean OpenCode runtime, or invoke `debug skill`; it is stopped by the common runtime gate first.

**Evidence:** Validation evidence from the local help probes establishes the actual `1.14.39` command surface. Production source in `cli/internal/replay/process.go` defines `HarnessOpenCode` using `runArgs`/`runResumeArgs`, uses `--model` for non-Codex/Cursor/Claude candidates, and has no OpenCode-specific completion validator. `cli/internal/evaluation/runtime.go` has no OpenCode branch. The direct-installation report records a prior-research floor of `v1.0.186+`, native skill roots, `permission.skill`, and a caveated `debug skill` surface.

**Implication:** Local OpenCode noninteractive/session/JSON capability is **validation evidence**; the version floor and discovery/permission claims are **prior research pending official re-check**. OpenCode is the only candidate with a locally discovered executable, but the present runner has no proof of clean/default isolation, discovery after temporary installation, stable raw-event completion semantics, attribution, cancellation, or a provider/authentication setup for any required model cell.

### 5. Kilo Code is represented only by registry/replay stubs and prior research; no local executable is available

The production replay map proposes the same programmatic `run --format json` and `run --session <id> --format json` pattern as OpenCode for Kilo Code. It supplies no Kilo-specific runtime branch, output validator, model/reasoning mapping, permissions profile, or authentication preflight. The local shell probe found no `kilo` executable, so the proposed command form has no local executable confirmation.

**Evidence:** Production source in `cli/internal/harness/harness.go` registers the native Kilo roots and executable, while `cli/internal/replay/process.go` registers the generic `run` commands. `cli/internal/evaluation/runtime.go` rejects the candidate before replay. The direct-installation report is prior research for `.kilo/skills`, `kilo --version`, `/reload`, the `skill` tool observation, permissions, and the local-versus-Cloud path conflict; the sub-agent report is prior research for `task` and `permission.task`.

**Implication:** All Kilo CLI execution, session, protocol, permission, and cleanup requirements remain **unresolved** locally. Later source work must establish the exact current Kilo CLI noninteractive syntax, JSON schema, session continuation, default isolation controls, safe permission configuration, and cancellation behavior before the generic replay stub can be treated as candidate evidence.

### 6. Generic process ownership helps with cancellation, but it does not establish candidate lifecycle completion

The generic process adapter creates an owned process group, waits for the command, stops the owned group after success or failure, and stops any pending group from `Close`. The replay runner also defers session close. The generic evaluation cleanup restores preexisting canonical Skill Issue directories or removes temporary directories, deletes private token mappings, and removes its private runtime directory. These mechanics are concrete production code, but candidate runs never reach them because runtime preparation fails first; no local smoke demonstrates that a Grok, OpenCode, or Kilo child process/session respects this ownership model.

**Evidence:** Production source in `cli/internal/replay/process.go`, `cli/internal/replay/replay.go`, `cli/internal/replay/process_group_unix.go`, `cli/internal/evaluation/evaluation.go`, and `cli/internal/installer/installer.go`. The retained smoke report validates comparable cleanup only for Codex, Cursor, Claude Code, and Pi.

**Implication:** Process-group cancellation and directory cleanup are **production source** at the generic layer, while candidate compatibility is **inference/unresolved**. A later probe must show that each candidate returns after a turn, exposes a durable session ID, terminates child/background work when its parent group is stopped, and leaves no candidate-owned temporary session/server state outside the bounded cleanup owner.

### 7. Candidate protocol validation is generic and under-specified

All candidate stubs must emit nonempty JSON objects or a JSON array; generic parsing rejects malformed output and generic session extraction searches `session_id`, `sessionId`, `thread_id`, or `threadId`. The adapter rejects explicit error events and enforces that a discovered session ID stays stable. Concrete completion-event validation exists only for Codex, Cursor, and Claude Code; Grok, OpenCode, and Kilo Code receive no candidate-specific completion marker check. Skill attribution is also Codex-specific in the evaluation callback, so candidate JSON output would currently not establish opaque-signal attribution.

**Evidence:** Production source in `cli/internal/replay/process.go` implements `parseEvents`, `findSessionID`, `validateSessionID`, and `validateHarnessOutput`; its switch has completion requirements only for Codex, Cursor, and Claude Code. `cli/internal/evaluation/evaluation.go` calls `recordCodexSignals` only when the selected harness is Codex.

**Implication:** The candidate protocol condition is **production-source scaffolding**, not a qualified protocol contract. Later research must obtain exact official/current event examples and prove initial/resume completion, tool/permission failures, model rejection, cancellation, and a candidate-visible yet private signal-attribution mechanism.

### 8. Model, reasoning, authentication, and Codex-subscription support are missing for these candidates

No candidate has an `EvaluationDefaults` entry, so `ResolveRequest` cannot select defaults. If a candidate could bypass that gate, the generic process adapter would append `--model <value>` to all three, but it does not translate `--reasoning` to a Grok, OpenCode, or Kilo native option. Runtime authentication status is checked only for Codex and Cursor. The only candidate-specific local model evidence is OpenCode help showing `--model` and `--variant`; it does not demonstrate provider availability, authentication, supported reasoning values, or any Codex subscription access. The explicit release matrix makes OpenCode a required minimum route for Codex Sol, Claude Opus 4.8, and Claude Fable, but it reserves Grok Build and Kilo Code for later expansion.

**Evidence:** Production source in `cli/internal/harness/harness.go`, `cli/internal/evaluation/evaluation.go`, `cli/internal/evaluation/runtime.go`, and `cli/internal/replay/process.go`. Local OpenCode help is validation evidence. The minimum matrix and tier appear in `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md` and the parent completion plan.

**Implication:** Candidate model/reasoning maps, normal login ownership, provider configuration, and every Codex-subscription claim are **unresolved**. Later research must distinguish officially supported provider/API-key mechanisms from copied credentials, home-directory substitution, permanent configuration edits, and undocumented proxies. No local source supports a claim that Grok Build, OpenCode, or Kilo Code can use a Codex subscription.

### 9. Existing candidate research is useful source targeting, but it is prior research rather than current qualification evidence

The local corpus identifies focused official-source targets for two later gates. For Grok Build: xAI Build overview, CLI reference, skills/plugins, settings, enterprise policy, and the xAI Grok Build repository's sub-agent guide/tool schema. For OpenCode: official Agent Skills, CLI, config, agents, and permissions documentation; `anomalyco/opencode` skill discovery/tool/config source; and released version history. For Kilo: official Skills, CLI, Agent Permissions, Custom Subagents/tool documentation, Cloud Agent documentation, and the Kilo Code repository/releases. Manual invocation records separate deterministic slash invocation (Grok) from agent-mediated natural-language selection (OpenCode and Kilo). Sub-agent records identify possible task/delegation tooling, subject to candidate permissions.

**Evidence:** The direct-installation, manual-invocation, and sub-agent assignment files listed in Sources preserve URLs, source revisions, commands, and caveats. Their reports also flag Grok's loose-skill schema/collision gaps, OpenCode's source-current-but-not-publicly documented surfaces, and Kilo's local `.kilo` versus Cloud `.kilocode` conflict.

**Implication:** Classify these as **prior research** and use them as a ranked source map, not current proof. The next technical deep dives should re-check only the named first-party documentation/repositories and the locally installed OpenCode `1.14.39` behavior where it differs from the earlier reported `v1.18.3` checkpoint.

### 10. Documentation and production source conflict in ways that must stay visible

`cli/README.md` describes headless/resumable or programmatic adapter forms for Grok Build, OpenCode, and Kilo Code, but it also explicitly limits current evaluation support to four harnesses. The source confirms the latter: candidates are blocked before runtime preparation. The README's phrase that adapters are implemented from automation contracts therefore describes partial registry/replay scaffolding, not a complete candidate evaluator. Separately, OpenCode's locally installed `1.14.39` is older than the `v1.18.3` release checkpoint recorded by prior research, and Kilo's prior research records incompatible local and Cloud skill roots.

**Evidence:** `cli/README.md`, `cli/internal/harness/harness.go`, `cli/internal/evaluation/runtime.go`, and the local OpenCode version probe. The direct-installation reports contain the older research checkpoints and caveats.

**Implication:** Mark the README adapter language as **stale/conflicting if read as qualification**, retain the source-backed four-harness evaluator boundary, and avoid transferring local OpenCode observations to newer versions or local Kilo paths to Kilo Cloud. These are evidence-classification issues, not a final support decision.

## Notes

- **Candidate prerequisite matrix:** Grok Build — `grok` executable, official headless/resume JSON contract, authenticated clean process, ambient-exclusion strategy, `grok inspect --json` discovery gate, permissions/sandbox-safe signal, stable session/events, and bounded cancellation/cleanup. OpenCode — the observed `1.14.39` executable plus official version/model/provider/auth validation, `--pure`/other isolation semantics, native temporary-skill discovery, `debug skill` behavior, permission-safe signal, JSON/session completion, and bounded server/session cleanup. Kilo Code — executable/version, exact noninteractive CLI protocol, model/provider/auth and reasoning mechanism, local-only skill discovery/reload, default isolation, permission/task behavior, session/event semantics, and cleanup.
- **Unsupported local claims:** no local Grok or Kilo executable; no candidate live model session; no candidate discovery/activation proof; no candidate signal attribution; no candidate cancellation/cleanup proof; no candidate authentication result; and no Grok/OpenCode/Kilo Codex-subscription proof.
- **Out-of-scope but relevant:** the product's required release tier includes OpenCode; Grok Build and Kilo Code are expansion candidates. This assignment does not recommend retaining, removing, or implementing any candidate.
