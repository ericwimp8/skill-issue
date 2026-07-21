# Skill Issue Local Testing Methodology Research

## Best-Supported Answer

The current development working tree defines a governed skill-calling methodology with three **30-turn** scenarios. At the inspected development snapshot, the gardening and archive scenarios each contain **44 unique expected turn-skill pairs**, the preparedness scenario contains **43**, and each scenario has expectations on **27 of its 30 turns**. These totals are current observations rather than frozen campaign constants because the scenario sources are still being edited and require final re-verification.

Thirty has two separate meanings in the current system:

- each governed scenario contains 30 conversational turns; and
- the planned campaign contains 30 full runs: ten harness/model configurations multiplied by three scenarios.

The CLI replays each scenario's fixed user prompts in order through one continued primary-agent session. It instruments disposable copies of the supplied skills with an opaque `signal` call, maps signals to skills outside model context, attributes them to the evaluator-owned active turn, and derives detailed and compact artifacts from the private answer sheet.

The defensible public account is therefore a description of the **current development methodology**, with explicit qualifications:

- the scenario inventory and scorecard totals must be re-counted after the development sources freeze;
- the selected known-good executable, committed `HEAD`, index snapshots, and earlier 12-turn analyses are stale relative to the intended method and cannot control current conclusions;
- that development-versus-known-good divergence blocks campaign freeze and publication, but it is not a reason to roll the scenarios back;
- Dictate Plan is named explicitly once in Turn 1 of every scenario and is scored;
- harness isolation is adapter-specific and retains documented product or host surfaces;
- the eight skills have bounded hardening evidence, while complete intake-to-generation lineage is not retained for every skill;
- no accepted public campaign artifacts currently exist, so observed cross-harness comparisons remain unpublished.

## Claim Validation

| Claim                                                                             | Current development truth                                                                                                                                                    | Qualification                                                                                                                                                       | Evidence                                                                                                                                                                                 |
| --------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| The method uses three 30-turn scenarios.                                          | Verified. All three governed built-ins contain 30 ordered user turns.                                                                                                        | Re-verify after scenario freeze because the development files are actively changing.                                                                                | `evaluations/skill-calling/built-ins/*.json`; `src/data/siteData.ts:75-83`                                                                                                               |
| Each scenario has approximately 15 to 20 expected skill calls.                    | Rejected. The inspected development snapshot contains 44, 44, and 43 unique expected turn-skill pairs.                                                                       | The final totals remain freeze-gated. “Substantially more than 15 to 20” is currently supported.                                                                    | `evaluations/skill-calling/built-ins/*.json`                                                                                                                                             |
| Every turn is scored.                                                             | Rejected. Each current scenario has expectations on 27 of 30 turns.                                                                                                          | A scored turn can expect more than one skill.                                                                                                                       | `evaluations/skill-calling/built-ins/*.json`                                                                                                                                             |
| Thirty means the campaign size.                                                   | Verified but incomplete. The campaign has ten configurations across three scenarios, for 30 runs.                                                                            | Thirty also means turns per scenario. Keep the two units explicit.                                                                                                  | `plans/skill-calling-evaluation-campaign/evaluation-progress.md:5-23`                                                                                                                    |
| Scorecards map turns to expected skills.                                          | Verified. Each built-in contains a private `answer_sheet.expected` collection of `{turn_id, skill}` pairs.                                                                   | Current pairs are unique within each scenario; re-check uniqueness at freeze.                                                                                       | `evaluations/skill-calling/built-ins/*.json`; `cli/internal/evaluation/evaluation.go:960-1059`                                                                                           |
| The CLI conducts controlled back-and-forth conversations.                         | Verified with qualification. One primary-agent session receives fixed user prompts verbatim and in order, and each turn completes before the next prompt is sent.            | User prompts are governed; assistant responses, tools, and workspace effects remain harness/model outputs.                                                          | `cli/internal/replay/replay.go:111-157`; `cli/internal/replay/process.go:137-192`; `cli/internal/replay/pi.go:180-236`                                                                   |
| Evaluated skills receive opaque recording calls translated outside model context. | Verified with qualification. Disposable skill copies receive a `signal` instruction whose token maps to a skill in private run state.                                        | The injected command and paths remain visible to the agent, so the method is attributable and minimally semantic rather than invisible or influence-free.           | `cli/internal/installer/installer.go:516-558`; `cli/internal/runstate/runstate.go:77-187`; `evaluations/skill-calling/instrumentation-contract.md`                                       |
| Runs use isolated, sandboxed, default-like environments.                          | Directionally supported. The evaluator creates controlled runtime state and the campaign requires fresh external workspaces.                                                 | The exact residuals differ by harness; a uniform confinement claim is unsupported until each lane passes its required probe.                                        | `cli/internal/evaluation/runtime.go`; `cli/internal/replay/process.go`; `cli/internal/replay/pi.go`; `plans/skill-calling-evaluation-campaign/evaluation-orchestration-prompt.md:93-116` |
| Skill generation and refinement harden the campaign skills.                       | Partly verified. The repository defines intake, generation, and two-loop refinement workflows, and the eight scenario skills have retained bounded refinement evidence.      | Complete current-workflow intake-to-generation provenance is not retained for every one of the eight skills; retained cases do not establish universal reliability. | `skills/skill-intake/`; `skills/skill-generation/`; `skills/skill-evaluation-and-refinement/`; `evaluations/scenario-skill-refinement/`                                                  |
| Dictate Plan is manually invoked exactly once and excluded from scoring.          | First clause verified; second rejected. Every Turn 1 begins with “Let's do Dictate Plan,” no later prompt names it, and every answer sheet expects `dictate-plan` on Turn 1. | The explicit prompt changes interpretation: this measures compliance with a named startup request, not spontaneous discovery.                                       | Turn 1 and answer sheet of each governed built-in                                                                                                                                        |
| Retained artifacts currently support public result inspection.                    | Rejected for campaign results. The accepted public collection is empty.                                                                                                      | Local refinement evidence exists, but campaign publication, detailed evidence links, and provenance-aware acceptance remain incomplete.                             | `src/data/publishedWebsiteArtifacts.json`; `src/data/evaluationData.ts:230-237`; `scripts/update-website-results.mjs:13-42`                                                              |

## Governed Evaluation Inputs

### Scenario inventory

| Built-in ID                                   | Turns | Current expected pairs | Expectation-bearing turns |
| --------------------------------------------- | ----: | ---------------------: | ------------------------: |
| `gardening-web-application`                   |    30 |                     44 |                        27 |
| `community-archive-desktop-application`       |    30 |                     44 |                        27 |
| `neighborhood-emergency-preparedness-program` |    30 |                     43 |                        27 |

These numbers were calculated directly from the current development built-ins. They describe the inspected snapshot and must be regenerated after the files stop changing. The executable inputs are the built-in JSON files embedded by `bundle.go:5` and loaded through `cli/internal/payload/payload.go:120-126`. The corresponding Markdown files under `evaluations/skill-calling/scenarios/` are human-readable governance views; replay consumes the embedded JSON.

### Expected-skill inventory

| Skill                            | Gardening | Archive | Preparedness |
| -------------------------------- | --------: | ------: | -----------: |
| `code-implementation-discipline` |        13 |      13 |           14 |
| `code-testing-discipline`        |         9 |       9 |            9 |
| `dictate-plan`                   |         1 |       1 |            1 |
| `document-update-discipline`     |         9 |       9 |            7 |
| `prompt-writing`                 |         5 |       5 |            5 |
| `skill-authoring-discipline`     |         2 |       2 |            2 |
| `system-change-ownership`        |         2 |       2 |            2 |
| `systematic-debugging`           |         3 |       3 |            3 |

The per-skill totals sum to 44, 44, and 43. Dictate Plan's single expected call is part of each total.

### Canonical skill sources

The embedded manifest, rather than a directory naming convention, owns each component's source:

- `dictate-plan` comes from `supporting-skills/dictate-plan`;
- the other seven scenario skills come from `evaluations/scenario-skill-refinement/<skill>/skill`.

`bundle.go:5` embeds those two source families. `cli/internal/payload/assets/manifest.json:18-49` records each ID-to-source mapping, and `cli/internal/payload/payload.go:43-70,129-181` validates and reads those mapped sources.

At the inspected snapshot, the canonical `SKILL.md` hashes are:

| Skill                            | SHA-256                                                               |
| -------------------------------- | --------------------------------------------------------------------- |
| `code-implementation-discipline` | `c556e5461a5202106822cf0e54906b561f225cb93729f06a008bf5af7bf3fb66`    |
| `code-testing-discipline`        | `72277b024cd12d7875792e3a677a73c9a13f294d266231c6fe5876930cfbdd55`    |
| `dictate-plan`                   | `44bc9abeeebc2200ec5c3b6ea7a65cab70551e72c1b9f7c06c281c2b2ec0fe2b`    |
| `document-update-discipline`     | `03668f94fc4b0d3b5293bd2c5212c20563cb340a158b4daa99274c91af289ca4`    |
| `prompt-writing`                 | `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`    |
| `skill-authoring-discipline`     | `7fcb2f074490433c6eb6dc3f288d3af1caf353f1cd86b8c65be91d6f017d8bb4`    |
| `system-change-ownership`        | `28659276d4de61422a233f6e909c075a611bab29c9c6fe5357db3fcd22fc3c82`    |
| `systematic-debugging`           | `008eb3082ab95495495453702d616dd82834f38dcd0639aa1d0d00ecf493254b0af` |

These hashes support source identity for this inspection. The campaign baseline still needs a frozen executable/version that embeds the same sources.

## Conversation Execution

`replay.Runner.Run` owns the common ordered sequence:

1. validate the scenario;
2. start one primary-agent session;
3. mark the next turn active;
4. send the exact governed prompt;
5. wait for terminal completion;
6. capture session and harness output;
7. close the active-turn window; and
8. repeat through Turn 30 using the same session.

The adapters use their native continuation mechanisms: Codex resume, Cursor `--resume`, Claude Code `--resume`, OpenCode/Kilo session identifiers, and one long-lived Pi RPC process. Full campaign runs omit `--turns`; truncation is a smoke/custom feature and changes the evaluated unit.

Governed surfaces include the scenario ID, prompt order, private answer sheet, supplied skill payload, selected harness/model/reasoning target, active-turn boundaries, and result derivation. Variable surfaces include assistant text, call order, tools, workspace effects, native compaction, provider behavior, and session implementation.

## Instrumentation And Attribution

The evaluator materializes disposable skill copies and injects one short instruction immediately after frontmatter. The instruction asks the agent to run `skill-issue signal <opaque-token>` when it uses the skill. Private run state maps each token back to one skill and records the evaluator-owned active turn.

This supports turn-level external attribution without placing the private answer sheet or skill name-token map in model context. It does not prove that instrumentation has zero behavioral effect: the agent can see the instruction, executable path, and state-related command surface. Public wording should use “opaque external attribution” or “minimally semantic signal,” accompanied by that limitation.

Scoring compares the observed turn-skill pairs with the private expected pairs. Expected-and-observed pairs are called; expected-and-unobserved pairs are missed; observed pairs outside the answer sheet are additional. Tooling or protocol failures affect run acceptance rather than being silently reclassified as missed skills.

## Dictate Plan

Dictate Plan is a distinct canonical component at `supporting-skills/dictate-plan`. Every governed scenario names it in ordinary language exactly once at Turn 1, and every answer sheet expects it there. It is therefore:

- an explicit startup instruction;
- instrumented like the other supplied skills;
- included in scoring; and
- evidence of compliance with an explicit invocation, not evidence of proactive discovery.

Its Codex metadata sets `allow_implicit_invocation: false`, which supports an explicit-only Codex surface. The portable skill bundle and other harness adapters do not establish an equivalent universal product policy.

The retained Dictate Plan status remains useful but historically qualified. `evaluations/skill-system-production-refinement/targets/dictate-plan/status.md` points to the stale target path `test-targets/skills/dictate-plan/SKILL.md`, while its retained content hash matches the current canonical skill. It also records a legacy 40-iteration campaign allowance. Current workflow authority instead pauses after five unsuccessful description rounds and stops after five unsuccessful body cycles unless the user authorizes continuation (`skills/skill-evaluation-and-refinement/SKILL.md:24,40`). The retained status is evidence of the earlier campaign, not authority for current iteration policy or source location.

## Harness Environment Controls

The campaign contract requires a fresh empty workspace outside this repository for each run, separate output storage, explicit executable/model/reasoning selection, and lane-specific confinement probes. Production adapters add controlled runtime configuration, but their remaining surfaces differ.

| Harness     | Implemented controls                                                                                                                                                                       | Residual surface that must remain explicit                                                                                                                                                    |
| ----------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Codex       | Disables plugins, ignores user config and rules, disables discovered ambient skills, uses `workspace-write`, and supplies evaluator configuration.                                         | Normal Codex authentication/session state remains necessary; the sandbox and approval behavior are Codex-managed.                                                                             |
| Claude Code | Restricts setting sources, disables project memory/marketplace/background features, supplies only the passed skill directory, limits allowed tools, and uses a generated workspace prompt. | Bundled Claude skills/commands may remain; normal authentication and Claude's managed permission policy remain product surfaces.                                                              |
| Cursor      | Uses a generated home/config/plugin, disables project configs, enables its sandbox, and allowlists evaluator-required operations.                                                          | Product-managed skills and account/team policy may remain; the generated runtime links the user's keychain for normal authentication.                                                         |
| OpenCode    | Uses generated XDG roots, disables plugins and Claude imports, and configures explicit skills and permissions.                                                                             | Provider authentication and the harness's own runtime enforcement remain relevant.                                                                                                            |
| Kilo        | Uses generated XDG roots, disables external/project skill sources and indexing, and configures only the supplied custom skill path.                                                        | Runtime exclusivity of that configured path has not yet been independently proven; provider authentication and host permissions remain.                                                       |
| Pi          | Disables extensions, ambient skills, templates, themes, context files, and sessions; explicitly enumerates supplied skills; sets `--offline`.                                              | Offline mode is a Pi content/update mode, not demonstrated network confinement. Bash and host filesystem permissions remain available, and an allowlist of provider credentials is forwarded. |

“Default-like” is defensible only when defined as a primary-agent run with known ambient customization removed and evaluator-required controls added. “Uniformly isolated” or “equivalently sandboxed” is unsupported. The orchestration contract correctly requires bounded lane probes rather than inferring confinement from a working directory or flag name.

## Skill Construction And Refinement

The repository has an explicit workflow chain:

1. `skill-intake` produces a build-ready A-to-B contract;
2. `skill-generation` builds the skill and hands it to evaluation; and
3. `skill-evaluation-and-refinement` runs separate description-selection and body-behavior loops with retained evidence and semantic-owner corrections.

The current scenario-skill campaigns support a bounded hardening claim. Their status and conclusion records retain natural selection trials, isolated behavior fixtures, native traces, audits, refinements where warranted, and evidence-level claim limits. Several conclusions explicitly reject universal routing or correctness claims outside the qualified model, harness, prompts, and fixtures.

The evidence does not establish complete generation lineage for every current scenario skill. Some retained campaigns begin from an existing target, and Dictate Plan retains an older source path. Public wording should say that the campaign skills were evaluated and, where evidence required it, semantically refined. It should describe the full intake-to-generation-to-refinement workflow as the current system workflow without implying that every retained target preserves that complete historical chain.

The current seven scenario-refinement statuses align their target paths with `evaluations/scenario-skill-refinement/<skill>/skill` and generally record the current five-failure boundaries. The legacy Dictate Plan status is the outlier described above.

## Artifacts And Publication

### Run artifacts and acceptance

The CLI writes a detailed result and a compact website artifact. Acceptance must also establish full-turn completion, one valid continued session, instrumentation completeness, private-state cleanup, temporary-skill cleanup, expected workspace behavior, correct model/reasoning resolution, and absence of run-owned processes. A clean exit alone is insufficient.

The planned campaign currently remains at 0/30 complete in `plans/skill-calling-evaluation-campaign/evaluation-progress.md:23-33`. Exact completion totals in that progress document may be transient while execution work is underway, but there is no accepted 30-run result set available for publication in the inspected tree.

### Website state

`src/data/publishedWebsiteArtifacts.json` is empty. Consequently `src/data/evaluationData.ts:230-237` falls back to illustrative data rather than accepted campaign evidence.

The illustrative layer is not aligned with the current scorecards:

- it contains only four sparse points at Turns 1, 11, 25, and 30 (`src/data/evaluationData.ts:169-200`), while each current scenario has 27 expectation-bearing turns;
- it retains the gardening label `GardenFlow planning` (`src/data/evaluationData.ts:45-58`), which is stale against the governed scenario presentation; and
- its `total_turns: 30` now agrees with the current development scenarios and should no longer be described as stale.

The public site copy in `src/data/siteData.ts:75-83` also says each scenario has 30 turns, which is aligned with current development intent. Publication still requires accepted artifacts generated from the frozen baseline.

### Import and provenance gap

`scripts/update-website-results.mjs:13-42` checks only a compact set of top-level types, sorts the supplied artifacts, and writes them directly into the accepted collection. It does not establish:

- correspondence with an accepted detailed `result.json`;
- campaign membership or a frozen baseline identity;
- harness and model version provenance;
- full-run and cleanup acceptance;
- reasoning-setting verification;
- scorecard/source hashes; or
- detailed evidence URLs.

Compact artifacts also omit the skill identities needed to explain Dictate Plan, additional calls, or unattributed activity independently. The location and privacy model for detailed public evidence links remain undefined.

## Development Versus Stale Baselines

The current development working tree is the authoritative methodology baseline for this research because it contains the intended 30-turn governed scenarios, revised scorecards, canonical skill-source manifest, and current runtime controls. Several other repository views lag it:

- the selected known-good executable was built from an earlier committed revision;
- committed `HEAD` and index-only snapshots can describe intermediate states rather than the current working files;
- earlier research and retained analyses describe superseded 12-turn scenarios; and
- some retained historical status documents preserve obsolete source paths or iteration authority.

These mismatches are freeze and publication blockers. The correct resolution is to finish and verify the intended development sources, commit the coherent baseline, build a new known-good executable from that revision, and confirm that its embedded units and hashes match the frozen sources. Rolling the current scenarios back to the stale executable would make the executable consistent by discarding the intended methodology and is therefore unsupported.

## Conditional Alternatives

### Before scenario freeze

Publish only structural method claims: three planned scenarios, 30 turns per scenario, continued-session replay, private turn-skill scorecards, opaque external attribution, and adapter-specific controls. Label the current 44/44/43 and 27-turn figures as inspected development counts pending freeze.

### After scenario freeze

Recompute and record turn counts, unique expected pairs, expectation-bearing turns, per-skill totals, prompt-level Dictate Plan occurrences, and canonical hashes from the frozen sources. Then build and inspect the known-good executable before starting campaign runs.

### Before accepted campaign results

Keep website charts explicitly illustrative. Avoid rankings, comparative percentages, or claims about observed harness performance.

### For environment wording

Describe concrete controls and residuals per harness. Use “fresh evaluator-owned workspace” and “controlled runtime configuration” rather than a blanket isolation claim.

## Rejected Or Lower-Fit Interpretations

- **The current scenarios contain 12 turns:** rejected by the development built-ins; this describes a stale baseline.
- **Thirty refers only to campaign runs:** rejected; current scenarios also contain 30 turns each.
- **Each scenario expects 15 to 20 calls:** rejected by the current 44/44/43 scorecards.
- **Every turn is scored:** rejected; 27 of 30 turns currently carry expectations.
- **The four illustrative website checkpoints are the current scorecard:** rejected; they are fallback presentation fixtures.
- **The website's `total_turns: 30` is stale:** rejected; it is aligned with current development scenarios.
- **Dictate Plan is excluded because it is manual:** rejected by every current answer sheet.
- **Dictate Plan demonstrates spontaneous startup discovery:** rejected because Turn 1 names it directly.
- **Dictate Plan is explicit-only on every harness:** unsupported beyond its Codex metadata and the explicit governed prompt.
- **Opaque instrumentation is invisible or behavior-free:** rejected by the visible injected instruction and command surface.
- **All harnesses have equivalent isolation:** rejected by their concrete runtime and authentication differences.
- **Pi offline mode proves network confinement:** rejected; the host Bash/file surface and network boundary require separate proof.
- **Kilo's supplied path proves exclusive runtime skill discovery:** not yet proven by accepted runtime evidence.
- **A clean process exit proves acceptance:** rejected by protocol, artifact, session, cleanup, and confinement gates.
- **Missing calls are tooling failures:** rejected when the run is otherwise accepted; they are measured outcomes.
- **A passed refinement campaign proves universal skill correctness:** rejected by the campaigns' own evidence boundaries.
- **The current website charts are observed campaign results:** rejected because the accepted collection is empty.
- **A compact artifact that passes the importer is accepted campaign evidence:** rejected because the importer lacks provenance and acceptance checks.
- **The stale known-good executable should override development intent:** rejected; the divergence blocks freeze and requires rebuilding the baseline.

## Canonical Repository Locations

| Evidence class                                 | Current canonical location                                                                    |
| ---------------------------------------------- | --------------------------------------------------------------------------------------------- |
| Embedded runtime scenarios and answer sheets   | `evaluations/skill-calling/built-ins/*.json`                                                  |
| Human-readable conversations and scorecards    | `evaluations/skill-calling/scenarios/<scenario>/`                                             |
| Instrumentation contract                       | `evaluations/skill-calling/instrumentation-contract.md`                                       |
| Evaluation orchestration and derivation        | `cli/internal/evaluation/evaluation.go`                                                       |
| Prompt replay and session continuity           | `cli/internal/replay/`                                                                        |
| Skill instrumentation/materialization          | `cli/internal/installer/installer.go`                                                         |
| Token mapping and invocation events            | `cli/internal/runstate/runstate.go`                                                           |
| Harness runtime controls                       | `cli/internal/evaluation/runtime.go`                                                          |
| Component source ownership                     | `cli/internal/payload/assets/manifest.json`                                                   |
| Dictate Plan canonical source                  | `supporting-skills/dictate-plan/`                                                             |
| Other seven scenario skill sources             | `evaluations/scenario-skill-refinement/<skill>/skill/`                                        |
| Current intake/generation/refinement workflows | `skills/skill-intake/`, `skills/skill-generation/`, `skills/skill-evaluation-and-refinement/` |
| Dictate Plan retained refinement evidence      | `evaluations/skill-system-production-refinement/targets/dictate-plan/`                        |
| Other scenario-skill refinement evidence       | `evaluations/scenario-skill-refinement/<skill>/`                                              |
| Campaign matrix and progress                   | `plans/skill-calling-evaluation-campaign/evaluation-progress.md`                              |
| Campaign execution and acceptance procedure    | `plans/skill-calling-evaluation-campaign/evaluation-orchestration-prompt.md`                  |
| Accepted compact website collection            | `src/data/publishedWebsiteArtifacts.json`                                                     |
| Website adaptation and illustrative fallback   | `src/data/evaluationData.ts`                                                                  |
| Website import command                         | `scripts/update-website-results.mjs`                                                          |

## Unresolved Blockers

1. **Scenario freeze is incomplete.** The current 30-turn sources are authoritative for intent but still changing; final totals and hashes need one post-freeze verification.
2. **Known-good parity is absent.** The selected known-good executable and earlier committed snapshots do not yet embed the current governed method.
3. **The campaign baseline is not immutable.** A coherent revision, rebuilt known-good executable, exact version record, and embedded-source verification are still required.
4. **No accepted campaign set exists.** The planned matrix is 0/30 complete in the inspected progress record.
5. **Lane confinement remains evidence-gated.** Each harness requires its bounded runtime probe, especially for product-managed surfaces, Kilo skill exclusivity, and Pi's host/network boundary.
6. **Exact configuration provenance remains incomplete.** Every run needs the effective harness version, executable, model identifier/version or alias, and reasoning setting.
7. **The accepted website collection is empty.** Current charts fall back to illustrative data.
8. **Illustrative website data is sparse and partly stale.** Four points cannot represent the 27 expectation-bearing turns, and the gardening label needs reconciliation; the 30-turn total is current.
9. **The importer is not acceptance-aware.** It does not bind compact artifacts to detailed accepted evidence or a frozen campaign identity.
10. **Detailed public evidence links are undefined.** Publication needs a stable location, privacy review, and linkage model.
11. **Generation lineage is incomplete.** Retained evidence supports bounded skill refinement but not a full current intake-to-generation chain for every target.
12. **Historical status drift remains.** Dictate Plan's retained status uses an obsolete target path and legacy 40-iteration authority.
13. **Durability is not established.** Current methodology sources, canonical skill moves, and this research may be unpublished, untracked, or otherwise absent from the eventual public revision until deliberately committed and pushed.
14. **Transcript privacy remains a publication gate.** Sanitization of known machine identifiers does not establish that arbitrary conversation content is safe to publish.

## Evidence Boundary

This synthesis establishes the inspected local development methodology and its divergence from stale baselines. It does not establish a frozen executable, completed campaign outcomes, accepted cross-harness comparisons, current public-repository contents, or universal harness confinement. The detailed inspection record remains at `research/testing-methodology/assignments/01-local-system-methodology.md`.
