# Generation-to-Evaluation Handoff

## Generated Target

- Skill: `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/SKILL.md`
- Reference: `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/references/web-accessibility-evidence.md`
- Report asset: `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/assets/accessibility-first-pass-report.md`
- OpenAI metadata: `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/agents/openai.yaml`

## Evaluation Contract

- **Goal:** create a bounded, evidence-labeled first-pass accessibility report that prioritizes user impact and actionable remediation without overstating completeness or conformance.
- **Intended use:** preliminary accessibility review of supplied web pages, features, implementations, and representative fixtures.
- **Expected behavior:** investigate available production source, rendered behavior, native tooling, and authoritative guidance; combine applicable automated and manual checks; classify evidence; identify affected users; prioritize findings; and direct unresolved checks to human or assistive-technology testing.
- **Expected result:** a report containing scope, methods, limitations, prioritized findings, passed checks within tested scope, follow-up checks, unknowns, and next actions.
- **Boundaries:** do not claim complete accessibility, certification, conformance, or absence of barriers; do not convert unverified behavior into findings or passes; do not install tooling or broaden review scope without authority.
- **Supported surface:** portable Agent Skills content with OpenAI Codex interface metadata. Runtime reliability outside the evaluated Codex environment is unclaimed.
- **Refinement mode:** automatic, as authorized by the initiating prompt, constrained to the showcase workspace and governed semantic-owner updates.

## Runtime Criteria

Evaluation must prove that an independent agent can:

- discover and apply the evidence reference rather than relying on the core body alone;
- distinguish observed, inferred, and unverified behavior in a real report;
- prioritize by user impact and evidence rather than copying automated severity;
- report relevant manual and assistive-technology follow-up;
- preserve scope and limitations while producing actionable remediation direction; and
- refuse unsupported accessibility or conformance conclusions even when automated checks are clean.

Continue directly into `skill-evaluation-and-refinement` using the intake plan as source of truth.
