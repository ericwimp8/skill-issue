# Release Readiness Checker Environment Qualification

## Qualified Surface

- Harness surface: Codex Desktop fresh sub-agents backed by Codex CLI `0.145.0-alpha.18`.
- Model: `gpt-5.6-sol` with medium reasoning.
- Qualification date: 2026-07-21.
- Scope: this target campaign, the project-local `skill-issue:release-readiness-checker` discovery entry, and fresh agents started with `fork_turns: "none"`.

## Trial Method

Fresh agent `rrc_environment_probe` received only `evaluation/environment-qualification/request.md` and the authority to use applicable project-local skills. The prompt did not name the candidate skill, quote its description, reveal an expected decision, or prescribe its output beyond the requested report path and safety boundary.

## Direct Evidence

- Native session `019f80a4-4869-71e1-8e83-52ae6ca2f537` identifies the Codex Desktop sub-agent, CLI version, `gpt-5.6-sol`, medium reasoning, and `fork_turns: "none"` lineage.
- Before inspecting the fixture or producing output, its first tool call read the exact target at `showcase-skills/release-readiness-checker/skill/release-readiness-checker/SKILL.md` in the same command as the unmodified request.
- The resulting `report.md` derived all four fixture-owned gates, ran the safe current check, preserved its exact scope and exit status, recorded rollback evidence limits, avoided release effects, and returned a supported `ready` decision.
- Native JSONL evidence is retained under `~/.codex/sessions/2026/07/21/` for this session. The agent's prose claim was not used as the invocation proof.

## Decision

The current Codex Desktop surface is qualified for this campaign's fresh GPT-5.6 Sol medium-reasoning proactive-selection and behavior trials. Qualification applies only while the discovery entry and target hash remain unchanged.
