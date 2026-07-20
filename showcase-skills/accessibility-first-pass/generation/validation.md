# Generation Validation

## Structural Result

- **Validator:** `skill-creator/scripts/quick_validate.py`
- **Target:** `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/`
- **Result:** PASS — `Skill is valid!`
- **Folder and frontmatter name:** `accessibility-first-pass`
- **Resources:** one routed reference, one consumed report asset, and OpenAI interface metadata; no unnecessary scripts or auxiliary documentation.
- **Reference closure:** all backtick-marked local resource paths from `SKILL.md` resolve inside the skill directory.

## Generated Content Hashes

| File | SHA-256 |
| --- | --- |
| `SKILL.md` | `c2cd6a758ce1c8de3cd5c10d2026d029c3248e29d0e3d89a6cfe65ebd2d49d8e` |
| `references/web-accessibility-evidence.md` | `c52f41f96a138d9cf8d891146a4482dd367a93d82c8fb43da83ee600575c03da` |
| `assets/accessibility-first-pass-report.md` | `9ccc238fa29f9ec076ddc17817faba03b2143b4eddafcadbcbceb69d72db4ac3` |
| `agents/openai.yaml` | `6b37dc960542529d0225762ed6255f334175ef04813d00abf497301178f0227c` |

## Intake Criteria Walkthrough

- PASS: accepts page, feature, implementation, and fixture review targets while requiring exact scope and unavailable surfaces.
- PASS: indexes the packaged evidence reference under `## Reference Documents` with a concise what-it-is and when-to-use selection rule; no inline instruction makes every review read it by default.
- PASS: traces production source, rendered behavior, native tooling, and authoritative guidance before claims.
- PASS: separates observed, inferred, and unverified behavior.
- PASS: combines applicable automated and manual checks without treating a scan as conformance evidence.
- PASS: connects findings to affected users, inspection steps, remediation direction, priority rationale, follow-up, and limitations.
- PASS: requires human or assistive-technology follow-up for behavior the environment cannot establish.
- PASS: uses a report asset that preserves scope, methods, findings, tested passes, unknowns, next actions, and the first-pass boundary.
- RUNTIME PROOF REQUIRED: description selection, reference traversal and use, evidence classification, priority behavior, report completeness, and refusal of unsupported conformance claims.

Structural validation is generation evidence only. Runtime claims are governed by `showcase-skills/accessibility-first-pass/evaluation/accessibility-first-pass/`.

## Post-Refinement Revalidation

- PASS: the refined target keeps the mandatory `## Reference Documents` index and removes inline default-read routing.
- PASS: `quick_validate.py` reports `Skill is valid!` for the refined bundle.
- PASS: post-refinement reference verification and confirmation use the refined target hash recorded above.
