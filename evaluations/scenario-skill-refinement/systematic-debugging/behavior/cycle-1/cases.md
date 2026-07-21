# Behavior Cycle 1 Cases

## Case 1: Shared Cache-Key Cause

Prompt:

> Read and apply the advertised systematic debugging skill. Work only in the assigned case-1 fixture. Two symptoms are reported: semantically identical query objects can miss the cache, and invalidating one ordering can leave the other ordering present. Establish intended behavior from production source, reproduce both symptoms, trace the concrete path to their first shared incorrect condition, and apply the smallest supported correction at its owner. Preserve pre-correction evidence, then run the focused reproduction and nearest related checks. Write a concise report with changed files and commands.

Ground truth and completion characteristics:

- The source contract says query property insertion order does not affect cache identity.
- Both `getCached` and `invalidate` depend on `cacheKey`; downstream Map operations only observe the key mismatch.
- `JSON.stringify(query)` preserves insertion order, producing different keys for equivalent objects.
- The causal correction belongs in `cacheKey`, with deterministic key ordering.
- The focused reproduction fails before the correction and passes after it.
- The related check preserves distinct values and path separation.

## Case 2: Insufficient Intermittent Evidence

Prompt:

> Read and apply the advertised systematic debugging skill. Work only from the assigned case-2 incident record and source snapshot. The report describes an intermittent duplicate notification but contains no captured failing run. Produce the next diagnostic action in a concise report. You may run non-mutating inspections. Do not change source unless the available evidence supports a root cause that explains every reported symptom.

Ground truth and completion characteristics:

- The incident contains competing timer-retry and subscriber-duplication hypotheses.
- No reproduction, event identity, subscriber count, retry count, or timestamp sequence distinguishes them.
- Correct behavior stops before proposing or applying a source correction.
- The report identifies the smallest missing evidence needed to distinguish the concrete paths.
- A speculative fix to either timer or subscription code is a material failure.

## Case 3: Repository Boundary Cause

Prompt:

> Read and apply the advertised systematic debugging skill. Work only in the assigned case-3 fixture. A profile loaded from disk renders `Welcome, undefined`, while the same profile supplied by the in-memory repository renders the expected name. Reproduce the failure, compare the working and failing concrete paths, identify where the incorrect condition first appears, and apply the smallest supported correction at the causal owner. Preserve pre-correction evidence, then run the focused reproduction and nearest related checks. Write a concise report with changed files and commands.

Ground truth and completion characteristics:

- Both repository implementations promise a profile object.
- The disk repository returns raw JSON text while the in-memory repository returns an object.
- The service and view correctly consume the promised object and only observe the disk boundary violation.
- The correction belongs in the disk repository's deserialization boundary.
- The focused reproduction and checks prove both repositories preserve the profile contract.

## Isolation

Each case has its own fixture and output directory. Fresh agents receive only their case prompt and fixture. Reports, native traces, audits, and command evidence remain case-owned.
