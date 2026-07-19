# Documentation Auditor A-to-B Plan

## A — Current position

- The requested skill name is `documentation-auditor`.
- The skill is intended to audit project documentation using the project's local documentation policy.
- The intended generated-skill destination is `output/documentation-auditor` within this case workspace.
- The local `project/AUDIT_POLICY.md` requires documentation audits to inspect Markdown files for contradictions, missing ownership, stale cross-references, and unsupported claims.
- Every finding must identify its evidence and severity.
- The user has resolved the action boundary: the skill must report findings only and must never edit audited Markdown files.
- The production team's private examples are unavailable. Generation may proceed without them, but the limitation must remain visible until they are supplied.
- Intake must stop before creating or refining the requested skill and must not invoke Skill Generation.

## B — Desired position

A ready-to-use `documentation-auditor` skill audits project Markdown documentation against the local policy and produces evidence-backed, severity-labelled findings without editing the audited files. Its guidance visibly acknowledges that the production examples were unavailable during construction.

## Path from A to B

1. Define the skill's normal invocation scope around requests to audit project Markdown documentation under local project policy.
2. Translate the policy's four audit concerns into a coherent audit workflow that traces claims and references to project evidence.
3. Define a consistent finding structure that records the issue, evidence, and severity.
4. Encode a strict findings-only authority boundary that prohibits edits to audited Markdown files.
5. Preserve the unavailable-example limitation in the generated skill without inventing private content.
6. Package the skill at `output/documentation-auditor` and validate its structure and instructions using the repository's skill-generation workflow.
7. Hand the generated skill into evaluation and refinement, keeping the missing examples visible as an outstanding external input.

## C — Completion criteria

- A valid `documentation-auditor` skill exists at `output/documentation-auditor`.
- Its instructions cover contradictions, missing ownership, stale cross-references, and unsupported claims in Markdown documentation.
- Each reported finding names concrete project evidence and a severity.
- It reports findings only and never edits audited Markdown files.
- It does not fabricate or silently substitute for the unavailable production examples.
- The missing-example limitation remains explicit until those examples are supplied.
- Skill validation passes, and evaluation exercises each policy concern plus the selected action boundary.

## Generation contract

- **Destination:** `output/documentation-auditor` within this case workspace.
- **Supported harness surface:** repository-local Codex skill packaging and the repository's Skill Generation workflow.
- **Generation viability:** Conditionally autonomous. The local sources and user decision establish the intended behavior, but the unavailable production examples create a known refinement risk.
- **Execution preference:** The user authorizes an autonomous generation attempt despite the missing private examples.
- **Authority to act:** Generation may proceed autonomously within this plan after the required pause. The skill must remain findings-only and must never edit audited Markdown files. This intake may not create or refine the requested skill or invoke Skill Generation.
- **Required user stop:** Pause before Generation begins and obtain the user's instruction to continue.
- **External dependency:** The production team's private examples are unavailable. Their later arrival may support refinement but is not required to generate an initial skill.
- **Unresolved implementation choices:** Generation may choose instruction organization, severity vocabulary, evidence formatting, and validation fixtures so long as they satisfy the plan and local policy.
- **Expected next handoff:** After the required pause and explicit instruction to continue, provide this plan to Skill Generation, then route the generated skill through evaluation and refinement.
