# Order Event Schema Change Impact Report

## Decision Summary

The proposal replaces `amount_cents` with `amount_micros` while reusing protobuf field number 2, adds `warehouse_id` at field 4, and adds `BACKORDERED = 2`. The checkout producer already emits the proposed V2 object to the existing `orders.placed` topic for a configured 5% cohort, alongside V1 events.

Overall risk is **critical**. Evidence does not support expanding the rollout. Reusing field 2 for a value whose unit is 10,000 times smaller is wire-readable but semantically breaking: an old field-2 consumer can interpret micros as cents, while object-shape consumers can receive no `amount_cents` at all. The highest-risk confirmed paths are checkout to billing, mixed archived events through replay, and archived analytics. The partner webhook is also release-blocking because it forwards the event unchanged and the partner versions are explicitly unknown.

Proceed only with containment and contract redesign. Pause V2 emission, preserve the cents meaning of field 2, introduce the new unit additively on an unused field or a separately versioned topic/message, and establish version-aware coexistence before resuming a canary.

## Scope and Evidence

- **Revision examined:** `bb77a9285996d528b1bfd9c8bb376be0f6751f20`.
- **Contract authority examined:** `contracts/order-old.proto` and `contracts/order-new.proto`. They define the old and proposed protobuf shapes, but the fixture does not show registration, generation, serialization, or runtime binding, so their relationship to the plain-object bus publisher is an evidence gap.
- **Production and configuration:** `checkout/publish.ts:1-17` and `config/features.yaml:1-3`.
- **Internal consumption and persistence-facing use:** `billing/consume.ts:1-3` and `analytics/consume.sql:1`.
- **Replay and external forwarding:** `replay/replay.ts:1-4` and `webhook/forward.ts:1-3`.
- **Methods:** complete fixture inventory and read; semantic contract comparison; producer-to-consumer, archive, replay, configuration, and external-boundary tracing; targeted negative searches.
- **Negative searches:** no generator or `protoc` owner, migration/backfill or dual-read/write implementation, schema-version/capability negotiation, reject/DLQ/parse-error handling, cache/index/queue implementation, or named team owner appears in the fixture.
- **Limitations:** bus serialization, subscriber routing, archive writer/schema, database constraints, gateway behavior, partner contract, deployment units, enum runtime, observability, and external consumers are unavailable. Repository silence does not establish that other consumers are absent.

## Contract Delta

| Element | Old contract | Proposed contract | Semantic effect and authority |
|---|---|---|---|
| Monetary field 2 | `int64 amount_cents = 2` | `int64 amount_micros = 2` | Name and unit change while preserving wire number/type. Protobuf bytes remain parseable, but the numeric meaning changes by a factor of 10,000. This is a semantic field-number reuse and is breaking for old consumers. Authority: the two contract files; runtime binding is unknown. |
| Warehouse | Absent | `string warehouse_id = 4` | Additive field. Older protobuf readers normally ignore an unknown field, but plain-object consumers and partner validation behavior are not shown. The checkout V2 producer supplies `order.warehouse`. |
| Status enum | `PENDING = 0`, `CONFIRMED = 1` | Adds `BACKORDERED = 2` | Additive on the wire, but compatibility depends on each consumer/runtime accepting and preserving unknown enum values. No handling or fallback is shown. |
| Event identity | `OrderPlaced` | `OrderPlaced` | Same message name; checkout also uses the same `orders.placed` topic for both shapes. There is no schema/version discriminator. |
| Optionality/defaults | Proto3-style scalar declarations shown without presence modifiers | Same style for retained fields | The fixture gives no presence policy. Missing object properties and protobuf scalar defaults can behave differently; no validation contract resolves this ambiguity. |

The checkout arithmetic confirms the intended unit conversion: V1 emits `order.total * 100` and V2 emits `order.total * 1000000` (`checkout/publish.ts:4-15`). Therefore the rename is not cosmetic.

## Impact Graph

### Live production and billing

`config/features.yaml:1` selects a 5% V2 cohort -> `checkout/publish.ts:1-16` branches on `useV2` -> both branches publish to `orders.placed` -> V1 includes `amount_cents`, while V2 includes `amount_micros` and `warehouse_id` -> `billing/consume.ts:1-2` accepts only `amount_cents` and sends it directly to `gateway.charge`.

Concrete effects depend on the unshown bus binding, and both evidenced interpretations are unsafe:

1. If the plain V2 object reaches billing, `event.amount_cents` is absent and the gateway receives `undefined` or the transport rejects before invocation.
2. If an old generated protobuf consumer decodes V2 bytes, tag 2 can populate its `amount_cents` field with the micros integer, causing a potential 10,000-times overcharge.

### Archive, analytics, and replay

The configured archive retention is 365 days (`config/features.yaml:2`). `analytics/consume.sql:1` reads `amount_cents` from `archived_order_events` and divides by 100. The archive writer and V2 storage mapping are absent, so V2 rows may be rejected, stored with a missing/null cents column, or store the micros value under the old field-2/cents interpretation. Those outcomes respectively lose events, produce missing revenue, or overstate revenue by 10,000.

`replay/replay.ts:1-3` reads every archived `orders.placed` event and republishes it unchanged to the same live topic. Because retention spans 365 days and neither event shape carries a version discriminator, replay can mix historical cents events with new micros events. A V2 consumer that interprets tag 2 as micros can understate replayed V1 amounts by 10,000; a V1 consumer can reject or misinterpret replayed V2 events. Replay therefore extends the coexistence obligation to the full retained history unless records are versioned or normalized before republishing.

### Partner forwarding and unknown consumers

`webhook/forward.ts:1-2` accepts an unknown event and posts it unchanged to `/order-events`. No adapter, filtering, unit conversion, version header, or capability check is present. `config/features.yaml:3` explicitly records `partner_consumer_versions: unknown`. Any partner expecting `amount_cents`, a closed field set, or only the two old enum values can reject or misinterpret V2. The repository cannot establish compatibility for this independently deployed boundary.

Other subscribers to the shared topic are an explicit unknown boundary. The fixture contains no registry or exhaustive consumer inventory.

### Generation, configuration, and deployment

The `.proto` files are schema inputs, but no generated artifacts or generation command are present. Generated clients remain a likely dependent surface whose languages and unknown-enum behavior require inventory. The 5% feature value proves mixed-version production, but the fixture does not show how it is evaluated, how quickly it can be disabled, whether selection is stable, or whether billing/webhook/replay deployments can be independently coordinated.

## Compatibility Matrix

| Material path | Classification | Versions or required conditions | Rationale | Confidence | Unresolved evidence |
|---|---|---|---|---|---|
| Checkout V2 -> billing object consumer -> gateway | **Breaking** | V2 producer with current billing | V2 omits `amount_cents`; billing reads it directly and passes it to charging. | High | Bus validation and gateway response to missing value. |
| Checkout V2 protobuf bytes -> old field-2 billing client | **Breaking** | New producer, old generated schema | Same tag/type parses, but micros are interpreted as cents, potentially charging 10,000x. | Medium | Actual serialization and generated binding are absent. |
| Checkout V1 -> current billing | **Compatible** | V1 producer and shown billing shape | Both use `amount_cents`; no changed field is required on this path. | High | End-to-end transport remains unshown. |
| `warehouse_id` -> old protobuf consumers | **Conditionally compatible** | Consumers must ignore unknown tag 4 | Additive unknown protobuf fields are normally tolerated; strict object/schema validators may reject. | Medium | Consumer runtimes and validation settings. |
| `BACKORDERED` -> old consumers | **Conditionally compatible** | Runtime must accept/preserve unknown enum values and business logic must define fallback | Field is additive, but old closed-enum logic can reject or mishandle value 2. | Medium-low | Generated languages, status branches, partner rules. |
| Mixed producer -> archive storage -> analytics | **Breaking** | Current query with any V2 storage lacking normalized cents | Query assumes an `amount_cents` column and cents semantics; V2 has neither by contract. | High for query mismatch; medium for exact failure | Archive writer/schema and null/reject behavior. |
| 365-day archive -> replay -> V1/V2 consumers | **Breaking** | Mixed retained history on the unversioned topic | Replay republishes unchanged; tag 2 has two meanings and no discriminator. | High | Exact archive encoding and consumer population. |
| Bus event -> unchanged partner webhook | **Conditionally compatible, unresolved** | Every partner must explicitly accept both shapes/units/statuses, or forwarding must negotiate/adapt | Forwarding is transparent and configured partner versions are unknown. Expansion is unsafe until confirmed. | High on evidence gap; low on partner behavior | Partner contracts, release state, validation, and rollback expectations. |
| V1/V2 on same topic under 5% flag | **Breaking** | Any consumer that is not dual-shape and unit-aware | Cohort produces simultaneous incompatible meanings without event-level versioning. | High | Flag evaluator and routing topology. |
| Other external/topic consumers | **Unknown; treat as release-blocking** | Exhaustive inventory and compatibility confirmation required | No consumer registry is available; absence from fixture is not evidence of absence. | High on unknown boundary | Subscriber inventory and owners. |

## Migration Sequence

1. **Feature-configuration owner — contain the rollout.** Set V2 emission to zero or otherwise stop new V2 events. Dependency: confirm the control path and propagation time. Risk addressed: live billing, analytics, replay, and partner exposure. Completion evidence: observed V2 count reaches zero after the maximum propagation window and no `amount_micros` events continue. Safe coexistence window: none for the current ambiguous field-2 design.
2. **Event-contract owner — redesign additively.** Preserve `amount_cents = 2` and never reuse tag 2 for micros semantics. Add `amount_micros` on an unused tag (for example tag 5, since tag 4 is proposed for `warehouse_id`) or publish a distinctly versioned message/topic. Define canonical amount precedence, rounding, status evolution, and event-version semantics. Dependency: consumer/partner inventory. Completion evidence: reviewed contract and compatibility checks proving old readers retain cents semantics.
3. **Checkout owner — implement coexistence production.** Emit a contract that old consumers can read correctly, such as retaining cents while adding micros, or route V2 only to a new versioned topic. Dependency: revised contract and serializers. Risk addressed: mixed-shape production. Completion evidence: producer integration evidence for exact values, tags, topic, and event version across feature states. Coexistence window: through all consumer upgrades plus retained/replayable history.
4. **Billing owner — add explicit version/unit handling.** Validate presence, unit, range, and supported schema before charging; support both formats only under the contract's unambiguous precedence rule. Dependency: revised contract. Completion evidence: cross-version boundary checks show identical charge amounts for V1 and V2 representations and rejection of ambiguous events.
5. **Archive/data owner — make stored representation explicit.** Inventory existing V2 canary records, record original schema/version, retain an unambiguous canonical unit, and backfill or quarantine ambiguous rows. Dependency: identify archive encoding and current V2 exposure. Completion evidence: counts reconcile by version, monetary totals reconcile, and no row's tag-2 value is interpreted without provenance. Coexistence window: at least 365 days, or until all ambiguous history is migrated/expired.
6. **Replay owner — version or normalize replay.** Prevent unchanged mixed-history republishing. Decode by recorded event version and publish a normalized compatible event or version-specific topic. Dependency: archive provenance/backfill. Completion evidence: replay tests over V1, proposed-canary, and revised V2 records yield identical business amounts and supported statuses.
7. **Analytics owner — migrate queries after storage is dual-readable.** Read a canonical amount or convert by explicit version; reconcile against known samples before switching. Dependency: archive migration. Completion evidence: V1/V2 revenue parity and historical aggregate reconciliation.
8. **Partner-relationship/webhook owner — establish external compatibility.** Obtain each partner's accepted fields, units, enum behavior, versioning mechanism, deployment date, and rollback needs; adapt or negotiate at the webhook boundary. Dependency: partner inventory and revised contract. Completion evidence: partner confirmation plus boundary test for both coexistence versions.
9. **Generated-client/deployment owners — regenerate and coordinate.** Locate generation commands and all independently deployed clients, regenerate from the revised schema, and verify unknown-field/enum behavior by language. Dependency: final schema. Completion evidence: no artifact drift and a deployment matrix showing every consumer safe before canary expansion.
10. **Feature-configuration owner — resume staged rollout and retire compatibility deliberately.** Start with internal version-aware consumers, then replay/analytics, then confirmed partners. Dependency: all preceding gates. Retain old-field support until the 365-day archive window is expired or migrated and all deployed/external consumers confirm retirement. Completion evidence: stage gates and observation signals remain healthy for an agreed window.

## Validation Plan

- **Contract:** assert field 2 retains cents semantics; prevent field-number semantic reuse; verify new amount and warehouse tags; verify enum evolution in every generated runtime; detect generated-artifact drift.
- **Producer:** for representative totals, assert V1 and revised V2 encode equivalent monetary value, correct rounding/overflow behavior, stable topic/version identity, and feature-state selection.
- **Billing:** exercise old/new producer against old/new consumer matrices at the actual bus boundary. Verify the gateway receives the same cents charge, and missing, ambiguous, unsupported-version, overflow, and `BACKORDERED` cases fail safely before charging.
- **Data and analytics:** inspect canary rows; validate archive version provenance; compare counts and sums before/after conversion; test null/reject paths; reconcile revenue for known V1/V2 samples and the complete backfill.
- **Replay:** replay mixed events spanning the retention window into isolated consumers; confirm no 10,000x over/understatement, duplicates, loss, unsupported enums, or forwarding to production partners.
- **Configuration and deployment:** test 0%, canary, and 100% matrices; prove flag propagation and emergency shutoff; verify deployment ordering across checkout, billing, archive, replay, analytics, and webhook.
- **External:** obtain partner-owned contract confirmation and test payload acceptance, amount meaning, enum handling, version negotiation, duplicate behavior, and partner-side rollback.

## Rollout and Observation

1. **Containment gate:** stop the current V2 cohort. Observe counts by detected shape, billing validation failures, gateway charge anomalies, archive rejects/nulls, webhook failures, and partner acknowledgements. Stop if V2 continues after the control propagation window.
2. **Dark compatibility gate:** deploy revised consumers, storage, replay normalization, and webhook adaptation while continuing V1 production. Require zero amount divergence and successful mixed-history replay.
3. **Internal canary gate:** emit revised V2 only to version-aware internal paths. Track event counts by version, parser/validation rejects, fallback use, charge amount reconciliation, archive row/version counts, analytics variance, replay failures, and `BACKORDERED` handling. Stop on any unclassified event, amount mismatch, charge anomaly, data loss, or replay divergence.
4. **Partner canary gate:** enable only confirmed partners with an observable version route. Track delivery status, partner rejects, partner-reported amount/status correctness, retries, and duplicates. Stop on missing confirmation or semantic disagreement.
5. **Expansion and cleanup gate:** expand gradually only while business totals and technical signals reconcile. Remove cents/old-version support only after consumer inventory is complete, partners confirm retirement, and 365-day history is migrated or expired.

Business-level signals must include charged amount versus order total, revenue totals by event version, and partner-reported order amount/status. Transport success alone cannot establish semantic correctness.

## Rollback Constraints

- Disabling the feature is reversible for future production once flag propagation is proven.
- Checkout, billing, webhook, and query code can be rolled back only while the old contract and cents data remain available and no deployed consumer depends exclusively on the revised form.
- Events already emitted with micros semantics at tag 2 are ambiguous without provenance. Replaying or re-decoding them under either schema can create a 10,000x error; they require quarantine plus forward repair or an evidence-backed canary cohort/version mapping.
- Gateway charges already made are external side effects. Incorrect charges require reconciliation/refund/forward repair rather than code rollback.
- Partner deliveries are externally observed and cannot be recalled. Correction events or partner-specific repair may be required.
- Destructive archive conversion or dropping `amount_cents` before all 365-day history is migrated removes simple rollback. Retain original values and version metadata through reconciliation.
- Once `BACKORDERED` is externally produced, reverting code does not make stored or partner-held value 2 disappear; old consumers must remain tolerant or records require normalization.
- Never return to the proposed field-2 semantic reuse after additive/versioned migration begins; the same bytes would remain ambiguous across deployed clients and history.

## Open Decisions and Unknowns

| Unknown or decision | Impact | Smallest resolving action |
|---|---|---|
| Is protobuf actually used on `bus.publish`, and which generated runtimes deserialize it? | Determines whether V2 is missing-property failure, 10,000x semantic misread, or both across consumers. | Inspect bus serializer/subscriber bindings and inventory generated client versions. |
| What has the 5% canary already emitted, and can the flag be shut off immediately? | Sets current financial/data/partner exposure and containment timing. | Query event telemetry/archive by shape and test flag propagation to zero. |
| How are events written into `archived_order_events`? | Determines reject, null, or unit-corruption behavior and backfill design. | Inspect archive writer and table schema; sample V2-cohort records with raw payload provenance. |
| Which consumers subscribe to `orders.placed`? | Unknown consumers can silently corrupt behavior during mixed emission or replay. | Obtain broker subscription registry and map each deployment owner/version. |
| Which partners receive the webhook and what do they accept? | External compatibility and rollback remain unproven. | Request per-partner contract/version confirmation and test payloads. |
| How do consumer runtimes handle enum value 2? | `BACKORDERED` may be rejected, dropped, or mishandled. | Generate/run a value-2 cross-runtime compatibility check for every deployed language. |
| What is the authoritative V2 migration design? | Safe coexistence cannot begin while tag 2 has conflicting meanings. | Contract owner chooses additive field versus versioned message/topic and records precedence/version rules. |
| Are canary events durably identifiable after emission? | Without provenance, existing field-2 records cannot be safely normalized or replayed. | Establish whether cohort, timestamp, producer revision, raw field name, or metadata can identify them; otherwise quarantine. |
| What owners and deployment order apply? | Required work cannot be scheduled or gated. | Assign owner types above to named teams and document deployment units/dependencies. |

Every material delta reaches either a concrete effect (billing charge input, analytics revenue, replayed payload, or partner post) or a named evidence boundary. No implementation was changed while producing this report.
