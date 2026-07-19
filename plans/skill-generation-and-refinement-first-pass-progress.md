# Skill Generation and Refinement First-Pass Progress

> **Document status — completed historical execution record.** This file records the written first pass under its original JetBrains-era candidate set and non-execution boundary. Current support and delivery requirements are owned by the [completed support contract](skill-issue-project-completion/01-reconcile-the-definitive-product-support-and-evidence-contract.md), [direct-install architecture](skill-issue-project-completion/02-research-and-define-direct-harness-installation-architecture.md), and [six-block completion plan](skill-issue-project-completion/skill-issue-project-completion-a-to-b-plan.md). Do not update this record to imitate later execution; follow the [documentation authority map](skill-issue-project-completion/document-authority-and-update-map.md).

## Goal

Write the complete first-pass set of skills and supporting references that take a user from skill intake through skill generation and into skill evaluation and refinement. This pass validates the written system against the plans without running its evaluation or generation workflows.

## Scope Boundary

- Included: the skill-evaluation-and-refinement skill, the skill-intake skill, the skill-generation skill, their required references, and supported harness packaging metadata needed to distribute or configure those skills.
- Included: a bounded requirements audit and two outside-in semantic walkthroughs of the complete intake-to-refinement flow.
- Excluded: executing the new workflows, forward-testing them, environment/model/harness reliability evaluation, standalone CLI implementation, website implementation, benchmarks, and publishing.
- Harness candidates: GitHub Copilot, Claude Code, OpenAI Codex, Cursor, Google Antigravity/Gemini CLI, JetBrains AI/Junie, OpenCode, Kilo Code, and Pi.
- Harness support rule: retain a harness surface only when authoritative evidence establishes a usable plugin, package, extension, registry, or equivalent distribution/configuration path for the required skills.

## Progress Status

- Current phase: complete.
- Last completed item: second outside-in semantic walkthrough reconciled and structurally validated.
- Next item: runtime execution and evaluation remain deferred to a later pass.
- Blocking conflicts: none recorded.

## Sequential Task List

### Phase 1: Establish the Execution Baseline

- [x] Read the parent plan, all expanded task plans, current nine-harness popularity report, packaging synthesis, and existing test-target skills.
- [x] Establish the first-pass scope, exclusions, harness-selection rule, and final completion boundary.
- [x] Create this durable dependency-ordered progress document.
- [x] Audit this task list against every parent and expanded-task completion criterion.
- [x] Reorder or add tasks required to close any coverage gap before implementation begins.
- [x] Record any conflicting or unreasonable completion criteria without allowing them to stall unrelated work.

### Phase 2: Establish Evaluation Targets and Shared Constraints

- [x] Read each complete local target and record its local path, canonical source, goal, intended use, expected behavior, and observable expected result.
- [x] Verify that each interpretation preserves the complete target contract without evaluation-driven expansion or narrowing.
- [x] Create a reusable target-outcome interpretation reference for later evaluation design.
- [x] Establish the shared semantic-refinement constraint set from document-update discipline, code-implementation discipline, system-change ownership, and skill-authoring discipline.
- [x] Decide and document whether each constraint is an installed dependency, an inline equivalent, or both, without duplicating ownership.

### Phase 3: Establish Supported Evaluation Harnesses

- [x] Reconcile the nine selected harness entries with the packaging research and retain only skill-capable surfaces that can carry the evaluation/refinement skill.
- [x] Verify each retained surface's current implicit-invocation controls from the research corpus or authoritative official documentation.
- [x] Verify each retained surface's independent sub-agent controls and the execution evidence available to prove skill invocation.
- [x] Classify each harness/model combination as primary, tentative, configuration-supported, unsupported, or deferred without implying unperformed reliability evaluation.
- [x] Define the description loop's evidence standard and bounded exhaustion behavior.
- [x] Create one curated evaluation-harness reference covering harness identification, implicit-invocation gates, sub-agent gates, evidence, fallbacks, and official sources.

### Phase 4: Write Skill Evaluation and Refinement

- [x] Create the `skill-evaluation-and-refinement` skill with a concise trigger description and a lean two-loop workflow.
- [x] Make the description loop operate independently from body evaluation.
- [x] Require two unleading initial trials followed by two fresh confirmation trials, each using an independent sub-agent.
- [x] Gate description trials on supported implicit invocation and sub-agent availability, with clear unsupported-state reporting.
- [x] Require evaluation-owned inputs, retained evidence, cleanup, semantic diagnosis, and generalized description updates.
- [x] Make the behavior loop derive its contract, observable completion criteria, surface classification, representative cases, and ground truth from the target skill's intended outcome.
- [x] Support code, document, artifact, single-turn chat, multi-turn conversation, and other observable skill surfaces without prescribing one exact output unnecessarily.
- [x] Require conversational evaluation through independent turn-by-turn interaction with a verbatim transcript or equivalent record.
- [x] Gate behavior evaluation on required execution capabilities and provide actionable user alternatives when the harness cannot supply them.
- [x] Offer automatic refinement or human review before body changes, applying the same semantic constraints to automatic edits and proposed chat changes.
- [x] Require generalized semantic updates, cycle cleanup, retained campaign evidence, and fresh execution after each update.
- [x] Enforce five unsuccessful behavior cycles per target, pausing the whole campaign for user direction and accepting only a user-authorized number of additional cycles.
- [x] Produce an explicit passing or stopped result for both description and behavior loops.
- [x] Index every supporting reference with concise selection guidance.

### Phase 5: Establish Explicit-Only Intake Support

- [x] Reconcile the nine selected harness entries with authoritative packaging and skill-control evidence for the intake skill.
- [x] Retain only harness surfaces with a credible way to install, register, or distribute the intake skill.
- [x] Verify the strongest enforceable explicit-only control for every retained surface.
- [x] Define explicit invocation guidance for surfaces without an enforceable control.
- [x] Keep supported configuration separate from evaluated invocation reliability.
- [x] Create a concise, extensible explicit-invocation reference with authoritative links, confirmed controls, unsupported controls, and required fallbacks.
- [x] Derive the shared and platform-specific intake requirements from that reference.

### Phase 6: Write Skill Intake

- [x] Create the stable user-facing `skill-intake` skill and make it explicit-only through its description and every supported platform control.
- [x] Support invocation with an initial skill description or invocation alone followed by a concise request for that description.
- [x] Adapt the living A-to-B model to skill creation: current state, desired ready-to-use skill, dependency-ordered creation path, and observable completion criteria.
- [x] Treat successive user messages as source material for one coherent living plan rather than a transcript.
- [x] Require relevant source code, documentation, configuration, tooling, and project context to be investigated before asking the user.
- [x] Preserve model responsibility for evidence-supported organization and decisions while reserving focused questions for unresolved user-owned intent.
- [x] Surface material omissions without deciding ordinary implementation details prematurely.
- [x] Establish plan readiness before generation handoff.
- [x] Assess autonomous, conditionally autonomous, or ongoing-participation generation modes and explain material risks.
- [x] Record the user's chosen working mode, including authorization to proceed despite known limitations.
- [x] Hand Skill Generation the build-ready plan, confirmed context, unresolved implementation-time matters, expected outcomes, completion criteria, and interaction boundary without creating the skill during intake.
- [x] Index the explicit-invocation and planning references with concise selection guidance.

### Phase 7: Write Skill Generation

- [x] Create a lean `skill-generation` skill that owns execution of the build-ready intake contract.
- [x] Require generation to use the confirmed project context and investigate unresolved implementation-time matters without rediscovering user intent.
- [x] Honor the recorded autonomous or user-involved working mode and stop only for decisions that remain user-owned under that boundary.
- [x] Generate an idiomatic, concise skill and only the references, scripts, assets, or platform metadata its behavior genuinely requires.
- [x] Apply skill-authoring, semantic document-update, code-implementation, system-ownership, and prompt-writing constraints where relevant.
- [x] Validate structural completeness without executing or forward-testing the generated skill during this first pass.
- [x] Produce a clear handoff into Skill Evaluation and Refinement containing the generated skill, intake contract, expected outcomes, completion criteria, and unresolved limitations.

### Phase 8: Establish First-Pass Packaging Metadata

- [x] Establish one canonical source tree for the three workflow skills and their shared references.
- [x] Add the proven Codex plugin wrapper used by the current implementation environment.
- [x] Record the native package path and differing invocation metadata for every other retained harness without creating divergent rendered skill copies during this skills-only pass.
- [x] Keep renderer implementation, release staging, the standalone CLI, global executable installation, account authorization, secrets, and unrelated host configuration outside this first-pass package work.
- [x] Mark platform limitations and explicit-only fallbacks accurately rather than claiming feature parity.
- [x] Verify that current metadata points to canonical skill content and that later adapters can render from that source.

### Phase 9: Validate Written Artifacts

- [x] Validate every skill's frontmatter, folder name, description form, reference index, and absence of body-level trigger guidance.
- [x] Check that each skill remains concise, behavior-changing, and free from auxiliary documentation or evaluation-case accumulation.
- [x] Check all local paths, cross-references, platform claims, and cited official sources.
- [x] Confirm that no workflow, evaluation, sub-agent trial, generated-skill run, CLI action, publishing action, or benchmark was executed.

### Phase 10: Perform the Bounded Requirements Audit

- [x] Audit the implementation once against every expanded task and parent completion criterion.
- [x] Fix each concrete unmet requirement through the semantic owner of the failure.
- [x] Record unresolved conflicts, unsupported criteria, or intentionally deferred runtime proof and continue without entering an audit loop.
- [x] Re-run structural validation only for files changed by audit fixes.

### Phase 11: Perform Two Outside-In Semantic Walkthroughs

- [x] Walk through the complete experience from a user's first intake invocation through generation and evaluation/refinement.
- [x] Correct unclear ownership, missing information, awkward interaction, broken handoffs, or inconsistent terminology using the relevant governing skills.
- [x] Repeat the full walkthrough from a fresh outside perspective.
- [x] Correct any remaining semantic-flow problems without expanding scope or running the workflows.
- [x] Confirm the three skills form one coherent first-pass system.

### Phase 12: Close the First Pass

- [x] Reconcile this progress document with the final artifact tree and audit record.
- [x] Confirm every sane parent and expanded-task completion criterion is met by the written system or explicitly recorded as runtime-deferred.
- [x] Confirm every task above is complete or has a documented accepted conflict.
- [x] Mark the first pass complete and provide the user with the artifact paths, supported harness scope, audit findings, and deferred runtime work.

## Requirement Coverage Audit

| Requirement source                                            | Covered by            |
| ------------------------------------------------------------- | --------------------- |
| Parent: evaluation and iterative improvement process          | Phases 2–4, 9–11      |
| Parent: intake gathers and clarifies build requirements       | Phases 5–6, 9–11      |
| Parent: generation executes intake contract in selected mode  | Phase 7, 9–11         |
| Parent: generated skill enters refinement or user stop        | Phases 4 and 7, 10–11 |
| Evaluation Task 1: target interpretation                      | Phase 2               |
| Evaluation Task 2: description loop and harness gates         | Phases 3–4            |
| Evaluation Task 3: behavior loop and safety controls          | Phases 2 and 4        |
| Intake Task 1: cross-harness explicit-only requirements       | Phase 5               |
| Intake Task 2: conversational planning and generation handoff | Phase 6               |
| Current nine-harness research and packaging boundary          | Phases 3, 5, and 8    |
| Bounded first-pass audit                                      | Phases 9–10           |
| Two complete semantic walkthroughs                            | Phase 11              |

## Conflicts and Deferred Proof

- Runtime proof that descriptions trigger proactively is deferred because this pass must not run the workflows. The written loop, evidence requirements, and stopping behavior can be audited structurally; actual pass evidence requires the later execution pass.
- Runtime proof that generated or refined skills perform their intended tasks is deferred for the same reason.
- Model identifiers and comparative model/harness reliability remain deferred environment-evaluation work. References may document configuration surfaces without claiming measured reliability.

## Completion Rule

The first pass is complete when every checklist item is checked, or an item has a documented accepted conflict; all three workflow skills and required references exist; retained harness support is evidence-backed; the bounded requirements audit has been reconciled; and two outside-in semantic walkthroughs confirm a coherent intake-to-generation-to-refinement flow.
