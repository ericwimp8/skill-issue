# Vite 6 Adoption Plan for the Synthetic Widget Library

## Upgrade contract

- **Observed project fact:** The workspace is `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/vite-library/`. It directly pins Vite `5.4.21` and requests the Vite 6 line; an exact Vite 6 patch is not selected yet (`package.json:7`, `package-lock.json:5-6`).
- **Observed project fact:** This is an npm lockfile-v3, TypeScript library build. Vite is a direct development/build-tool dependency, while Rollup `4.24.0` is currently transitive through Vite (`package-lock.json:3-8`).
- **Constraint:** This document is planning-only. No manifest, lockfile, configuration, source, generated artifact, install state, or deployment declaration may change without separate implementation authorization.
- **Unknown:** The repository does not identify the exact Vite 6 patch, the replacement Node release, or whether preserving the existing `style.css` consumer contract is preferred over changing that contract. Those choices must be settled at the gates below.

## Current graph and concrete use

| Relationship | Evidence | Vite 6 relevance |
| --- | --- | --- |
| Direct build tool | `vite` `5.4.21` and `vite build` in `package.json:6-7` | Requested dependency change and build entry point. |
| Direct build dependency | `terser` `5.15.1` in `package.json:7`; `build.minify: "terser"` in `vite.config.ts:6` | Installed version is below Vite 6's documented minimum for this configured minifier. |
| Direct stylesheet compiler | `sass` `1.77.8` in `package.json:7`; SCSS API forced to `legacy` in `vite.config.ts:5` | Vite 6 changes the default to the modern API, while allowing the explicit legacy setting for Vite 6. |
| Transitive build tool | Vite currently selects Rollup `4.24.0` in `package-lock.json:6-7` | The authorized lock refresh should select the Vite 6-compatible Rollup; Rollup must not be upgraded independently without an evidenced reason. |
| Platform | Node `21.x` in `package.json:5`; deployed Node `21.7.3` in `package-consumer.json:3` | Vite 6 drops Node 21 support, so platform ownership must be resolved first. |
| Resolution configuration | Custom `resolve.conditions: ["widget"]` in `vite.config.ts:4` | Vite 6 requires custom lists to include conditions that previously were added internally. |
| Produced artifact and consumer | `src/index.ts:1` imports `theme.scss`; `package-consumer.json:2` exports `./dist/style.css` | Vite 6 library CSS naming can change the emitted filename and break this export. |

## Applicable upstream requirements

- **Authoritative upstream fact:** Vite 6 supports Node 18, 20, and 22+ and drops Node 21. This applies directly to both Node 21 declarations in the fixture. [Vite 6 announcement — Node.js support](https://v6.vite.dev/blog/announcing-vite6#node-js-support)
- **Authoritative upstream fact:** With a custom `resolve.conditions`, Vite 6 no longer adds `module`, `browser`, and `development|production` internally; custom client conditions must include the new defaults. Vite exports `defaultClientConditions` for this purpose. This applies because the fixture replaces the list with `['widget']`. [Vite 6 migration — default value for resolve.conditions](https://v6.vite.dev/guide/migration.html#default-value-for-resolve-conditions)
- **Authoritative upstream fact:** Vite 6 uses Sass's modern API by default, still permits an explicit `scss.api: 'legacy'`, and warns that legacy support is removed in Vite 7. The existing explicit setting is therefore a deliberate compatibility choice, not an immediate Vite 6 blocker. [Vite 6 migration — Sass modern API](https://v6.vite.dev/guide/migration.html#sass-now-uses-modern-api-by-default)
- **Authoritative upstream fact:** Vite 6 library CSS defaults to a name derived from the package/library filename contract; projects relying on `style.css` can set `build.lib.cssFileName: 'style'`. This applies to the consumer export of `./dist/style.css`. [Vite 6 migration — library CSS filename](https://v6.vite.dev/guide/migration.html#customize-css-output-file-name-in-library-mode)
- **Authoritative upstream fact:** Vite 6 raises the minimum supported Terser version for `build.minify: 'terser'` to `5.16.0`. The fixture pins `5.15.1`, so Terser is a prerequisite direct upgrade. [Vite 6 migration — advanced changes](https://v6.vite.dev/guide/migration.html#advanced)

## Dependency-ordered migration stages

### Stage 1 — Decide the supported Node platform

**Owner:** Runtime/deployment policy represented by `package.json` and `package-consumer.json`.

1. Select one exact Node release line supported by Vite 6 and available in both build and deployment environments.
2. Plan aligned updates to `engines.node` and `deploymentNode`; do not leave build policy and deployed runtime on different unsupported Node 21 declarations.

**Evidence gate:** Record the selected Node line, Vite 6 support evidence, and proof that the build runner and deployment platform can provision it.

**Stop condition:** Stop before selecting Vite 6 if the deployment platform cannot leave Node 21, or if build and deployment owners cannot agree on one supported contract.

**Rollback:** Before implementation, retain the original `package.json` and `package-consumer.json`. If the platform probe fails, restore both Node 21 declarations together and do not proceed to dependency resolution.

### Stage 2 — Settle configuration and artifact contracts

**Owner:** `vite.config.ts` owns build behavior; `package-consumer.json` owns the consumer-facing CSS export.

1. Preserve the custom `widget` condition while adding Vite 6's client defaults, preferably using Vite's exported `defaultClientConditions` rather than copying an unowned list.
2. Choose the Sass path. The smallest Vite 6 adoption scope may retain the explicit legacy API temporarily, but the decision must be recorded and Sass output must be gated; migration to the modern API is separate work with its own output comparison.
3. Preserve the current consumer contract by planning `build.lib.cssFileName: 'style'`, unless the consumer owner explicitly chooses a coordinated export change instead.
4. Select an exact Terser version at or above `5.16.0` and an exact Vite 6 patch using authoritative package metadata at implementation time.

**Evidence gate:** A reviewed change set must show one owner for each condition list, Sass API choice, CSS filename, and consumer export. Every proposed dependency version must satisfy the documented Vite 6 edge.

**Stop condition:** Stop if the CSS producer and consumer names disagree, if standard client conditions would still be omitted, or if the chosen Terser remains below `5.16.0`.

**Rollback:** Restore `vite.config.ts`, `package.json`, and `package-consumer.json` as one configuration boundary. Discard any generated `dist/` produced by a rejected configuration.

### Stage 3 — Apply direct changes and refresh resolution

**Owner:** npm manifest and lockfile state.

After separate authorization, update the aligned Node declarations, chosen Vite 6 and Terser versions, and reviewed Vite configuration. Then perform a clean npm resolution on the selected supported Node runtime. Allow npm/Vite to select Rollup transitively; do not edit the Rollup lock entry directly.

**Evidence gate:** The regenerated lockfile must resolve the chosen Vite 6 patch, Terser `>=5.16.0`, and Vite's compatible Rollup graph without unsupported-engine or peer-resolution failures. Review the lockfile diff before building.

**Stop condition:** Stop on resolution failure, an unsupported Node warning, an unexplained new package edge, or a Rollup override/direct pin that upstream evidence did not require.

**Rollback:** Restore the pre-stage manifests, `vite.config.ts`, `package-consumer.json`, and `package-lock.json`; remove disposable install and build state before retrying.

### Stage 4 — Validate build, artifacts, resolution, and deployment

**Owner:** The build pipeline proves produced artifacts; the deployment owner proves the runtime.

1. Run the configured `vite build` on the selected supported Node runtime.
2. Inspect `dist/` and prove that the CSS filename matches `package-consumer.json`, the JavaScript library artifacts exist, and the SCSS import emitted the expected `.widget` styling.
3. Exercise a representative conditional-export dependency or disposable resolution probe that distinguishes `widget` and Vite's standard client conditions. The current fixture contains no such dependency, so configuration inspection alone cannot prove runtime resolution behavior.
4. Run the packaged consumer/deployment check on the selected supported Node runtime and verify the exported stylesheet path resolves.

**Evidence gate:** Retain command identities, exit status, supported Node version, lockfile identity, artifact listing, CSS inspection, condition-resolution result, and consumer/deployment result.

**Stop condition:** Stop release adoption on any failed build, missing or renamed stylesheet, unresolved export, changed Sass semantics, unproven condition behavior, or deployment execution on Node 21.

**Rollback:** Restore the Stage 2 boundary, its lockfile, and previous dependency constraints; discard the Vite 6 `node_modules/` and `dist/`; rebuild the Vite 5 state only if rollback verification is authorized.

## Explicit exclusions and unresolved risks

- No dependency installation, lockfile regeneration, source/config edit, build, test, codemod, or deployment is performed by this plan.
- No successful Vite 6 resolution, build, artifact, conditional resolution, Sass output, consumer, or deployment compatibility is claimed.
- The exact Vite 6, Terser, and supported Node patch versions remain implementation-time decisions.
- TypeScript `5.4.5` and Sass `1.77.8` have no demonstrated incompatibility in the inspected Vite 6 sources; upgrading them is excluded unless resolution or a focused gate supplies new evidence.
- The fixture has no concrete conditional-export dependency, so the condition behavior requires a representative probe before release confidence.

## Smallest safe first stage

Decide and document the supported replacement for Node 21 across both `package.json` and `package-consumer.json`. Do not alter dependencies yet. Proceed only after evidence shows the same supported Node line is provisionable in build and deployment; otherwise Vite 6 adoption remains blocked.
