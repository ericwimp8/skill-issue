# API Change Impact Mapper Evaluation Contract

- **Target:** `showcase-skills/api-change-impact-mapper/skill/api-change-impact-mapper/`
- **Goal:** trace supplied contract changes through concrete repository and external boundaries and produce an evidence-bounded migration map.
- **Intended use:** compatibility and migration analysis for API, schema, event, protocol, or public-interface changes.
- **Expected behavior:** establish semantic deltas; trace concrete producers, transformations, consumers, generated artifacts, persistence, configuration, and external boundaries; classify each path; preserve unknowns; and derive owner-oriented migration, validation, rollout, and rollback work.
- **Expected result:** an actionable report containing an evidence graph, per-path compatibility matrix, ordered migration sequence, validation plan, rollout observation, rollback constraints, and open decisions.
- **Boundary:** read-only analysis; source-backed claims; no unsupported absence of external consumers; no global compatibility label when paths differ; no implementation modification.
- **Evaluation surface:** generated Markdown analysis from isolated synthetic repository fixtures.

## Observable Criteria

1. Natural change-impact requests select the target without naming it.
2. Semantic contract deltas distinguish facts, ambiguity, and unavailable detail.
3. Evidence tracing reaches concrete producers, transformations, and consumers.
4. Relevant generation, persistence, configuration, and external surfaces are traced or explicitly bounded.
5. Compatibility is classified per material path with versions, conditions, rationale, and confidence.
6. Unknown consumers and checked-negative evidence remain visible.
7. Migration actions identify behavior owners or owner types and dependency order.
8. Validation, staged rollout signals, and rollback constraints derive from traced paths.
9. Claims link to repository-relative evidence and preserve fact, inference, assumption, and unknown distinctions.
10. The analyzed fixture remains unmodified.

## Environment Qualification

- **Harness surface:** Codex fresh sub-agents in the repository workspace.
- **Model and reasoning:** `gpt-5.6-sol`, medium.
- **Qualification basis:** the repository's Codex custom evaluation surface was locally qualified on 2026-07-20 for this model and reasoning in `cli/README.md`, `plans/harness-setup.md`, and retained showcase evaluation evidence.
- **Trial method:** one fresh `fork_turns: "none"` agent per isolated prompt, with exact candidate selection and pre-output target reads retained in each trial record.
- **Direct evidence rule:** accept the native agent identity plus its pre-output read of the exact candidate; prose claims and answer similarity are insufficient.

## Refinement Mode

Automatic semantic refinement is authorized only inside `showcase-skills/api-change-impact-mapper/`. A retained material failure must be corrected at its semantic owner and affected trials repeated from clean fixtures.
