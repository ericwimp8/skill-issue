---
name: accessibility-first-pass
description: Evidence-bounded first-pass accessibility review for web pages, features, and implementations. Use when investigating accessibility risks and actionable next steps from available source, rendered behavior, project tooling, and authoritative guidance.
---

# Accessibility First Pass

## Establish the Review Boundary

- Record the target, included user journeys, states, routes, breakpoints, inputs, and exclusions before testing.
- Inventory available production source, runnable or rendered access, browser capability, project accessibility tooling, and relevant project commands. Trace behavior through concrete production owners; use tests only as validation aids.
- Identify the people who may be affected by the reviewed behavior. Tie user impact to evidence rather than treating a disability-category list as proof of lived experience.
- State the standard or guidance baseline only when the request or project establishes one. Read `references/web-accessibility-evidence.md` before mapping findings to WCAG or making coverage statements.

## Gather Evidence

1. Inspect the production implementation and configuration that create the target behavior. Record source-backed facts separately from expectations or comments.
2. Exercise the rendered surface when available. Cover the included initial and changed states, meaningful interactions, viewport conditions, and failure paths before calling behavior observed.
3. Run compatible existing automation against every exercised state it can inspect. Preserve the command, configuration or rule scope, target state, and complete violation and incomplete counts. Treat unavailable or failed tooling as a limitation.
4. Manually inspect the highest-risk behavior that automation cannot establish, including perception, structure, names and instructions, keyboard and focus behavior, reflow or zoom, motion, errors, and dynamic updates as applicable to the target.
5. Route checks that require a qualified human, assistive technology, or disabled-user evaluation to follow-up. Never simulate certainty from markup, an accessibility tree, or generic model knowledge.

Label evidence consistently:

- **Observed:** directly reproduced in the rendered surface during this review.
- **Tool:** emitted by an identified automated check for a recorded state.
- **Source-backed:** established by traced production source but not rendered execution.
- **Inference:** reasoned risk supported by stated evidence and still requiring verification.
- **Unverified:** material behavior the available environment or scope could not test.

## Analyze and Prioritize

- Consolidate duplicate symptoms under the behavior owner that should be remediated.
- Prioritize by user impact, task blockage, reach, recurrence, and evidence confidence. Do not inflate priority from a scanner severity or a possible WCAG mapping alone.
- Map authoritative guidance narrowly. Identify the relevant criterion or resource and explain the relationship; mark uncertain mappings as interpretation requiring specialist review.
- Give remediation direction at the owning source or interaction design. Preserve multiple valid implementation choices unless one concrete correction is required by the evidence.
- Keep automated passes, manual passes, and absence of observed failures scoped to the exact states and checks executed.

## Report

Use `assets/accessibility-first-pass-report.md` as the output structure. For every material finding include affected users, evidence label and artifact, reproduction or inspection steps, remediation direction, verification route, and applicable authoritative guidance.

End with:

- checks requiring human, assistive-technology, or disabled-user testing;
- untested routes, states, viewports, environments, content, and technologies;
- tool failures, unavailable inputs, and evidence-quality limits;
- an explicit conclusion that the first pass does not prove accessibility, completeness, certification, or standards conformance.

If rendered execution or automation is unavailable, produce a source-bounded report and an actionable external verification route. Never convert missing evidence into a pass.

## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/web-accessibility-evidence.md`: Authoritative web accessibility sources and evidence boundaries. Use when selecting guidance, mapping findings, or explaining the limits of automated and preliminary review.

## Assets

- `assets/accessibility-first-pass-report.md`: Evidence-bounded report structure. Use for every completed review.
