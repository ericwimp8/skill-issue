---
name: api-change-impact-mapper
description: Evidence-driven mapping of repository and boundary impacts from API or contract changes. Use when assessing compatibility and migration work for a proposed or observed API, schema, event, protocol, or public-interface change.
---

# API Change Impact Mapper

## Establish the Change Boundary

- Read the supplied old and new contracts before tracing usages. Record semantic deltas in names, shapes, types, constraints, defaults, optionality, ordering, lifecycle, errors, and behavioral guarantees; do not rely on textual difference alone.
- Identify the changed contract's authority, versions, direction, and deployment unit. Mark ambiguous or unavailable contract meaning as unknown rather than choosing a convenient interpretation.
- Keep confirmed repository facts, evidence-supported inferences, assumptions, and unknown external behavior distinct. State the scope and revision examined.
- Keep the work read-only. Do not edit generated files, schemas, migrations, configuration, or implementation while producing the map.

## Trace the Concrete Impact Graph

Trace outward from every changed contract element and inward from each observed use until reaching concrete effects or an explicit evidence boundary.

- Find producers that create, serialize, publish, return, write, or expose the changed value and consumers that parse, validate, branch on, display, persist, forward, or depend on it.
- Follow aliases, re-exports, wrappers, adapters, mappers, serializers, validation layers, transport bindings, and dependency injection to their concrete implementations. A declaration or search match is a lead, not proof of impact.
- Identify generated clients, models, bindings, or documentation together with their source schema and generation command or owner. Treat a generated artifact as a dependent surface, not the semantic owner.
- Trace stored representations, historical records, caches, queues, replay or backfill paths, indexes, and migrations when old data can outlive deployment.
- Trace configuration, feature flags, version negotiation, capability detection, environment values, routing, and deployment topology that change which contract version runs.
- Locate external boundaries such as public clients, webhooks, partner integrations, plugins, packages, mobile releases, scheduled jobs, and independently deployed services. Preserve unobservable consumers as unknown; repository silence does not prove their absence.
- Record relevant negative searches and inapplicable surface categories so readers can distinguish checked absence from unexamined scope.

## Classify Compatibility Per Path

Classify each affected producer-to-consumer or storage path independently.

- **Compatible:** the path remains correct across the stated old/new versions without coordinated action or hidden assumptions.
- **Conditionally compatible:** correctness depends on a named condition such as deployment order, tolerant parsing, defaults, version negotiation, feature state, data age, regeneration, or a bounded consumer set.
- **Breaking:** a supported path can reject, misinterpret, lose, corrupt, or expose incorrect behavior under the proposed transition.

Support every classification with contract and repository evidence. Name the exact versions, conditions, and confidence. Do not apply one global label when paths differ, equate compilation with behavioral compatibility, or downgrade an unknown external boundary to compatible.

## Build the Migration Map

Derive actions from the impact graph rather than proposing a generic rollout.

1. Group work by the component or team that owns the concrete behavior, contract, generator, stored data, deployment configuration, or external relationship. Use an explicit owner type when a named owner is unavailable.
2. Order actions by prerequisites. Identify dual-read, dual-write, additive-first, regeneration, backfill, negotiation, deprecation, consumer-release, and cleanup phases only where the traced paths require them.
3. For each action, state the affected path, compatibility risk addressed, dependency, completion evidence, and safe coexistence window.
4. Define focused validation at the actual boundaries: contract checks, producer and consumer integration, generated-artifact drift, historical-data or replay tests, configuration matrices, cross-version tests, and external confirmation as applicable.
5. Define rollout observation using signals tied to the changed behavior, including reject or parse failures, version mix, fallback use, data divergence, queue or replay failures, and business-level correctness.
6. Define rollback from the post-change state. Name irreversible writes, destructive migrations, expired compatibility windows, deployed clients, schema reuse, or externally observed behavior that requires forward repair instead of simple reversal.

## Deliver the Impact Report

Use these sections:

- **Decision Summary:** proposed change, overall risk, highest-risk paths, and whether evidence supports proceeding.
- **Scope and Evidence:** contracts, revision, inspected areas, methods, limitations, negative searches, and unknown consumers.
- **Contract Delta:** old-to-new semantic changes with authority and ambiguity noted.
- **Impact Graph:** producer, transformation, consumer, persistence, configuration, generation, and external-boundary paths with repository-relative evidence.
- **Compatibility Matrix:** one row per material path with classification, versions or conditions, rationale, confidence, and unresolved evidence.
- **Migration Sequence:** ordered owner-oriented actions, dependencies, coexistence windows, and completion evidence.
- **Validation Plan:** checks at contract, code, data, deployment, and external boundaries.
- **Rollout and Observation:** stages, gates, signals, and stop conditions.
- **Rollback Constraints:** reversible actions, unsafe reversals, retained compatibility requirements, and forward-repair triggers.
- **Open Decisions and Unknowns:** unresolved owner choices or evidence gaps, their impact, and the smallest action that can resolve each.

Before concluding, verify that each material delta reaches a concrete effect or a named evidence boundary, every classification is conditional on the evidence actually available, and the plan changes no implementation.
