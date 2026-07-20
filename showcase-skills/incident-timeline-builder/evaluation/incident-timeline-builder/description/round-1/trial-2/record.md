# Trial Record

## Identity

- Trial identity: `1CB84CE4-5A59-46D0-80A5-8A63B0DFBA96`
- Trial: `description/round-1/trial-2`
- Execution: fresh independent description trial 2

## Selection Basis

The candidate was selected from its frontmatter because the request asks for a sourced incident chronology that reconciles deployment, monitoring, and operator records. This directly matches: “Use when logs, alerts, deployments, or operator notes must be reconciled into a sourced incident chronology.” The complete candidate was read before fixture inspection.

## Exact Inputs

- Candidate: `<repo-root>/showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- Request: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-2/request.md`
- Fixture directory: `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/`
- Fixtures:
  - `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/deploy.md`
  - `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/monitor.jsonl`
  - `<repo-root>/showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/operator.md`

## Source Preservation

| Source | Before SHA-256 | After SHA-256 | Result |
| --- | --- | --- | --- |
| Candidate `SKILL.md` | `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1` | `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1` | Preserved |
| `request.md` | `ac0098822267da08930e06ec769beec3adefa247e04ece6fb418f24f0f651a4b` | `ac0098822267da08930e06ec769beec3adefa247e04ece6fb418f24f0f651a4b` | Preserved |
| `deploy.md` | `53ececc56507b832b180f21706364f24e5697c6a994553f851c96731a9330d7a` | `53ececc56507b832b180f21706364f24e5697c6a994553f851c96731a9330d7a` | Preserved |
| `monitor.jsonl` | `388ae73d444f887f91e612b83a551ec0b6b2691b7664c4f8fb8416ea06d63d55` | `388ae73d444f887f91e612b83a551ec0b6b2691b7664c4f8fb8416ea06d63d55` | Preserved |
| `operator.md` | `a77ab79ceb083ee6a976e2004d57cfe087ba2fbb3d10294a7fccc0ea6e02067d` | `a77ab79ceb083ee6a976e2004d57cfe087ba2fbb3d10294a7fccc0ea6e02067d` | Preserved |

## Commands and Outputs

Exact pre-output request, frontmatter, and complete-candidate read commands and their full outputs are preserved in `native-evidence.log`.

### Candidate tool contract

```text
$ python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py --help
usage: build_timeline.py [-h] [--output OUTPUT] input

Normalize explicit timestamps and stably order incident JSONL records.

positional arguments:
  input            UTF-8 JSONL with source_id, source_path, locator,
                   classification, summary, and optional raw_timestamp and
                   source_timezone fields

options:
  -h, --help       show this help message and exit
  --output OUTPUT  write JSON to a new file instead of stdout; existing files
                   are refused
```

### Timeline generation

```text
$ python3 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/scripts/build_timeline.py showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-2/records.jsonl --output showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-2/timeline.json
<no stdout; exit 0>
```

### Derived JSON validation

```text
$ python3 -c 'import json, pathlib; p=pathlib.Path("showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-2/records.jsonl"); records=[json.loads(line) for line in p.read_text().splitlines() if line.strip()]; t=json.loads(pathlib.Path("showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-2/timeline.json").read_text()); assert len(records)==4; assert len(t["resolved_events"])==4; assert t["unresolved_events"]==[]; assert [e["normalized_utc"] for e in t["resolved_events"]]==["2026-08-12T08:10:00Z","2026-08-12T08:15:00Z","2026-08-12T08:15:00Z","2026-08-12T08:20:00Z"]; assert [e["input_order"] for e in t["resolved_events"][1:3]]==[2,3]; print("JSON validation: PASS (4 resolved, 0 unresolved, tie preserved)")'
JSON validation: PASS (4 resolved, 0 unresolved, tie preserved)
```

### Source-hash comparison

```text
$ shasum -a 256 showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-2/request.md showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/*
34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1  showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md
ac0098822267da08930e06ec769beec3adefa247e04ece6fb418f24f0f651a4b  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/description/round-1/trial-2/request.md
53ececc56507b832b180f21706364f24e5697c6a994553f851c96731a9330d7a  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/deploy.md
388ae73d444f887f91e612b83a551ec0b6b2691b7664c4f8fb8416ea06d63d55  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/monitor.jsonl
a77ab79ceb083ee6a976e2004d57cfe087ba2fbb3d10294a7fccc0ea6e02067d  showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/description/trial-2/operator.md
```

## Output Hashes

| Output | SHA-256 |
| --- | --- |
| `package.md` | `d9654adc3126a249f7b3d0a4a7db006883750a79855086640cbfb06ce302b4b2` |
| `records.jsonl` | `d0691f5610ac952aafa266e22c201b64a6e716b0f083d591aaaf370fc9c1f161` |
| `timeline.json` | `0b3cd1494c09790614e117c55dc53b91ac274cf2d4003d756ace6f909a483ee0` |
| `native-evidence.log` | `3bd90afc7555dfeb4c43c98208207bf5dfd4fbdb832ff823d964d4d909167d42` |

## Observable Result

The trial produced a sourced four-event incident chronology in UTC. It preserves the two `08:15:00Z` monitor observations as an exact tie, keeps the operator statement classified as a report, separates three bounded investigative hypotheses from recorded facts, identifies no direct contradiction, and lists evidence gaps and follow-up actions. Derived JSON validation passed with four resolved events, zero unresolved events, and stable input ordering for the tie. Candidate, request, and fixture hashes are unchanged.
