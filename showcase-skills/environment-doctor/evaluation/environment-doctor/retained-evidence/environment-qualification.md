# Campaign Environment Qualification

## Qualified Surface

- Harness surface: Codex fresh sub-agents in the current repository session.
- Model: `gpt-5.6-sol` with medium reasoning.
- Qualification date: 2026-07-21.
- Scope: this target campaign, exact repository-local candidate reads, and fresh agents started with `fork_turns: "none"`.

## Trial Method

Give a fresh agent a natural request and only the candidate frontmatter, decide selection from those meanings, then retain the exact complete-target read and content hash. Answer similarity and prose claims are insufficient without the read commands.

## Direct Evidence

- Replacement probe: `showcase-skills/environment-doctor/evaluation/environment-doctor/qualification/probe-2/record.md`.
- Native read evidence: `showcase-skills/environment-doctor/evaluation/environment-doctor/qualification/probe-2/native-evidence.log`.
- Exact loaded target hash: `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203`.
- The fresh agent selected the candidate from the natural environment-diagnosis request and its frontmatter, then read the exact complete target.

## Interrupted Probe

The first probe under `qualification/probe-1/` executed the candidate and retained useful synthetic evidence but failed to conclude after two bounded-finish instructions. It was interrupted and is excluded from qualification and trial counts.

## Decision

The exact GPT-5.6 Sol medium fresh-agent surface is qualified for candidate selection and direct target-load observation in this campaign.
