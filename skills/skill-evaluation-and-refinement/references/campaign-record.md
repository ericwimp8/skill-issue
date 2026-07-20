# Evaluation Campaign Record

## Required Structure

Keep one directory per target and separate description and behavior loops:

```text
<campaign>/<target>/
  contract.md
  status.md
  description/round-<n>/trial-<n>.md
  references/<reference-path>/round-<n>/trial-<n>.md
  behavior/cycle-<n>/case-<n>.md
  fixtures/
  retained-evidence/
```

Store generated outputs inside the owning round or cycle. Do not let a later trial read a prior output unless the evaluation explicitly tests follow-up behavior.

## Target Status

Record:

- target path and version or content hash;
- environment-qualification record used;
- refinement mode;
- description state as `unvalidated`, `passed`, `not-applicable`, `failed`, or `blocked`, with four trial-record paths for `passed` or an explicit-only reason for `not-applicable`;
- reference state as `not-applicable`, `unvalidated`, `passed`, `failed`, or `blocked`, with one entry per current reference file and the trial-record paths supporting each passed entry;
- current loop, round or cycle, and per-loop failure count;
- total campaign allowance and remaining attempts;
- last completed trial and next action;
- current pass, stopped, or blocked state.

## Trial Evidence

Each trial or case records its unmodified prompt, fresh-agent identity, fixture paths, direct invocation evidence when applicable, observable outputs, ground-truth comparison, criterion result, and cleanup ownership.

Keep conclusions separate from raw evidence. Preserve the evidence needed to reproduce a pass or understand a stop.

## Cleanup

Before rerunning a loop:

1. inventory files and changes owned by the completed round or cycle;
2. retain trial records, direct evidence, audit findings, and the target refinement;
3. remove generated fixtures, outputs, transcripts exposed to later agents, and transient workspace changes;
4. recreate clean fixtures from their recorded specification;
5. record cleanup completion before starting the next agent.

## Resume Gate

Resume only when `status.md` identifies the exact target, loop, count, last evidence, applied refinement, completed cleanup, and next action. When any item is uncertain, reconstruct it from retained evidence before running another trial.

Body work may resume only when `status.md` also satisfies the description-state evidence requirement above. General body cases may begin only when every current reference entry is `passed` or the reference state is `not-applicable` because the target has no `references/` files.
