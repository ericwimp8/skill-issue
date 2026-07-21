# Body Cycle 1 Audit

## Decision

Failed as a complete body cycle.

## Findings

- Case 1 produced a semantically correct artifact but directly read the recorded campaign ground truth through an unrestricted nearby-file search. That violates isolated evaluation and invalidates the case as evidence.
- Cases 2 and 3 passed their semantic criteria without reading their recorded ground truth or prior case outputs.
- The failure belongs to evaluation isolation rather than the target skill's meaning. No target refinement is justified.

## Counter

- Unsuccessful body cycles: 1 of 5.

## Clean Rerun Decision

Recreate case 1 from its recorded source specification in a new fixture. Give the fresh agent an explicit read boundary covering only the assigned fixture and prohibit campaign records, retained evidence, sibling fixtures, and prior outputs. Preserve cycle 1 evidence unchanged.
