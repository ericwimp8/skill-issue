# Code Implementation Discipline Evaluation Contract

## Target

- Canonical path: `evaluations/scenario-skill-refinement/code-implementation-discipline/skill/`
- Evaluation surface: code implementation and code-change planning.
- Refinement mode: automatic semantic refinements supported by retained evidence.
- Environment qualification: `evaluations/skill-system-production-refinement/environment-qualification.md`.

## Goal

Prevent an agent from choosing the nearest visible edit site before tracing the behavior to the concrete location that owns it.

## Intended Use

Any task that implements or edits code, including bug fixes, feature changes, and implementation planning that stops before editing.

## Expected Behavior

The agent separates the requested outcome from the observation point, traces the producing and consuming path to concrete implementations, identifies the behavior owner and affected paths, selects the smallest complete owner-level approach, and reconciles and verifies the result.

## Expected Result

The resulting code or implementation plan places the change at the behavior owner, resolves the original observation, preserves affected shared paths, and leaves no compensating workaround at a caller, wrapper, relay, or seam.

## Boundary

- Preserve valid local changes when the local location genuinely owns the behavior.
- Preserve explicit user requests to stop before implementation.
- Do not widen the target into a general coding manual or require exhaustive ceremony.
- Do not infer correctness from tests before tracing production source.

## Completion Criteria

1. Four fresh, unleading description trials retain native evidence that the exact candidate loaded before output.
2. Reference qualification is recorded as not applicable because the target has no `references/` files.
3. Varied isolated body cases show the required decision and workflow in observable output or edits.
4. Every body output identifies and changes or proposes the concrete behavior owner rather than a merely visible seam.
5. Every implementation case reconciles affected callers and verifies owner-level and observation-level behavior.
6. Any refinement is generalized, applied at the semantic owner, and followed by clean reruns.
