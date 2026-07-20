# Accessibility First Pass A-to-B Plan

## A — Current Position

- A supplied web page, feature, or implementation may include source files, a runnable or already-rendered surface, project-specific tooling, and incomplete contextual information.
- The reviewer can inspect available production source, rendered states, interaction behavior, project commands and dependencies, and public authoritative accessibility guidance.
- Automated accessibility tools find only some classes of issues. Manual inspection, assistive-technology testing, and testing with disabled users remain necessary for many barriers and for reliable conformance claims.
- Review inputs may be partial, unrunnable, authentication-gated, dynamic, responsive, stateful, or dependent on capabilities unavailable in the current environment.
- The finished skill must preserve evidence boundaries: observed behavior, source-backed findings, tool output, inference, and unverified behavior must remain distinguishable.
- Review artifacts must use synthetic or repository-owned inputs and public sources. Durable public artifacts must use repository-relative paths and contain no secrets, personal or business identities, usernames, home-directory names, or machine-specific checkout paths.
- The intended destination is `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/`.
- The canonical payload targets portable Agent Skills content and OpenAI Codex project or user skill delivery without requiring host-specific metadata.

## B — Desired Position

A ready-to-use `accessibility-first-pass` agent skill guides a responsible, bounded first-pass review of a web surface. It investigates the available implementation and rendered experience, selects project-compatible automated checks, performs targeted manual inspection, relates findings to authoritative guidance, and produces a prioritized actionable report without representing limited evidence as proof of accessibility or standards conformance.

## Path From A to B

1. Establish a concise review workflow that records scope, tested states, available evidence, affected user needs, and material environmental limitations before drawing conclusions.
2. Direct the reviewer to trace relevant production source and project tooling, exercise the rendered surface when available, and use authoritative accessibility guidance appropriate to the review.
3. Define evidence handling that separates directly observed results, tool findings, source-backed facts, reasoned inference, and behavior that remains unverified.
4. Combine compatible automated checks with manual inspection of high-impact interaction, perception, structure, content, responsive, and state-change behavior without treating either method as exhaustive.
5. Require explicit routing of checks that need keyboard-only, zoom/reflow, high-contrast, reduced-motion, screen-reader, speech-input, switch, cognitive, or disabled-user testing to qualified human follow-up when the current environment cannot establish them.
6. Create a report contract that prioritizes actionable findings and records scope, affected users, evidence, reproduction or inspection steps, remediation direction, follow-up testing, and limitations.
7. Keep WCAG or other standards references accurately scoped to the evidence and prohibit conformance, certification, completeness, and accessibility claims from a first-pass review.
8. Validate the skill structure, reference closure, authoring discipline, completion criteria, and privacy constraints, then hand it to a governed evaluation campaign.
9. Evaluate representative source-only, rendered-interaction, and constrained-environment cases where qualified harness evidence is available, refining only generalized meaning that materially fails the target contract.

## C — Completion Criteria

- The skill begins by defining the review target, included states and viewports, available source and rendered access, relevant project tooling, known user journeys, and exclusions.
- The skill inspects production source through concrete behavior owners rather than relying on tests, screenshots, or declarations as behavioral truth.
- The skill exercises available rendered states and interactions, including meaningful state changes, before reporting them as observed.
- The skill uses appropriate project-compatible automated checks and preserves their raw or summarized evidence without treating zero reported violations as proof of accessibility.
- The skill performs manual inspection targeted to the reviewed surface and identifies affected users without claiming that a generic checklist represents lived experience.
- Every material finding distinguishes evidence from inference and includes priority, affected users, evidence, reproduction or inspection steps, and remediation direction.
- The report separately records checks requiring human, assistive-technology, or disabled-user testing and explains why they remain open.
- The report states material limitations, unreviewed states, unavailable capabilities, and the exact evidence level established.
- References to WCAG or other guidance identify the applicable source and avoid claiming full standards coverage or conformance from partial checks.
- The skill remains useful when rendered execution or a preferred scanner is unavailable by producing a source-bounded review and an actionable external verification route.
- The skill folder contains only the canonical skill document and genuinely required supporting resources, with valid relative references and no public-repository privacy violations.
- Structural validation passes, and behavior evaluation either passes every governed case or stops with the exact retained gate and evidence.

## Generation Contract

- **Destination:** `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/`
- **Generation viability:** Autonomous. The intended behavior, evidence boundaries, destination, authoritative public sources, repository conventions, and evaluation prerequisites are available without an expected user-owned design decision.
- **Execution preference:** Autonomous continuation through generation and evaluation/refinement, as explicitly authorized by the user.
- **Authority to act:** Create and refine only the new `showcase-skills/accessibility-first-pass/` workspace and artifacts governed by this task.
- **Required user stops:** Stop only if a workflow gate requires a user-owned choice, the qualified harness cannot produce required evidence, or safe evaluation would require unavailable credentials, private data, or external actions.
- **External dependencies or unavailable inputs:** Live rendered review depends on a runnable fixture and browser capability; assistive-technology and disabled-user findings cannot be inferred when those testers or technologies are unavailable.
- **Unresolved implementation choices:** Generation may decide the smallest useful split between the canonical skill body, conditional reference material, and an output template; it may select repository-available validators and synthetic fixtures.
- **Expected handoff:** Continue directly to `skill-generation`, then to `skill-evaluation-and-refinement` in automatic-refinement mode when its qualification and evidence gates are satisfied.
