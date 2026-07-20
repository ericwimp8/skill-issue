# Local Target Outcome Interpretations

## Document Update Discipline

- Local path: `test-targets/skills/document-update-discipline/SKILL.md`
- Canonical source: Eric Wimp Toolkit `document-update-discipline`
- Goal: prevent proximal document patching by making semantic ownership determine where an update belongs.
- Intended use: planning or applying a document update whose correct semantic home must be established.
- Expected behavior: map the document's purpose, observation point, semantic owner, related manifestations, and preserved meanings; apply the smallest complete change at the owner; reconcile the whole document.
- Expected result: one coherent document in which the requested meaning lives at its semantic owner, contradictions and duplicate ownership are resolved, and unrelated meaning remains intact.
- Boundary: the skill governs semantic placement and reconciliation, not the document's subject-matter decisions.

## Prompt Writing

- Local path: `test-targets/skills/prompt-writing/SKILL.md`
- Canonical source: Eric Wimp Toolkit `prompt-writing`
- Goal: produce task prompts containing only the goal, completion criteria, authority boundary, and task-specific context the receiving agent genuinely needs.
- Intended use: writing prompts for agents, sub-agents, invoked skills, or authoritative plans.
- Expected behavior: rely on named skills or documents for what they already own, request the smallest useful deliverable, keep open discovery categorical, and remove redundant or over-prescriptive context.
- Expected result: a concise, self-sufficient prompt that supplies missing task authority and context without reteaching the receiving agent or prescribing unnecessary method.
- Boundary: the skill shapes prompts; it does not own the underlying task procedure already supplied by another authority.

## Retired Dictate Plan Campaign Target

- Historical local path: `test-targets/skills/dictate-plan/SKILL.md`
- Canonical source: Eric Wimp Toolkit `dictate-plan`
- Current state: the copied target was removed after the completed production-refinement campaign; retained campaign evidence records the evaluated content hash and results.
- Goal: maintain a living, dependency-ordered A-to-B plan from successive conversational dictation.
- Intended use: multi-message planning in which the user develops a broad task sequence conversationally.
- Expected behavior: classify meaning into A, B, path, and C; integrate successive dictation semantically; preserve intent; surface planning-level gaps; keep implementation choices open; maintain the durable plan and expansion handoff.
- Expected result: one coherent plan whose ordered broad tasks transform the established current position into the desired position and whose observable criteria demonstrate completion.
- Boundary: the skill plans broad work and does not execute it or turn the plan into an exhaustive implementation prescription.
