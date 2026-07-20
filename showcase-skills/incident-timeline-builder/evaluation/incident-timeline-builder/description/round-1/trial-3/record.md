# Trial Record

## Identity

- Trial identity: `80920012-2960-4919-A59F-043DCBBEB647`
- Execution: fresh independent description trial 3
- Observable result: produced an evidence-led incident package with two resolved UTC observations, one repeated-hour change report bounded to two possible UTC instants, two additional visible unresolved note records, and explicit causal limitations.

## Selection Basis

The candidate was selected because its frontmatter says to use it when logs, alerts, deployments, or operator notes must be reconciled into a sourced incident chronology. The request asks for a trustworthy incident timeline from a system log and on-call notes around a daylight-saving transition, so the advertised scope directly applies.

## Exact Evidence Paths

- Candidate: `showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- Request: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-3/request.md`
- Fixture directory: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/`
- Fixture: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/on-call.md`
- Fixture: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/system.log`

## Source Preservation

| Source | Before SHA-256 | After SHA-256 | Result |
| --- | --- | --- | --- |
| `on-call.md` | `635224c544727055c165a0f288a10c21683a7472f45a5312e216784dcdc6d2ba` | `635224c544727055c165a0f288a10c21683a7472f45a5312e216784dcdc6d2ba` | Preserved |
| `system.log` | `88327e7bd50f3d0be338d050b3faffcd5305b11f32e3ad1b8e5149846b1baf2c` | `88327e7bd50f3d0be338d050b3faffcd5305b11f32e3ad1b8e5149846b1baf2c` | Preserved |

`git diff -- showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3 showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-3/request.md` produced no output.

## Commands and Outputs

The exact request, candidate-frontmatter, and complete-candidate read commands and outputs are retained in `native-evidence.log`.

### Inventory and before hashes

```text
$ find showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3 -type f -print | sort
showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/on-call.md
showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/system.log

$ find showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3 -type f -print0 | sort -z | xargs -0 shasum -a 256
635224c544727055c165a0f288a10c21683a7472f45a5312e216784dcdc6d2ba  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/on-call.md
88327e7bd50f3d0be338d050b3faffcd5305b11f32e3ad1b8e5149846b1baf2c  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/system.log
```

### Derived timeline

```text
$ python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-3/records.jsonl --output showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-3/timeline.json
[exit 0; no stdout or stderr]
```

The repeated-hour candidates were reproduced from the named zone:

```text
fold=0 offset=-1 day, 20:00:00 utc=2026-11-01T05:30:00Z
fold=1 offset=-1 day, 19:00:00 utc=2026-11-01T06:30:00Z
```

### JSON validation

The validation parsed every nonblank JSONL line with `json.loads`, parsed `timeline.json`, checked the record and resolved/unresolved counts, compared `timeline.json`'s input hash to the SHA-256 of `records.jsonl`, checked deterministic UTC order, and verified the ambiguous timestamp status.

```text
PASS: 5 JSONL records parsed; timeline JSON parsed; 2 resolved and 3 unresolved events; input hash and deterministic UTC order verified.
```

### After hashes

```text
$ find showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3 -type f -print0 | sort -z | xargs -0 shasum -a 256
635224c544727055c165a0f288a10c21683a7472f45a5312e216784dcdc6d2ba  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/on-call.md
88327e7bd50f3d0be338d050b3faffcd5305b11f32e3ad1b8e5149846b1baf2c  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-3/system.log
```

## Produced Artifacts

- `package.md`: investigator-facing incident record.
- `records.jsonl`: source-derived event ledger.
- `timeline.json`: deterministic normalization output.
- `native-evidence.log`: pre-output candidate applicability and full-read evidence.
- `record.md`: this reproducibility and validation record.
