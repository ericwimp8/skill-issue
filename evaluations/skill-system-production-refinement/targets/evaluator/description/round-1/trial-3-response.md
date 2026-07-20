# Dictate Plan Read-Only Assessment

## Decision

- **Description selection:** incomplete formal result. Three fresh Codex agents independently received natural, successive-message planning requests and each read the exact candidate `test-targets/skills/dictate-plan/SKILL.md` before producing an answer. This is three direct selection successes, but the required four-trial protocol was stopped before completion.
- **Loaded behavior:** unassessed. All three agents were interrupted before returning an observable response or plan, and the evaluator contract prohibits body execution until four retained description trials pass. No claim about instruction effectiveness is supported by the current runtime evidence.
- **Current campaign state:** description remains `unvalidated`; body remains blocked by the description gate. The evidence supports a bounded statement that the current harness selected the candidate in three natural requests, not a complete reliability pass.
- **Changes:** no production skill or copied target was modified. No refinement was performed.

## Target Contract

- **Target:** `test-targets/skills/dictate-plan/SKILL.md`
- **SHA-256:** `44bc9abeeebc2200ec5c3b6ea7a65cab70551e72c1b9f7c06c281c2b2ec0fe2b`
- **Goal:** maintain one living sequence of broad tasks that transforms the user's starting position A into desired position B and creates observable outcomes C that demonstrate B.
- **Intended use:** planning developed through successive conversational messages, treated as source material to integrate rather than text to transcribe.
- **Required behavior:** maintain distinct A, B, path, and C meanings; keep tasks broad and dependency-ordered; omit implementation procedure; create and reread one plan document; integrate superseding dictation; retain material uncertainty; confirm resolved vague details; and use the exact brief update plus Expand Dictate Plan handoff.
- **Expected result:** a coherent living `plans/<task-slug>/<task-slug>-a-to-b-plan.md` whose current state, desired state, broad path, and observable completion criteria remain semantically consistent across turns.
- **Preserved boundaries:** existing conditions remain in A rather than being duplicated in C; execution choices remain open; vague details and proposed missing C items require confirmation; the user controls completion and expansion.

## Environment And Evidence Standard

- **Qualified surface:** Codex Desktop fresh sub-agents backed by Codex CLI `0.145.0-alpha.18`, using `gpt-5.6-sol` with high reasoning, as recorded in `evaluations/skill-system-production-refinement/environment-qualification.md` on 2026-07-19.
- **Discovery path:** `.codex/skills/dictate-plan` is a project-local link to `test-targets/skills/dictate-plan`.
- **Isolation:** each observed request ran in a fresh native sub-agent with `fork_turns: "none"` and a distinct `<temporary-output-root-N>`.
- **Accepted selection proof:** the qualification record accepts a native pre-output tool trace reading the exact candidate `SKILL.md`. Agent prose, answer similarity, and a `Skills Used` list are excluded.
- **Body proof requirement:** a complete multi-turn transcript plus the produced plan must be audited against pre-recorded semantic criteria. Loading alone cannot prove correct behavior.

## Retained Selection Evidence

The three prompts below did not name the skill, request invocation, quote its description, or state the expected selection. Runtime locations use stable placeholders; all other prompt content is retained. The agents were interrupted after direct load evidence was emitted, so none of these trials has an observable behavior result.

### Probe 1 — Equipment Request Portal

- **Fresh agent:** `<fresh-agent-1>`
- **Session:** `019f7772-e553-7853-ac9f-d014c12519d3`
- **Prompt:** "Work only inside `<temporary-workspace-1>`; do not read or modify any files in the Skill Issue repository except repository guidance or skills that the harness selects normally. I want to work out a plan with you across a few messages for replacing a team's spreadsheet-based equipment requests with a small internal request portal. Today requests arrive through a shared spreadsheet, managers approve rows manually, and the IT team fulfills approved items. We want staff to submit requests, managers to approve them, and IT to track fulfillment in one place, with a visible audit trail showing who approved and fulfilled each request. This is my first block of thoughts; begin the living planning artifact and tell me what you changed. I will send more details later. Keep all generated artifacts under the stated temporary workspace."
- **Native evidence:** the first attempted command read the exact target from `test-targets/skills/dictate-plan/SKILL.md`; after discovering that the temporary working directory did not yet exist, the agent created only that temporary directory and read the exact target again.
- **Selection result:** pass for direct candidate loading.
- **Behavior result:** unavailable because the agent was interrupted before output or plan creation.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-27-26-019f7772-e553-7853-ac9f-d014c12519d3.jsonl`

### Probe 2 — Automated Customer-Health Report

- **Fresh agent:** `<fresh-agent-2>`
- **Session:** `019f7773-0688-75c1-8a98-a81b65e1defa`
- **Prompt:** "Work only inside `<temporary-workspace-2>`; do not read or modify any files in the Skill Issue repository except repository guidance or skills that the harness selects normally. Help me shape the work over this and later messages for moving a weekly customer-health report from manual analyst preparation to a reliable internal service. Analysts currently combine exports from three systems every Friday. The intended result is an automatically prepared report that analysts review before distribution, and we must be able to observe whether source data is fresh and whether each report was reviewed. Start organizing the work from this first batch of notes, maintain whatever planning artifact is appropriate, and briefly tell me what changed. I will add constraints in a later message. Keep all generated artifacts under the stated temporary workspace."
- **Native evidence:** before any final response, the agent read `test-targets/skills/dictate-plan/SKILL.md` from the repository root.
- **Selection result:** pass for direct candidate loading.
- **Behavior result:** unavailable because the agent was interrupted before output or plan creation.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-27-34-019f7773-0688-75c1-8a98-a81b65e1defa.jsonl`

### Probe 3 — Release Communication Process

- **Fresh agent:** `<fresh-agent-3>`
- **Session:** `019f7773-1e67-7c33-95f5-803340bfd078`
- **Prompt:** "Work only inside `<temporary-workspace-3>`; do not read or modify any files in the Skill Issue repository except repository guidance or skills that the harness selects normally. I am going to talk through the work needed to replace ad-hoc release notes with a consistent product release communication process. Right now each product squad writes notes differently and support cannot reliably tell what customers were told. We need one maintained release record per release, a review before publication, and a way for support to verify which changes were communicated. This is the opening set of notes. Turn it into the initial planning artifact and tell me how you incorporated it; expect another message from me. Keep all generated artifacts under the stated temporary workspace."
- **Native evidence:** before any final response, the agent read `test-targets/skills/dictate-plan/SKILL.md` from the repository root.
- **Selection result:** pass for direct candidate loading.
- **Behavior result:** unavailable because the agent was interrupted before output or plan creation.
- **Session record:** `~/.codex/sessions/2026/07/19/rollout-2026-07-19T08-27-40-019f7773-1e67-7c33-95f5-803340bfd078.jsonl`

## Completion Framework

### Description Loop

1. Start a fresh `fork_turns: "none"` agent for each trial and give it a unique temporary output root.
2. Use two representative natural prompts followed by two distinct confirmation prompts. Each prompt should clearly describe planning developed over successive messages without naming the skill, requesting invocation, quoting metadata, or disclosing expected selection.
3. Preserve the prompt content with runtime locations represented by stable placeholders, plus agent identity, exact model and harness metadata, session record, and native pre-output load trace.
4. Count a trial only when the trace identifies the exact candidate. Ignore response quality and self-report as selection proof.
5. Pass description selection only at four successes across the two-stage protocol. On any failure, retain the evidence and keep body evaluation blocked. This read-only assessment does not diagnose or refine the description.

One additional distinct confirmation trial is required to complete the current four-trial selection measurement. Because the three existing agents were interrupted, a strict campaign record should also state that selection loaded but trial execution did not reach an observable output.

### Body Loop

Begin only after the description state records four retained selection successes. Body cases should explicitly load the exact target so the result measures instructions rather than discovery, and each conversation should proceed one user turn at a time in its own temporary root.

Use at least these varied cases:

1. **Whole-plan reintegration:** begin with A, B, an initial broad path, and observable C; add a later message that supersedes an earlier assumption and inserts a new dependency. Verify that the entire document is reread, superseded meaning is replaced, task order is corrected, repetition is consolidated, and unresolved material remains visible.
2. **A/path/C ownership:** supply an external precondition that the finished system consumes, work that creates the result, and a distinct produced verification capability. Verify that the precondition appears only in A, construction appears in the path, the produced verification appears in C, and tasks remain broad rather than procedural.
3. **Confirmation and completion:** introduce a vague artifact reference that affects task boundaries, then indicate completion after the detail is confirmed. Verify that likely detail is presented before inclusion, missing C items are proposed but not added without confirmation, the final A-to-B route is coherent, and the response ends only with the required change note and Expand Dictate Plan handoff.

For every case, preserve:

- the verbatim turn-by-turn transcript;
- the exact loaded-target trace;
- every version of the temporary plan;
- ground truth stated as semantic conditions and result characteristics rather than one required wording;
- an audit of every contract criterion, including file path, A/B/path/C separation, dependency order, whole-document integration, confirmation gates, and closing-response boundary.

Pass loaded behavior only when every required case and criterion passes without a retained material failure. Output similarity or a plausible-looking plan is insufficient.

## Limitations And Cleanup

- The selection sample is three direct loads, short of the evaluator's four-trial pass threshold.
- No complete child response, plan artifact, follow-up turn, or completion turn exists, so instruction effectiveness remains unknown.
- The direct reads establish selection of the exact candidate on these sessions; they do not establish behavior reliability across models, Codex versions, repositories, or other harnesses.
- The temporary output roots contained no plan files when inspected after interruption. They are transient evaluation locations rather than campaign evidence owners.
- No production or copied target content changed, and no refinement proposal is made.

## Skills Used

- `skill-issue:skill-evaluation-and-refinement`
- `skill-issue:document-update-discipline`
