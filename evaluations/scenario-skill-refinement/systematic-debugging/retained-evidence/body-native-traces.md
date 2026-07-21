# Body Native Traces

## Case 1

- Session: `019f8278-9e7d-77a3-b3d6-75b3eab67c37`
- First tool call: exact `cat supporting-skills/systematic-debugging/SKILL.md` before fixture inspection.
- Complete JSONL: `~/.codex/sessions/2026/07/21/rollout-2026-07-21T11-49-30-019f8278-9e7d-77a3-b3d6-75b3eab67c37.jsonl`.

## Case 2

- Session: `019f8278-ae7b-7653-ad37-4c4b10165310`
- First tool call: exact `cat supporting-skills/systematic-debugging/SKILL.md` followed by the complete assigned fixture.
- Complete JSONL: `~/.codex/sessions/2026/07/21/rollout-2026-07-21T11-49-34-019f8278-ae7b-7653-ad37-4c4b10165310.jsonl`.

## Case 3

- Session: `019f8279-f6ed-7b93-a7bc-f2875272d584`
- First tool call: exact pre-output `sed` read of `supporting-skills/systematic-debugging/SKILL.md` before fixture inspection.
- Complete JSONL: `~/.codex/sessions/2026/07/21/rollout-2026-07-21T11-50-59-019f8279-f6ed-7b93-a7bc-f2875272d584.jsonl`.

These retained excerpts preserve the target-read evidence. The body prompts explicitly required the target, so these traces establish application context rather than proactive description selection.
