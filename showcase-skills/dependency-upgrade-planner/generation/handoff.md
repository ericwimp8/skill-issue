# Generation-to-Evaluation Handoff

## Generated Target

- Skill: `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/SKILL.md`
- OpenAI metadata: `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/agents/openai.yaml`
- Packaged references, scripts, and assets: none; project-specific evidence belongs to each inspected upgrade.

## Evaluation Contract

- **Goal:** produce a source-backed, dependency-ordered migration plan without changing dependency state or overstating compatibility.
- **Intended use:** requested dependency upgrade planning or impact analysis in an existing project.
- **Expected behavior:** establish the upgrade contract; trace manifests, lockfiles, configuration, platform declarations, and concrete production usages; verify authoritative upstream requirements; classify effects; order prerequisites and work by dependency edges; and define evidence gates and rollback points.
- **Expected result:** a coherent plan separating project facts, upstream facts, implications, and unknowns while covering current graph, applicable changes, impact classes, ordered stages, validation, rollback, stops, and exclusions.
- **Boundaries:** planning and read-only inspection only; no dependency edits, installation, lockfile regeneration, codemod execution, migration execution, or unsupported compatibility claims.
- **Supported surface:** portable Agent Skills content with OpenAI Codex interface metadata. Runtime reliability outside the evaluated fresh-agent candidate-selection surface is unclaimed.
- **Refinement mode:** automatic within `showcase-skills/dependency-upgrade-planner/` under semantic-owner constraints.

## Runtime Criteria

Independent evidence must establish description selection and all completion criteria from the intake plan across varied application, build-tool/platform, and Python service fixtures. The target has no references, so per-reference qualification is not applicable.
