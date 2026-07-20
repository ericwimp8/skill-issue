# Description Trial 4 Record

## Trial identity

- **Fresh agent identity:** `/root/dependency_upgrade_planner/dep_desc_4`
- **Model:** `gpt-5.6-sol`
- **Reasoning:** medium
- **Fixture:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/vite-library/`
- **Prompt:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-4/prompt.md`
- **Output:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-4/output.md`

## Candidate selection

- Inspected the YAML frontmatter descriptions for all four supplied candidates before reading a complete candidate.
- Selected `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/SKILL.md` because the prompt requests source-backed planning for a dependency major-version change without implementation.
- Complete-read evidence and hashes are retained in `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-4/native-evidence.log`.
- This is direct candidate-selection and read evidence on the qualified fresh-agent surface. It makes no native automatic-injection claim.

## Sources inspected

- Fixture production manifest, active lockfile, consumer/deployment declaration, Vite configuration, TypeScript entry point, and SCSS source under `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/vite-library/`.
- Target evaluation contract and behavior ground truth under `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/`.
- Live official Vite 6 migration guide: `https://v6.vite.dev/guide/migration.html`.
- Live official Vite 6 announcement: `https://v6.vite.dev/blog/announcing-vite6`.

## Eight-criterion contract audit

1. **PASS — Natural selection:** The unleading upgrade-planning prompt selected and completely loaded the target skill; `native-evidence.log` retains direct evidence.
2. **PASS — Project inspection:** The plan cites the manifest, lockfile, configuration, platform declarations, consumer declaration, entry point, and stylesheet source.
3. **PASS — Effect classification:** The impact table separates direct build dependencies, transitive Rollup, build-tool configuration, Node platform, and produced-artifact/consumer effects.
4. **PASS — Live authoritative sources:** The plan uses the live official Vite 6 migration guide and release announcement and explains why each documented change applies.
5. **PASS — Breaks, prerequisites, unknowns:** Node 21, Terser, custom conditions, CSS naming, Sass choice, target patches, and the missing condition probe are bounded without invented certainty.
6. **PASS — Dependency and ownership order:** Node platform policy precedes configuration contracts, direct dependency resolution, transitive lock resolution, build validation, and deployment proof.
7. **PASS — Gates, rollback, stops, exclusions:** Every state-changing stage has evidence, stop, and rollback coverage; exclusions and the smallest safe first stage are explicit.
8. **PASS — Evidence labels and planning boundary:** Project facts, upstream facts, implications, and unknowns are labeled, and the output makes no build, test, runtime, or compatibility success claim.

## Cleanup and result

- Fixture files were not modified.
- No dependency state, generated artifact, install state, or disposable evaluation output was created.
- Trial artifacts are deliberately retained in the assigned trial directory.
- **Overall: PASS**
