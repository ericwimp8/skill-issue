# Replay Scoring And Turn Attribution

## Assignment

**Goal:** Trace production replay capture, skill-invocation event normalization, evaluation comparison, and turn attribution end to end so the website chart can consume the implemented evidence without changing its meaning.

**Scope:** The CLI command route, disposable skill instrumentation, replay boundary callbacks, private run state, harness capture normalization, expected-call validation and correlation, detailed result derivation, compact `website.json` projection, governing instrumentation/schema documents, built-in expected-call methodology, and chart-relevant granularity.

**Exclusions:** Frontend implementation, campaign execution, semantic review of whether the governed answer sheets chose the right expected calls, unrelated installer/runtime behavior, and any edits to production sources or existing plans.

## Sources

- `cli/internal/command/command.go:38-88`, especially `App.Run`, `runMarker`, and `runEvaluation`.
- `cli/internal/lifecycle/lifecycle.go:146-167`, especially `runEvaluation` and `mark`.
- `cli/internal/installer/installer.go:77-117,267-311`, especially `PrepareEvaluation`, `instrument`, and `inject`.
- `cli/internal/replay/replay.go:27-58,61-88,97-159`, especially `Scenario.Validate`, `Capture`, `Runner.Run`, and `Runner.notify`.
- `cli/internal/replay/process.go:133-218,385-492`, especially `processSession.SendPrompt`, `processSession.Wait`, `parseEvents`, and the harness command specifications.
- `cli/internal/replay/pi.go:145-214`, especially `piSession.SendPrompt` and `piSession.Wait`.
- `cli/internal/runstate/runstate.go:15-48,74-95,125-213`, especially `Run`, `Event`, `Store.SetActiveTurn`, `Store.Mark`, and `Store.Events`.
- `cli/internal/evaluation/evaluation.go:20-113,131-365,449-525,590-737`, especially `SkillCall`, `AnswerSheet`, `Result`, `WebsiteResult`, `Service.Run`, `recordCodexSignals`, `validateAnswerSheetForSkills`, `deriveResult`, and `deriveWebsiteResult`.
- `evaluations/skill-calling/event.schema.json:1-52` and `evaluations/skill-calling/instrumentation-contract.md:3-56`.
- `evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:3-30`, with the archive and emergency equivalents inspected for consistency.
- `evaluations/skill-calling/built-ins/gardening-web-application.json`, with all three built-in answer sheets inspected; each has the same four expected `{turn_id, skill}` pairs.
- `cli/README.md:120-173` and `plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:19-71,78-94,104-128` for the public contract and planned methodology.
- `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:85-95` for the downstream website migration requirement.
- `src/data/siteData.ts:1-80` and `src/components/EvaluationChart.tsx:17-106` to identify the current mock chart shape that will consume the artifact later.
- Validation only, after establishing behavior from production source: `cli/internal/evaluation/evaluation_test.go:79-146`.

## Findings

### The production path records instrumented skill-entry signals, not transcript interpretation

Each selected evaluation skill receives one 32-byte opaque token. The installer copies the complete skill and inserts `Run "<cli>" signal "<token>" "<state-root>", then continue normally.` immediately after the closing `SKILL.md` frontmatter. The evaluated agent receives the token but not its skill mapping, expected turn, or scoring rule. A successful signal resolves the token in private state and appends a `skill_invocation` event. The evaluator never infers a call by reading response prose or checking whether the rest of the skill body was followed.

**Evidence:** `skillTokens` creates one token for every supplied skill and maps token to skill name (`cli/internal/evaluation/evaluation.go:630-639`). `instrument` finds that token by skill name and `inject` inserts the command without changing other files (`cli/internal/installer/installer.go:267-311`). `command.App` routes `signal` silently through lifecycle marking (`cli/internal/command/command.go:52-66`), and lifecycle resolves the caller-supplied absolute state root before calling `evaluation.Service.Mark` (`cli/internal/lifecycle/lifecycle.go:156-166`). `runstate.Store.Mark` resolves token to run and skill, creates the event, and appends one JSON line (`cli/internal/runstate/runstate.go:147-191`). The public contract explicitly says the token reveals no expected call, turn, or scoring rule (`cli/README.md:120-126`).

**Implication:** A chart label such as “called” means that the instrumented skill entrypoint's signal was recorded for that skill and turn. It does not prove skill-body compliance, output quality, task success, or semantic usefulness.

### Replay owns one active turn around the complete terminal harness response

The scenario's ordered `Turns` slice is the replay authority. For each item, `Runner.Run` emits a before boundary, sends the prompt verbatim, waits until the adapter reports terminal completion, emits an after boundary with that turn's capture, and only then advances. `evaluation.Service.Run` sets private `ActiveTurn` at the before boundary. At the after boundary it stores the harness session, performs Codex fallback normalization while the turn is still active, and finally clears `ActiveTurn`.

**Evidence:** Scenario validation requires schema version 1, a nonempty scenario ID, at least one turn, and unique nonempty turn IDs (`cli/internal/replay/replay.go:38-58`). The sequential boundary/send/wait/after loop is `replay.Runner.Run` (`cli/internal/replay/replay.go:112-149`). The concrete boundary callback sets `ActiveTurn` before prompt delivery and clears it after capture handling (`cli/internal/evaluation/evaluation.go:262-279`). `Store.SetActiveTurn` persists that field under the run lock (`cli/internal/runstate/runstate.go:125-133`). Process adapters reset per-turn stdout/stderr, wait for the native command to finish, parse structured events, and return one `Capture` (`cli/internal/replay/process.go:171-218`); Pi similarly waits for prompt acceptance plus `agent_settled` before returning (`cli/internal/replay/pi.go:164-214`).

**Implication:** Attribution covers activity occurring from just before prompt submission through the adapter's terminal response boundary. It is CLI-owned rather than inferred from timestamps or agent text. A late signal after the active turn is cleared is intentionally unattributed instead of assigned to the nearest turn.

### Attributed and unattributed events are decided at append time

`Store.Mark` copies the current `ActiveTurn` into `Event.TurnID` and sets `Attributed` to whether that string is nonempty. `deriveResult` routes unattributed events into `Unattributed` and excludes them from both `Observed` and website scoring. There is no later timestamp-based reassignment or guessed turn recovery.

**Evidence:** `runstate.Event` carries `TurnID` plus a separate `Attributed` boolean (`cli/internal/runstate/runstate.go:36-48`). `Store.Mark` assigns both directly from the loaded run (`cli/internal/runstate/runstate.go:165-176`). `deriveResult` checks `event.Attributed` before adding an event to observed matching (`cli/internal/evaluation/evaluation.go:651-659`). The instrumentation contract explicitly retains out-of-turn signals as unattributed tooling evidence (`evaluations/skill-calling/instrumentation-contract.md:32-36`).

**Implication:** `website.json` cannot surface unattributed calls, and an unattributed event cannot satisfy an expected call even if its timestamp appears near that turn. Investigating such evidence requires `result.json` and, when deliberately enabled, raw events or transcript diagnostics.

### Codex has a structured command-event fallback with per-skill-per-turn normalization

Non-Codex harnesses rely on the injected `signal` command successfully appending private state. Codex additionally scans that turn's structured native capture for `item.type == "command_execution"` and any opaque token substring. If no direct event for that skill already exists on the turn, it calls the same `Store.Mark` while `ActiveTurn` is still set. Within one captured turn it records at most one fallback event per token. Thus a denied Codex signal attempt can still count as invocation evidence, while unrelated denied commands remain tooling errors under the documented runtime contract.

**Evidence:** `recordCodexSignals` first indexes already stored skills for the current turn, then scans capture events and suppresses a token when it was already recorded or its skill already exists (`cli/internal/evaluation/evaluation.go:331-365`). It is invoked only for Codex in the after boundary before active-turn clearing (`cli/internal/evaluation/evaluation.go:273-278`). Process capture preserves each parsed native JSON object (`cli/internal/replay/process.go:186-217,385-417`). The Codex contract identifies a denied opaque signal command as attributable from its structured command event (`evaluations/skill-calling/instrumentation-contract.md:42-46`; `cli/README.md:128-134`).

**Implication:** The common chart meaning is an observed instrumentation activation, but capture mechanics differ: ordinary successful signals are execution evidence, while Codex fallback can be attempted-command evidence. Repeated denied Codex attempts for one skill in one turn collapse to one fallback event; repeated successfully appended signals may remain as multiple detailed events.

### Expected calls are exact turn-and-skill pairs, governed as required first activations

The answer sheet contains an ordered list of `{turn_id, skill}` pairs. Validation confirms schema version 1, exact scenario identity, at least one expected entry, an existing scenario turn, and an installed skill name. It does not judge whether the expectation is semantically correct and does not reject duplicate pairs. Comparison correlates events and expectations only by the composite string `turn_id + NUL + skill`; tokens and timestamps do not appear in the comparison key.

For all three governed 30-turn evaluations, the methodology treats the answer sheet as required first activations: `document-update-discipline` on turn 1, `prompt-writing` on 11, `skill-authoring-discipline` on 25, and `system-change-ownership` on 30. Later turns where a skill remains applicable are deliberately not additional expected opportunities. A later reload is “additional,” and failure to reload is not “missed” after the governed first activation.

**Evidence:** `AnswerSheet` and `SkillCall` define the pair shape (`cli/internal/evaluation/evaluation.go:20-29`). `validateAnswerSheetForSkills` performs only structural membership checks (`cli/internal/evaluation/evaluation.go:610-627`), and custom-mode warning copy states that answer-sheet correctness is caller-owned (`cli/internal/command/command.go:91-105`). `deriveResult` creates its matching sets from the pair key (`cli/internal/evaluation/evaluation.go:642-669`). Each built-in answer sheet contains the same five pairs (`evaluations/skill-calling/built-ins/gardening-web-application.json:130-155` and corresponding files). The governance view names them “Required First Activations,” lists later applicable turns, and defines later reloads as additional (`evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:5-30`).

**Implication:** “Missed” means a governed first-activation pair had no recorded event on that exact turn. A call one turn early or late is additional and does not repair the miss. The chart measures these five selected opportunities per scenario, not every turn where a skill could reasonably apply and not every skill decision made by the model.

### Detailed results retain event multiplicity; compact points deliberately collapse it

For every attributed event, `deriveResult` appends one entry to `Observed` in event-log order. It also appends every unexpected event to `Additional`; unattributed multiplicity is likewise retained. A set is used only to decide whether at least one matching event exists. Consequently, two signals for the same skill and turn remain two detailed observations, but either one satisfies the expected pair.

`deriveWebsiteResult` independently converts both expected and observed calls into sets. Its `called` count is the number of unique expected skill names with at least one observed event on that expected turn. `missed` is the number of unique expected skill names on that turn minus `called`. Repeated observed events and duplicate expected pairs cannot inflate the compact counts.

**Evidence:** `Store.Mark` always appends and `Store.Events` preserves every decoded line in order (`cli/internal/runstate/runstate.go:182-213`). `deriveResult` appends before assigning its set key (`cli/internal/evaluation/evaluation.go:647-663`). `deriveWebsiteResult` uses a per-turn skill set and an observed pair set (`cli/internal/evaluation/evaluation.go:696-727`). The duplicate-event test validates that two observed `alpha` calls still produce one called expected skill (`cli/internal/evaluation/evaluation_test.go:79-118`).

**Implication:** Use `website.json` for binary expected-opportunity scoring and `result.json` or `events.jsonl` for invocation frequency. The compact chart cannot answer “how many times was this skill loaded on this turn?”

### Numeric `turn` is the one-based scenario position, independent of `turn_id`

The compact `turn` value is `index + 1` while iterating the scenario's ordered turns. It is not parsed from `turn_id`, a harness-native conversation counter, elapsed time, token count, or context-consumption percentage. `turn_id` preserves the source identifier for detailed correlation. `total_turns` is the full scenario slice length.

Only scenario positions containing at least one expected skill receive a point. The built-in artifact therefore has points at numeric turns 1, 11, 25, and 30, even though the replay contains 30 turns. Turn 1 has denominator two; each other point has denominator one.

**Evidence:** The scenario shape and ordering are `replay.Scenario.Turns` (`cli/internal/replay/replay.go:27-36`). `deriveWebsiteResult` iterates that slice, skips turns absent from `expectedByTurn`, assigns `Turn: index + 1`, and sets `TotalTurns: len(scenario.Turns)` (`cli/internal/evaluation/evaluation.go:710-736`). The README explicitly defines `turn`, `turn_id`, sparse points, and sample-size derivation (`cli/README.md:150-173`). The built-in answer sheets place five expectations on four positions (`evaluations/skill-calling/built-ins/gardening-web-application.json:130-155`).

**Implication:** The Recharts x-axis must be numeric with a domain from 1 through `total_turns`; a categorical axis would space turns 1, 11, 25, and 30 evenly and misrepresent their positions. Intervening turns are unscored, not zero-call or zero-miss observations.

### One `website.json` is one run of one scenario, not a campaign aggregate

`RunRequest` selects one built-in evaluation ID or one custom scenario/answer-sheet set. `Service.Run` produces one replay result, one detailed result, and one compact artifact with one `run_id` and one `scenario_id`. No production aggregation path combines the three governed scenarios, repeated trials, or multiple harness/model cells into one chart series.

For one governed built-in run, sample size is `sum(called + missed) = 5`, independent of whether any signals were observed. Across the planned three-scenario suite it would be 15 only if a downstream consumer explicitly combines the three separate artifacts. The current one-suite MVP methodology specifies one tooling-complete run of all three scenarios for each matrix cell and leaves repeated trials for later, but execution and artifact generation remain per scenario.

**Evidence:** `RunRequest` has singular evaluation/scenario/answer-sheet selectors (`cli/internal/evaluation/evaluation.go:31-47`); `loadEvaluationInputs` returns one scenario and answer sheet (`cli/internal/evaluation/evaluation.go:472-525`); `Service.Run` writes one `result.json` and one `website.json` from that scenario (`cli/internal/evaluation/evaluation.go:281-310`). `WebsiteResult` contains singular run and scenario identifiers (`cli/internal/evaluation/evaluation.go:105-113`). The plan sets three 30-turn scenarios and one tooling-complete three-scenario suite per matrix cell, with repeated trials deferred (`plans/skill-issue-project-completion/05-define-the-skill-calling-evaluation-contract-and-campaign-assets.md:19-24,52-64`).

**Implication:** The website migration needs an explicit presentation choice outside the artifact contract: show each run/scenario separately, or define a transparent campaign aggregation. It must not silently treat one scenario's five opportunities as a full harness/model campaign or merge runs without preserving scenario provenance.

### Raw called/missed lines have sparse and unequal denominators

Every compact point satisfies `called + missed = number of unique expected skills on that turn`. In the governed scenarios, the denominator is two at turn 1 and one at turns 11, 25, and 30. Drawing continuous lines through these sparse raw counts can visually imply observations on the unscored intervening turns and can make the turn-1 value appear inherently larger because it has twice the number of governed opportunities.

**Evidence:** Per-turn expected skills are deduplicated into `expectedByTurn`, and `Missed` is `len(skills) - called` (`cli/internal/evaluation/evaluation.go:696-727`). The governed expected-call tables show two skills at turn 1 and one at each later checkpoint (`evaluations/skill-calling/scenarios/gardening-web-application/expected-calls.md:5-13`). The downstream plan requires using artifact points directly, a numeric full-turn domain, derived sample size, and data-driven limits (`plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:93`).

**Implication:** Public copy should describe “expected first-activation opportunities at selected turns,” retain visible point markers or otherwise disclose sparsity, and avoid interpreting interpolation as measured decay. Cross-turn raw heights should be read with their per-point denominator; a normalized rate would be a new derived view and must remain transparently traceable to `called + missed`.

### The compact artifact intentionally omits evidence needed for deeper interpretation

`website.json` includes schema version, run/scenario identity, harness, model, total turns, and compact points. It omits reasoning, evaluation ID distinct from scenario ID, timestamps, skill names, additional calls, unattributed calls, harness version, operating system, transcript, and raw events. The detailed result adds reasoning, evaluation ID, timestamps, and skill-level classifications, while harness version and OS remain external campaign metadata.

**Evidence:** Compare `Result` and `WebsiteResult` fields (`cli/internal/evaluation/evaluation.go:79-113`). The README states what the compact artifact excludes (`cli/README.md:171-173`). The scenario evidence instructions request harness version and operating system separately (`evaluations/skill-calling/scenarios/gardening-web-application/instructions.md:15-21`).

**Implication:** A credible website card cannot derive all tested-environment and limitation labels from `website.json` alone. It needs vetted campaign metadata and links to `result.json`; skill-level explanations and additional/unattributed evidence must remain in the detailed view rather than being invented from compact counts.

### The event JSON Schema currently conflicts with emitted production events and is not runtime-enforced

The checked-in event schema allows only `schema_version`, `event`, `run_id`, optional `turn_id`, `attributed`, `harness`, `model`, `skill`, and `recorded_at`, with `additionalProperties: false`. Current production `runstate.Event` also always serializes `reasoning` and `evaluation_id`. Therefore a newly emitted current event violates the published schema because those two properties are forbidden. The schema also does not conditionally require `turn_id` when `attributed` is true or forbid it when false.

No production path loads or validates `event.schema.json`. `Store.Events` merely unmarshals each line into the Go struct; it does not check schema version, event name, run ID, attribution consistency, or unknown JSON properties. The Go struct and current code are therefore the executable behavior, while the JSON Schema is a drifted contract artifact.

**Evidence:** `additionalProperties: false`, the required list, and the complete property list appear in `evaluations/skill-calling/event.schema.json:5-50`. The emitted Go event fields include `Reasoning` and `EvaluationID` (`cli/internal/runstate/runstate.go:36-48`) and `Store.Mark` populates them (`cli/internal/runstate/runstate.go:165-177`). `Store.Events` only calls `json.Unmarshal` (`cli/internal/runstate/runstate.go:194-213`). Repository references to `event.schema.json` are documentation/planning references rather than a validation call.

**Implication:** Downstream tooling should not validate current `events.jsonl` against the checked-in schema without first reconciling this drift. This does not change `website.json` derivation, because the CLI reads its own Go event shape directly, but it weakens the claim that raw events presently satisfy the documented portable schema.

### Implemented artifact semantics supersede the current website mock, while campaign publication remains planned

The current React data and chart still use mock `contextConsumed`, `skillCalls`, `skillMisses`, six evenly spaced percentage checkpoints, fixed `n = 120`, and a fixed y-axis maximum of 20. The implemented CLI artifact instead exposes numeric scenario positions, sparse `called`/`missed` counts, per-run sample size five for governed scenarios, and a full domain of 30. The authoritative completion plan explicitly instructs the website to replace the mock through `siteData.ts`, use `turn`/`called`/`missed`, derive sample size, and use data-driven axis limits.

**Evidence:** Current mock types and values are in `src/data/siteData.ts:1-74`; current chart keys and fixed domain are in `src/components/EvaluationChart.tsx:47-86`. Production artifact fields are in `cli/internal/evaluation/evaluation.go:98-113,696-737`. The downstream migration requirement is `plans/skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md:93`.

**Implication:** Context-consumption percentages and `n = 120` have no derivation from the production replay/scoring path and should not survive the migration. Public campaign claims still depend on tooling-complete scenario runs and a disclosed aggregation/presentation method; code support for per-run artifacts is implemented, while qualified multi-run published data is a later workflow outcome.

## Notes

- A focused `go test ./cli/internal/evaluation ./cli/internal/replay ./cli/internal/runstate` validation attempt could not initialize the default Go build cache because the sandbox denied writes under `~/Library/Caches/go-build`. This did not block the source trace. Existing tests were used only as secondary validation after production behavior was established.
- The current working tree contains unrelated in-progress modifications to evaluation/run-state/contracts. Findings describe the inspected production working tree and preserve those changes without editing them.
- The three built-in human-readable expected-call governance documents are materially consistent on first activations, later applicability, and descriptive additional/missing classification.
- Useful search terms: `deriveWebsiteResult`, `deriveResult`, `recordCodexSignals`, `SetActiveTurn`, `skill_invocation`, `Required First Activations`, `website.json`, `total_turns`.
