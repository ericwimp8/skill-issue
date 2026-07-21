# Description Trial 3

- Target version: `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`
- Stage: confirmation pair
- Fresh agent: `/root/eval_prompt_writing/prompt_desc_r1_t3`
- Session ID: `019f825f-77a3-7f70-b91f-e2459d0226fd`
- Fixture: `fixtures/description/trial-3-source.md`
- Observable output: `description/round-1/trial-3/output.md`
- Cleanup owner: this campaign; output is isolated from every later trial

## Unmodified Prompt

> Work in <repo-root>. Read evaluations/scenario-skill-refinement/prompt-writing/fixtures/description/trial-3-source.md. Produce the exact task prompt described by that source. Write only that prompt to evaluations/scenario-skill-refinement/prompt-writing/description/round-1/trial-3/output.md. Do not perform the documentation update. Do not modify any other file. You are not alone in the repository; preserve concurrent work. In your final response, report the output path.

## Native Invocation Evidence

The session trace records a pre-output `exec` call at `2026-07-21T01:52:23.206Z` reading `supporting-skills/prompt-writing/SKILL.md`. See `retained-evidence/description-trial-3-native-trace.json`.

## Audit

- Selection: pass; exact canonical target read before output.
- Authority reuse: pass; points to the document discipline without reteaching it.
- Goal, autonomy, and completion: pass; identifies the update, automatic edit boundary, stop boundary, and smallest handoff.
- Result: **pass**.

