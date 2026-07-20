# Description Selection Evaluation Record

## User Task

Assess and sequence the migration of `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/vite-library/` from Vite 5.4 to Vite 6, including build and deployment constraints, before any dependency is changed.

## Evaluation Identity

- **Fresh agent:** `/root/dependency_upgrade_planner/dep_desc_2`
- **Model:** `gpt-5.6-sol`
- **Reasoning:** medium
- **Candidate set:**
  - `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/SKILL.md`
  - `skills/skill-generation/SKILL.md`
  - `supporting-skills/document-update-discipline/SKILL.md`
  - `supporting-skills/prompt-writing/SKILL.md`
- **Selection decision:** Selected `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/SKILL.md`. Its description directly matches source-backed planning for a requested dependency version change without changing dependency state. The other candidates own skill generation, document editing discipline, or agent-prompt authoring rather than the requested migration assessment.
- **Evidence log:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-2/native-evidence.log`
- **Fixture:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/vite-library/`
- **Prompt:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-2/prompt.md`
- **Output:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-2/output.md`

## Sources

- Fixture `package.json`, abbreviated `package-lock.json`, `vite.config.ts`, `package-consumer.json`, `src/index.ts`, and `src/theme.scss`.
- Vite 6 migration guide for resolver defaults, Sass API behavior, library CSS naming, and Terser minimum: `https://v6.vite.dev/guide/migration.html`.
- Vite 6 build options for the installed-Terser requirement: `https://v6.vite.dev/config/build-options.html#build-minify`.
- Vite 6.0.0 package metadata for the Node engine range: `https://github.com/vitejs/vite/blob/v6.0.0/packages/vite/package.json`.
- Node.js EOL policy and release table for Node 21 status: `https://nodejs.org/en/about/eol`.

## Eight-Criterion Contract Audit

1. **Natural selection and complete load:** PASS. The unleading Vite migration request selected the dependency upgrade planner; candidate hashes and the complete selected-skill read hash were retained before output.
2. **Project inspection and citation:** PASS. The plan cites the manifest, lockfile, Vite configuration, deployment/consumer declaration, entry source, and SCSS usage.
3. **Effect classification:** PASS. The graph and impact sections distinguish direct, transitive, build-tool, runtime, platform, packaging, and overlapping effects with project evidence.
4. **Live authoritative upstream sources:** PASS. The plan applies Vite-owned Vite 6 migration/configuration/package metadata and Node-owned lifecycle evidence to the exact observed surfaces.
5. **Breaks, edges, prerequisites, unknowns:** PASS. Node 21, Terser, resolver conditions, Sass mode, CSS naming, transitive-lock limits, exact target selection, and absent deployment evidence remain explicit and appropriately bounded.
6. **Dependency and behavior-owner ordering:** PASS. Runtime/contract decisions and a baseline precede reversible configuration trials, runtime/Terser alignment, the Vite/lock transition, and packaging/deployment validation.
7. **Validation, rollback, stops, exclusions, smallest stage:** PASS. Every state-changing stage has evidence gates, rollback scope, stop conditions, and the plan ends with a no-dependency-change Stage 0.
8. **Evidence labels and planning boundary:** PASS. Observed project facts, authoritative upstream facts, reasoned implications, and unknowns are labeled; no dependency, fixture, build, test, package, or deployment state was changed or claimed successful.

## Cleanup Ownership

The evaluation owner may remove the four temporary trial artifacts in `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-2/` when this evaluation no longer has downstream value: `prompt.md`, `native-evidence.log`, `output.md`, and `record.md`.

## Result

PASS
