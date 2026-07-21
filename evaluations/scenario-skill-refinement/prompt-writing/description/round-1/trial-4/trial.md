# Description Trial 4

- Target version: `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`
- Stage: confirmation pair
- Fresh agent: `/root/eval_prompt_writing/prompt_desc_r1_t4`
- Session ID: `019f8260-6468-7d63-b065-5433b7d424ee`
- Fixture: `fixtures/description/trial-4-source.md`
- Observable output: `description/round-1/trial-4/output.md`
- Cleanup owner: this campaign; output is isolated from every later trial

## Unmodified Prompt

> Work in <repo-root>. Read evaluations/scenario-skill-refinement/prompt-writing/fixtures/description/trial-4-source.md. Produce the exact commissioning prompt described by that source. Write only that prompt to evaluations/scenario-skill-refinement/prompt-writing/description/round-1/trial-4/output.md. Do not create the implementation plan or change production files. Do not modify any other file. You are not alone in the repository; preserve concurrent work. In your final response, report the output path.

## Native Invocation Evidence

The session trace records a pre-output `exec` call at `2026-07-21T01:53:15.335Z` reading `supporting-skills/prompt-writing/SKILL.md`. See `retained-evidence/description-trial-4-native-trace.json`.

## Audit

- Selection: pass; exact canonical target read before output.
- Authority and evidence: pass; directs the agent to read the research and verify it against production source without restating the report.
- Smallest deliverable: pass; an executor-ready plan only.
- Autonomy boundary: pass; plan writing allowed while implementation, test alteration, commit, and push are excluded.
- Result: **pass**.

