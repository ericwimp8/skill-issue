# Vite 5.4 to Vite 6 Migration Plan

## Upgrade Contract

- **Requested change — observed task:** Plan the migration of `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/vite-library/` from direct development dependency Vite `5.4.21` to the Vite 6 major line.
- **Authority boundary — observed task:** Planning and read-only assessment only. No manifest edit, dependency installation, lockfile regeneration, codemod, build, or deployment was performed.
- **Target precision — unknown:** The exact Vite 6 patch is unspecified. Select and record an exact 6.x release before implementation; resolve its package metadata and release notes rather than assuming every Vite 6 patch has the same transitive graph.
- **Delivery constraint — observed project fact:** The library builds through `vite build`, uses library mode, imports SCSS into its entry point, selects Terser minification, exposes `./dist/style.css` to consumers, and records Node 21 for both the package and deployment surfaces.
- **Evidence inspected:** `package.json`, `package-lock.json`, `vite.config.ts`, `package-consumer.json`, `src/index.ts`, and `src/theme.scss` under the fixture directory.

## Current Graph and Concrete Usage

| Relationship | Classification | Observed project fact | Migration significance |
| --- | --- | --- | --- |
| `synthetic-widget-library` -> Vite `5.4.21` | Direct, build-tool | `package.json:6-7` runs `vite build` and pins Vite; `package-lock.json` resolves `5.4.21` | The requested major upgrade changes the production packaging tool and generated artifacts. |
| Vite -> Rollup `^4.20.0`, resolved `4.24.0` | Transitive, build-tool | `package-lock.json` records this sole Vite dependency edge | A lock refresh may select a different Rollup graph for the exact Vite 6 patch; the supplied synthetic lockfile is incomplete and cannot establish the full future graph. |
| Library build -> Terser `5.15.1` | Direct, build-tool | `vite.config.ts:6` sets `minify: "terser"`; manifest and lockfile pin `5.15.1` | Vite 6 raises the minimum supported Terser version to `5.16.0`. |
| Library build -> Sass `1.77.8` | Direct, build-tool | `src/index.ts:1` imports `theme.scss`; `vite.config.ts:5` explicitly selects Sass's legacy API | Vite 6 defaults to the modern Sass API, although this explicit legacy setting retains old behavior for Vite 6 and carries a later-removal risk. |
| Resolver -> custom `widget` condition | Direct configuration, build-tool | `vite.config.ts:4` replaces `resolve.conditions` with `['widget']` | Vite 6 requires custom condition arrays to include the exported default client conditions that Vite 5 added internally. |
| Library mode -> consumer CSS export | Direct packaging and deployment contract | `vite.config.ts:6` enables library mode; `package-consumer.json:2` exports `./style.css` from `./dist/style.css` | Vite 6 changes library-mode CSS's default filename to the package name, which would break this declared consumer path unless the filename or export changes deliberately. |
| Package/deployment -> Node 21 | Runtime and platform | `package.json:5` requires `21.x`; `package-consumer.json:3` records `21.7.3` | Vite 6.0.0's engine range excludes Node 21, and Node 21 is end-of-life. Runtime alignment is a prerequisite, even if Vite executes only during the build in the eventual deployment design. |
| Source API | Runtime output | `src/index.ts:2` exports only `widgetVersion`; `theme.scss` defines one CSS rule | No Vite runtime API or plugin usage is observed. Generated JS/CSS names and contents still require comparison because consumers depend on build output. |

The lockfile does not contain package entries for the declared Sass and TypeScript dependencies and contains only abbreviated Vite, Rollup, and Terser entries. **Unknown:** it is unsuitable as proof of a complete reproducible installation. Preserve this limitation when reviewing lockfile diffs.

## Applicable Upstream Requirements

| Upstream fact | Version applicability | Project implication |
| --- | --- | --- |
| Vite 6.0.0 declares Node `^18.0.0 || ^20.0.0 || >=22.0.0`. [Vite 6.0.0 package metadata](https://github.com/vitejs/vite/blob/v6.0.0/packages/vite/package.json#L60-L62) | Exact to the first Vite 6 release; recheck the chosen 6.x patch | **Reasoned implication:** Node `21.x` and deployment Node `21.7.3` do not satisfy this range. |
| With custom `resolve.conditions`, Vite 6 requires adding `module`, `browser`, and `development|production`, available through `defaultClientConditions`. [Vite 6 migration guide](https://v6.vite.dev/guide/migration.html#default-value-for-resolve-conditions) | Vite 5 -> 6 and directly applicable because the fixture configures the option | **Reasoned implication:** retain `widget` and add Vite's client defaults before trusting conditional-export resolution. |
| Vite 6 uses Sass's modern API by default; explicit `scss.api: 'legacy'` remains supported in Vite 6 but is scheduled for removal in Vite 7. [Vite 6 migration guide](https://v6.vite.dev/guide/migration.html#sass-now-uses-modern-api-by-default) | Vite 5 -> 6; directly applicable to the explicit SCSS configuration | **Reasoned implication:** legacy mode can isolate the Vite 6 upgrade, but modern-mode validation on Vite 5.4 is the cleaner prerequisite and removes a known future stop. |
| In Vite 6 library mode, CSS defaults to a name derived from `package.json`; `build.lib.cssFileName` can preserve another filename. [Vite 6 migration guide](https://v6.vite.dev/guide/migration.html#customize-css-output-file-name-in-library-mode) | Vite 5 -> 6 and directly applicable to this library build | **Reasoned implication:** preserve the public `./style.css` export by setting `cssFileName: 'style'`, unless owners intentionally version a consumer-facing filename change. |
| Vite 6 raises the minimum Terser version for `build.minify: 'terser'` to `5.16.0`; Vite requires Terser to be installed when selected. [Vite 6 migration guide](https://v6.vite.dev/guide/migration.html#advanced), [Vite 6 build options](https://v6.vite.dev/config/build-options.html#build-minify) | Vite 5 -> 6 and directly applicable to `minify: 'terser'` | **Reasoned implication:** Terser must be upgraded with or before Vite; the current `5.15.1` is below the supported floor. |
| Node 21 reached end-of-life on April 10, 2024 and no longer receives security fixes. [Node.js EOL policy and table](https://nodejs.org/en/about/eol) | Directly applicable to the declared and recorded deployment runtime | **Reasoned implication:** choose a supported Node line accepted by the exact Vite 6 target rather than carrying Node 21 forward. |

No production use of Vite's experimental Runtime API, SSR, PostCSS configuration, JSON transforms, proxying, HTML transforms, glob syntax, or CommonJS entry points was found. Their Vite 6 migration notes do not create tasks for this fixture unless broader source or deployment evidence appears.

## Impact and Risk Decisions

1. **Platform prerequisite — blocking:** Node 21 conflicts with Vite 6. Choose one supported Node version for local builds, CI, packaging, and deployment; update both declarations together during implementation. **Stop condition:** any required deployment platform cannot run the chosen Node line.
2. **Terser compatibility — blocking:** `5.15.1` is below Vite 6's supported minimum. Select a compatible Terser version and review its own release span before changing the Vite constraint. **Stop condition:** the selected Terser release changes required output semantics or cannot run on the chosen Node line.
3. **Conditional resolution — blocking correctness risk:** update the custom array to retain `widget` while including `defaultClientConditions`. **Unknown:** there is no fixture dependency with conditional exports, so the current fixture alone cannot prove which branch `widget` must win; validate against a representative package or fixture before release.
4. **Sass API — reversible decision:** prefer removing `api: 'legacy'` and proving modern-mode output while still on Vite 5.4, because Vite 5.4 supports the modern API. If modern mode exposes an unrecorded importer or option dependency, retain explicit legacy mode for the Vite 6 stage and record a separate Vite 7 prerequisite.
5. **CSS artifact contract — blocking deployment risk:** the consumer path requires `dist/style.css`. Preserve it in Vite 6 with `build.lib.cssFileName: 'style'` unless package owners approve a breaking export migration to the new package-derived filename. **Stop condition:** a candidate build omits `dist/style.css` or the export resolves to a missing file.
6. **Transitive graph — unknown until exact selection:** the abbreviated current lock cannot predict Vite 6's complete Rollup/esbuild/optional platform packages. Review the regenerated lock as a state-changing result, not as a mechanical formality.

## Dependency-Ordered Migration Stages

### Stage 0 — Freeze the Decisions and Baseline

**Behavior owners:** runtime policy, package contract, and release process.

1. Select the exact Vite 6 patch, a supported Node line accepted by that patch, and a Terser version at or above Vite's floor.
2. Decide that `./style.css` remains the stable consumer export, unless an explicitly coordinated breaking package release owns a rename.
3. On the unchanged Vite 5.4 tree in its currently supported environment, capture a clean build's file list, hashes or normalized content for JS and CSS, package export resolution, and command/runtime versions. The current assessment did not run this baseline.
4. Confirm where `deploymentNode` is enforced and enumerate CI/container/host runtime declarations absent from the fixture.

**Gate:** proceed only with an exact version set, an identified deployment owner, and retained baseline evidence. **Rollback:** no dependency state has changed; discard only newly captured disposable artifacts. **Stop:** missing deployment authority, inability to reproduce the current build, or evidence of additional uninspected manifests/configuration.

### Stage 1 — Prove the Reversible Source and Configuration Prerequisites

**Behavior owners:** Vite resolver configuration and Sass compilation configuration.

1. While Vite remains at `5.4.21`, trial Sass modern mode by removing the explicit legacy override; compare emitted CSS and warnings with Stage 0. Revert if modern behavior depends on unrecorded Sass options or importers.
2. Prepare the Vite 6 resolver form using `defaultClientConditions`, retaining `widget` before the defaults as shown by the migration guidance. Validate resolution with a representative conditional-export dependency; the current source has none.
3. Prepare the Vite 6 library configuration that sets `build.lib.cssFileName: 'style'` so the existing consumer export remains true.

**Gate:** modern Sass produces accepted CSS, or an explicit temporary legacy exception is approved; resolution proves both `widget` and normal browser/module conditions; the planned CSS filename matches the package contract. **Rollback:** restore only the configuration trial. **Stop:** resolution changes the selected package entry unexpectedly or CSS semantics differ without owner approval.

### Stage 2 — Align Runtime and Direct Compatibility Dependencies

**Behavior owners:** platform/runtime declarations and build-tool manifest.

1. Change the package engine and deployment runtime record to the single supported Node decision from Stage 0; update any discovered CI or container owners in the same stage.
2. Change Terser to the reviewed compatible version. Keep Vite at `5.4.21` for this checkpoint so runtime and minifier effects are isolated.
3. Regenerate dependency state with the repository's package-manager workflow only after both manifest decisions are explicit.

**Gate:** clean install under the chosen Node version; current-Vite production build succeeds; Terser-minified JS and CSS artifacts satisfy the Stage 0 comparison and consumer export probe. **Rollback:** restore manifests and lockfile together, reinstate the prior runtime only in the isolated baseline environment, and discard generated `node_modules`/`dist` state. **Stop:** install engine rejection, minifier failure, unexplained artifact change, or deployment platform rejection.

### Stage 3 — Upgrade Vite and Refresh Its Transitive Graph

**Behavior owners:** build-tool manifest, Vite configuration, and generated dependency state.

1. Apply the exact Vite 6 constraint selected in Stage 0 together with the prepared resolver defaults and stable CSS filename configuration.
2. Regenerate the lockfile once; inspect Vite, Rollup, esbuild/platform packages, and optional peer changes against the chosen release metadata. Resolve unexpected changes before building.
3. Run a clean production library build on the chosen Node line. Verify `dist/style.css` exists, the configured conditional-resolution probe selects the intended entries, and JS/CSS output remains semantically acceptable.

**Gate:** clean install and build, reviewed lock graph, expected artifact names, passing export-resolution probe, and accepted output comparison. **Rollback:** restore the Stage 2 Vite/configuration/lockfile snapshot together and discard the Vite 6 install/build outputs. **Stop:** missing or renamed consumer artifacts, unresolved conditional exports, unsupported peer/engine warnings, or unexplained transitive packages.

### Stage 4 — Validate Packaging and Deployment Before Release

**Behavior owners:** package consumer contract and deployment platform.

1. Pack the library from a clean checkout and verify the published file set contains the JS entry artifacts and `dist/style.css` addressed by `./style.css`.
2. Install or consume the packed artifact through the representative consumer path; import both the library entry and `./style.css` rather than reading the build directory directly.
3. Run the deployment/package job under the chosen Node runtime and confirm the deployed artifact comes from the reviewed clean build.
4. Retain exact Node/npm/Vite/Terser versions, lockfile identity, build logs, package contents, export checks, and deployment probe results as the release evidence.

**Gate:** package and deployment consumers resolve every declared export, and the supported runtime is observed end to end. **Rollback:** redeploy the last accepted Vite 5 artifact and restore its manifest/lock snapshot if the new package cannot be consumed. **Stop:** deployment still selects Node 21, published files differ from the reviewed package, or any export resolves to a missing artifact.

## Validation Coverage

| Gate | Required evidence |
| --- | --- |
| Static | Type-check the Vite configuration under the selected Vite 6 types; inspect engine and consumer-export declarations for one aligned runtime/artifact contract. |
| Dependency | Clean package-manager install, no unsupported engine/peer diagnostics, and human-reviewed Vite/Rollup/Terser/platform lockfile edges. |
| Build | `vite build` success from a clean checkout under the chosen Node line; retained artifact list and semantic comparison against the Vite 5 baseline. |
| Resolver | A representative conditional-export package proves `widget` precedence and normal module/browser/development-or-production fallback. |
| Style | Modern-Sass compilation result or an explicitly approved legacy exception; emitted `.widget` rule remains acceptable. |
| Package | Packed artifact contains the files addressed by declared exports; a consumer resolves `./style.css`. |
| Deployment | The actual packaging/deployment path observes the chosen supported Node version and deploys the reviewed package identity. |

No test suite, CI workflow, deployment script, browser support policy, or real consumer package is present in the fixture. Those are **unknown validation surfaces**, not inferred passes.

## Exclusions and Follow-Up Questions

- No dependency or fixture file changes, installs, builds, tests, packaging, or deployment execution are part of this plan.
- No Vite 7/8 migration, framework/plugin migration, new deployment design, or public export rename is included.
- Which exact Vite 6 patch must be adopted, and is adopting an old major line still required by delivery policy?
- Which currently supported Node line can every build and deployment surface use, and where are those surfaces configured outside this fixture?
- Is `./style.css` a published compatibility promise? This plan preserves it unless owners explicitly approve a breaking release.
- Is there a real conditional-export dependency that uses `widget`, and must that condition take precedence over Vite's client defaults?
- Is temporary Sass legacy mode acceptable if the modern-mode preflight changes output, or must modern Sass be a hard prerequisite?
- What clean install/build/package commands and artifact-diff policy does the owning repository use?

## Smallest Safe First Stage

Complete Stage 0 only: choose the exact Vite 6, supported Node, and compatible Terser versions; confirm the stable CSS export and deployment owner; then capture a reproducible unchanged-Vite baseline. Do not change dependency state until those decisions and baseline artifacts exist. The evidence required to proceed is a successful clean Vite 5.4 build under its valid baseline runtime, a retained output/package inventory, and explicit ownership of the Node and CSS-export decisions.
