# Behavior Case 3 Audit

## Run Identity

- **Fresh identity:** `/root/api_change_impact_mapper/api_body_case_3`
- **Model and reasoning:** `gpt-5.6-sol`, medium
- **Target hash opened:** `7b5dfd09fae4349f467f13aa4d7a85ddae21b831eef6dee0f93421d03cd4876e`
- **Fixture tree hash before and after:** `91b7e73f05a0062d2e91dadba11097a51fda8f7296d3036b973070b40922a268`
- **Output hash:** `781d89ac2f72a77bfbf1b6c87b7fb4c66f73894723d3e0a6eeac87e3715e5589`
- **Evidence:** `native-evidence.log`

## Ground-Truth Comparison

- PASS — treats positional reordering, asynchronous delivery, missing-result removal, and error behavior as independent compatibility changes.
- PASS — traces through the re-export to concrete router behavior and follows CLI, service, legacy adapter, TypeDoc generation, package identity, backend configuration, and external package/plugin boundaries.
- PASS — distinguishes source, runtime, behavioral, backend-conditional, documentation, and cross-version impacts.
- PASS — migration sequencing preserves a truthful legacy surface, resolves backend behavior, aligns generated artifacts, migrates callers and adapters, stages publication, and delays removal to a major-version gate.
- PASS — rollback addresses new consumers against old runtimes, propagated asynchronous caller contracts, published packages, and possible backend data effects.

## Completion-Criterion Audit

All ten observable criteria pass. Unknown external consumers remain explicit, concrete behavior is separated from declaration authority, and the fixture stayed unchanged.

## Result

PASS
