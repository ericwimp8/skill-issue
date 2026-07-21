# Body Case 1 Record

- Fresh agent: `/root/eval_systematic_debugging/sysdebug_body_1` (`Planck`)
- Session: `019f8278-9e7d-77a3-b3d6-75b3eab67c37`
- Fixture: `fixtures/behavior/case-1/`
- Output: `case-1-output.md`
- Result: pass

## Unmodified Prompt

> Work in <repo-root>. Read and apply supporting-skills/systematic-debugging/SKILL.md. Work only in evaluations/scenario-skill-refinement/systematic-debugging/fixtures/behavior/case-1. Two symptoms are reported: semantically identical query objects can miss the cache, and invalidating one ordering can leave the other ordering present. Establish intended behavior from production source, reproduce both symptoms, trace the concrete path to their first shared incorrect condition, and apply the smallest supported correction at its owner. Preserve pre-correction evidence, then run the focused reproduction and nearest related checks. Write a concise report with changed files and commands to evaluations/scenario-skill-refinement/systematic-debugging/behavior/cycle-1/case-1-output.md. Do not inspect other campaign cases or outputs. You are not alone in the repository; do not revert or alter unrelated work.

## Audit

The native first tool call read the exact target. The report established the source contract, retained both pre-correction symptoms, traced `getCached` and `invalidate` through their shared `cacheKey` owner, corrected deterministic serialization there, and passed the focused and related checks. The change preserved distinct query values and paths. All completion criteria pass.
