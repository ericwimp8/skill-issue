# Campaign Environment Qualification

## Qualified Surface

- Harness surface: Codex Desktop fresh sub-agents backed by Codex CLI `0.145.0-alpha.18`.
- Model: `gpt-5.6-sol` with high reasoning.
- Qualification date: 2026-07-19.
- Scope: this local campaign, project-local skills advertised from `.codex/skills/`, and fresh agents started with `fork_turns: "none"`.

## Trial Method

Give fresh agents unleading tasks that naturally require an advertised skill. Accept only a native injection event or a native pre-output tool trace reading the exact candidate `SKILL.md`. Ignore answer similarity and final prose claims.

## Direct Evidence

- Evaluator probe `019f7766-23a9-7ce1-b0c1-aa66d7e7ca39` independently read `skills/skill-evaluation-and-refinement/SKILL.md` before producing its assessment.
- Evaluator probe `019f7766-38cb-7752-bead-5004c4a6424b` independently read the same evaluator before producing its assessment. Its session context identifies `gpt-5.6-sol` and high reasoning.
- Prompt-writing assessment retained four fresh-agent traces that read the exact local candidate before output.
- Document-update assessment retained a native `codex.skill.injected` event for `skill-issue:document-update-discipline` and an exact candidate read.

## Qualification Decision

The current Codex surface is qualified for campaign-local proactive-selection trials with direct evidence. Duplicate advertised identities remain a routing risk and must be isolated or disclosed for any affected target.
