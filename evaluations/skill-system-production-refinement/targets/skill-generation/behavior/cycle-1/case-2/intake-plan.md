# Documentation Auditor Generation Contract

## A — Current position

- The local policy requires Markdown audits for contradictions, missing ownership, stale cross-references, and unsupported claims.
- Findings require evidence and severity.
- The user selected report-only behavior; the skill never edits audited Markdown.
- Private production examples are unavailable and their absence must remain visible.

## B — Desired position

A valid Codex `documentation-auditor` skill produces evidence-backed, severity-labelled findings and remediation guidance without editing source documentation.

## Path from A to B

1. Create the minimal skill and concise invocation metadata.
2. Encode source-first audit and findings-only behavior.
3. Preserve the missing-example limitation.
4. Validate structure and hand runtime criteria to Evaluation.

## C — Completion criteria

- All four policy concerns are covered.
- Every finding names evidence and severity.
- Audited Markdown is never edited.
- Missing examples remain explicit.
- Only required resources are created.
- Structural validation passes; runtime proof is deferred.

## Generation contract

- Destination: `evaluations/skill-system-production-refinement/targets/skill-generation/behavior/cycle-1/case-2/output/documentation-auditor`.
- Generation viability: Conditionally autonomous.
- Selected execution preference: Authorized autonomous attempt.
- Authority: create only inside the destination and run structural checks.
- Limitation: private examples unavailable.
- Evaluation handoff: provide goal, behavior, boundaries, runtime criteria, and limitation.
