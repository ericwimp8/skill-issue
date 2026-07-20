# Vite 5.4 to Vite 6 Migration Plan

## Upgrade Contract

- **Observed project fact:** `fixtures/vite-library/package.json` directly pins Vite `5.4.21` as a development build tool; `package-lock.json` resolves the same version. The requested target is Vite 6, but an exact 6.x release is not specified.
- **Observed project fact:** The isolated workspace is a private ESM TypeScript library. Its only build command is `vite build`; implementation, installation, lockfile regeneration, and fixture mutation are outside this planning run.
- **Observed project fact:** `src/index.ts` imports `src/theme.scss`, and `vite.config.ts` owns library bundling, custom resolution conditions, SCSS processing, and Terser minification.
- **Constraint:** Upstream conclusions are limited to the two official Vite 6 routes listed in `fixtures/authoritative-sources.md`: the [Vite 5 migration guide](https://v6.vite.dev/guide/migration.html) and [Vite 6 release announcement](https://v6.vite.dev/blog/announcing-vite6).
- **Unknown:** The exact Vite 6.x and compatible Terser patch versions to pin, the available deployment runtime, and any CI or publishing workflow are absent from the permitted evidence.

## Current Dependency Graph and Usage

| Edge | Classification | Evidence and concrete effect |
| --- | --- | --- |
| library -> Vite `5.4.21` | direct, build-tool | `package.json` invokes `vite build`; `package-lock.json` resolves Vite `5.4.21`. |
| Vite -> Rollup `4.24.0` | transitive, build-tool | The active lockfile records Vite depending on Rollup `^4.20.0` and resolving `4.24.0`. No source imports Rollup directly. |
| library -> Terser `5.15.1` | direct, build-tool | `vite.config.ts` selects `build.minify: "terser"`; both manifest and lockfile pin `5.15.1`. |
| library -> Sass `1.77.8` | direct, build-tool | `src/index.ts` imports SCSS and the config explicitly requests the legacy SCSS API. |
| build/deployment -> Node `21.x` / `21.7.3` | runtime, platform | `package.json` declares Node `21.x`; `package-consumer.json` records deployment Node `21.7.3`. |
| built CSS -> consumer export | runtime packaging contract | `package-consumer.json` exports `./style.css` from `./dist/style.css`. |

TypeScript `5.4.5` is directly present for source compilation, but the permitted upstream sources establish no Vite 6 incompatibility for it. There are no plugins, overrides, resolutions, patches, registries, tests, generated artifacts, SSR entry points, PostCSS configuration, HTML entry points, or Vite Runtime API calls in the fixture.

## Applicable Upstream Requirements

| Source-backed statement | Classification | Project applicability |
| --- | --- | --- |
| Vite 6 supports Node 18, 20, and 22+, and drops Node 21 support. | authoritative upstream fact ([release announcement](https://v6.vite.dev/blog/announcing-vite6#node-js-support)) | The declared and deployment Node 21 lines are blockers that must be resolved before Vite 6 is installed or run. |
| A custom `resolve.conditions` value must include Vite 6's default client conditions, available as `defaultClientConditions`. | authoritative upstream fact ([migration guide](https://v6.vite.dev/guide/migration.html#default-value-for-resolve-conditions)) | `vite.config.ts` supplies only `["widget"]`; retaining that array unchanged would omit Vite 6 defaults. |
| Library-mode CSS defaults to a name based on the package name in Vite 6; `build.lib.cssFileName` can preserve an explicit name. | authoritative upstream fact ([migration guide](https://v6.vite.dev/guide/migration.html#customize-css-output-file-name-in-library-mode)) | The consumer contract still points to `dist/style.css`, while the package name is `synthetic-widget-library`; the output name must be made explicit or the consumer contract deliberately changed. |
| Vite 6 uses Sass's modern API by default, but explicit `scss.api: "legacy"` remains supported until Vite 7. | authoritative upstream fact ([migration guide](https://v6.vite.dev/guide/migration.html#sass-now-uses-modern-api-by-default)) | The fixture explicitly selects the legacy API, so Vite 6 does not require a Sass API migration. Retaining it is the smallest scoped Vite 6 plan. |
| Terser must be at least `5.16.0` when Vite uses `build.minify: "terser"`. | authoritative upstream fact ([migration guide](https://v6.vite.dev/guide/migration.html#advanced)) | The direct `5.15.1` pin is incompatible with the configured minifier path and must move before a Vite 6 build gate can pass. |

The Environment API, Runtime-to-Module-Runner change, JSON options, HTML asset handling, PostCSS config loading, SSR CSS behavior, proxy changes, CommonJS entry behavior, glob changes, and SSR resolution changes have no concrete fixture usage and produce no migration task.

## Impact Matrix

| Surface | Effect | Required disposition |
| --- | --- | --- |
| Direct dependency | Vite moves from exact `5.4.21` to one approved exact 6.x release. | Select the exact 6.x version before implementation; preserve the existing exact-pin policy. |
| Direct companion dependency | Terser `5.15.1` is below Vite 6's minimum for the configured minifier. | Select and pin an exact Terser version `>=5.16.0`. |
| Transitive dependency | Rollup is selected through Vite and currently resolves to `4.24.0`. | Let the package manager resolve Vite 6's declared graph; do not independently pin Rollup without new evidence. |
| Build configuration | Custom `resolve.conditions` would replace required client defaults. | Import `defaultClientConditions` from Vite 6 and configure `["widget", ...defaultClientConditions]`. |
| Packaging | Vite 6's library CSS default conflicts with the observed `dist/style.css` export. | Preserve the existing consumer contract by adding `build.lib.cssFileName: "style"`. |
| Sass build path | Explicit legacy mode bypasses the new modern default. | Retain `scss.api: "legacy"` for this migration and record its Vite 7 removal as future work. |
| Runtime/platform | Node 21 is unsupported. | Move both declared engine and deployment runtime to one available supported line: Node 20 or Node 22+. |

## Dependency-Ordered Migration Stages

### Stage 0 — Resolve Version and Runtime Decisions

**Behavior owner:** release/deployment decision makers.

1. Choose a deployment runtime supported by Vite 6 and available in the real environment: Node 20 or Node 22+. Update planning authority must choose one line before any dependency edit; Node 21 cannot remain.
2. Select an exact Vite 6.x release and an exact Terser release `>=5.16.0`, preserving the fixture's exact-pin convention. Because registry metadata is outside the permitted evidence, this selection remains blocked in this plan.

**Gate:** Record the selected Node line and exact package versions, plus evidence that the build and deployment environments can provide that Node line.

**Stop condition:** Stop if deployment cannot leave Node 21, the exact packages cannot be authorized, or their metadata introduces requirements not covered by this plan.

**Rollback point:** No repository state changes occur in this stage; discard the proposed selections.

### Stage 1 — Establish the Supported Runtime Baseline

**Behavior owners:** `package.json` for declared engine policy and the deployment configuration represented by `package-consumer.json` for runtime selection.

1. Change `engines.node` from `21.x` to the selected supported line.
2. Change `deploymentNode` from `21.7.3` to an available version on that same supported line.

**Gate:** On that runtime, record `node --version`, package-manager version, and a clean Vite 5.4.21 baseline `npm run build`. Confirm the build produces the currently contracted `dist/style.css`. This is implementation-time evidence; this planning run claims no successful build.

**Stop condition:** Stop if the supported runtime cannot install the unchanged lockfile or reproduce the Vite 5 baseline. Diagnose that runtime transition separately before adding the major upgrade.

**Rollback point:** Restore the two Node declarations to `21.x` and `21.7.3`; discard any generated installation or build state.

### Stage 2 — Apply the Version-Coupled Vite 6 Configuration and Manifest Changes

**Behavior owners:** `vite.config.ts` for build behavior and `package.json` for direct dependency policy.

Treat the Vite import/config edits and the Vite manifest bump as one reversible checkpoint because `defaultClientConditions` is a Vite 6 API.

1. In `vite.config.ts`, import `defaultClientConditions` alongside `defineConfig` and change client conditions to `["widget", ...defaultClientConditions]`.
2. In `vite.config.ts`, set `build.lib.cssFileName` to `"style"` so the established consumer export continues to resolve to `dist/style.css`.
3. Retain `css.preprocessorOptions.scss.api: "legacy"`; migrating Sass APIs is not required for Vite 6 and lacks permitted Sass-source evidence.
4. In `package.json`, replace Vite `5.4.21` with the approved exact 6.x version and replace Terser `5.15.1` with the approved exact version `>=5.16.0`.

**Gate:** Review the diff before package-manager operations. Every edit must map to the Node, conditions, CSS-name, Vite, or Terser edges above; Sass, TypeScript, source exports, and consumer paths remain unchanged.

**Stop condition:** Stop if the selected Vite release does not export `defaultClientConditions`, the consumer intentionally wants the new package-derived CSS filename, or additional direct dependency changes appear necessary without authoritative evidence.

**Rollback point:** Restore `vite.config.ts` and `package.json` to their Stage 1 checkpoint.

### Stage 3 — Refresh Dependency State

**Behavior owner:** the package manager and `package-lock.json`.

1. Regenerate the lockfile from the approved manifest on the supported Node runtime; do not hand-edit transitive Rollup state.
2. Inspect the resulting graph to confirm the root resolves the exact selected Vite and Terser versions, Terser is `>=5.16.0`, and Rollup is derived from Vite's declared dependency range.
3. Reject unrelated dependency churn unless the package manager demonstrates it is required by the selected direct versions.

**Gate:** Save the manifest/lockfile diff and dependency-tree evidence for Vite, Terser, Rollup, Sass, and TypeScript. Run the package manager's lockfile-consistency/install check without changing the approved versions.

**Stop condition:** Stop on peer conflicts, engine warnings, an unexpected registry, unexplained graph churn, or a resolved direct version different from the approved exact pin.

**Rollback point:** Restore `package.json` and `package-lock.json` from the Stage 1 checkpoint, restore `vite.config.ts`, and discard installed modules and generated package-manager state.

### Stage 4 — Validate Build and Published Contract

**Behavior owners:** Vite library build output and the consumer export contract.

1. Run the focused static/configuration check available to the project, then `npm run build` on the selected supported Node runtime.
2. Confirm the build emits JavaScript library output and exactly `dist/style.css`; verify the consumer export `./style.css -> ./dist/style.css` resolves to the emitted file.
3. Confirm the SCSS import is compiled and `.widget` retains the expected accent color semantics.
4. Confirm the JavaScript output is minified through the configured Terser path and the library still exports `widgetVersion` with value `"1"`.
5. Add a focused resolution probe that exercises a dependency with standard client conditions plus the custom `widget` condition; a build that never traverses such a package is insufficient evidence for the conditions change.

**Gate:** Retain command output, artifact inventory, export-resolution result, CSS content check, JavaScript export probe, and resolution-condition probe. There is no test suite or release command in the fixture, so none is claimed.

**Stop condition:** Stop on a build error, missing or renamed CSS, broken export resolution, changed public export, failed minification, Sass failure, or condition-resolution regression.

**Rollback point:** Restore all Stage 1 source/configuration, manifest, and lockfile checkpoints; discard `dist`, installed modules, and other generated state; reinstall the restored lockfile on the supported runtime.

## Risks, Exclusions, and Follow-Up Questions

- **Unknown:** Which supported Node line is available for deployment: Node 20 or Node 22+?
- **Unknown:** Which exact Vite 6.x and Terser `>=5.16.0` versions are approved? Their package metadata and full resolved graph must be reviewed during implementation.
- **Reasoned implication:** Preserving `style.css` in Vite configuration is safer than changing the observed consumer export because the fixture exposes that path as a contract. If product intent is to adopt `synthetic-widget-library.css`, revise both the config decision and consumer contract together.
- **Explicit exclusion:** A Sass modern-API migration is deferred. The fixture's explicit legacy setting remains valid in Vite 6; Vite 7 removal is a separate upgrade concern.
- **Explicit exclusion:** No task is created for TypeScript, Rollup, JSON, HTML, PostCSS, SSR, proxy, CommonJS, globs, Environment API, or Module Runner without concrete fixture usage or applicable compatibility evidence.
- **Inspection limitation:** CI, registry configuration, publisher configuration, deployment tooling, runtime availability, package metadata, and consumer source are absent from the permitted evidence.

## Smallest Safe First Stage

Begin with Stage 0 only: choose one Vite-6-supported Node line and exact Vite 6.x and Terser `>=5.16.0` pins. Implementation must not begin until there is recorded evidence that the real build and deployment environments provide the chosen Node line and the selected package metadata introduces no additional prerequisite.
