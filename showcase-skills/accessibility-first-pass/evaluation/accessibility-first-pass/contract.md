# Accessibility First Pass Evaluation Contract

- **Target:** `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/SKILL.md`
- **Goal:** produce a responsible, evidence-bounded first-pass accessibility review that identifies actionable web accessibility risks and required follow-up.
- **Intended use:** review of a supplied web page, feature, or implementation with any combination of production source, rendered access, project tooling, and public authoritative guidance.
- **Expected behavior:** establish scope; trace production owners; exercise available rendered states; combine compatible automation with risk-based manual inspection; separate observed, tool, source-backed, inferred, and unverified evidence; prioritize user impact; and route unresolved human or assistive-technology checks.
- **Expected result:** a prioritized report with affected users, evidence, reproduction or inspection steps, remediation direction, verification routes, authoritative mappings, human or assistive-technology follow-up, and material limitations.
- **Preserved boundaries:** no claim of accessibility, completeness, certification, or standards conformance; no invented observation; no scanner-only conclusion; no assumption that unavailable behavior passes; no use of private or non-permitted fixtures.
- **Evaluation surface:** document output produced from isolated synthetic web fixtures, source inspection, and rendered or automated capability available to each trial.

## Observable Criteria

1. The report defines included journeys, states, viewports, inputs, tooling, and exclusions before broad conclusions.
2. Findings distinguish observed, tool, source-backed, inferred, and unverified evidence without upgrading one level into another.
3. The review uses available project-compatible automation and records the command, scope, state, and result, or records the precise tooling limitation.
4. Manual inspection covers the highest-risk behavior applicable to the target and does not substitute a generic checklist for executed evidence.
5. Every material finding includes affected users, evidence, inspection steps, remediation direction, verification route, and narrowly scoped authoritative guidance.
6. Priority follows user impact, task blockage, reach, recurrence, and evidence confidence rather than scanner severity alone.
7. Human, assistive-technology, and disabled-user follow-up is explicit where the current review cannot establish the experience.
8. The report records material limitations and states that the first pass does not prove accessibility, completeness, certification, or standards conformance.
9. A source-only or capability-constrained review remains useful and gives an actionable external verification route without converting missing evidence into a pass.
