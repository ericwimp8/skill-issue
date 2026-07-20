# Evaluation Record

## User Task

Use index.html and scan-results.json to prepare a first-pass accessibility report for the service dashboard. Explain why each item has its assigned priority, distinguish what the scan proves from what it does not, and identify affected users and follow-up testing.

## Fixtures

- `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dashboard/index.html`
- `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/fixtures/dashboard/scan-results.json`

## Evaluation Evidence

- **Evidence log:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/references/web-accessibility-evidence/round-1/trial-1/native-evidence.log`
- **Observable output:** `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/references/web-accessibility-evidence/round-1/trial-1/output.md`
- **Reference-owned ground truth used:** Priority was independently assigned from task impact, reach, workaround burden, and confidence rather than copied from tool severity. The color-only service state and heading-structure risk were evaluated as separate user impacts. Observation, inference, unverified behavior, affected users, exact follow-up, and the first-pass boundary were required and audited.

## Criterion Audit

| Criterion                                 | Result | Evidence                                                                                                                                                                                                                         |
| ----------------------------------------- | ------ | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Uses both named fixtures                  | PASS   | `output.md` identifies and analyzes the HTML/CSS and both reported scan results.                                                                                                                                                 |
| Explains each assigned priority           | PASS   | Color-only status is High because it affects every row and the dashboard's primary task with no direct cue; heading semantics are Medium because navigation is degraded but linear reading remains available on this small view. |
| Does not inherit tool severity            | PASS   | The scan's `minor` and `serious` labels are reported as tool output, while report priority reverses their order based on user impact and confidence.                                                                             |
| Distinguishes scan proof and limitations  | PASS   | Methods and limitations state that the scan proves only what `synthetic-checker` reported and does not prove impact, remediation, completeness, or conformance.                                                                  |
| Distinguishes evidence classes            | PASS   | Each finding separates observed source/tool evidence from inferred runtime effects and enumerates unverified behavior.                                                                                                           |
| Identifies affected users and task impact | PASS   | Findings name color-perception, low-vision, screen-reader, braille, and heading-navigation impacts without asserting universal experience.                                                                                       |
| Provides exact follow-up                  | PASS   | Findings specify browsers, assistive technologies, states, journeys, and expected outcomes; unknowns include keyboard, high contrast, zoom/reflow, and dynamic updates.                                                          |
| Uses authoritative public guidance        | PASS   | Material mappings link exact W3C WCAG 2.2 criteria and Understanding documents; W3C evaluation-tool guidance supports the automation boundary.                                                                                   |
| Preserves first-pass boundary             | PASS   | Scope, limitations, unknowns, and the final boundary avoid completeness, certification, and conformance claims.                                                                                                                  |
| Uses publication-safe durable paths       | PASS   | Report and record use repository-relative paths and public W3C URLs; no private identity or absolute filesystem path appears.                                                                                                    |

## Cleanup Ownership

- This trial owns only `native-evidence.log`, `output.md`, and `record.md` in `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/references/web-accessibility-evidence/round-1/trial-1/`.
- These evaluation artifacts may be removed together when the round-1 trial evidence is no longer needed.

## Result

**PASS**
