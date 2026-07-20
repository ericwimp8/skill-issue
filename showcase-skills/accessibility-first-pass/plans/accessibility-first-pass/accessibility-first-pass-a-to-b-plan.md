# Accessibility First-Pass A-to-B Plan

## A — Current Position

- The requested reusable skill is named `accessibility-first-pass` and belongs under `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/`.
- The supplied review target may be a web page, feature, or implementation with some combination of source, rendered behavior, project tooling, and user-provided evidence available.
- The skill must use repository-owned or synthetic review material and public authoritative accessibility guidance.
- W3C guidance establishes that automated tools can assist evaluation but cannot determine accessibility on their own, and that human judgment remains necessary.
- Durable artifacts must use repository-relative paths and exclude secrets, personal or business identities, usernames, home-directory names, and machine-specific checkout paths. The only permitted public project identities are `Eric Wimp` and `ericwimp8`.
- Production workflow skills, canonical supporting skills, CLI code, website code, and other showcase work are outside the edit boundary.

## B — Desired Position

A ready-to-use agent skill conducts a bounded, responsible first-pass accessibility review; connects source, rendered, automated, manual, and authoritative evidence; separates observation from inference and unverified behavior; and delivers a prioritized report that helps a team investigate and remediate barriers without claiming accessibility or conformance proof.

## Path from A to B

1. Define the review sequence and evidence classifications that govern source inspection, rendered inspection, project-native automation, manual checks, and authoritative guidance.
2. Define risk-based prioritization around affected users, task impact, reach, and evidence confidence rather than treating tool severity as the final priority.
3. Define reporting fields for scope, evidence, reproduction or inspection steps, affected users, remediation direction, human or assistive-technology follow-up, and limitations.
4. Package stable accessibility-evidence guidance in one routed reference and a reusable report structure as an output asset.
5. Add portable skill frontmatter and OpenAI Codex interface metadata without making unsupported harness claims.
6. Structurally validate the generated bundle and record criteria requiring behavioral proof.
7. Evaluate the description, every packaged reference, and representative body behavior through the current evaluation workflow, refining only at the semantic owner of retained failures.
8. Audit all retained publication candidates for privacy and record workflow weaknesses without changing production workflow owners.

## C — Completion Criteria

- The skill accepts a supplied page, feature, implementation, or repository-owned fixture and states the exact review scope and unavailable surfaces.
- The skill inspects available source, rendered behavior, and project-native tooling before selecting checks, and records which evidence actually ran or was observed.
- The skill uses authoritative guidance appropriate to each claim and distinguishes normative requirements from informative implementation guidance.
- Automated findings are verified where possible and are never represented as complete accessibility or conformance evidence.
- Manual inspection covers applicable keyboard, focus, structure, names and roles, content alternatives, forms and errors, visual presentation, motion, reflow or zoom, and dynamic-state concerns.
- Findings identify affected users, evidence, reproduction or inspection steps, remediation direction, priority rationale, and follow-up requiring human or assistive-technology testing.
- Findings visibly distinguish observed facts, source-backed inference, and unverified behavior.
- The report begins with scope and limitations, includes prioritized actionable findings, records passed checks only at the tested scope, and identifies material unknowns.
- The skill refuses to claim that a first-pass review, clean automated scan, or limited manual check proves accessibility or WCAG conformance.
- The generated skill passes structural validation and the available governed description, reference, and behavior evaluations, or records the exact stopping gate and retained evidence.
- All retained artifacts remain under `showcase-skills/accessibility-first-pass/` and pass the repository privacy audit.

## Generation Contract

- **Destination:** `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/`
- **Supported surface:** portable Agent Skills content with OpenAI Codex interface metadata; no broader runtime reliability claim.
- **Generation viability:** autonomous.
- **Execution preference:** autonomous continuation through generation and evaluation.
- **Authority to act:** create and refine only artifacts owned by `showcase-skills/accessibility-first-pass/`.
- **Required user stops:** none identified; stop only if a governing evaluation prerequisite or unavailable capability prevents reliable evidence.
- **External dependencies:** public W3C WAI and WCAG guidance; local review tools only when the target project already provides them or the environment safely exposes them.
- **Implementation-time choices:** exact fixture mix, report-template layout, and concise routing language may be decided during generation without changing intent.
- **Evaluation handoff:** continue directly into `skill-evaluation-and-refinement`; retain structural, reference-traversal, behavior, refinement, and stopping evidence under the showcase workspace.

## Authoritative Context

- [W3C WAI: Selecting Web Accessibility Evaluation Tools](https://www.w3.org/WAI/test-evaluate/tools/selecting/)
- [W3C: Web Content Accessibility Guidelines 2.2](https://www.w3.org/TR/WCAG22/)
- [W3C WAI: Easy Checks](https://www.w3.org/WAI/test-evaluate/preliminary/)
- [W3C WAI-ARIA Authoring Practices Guide](https://www.w3.org/WAI/ARIA/apg/)
