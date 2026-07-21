# Body Case 1: Repair An Over-Prescribed Handoff

- Fresh agent: `/root/eval_prompt_writing/prompt_body_c1`
- Session ID: `019f8267-666b-76d2-b9c1-96e330f67194`
- Target version: `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`
- Fixture: `fixtures/body/case-1.md`
- Output: `behavior/cycle-1/case-1/output.md`
- Isolation: the agent was instructed to read only this fixture and no other campaign outputs
- Cleanup ownership: this campaign

## Unmodified Prompt

> Work in <repo-root>. Read and apply supporting-skills/prompt-writing/SKILL.md. Then read only evaluations/scenario-skill-refinement/prompt-writing/fixtures/body/case-1.md and write the revised recipient prompt to evaluations/scenario-skill-refinement/prompt-writing/behavior/cycle-1/case-1/output.md. Do not execute that prompt or inspect other campaign outputs. Do not modify any other file. You are not alone in the repository; preserve concurrent work. Report the output path.

## Ground Truth And Audit

- Fix the failed ownership decision: pass; the prompt requires tracing to the concrete owner.
- Remove broad accumulated instructions: pass; the output drops the exhaustive reading, report structure, and implementation request.
- Preserve needed source access: pass; production source and focused Git history remain available.
- Smallest deliverable: pass; ownership finding, source locations, and unresolved uncertainty only.
- Autonomy boundary: pass; no edits, tests, or adjacent fixes.
- Material failure: none.
- Result: **pass**.

