# API Change Impact Mapper Generation Handoff

- **Generated skill:** `showcase-skills/api-change-impact-mapper/skill/api-change-impact-mapper/`
- **Canonical target:** `showcase-skills/api-change-impact-mapper/skill/api-change-impact-mapper/SKILL.md`
- **Intake contract:** `showcase-skills/api-change-impact-mapper/plans/api-change-impact-mapper/api-change-impact-mapper-a-to-b-plan.md`
- **Supported surfaces:** portable Agent Skills content and OpenAI Codex project or user delivery.
- **Goal:** trace supplied contract changes through concrete repository and external boundaries and produce an evidence-bounded migration map.
- **Intended use:** compatibility and migration analysis for API, schema, event, protocol, or public-interface changes.
- **Expected behavior:** establish semantic deltas; trace producers, transformations, consumers, generated artifacts, persistence, configuration, and external boundaries; classify paths; preserve unknowns; and derive owner-oriented validation, rollout, and rollback work.
- **Expected result:** an actionable report containing an evidence graph, per-path compatibility matrix, ordered migration sequence, validation, rollout observation, rollback constraints, and open decisions.
- **Preserved boundaries:** read-only analysis, no inferred absence of external consumers, no unsupported compatibility claim, and no generic plan detached from repository evidence.
- **Runtime criteria:** all ten intake completion criteria require representative execution evidence.
- **Resources:** no reference, script, or asset is required; the behavior is contextual reasoning over supplied contracts and repository evidence.
- **Known limitations:** conclusions remain bounded by supplied contract authority, repository revision, accessible deployments and configuration, and externally visible consumer evidence.
- **Refinement mode:** automatic semantic refinement within the showcase workspace when retained evidence establishes a material failure.
- **Evaluation route:** `showcase-skills/api-change-impact-mapper/evaluation/api-change-impact-mapper/`.
- **Generation decision:** continue directly into evaluation; no user-owned stop is active.
