---
name: systematic-debugging
description: Systematic root-cause debugging process. Use when encountering any bug, test failure, or unexpected behavior, before proposing fixes.
---

# Systematic Debugging

## Root-Cause Gate

Do not propose or apply a fix until the evidence supports a root cause.

Before intervening, establish:

- the observed failure and the behavior that should hold;
- a reliable reproduction or the missing evidence needed to obtain one;
- the concrete path that produces the failure;
- where the incorrect condition first appears;
- why that condition produces every relevant symptom;
- the smallest correction at the causal owner;
- the focused check that will prove the correction.

## Workflow

1. Read the complete error or report and reproduce the failure without changing
   the system.
2. Trace the failing value, state, or control flow backward through concrete
   implementations. Separate the observation point from the causal origin.
3. Compare the failing path with the nearest working path and test one explicit
   causal hypothesis against the evidence.
4. Add the smallest failing reproduction when an automated check can capture
   the behavior.
5. Correct the causal owner, then rerun the focused reproduction and the
   nearest related checks.
6. Stop and gather more evidence when the failure is intermittent, the proposed
   cause explains only one symptom, or the intervention merely hides the
   observation.
