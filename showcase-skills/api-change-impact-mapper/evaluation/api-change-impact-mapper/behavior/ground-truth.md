# Behavior Ground Truth

## Case 1 — HTTP Account Contract

- Removing `display_name`, requiring `profile.label`, narrowing `status`, and changing `404` to `200` plus `null` have distinct semantic effects.
- The API handler is a producer; the generated TypeScript client and mobile decoder are consumers; the compatibility adapter maps only outgoing v2 responses; stored JSON and cache keys preserve old shape; generation ownership belongs to the OpenAPI source and command.
- The public SDK and old mobile releases are external or independently deployed consumers and cannot be declared safe from repository-only evidence.
- A safe sequence must account for additive producer support, consumer tolerance, stored records, regeneration, cross-version validation, observation of decode/null handling, and rollback after new-shape writes.

## Case 2 — Order Event Contract

- Adding a required `warehouse_id`, changing amount units, adding an enum value, and reusing a protobuf field number create different producer, consumer, replay, and rollback risks.
- The checkout publisher and replay job produce events; billing, analytics, notification, and partner webhook paths consume or forward them; archived events outlive deployment; a flag changes event version.
- Unknown partner consumers remain an explicit evidence boundary.
- Rollback must consider newly emitted unit semantics, archived v2 events, and field-number reuse rather than assuming code rollback restores correctness.

## Case 3 — Library Public Interface

- Reordering positional parameters, changing sync return to async, and changing error behavior affect source, runtime, and behavioral compatibility independently.
- The facade re-export is not the concrete implementation; internal CLI and service call sites, plugin adapter, generated documentation, configuration-selected backend, and external package users require distinct treatment.
- Repository search can establish checked internal uses but cannot enumerate external package consumers.
- Migration ordering must address additive compatibility or a major version, adapter behavior, caller awaiting, documentation generation, validation across backends, staged publication, and rollback after downstream adoption.
