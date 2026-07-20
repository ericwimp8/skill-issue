# Campaign Environment Qualification

## Qualified Surface

- **Harness surface:** Codex collaboration fresh sub-agent started with `fork_turns: "none"`.
- **Model:** `gpt-5.6-sol` with medium reasoning.
- **Qualification date:** 2026-07-21.
- **Scope:** this incident-timeline-builder campaign, its exact candidate at `showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`, and fresh agents started through the same collaboration surface with identical model, reasoning, and clean-context settings.

## Trial Method

The probe received a natural, unmodified incident-chronology request without the skill name, description text, expected answer, or prior conclusion. The harness separately advertised the candidate path. The probe inspected frontmatter, recorded its selection decision, then read the complete exact candidate before producing output. Selection evidence requires the pre-output candidate read and cannot be inferred from answer similarity or agent prose.

## Direct Evidence

- Fresh agent: `/root/incident_timeline_builder/qualification_probe_1`.
- Native spawn configuration: `fork_turns: "none"`, model `gpt-5.6-sol`, reasoning `medium`.
- Exact request: `description/round-1/trial-1/request.md`.
- Candidate load trace: `qualification/probe-1/native-evidence.log` records the frontmatter read used for selection and the complete candidate read before source inspection or output creation.
- Selection and preservation record: `qualification/probe-1/record.md` records the semantic selection basis, output paths, and matching before/after SHA-256 values for all four source files.
- Observable result: `qualification/probe-1/output.md` and `timeline.json` correctly normalize all explicit offsets, retain two untimed reports as unplaced evidence, preserve provenance, and bound the deployment hypothesis without asserting causation.

## Qualification Decision

The specified clean-context GPT-5.6 Sol medium collaboration surface is qualified for this campaign's proactive-selection and body trials. Qualification is candidate- and campaign-local; it does not establish reliability for other skills, models, reasoning settings, or harness surfaces.
