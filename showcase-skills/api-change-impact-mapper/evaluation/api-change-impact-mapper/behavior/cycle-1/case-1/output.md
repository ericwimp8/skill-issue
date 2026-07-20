# API Change Impact Report: HTTP Account Contract

## Decision Summary

The proposed `GET /accounts/{id}` v2 contract replaces the required top-level `display_name` field with required `profile.label`, removes `closed` from the allowed status values, and represents a missing account as `200` with a JSON `null` body instead of `404`. The fixture config selects v2, while the checked-in TypeScript SDK and mobile decoder still implement the old contract.

Overall transition risk is **high**. The highest-risk paths are the configured v2 response into the old mobile decoder, the public SDK surface generated from the old schema, and stored `closed` accounts flowing unchanged into a v2 response that forbids that value. Evidence does **not** support a general v2 rollout yet. Proceed only after the contract owner resolves `closed` handling, known consumers support the v2 body and missing-account semantics, the public SDK is regenerated, and routing can preserve v1 for consumers that have not migrated.

## Scope and Evidence

- **Scope and revision:** all seven files in `showcase-skills/api-change-impact-mapper/evaluation/api-change-impact-mapper/fixtures/http-contract/` as supplied for this evaluation. No repository revision or OpenAPI `info.version` is present, so the examined old/new files are referred to as v1 and v2 based on handler and route configuration evidence.
- **Contract authority:** `contracts/openapi-old.yaml` and `contracts/openapi-new.yaml` are the supplied contract authorities. Their organizational owner and publication process are unknown.
- **Inspected paths:** both OpenAPI contracts, route configuration, server handler, storage type/loader/cache key, generated TypeScript client, and Kotlin mobile decoder.
- **Methods:** semantic contract comparison; producer-to-consumer and storage tracing; exact identifier/category searches across the fixture; inspection of generator provenance and version configuration.
- **Concrete-runtime limitation:** `storage/account_store.ts:6`-`9` always returns `undefined`, so the only concretely reachable handler outcome in this isolated fixture is the missing-account branch. Present-account and historical-record findings are source-supported structural paths whose data source is absent.
- **Transport limitation:** `generated/account_client.ts:7`-`8` throws without performing HTTP I/O. Its public type is evidence of a stale generated surface, while its runtime parse behavior is unavailable.
- **Routing limitation:** `config/routes.yaml:1` selects v2 and `server/account_handler.ts:3` accepts a version, but the binding that turns configuration or a request into that argument is absent.
- **Checked absences:** no webhook, partner, plugin, queue, replay, backfill, migration, index, feature-flag, capability-negotiation, deprecation, owner/team, rollback, metric, telemetry, or logging definitions were found in the fixture. These categories are absent from the inspected fixture, not proven absent from the surrounding system.
- **Known external boundary:** `config/routes.yaml:2` names public package `@example/accounts-client`. Its published versions, dependants, and release cadence are unknown. `config/routes.yaml:3` explicitly leaves the minimum mobile version unknown. Other external consumers remain unknown; repository silence does not establish that none exist.

## Contract Delta

| Element | Old contract | New contract | Semantic effect / ambiguity |
|---|---|---|---|
| Success body | `200` requires an `Account` object (`contracts/openapi-old.yaml:6`, `contracts/openapi-old.yaml:12`) | `200` is `Account` or JSON `null` (`contracts/openapi-new.yaml:6`, `contracts/openapi-new.yaml:15`) | Consumers must accept and branch on a nullable success body. The null media representation and SDK mapping are not demonstrated by a generated v2 client. |
| Missing account | Explicit `404` (`contracts/openapi-old.yaml:15`) | No `404`; description says account or null and schema admits null (`contracts/openapi-new.yaml:8`, `contracts/openapi-new.yaml:18`) | Absence moves from HTTP status control flow to body-value control flow. Error-handling, caching, and observability semantics can change. |
| Display name | Required top-level `display_name: string` (`contracts/openapi-old.yaml:20`, `contracts/openapi-old.yaml:23`) | Required nested `profile.label: string` (`contracts/openapi-new.yaml:28`, `contracts/openapi-new.yaml:31`) | Wire shape and field path change; this is not a rename that old decoders can tolerate. |
| Status | Required enum `active`, `suspended`, `closed` (`contracts/openapi-old.yaml:24`) | Required enum `active`, `suspended` (`contracts/openapi-new.yaml:35`) | `closed` becomes invalid on v2. The contract does not define mapping, rejection, or replacement behavior. |
| Identifier | Required string `id` (`contracts/openapi-old.yaml:20`, `contracts/openapi-old.yaml:22`) | Required string `id` (`contracts/openapi-new.yaml:28`, `contracts/openapi-new.yaml:30`) | Semantically unchanged. |
| Contract version identity | No OpenAPI `info.version`; old behavior corresponds to handler v1 | No OpenAPI `info.version`; new behavior corresponds to handler v2 | Version names are inferred from repository wiring, and the public/version-negotiation mechanism is unknown. |

The new contract does not state defaults, deprecation windows, lifecycle guarantees, or how an existing `closed` account should be represented. Those meanings remain unknown rather than inferred.

## Impact Graph

### Server and Storage Paths

1. **Missing account, v1:** `loadAccount(id)` returns `undefined` in the supplied implementation (`storage/account_store.ts:6`-`9`) -> `getAccount` chooses the missing branch (`server/account_handler.ts:4`-`6`) -> v1 returns `404` -> matches the old contract's missing response (`contracts/openapi-old.yaml:15`).
2. **Missing account, v2:** the same storage result -> v2 returns `{ status: 200, body: null }` (`server/account_handler.ts:5`-`6`) -> matches the new nullable `200` response (`contracts/openapi-new.yaml:6`-`19`) -> reaches consumers that must tolerate null.
3. **Present account, v1:** `StoredAccount` uses `id`, `display_name`, and the three-value old status domain (`storage/account_store.ts:1`-`5`) -> the v1 handler returns it unchanged (`server/account_handler.ts:16`) -> structurally matches the old `Account`. This branch is unreachable with the supplied loader implementation.
4. **Present active/suspended account, v2:** stored `display_name` -> handler adapter -> `profile.label`; stored status passes through (`server/account_handler.ts:7`-`14`) -> structurally matches new `Account` for `active` or `suspended`. This branch is also unreachable with the supplied loader implementation.
5. **Present closed account, v2:** storage type permits `closed` (`storage/account_store.ts:4`) -> handler forwards status without validation or mapping (`server/account_handler.ts:13`) -> v2 emits a value excluded by the new enum (`contracts/openapi-new.yaml:35`). Any historical or future closed record would violate the advertised contract.
6. **Cache representation:** `cacheKey` includes old storage `status` and `display_name` (`storage/account_store.ts:11`). The response-shape adapter does not require a cache-key change. A later `closed`-record rewrite would change the status component and therefore requires cache invalidation or compatibility evidence.

### Consumer, Generation, Configuration, and External Paths

7. **Generated TypeScript surface:** the checked-in client declares the old fields and `closed` status and promises a non-null `Account` (`generated/account_client.ts:2`-`7`). Its provenance names `npm run generate:accounts` and `contracts/openapi-old.yaml` (`generated/account_client.ts:1`). Against v2, both present-account shape and missing-account nullability are stale. The generated artifact owner and a runnable generation definition are not present.
8. **Mobile decoder:** the Kotlin model expects `displayName`, and decoding force-unwraps top-level `id`, `display_name`, and `status` (`mobile/account_decoder.kt:1`-`2`). A v2 account object lacks `display_name`, and a v2 null body is not a `JsonObject`; both v2 outcomes are incompatible with this decoder.
9. **Route selection:** configuration sets `accounts_api_version: v2` (`config/routes.yaml:1`) -> an absent routing/binding layer presumably selects the handler's v2 branch -> known old consumers become exposed to v2. That connection is an evidence-supported inference, not a directly visible call path.
10. **Public SDK:** configuration names `@example/accounts-client` (`config/routes.yaml:2`) -> independently released consumers may depend on old types or missing-account behavior -> consumer inventory and deployed version mix are an external evidence boundary.
11. **Mobile releases:** minimum supported mobile version is explicitly unknown (`config/routes.yaml:3`) -> the repository cannot establish when the old decoder is no longer deployed -> safe v1 retirement remains externally bounded.

## Compatibility Matrix

| Material path | Classification | Versions / conditions | Rationale | Confidence | Unresolved evidence |
|---|---|---|---|---|---|
| Missing store result -> v1 handler -> old HTTP consumer | **Compatible** | v1 producer with old contract consumer | Concrete handler returns the old contract's `404`. | High | Actual HTTP routing and response serialization are absent. |
| Missing store result -> v2 handler -> v2-aware consumer | **Compatible** | v2 consumer must accept `200` plus JSON null | Concrete handler output matches the new nullable response. | High | Framework serialization of `body: null` is absent. |
| v2 missing response -> current Kotlin decoder | **Breaking** | Current decoder against v2 | Decoder requires a `JsonObject` and force-unwraps fields; v2 returns null for the concretely reachable missing path. | High | Mobile call-site error handling is absent. |
| v2 present response -> current Kotlin decoder | **Breaking** | Any v2 object against current decoder | Decoder requires top-level `display_name`; v2 supplies `profile.label`. | High on shape; medium on runtime reachability | Real storage implementation and mobile version distribution are absent. |
| v2 responses -> checked-in generated TypeScript API | **Breaking** | Old generated types against v2 | Type requires old shape, admits `closed`, and excludes null. | High on API contract; low on runtime | Client transport/parser is a throwing stub; published package contents are unknown. |
| Stored active/suspended -> v2 adapter -> new contract | **Conditionally compatible** | Record exists; status is `active` or `suspended`; serializer preserves shape | Adapter creates required `profile.label` and forwards an allowed status. | Medium | Supplied loader never returns a record; serialization is absent. |
| Stored closed -> v2 adapter -> new contract | **Breaking** | Any `closed` record exposed through v2 | Handler forwards `closed`, which v2 forbids. | High on code/schema; medium on occurrence | Historical record population and desired replacement semantics are unknown. |
| Stored record -> v1 adapter -> old consumer | **Conditionally compatible** | v1 remains routable; stored record follows declared type | Handler returns old storage shape matching old schema. | Medium | Present-record runtime and routing are absent. |
| Existing storage/cache -> additive response adapter | **Conditionally compatible** | Storage stays old-shaped and status is not rewritten | Nested v2 presentation is derived without changing stored `display_name`; current cache key remains stable. | Medium | Persistent store, cache engine, record age, and invalidation behavior are absent. |
| Route configuration -> known old consumers | **Breaking** | v2 selected while old mobile/client remains in use | Config selects v2 and both known consumer surfaces implement v1. | High on configured mismatch; medium on binding | Config-to-handler wiring and deployed consumer versions are absent. |
| v2 -> public SDK/external consumers | **Conditionally compatible** | Only if each consumer supports null, `profile.label`, and two-value status | The boundary is named but consumer implementations and version mix are unobservable. It cannot be treated as compatible without confirmation. | Low | Package releases, dependants, usage, and rollout control. |

## Migration Sequence

1. **Contract owner — resolve v2 semantics before generation.** Decide how `closed` accounts are represented, confirm that missing accounts intentionally become `200`/null, assign an explicit contract version, and define v1/v2 coexistence. Dependency: none. Completion evidence: approved OpenAPI plus decision record covering `closed`, absence, and version negotiation. Safe coexistence: keep v1 unchanged.
2. **Storage and server owners — make every v2 producer path conform.** Inventory actual `closed` records and implement the chosen mapping, rejection, or data transition before forwarding status; retain the existing v1 projection while old consumers remain. Dependency: step 1. Completion evidence: boundary tests showing v1 old shape/404 and v2 nested shape/null for missing, active, suspended, and the decided closed behavior. Safe coexistence: dual response projections over the current stored representation; any record rewrite must account for cache-key changes.
3. **SDK generator owner — regenerate and publish the public client.** Run the identified `npm run generate:accounts` pipeline from the finalized new schema, ensuring the generated return type represents `Account | null`, `profile.label`, and the final status domain. Dependency: steps 1-2. Completion evidence: reproducible generation command, clean generated-artifact drift check, published package version, and cross-version HTTP contract tests. Safe coexistence: preserve the old major/version for v1 consumers.
4. **Mobile owner — release a v2-capable decoder.** Add explicit null handling and decode `profile.label`; if one app build may reach both versions, use explicit version negotiation or a bounded dual-format adapter rather than guessing from missing fields. Dependency: step 1 and a stable v2 contract. Completion evidence: tests for null, active, suspended, the chosen closed behavior, malformed payloads, and confirmed production adoption above a defined minimum version. Safe coexistence: route older mobile versions to v1.
5. **External relationship/package owner — close the consumer evidence gap.** Inventory `@example/accounts-client` dependants and any unrepresented direct HTTP consumers; obtain readiness confirmation for body shape, status domain, and missing-account semantics. Dependency: published v2 client or equivalent migration guide. Completion evidence: consumer/version matrix with accountable confirmations. Safe coexistence: v1 remains available to unconfirmed consumers.
6. **Routing/deployment owner — stage v2 with observable selection.** Wire an explicit request/client capability to `apiVersion`, begin with allowlisted migrated consumers, and replace `minimum_mobile_version: unknown` with a verified gate. Dependency: steps 2-5. Completion evidence: configuration-to-handler integration tests, version-mix visibility, and a deployment manifest showing only ready consumers receive v2. Safe coexistence: additive v2 rollout with immediate routing fallback to v1.
7. **Contract and deployment owners — deprecate and clean up v1.** Retire v1, old generated surfaces, and any dual-read projection only after no supported old consumers remain and historical `closed` behavior is resolved. Dependency: sustained observation and external confirmation. Completion evidence: zero supported v1 traffic for the agreed window, consumer sign-off, no fallback use, and no incompatible stored records. Safe coexistence window ends only at this gate.

## Validation Plan

- **Contract:** validate both OpenAPI documents; add semantic compatibility checks for response status, nullability, required property paths, and enum narrowing; verify the finalized schema explicitly documents `closed` behavior.
- **Server boundary:** exercise `getAccount` through the real HTTP binding for missing, active, suspended, and closed records under both v1 and v2; verify status code, serialized JSON, and content type rather than only TypeScript shapes.
- **Generation:** run the owned generator from the finalized v2 schema; fail on generated drift; compile a consumer that branches on null and reads `profile.label`; preserve a v1 client test during coexistence.
- **Mobile boundary:** decode real serialized responses for both versions, including JSON null and malformed/unknown enum cases; verify business behavior for absence rather than only parse success.
- **Data and cache:** query actual status distribution and record age; replay representative historical records through both projections; if statuses are rewritten, verify cache-key migration/invalidation and stale-cache behavior.
- **Deployment:** test configuration-to-version binding, capability/version routing, v1 fallback, and a matrix of old/new producer and consumer versions. Confirm the minimum deployed mobile version from release telemetry or store/device evidence.
- **External:** obtain contract-test or release confirmation from public SDK consumers and any direct HTTP integrations; treat missing confirmation as a rollout blocker for that consumer path.

## Rollout and Observation

1. **Preflight:** keep v1 as the general route; finalize semantics, conform the server, regenerate clients, and establish consumer inventory. Gate: all boundary tests pass and every initial v2 consumer is identified.
2. **Allowlisted canary:** send only a small set of confirmed v2-capable SDK/mobile consumers to v2. Observe counts by selected API version, `404` versus `200`/null outcomes, response-schema violations, client parse failures, null-handling failures, unknown status values, and v1 fallback use. Gate: no contract violations or business correctness regression through an agreed traffic/time window.
3. **Progressive expansion:** increase v2 traffic by confirmed consumer version, retaining v1 routing. Also observe account-not-found business outcomes, support/error rates around account display, distribution of stored statuses, `closed` handling, and cache divergence after any data change. Gate: verified minimum mobile version and external consumer sign-off match the routed population.
4. **Default v2:** change the default only when unversioned or default-routed clients are proven capable. Stop and return affected traffic to v1 on schema reject/parse failures, nonzero forbidden `closed` responses, unexplained null/404 behavior changes, rising fallback, cache inconsistency, or account-display correctness regressions.
5. **v1 retirement:** begin the deprecation clock only after version-mix observation shows no supported v1 consumers. Continue monitoring late traffic and external confirmations through the retained compatibility window.

## Rollback Constraints

- **Reversible:** before v1 retirement, route consumers from v2 back to v1 and roll back the v2 handler projection, provided storage still retains the old `display_name` and old status semantics.
- **Retained compatibility requirement:** keep the v1 contract, handler path, and old SDK version available for the full mobile/external adoption window. A client already deployed expecting `200`/null and `profile.label` must remain routed to v2 or receive equivalent negotiated behavior.
- **Unsafe reversal:** rewriting or deleting `closed` records, changing cache keys without retained mappings, removing `display_name`, unpublishing/reusing a public SDK version, or removing v1 after independently deployed clients still depend on it cannot be safely reversed by a server toggle alone.
- **Forward-repair triggers:** externally deployed v2-only clients after v1 restoration, irreversible data normalization, expired v1 infrastructure, or public observation of reused schema/package versions require a forward-compatible server/client repair rather than simple rollback.

## Open Decisions and Unknowns

| Decision or unknown | Impact | Smallest resolving action |
|---|---|---|
| What should v2 return for stored `closed` accounts? | Blocks schema-conforming production and data migration design. | Contract owner selects and documents mapping, rejection, replacement status, or exclusion behavior. |
| Is `200`/null the intentional public absence semantic, including cache and business behavior? | Determines consumer control flow, telemetry, and compatibility guarantees. | Approve the semantic change and add one HTTP boundary example/test. |
| How is `apiVersion` selected and can versions coexist per consumer? | Determines whether old consumers can be protected during rollout. | Trace or add the routing binding and prove a v1/v2 selection integration test. |
| What real store/cache supplies `StoredAccount`, and do `closed` records exist? | Determines reachability, backfill need, and cache risk. | Query production-like status counts and identify the concrete persistence/cache owner. |
| Which mobile versions contain the current decoder? | Unknown adoption makes default v2 unsafe. | Produce the deployed-version distribution and set a verified minimum version gate. |
| Who consumes `@example/accounts-client` or calls the HTTP API directly? | Unknown external paths cannot be classified as ready. | Build a consumer/version inventory and request explicit migration confirmation. |
| Who owns and runs `npm run generate:accounts`? | Regeneration and drift control lack an accountable completion path. | Locate the script/CI owner and reproduce generation from the finalized schema. |
| How are contract violations and version mix observed? | Canary gates and rollback triggers cannot currently be enforced. | Define metrics/logs for route version, response validation, parse failures, null outcomes, forbidden statuses, and fallback. |

Every material delta reaches either a concrete handler/consumer/storage effect or a named repository/external evidence boundary. No implementation or fixture file was changed as part of this report.
