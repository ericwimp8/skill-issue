# Intake Plan Contract

## Readiness Standard

The plan is ready when another agent can create the intended skill without rediscovering user intent or known project context.

Confirm that it establishes:

- the user-visible purpose and expected outcome;
- the real project state, inputs, constraints, and destination;
- the intended invocation boundary;
- broad dependency-ordered creation tasks;
- observable completion criteria;
- material exclusions and preserved behavior;
- the target harness surface where known and whether Codex metadata is required;
- unresolved implementation choices that generation may decide;
- unresolved user-owned decisions that require a stop;
- the generation-viability assessment, the user's execution preference, and the authority boundary.

## Question Boundary

Ask the user only when all three conditions hold:

1. available project and authoritative platform context does not answer the question;
2. different answers would materially change the intended skill;
3. generation cannot safely treat the choice as an implementation detail.

Combine closely related gaps into one focused question. Integrate the answer into the semantic section that owns it.

## Handoff Shape

Deliver one coherent A-to-B plan followed by a compact generation contract containing:

- destination;
- generation viability;
- selected execution preference;
- authority to act;
- required user stops;
- external dependencies or unavailable inputs;
- expected handoff into evaluation and refinement.

Do not attach a second shadow specification. The plan is the source of truth.

The Generation contract is the single owner of generation viability, execution preference, and intake-to-generation authority stops. Keep those orchestration fields out of A, B, the construction path, and C.
