# Description Trial 2

- Target: `supporting-skills/systematic-debugging/`
- Loop: description
- Round: 1
- Trial: 2, representative initial pair
- Fresh agent: `/root/eval_systematic_debugging/sysdebug_desc_2` (`Sagan`)
- Session: `019f8261-457a-7a73-b78b-98aa0ee0816f`
- Fixture: `fixtures/description/trial-2/`
- Output: `trial-2-output.md`
- Result: pass

## Unmodified Prompt

> Work in <repo-root>. A small JavaScript reproduction under evaluations/scenario-skill-refinement/systematic-debugging/fixtures/description/trial-2 routes a lowercase EU request to the default destination. Investigate the failure from source and runtime evidence, then write a concise diagnosis and smallest supported correction proposal to evaluations/scenario-skill-refinement/systematic-debugging/description/round-1/trial-2-output.md. Do not modify the fixture or any production skill. Include the reproduction command and evidence that distinguishes the failure surface from its cause. You are not alone in the repository; do not revert or alter unrelated work.

## Native Invocation Evidence

The session's first tool call read `supporting-skills/systematic-debugging/SKILL.md` before it inspected or executed the fixture. That exact pre-output target read satisfies the qualified evidence rule.

- Native trace retained in `retained-evidence/description-trial-2-native-trace.md`.
- Original runtime trace: `~/.codex/sessions/2026/07/21/rollout-2026-07-21T11-24-00-019f8261-457a-7a73-b78b-98aa0ee0816f.jsonl`.

## Observable Result

The agent reproduced the failure, compared lowercase, uppercase, other, and absent header values, located the first incorrect control decision at the case-sensitive comparison, and tested normalization without changing the fixture. The proposed correction is placed at the route decision that owns case handling.

## Criterion Result

- Proactive target selection: pass.
- Description boundary match: pass; the task naturally requested diagnosis of unexpected routing behavior without naming or hinting at a skill.
- Task quality: pass against the campaign contract.
- Cleanup ownership: the fixture remains unchanged; the retained report and trace belong to this trial.
