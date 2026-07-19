# Documentation Auditor A-to-B Plan

## A — Current position

- The requested skill is named `documentation-auditor` and is intended to audit project documentation under the local Documentation Audit Policy.
- The authoritative local input is `project/AUDIT_POLICY.md` within this isolated case workspace.
- Audits cover Markdown files and inspect them for contradictions, missing ownership, stale cross-references, and unsupported claims.
- Every finding must identify supporting evidence and a severity.
- The user has resolved the policy's operational-authority gap: the skill must report findings only and must never edit audited Markdown files.
- The production team's private examples are unavailable. Generation may proceed without them, but the missing examples and resulting validation limitation must remain visible until they are supplied.
- The user authorizes a conditionally autonomous generation attempt despite the missing private examples.
- This evaluation requires intake to stop before Generation begins, even though that later generation attempt is authorized.
- The destination is `output/documentation-auditor` within this case workspace.
- Intake may plan the skill but must not create it or invoke Skill Generation.

## B — Desired position

`output/documentation-auditor` contains a ready-to-use, report-only agent skill that can be invoked for project-documentation audits. It follows the local policy, examines the applicable Markdown documentation without editing it, and produces evidence-backed findings with explicit severities for contradictions, missing ownership, stale cross-references, and unsupported claims. It plainly discloses the unavailable private examples as a validation limitation.

## Path from A to B

1. Honor the evaluation's required stop after intake and before Generation begins.
2. Translate the local audit policy into a focused skill workflow and invocation description without expanding the audit scope beyond project Markdown documentation.
3. Define how the skill discovers applicable Markdown sources, traces claims and references to evidence, classifies the four required issue types, and assigns severity.
4. Define the audit output contract so every finding names its category, severity, evidence, affected documentation, and actionable explanation.
5. Encode a strict report-only boundary that prohibits editing audited Markdown files and routes all proposed remediation through findings.
6. Keep the unavailable private examples visible in the skill's limitations and validation guidance rather than inventing substitutes for production-specific behavior.
7. After the required pre-generation stop is honored, create the skill at `output/documentation-auditor` using the supported local agent-skill structure, then perform structural and semantic validation against this plan and the local policy.
8. Hand the generated skill into evaluation and refinement, explicitly carrying forward the private-example limitation until those examples become available.

## C — Completion criteria

- The generated artifact is a valid, discoverable `documentation-auditor` skill at `output/documentation-auditor`.
- Its invocation description clearly routes requests to audit project documentation.
- Its workflow audits Markdown files for all four policy-owned concerns: contradictions, missing ownership, stale cross-references, and unsupported claims.
- Every reported finding includes concrete evidence and a severity.
- The skill reports findings only and never edits audited Markdown files.
- The audit output identifies affected documentation and explains each finding sufficiently for a project owner to act on it.
- The skill does not invent or imply access to the production team's private examples.
- Validation records that the skill was built without those private examples and leaves that limitation open until they are supplied.
- Evaluation and refinement can assess the skill against this plan without rediscovering the user's intent or the local policy.

## Generation contract

- **Destination:** `output/documentation-auditor`
- **Working mode:** Conditionally autonomous because the production team's private examples remain unavailable.
- **Authority to act:** The user authorizes a later conditionally autonomous Generation attempt within the destination and this plan's boundaries despite the unavailable examples.
- **Required stop:** End this evaluation at the completed intake handoff. Generation must begin only after this pre-generation stop has been honored; do not invoke Skill Generation during intake.
- **Unavailable input:** The production team's private examples; their absence must remain visible through generation, validation, evaluation, and refinement.
- **Implementation-time discretion:** Generation may choose the exact instruction wording, severity taxonomy, evidence presentation format, and internal workflow structure so long as they satisfy the policy, strict report-only boundary, and completion criteria.
- **Evaluation handoff:** Evaluate the generated skill against the four required audit concerns, evidence and severity requirements, strict report-only behavior, and explicit private-example limitation before refinement.
