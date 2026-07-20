# Trial Record

- **Agent identity:** `/root/dependency_upgrade_planner/dep_desc_1`
- **Model:** inherited runtime model; exact identifier was not exposed to this agent
- **Reasoning setting:** inherited; exact setting was not exposed to this agent
- **User prompt:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-1/prompt.md`
- **Fixture:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/react-app/`
- **Production evidence inspected:** `package.json`, `package-lock.json`, `tsconfig.json`, `src/main.tsx`, `src/legacy-panel.tsx`
- **Selection:** `dependency-upgrade-planner`; selected by comparing candidate YAML frontmatter descriptions. The selected skill was then read completely and its hash recorded in `native-evidence.log`. This record makes no claim of native automatic injection.
- **Supporting discipline used:** `document-update-discipline` for coherent ownership and reconciliation of the generated documents.

## Authoritative public sources

- https://react.dev/versions
- https://react.dev/blog/2024/04/25/react-19-upgrade-guide
- https://react.dev/blog/2025/10/01/react-19-2
- https://react.dev/community/versioning-policy
- https://registry.npmjs.org/react/latest
- https://registry.npmjs.org/react-dom/latest
- https://registry.npmjs.org/react/18.3.1
- https://registry.npmjs.org/react-dom/18.3.1
- https://registry.npmjs.org/%40types%2freact/latest
- https://registry.npmjs.org/%40types%2freact-dom/latest
- https://registry.npmjs.org/vite/5.4.21

## Observable-criteria audit

1. **PASS — natural selection:** the request is explicitly an existing-project dependency upgrade plan, matching the selected skill description.
2. **PASS — project inspection:** the plan cites the manifest, lockfile, TypeScript configuration, absent platform/config declarations, and both concrete production source files.
3. **PASS — effect classification:** the graph distinguishes direct, transitive, build-tool, runtime, and platform edges and preserves overlaps.
4. **PASS — live upstream authority:** React owner documentation and live npm package metadata establish the stable line, exact current patches, migration requirements, peers, transitive scheduler range, type levels, and Vite Node range; applicability is explained beside each edge.
5. **PASS — changes and uncertainty:** removed APIs, function `defaultProps`, JSX transform, error reporting, peer requirements, nullable root, absent callers, incomplete lock, missing entry, and environment unknowns are separated without compatibility overstatement.
6. **PASS — dependency ordering:** environment qualification precedes dual-compatible source ownership, the 18.3 warning bridge precedes React 19 resolution, and release proof follows validated graph changes.
7. **PASS — safety controls:** every stage has evidence gates, rollback coverage, and stop conditions; exclusions and a smallest safe pre-dependency stage are explicit.
8. **PASS — evidence labels and boundary:** project facts, upstream facts, implications, and unknowns are labeled; no fixture file or dependency state was changed.

## Cleanup ownership

The dependency-upgrade-planner evaluation campaign owner owns retention and later cleanup of this trial directory. Keep these four evidence artifacts while the description campaign is active; remove them only through deliberate evaluation cleanup after their downstream comparison value ends.

## Result

**PASS**
