# Description Round 1 Evidence

## Initial Trials

- Trial 1 used the unmodified prompt at `trial-1/request.md`. Fresh session `019f7ce2-c958-7651-8b7e-b512bf209be2` read the exact target in a pre-output tool call and produced `trial-1/package.md`. The package reproduced the documented config-key mismatch, retained the 3/3 observation limit, cited environment sources, and produced a standalone issue.
- Trial 2 used the unmodified prompt at `trial-2/request.md`. Fresh session `019f7ce2-dc2f-7853-9827-db1c48d4e48f` read the exact target in a pre-output tool call and produced `trial-2/package.md`. The package reproduced the missing-field crash, separated the reporter's uncertainty from local evidence, and recorded comparison behavior and material gaps.

## Confirmation Trials

- Trial 3 used the distinct unmodified prompt at `trial-3/request.md`. Fresh session `019f7ce4-d0ca-7533-93e7-954174847ece` read the exact target in a pre-output tool call and produced `trial-3/package.md`. The package correctly reported a 3/3 non-reproduction, preserved the original claim separately, and requested only evidence that could distinguish input, invocation, version, output, and environment causes.
- Trial 4 used the distinct unmodified prompt at `trial-4/request.md`. Fresh session `019f7ce4-e6f8-7aa3-8312-85a8677b3db4` read the exact target in a pre-output tool call and produced `trial-4/package.md`. The package respected the production-data boundary, distinguished the retained upstream timeout from the unobserved user-interface hang, and returned a blocked but actionable handoff.

## Decision

Description evaluation passed at 4/4. Every fresh agent selected and read the exact candidate before output, and all four outputs stayed inside the intended incomplete-bug-report boundary. No description refinement is supported by the retained evidence.
