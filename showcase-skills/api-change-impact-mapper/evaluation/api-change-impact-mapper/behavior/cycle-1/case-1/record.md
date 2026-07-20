# Behavior Case 1 Audit

## Run Identity

- **Fresh identity:** `/root/api_change_impact_mapper/api_body_case_1`
- **Model and reasoning:** `gpt-5.6-sol`, medium
- **Target hash opened:** `7b5dfd09fae4349f467f13aa4d7a85ddae21b831eef6dee0f93421d03cd4876e`
- **Fixture tree hash before and after:** `3354c05dffb6aa15259fe0aa276283c0cb152b131641857dd1e9b4f6d2e21c4c`
- **Output hash:** `8e440e171d0461f6056abca78c59f8fd0b4acdf198df3868735eb4bb463129f4`
- **Evidence:** `native-evidence.log`

## Ground-Truth Comparison

- PASS — identifies field nesting, status narrowing, and missing-response semantics as distinct deltas.
- PASS — traces the handler producer, v2 response adapter, stale generated client and generator provenance, mobile consumer, stored old shape, cache key, route configuration, public package, and unknown deployed mobile/external consumers.
- PASS — classifies paths independently, including conditional stored-record compatibility and breaking `closed`, mobile, generated-client, and configured-route paths.
- PASS — migration order retains v1, resolves `closed`, regenerates clients, migrates mobile and external consumers, gates routing, and delays cleanup.
- PASS — validation, rollout signals, and rollback constraints cover null handling, enum/status data, version mix, cache/data effects, and deployed clients.

## Completion-Criterion Audit

All ten observable criteria pass. Claims use repository-relative evidence, material limitations and checked absences are explicit, and the fixture stayed unchanged.

## Result

PASS
