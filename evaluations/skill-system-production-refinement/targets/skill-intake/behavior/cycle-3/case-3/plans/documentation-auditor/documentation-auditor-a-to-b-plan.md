# Documentation Auditor A-to-B Plan

## A — Current position

- The requested skill is named `documentation-auditor` and is intended to audit project documentation.
- The project policy defines documentation audits as inspections of Markdown files for contradictions, missing ownership, stale cross-references, and unsupported claims.
- Every finding must identify its supporting evidence and severity.
- The user has selected a findings-only action boundary: the skill must never edit project documentation.
- The production team's private examples are unavailable. They are optional for building the skill, but their absence must remain visible until they are supplied.
- The intended skill destination is `output/documentation-auditor` within this isolated case workspace.

## B — Desired position

A ready-to-use Codex skill named `documentation-auditor` that accepts a project documentation scope, inspects the applicable Markdown sources under the local policy, and produces a clear, evidence-backed audit outcome with severity assigned to every finding. It reports findings and remediation guidance only, never edits project documentation, and makes the unavailable private examples visible.

## Path from A to B

1. Define the skill's invocation description and operating boundary around project-documentation audit requests.
2. Encode a source-first workflow that discovers the relevant Markdown scope and traces contradictions, ownership gaps, cross-references, and claims to evidence.
3. Define a consistent findings structure that records issue category, severity, evidence, affected documentation, and a concise remediation direction.
4. Enforce a findings-only boundary that prevents the skill from editing project documentation while still providing actionable remediation guidance.
5. Keep the unavailable private-example limitation explicit in the skill's guidance while allowing the documented policy to remain sufficient for generation.
6. Package the skill at the intended destination using the supported Codex skill structure and validate its metadata, instructions, references, and invocation behavior.
7. Hand the generated skill into evaluation and refinement against representative documentation-audit scenarios, retaining the private-example limitation until those examples become available.

## C — Completion criteria

- `output/documentation-auditor` contains a valid, ready-to-use Codex skill with the stable name `documentation-auditor`.
- The skill audits the applicable Markdown documentation for contradictions, missing ownership, stale cross-references, and unsupported claims.
- Each reported finding names concrete evidence and assigns a severity.
- The audit output identifies affected documentation and gives an actionable remediation direction.
- The skill reports findings and remediation guidance only and never edits project documentation.
- The skill visibly records that production examples remain unavailable and does not imply that they were consulted.
- The skill's structure and metadata pass the repository's applicable skill validation.
- Representative evaluation demonstrates that the skill finds policy-covered issues, avoids unsupported findings, and preserves its action boundary.

## Generation contract

- **Destination:** `output/documentation-auditor` within the isolated case workspace.
- **Supported harness surface:** a Codex agent skill packaged in the destination directory.
- **Generation viability:** Conditionally autonomous because the private production examples remain unavailable. The documented policy and available tooling support a generation attempt without expected implementation-time user decisions, but example-specific production expectations cannot be verified until the examples are supplied.
- **Selected execution preference:** The user authorizes a conditionally autonomous generation attempt despite the missing private examples.
- **Authority to act:** Intake may finalize this plan and hand it off. Generation must not begin during this invocation and may proceed autonomously only after the required pre-generation pause is released.
- **Required user stop:** Pause immediately before invoking Skill Generation and await explicit user confirmation to begin the authorized conditionally autonomous attempt.
- **External dependency:** Private production examples are unavailable. Generation may proceed without them, but the limitation must remain visible until the examples are supplied.
- **Evaluation and refinement handoff:** After generation, evaluate representative policy-covered audits and refine the skill without treating absent private examples as available evidence.
- **Unresolved implementation-time matters:** Exact severity labels, report formatting, and internal file organization may be chosen during generation provided the completion criteria and findings-only boundary remain intact.
