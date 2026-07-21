# Local Skill-Calling Evaluation Methodology

## Assignment

**Goal:** Validate the proposed public account of Skill Issue's testing methodology by tracing the current local production implementation, governed evaluation inputs, campaign controls, skill-building workflow, retained evidence, and website publication path end to end. Independently establish how Dictate Plan participates in the governed scenarios.

**Scope:** Local repository sources only. Production Go and website source are behavioral authority. Governed built-in JSON, human-readable scenario views, campaign records, canonical skills, and retained evaluation material are supporting repository contracts. Tests were not used as behavioral authority.

**Exclusions:** No production, plan, scenario, skill, result, or website files were edited. No campaign run was started. No internet source or current GitHub page was inspected. This assignment does not decide page layout or write audience-facing copy.

## Sources

### Production implementation

- `bundle.go:5-6` and `cli/internal/payload/payload.go:96-169` — embedded canonical skill and built-in evaluation loading.
- `cli/internal/evaluation/evaluation.go:181-459, 938-1060` — run preparation, runtime handoff, turn boundaries, result derivation, website artifact derivation, and completion.
- `cli/internal/replay/replay.go:111-157` — one-session ordered prompt replay.
- `cli/internal/replay/process.go:137-192, 337-349, 485-505, 797-839` and `cli/internal/replay/pi.go:48-101, 156-236` — concrete harness process/session continuation.
- `cli/internal/installer/installer.go:468-558` — harness metadata transformation and injected signal instruction.
- `cli/internal/runstate/runstate.go:48-220` — random tokens, private token-to-skill mapping, active-turn attribution, and event persistence.
- `cli/internal/evaluation/runtime.go:34-455, 506-560` and `cli/internal/harness/harness.go:45-103` — per-harness defaults, generated configuration, environment isolation, permissions, and skill roots.
- `cli/internal/lifecycle/lifecycle.go:94-185` and `cli/internal/command/command.go:305-340` — pre-run lifecycle, completed-run result, and silent `signal` route.

### Governed inputs and contracts

- `evaluations/skill-calling/built-ins/gardening-web-application.json`
- `evaluations/skill-calling/built-ins/community-archive-desktop-application.json`
- `evaluations/skill-calling/built-ins/neighborhood-emergency-preparedness-program.json`
- `evaluations/skill-calling/scenarios/*/{conversation.md,expected-calls.md,instructions.md}`
- `evaluations/skill-calling/instrumentation-contract.md:3-56`
- `README.md:33-61` and `cli/README.md:61-221`
- `plans/skill-calling-evaluation-campaign/evaluation-progress.md:1-154`
- `plans/skill-calling-evaluation-campaign/evaluation-orchestration-prompt.md:1-148`

### Skills and retained refinement evidence

- `supporting-skills/dictate-plan/SKILL.md` and `supporting-skills/dictate-plan/agents/openai.yaml`
- The seven other scenario skills under `supporting-skills/`.
- `skills/skill-intake/SKILL.md`, `skills/skill-generation/SKILL.md`, and `skills/skill-evaluation-and-refinement/SKILL.md`.
- `evaluations/skill-system-production-refinement/targets/dictate-plan/`
- `evaluations/scenario-skill-refinement/<skill>/` for the seven non-Dictate-Plan scenario skills.

### Website publication path

- `src/data/evaluationData.ts:1-239`
- `src/data/publishedWebsiteArtifacts.json`
- `scripts/update-website-results.mjs:1-42`
- `src/data/siteData.ts:80-128`

### Validation performed

- Parsed all three built-ins with `jq` to count turns, expected calls, and calls per skill.
- Compared the current index with `HEAD` using `git diff --cached`; the 12-turn scenarios, their scorecards, Dictate Plan, and much retained refinement material are staged changes at local `HEAD` `e1f7e1062afe3d507698975c6517b90b84467a12`, whose `origin/main` points to the same commit without the staged content.
- Compared the former committed 30-turn gardening unit with the staged 12-turn unit. The former scorecard had four expected calls; the current staged unit has 20.
- Confirmed the retired Dictate Plan evaluation copy is absent and that the current canonical `supporting-skills/dictate-plan/SKILL.md` hash, `44bc9abeeebc2200ec5c3b6ea7a65cab70551e72c1b9f7c06c281c2b2ec0fe2b`, matches the content hash recorded by its retained campaign.

## Findings

### Finding 1 — The current campaign has three 12-turn scenarios, not three 30-turn conversations

The current governed runtime units are three scripted scenarios with 12 fixed user turns each. Their scorecards contain 20, 20, and 19 unique expected `{turn_id, skill}` pairs. The “approximately 15 to 20 expected calls each” statement is supported by the current units, but “30 conversational turns each” is rejected. Thirty is still meaningful elsewhere: the campaign matrix plans ten harness/model configurations times three scenarios, producing 30 separate evaluation runs.

| Built-in evaluation | Scenario product | Turns | Expected calls |
| --- | --- | ---: | ---: |
| `gardening-web-application` | SproutCheck | 12 | 20 |
| `community-archive-desktop-application` | BoxIndex | 12 | 20 |
| `neighborhood-emergency-preparedness-program` | ReadyCard | 12 | 19 |

The current README states the 12-turn contract (`README.md:35-39`), and the CLI README inventories the exact counts (`cli/README.md:61-80`). The built-in JSON is the executable input; for example, gardening defines 12 prompts at `evaluations/skill-calling/built-ins/gardening-web-application.json:7-56` and 20 expectations at lines 58-142. The human-readable `conversation.md` and `expected-calls.md` files mirror those embedded inputs but do not drive replay.

**Evidence:** Direct JSON counts; `cli/internal/payload/payload.go:96-105`; `cli/README.md:61-80`; `plans/skill-calling-evaluation-campaign/evaluation-progress.md:3-20`.

**Implication:** Public copy should say “three 12-turn scenarios, yielding 30 planned runs across ten configurations.” It should not reuse “30 turns” from the previous committed scenario generation or from stale website fixtures.

### Finding 2 — Scorecards map expected skills to turns, and the scored unit is a unique turn-skill pair

Each built-in embeds an `answer_sheet` beside its scenario. Every current turn has at least one expected pair, and a turn may have several. Gardening Turn 8, for example, expects `systematic-debugging`, `code-testing-discipline`, and `code-implementation-discipline`; Turn 1 expects both Dictate Plan and document update discipline.

The evaluator validates that each expected turn exists and each skill is present, then compares recorded attributed events by the composite key `turn_id + skill` (`cli/internal/evaluation/evaluation.go:938-1007`). Missing is the absence of that key. Additional is an attributed event whose composite key is absent from the scorecard. Signals outside an active turn are `unattributed`. The detailed `observed` and `additional` arrays can retain repeated raw events; the missing calculation and compact website counts use sets, so same-turn duplicates do not increase `called` or `missed` (`cli/internal/evaluation/evaluation.go:1016-1059`).

For these current answer sheets, every expected pair is unique and therefore has a Boolean called/missed outcome. Repeated applicability of the same skill on later turns is intentionally scored as a separate decision. There is no pass/fail label in `result.json`.

**Evidence:** All three `expected-calls.md` tables; `evaluations/skill-calling/built-ins/gardening-web-application.json:58-142`; `cli/internal/evaluation/evaluation.go:938-1059`; `evaluations/skill-calling/instrumentation-contract.md:46-50`.

**Implication:** The defensible public unit is “expected skill on an expected turn,” not conversational turn alone and not raw invocation count. Charts can aggregate called/missed pairs by turn, while detailed interpretation must use `result.json`.

### Finding 3 — The CLI conducts controlled, stateful conversations with fixed user prompts

The replay runner starts one primary-agent session, iterates the ordered scenario, marks a before-turn boundary, sends the exact prompt, waits for terminal completion, marks an after-turn boundary, and then proceeds (`cli/internal/replay/replay.go:111-147`). The same harness session identifier is captured and reused on later prompts through each concrete adapter. Agent responses and structured harness events are captured as evidence but never generate, rewrite, or skip later governed prompts.

“Controlled back-and-forth” is accurate if it means fixed user messages interleaved with the model's unconstrained completed responses in one continued session. It would be misleading if it implied that both sides are scripted or that later user turns react to response content. The runner controls boundaries and order, while the selected model and harness control the responses, tool use, and workspace effects.

The CLI allows `--turns` truncation for smoke and custom work, but the campaign contract explicitly forbids it for full runs. Inputs are parsed and cached before confirmation so execution uses what the operator reviewed (`cli/internal/evaluation/evaluation.go:206-209`).

**Evidence:** `cli/internal/replay/replay.go:111-157`; `cli/internal/replay/process.go:137-192, 337-349, 485-505, 813-839`; `cli/internal/replay/pi.go:156-236`; `README.md:35`; `plans/skill-calling-evaluation-campaign/evaluation-orchestration-prompt.md:61-69`.

**Implication:** The page should distinguish governed user turns, generated assistant responses, stable session continuity, and deterministic versus harness-dependent behavior.

### Finding 4 — Instrumentation is one injected `signal` instruction with an opaque external mapping, but it is unobtrusive rather than invisible

Every selected skill is copied and receives exactly one instruction immediately after frontmatter:

```text
Run "<absolute-cli-path>" signal "<opaque-token>" "<absolute-output-state-path>", then continue normally.
```

`runstate.NewToken` generates 32 random bytes encoded as 64 hexadecimal characters. The evaluated skill contains the token but not the skill-name mapping, answer sheet, expected turn, or scoring rule. The private token map lives under the output-owned state root outside the evaluated workspace. `Store.Mark` resolves the token to the skill, reads the evaluator-owned active turn, and appends a structured event. The `signal` CLI route emits no user output on success.

Codex is a concrete exception in recording mechanics: the nested command need not write the private state. The outer evaluator recognizes the exact attempted signal command in Codex structured output and calls the same private marker path while the turn remains active. Cursor requires the executed command to have recorded its marker. Other process adapters permit the exact signal command.

The owner statement that opaque identifiers are translated outside model context is supported. Calling the command “innocuously named” should be qualified: `signal` is semantically neutral, but the model can see the absolute executable and state path, the command shape, and the injected instruction. The executable or path can contain `skill-issue`, and an inspecting agent can infer evaluation machinery. The repository contract expressly disclaims perfect invisibility or influence-free instrumentation.

**Evidence:** `cli/internal/installer/installer.go:516-558`; `cli/internal/runstate/runstate.go:77-90, 96-187`; `cli/internal/evaluation/evaluation.go:475-597`; `cli/internal/lifecycle/lifecycle.go:176-185`; `evaluations/skill-calling/instrumentation-contract.md:16-36, 56`.

**Implication:** Public disclosure should say the design minimizes semantic clues and keeps scoring meaning outside the model context, while acknowledging that the extra instruction and command may affect behavior.

### Finding 5 — Dictate Plan is explicitly requested at startup and included in the scored expectations

The canonical current skill is `supporting-skills/dictate-plan/`. Its portable name is `dictate-plan`; Codex metadata displays **Dictate Plan** and sets `policy.allow_implicit_invocation: false` (`supporting-skills/dictate-plan/agents/openai.yaml:1-6`). The portable skill description says it handles living A-to-B task sequencing over successive messages (`supporting-skills/dictate-plan/SKILL.md:1-4`).

Each scenario's Turn 1 begins with the ordinary-language phrase “Let's do Dictate Plan” and asks for one living A-to-B plan. No later governed prompt names Dictate Plan. Thus the scenario manually requests it exactly once, at scenario startup, through explicit prose rather than a harness-specific slash command or picker.

The “explicit-only” claim is only fully established for the Codex metadata surface. `agents/openai.yaml` is included only for Codex installations. The portable `SKILL.md` has no cross-harness disable flag, and the installer adds `disable-model-invocation: true` only to `skill-intake`, not Dictate Plan, for harnesses supporting that metadata (`cli/internal/installer/installer.go:468-514`). Other harnesses still receive an explicitly worded Turn 1 prompt, but this repository does not configure Dictate Plan as technically non-implicit on every harness.

Dictate Plan is not excluded from scoring. Every answer sheet contains `{turn-1, dictate-plan}`. It receives the same injected instrumentation as every other supplied skill and appears by name in detailed `expected`, `observed`, `missing`, or `additional` result arrays. `website.json` does not include skill names: the Dictate Plan outcome is merged into Turn 1's aggregate `called`/`missed` count alongside document update discipline.

No accepted campaign result currently proves how Dictate Plan appeared in a completed governed run. Its retained pre-campaign evaluation did pass four natural description trials and three multi-turn behavior cases at the recorded content hash (`evaluations/skill-system-production-refinement/targets/dictate-plan/status.md`, `description/round-1/evidence.md`, and `behavior/cycle-1/audit.md`).

**Evidence:** Turn 1 in each built-in JSON and conversation; each `expected-calls.md`; `supporting-skills/dictate-plan/`; `cli/internal/installer/installer.go:468-514`; `cli/internal/evaluation/evaluation.go:129-163, 981-1059`.

**Implication:** The methodology page needs a prominent startup disclosure: Dictate Plan is deliberately named by the user and scored as an expected Turn 1 call. It is not evidence of spontaneous discovery. Cross-harness “explicit-only configuration” should not be claimed.

### Finding 6 — Isolation is substantial and harness-specific; “isolated, sandboxed, default-like” is not a uniform proven property

The campaign orchestration requires a newly allocated external `chat-<number>/workspace` that is empty and contains no Git metadata, instructions, rules, configuration, results, or prior files. It also requires a separate sibling output root and a confinement smoke proving the harness cannot read sibling containers, output, or the repository. The production CLI itself enforces only that the workspace is an existing directory and the output root is outside it; emptiness, neutrality, and confinement are orchestration responsibilities.

| Harness | Default production target | Main isolation controls | Material retained ambient surface or limitation |
| --- | --- | --- | --- |
| Codex | `gpt-5.6-sol`, medium | `--ignore-user-config`, `--ignore-rules`, plugins/apps disabled, `project_doc_max_bytes=0`, ambient skill deny-list, explicit `workspace-write`, `on-request` plus `auto_review` | Uses normal Codex home for authentication and retained session history; managed/system requirements remain; environment is overlaid rather than fully clean |
| Claude Code | `opus`, medium | Private launch directory and passed-skill tree, project-only setting source, strict MCP, Chrome disabled, auto-memory/CLAUDE.md/background features disabled, allowlisted tools and exact signal command | Caller environment is merged; workspace is an added directory named by appended system prompt; no repository-owned OS sandbox claim is established |
| Cursor | `auto`, model-native reasoning | Private home/config/store/plugin, clean environment, project configs disabled, sandbox enabled, allowlist, isolated skill plugin | Links operator keychain for auth; network uses `user_config_with_defaults`; Cursor exposes no independent reasoning override |
| OpenCode 1.18.4 | `openai/gpt-5.6-sol`, medium | Private XDG config/state/cache/temp, pure JSON run, project/external skills/plugins/instructions/MCP disabled, deny-first tools and skills, exact signal Bash permission | Retains operator-owned authentication data root and compiled internal plugins; process environment is merged despite controlled XDG overrides |
| Pi | `openai-codex/gpt-5.6-sol`, medium | Clean private home/session roots, no session file, approvals/extensions/ambient skills/templates/themes/context disabled, explicit skills only, offline mode, limited tool list | Uses operator `PI_CODING_AGENT_DIR` for authentication/provider state and forwards an explicit credential allowlist; repository source does not establish an OS filesystem sandbox |

These controls support “as default-like and isolated as the qualified adapter permits,” not “identical defaults” or “uniform sandboxing.” Necessary evaluator controls deliberately change defaults. Provider authentication, organization policy, compiled authentication plugins, native session behavior, and harness-specific permissions remain different.

The current campaign does not yet supply accepted full-run confinement evidence. Its progress is 0/30, so source configuration and bounded qualification records establish design and preflight behavior, not final campaign comparability.

**Evidence:** `plans/skill-calling-evaluation-campaign/evaluation-orchestration-prompt.md:73-127`; `cli/internal/evaluation/evaluation.go:181-205, 867-904`; `cli/internal/evaluation/runtime.go:34-455, 506-560`; `cli/internal/harness/harness.go:45-73`; `cli/internal/replay/process.go:485-505, 813-839`; `cli/internal/replay/pi.go:48-77`; `cli/README.md:146-172`.

**Implication:** The page should explain each harness's controls and residual surfaces rather than present one blanket isolation claim. Campaign publication should wait for the required per-route confinement and cleanup evidence.

### Finding 7 — Skill generation and two-loop refinement are real workflows, but generation provenance for the eight campaign skills is not fully established

The public workflow is explicit and dependency ordered: `skill-intake` creates a build-ready A-to-B contract; `skill-generation` turns that contract into the skill, validates structure and criteria, and must hand the result to `skill-evaluation-and-refinement`; evaluation then separately tests description selection and skill-body behavior.

Description evaluation uses two initial plus two confirmation prompts with fresh agents and native load evidence. Body evaluation derives observable criteria, qualifies references, runs varied isolated cases, audits outputs, applies one semantic refinement at the meaning owner when warranted, recreates clean fixtures, and reruns. Five unsuccessful description rounds or five unsuccessful body cycles trigger a user-controlled stop rather than endless automatic refinement (`skills/skill-evaluation-and-refinement/SKILL.md:15-67`). Explicit-only skills are allowed to mark description evaluation not applicable, although the retained Dictate Plan campaign instead records four natural selection trials.

Tracked local evidence exists for all eight scenario skills. Dictate Plan's earlier production-refinement campaign passed. Seven other supporting skills have completed `evaluations/scenario-skill-refinement/<skill>/` campaign records with status and conclusions; some passed unchanged, while `document-update-discipline` retained and passed an evidence-supported body refinement. These records establish bounded cases on qualified local Codex surfaces, not universal correctness across campaign harnesses.

What is not established is that each of these eight particular skills originally came through the current `skill-intake` plus `skill-generation` workflow. The workflow exists and the skills were evaluated/refined, but source lineage from a retained intake contract through generation is incomplete for the scenario set.

**Evidence:** `skills/skill-intake/SKILL.md`; `skills/skill-generation/SKILL.md:8-60`; `skills/skill-evaluation-and-refinement/SKILL.md:8-74`; each retained target `status.md` and `conclusion.md`; Dictate Plan retained campaign files.

**Implication:** A safe public claim is that the campaign skills underwent the project's evaluation/refinement process before campaign use. Claiming that every one was generated end to end by the current generation workflow requires additional provenance.

### Finding 8 — Run acceptance is tooling-based and descriptive, while campaign acceptance remains unfinished

The CLI returns a completed result after the ordered replay, protocol validation, result and website writes, state update, and cleanup. A run with zero observed expected skills is still valid evidence. Launch, authentication, permissions, protocol, session, marker, artifact, process, or cleanup failures are tooling failures and must be repaired and rerun. The result itself has no pass/fail field.

The interactive CLI prints the selected evaluation, effective turn count, harness, model, reasoning, workspace, output, executable, and custom paths before starting and defaults cancellation to no. Full campaign acceptance is stricter: each of 30 matrix runs must be tooling-complete, retained, cleaned, associated with exact effective settings, and then frozen as the website's accepted set.

Current progress is 0/30 complete. `CLA-COD-01` has one failed attempt because the operator interrupted it during Turn 7; cleanup completed, and the whole 12-turn scenario must be rerun in a fresh workspace. Exact model identifiers remain unresolved for most matrix cells. The only current campaign note with exact values is Claude Code through the Codex proxy: `gpt-5.6-sol`, Claude Code `2.1.205`, CLIProxyAPI `7.2.91`.

**Evidence:** `cli/internal/evaluation/evaluation.go:351-459`; `README.md:55-61`; `plans/skill-calling-evaluation-campaign/evaluation-progress.md:22-154`; `plans/skill-calling-evaluation-campaign/campaign-orchestration-prompt.md:3-75`.

**Implication:** The methodology can be published before results if clearly labeled as the designed method. Observed-result claims, comparative conclusions, and accepted run links remain blocked.

### Finding 9 — Detailed and compact artifacts exist, but the public website currently contains no accepted evaluation artifact

Every tooling-complete run writes:

- `result.json`: authoritative run identity plus expected, observed, missing, additional, and unattributed named calls;
- `website.json`: compact `{turn, turn_id, called, missed}` points plus run/scenario/harness/model and `total_turns`;
- optional `events.jsonl` with `--events`;
- optional sanitized `transcript.json` with `--transcript`.

The campaign orchestration intends to retain logs and transcripts for every official run even though the production flags default off. Transcript sanitization replaces known paths and machine identifiers but cannot infer or remove arbitrary personal/confidential conversation content, so review before public sharing is mandatory.

The website's checked-in public artifact owner is `src/data/publishedWebsiteArtifacts.json`, currently `[]`. `src/data/evaluationData.ts` therefore falls back to generated illustrative records. Those fixtures still model 30 turns at Turns 1, 11, 25, and 30 and use real-looking harness/model/scenario identities. `src/data/siteData.ts` also says each scenario contains 30 turns. These are stale against the staged 12-turn scenario source. The importer accepts any loosely shaped `website.json`, sorts it, and writes it directly; it does not join `result.json`, verify campaign acceptance, preserve reasoning, or validate scenario-specific denominators.

Dictate Plan cannot be identified from `website.json`; only the detailed result preserves skill names. Public inspection therefore needs both accepted detailed evidence and the compact artifact, or explicit links from the aggregate chart to the detailed run.

**Evidence:** `cli/internal/evaluation/evaluation.go:129-163, 415-459, 1016-1059`; `cli/README.md:174-221`; `src/data/publishedWebsiteArtifacts.json`; `src/data/evaluationData.ts:36-208, 230-239`; `scripts/update-website-results.mjs:1-42`; `src/data/siteData.ts:80-83`.

**Implication:** “Retained artifacts will support public inspection” is an intended and technically plausible direction, not a current achieved state. Publication needs an accepted-set owner, paired detailed/compact validation, provenance, and corrected illustrative/stale copy.

### Finding 10 — Canonical repository locations are clear, but current public-link readiness is blocked by staged source

The semantic owners are:

| Material | Canonical repository location |
| --- | --- |
| Runtime built-in inputs | `evaluations/skill-calling/built-ins/*.json` embedded through `bundle.go` |
| Human-readable scenarios | `evaluations/skill-calling/scenarios/<scenario>/conversation.md` |
| Human-readable scorecards | `evaluations/skill-calling/scenarios/<scenario>/expected-calls.md` |
| Instrumentation contract | `evaluations/skill-calling/instrumentation-contract.md` |
| Production evaluator | `cli/internal/evaluation/`, `cli/internal/replay/`, `cli/internal/runstate/`, `cli/internal/installer/`, `cli/internal/harness/` |
| Scenario skills | `supporting-skills/<skill>/` |
| Generation/refinement workflows | `skills/skill-intake/`, `skills/skill-generation/`, `skills/skill-evaluation-and-refinement/` |
| Dictate Plan retained evaluation | `evaluations/skill-system-production-refinement/targets/dictate-plan/` |
| Seven other retained evaluations | `evaluations/scenario-skill-refinement/<skill>/` |
| Campaign state | `plans/skill-calling-evaluation-campaign/evaluation-progress.md` |
| Accepted website artifact collection | `src/data/publishedWebsiteArtifacts.json` |

The current 12-turn built-ins, matching human-readable documents, Dictate Plan source/metadata, many refinement records, and progress corrections are staged but absent from `origin/main` at the inspected commit. The public repository at that commit therefore still represents the former 30-turn/four-expectation scenario generation. Canonical GitHub evidence links for the new method will be wrong or missing until the staged sources are reviewed, privacy-checked, committed, and pushed.

**Evidence:** `git status --short`, `git diff --cached`, `git log -1`, `git remote -v`, `.repository-privacy.md`, `bundle.go:5-6`.

**Implication:** Link creation should follow source publication, not precede it. The methodology page should link to committed immutable paths or commits that contain the same units used by the accepted campaign CLI baseline.

### Finding 11 — Claim-by-claim disposition

| Project-owner statement | Disposition | Supported wording |
| --- | --- | --- |
| Three scenarios with 30 conversational turns and about 15-20 expected calls each | Partly rejected | Three current scenarios have 12 turns and 20/20/19 expected calls; 30 is the number of planned matrix runs |
| Scorecards map turns to expected skills | Verified | Embedded private answer sheets and mirrored Markdown map unique expected turn-skill pairs |
| CLI conducts controlled back-and-forth conversations | Verified with qualification | Fixed user prompts are sent in order in one resumed primary-agent session; assistant responses vary and never alter later prompts |
| Evaluated skills receive an innocuously named recording call with opaque identifiers translated outside model context | Verified with qualification | One `signal` instruction is injected; token meaning is private and external, but the visible instruction/path can reveal instrumentation |
| Comparable runs use isolated, sandboxed, default-like workspaces without unrelated skills or ambient instructions | Direction supported, blanket claim rejected | Orchestration requires empty neutral workspaces and adapters suppress many ambient surfaces; enforcement and sandbox strength differ, and accepted confinement evidence is incomplete |
| Skill generation plus evaluation/refinement hardens descriptions and bodies before campaign use | Partly verified | The workflow exists and all eight current scenario skills have bounded refinement evidence; complete generation lineage for each skill is not retained |
| Retained artifacts will support public inspection | Intended, currently blocked | Writers and retained local material exist, but the official campaign is 0/30, accepted website artifacts are empty, and key current sources are not yet on `origin/main` |
| Dictate Plan is explicit-only, invoked once at startup, and excluded from scored spontaneous calls | Mixed | Codex metadata is explicit-only; every scenario names it once at Turn 1; it is included in the scored expectations and must be disclosed as a manual exception |

**Evidence:** Findings 1-10.

**Implication:** The public methodology should lead with the verified 12-turn design and disclose Dictate Plan, instrumentation, harness-specific controls, campaign incompleteness, and bounded evidence without smoothing the mixed dispositions.

### Finding 12 — Unresolved publication blockers and unsupported claims

1. **No accepted campaign set:** 0/30 runs are complete; one first-lane attempt failed by operator interruption.
2. **Current governed sources are unpublished:** the 12-turn units, revised scorecards, Dictate Plan, and much refinement evidence are staged and absent from current `origin/main`.
3. **Stale public data/copy:** `evaluationData.ts`, `siteData.ts`, and the `cli/README.md` example still use 30-turn assumptions; current scenario display label `GardenFlow planning` no longer matches SproutCheck.
4. **No acceptance-aware importer:** `scripts/update-website-results.mjs` can publish loosely shaped compact artifacts without detailed-result, reasoning, campaign, environment, or acceptance validation.
5. **No accepted detailed evidence links:** `publishedWebsiteArtifacts.json` is empty, and compact artifacts cannot disclose Dictate Plan or additional/unattributed calls.
6. **Incomplete exact configuration inventory:** most campaign model aliases, versions, and effective reasoning values remain unresolved.
7. **Workspace neutrality is not a CLI invariant:** empty-workspace allocation and sibling/repository confinement depend on orchestration and still require per-route proof.
8. **Uniform sandboxing is unsupported:** controls and residual authentication/configuration surfaces differ materially by harness.
9. **Instrumentation neutrality is bounded:** the injected instruction may influence behavior and can be inferred by environment inspection.
10. **Event schema mismatch:** `runstate.Event` emits `reasoning` and `evaluation_id`, but `evaluations/skill-calling/event.schema.json` forbids undeclared properties, so emitted raw events do not conform to the committed schema.
11. **Transcript publication needs manual privacy review:** path sanitization is not content redaction.
12. **Skill-generation provenance is incomplete:** retained evidence supports refinement, not the full creation lineage of every campaign skill.

**Evidence:** The cited source owners above; `evaluations/skill-calling/event.schema.json`; `cli/internal/runstate/runstate.go:48-60`; `.repository-privacy.md`.

**Implication:** These are true blockers or required qualifications, not deferred validation tasks. Until resolved, public claims should remain about the designed method and bounded pre-campaign skill evidence rather than completed comparative results.

## Notes

- The former committed built-ins contained 30 turns but only four expected calls at selected checkpoints. The current staged units changed both the story and the measurement design: 12 executable application-building turns with an expectation on every turn. Combining “30 turns” from the old generation with “15-20 calls” from the new generation creates a method that never existed as one governed unit.
- The current `cli/README.md` compact artifact example still shows `"total_turns": 30` even though the surrounding text now states 12. The production writer correctly derives `total_turns` from the actual scenario.
- Some human scenario instructions say tooling-complete evidence retains `events.jsonl` and `transcript.json`; production keeps them optional. The official campaign prompt makes transcripts/logs required for campaign retention, so the campaign command and acceptance record must explicitly enable and verify them.
- `prompt-writing` and `skill-authoring-discipline` Codex metadata omit an explicit `allow_implicit_invocation` field, whereas most other proactively expected supporting skills set it to `true`. This does not invalidate the current scorecard, but the page should avoid implying identical metadata across all skills.
- The research found no production schema binding three scenario runs into one accepted configuration suite or all suites into an accepted campaign. The campaign progress/freeze process is currently the intended acceptance owner.
