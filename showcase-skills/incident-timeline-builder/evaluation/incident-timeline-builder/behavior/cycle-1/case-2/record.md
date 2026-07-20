# Evaluation Record

## Run Identity

- Identity: `body-case-2-gpt-5.6-sol-medium-2026-07-20T18:38:07Z`
- Execution mode: fresh independent body case 2
- Model: GPT-5.6 Sol
- Reasoning effort: medium
- Working directory: `<repo-root>`
- Completed at: `2026-07-20T18:38:07Z`

## Inputs

- Target: `<repo-root>/showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- Target SHA-256: `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1`
- Request: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/request.md`
- Request SHA-256: `43306491a963f363eb2b45fb53ed5344ecda9c35a58e869d63b2d44612f4e67d`
- Fixture directory: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-2/`

## Fixture Preservation

| Fixture | Before SHA-256 | After SHA-256 | Result |
| --- | --- | --- | --- |
| `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-2/events.jsonl` | `74546b545c1ecd9bf058bbd5172231be9d480198b58ad6912ad53af287ba8654` | `74546b545c1ecd9bf058bbd5172231be9d480198b58ad6912ad53af287ba8654` | Preserved |
| `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-2/operator.md` | `9ad254a303722604a744626381484e482c90a70f0d4530ecec4c4ec6ea585471` | `9ad254a303722604a744626381484e482c90a70f0d4530ecec4c4ec6ea585471` | Preserved |

## Outputs

| Artifact | SHA-256 |
| --- | --- |
| `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/package.md` | `eda370293f7b5cfa46a7a25a76bbca2f1a6c328a1ec1336c22e2400c11fa527d` |
| `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/records.jsonl` | `eb225e181ecfbec3c6e626cc1b7131ba366d0ea11c7711621d6b820cae6b41cb` |
| `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/timeline.json` | `f1c2b4df36914a5e3d06573b66cea46371ef5f707d6c39d7637ff1c4d4521fb0` |
| `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/native-evidence.log` | `ff9b1fc7aaf014c9ad1e742cee65c6f791cedb5bc734ef8665c2e4f8f44aa265` |

## Commands and Observable Outputs

The complete target and request reads, with their exact commands and output, are retained in `native-evidence.log`.

```text
$ shasum -a 256 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/request.md
34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1  showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md
43306491a963f363eb2b45fb53ed5344ecda9c35a58e869d63b2d44612f4e67d  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/request.md

$ python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/records.jsonl --output showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/timeline.json
[no stdout; exit 0]

$ python3 -m json.tool showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-2/timeline.json >/dev/null
[no stdout; exit 0]

$ python3 -c '<JSONL parsing, input hash, resolved ordering, and unresolved-status assertions>'
JSON valid: records=5, resolved=2, unresolved=3
input_sha256 match: eb225e181ecfbec3c6e626cc1b7131ba366d0ea11c7711621d6b820cae6b41cb
resolved order: input_order 1,2 at exact tie; unresolved statuses: ambiguous,missing,missing

$ shasum -a 256 showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-2/events.jsonl showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-2/operator.md
74546b545c1ecd9bf058bbd5172231be9d480198b58ad6912ad53af287ba8654  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-2/events.jsonl
9ad254a303722604a744626381484e482c90a70f0d4530ecec4c4ec6ea585471  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-2/operator.md
```

## Criterion-Relevant Result

- Both machine events normalize from `2026-10-25T01:15:00+01:00` to the exact tie `2026-10-25T00:15:00Z`; deterministic input order is retained without claiming finer chronology.
- The operator's `2026-10-25 01:30:00` in `Europe/London` is classified as ambiguous and remains unresolved.
- The untimed second restart remains unresolved while preserving only the reported relation that it happened after the first restart.
- The operator's recovery belief is retained as a reported hypothesis; missing identifiers and outcome evidence prevent a causal conclusion.
- JSON syntax, generated input hash, resolved-event ordering, unresolved statuses, and byte preservation of both fixtures were validated successfully.
