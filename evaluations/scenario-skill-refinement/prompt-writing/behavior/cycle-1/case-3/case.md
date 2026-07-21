# Body Case 3: Open-Ended Discovery

- Fresh agent: `/root/eval_prompt_writing/prompt_body_c3`
- Session ID: `019f826c-609d-7710-8efe-329654918849`
- Target version: `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`
- Fixture: `fixtures/body/case-3.md`
- Output: `behavior/cycle-1/case-3/output.md`
- Isolation: the agent was instructed to read only this fixture and no other campaign outputs
- Cleanup ownership: this campaign

## Unmodified Prompt

> Work in <repo-root>. Read and apply supporting-skills/prompt-writing/SKILL.md. Then read only evaluations/scenario-skill-refinement/prompt-writing/fixtures/body/case-3.md and write the requested discovery prompt to evaluations/scenario-skill-refinement/prompt-writing/behavior/cycle-1/case-3/output.md. Do not execute that prompt or inspect other campaign outputs. Do not modify any other file. You are not alone in the repository; preserve concurrent work. Report the output path.

## Ground Truth And Audit

- Discovery category: pass; asks what makes evaluation evidence durable for later audit.
- Unseeded investigation: pass; requests categories, owners, contradictions, and missing proof without naming expected artifact types.
- Source and inference discipline: pass; tracked files only, source paths, and separation of repository guidance from inference and unresolved questions.
- Smallest deliverable: pass; concise evidence map for a later design decision.
- Autonomy boundary: pass; read-only repository work.
- Material failure: none.
- Result: **pass**.

