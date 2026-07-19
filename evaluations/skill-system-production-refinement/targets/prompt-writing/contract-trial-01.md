# Prompt-Writing Target Outcome Interpretation

## Goal

Create prompts that give agents only the task-specific direction needed to act successfully, while supplying goal, completion criteria, and autonomy or approval boundaries when the task needs them and no named authority already owns them.

## Intended Use

Use this discipline when writing prompts for agents, sub-agents, invoked skills, or authoritative plans, including open-ended discovery prompts and revisions after a prompt failure.

## Expected Behavior

- Read any named skill or document before writing the prompt, point the agent to that authority, and add only missing task-specific information.
- Request the smallest deliverable that enables the next action; for discovery, name the investigation category rather than prescribing expected findings.
- Remove context, instructions, and structure the agent already receives from its environment or named authority.
- When a prompt fails, correct the failed decision specifically rather than accumulating broad instructions.

## Observable Expected Result

The resulting prompt is concise but sufficient for the task: it identifies the necessary outcome and completion or authority conditions where ownership is otherwise absent, relies on named sources without duplicating them, requests a minimal actionable deliverable, preserves open-ended discovery, and responds to prior failure with a targeted correction.

## Preserved Boundaries

- Do not impose a universal prompt template or require goal, completion criteria, and autonomy boundaries when the task does not need them or another authoritative source owns them.
- Do not reteach, summarize, or invent structure for named skills or documents.
- Do not widen an open-ended discovery request into a checklist of anticipated findings.
- Do not narrow the discipline beyond its stated prompt targets or extend it to general prose unrelated to agent, skill, or authoritative-plan prompting.
- Do not treat evaluation-friendly wording, exhaustive instructions, or instruction accumulation as intended target behavior.

## Sources Consulted

- `skills/skill-evaluation-and-refinement/references/target-outcome-interpretation.md` — interpretation fields and contract-preservation rule.
- `test-targets/skills/prompt-writing/SKILL.md` — complete target contract, including invocation boundary, conditional authority content, prompt discipline, and failure-repair behavior.

## Ambiguity

The target leaves “when the task needs them” contextual rather than defining a mechanical test for when goal, completion criteria, or an autonomy or approval boundary must be explicit. The controlling constraint is clear: add them only when needed and when no named skill or authoritative document already owns them.
