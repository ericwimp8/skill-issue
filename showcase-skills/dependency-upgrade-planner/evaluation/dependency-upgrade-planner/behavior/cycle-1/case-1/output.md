# React 19 Migration Plan

## Upgrade Contract

- **Observed project fact:** The bounded workspace is `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/react-app/`. Its production manifest pins `react` and `react-dom` to `18.2.0`, and its lockfile resolves both to `18.2.0`.
- **Requested target:** Stable React 19 for the app, including the matching React DOM runtime and TypeScript declaration packages.
- **Authoritative upstream fact:** React announced React 19 as stable on 2024-12-05. The React 19 Upgrade Guide instructs TypeScript projects to upgrade `react`, `react-dom`, `@types/react`, and `@types/react-dom` together to the React 19 line. Sources: [React 19 stable release](https://react.dev/blog/2024/12/05/react-19), [React 19 Upgrade Guide](https://react.dev/blog/2024/04/25/react-19-upgrade-guide).
- **Constraint:** This document is planning-only. No fixture file, package constraint, installed dependency, lockfile resolution, generated artifact, or build output was changed.
- **Unknown:** The request identifies the stable React 19 major line but does not select a concrete 19.x patch. Because the allowed source ledger supplies no registry metadata, implementation must record the resolved stable 19.x versions and keep each runtime/type pair aligned rather than assume a patch.

## Inspected Evidence

- **Observed project fact:** `package.json` is the only production manifest. It defines a single `tsc && vite build` script, exact React 18.2 runtime constraints, React 18 declaration packages, TypeScript 5.4.5, and Vite 5.4.21.
- **Observed project fact:** `package-lock.json` is lockfile version 3. It records direct React and React DOM resolutions and the `react-dom -> scheduler ^0.23.0 -> scheduler 0.23.2` transitive path. The synthetic lockfile does not contain entries for the declared type packages, TypeScript, or Vite, so it cannot prove their resolved versions or complete dependency trees.
- **Observed project fact:** `tsconfig.json` enables strict TypeScript and the `react-jsx` JSX runtime for `src/`.
- **Observed project fact:** `src/main.tsx` imports the legacy `render` API from `react-dom` and mounts into `document.getElementById("root")`.
- **Observed project fact:** `src/legacy-panel.tsx` uses `React.createFactory` and assigns `defaultProps` to a function component.
- **Observed project fact:** The fixture contains only the five inspected files. It declares no workspace, overrides, resolutions, patches, registry, package-manager version, Node engine, browser support, deployment configuration, generated code, test command, or HTML entry document.

## Current Dependency Graph and Usage

```text
application runtime
├── react 18.2.0 (direct runtime; JSX/component APIs)
└── react-dom 18.2.0 (direct runtime; legacy DOM render entry point)
    └── scheduler ^0.23.0 -> 0.23.2 (transitive runtime in supplied lockfile)

build/type path
├── @types/react 18.2.79 (direct development/type dependency)
├── @types/react-dom 18.2.25 (direct development/type dependency)
├── typescript 5.4.5 (direct build tool)
└── vite 5.4.21 (direct build/package tool)
```

The app bootstrap in `src/main.tsx` owns DOM-root creation and rendering. `LegacyPanel` in `src/legacy-panel.tsx` owns its element construction and default-title behavior. `package.json` owns requested direct constraints; npm and `package-lock.json` own concrete resolutions and the transitive graph.

## Applicable React 19 Requirements

- **Authoritative upstream fact:** `ReactDOM.render` was removed in React 19; the guide requires `createRoot` from `react-dom/client` followed by `root.render(...)`. **Project implication:** `src/main.tsx` must migrate before the React 19 runtime is selected. The bootstrap must also define behavior when `#root` is absent so strict TypeScript and runtime behavior are explicit rather than relying on a nullable lookup. [Upgrade Guide: removed ReactDOM.render](https://react.dev/blog/2024/04/25/react-19-upgrade-guide#removed-reactdom-render).
- **Authoritative upstream fact:** `React.createFactory` was removed and should be replaced with JSX. **Project implication:** `LegacyPanel` should return `<section>...</section>` directly. [Upgrade Guide: removed React.createFactory](https://react.dev/blog/2024/04/25/react-19-upgrade-guide#removed-createfactory).
- **Authoritative upstream fact:** Function-component `defaultProps` support was removed in favor of ES default parameters. **Project implication:** `LegacyPanel` should destructure `title = "Overview"` and remove the `LegacyPanel.defaultProps` assignment while retaining the current nullish behavior decision explicitly. [Upgrade Guide: function defaultProps](https://react.dev/blog/2024/04/25/react-19-upgrade-guide#removed-proptypes-and-defaultprops).
- **Authoritative upstream fact:** TypeScript projects need React 19 versions of both React declaration packages; the guide also documents broader React 19 type changes and a separate types codemod. **Project implication:** upgrade the declaration packages with the runtime pair, then let strict compilation reveal fixture-specific type work. No inspected source uses the guide's changed `useRef`, `ReactElement`, global `JSX`, or `useReducer` surfaces, so those changes do not become speculative migration tasks. [Upgrade Guide: TypeScript changes](https://react.dev/blog/2024/04/25/react-19-upgrade-guide#typescript-changes).
- **Authoritative upstream fact:** React 19 changes render-error reporting and provides root error callbacks for custom reporting. **Project implication:** no error reporter or error boundary exists in the fixture, so there is no concrete reporting integration to migrate. Record this as an excluded behavior review unless additional production files appear. [Upgrade Guide: render errors](https://react.dev/blog/2024/04/25/react-19-upgrade-guide#errors-in-render-are-not-re-thrown).

## Impact Matrix

| Effect | Evidence | Required treatment |
| --- | --- | --- |
| Direct runtime | `react` and `react-dom` are exact direct dependencies; both are imported in `src/` | Upgrade as one compatible React 19 pair after source compatibility work |
| Direct type/build | React 18 declaration packages are exact dev dependencies; strict TypeScript compiles `src/` | Upgrade both declaration packages with the runtime pair and run strict compilation |
| Transitive runtime | The lockfile records React DOM's `scheduler` edge | Let npm resolve the React 19 transitive graph; inspect the diff and reject unrelated churn |
| Build tool | TypeScript 5.4.5 and Vite 5.4.21 implement the only build pipeline | Keep constraints unchanged initially; validate compilation and bundling after lock refresh |
| Runtime | Browser DOM startup calls a removed renderer and assumes an element with id `root` | Move bootstrap to `createRoot`, define missing-container behavior, and perform a browser smoke probe when an HTML host is available |
| Platform | No Node engine, package-manager version, browser matrix, HTML host, CI, or deploy surface is supplied | Treat environment compatibility and production release validation as unknown; do not claim them from this fixture |

## Dependency-Ordered Migration Stages

### Stage 0 — Preserve a Reproducible Baseline

**Owner:** implementer operating the fixture; no file owner changes.

1. Record the current hashes of all five fixture files and the actual Node/npm versions intended for implementation.
2. If dependencies already exist outside the supplied fixture state, run `npm run build` and retain the complete output. Do not install merely to manufacture a baseline.
3. If no runnable installed state exists, record the baseline build as unavailable. If the build reaches Vite and fails because the supplied fixture has no HTML entry, retain that as a known fixture limitation rather than a React failure.

**Gate:** Proceed only with clean fixture hashes and an explicit baseline status (pass, fail with evidence, or unavailable).

**Rollback:** None; this stage is read-only.

### Stage 1 — Remove APIs Already Incompatible with React 19

**Owners:** `src/main.tsx` for root startup; `src/legacy-panel.tsx` for component construction and defaults.

1. In `src/main.tsx`, replace the `react-dom` `render` import and call with `createRoot` from `react-dom/client`, an explicit `#root` container check, and `root.render(<ConsoleApp />)`.
2. In `src/legacy-panel.tsx`, replace `React.createFactory("section")` with JSX.
3. Express the title default in the function parameter and remove the function `defaultProps` assignment. Preserve deliberate handling of an explicitly supplied `null` only if the component's public type is widened to allow it; the current type accepts `string | undefined`.

**Validation gate:** Run strict TypeScript compilation against the unchanged React 18 dependency state when a runnable installed state exists. Inspect the diff to confirm only the two behavior owners changed and the displayed `Ready`/`Overview` values remain represented.

**Rollback:** Restore both source files to their recorded Stage 0 hashes. Stop if the bootstrap's missing-container behavior cannot be agreed or the compatibility edit changes user-visible semantics.

### Stage 2 — Select the React 19 Runtime and Type Set

**Owners:** `package.json` for direct constraints; npm/`package-lock.json` for resolved state.

1. Resolve a concrete stable React 19.x patch through the implementation environment's configured npm registry, recording the selection time and registry evidence.
2. Update `react` and `react-dom` together to that React 19 line. Update `@types/react` and `@types/react-dom` together to compatible React 19 declarations as required by the guide. Preserve exact-pin policy unless the maintainer explicitly changes it.
3. Regenerate `package-lock.json` with the repository's selected npm version. Review the complete lockfile diff, including React DOM's new transitive dependencies and integrity metadata. The supplied lockfile is incomplete relative to the manifest, so stop rather than accepting unexplained unrelated reconstruction.
4. Leave TypeScript and Vite constraints unchanged unless Stage 3 produces evidence of a concrete incompatibility. A React major upgrade alone is not evidence that either tool needs a major upgrade.

**Validation gate:** Confirm manifest/runtime/type alignment, a lockfile generated by the selected package manager, no React 18 resolution remaining in the effective graph, and no unrequested direct dependency changes.

**Rollback:** Restore `package.json` and `package-lock.json` to their Stage 0 hashes and discard installed-state changes produced by the attempted resolution. Stop on peer-resolution errors, unavailable stable patch evidence, unexplained lockfile churn, or a requirement to upgrade an unscoped toolchain.

### Stage 3 — Validate Compile, Bundle, and Runtime Behavior

**Owners:** TypeScript for source/type compatibility, Vite for bundling, browser host for runtime behavior.

1. Run strict TypeScript compilation first; retain diagnostics and resolve only errors traced to the React 19 runtime/type change.
2. Run the declared `npm run build`; retain the TypeScript and Vite output separately. If Vite cannot build because the fixture lacks an HTML entry, classify bundling as unvalidated and stop before release claims.
3. When a valid HTML host with `#root` is supplied, load the built app and verify that `Ready` mounts once without console errors. Exercise `LegacyPanel` with an omitted title and an explicit title, verifying `Overview` and the supplied title respectively.
4. Inspect runtime console/error-reporting behavior. Add React 19 root error callbacks only if a concrete application reporting owner and expected behavior are supplied.

**Gate:** Completion requires clean strict compilation, a successful declared build from a complete host fixture, and the two focused browser observations. Missing host/platform evidence leaves the migration technically blocked rather than implicitly successful.

**Rollback:** Restore Stage 1 source files plus Stage 2 manifest and lockfile, then restore installed state from the React 18 lock. Roll back on type failures without a bounded source fix, startup failure, duplicate mount, changed panel defaults, or newly observed uncaught errors.

## Risks, Stops, and Exclusions

- **Unknown:** Exact React 19.x and declaration-package patch versions remain implementation-time selections because registry metadata was outside the supplied authoritative routes.
- **Unknown:** The incomplete synthetic lockfile cannot establish all current resolutions or predict the full regenerated diff.
- **Unknown:** Node/npm versions, browser targets, CI, deployment packaging, and the HTML mount host are absent. Release readiness cannot be established until those owners provide evidence.
- **Stop condition:** Do not combine this work with Vite 6, TypeScript, Node, package-manager, or platform upgrades unless a concrete React 19 incompatibility is demonstrated and separately authorized.
- **Excluded:** React 19 feature adoption, server components, actions, refactoring unrelated source, broad codemods, testing-framework additions, deployment work, and changes to Pydantic/FastAPI or the other fixtures.
- **Follow-up question:** Which exact stable React 19.x pin and npm version should implementation use, and where is the HTML/runtime host that owns the `#root` element?

## Smallest Safe First Stage

Perform Stage 0 only: capture clean fixture hashes, the intended Node/npm environment, and any already-runnable baseline build output without installing or editing. Implementation may begin only after that evidence is retained and the exact React 19.x pin plus runtime host are identified.
