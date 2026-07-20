# Prompt-Writing Production Assessment

## Result

The content in `test-targets/skills/prompt-writing/SKILL.md` showed no material description-selection or loaded-behavior failure in this Codex session.

- **Automatic selection:** four fresh agents received distinct, unleading prompt-writing tasks. All four selected prompt-writing guidance and loaded the exact target before producing their final prompt. Native child-session tool traces, rather than agent self-report or output similarity, establish the loads.
- **Behavior after loading:** three additional fresh agents explicitly loaded the target and exercised named-authority use, open-ended discovery, and repair of an over-broad failed prompt. All three outputs satisfied the target contract with no retained material failure.
- **Static authoring review:** the description is a concise what-it-is sentence followed by a when-to-use sentence, and the body contains direct behavior-changing guidance without repeating activation rules, adding a universal template, or enumerating exhaustive cases.
- **Production conclusion:** the skill's description and body are semantically suitable for production on the evidence observed. A formal cross-environment production pass remains unestablished because the repository does not contain the evaluator's required durable environment-qualification record with an exact model identifier, qualification date, trial method, and prior direct invocation evidence.
- **Primary residual risk:** duplicate skill identity. This session also advertised an installed `eric-wimp-toolkit:prompt-writing` skill with the same description. One selection agent first addressed that installed path before resolving to the test target. The evidence proves reliable selection of the prompt-writing meaning, while unique routing to one packaged copy requires a trial environment that exposes only the production candidate.

No target or production skill was modified.

## Target Contract

- **Target:** `test-targets/skills/prompt-writing/SKILL.md`
- **SHA-256:** `b5b2f86af6c8b0690b16702a86df2acc8ab244d1f4bc64eff44f7dc4035ea541`
- **Goal:** produce prompts containing only the task-specific direction an agent needs, while supplying goal, completion criteria, and autonomy or approval boundaries when needed and not already owned by a named authority.
- **Intended use:** prompts for agents, sub-agents, invoked skills, and authoritative plans, including discovery prompts and prompt-failure repairs.
- **Required behavior:** read named authorities before prompting; point to rather than reteach them; request the smallest next-action deliverable; preserve open discovery; remove inherited information; repair the failed decision specifically.
- **Expected result:** a concise, sufficient prompt that uses existing authority, defines the necessary outcome and stopping or approval boundary, and avoids inventories, duplicated instructions, and broad corrective accumulation.
- **Preserved boundary:** the skill does not impose a universal template or require goal, completion, and autonomy fields when the task or a named authority already supplies them.

## Evidence Standard And Environment

- **Harness:** OpenAI Codex Desktop / Codex CLI `0.145.0-alpha.18`.
- **Provider exposed by native session metadata:** `openai`; the session record does not expose the exact model identifier required by the evaluator's qualification gate.
- **Isolation:** every probe ran through a fresh native sub-agent with `fork_turns: "none"`.
- **Accepted native evidence:** a child-session tool call reading `test-targets/skills/prompt-writing/SKILL.md` before the final response.
- **Rejected as proof:** final-answer quality, a prose statement that the skill was used, or a `Skills Used` list.
- **Qualification limitation:** `evaluations/skill-system-production-refinement/progress.md` states that independent agents are available, but neither it nor the campaign log is a completed environment-qualification record. The assessment therefore reports current-session measurements without converting them into a general reliability claim.

## Automatic-Selection Evidence

All prompts below avoided the skill name, invocation request, description wording, and expected selection.

### Selection 1 — Source-Owned Discovery Prompt

- **Fresh agent/session:** `description_probe_1` / `019f7767-3b0d-78e0-9428-c52e80ff7cda`.
- **Prompt:** “Work in the Skill Issue repository root. Do not edit files. Write a concise task prompt for a fresh agent to inspect how the Vite website obtains benchmark content and identify only the source-owned evidence needed to decide where a new benchmark belongs. The prompt must define a useful stopping point and preserve open-ended discovery. Return only the prompt you would send.”
- **Native load evidence:** before answering, the child emitted an `exec` tool call whose command was `cat test-targets/skills/prompt-writing/SKILL.md`.
- **Observable result:** the prompt requested a read-only trace to the authoritative content owner, limited the deliverable to the minimal supporting path, and defined a stopping condition without prescribing findings.
- **Criterion result:** pass.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-14-41-019f7767-3b0d-78e0-9428-c52e80ff7cda.jsonl`.

### Selection 2 — Failed-Prompt Repair

- **Fresh agent/session:** `description_probe_2` / `019f7767-4cdd-7d01-a517-15cb4bbe10b6`.
- **Prompt:** “Work in the Skill Issue repository root. Do not edit files. A previous delegated investigation returned a broad inventory instead of the decision-enabling evidence. Write the revised prompt you would send to a fresh agent so it determines which production module owns reusable evaluation-chart behavior and returns the smallest evidence needed for the next implementation decision. Return only that prompt.”
- **Native load evidence:** the child first attempted to read the advertised installed `eric-wimp-toolkit` prompt-writing path, found that cached path had drifted, located available prompt-writing copies, then read `test-targets/skills/prompt-writing/SKILL.md` before answering.
- **Observable result:** the revised prompt corrected the failed deliverable decision, bounded evidence to at most three references, and excluded inventory, planning, and recommendations.
- **Criterion result:** pass for semantic selection and body behavior; qualified for unique package routing because two same-named/same-described copies were advertised.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-14-46-019f7767-4cdd-7d01-a517-15cb4bbe10b6.jsonl`.

### Confirmation 1 — Named Authoritative Plan

- **Fresh agent/session:** `description_probe_3` / `019f7768-71f2-7212-8c46-121f683b9ca1`.
- **Prompt:** “Work in the Skill Issue repository root. Do not edit files. Write the instruction you would give another agent to use `evaluations/skill-system-production-refinement/progress.md` as the authority for identifying the campaign's next unfinished action. The delegate should return only the decision and the smallest supporting evidence needed to resume work. Return only the instruction you would send.”
- **Native load evidence:** the child read the exact target and the named progress document in one `exec` tool call before answering.
- **Observable result:** the prompt pointed to the authoritative document, requested the next unfinished action and minimal support, and did not reteach the document.
- **Criterion result:** pass.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-16-01-019f7768-71f2-7212-8c46-121f683b9ca1.jsonl`.

### Confirmation 2 — Open-Ended Layer Diagnosis

- **Fresh agent/session:** `description_probe_4` / `019f7768-8a59-7003-a474-9e05a3473e33`.
- **Prompt:** “Work in the Skill Issue repository root. Do not edit files. You need to delegate an open-ended investigation into why a skill might appear suitable from its metadata yet behave poorly after activation. Write the prompt for a fresh agent so the result supports deciding which layer owns the failure, without prescribing an inventory of expected causes. Return only the prompt.”
- **Native load evidence:** the child read `test-targets/skills/prompt-writing/SKILL.md` before answering.
- **Observable result:** the prompt defined the investigation category, evidence-backed ownership decision, remaining-uncertainty stop, and read-only authority boundary without listing suspected causes.
- **Criterion result:** pass. The output's `Skills Used` self-report was ignored as invocation evidence.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-16-07-019f7768-8a59-7003-a474-9e05a3473e33.jsonl`.

### Description Decision

Current-session measurement is **4/4 proactive loads and 4/4 contract-satisfying outputs** across two initial prompts and two distinct confirmations. The description covers concrete delegation, named-authority prompting, open-ended discovery, and failed-prompt repair without apparent widening or narrowing. No description refinement is supported by the evidence.

The formal evaluator campaign should record this as a measured current-session pass with a production-qualification hold, rather than a general harness-reliability pass, until the missing qualification record and duplicate-package isolation are resolved.

## Loaded-Behavior Evidence

These cases were separate from description selection: each fresh agent was explicitly instructed to read and apply the target before writing its prompt. The native trace confirms that each did so.

### Body Case 1 — Named Authority

- **Fresh agent/session:** `body_probe_authority` / `019f7769-a87f-7cc0-b349-5f8220308053`.
- **Native load evidence:** the child read the exact target and `evaluations/skill-system-production-refinement/progress.md` before answering.
- **Output:** “Work in the Skill Issue repository root. Read `evaluations/skill-system-production-refinement/progress.md` and use it as the authority to determine the campaign’s next action. Do not edit files. Return only the next action, or the smallest authority conflict that must be resolved before selecting it.”
- **Audit:** points to and relies on the authority; does not summarize or invent its structure; asks for the smallest next-action result; supplies the read-only boundary and an ambiguity stop.
- **Criterion result:** pass.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-17-20-019f7769-a87f-7cc0-b349-5f8220308053.jsonl`.

### Body Case 2 — Open-Ended Discovery

- **Fresh agent/session:** `body_probe_discovery` / `019f7769-b8c3-7371-90ac-fd9e6248c9b1`.
- **Native load evidence:** the child read the exact target before answering.
- **Output characteristics:** asks which observable evidence Codex exposes for proactive activation; requests a concise source-backed assessment of evidentiary strength; stops when the evidence category and limits are explainable; retains read-only authority; avoids an inventory of expected evidence types or causes.
- **Audit:** preserves open discovery, defines the category rather than anticipated findings, supplies a usable outcome and stopping point, and requests only the decision-enabling deliverable.
- **Criterion result:** pass.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-17-25-019f7769-b8c3-7371-90ac-fd9e6248c9b1.jsonl`.

### Body Case 3 — Targeted Failure Repair

- **Fresh agent/session:** `body_probe_repair` / `019f7769-c86c-7293-9c77-ce622ef0d499`.
- **Native load evidence:** the child read the exact target before answering.
- **Output:** “Determine whether native harness or session evidence can directly prove that a skill was proactively loaded before its instructions influenced an agent’s answer. Inspect only the evidence needed to resolve that question. Return a concise conclusion, identify the exact evidence surface that supports it, distinguish direct proof from inference, and state the precise evidentiary gap if proof is unavailable. Keep the work read-only; do not propose fixes or a broader evaluation campaign.”
- **Audit:** corrects the failed scope and deliverable decisions directly; removes the broad inventory, fix proposals, and complete-campaign planning; preserves the actual question and supplies a precise evidence stop.
- **Criterion result:** pass.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-17-29-019f7769-c86c-7293-9c77-ce622ef0d499.jsonl`.

### Body Decision

Loaded behavior is **3/3 passing cases with no retained material failure**. The body produces concise but sufficient prompts, respects named authority, preserves open-ended discovery, requests minimal actionable results, and repairs only the failed prompt decision. No body refinement is supported by the evidence.

## Production Recommendation

Promote the target content unchanged when the production package preserves this description and body. Before representing automatic invocation as generally qualified, create the evaluator-required environment record and rerun the four description trials with the production candidate as the only advertised prompt-writing skill. Those steps isolate packaging and harness reliability from the semantic content that passed here.
