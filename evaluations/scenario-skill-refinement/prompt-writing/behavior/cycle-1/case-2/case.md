# Body Case 2: Named Authority And Approval Boundary

- Fresh agent: `/root/eval_prompt_writing/prompt_body_c2`
- Session ID: `019f8269-31a5-7e81-99ef-b724488e8b5a`
- Target version: `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`
- Fixture: `fixtures/body/case-2.md`
- Output: `behavior/cycle-1/case-2/output.md`
- Isolation: the agent was instructed to read only this fixture and no other campaign outputs
- Cleanup ownership: this campaign

## Unmodified Prompt

> Work in <repo-root>. Read and apply supporting-skills/prompt-writing/SKILL.md. Then read only evaluations/scenario-skill-refinement/prompt-writing/fixtures/body/case-2.md and write the requested recipient prompt to evaluations/scenario-skill-refinement/prompt-writing/behavior/cycle-1/case-2/output.md. Do not execute that prompt or inspect other campaign outputs. Do not modify any other file. You are not alone in the repository; preserve concurrent work. Report the output path.

## Ground Truth And Audit

- Task goal: pass; update the deployment step in `plans/release.md`.
- Named authority: pass; directs the recipient to read and follow the canonical release README without restating it.
- Autonomy: pass; plan edit and Markdown validation are allowed.
- Approval and stop boundary: pass; approval precedes deployment, and commit or push remains outside scope.
- Smallest handoff: pass; changed file and validation result.
- Material failure: none.
- Result: **pass**.

