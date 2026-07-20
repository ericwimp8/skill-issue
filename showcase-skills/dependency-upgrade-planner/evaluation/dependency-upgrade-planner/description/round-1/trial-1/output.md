# React 18.2 to Stable React 19 Upgrade Plan

## Upgrade contract

- **Requested change:** move the synthetic application from exact `react@18.2.0` and `react-dom@18.2.0` pins to the stable React 19 line. This is a review plan; dependency edits, installs, lockfile regeneration, codemods, and source migration remain outside this task.
- **Authoritative upstream fact:** React documents 19.2 as its latest stable minor, while npm currently tags `react@19.2.7` and `react-dom@19.2.7` as `latest`. Preserve the fixture's exact-pin policy by proposing `19.2.7` for review; re-query npm immediately before implementation because the stable patch can advance. Sources: [React versions](https://react.dev/versions), [react latest metadata](https://registry.npmjs.org/react/latest), [react-dom latest metadata](https://registry.npmjs.org/react-dom/latest).
- **Scope:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/react-app/` only.
- **Exclusions:** React 19 feature adoption, Vite or TypeScript modernization, test-framework introduction, application redesign, and fixture repair unrelated to proving the React upgrade.

## Inspected project evidence

- **Observed project fact:** `package.json` directly pins `react` and `react-dom` to `18.2.0`; it directly pins React type packages, TypeScript 5.4.5, and Vite 5.4.21 as development/build dependencies. The only script is `tsc && vite build`.
- **Observed project fact:** `package-lock.json` is lockfile v3 and resolves `react@18.2.0`, `react-dom@18.2.0`, and transitive runtime package `scheduler@0.23.2`. It declares the dev dependencies in the root package but contains no resolved entries for them, so it is not evidence of a complete reproducible toolchain resolution.
- **Observed project fact:** `src/main.tsx` imports removed `render` from `react-dom` and mounts into `document.getElementById("root")` without proving that the container exists.
- **Observed project fact:** `src/legacy-panel.tsx` uses removed `React.createFactory` and assigns `defaultProps` to a function component. No production file imports `LegacyPanel`; TypeScript still includes it through `tsconfig.json`, so it remains a compile-time migration surface even though runtime reachability is unproven.
- **Observed project fact:** `tsconfig.json` uses `jsx: "react-jsx"`, which is the modern JSX transform React 19 requires. No workspace declaration, override, resolution, patch, registry override, Vite config, HTML entry, Node declaration, browser target, deployment config, or test configuration exists in the fixture.

## Current and target dependency graph

| Edge | Classification | Current evidence | Proposed target / consequence |
| --- | --- | --- | --- |
| app -> `react` | direct + runtime | exact `18.2.0` | exact `19.2.7`, subject to implementation-day registry recheck |
| app -> `react-dom` | direct + runtime | exact `18.2.0` | exact `19.2.7`; its npm peer requires `react@^19.2.7` |
| `react-dom` -> `scheduler` | transitive + runtime | `^0.23.0`, resolved `0.23.2` | npm metadata declares `^0.27.0`; let npm resolve and record it rather than editing it directly |
| app -> `@types/react` | direct + build-tool | exact `18.2.79` declared; unresolved in the lock | `19.2.17` currently latest; npm metadata reports TypeScript 5.3 definitions |
| app -> `@types/react-dom` | direct + build-tool | exact `18.2.25` declared; unresolved in the lock | `19.2.3` currently latest; peer requires `@types/react@^19.2.0` |
| app -> TypeScript | direct + build-tool | exact `5.4.5` | retain: it meets the current type packages' published TypeScript 5.3/5.2 levels; compilation must confirm |
| app -> Vite / Node | build-tool + platform | Vite `5.4.21`; Node unspecified | retain Vite; its package metadata requires Node `^18.0.0 || >=20.0.0`, so qualify the implementation environment first |

Type package facts come from [@types/react latest metadata](https://registry.npmjs.org/%40types%2freact/latest) and [@types/react-dom latest metadata](https://registry.npmjs.org/%40types%2freact-dom/latest). Vite's current fixture constraint comes from [vite 5.4.21 metadata](https://registry.npmjs.org/vite/5.4.21).

## Applicable breaking changes and implications

1. **Authoritative upstream fact:** React recommends an intermediate React 18.3 release, behaviorally matching 18.2 while warning about React 19 incompatibilities. It also requires the modern JSX transform. The fixture already has the transform, but the 18.3 warning bridge remains applicable. [React 19 upgrade guide](https://react.dev/blog/2024/04/25/react-19-upgrade-guide)
2. **Authoritative upstream fact:** React 19 removes `ReactDOM.render`; the owner-provided replacement is `createRoot` from `react-dom/client`. **Reasoned implication:** `src/main.tsx` must obtain and validate the root element, create the root, then render. The guard is necessary to satisfy the nullable DOM lookup under strict TypeScript rather than hiding absence with an assertion.
3. **Authoritative upstream fact:** React 19 removes `React.createFactory`; JSX is the replacement. **Reasoned implication:** `LegacyPanel` should directly return `<section>{...}</section>`.
4. **Authoritative upstream fact:** React 19 removes `defaultProps` behavior for function components in favor of default parameters. **Reasoned implication:** change the `LegacyPanel` parameter to `{ title = "Overview" }` and remove the assignment. Preserve the existing null behavior deliberately if callers may pass `null`; current production evidence shows no callers and therefore cannot settle that contract.
5. **Authoritative upstream fact:** React 19 changes render-error reporting: uncaught errors go to `window.reportError` and caught errors go to `console.error` instead of being rethrown. **Unknown:** the fixture contains no monitoring or error-boundary integration, so no migration task is justified; a browser error-path probe is required only if an external host supplies such integration.
6. **Observed project fact:** the fixture has no third-party React libraries or React-internals usage, so there are no evidenced library peer edges beyond `react-dom`. Do not infer ecosystem compatibility from this synthetic fixture.

## Dependency-ordered implementation stages

### Stage 0 — qualify a reproducible baseline

**Owner:** build environment and package state.

- Record Node and npm versions; stop if Node is outside Vite 5.4.21's declared range.
- In a disposable copy, verify whether the checked-in lock supports a clean install. Expect investigation because resolved dev-tool entries are absent; do not normalize the lock silently.
- Run `npx tsc --noEmit` using the fixture's installed current toolchain if one already exists. The checked-in build script emits JavaScript before Vite runs, so use `--noEmit` for the baseline evidence.
- Determine the Vite entry supplied by the evaluation harness. The fixture itself has neither `index.html` nor a Vite config; stop before using `npm run build` as an upgrade signal until that ownership is resolved.

**Proceed evidence:** supported Node/npm recorded, package-manager ownership known, current TypeScript result captured, and a concrete build/runtime entry identified. **Rollback:** none; this stage is read-only. **Stop:** any unexplained baseline failure, implicit dependency repair, or missing entry ownership.

### Stage 1 — make production source dual-compatible while still on 18.2

**Owner:** React API usage in `src/main.tsx` and `src/legacy-panel.tsx`.

- Replace legacy DOM mounting with guarded `createRoot(...).render(...)`.
- Replace `createFactory` with JSX and function `defaultProps` with a parameter default, resolving the `null` contract explicitly.
- Remove React namespace imports that become unused under `react-jsx`.
- Validate with `npx tsc --noEmit`, the identified Vite build path, and a browser probe that proves `ConsoleApp` renders. If a harness can reach `LegacyPanel`, prove its omitted-title fallback; otherwise record that runtime behavior as untested.

**Proceed evidence:** current React 18.2 still compiles, builds through the real entry, and renders the root without warnings introduced by the source migration. **Rollback:** restore only the two source files. **Stop:** changed visible fallback semantics, missing root ownership, or any baseline regression.

### Stage 2 — use the React 18.3 warning bridge

**Owner:** direct runtime dependency pair and package lock.

- In one package-manager transaction, pin `react` and `react-dom` to `18.3.1` and regenerate the entire lockfile; never hand-edit `scheduler`.
- Review the lock diff for only expected React runtime graph changes and restoration of complete direct dev-tool resolutions.
- Repeat Stage 1 static, build, and browser checks; capture every React deprecation warning and stop until each warning is either fixed or evidenced as irrelevant to the fixture.

**Proceed evidence:** exact 18.3 pair resolved, peer graph clean, lock complete, build/runtime evidence passes, and warning inventory is empty or dispositioned. **Rollback:** restore `package.json` and `package-lock.json` together to the reviewed Stage 1 checkpoint and reinstall from that lock. **Stop:** peer conflict, unrelated lock churn, warning without an owner, or clean-install failure.

### Stage 3 — move the runtime and type graph to React 19

**Owner:** four direct React packages and their resolved transitive graph.

- Re-query the four `latest` registry endpoints. If stable versions differ from this plan, pause for review rather than drifting the approved target.
- In one transaction, pin the matched runtime pair (`react` and `react-dom`) to the approved stable patch and the compatible React 19 type pair (`@types/react` and `@types/react-dom`) to reviewed exact versions.
- Regenerate the lock; require `react-dom`'s peer to accept the selected React version, `@types/react-dom`'s peer to accept the selected React types, and `scheduler` to resolve through `react-dom`'s declared range.
- Keep TypeScript 5.4.5 and Vite 5.4.21 unchanged unless their own diagnostics provide concrete evidence that a separate prerequisite is necessary.

**Proceed evidence:** clean install from the regenerated lock, no peer errors, `npx tsc --noEmit`, real Vite production build, browser root render, console free of React migration warnings/errors, and any owned error-reporting probe. **Rollback:** restore the complete Stage 2 manifest/lock checkpoint and reinstall; restore source files only if investigation proves the dual-compatible migration caused the regression. **Stop:** mismatched React pair, type-peer conflict, unexpected transitive package, new compile error, build failure, blank root, console error, or unresolved behavior change.

### Stage 4 — release evidence and cleanup

**Owner:** delivery pipeline, currently unknown.

- Compare the React 18.3 and React 19 browser output for the rendered root and any reachable `LegacyPanel` state.
- Confirm no emitted `.js` beside TypeScript sources or stale `dist/` output is included accidentally; retain only artifacts required by the real delivery process.
- Record the exact approved versions and validation environment. Release only through an identified project pipeline; the fixture provides none.

**Rollback:** redeploy the last validated React 18.3 artifact and restore its exact manifest/lock checkpoint. **Stop:** absent release owner, unrepeatable build, or runtime evidence that cannot be compared to the checkpoint.

## Unknowns requiring review

- Which Node/npm versions and package-manager policy own lock regeneration?
- Is the incomplete lock intentional synthetic evidence or a defect to repair during the first dependency transaction?
- What supplies the missing browser/Vite entry and root element?
- Is `LegacyPanel` intentionally dead production code, and should `null` mean “use Overview” or render empty content?
- Does an external host own error monitoring that depends on React 18's render-error behavior?

## Smallest safe first stage

Approve **Stage 0**, then **Stage 1** before changing any dependency. The first source edits use APIs supported by React 18.2 and React 19, isolate behavior-owner changes from resolution changes, and create a reversible checkpoint. No dependency mutation should begin until the current TypeScript result, actual Vite entry, supported Node/npm environment, root render, and `LegacyPanel` fallback decision are recorded.
