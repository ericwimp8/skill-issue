# Description Trial 2

- Target: `supporting-skills/prompt-writing/`
- Target version: `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`
- Stage: initial pair
- Fresh agent: `/root/eval_prompt_writing/prompt_desc_r1_t2`
- Session ID: `019f825e-0542-7680-8e09-ef51b4177f4e`
- Harness: Codex Desktop sub-agent, Codex CLI `0.145.0-alpha.18`
- Fixture: `fixtures/description/trial-2-source.md`
- Observable output: `description/round-1/trial-2/output.md`
- Cleanup owner: this campaign; output is isolated from every later trial

## Unmodified Prompt

> Work in <repo-root>. Read evaluations/scenario-skill-refinement/prompt-writing/fixtures/description/trial-2-source.md. Produce the exact prompt the product lead should send to the independent discovery agent described there. Write only that prompt to evaluations/scenario-skill-refinement/prompt-writing/description/round-1/trial-2/output.md. Do not conduct the investigation. Do not modify any other file. You are not alone in the repository; preserve concurrent work. In your final response, report the output path.

## Native Invocation Evidence

Before writing the output, the native session trace records an `exec` tool call at `2026-07-21T01:50:36.687Z` whose command reads the exact canonical file `supporting-skills/prompt-writing/SKILL.md`. The retained excerpt is `retained-evidence/description-trial-2-native-trace.json`.

## Audit

- Selection: pass; exact canonical target read before output.
- Task-specific goal: pass.
- Discovery framing: pass; requests the governing category, owners, conflicts, and questions without seeding presumed mechanisms.
- Deliverable and autonomy: pass; smallest planning handoff and no implementation.
- Scope and prescription: pass; concise and directly usable.
- Result: **pass**.

