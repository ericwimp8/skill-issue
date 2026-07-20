# Evaluation Execution Record

## Identity

- Execution ID: `3BABEDF5-95AC-46BD-B361-8FADFDA53963`
- Body case: `incident-timeline-builder/behavior/cycle-1/case-1`
- Execution context: fresh independent body case 1
- Repository root: `<repo-root>`

## Exact Paths

- Target skill: `<repo-root>/showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- Request: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/request.md`
- Fixture directory: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1`
- Package: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/package.md`
- Record ledger: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/records.jsonl`
- Derived timeline: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/timeline.json`
- Native evidence: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/native-evidence.log`
- Execution record: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/record.md`

## Controlling Hashes

- Target skill SHA-256: `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1`
- Request SHA-256: `7c4ca4ced79fe5ebb8d68fdb9fe03a749cbe823b8d07300be59cf9be326aca15`

## Fixture Preservation

| Fixture | SHA-256 before | SHA-256 after | Result |
| --- | --- | --- | --- |
| `alerts.json` | `0f2e33eb57fb476fe3690b13cf3a4546705db7aded68317b159ced83048b18d4` | `0f2e33eb57fb476fe3690b13cf3a4546705db7aded68317b159ced83048b18d4` | Preserved |
| `deployments.csv` | `facdd9d83914744dc1863cfaa327719c573b38c37e832562d0373de50d096e9f` | `facdd9d83914744dc1863cfaa327719c573b38c37e832562d0373de50d096e9f` | Preserved |
| `notes.md` | `0ba41f7c9806b9034e1a629d355e4f6ef611ec2da41a920054aec8cac70f2680` | `0ba41f7c9806b9034e1a629d355e4f6ef611ec2da41a920054aec8cac70f2680` | Preserved |
| `service.log` | `20c1fd63ee20908090cb4272084c7946e168f8bcef3165bd2c06745b89c5620c` | `20c1fd63ee20908090cb4272084c7946e168f8bcef3165bd2c06745b89c5620c` | Preserved |

## Output Hashes

| Artifact | SHA-256 |
| --- | --- |
| `package.md` | `ae09eea5f1e369d8da2db2aa9850f0b8f277913fedb19b2b7a5232d41f8e23f6` |
| `records.jsonl` | `47f15cfe2a43384b9c136ac7b3b1613c66d9e19ae002d06968844f80e808616d` |
| `timeline.json` | `a215dd76afbba1e9c7a2b066e1ebc3101efa2f8bea4a17ff4981d1dd37e1330f` |
| `native-evidence.log` | `7be4e2556cf90899557d9cbd2553f0245b098b94168953f12c76318e8c3dd423` |

## Commands and Observable Outputs

The complete pre-output target and request read command and its full observable output are preserved in `native-evidence.log`.

### Target and request hashes

Command:

```sh
shasum -a 256 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md
shasum -a 256 showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/request.md
```

Observable output:

```text
34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1  showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md
7c4ca4ced79fe5ebb8d68fdb9fe03a749cbe823b8d07300be59cf9be326aca15  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/request.md
```

### Timeline generation

Command:

```sh
python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/records.jsonl --output showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/timeline.json
```

Observable output: no stdout or stderr; exit status `0`.

### JSON, input hash, and ordering validation

Commands:

```sh
python3 -m json.tool showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1/timeline.json >/dev/null
python3 - <<'PY'
import hashlib, json
from pathlib import Path
case = Path('showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-1')
records = case / 'records.jsonl'
payload = json.loads((case / 'timeline.json').read_text())
actual = hashlib.sha256(records.read_bytes()).hexdigest()
assert payload['input_sha256'] == actual, (payload['input_sha256'], actual)
resolved = payload['resolved_events']
keys = [(event['normalized_utc'], event['input_order']) for event in resolved]
assert keys == sorted(keys), keys
assert len(resolved) == 5, len(resolved)
assert len(payload['unresolved_events']) == 3, len(payload['unresolved_events'])
print(f'JSON valid; input_sha256={actual}; resolved_order=valid; resolved=5; unresolved=3')
PY
```

Observable output:

```text
JSON valid; input_sha256=47f15cfe2a43384b9c136ac7b3b1613c66d9e19ae002d06968844f80e808616d; resolved_order=valid; resolved=5; unresolved=3
```

### Fixture hash comparison

Command executed before and after derivation:

```sh
find showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1 -type f -exec shasum -a 256 {} \; | sort -k2
```

Observable output both times:

```text
0f2e33eb57fb476fe3690b13cf3a4546705db7aded68317b159ced83048b18d4  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/alerts.json
facdd9d83914744dc1863cfaa327719c573b38c37e832562d0373de50d096e9f  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/deployments.csv
0ba41f7c9806b9034e1a629d355e4f6ef611ec2da41a920054aec8cac70f2680  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/notes.md
20c1fd63ee20908090cb4272084c7946e168f8bcef3165bd2c06745b89c5620c  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/service.log
```

## Criterion-Relevant Result

- All four supplied sources are inventoried with stable identifiers, exact paths, source types, locators, and time-zone context.
- Five complete instants are normalized to UTC and deterministically ordered. Three note-derived records remain unplaced because their time evidence is incomplete or absent.
- Automated observations, operator reports, a gap, and bounded hypotheses remain distinct. The package preserves the reported rollback alongside the missing completion record without converting either into a proven fact.
- The suspected deployment relationship is stated as a hypothesis supported only by sequence and proximity, with alternative explanations retained.
- Follow-up actions identify the version-tagged traces, comparative metrics, deployment diff, rollback audit records, and alert evaluation history needed to assess causation.
- JSON syntax, derived-input hash, resolved-event ordering, event counts, and byte-for-byte fixture preservation passed the bounded validations.
