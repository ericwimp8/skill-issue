# Skill Intake Cycle 2 Audit

## Result

Material failure retained.

- Case 1 passed.
- Case 2 passed the corrected distinction: Generation remains `Autonomous` while `Step-by-step` is recorded separately as the user's execution preference.
- Case 3 keeps the pause out of the construction path, but both case 2 and case 3 duplicate execution preference and authority-stop state into section A as well as the Generation contract. The Generation contract is not yet the single semantic owner.

## Diagnosis

“Keep ... in the Generation contract” did not explicitly prohibit repeating intake-to-generation control state inside A. Agents treated the current pause and selected preference as current-position facts even though those meanings belong to the handoff contract.

## Generalized Refinement

Make the Generation contract the only owner of viability, execution preference, and intake-to-generation authority stops. Keep intrinsic external constraints on the requested skill in A, while excluding orchestration state from A, B, Path, and C.
