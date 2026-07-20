# API Change Migration Impact Report

## Decision Summary

The proposal changes the public `@example/router` function `resolveRoute` from
`(path: string, strict: boolean) => Route | undefined` to
`(strict: boolean, path: string) => Promise<Route>`. This simultaneously swaps
two positional parameters, changes synchronous delivery to asynchronous
delivery, and removes `undefined` from the declared success result. The
fixture's concrete implementation and every observed caller still implement
the old contract.

Overall risk is **high and breaking**. The highest-risk paths are the three
positional internal calls, the synchronous `legacyResolve` adapter, unknown
external plugin/package consumers, and the unobserved `lookup` implementation
whose backend and missing-route/error semantics determine whether the new
promise contract is truthful. A declaration-only publication would also leave
generated documentation tied to old source behavior.

Evidence does **not** support publishing the replacement in place. Proceed only
after the missing concrete backend behavior and external consumer set are
resolved, an additive/versioned transition is chosen, callers are migrated,
and a prerelease passes cross-version and backend-specific validation.

## Scope and Evidence

- **Change authority examined:** the supplied old and new declaration files at
  `contracts/old.d.ts` and `contracts/new.d.ts`. They describe the proposed
  public surface, but the fixture does not identify who owns them or how they
  are generated.
- **Current package evidence:** `package.json:2-5` names `@example/router`,
  reports version `3.4.0`, and generates TypeDoc from `src/index.ts`.
- **Concrete source traced:** `src/index.ts:1` re-exports `resolveRoute` from
  `src/router.ts`; `src/router.ts:1-3` owns the observed implementation and
  delegates to `lookup(backend, path, strict)`.
- **Consumers traced:** `cli/run.ts`, `service/route_service.ts`, and
  `plugins/legacy_adapter.ts` all import through the `src` re-export and call
  the old positional signature.
- **Configuration traced:** `config/backends.env:1-2` selects `redis` and
  explicitly records external plugin consumers as unknown. The implementation
  also has a `memory` fallback at `src/router.ts:2`.
- **Revision:** isolated fixture snapshot supplied by the case. No VCS revision,
  target package version, deployment topology, or release channel is present.
- **Methods:** complete file inventory and read; semantic old/new contract
  comparison; symbol, caller, adapter, documentation, package, and
  configuration searches; trace through the re-export to concrete source.
- **Checked absences:** the fixture contains no tests, `await`, exception
  handling, feature/capability negotiation, package export/type entry metadata,
  lockfile/workspace metadata, migration/backfill code, cache/queue/replay
  paths, generated documentation artifact, or publication configuration.
- **Evidence boundaries:** `lookup`, `Route`, backend implementations, package
  registry consumers, plugin consumers, generator configuration beyond the
  script, release automation, and runtime deployment state are absent.
  Repository silence cannot establish that these external consumers or data
  paths do not exist.

## Contract Delta

| Element | Old contract | New contract | Semantic effect and ambiguity |
|---|---|---|---|
| Parameter 1 | `path: string` | `strict: boolean` | Existing positional calls send a path where the new implementation expects strictness. |
| Parameter 2 | `strict: boolean` | `path: string` | Existing positional calls send a boolean where the new implementation expects a path. |
| Delivery | Direct return | `Promise` | Every consumer must adopt asynchronous control flow; truthiness and property access change before resolution. |
| Missing route | Declared `undefined` | No `undefined` in fulfillment type | The new outcome could reject, fabricate a route, or guarantee lookup success; the fixture does not define which. |
| Errors | Unspecified synchronous behavior | Unspecified promise behavior | Throw-versus-rejection semantics, error types, and retry guarantees are unknown. |
| Name | `resolveRoute` | `resolveRoute` | Same-name replacement reaches all public consumers without an additive opt-in boundary. |
| Ordering/lifecycle | Result available in current call stack | Result available after promise settlement | CLI exit timing, service return type, adapter compatibility, and caller error handling all change. |

The declaration files are the supplied contract authority for comparison.
However, `src/router.ts:1-3` remains the concrete behavior owner in this
snapshot, and it still has the old parameter order and direct return. The
relationship between hand-authored declarations and source generation is
unknown. The target version is also unknown; because `3.4.0` is current and the
change breaks a public function, an in-place minor or patch publication would
be incompatible with conventional package-version expectations.

## Impact Graph

1. **Public declaration to concrete behavior:** `contracts/new.d.ts:1-4` declares
   boolean-first asynchronous behavior -> public import reaches
   `src/index.ts:1` -> re-export reaches `src/router.ts:1-3` -> current
   implementation remains path-first and returns `lookup` directly. The
   declaration and implementation are presently divergent.
2. **CLI path:** `cli/run.ts:1` -> `src/index.ts:1` ->
   `src/router.ts:1-3` -> unobserved `lookup`. At `cli/run.ts:2`, the old call
   passes `(string | undefined, true)`; under the new signature those positions
   are reversed. If it returned a promise, `cli/run.ts:3` would test the
   promise object's truthiness rather than the resolved route, so the old
   missing-route exit path would not run. No rejection handling exists.
3. **Service path:** `service/route_service.ts:1-3` -> re-export -> concrete
   router -> `lookup`. The old call uses `(path, false)`. Under the new contract
   it has reversed arguments, and `?.target` is applied to a promise rather
   than a resolved `Route`. In unchecked JavaScript this falls through to
   `'/404'`; under appropriate TypeScript declarations it should be rejected by
   type checking. Migrating it to `await` also changes `routeRequest` into an
   asynchronous API, extending impact to its unobserved callers.
4. **Legacy adapter path:** `plugins/legacy_adapter.ts:1-2` -> re-export ->
   concrete router -> `lookup`. `legacyResolve(path)` promises an old-order,
   apparently synchronous adapter surface. The new call order breaks its
   forwarded values, while the promise return breaks synchronous adapter
   consumers. An adapter cannot turn a genuinely asynchronous operation back
   into the same synchronous contract without a different implementation or
   behavior.
5. **Backend path:** `src/router.ts:2-3` reads `ROUTER_BACKEND`, defaults to
   `memory`, and passes the backend plus route inputs to `lookup`.
   `config/backends.env:1` selects `redis` for the supplied configuration.
   Neither `lookup` nor either backend implementation is present, so latency,
   missing-route, error, persistence, and rollback effects terminate at this
   explicit evidence boundary.
6. **Documentation path:** `package.json:5` runs `typedoc src/index.ts` ->
   `src/index.ts:1` -> `src/router.ts:1-3`. Therefore documentation generation
   is sourced from the current implementation/re-export, not from either
   supplied declaration file. Changing only `contracts/new.d.ts` can publish
   new types while generated docs continue to describe old source behavior.
   No generated artifact is present to inspect.
7. **Package and external path:** `package.json:2-3` identifies a public package
   at `3.4.0`. Package exports/type entry points are absent, so the exact packed
   artifact cannot be confirmed. `config/backends.env:2` explicitly says
   external plugin consumers are unknown. Registry consumers and independently
   released clients remain an unknown external boundary and must be assumed
   capable of using the old positional/synchronous contract until inventoried.
8. **Persistence and replay path:** no storage, cache, queue, replay, migration,
   or backfill surface appears in the fixture. Backend persistence cannot be
   excluded because backend implementations are outside the evidence boundary.

## Compatibility Matrix

| Material path | Classification | Versions/conditions | Rationale | Confidence | Unresolved evidence |
|---|---|---|---|---|---|
| New declaration -> current `src/router.ts` | **Breaking** | Proposed contract against fixture snapshot | Parameter order and return delivery disagree at the concrete implementation. | High | Declaration generation/packaging owner |
| CLI -> router -> `lookup` | **Breaking** | Old caller with proposed same-name API | Arguments reverse; a promise is always truthy; missing-route exit and rejection handling fail. | High through router | `lookup`, CLI contract, desired rejection exit code |
| Route service -> router -> `lookup` | **Breaking** | Old caller with proposed same-name API | Arguments reverse; property access occurs before promise fulfillment; migration propagates async to callers. | High through router | Callers of `routeRequest`, missing-route policy |
| `legacyResolve` -> router -> external plugin caller | **Breaking** | Old adapter consumers with proposed API | Arguments reverse and synchronous compatibility cannot be preserved by simple forwarding. | High locally; low externally | External consumer inventory and supported adapter promise |
| Redis-configured backend | **Conditionally compatible** | Only if Redis lookup accepts corrected inputs and implements the new fulfillment/rejection guarantee | Backend is selected, but its implementation and semantics are absent. | Low | Redis lookup behavior, failures, data effects |
| Memory fallback backend | **Conditionally compatible** | Only if fallback lookup implements the same new guarantee | Fallback is concrete configuration behavior, but implementation is absent. | Low | Memory lookup behavior and parity with Redis |
| TypeDoc generation -> published docs | **Conditionally compatible** | Source must be updated first; docs must be regenerated and reviewed from the packed release | Current command follows old concrete source, so declaration-only publication creates drift. | High for source path | TypeDoc config/output and release inclusion |
| Existing `@example/router` consumers | **Breaking** | Any consumer compiled or coded to the old signature | Same-name positional and async changes require source/control-flow migration. | High for contract; unknown population | Registry/package consumer inventory |
| New-only consumer against old runtime | **Breaking** | New declarations paired with old implementation or rollback to v3 | It supplies boolean-first values to a path-first runtime and expects a promise. | High | Actual packaging resolution |
| Storage/history/replay | **Conditionally compatible** | Only if backend lookup is read-only and route representation is unchanged | No such path is visible, but backend internals are absent. | Low | Backend persistence and historical-data behavior |

## Migration Sequence

1. **Public API owner — settle contract semantics and version strategy.** Define
   missing-route fulfillment versus rejection, error types, cancellation and
   ordering guarantees, and the target package version. Prefer an additive
   async API or versioned entry point in the `3.x` line while retaining the old
   `resolveRoute`; reserve removal or same-name replacement for a major
   version. Dependency: none. Completion evidence: approved contract examples
   for success, miss, and backend failure plus an explicit semver decision.
   Safe coexistence: entire deprecation window.
2. **Backend owner — resolve the `lookup` evidence boundary.** Inspect Redis and
   memory implementations, confirm whether each is synchronous or asynchronous,
   define miss/error mapping, and identify reads/writes or externally visible
   effects. Dependency: settled semantics. Completion evidence: concrete path
   trace and backend contract tests for both configured values. Safe
   coexistence: both old and additive APIs must map outcomes consistently.
3. **Library implementation owner — implement additive/versioned behavior.**
   Keep the old path-first synchronous export stable if its backend permits;
   add the new boolean-first async surface at a distinct opt-in boundary. If a
   same-name replacement is mandatory, implement it only on the new major and
   retain the prior major for old consumers. Dependency: backend findings.
   Completion evidence: source and emitted declarations agree, and packed API
   inspection shows the intended exports. Safe coexistence: until supported old
   consumers have migrated.
4. **Internal caller owners — migrate explicit control flow.** Update the CLI to
   await settlement and map miss/rejection to defined exit codes; update the
   route service to await before reading `target` and trace the resulting
   promise through all service callers. Dependency: additive API available.
   Completion evidence: type checking plus success, miss, and failure
   integration checks against both backends. Safe coexistence: callers can move
   independently while the old API remains published.
5. **Plugin/adapter owner — choose a truthful adapter contract.** Keep
   `legacyResolve` bound to the old implementation during coexistence. Offer a
   separately named async adapter for migrated plugins; do not silently change
   the legacy adapter's return lifecycle. Dependency: external contract and
   consumer inventory. Completion evidence: adapter contract tests and
   confirmation from every supported plugin owner. Safe coexistence: through
   the announced plugin support/deprecation window.
6. **Documentation/package owner — align artifacts.** Make source,
   declarations, TypeDoc output, examples, package export/type metadata, and
   migration guide describe the same versioned surfaces. Dependency:
   implementation and adapter decisions. Completion evidence: generated-doc
   diff and inspection of the exact packed tarball. Safe coexistence: docs must
   show both APIs throughout deprecation.
7. **Release owner — stage publication.** Publish a prerelease; exercise the
   packed package with old and new consumer fixtures; canary against memory and
   Redis; expand only after error, miss, fallback, and version-mix gates pass.
   Dependency: steps 1-6. Completion evidence: prerelease install results,
   canary telemetry, and external plugin confirmations. Safe coexistence: keep
   the last compatible `3.x` available and supported.
8. **Public API owner — remove legacy surface only in the planned major.** Wait
   until internal migration is complete, supported external consumers confirm
   migration, and the deprecation window expires. Completion evidence: zero
   observed legacy use for the agreed period, consumer sign-off, final API/doc
   review, and major-version publication approval.

## Validation Plan

- **Contract:** mechanically compare old/new emitted API shape; compile examples
  for both argument orders; assert exact success, miss, and failure types; check
  source declarations against the packed package rather than the fixture file
  alone.
- **Concrete implementation:** test argument propagation from each public entry
  through `src/index.ts`, `src/router.ts`, and the actual `lookup`; verify that
  strictness and path cannot be silently swapped.
- **Internal callers:** run CLI success/miss/backend-failure cases and assert
  exit codes; run route-service success/miss/rejection cases and assert its
  resulting sync/async contract; validate both legacy and new adapters.
- **Backend/configuration:** execute the same matrix with
  `ROUTER_BACKEND=redis`, `ROUTER_BACKEND=memory`, and the unset fallback.
  Compare route, miss, latency, rejection, retry, and any data-side effects.
- **Cross-version:** test old consumer + old runtime, old consumer + transition
  runtime, new consumer + transition runtime, and new consumer + new-major
  runtime. Explicitly prove that unsupported new-consumer + old-runtime and
  old-consumer + new-only-runtime combinations fail before release selection.
- **Generated artifacts:** run `npm run docs`, review the TypeDoc signature and
  examples, detect generated drift in CI, and inspect the release tarball's
  JavaScript, declarations, exports, and documentation.
- **External boundary:** inventory registry dependants and plugin owners; obtain
  explicit compatibility confirmation or bound the supported consumer set.
  Treat silence as unresolved rather than success.
- **Data/history:** after backend inspection, add replay/historical-data checks
  only if lookup reads durable route records or writes observations/caches.

## Rollout and Observation

1. **Prerelease gate:** publish the additive/versioned API as a prerelease.
   Stop on source/declaration/doc drift, any swapped-argument observation, or a
   mismatch between memory and Redis outcomes.
2. **Internal canary:** migrate the CLI and route service separately. Observe
   resolved-route correctness, miss rate, promise rejection count and class,
   CLI exit-code distribution, `/404` fallback rate, backend latency, and
   Redis-versus-memory divergence. Stop on unexplained fallback increases,
   unhandled rejections, or route correctness regressions.
3. **Plugin canary:** release the async adapter to confirmed plugin consumers.
   Observe legacy/new adapter use, consumer version mix, parse/type failures,
   and plugin-reported route correctness. Stop if any supported consumer cannot
   migrate within the coexistence window.
4. **Stable additive release:** publish only after canary gates pass and packed
   documentation matches behavior. Continue monitoring old/new API usage and
   backend-specific failures throughout deprecation.
5. **Major cutover:** remove the old surface only after the consumer gate and
   deprecation window complete. Stop or delay if unknown plugin consumers remain
   material, legacy use persists, or rollback compatibility is no longer
   available.

## Rollback Constraints

- Before consumers adopt the new API, rollback is limited to withdrawing the
  prerelease/additive export and returning callers to the retained old surface.
- After any caller adopts boolean-first promise behavior, rolling its runtime
  back to the old path-first synchronous implementation is unsafe; that caller
  must be rolled back in coordination or kept on a compatible transition build.
- After `routeRequest` or adapter callers propagate asynchronous return types,
  reverting only the library leaves their control flow incompatible. Retain the
  additive API until downstream rollback windows close.
- Published clients and external plugins may remain on old versions
  indefinitely. Preserve the previous major and its documentation for the
  declared support period; package unpublication is not a reliable rollback.
- No irreversible writes or destructive migrations are evidenced in the
  fixture. If Redis lookup writes cache or route state, rollback safety must be
  reassessed from its concrete schema and write behavior before rollout.
- A same-name major cutover observed by external consumers, expired
  coexistence window, or incompatible backend data effect requires forward
  repair (compatibility release or corrected adapter), not a simple binary
  rollback.

## Open Decisions and Unknowns

| Decision or unknown | Impact | Smallest resolving action |
|---|---|---|
| What guarantees replaces `undefined`? | Determines API truthfulness, CLI exit behavior, service fallback, and error telemetry. | Specify success, miss, and backend-failure examples in the new contract. |
| Where are `lookup`, `Route`, Redis, and memory implementations? | Prevents confirmation of async behavior, backend parity, persistence, and rollback safety. | Inspect and trace those concrete implementations. |
| Are declaration files generated or hand-authored? | Determines the semantic owner and drift-prevention step. | Identify their build command and compare emitted declarations with source. |
| Which package version will carry the change? | Determines semver and coexistence obligations. | Record the target version and whether the same-name change is deferred to a major. |
| Which consumers call `routeRequest`? | Its async migration may widen the internal impact graph. | Search the owning repository/package for all imports and calls. |
| What contract does `legacyResolve` promise? | Determines whether it can remain old or must version its adapter surface. | Confirm adapter support policy and add explicit sync/async contract tests. |
| Which external package/plugin consumers are supported? | Unknown old callers block a safe removal date. | Inventory registry dependants and obtain owner/version confirmations. |
| What exactly is published? | Missing `exports`, `types`, and release metadata prevents verification of runtime/type pairing. | Inspect the packed tarball and release pipeline. |
| How are docs published and retained per version? | A declaration-only change can conflict with TypeDoc generated from old source. | Run the docs command on the release candidate and verify versioned output. |
| Does Redis lookup read or write durable state? | Could add historical-data, cache, and forward-repair requirements. | Trace Redis commands and schemas, then add data compatibility checks if needed. |

Every material delta has been traced to a concrete caller/effect or to the
named `lookup`, packaging, documentation, or external-consumer evidence
boundary. This report changes no fixture implementation.
