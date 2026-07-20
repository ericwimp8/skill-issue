# Evaluation Record

- Run identity: `700014e7-59fa-4f8d-912d-683ff2c4953c`
- Recorded at: `2026-07-20T18:38:11Z`
- Case directory: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/`
- Target skill: `<repo-root>/showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- Target SHA-256: `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1`
- Request: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/request.md`
- Request SHA-256: `6351851a8ae8698bd6af48a64f243581b981e9f9e4b354b08e774cd51d3b9733`
- Fixture directory: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-3/`

## Outputs

- `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/package.md` — SHA-256 `ba925f9d661a694bd937ac6af7a734574e43f96c1707f9f410af5ee6df78af75`
- `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/records.jsonl` — SHA-256 `9157d738e344551e62efdc14d3c83b062456f2c020fd8c08781f517f15ab669a`
- `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/timeline.json` — SHA-256 `838308e69de5b845f6d476b5b9ee04d00e8510761333ae8bfd8111dfce2fec91`
- `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/native-evidence.log` — SHA-256 `ad6ea82b8865a7895bbd49e3b135c7ef89e3754b743885137b347221a0a3ffc2`

## Fixture Preservation

| Fixture | Before SHA-256 | After SHA-256 | Result |
| --- | --- | --- | --- |
| `gateway.log` | `9639936e3a0ba7edc1c28af5a15b0081d15060e76a2cbacc96a7c3664c57f92e` | `9639936e3a0ba7edc1c28af5a15b0081d15060e76a2cbacc96a7c3664c57f92e` | Preserved |
| `private-note.md` | `915c1e8d67260edec9470f1671968f1c0593f90416f479a44bdb8373d4384c13` | `915c1e8d67260edec9470f1671968f1c0593f90416f479a44bdb8373d4384c13` | Preserved |
| `status-page.md` | `28c1e1276a6a54d57a2ff21aedc76c968e25800e87d8f1edb17771de6adedeec` | `28c1e1276a6a54d57a2ff21aedc76c968e25800e87d8f1edb17771de6adedeec` | Preserved |

## Commands and Observable Outputs

### Target and request reads

The exact pre-output commands and their complete observable outputs are preserved in `native-evidence.log`:

```text
cat showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md
cat showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/request.md
```

### Timeline generation

```text
python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/records.jsonl --output showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/timeline.json
```

Observable output: no stdout or stderr; exit status `0`; `timeline.json` created.

### Validation

```text
python3 -c 'import json, pathlib; p=pathlib.Path("showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/records.jsonl"); [json.loads(line) for line in p.read_text().splitlines() if line.strip()]; print("records JSONL: PASS")'
python3 -m json.tool showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3/timeline.json >/dev/null
python3 -c 'import hashlib,json,pathlib; base=pathlib.Path("showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/behavior/cycle-1/case-3"); raw=(base/"records.jsonl").read_bytes(); data=json.loads((base/"timeline.json").read_text()); assert data["input_sha256"]==hashlib.sha256(raw).hexdigest(); times=[e["normalized_utc"] for e in data["resolved_events"]]; assert times==sorted(times); assert len(data["resolved_events"])==2 and len(data["unresolved_events"])==3; print("timeline JSON: PASS"); print("input hash: PASS", data["input_sha256"]); print("resolved ordering: PASS", times); print("partition: PASS 2 resolved, 3 unresolved")'
find showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-3 -type f -exec shasum -a 256 {} \; | sort
```

Observable output:

```text
records JSONL: PASS
timeline JSON: PASS
input hash: PASS 9157d738e344551e62efdc14d3c83b062456f2c020fd8c08781f517f15ab669a
resolved ordering: PASS ['2026-02-02T04:00:00Z', '2026-02-02T04:06:00Z']
partition: PASS 2 resolved, 3 unresolved
28c1e1276a6a54d57a2ff21aedc76c968e25800e87d8f1edb17771de6adedeec  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-3/status-page.md
915c1e8d67260edec9470f1671968f1c0593f90416f479a44bdb8373d4384c13  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-3/private-note.md
9639936e3a0ba7edc1c28af5a15b0081d15060e76a2cbacc96a7c3664c57f92e  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-3/gateway.log
```

## Criterion-Relevant Result

The shareable package inventories all three supplied sources, preserves provenance at line-level locators, and keeps the two fully dated gateway observations in deterministic UTC order. The two date-less status-page reports and the untimed credential-rotation report remain unplaced rather than receiving invented dates or sequence. The package avoids credential-value reconstruction, limits causal statements to bounded hypotheses, and requests only actionable non-secret audit metadata. JSON structure, timeline input hash, resolved ordering, unresolved partition, and byte-for-byte fixture preservation all pass.
