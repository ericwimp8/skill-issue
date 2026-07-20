# Local GitHub Copilot CLI and Gemini CLI Candidate Evidence

## Assignment

**Goal:** Map all existing local evidence relevant to technical qualification of GitHub Copilot CLI and Gemini CLI, separating runnable production behavior from contracts, prior research, validation evidence, inferences, stale or conflicting material, and unresolved requirements.

**Scope:** Repository production sources, CLI contracts, setup and completion documents, retained direct-installation/manual-invocation/sub-agent research, smoke records, and a read-only executable-path probe. This map identifies source targets for later official-web research; it does not decide candidate support.

**Exclusions:** Internet research, modification of candidate adapters, installation or authentication of either candidate, and a final qualification decision.

## Sources

- **Production source:** `cli/internal/harness/harness.go` — candidate IDs, executable names, skill roots, and evaluation allow-list.
- **Production source:** `cli/internal/evaluation/evaluation.go`, `cli/internal/evaluation/runtime.go`, and `cli/internal/installer/installer.go` — evaluation entry, runtime preparation, temporary skills, cleanup, and attribution ownership.
- **Production source:** `cli/internal/replay/replay.go` and `cli/internal/replay/process.go` — process adapter, headless argument stubs, process ownership, session extraction, protocol parsing, and cancellation.
- **Local contract:** `cli/README.md`, `plans/harness-setup.md`, `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`, and `skills/skill-evaluation-and-refinement/SKILL.md` — current user-facing boundary and qualification gates.
- **Prior research:** `research/deep-research/harness-direct-installation-architecture/assignments/01-github-copilot-direct-installation.md` and `05-google-antigravity-gemini-cli-direct-installation.md` — first-party findings retained locally, inspected 2026-07-19.
- **Prior research:** `research/deep-research/manual-skill-invocation-nine-harnesses/assignments/01-github-copilot.md` and `05-google-antigravity-gemini-cli.md`; `research/harness-subagent-launch/assignments/01-github-copilot.md` and `05-google-antigravity-gemini-cli.md` — explicit invocation and delegated-work findings.
- **Validation evidence:** `evaluations/skill-calling/smoke/real-harness-smoke-report.md` and the 2026-07-21 local shell probe: neither `copilot` nor `gemini` resolves on `PATH` in this environment.

## Findings

### Candidate Registration Is Installation-Only

**Classification:** Production source.

`harness.Spec` registers Copilot as executable `copilot` with project root `.github/skills` and user root `~/.copilot/skills`; it registers Gemini CLI as `gemini` with `.gemini/skills` and `~/.gemini/skills`. The general `ParseID` path accepts both, so ordinary installation and uninstall can resolve their native roots. `evaluationDefaults`, however, contains only Claude Code, Codex, Cursor, and Pi. `ParseEvaluationID`, which is called by the evaluate command path, rejects both candidates before runtime preparation.

**Evidence:** `cli/internal/harness/harness.go` defines both `Spec` entries, while `evaluationDefaults`, `ParseEvaluationID`, and `EvaluationDefaultsFor` exclude them. `cli/internal/lifecycle/lifecycle.go` routes `evaluate run` through `ParseEvaluationID`; `cli/README.md` states that installation supports nine harnesses and evaluation currently supports only four.

**Implication:** The local executable names and project skill roots are production-backed installation data. Neither candidate currently owns an accepted evaluation request, a default model, or a default reasoning setting.

### Replay Contains Unqualified Headless Stubs

**Classification:** Production source for the existence of stubs; inference for runnable capability.

The shared process adapter has entries for both candidates. Copilot initial turns use `copilot -p <prompt>` and resumed turns use `copilot --resume <session-id> -p <prompt>`. Gemini initial turns use `gemini -p <prompt> --output-format stream-json`; resumed turns use `gemini --resume <session-id> -p <prompt> --output-format stream-json`. The adapter resolves an executable with `exec.LookPath`, runs one owned process per turn, retains stdout/stderr, parses JSON objects when available, and requires a stable session ID. Copilot alone can fall back from invalid structured stdout to a synthetic `result` event and then searches stdout/stderr for a `--resume` suggestion.

**Evidence:** `cli/internal/replay/process.go` defines `commandSpecs`, `SendPrompt`, `Wait`, `findSessionID`, `findCopilotSessionID`, `resolveExecutable`, and `validateSessionID`; `cli/internal/replay/replay.go` defines sequential multi-turn replay.

**Implication:** Later official-source deep dives must verify exact noninteractive syntax, output modes, emitted session identifiers, and resume semantics for the installed versions. Copilot's permissive non-JSON fallback is an especially important local behavior to audit: it is implementation tolerance, not evidence that the candidate supplies a qualifying native trace.

### Candidate Isolation Has No Runtime Owner

**Classification:** Production source and unresolved gap.

`prepareRuntime` has bespoke preparation only for Codex, Cursor, Claude Code, and Pi. Its non-Codex branch creates a private root, but the switch returns `unsupported evaluation harness` for Copilot and Gemini. Therefore neither candidate has a candidate-specific `HOME`, config/state root, temporary directory layout, controlled environment, or documented ambient-skill exclusion in the current code. The shared process adapter would merge the ambient environment for each candidate because only Cursor and Pi set `CleanEnvironment`.

**Evidence:** `cli/internal/evaluation/runtime.go` contains only `prepareCursorRuntime`, `prepareClaudeRuntime`, and `preparePiRuntime`; `cli/internal/evaluation/evaluation.go` sets `CleanEnvironment` only for Cursor and Pi. `controlledEnvironment` exists but is not invoked for either candidate.

**Implication:** Isolation, ambient-skill/configuration exclusion, workspace trust, model access, and home-directory credential behavior are unresolved for both candidates. An official deep dive must establish whether an isolated home can preserve each CLI's supported authentication or whether the default-environment attestation contract is the only safe path.

### Temporary Skills and Cleanup Are Generic but Unproven Per Candidate

**Classification:** Production source for generic lifecycle; unresolved candidate behavior.

The evaluation installer can materialize instrumented skills at any supplied project root, preserves pre-existing canonical Skill Issue payload names, removes temporary-only names during cleanup, and restores ordinary copies when required. The evaluation service retains per-run installation state, cleans on success/error, deletes private runtime material, and the replay session kills its owned process group on close. Since candidate evaluation requests cannot pass the allow-list and neither candidate has runtime preparation, this lifecycle has never reached Copilot or Gemini through the production entry path.

**Evidence:** `cli/internal/installer/installer.go` defines `PrepareEvaluation`, `CleanupEvaluation`, and complete-directory materialization; `cli/internal/evaluation/evaluation.go` persists installation state and defers cleanup; `cli/internal/replay/process.go` owns process termination in `Close` and after `Wait`.

**Implication:** Temporary project skill placement can be reused only after host discovery, trust/activation, signal permission, cancellation, and cleanup behavior are proven on the exact candidates. No local evidence proves that killing the foreground CLI also cleans candidate-spawned workers, preserves a resumable session, or removes candidate-created state.

### Attribution Is Codex-Specific

**Classification:** Production source and unresolved gap.

Skill Issue's generated evaluation skills signal activation through a private-state command, but the evaluation service reads native replay events only for Codex: `recordCodexSignals` accepts a structured command-execution attempt as activation evidence. No Copilot or Gemini event parser, trace extractor, marker collector, or equivalent attribution route exists. The generic capture stores raw stdout/stderr/events but has no candidate-specific semantics beyond process success and session continuity.

**Evidence:** `cli/internal/evaluation/evaluation.go` calls `recordCodexSignals` only when `request.Harness == harness.Codex`; `cli/internal/replay/process.go` has strict output checks only for Codex, Cursor, and Claude.

**Implication:** A later deep dive must find a native, persistent, candidate-identifying skill-load/activation record. A successful prompt, final prose, the generic Copilot synthetic event, or a signal command alone is insufficient under the evaluator's native-evidence requirement.

### Authentication and Codex-Subscription Access Are Unowned

**Classification:** Production source and unresolved gap.

The only preflight authentication checks are Codex (`login status`) and Cursor (`status`). The candidate replay stubs can accept an executable override and inherit ambient environment, but provide no authentication check, account/provider boundary, token policy, supported model/provider selection, or Codex-subscription mechanism. The current default-model map likewise omits both candidates.

**Evidence:** `cli/internal/replay/process.go` limits `CheckAuthentication` to `HarnessCodex` and `HarnessCursor`; `cli/internal/harness/harness.go` limits `evaluationDefaults` to four other harnesses. The smoke report records only Codex ChatGPT, Cursor Keychain, Claude proxy, and Pi authentication.

**Implication:** Treat all claims that Copilot or Gemini can use a Codex subscription, Codex OAuth, copied credentials, a proxy, or a particular model/reasoning tier as unsupported locally. Official-web research must separate each candidate's supported account model from any proposed access bridge and keep Skill Issue outside credential transfer or configuration mutation.

### Local Executable Evidence Blocks Live Qualification Here

**Classification:** Validation evidence.

On 2026-07-21, `command -v copilot` and `command -v gemini` returned no executable on the current `PATH`; no version/help, authentication, isolated launch, protocol, permission, cancellation, cleanup, session-resume, or trace probe was attempted. The retained real-harness smoke report records only Codex, Cursor, Claude Code, and Pi, and explicitly limits its qualification to two-turn smokes rather than broad campaign proof.

**Evidence:** Read-only local shell probe; `evaluations/skill-calling/smoke/real-harness-smoke-report.md` sections “Qualified Launchers,” “Scope,” and “Cleanup Evidence.”

**Implication:** There is no local live candidate evidence. The absence is an environment limitation rather than evidence about either product's documented capabilities.

### Direct Installation and Explicit Invocation Are Prior Research, Not Runtime Proof

**Classification:** Prior research; some time-sensitive claims require revalidation.

The retained Copilot research states that direct skills use `.github/skills/<name>/SKILL.md` or `~/.copilot/skills/<name>/SKILL.md`, can reload/list/info skills, and support `/skill-name` manual invocation when user-invocable; it records trust, organization-policy, and `allowed-tools` constraints. The retained Gemini research records `.gemini/skills` and `~/.gemini/skills`, `/skills reload` and `/skills list`, matching-task activation consent, workspace trust, `skills.enabled`, collision precedence, and natural-language explicit naming rather than a documented per-skill slash invocation. Both direct-installation documents preserve official URLs for later re-checking.

**Evidence:** `research/deep-research/harness-direct-installation-architecture/assignments/01-github-copilot-direct-installation.md`; `research/deep-research/harness-direct-installation-architecture/assignments/05-google-antigravity-gemini-cli-direct-installation.md`; `research/deep-research/manual-skill-invocation-nine-harnesses/assignments/01-github-copilot.md`; `research/deep-research/manual-skill-invocation-nine-harnesses/assignments/05-google-antigravity-gemini-cli.md`.

**Implication:** These findings support source-target selection for discovery/trust/activation research, but do not prove noninteractive evaluation compatibility, protocol output, signal permission, session persistence, or cleanup in the current unavailable local binaries.

### Sub-Agent Claims Do Not Establish Evaluation Isolation

**Classification:** Prior research; inference if applied to this runner.

Local research reports that Copilot CLI can use model-driven delegation and `/fleet`, while Gemini CLI can expose named sub-agents and accept explicit `@name` requests. These descriptions concern in-session agent delegation. They do not show that the outer CLI launch is independent, that a temporary skill is preloaded in a child, that a child emits native skill evidence, or that child processes/session state are controlled by the current process adapter.

**Evidence:** `research/harness-subagent-launch/assignments/01-github-copilot.md` and `research/harness-subagent-launch/assignments/05-google-antigravity-gemini-cli.md`; the qualification and native-evidence requirements in `skills/skill-evaluation-and-refinement/SKILL.md`.

**Implication:** Later research must distinguish native delegated-work capability from the required fresh independent-agent or isolated-equivalent evaluation route. No candidate can inherit the qualified Codex/Cursor/Claude/Pi implementation merely because it offers sub-agents.

### Local Contracts Set the Qualification Bar and Product Boundary

**Classification:** Local contract.

The current product contract makes Copilot and Gemini additional selected targets outside the five-harness minimum tier. The evaluator requires a qualified exact environment, implicit invocation where needed, a fresh independent agent or isolated equivalent, and native evidence identifying the skill before description trials. The CLI README further labels adapters on unavailable harnesses unqualified and calls launch, permission, session, marker, timeout, cancellation, or protocol failure a tooling failure.

**Evidence:** `plans/skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md`; `skills/skill-evaluation-and-refinement/SKILL.md`; `cli/README.md` “Harness Boundary.”

**Implication:** The later technical investigation must cover default/isolated environment, temporary skill discovery, noninteractive mode, resumable session, native protocol/activation evidence, permissions, cancellation, cleanup, executable discovery, model/reasoning configuration, and authentication separately. Passing only a direct-install or manual-invocation check cannot qualify either candidate.

### Time-Sensitive and Conflicting Local Material Needs Reconciliation

**Classification:** Stale/conflicting prior research and unresolved.

The direct-installation research treats Gemini CLI as the selected Google adapter and records a current v0.50.0 skill feature claim, while the manual-invocation research says a May 2026 transition made Antigravity CLI the general consumer terminal surface and retained Gemini CLI only for enterprise/paid API-key access. The active research map still names Gemini CLI as this campaign candidate, and the CLI README says Antigravity automation is outside the runner. Copilot's retained direct-installation research is explicitly current only as of 2026-07-19 and marks `gh skill` preview. No live binary is available to resolve either product-version/access state locally.

**Evidence:** `research/deep-research/harness-direct-installation-architecture/assignments/05-google-antigravity-gemini-cli-direct-installation.md`; `research/deep-research/manual-skill-invocation-nine-harnesses/assignments/05-google-antigravity-gemini-cli.md`; `research/harness-portability-qualification/research-map.md`; `cli/README.md`; and the Copilot assignment's “Currentness caveat.”

**Implication:** Official-web deep dives must first confirm the exact currently supported Copilot CLI and Gemini CLI distributions, auth/access eligibility, feature/version gates, and documented automation mode. The campaign should retain Gemini as the named research target until that evidence is reconciled, without transferring Antigravity behavior into Gemini.

## Notes

- **Official-web targets — Copilot:** GitHub Copilot CLI command/reference pages for `-p`, `--resume`, output/streaming protocol, session inspection, noninteractive authentication, trust/approval, skill-load diagnostics, `/fleet`, and cancellation; GitHub CLI `gh skill` remains a separate installation lifecycle, not evidence for the `copilot` executable.
- **Official-web targets — Gemini:** Gemini CLI reference and primary `google-gemini/gemini-cli` releases/source for `-p`, `--output-format stream-json`, `--resume`, session storage, noninteractive/auth modes, trusted folders, skills activation/consent, sub-agent isolation, structured traces, model/reasoning controls, process cancellation, and cleanup; verify product/access status before relying on older docs.
- **Unsupported locally:** Native skill-load trace, candidate-specific temporary-home strategy, ambient exclusion, safe signal permission, strict protocol schema, subprocess cleanup, session deletion/retention policy, model/reasoning values, account support, and every Codex-subscription claim.
- **No final support decision:** This assignment maps evidence and gaps only.
