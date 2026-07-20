# Behavior Cycle 1 Audit

## Result

Passed all three isolated body cases on the qualified fresh GPT-5.6 Sol medium surface. No material contract failure was retained, so no semantic refinement or post-refinement rerun was required.

| Case | Stress surface | Evidence | Result |
| --- | --- | --- | --- |
| 1 | Offset normalization, deployment proximity, missing rollback record, and alternative explanations | `case-1/record.md`, `case-1/native-evidence.log`, `case-1/package.md`, `case-1/timeline.json` | Passed |
| 2 | Exact tie, repeated-hour ambiguity, relative untimed ordering, and unsupported recovery belief | `case-2/record.md`, `case-2/native-evidence.log`, `case-2/package.md`, `case-2/timeline.json` | Passed |
| 3 | Partial status-page times, untimed credential rotation, shareable output, and non-secret follow-up | `case-3/record.md`, `case-3/native-evidence.log`, `case-3/package.md`, `case-3/timeline.json` | Passed |

## Criterion Audit

1. **Immutable evidence:** passed. Every record reports matching before/after SHA-256 values for all fixtures.
2. **Provenance:** passed. Material rows retain stable source identifiers and line, row, or field locators.
3. **Zone normalization:** passed. Explicit offsets normalize to the expected UTC instants while raw values remain in the ledger.
4. **Unresolved time:** passed. A date-less Adelaide clock value, repeated London local hour, date-less status times, and absent credential time remain unresolved.
5. **Deterministic ties:** passed. Case 2 retains the two `00:15:00Z` events in input order and explicitly denies finer sequence.
6. **Epistemic classes:** passed. Machine records remain observed, operator claims reported, missing records gaps, and hypotheses bounded in prose.
7. **Causal restraint:** passed. Deployment, restart, and credential-rotation proximity are framed as investigative hypotheses with alternatives and missing proof.
8. **Contradictions:** passed. Case 1 preserves reported rollback versus absent completion evidence without forcing conflict; exact disagreements and resolution evidence remain explicit.
9. **Deliverable:** passed. All packages contain the seven required sections with prioritized follow-up tied to gaps.
10. **Helper traceability:** passed. Every case retains parsed JSONL, helper output, matching `input_sha256`, valid resolved ordering, and unresolved partitions.

## Refinement Decision

- Material failures: 0.
- Body failure count: 0.
- Target refinement: none.
- Target hash remains `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1`.
