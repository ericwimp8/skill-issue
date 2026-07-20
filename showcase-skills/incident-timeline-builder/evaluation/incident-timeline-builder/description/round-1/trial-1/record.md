# Trial Record

## Identity

- Trial identity: `incident-timeline-builder-description-round-1-trial-1`
- Execution identity: fresh independent description trial 1, GPT-5.6 Sol, medium reasoning
- Completed at: `2026-07-20T18:29:18Z`
- Working directory: `<repo-root>`

## Selection Basis

The candidate was selected because its frontmatter says it applies when logs, alerts, deployments, or operator notes must be reconciled into a sourced incident chronology. The request asks for exactly that reconciliation across operational records, with chronology, provenance, uncertainty, and source preservation.

## Exact Inputs

- Candidate: `showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- Request: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-1/request.md`
- Fixture directory: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/`
- Fixture: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/alert.json`
- Fixture: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/deploy.csv`
- Fixture: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/notes.md`
- Fixture: `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/service.log`

The request SHA-256 was `82628b87806750bec5988d9e2bca002d7aad02d388133349963f1d5a51bb3208`. The candidate SHA-256 was `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1`.

## Native Read Evidence

`native-evidence.log` contains the exact commands and complete terminal output for the pre-output reads, in order:

1. `pwd`
2. `cat showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-1/request.md`
3. `sed -n '1,/^---$/p' showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
4. `cat showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`

The request and candidate frontmatter were read before applicability was decided. The complete candidate was then read before any fixture inspection.

## Source Hash Preservation

| Fixture | SHA-256 before | SHA-256 after | Result |
| --- | --- | --- | --- |
| `alert.json` | `08df777eacff74853175ffe8837104e41eab467ac69bc1a68f91446adef9cce5` | `08df777eacff74853175ffe8837104e41eab467ac69bc1a68f91446adef9cce5` | Preserved |
| `deploy.csv` | `e82e6d6977f91c1e00aa6ace05f4ec57a1584037d3d9dc5ab39d151a6071881f` | `e82e6d6977f91c1e00aa6ace05f4ec57a1584037d3d9dc5ab39d151a6071881f` | Preserved |
| `notes.md` | `7d9663c985ef28d7f49724b1f9fa10c51d31a041573304fde18e19aca737fb7d` | `7d9663c985ef28d7f49724b1f9fa10c51d31a041573304fde18e19aca737fb7d` | Preserved |
| `service.log` | `b470eb018b4cd97e31d23f3acc49b891851a3eb535011773c42924c6dab15139` | `b470eb018b4cd97e31d23f3acc49b891851a3eb535011773c42924c6dab15139` | Preserved |

## Command and Output Ledger

### Evidence inventory and first hashes

Command:

```sh
find "$fixture" -type f -print | sort
find "$fixture" -type f -print0 | sort -z | xargs -0 shasum -a 256
while IFS= read -r file; do printf '\n===== %s =====\n' "$file"; cat "$file"; done < <(find "$fixture" -type f -print | sort)
```

Output: four fixture paths were listed; the before hashes are recorded in the preservation table; all four file contents were emitted and used to construct the evidence ledger.

### Normalizer contract

Command:

```sh
python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py --help
```

Output:

```text
usage: build_timeline.py [-h] [--output OUTPUT] input

Normalize explicit timestamps and stably order incident JSONL records.
```

The remaining help output required UTF-8 JSONL fields `source_id`, `source_path`, `locator`, `classification`, and `summary`, with optional `raw_timestamp` and `source_timezone`, plus a new-file `--output` path.

### First normalization attempt

Command:

```sh
python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-1/records.jsonl --output showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-1/timeline.json
```

Output:

```text
error: line 1: classification must be one of contradiction, gap, inferred, observed, reported
```

The derived records were corrected to evidence-class labels without changing source evidence.

### Successful normalization

Command: the same normalizer command shown above.

Output: exit status 0; `timeline.json` was created with five resolved events in UTC order and two unresolved reported events. Its embedded input hash is `184093e36453bbca43b62101c166266779c6e23ea72384c90a8c17a37bf73f82`, matching `records.jsonl`.

### JSON and preservation validation

Commands:

```sh
python3 -c 'import json, pathlib; p=pathlib.Path("showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-1/records.jsonl"); rows=[json.loads(line) for line in p.read_text().splitlines()]; assert len(rows)==7; print(f"records.jsonl: valid ({len(rows)} records)")'
python3 -m json.tool showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-1/timeline.json >/dev/null
printf '<recorded fixture hashes>' | shasum -a 256 -c -
```

Output:

```text
records.jsonl: valid (7 records)
timeline.json: valid JSON
showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/alert.json: OK
showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/deploy.csv: OK
showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/notes.md: OK
showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-1/service.log: OK
```

## Output Hashes

- `records.jsonl`: `184093e36453bbca43b62101c166266779c6e23ea72384c90a8c17a37bf73f82`
- `timeline.json`: `1ef307bc7bc2c695033b77c72f1bf80d0417e7ded9d3d87c99eed604101ca120`
- `package.md`: `737a2fbd2525dce51a4914e7a8496b8264a6cc85eafd4f0b4caf7214118ce9db`
- `native-evidence.log`: `7f8703213c297fff1dfb49403335c7477bdc7c97ee883c62b1d5a7bcf9d01fe5`

## Observable Result

The trial produced a handoff package with an evidence inventory, explicit UTC normalization, a five-event sourced chronology, two unresolved records kept unplaced, bounded hypotheses, stated absence of exact contradictions, material evidence gaps, and prioritized follow-up actions. The accompanying JSONL and JSON preserve seven extracted records and deterministic normalization. All fixture hashes match before and after execution.
