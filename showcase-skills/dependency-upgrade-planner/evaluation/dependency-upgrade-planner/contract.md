# Dependency Upgrade Planner Evaluation Contract

- **Target:** `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/`
- **Goal:** produce a source-backed, dependency-ordered migration plan without changing dependency state or overstating compatibility.
- **Intended use:** requested dependency upgrade planning or impact analysis in an existing project.
- **Expected behavior:** trace current project state and concrete usages, verify applicable authoritative upstream requirements, classify effects, order prerequisites and migration stages by dependency edges, and preserve evidence boundaries.
- **Expected result:** a coherent plan covering scope, graph, usages, sources, effects, breaking changes, prerequisites, ordered work, validation, rollback, stops, exclusions, and unresolved risk.
- **Boundary:** read-only planning; no dependency mutation, migration execution, or unsupported compatibility claim.

## Observable Criteria

1. Natural upgrade-planning prompts select and load the target.
2. Plans inspect and cite manifests, lockfiles, configuration, platform declarations, and concrete production usage.
3. Plans distinguish direct, transitive, build-tool, runtime, and platform effects with evidence.
4. Plans use live authoritative upstream sources and explain version applicability.
5. Plans identify breaking changes, compatibility edges, prerequisites, and unknowns without inventing certainty.
6. Plans order stages by dependency edges and behavior owners.
7. Plans define relevant validation evidence, rollback points, stop conditions, exclusions, and a smallest safe first stage.
8. Plans label project facts, upstream facts, implications, and unknowns and remain planning-only.

## Evaluation Configuration

- **Environment qualification:** `showcase-skills/dependency-upgrade-planner/evaluation/environment-qualification.md`
- **Refinement mode:** automatic within the assigned showcase workspace.
- **Description protocol:** two representative trials followed by two different confirmation trials.
- **Reference protocol:** not applicable because the target packages no references.
- **Body surface:** generated document plans from isolated connected fixtures.
