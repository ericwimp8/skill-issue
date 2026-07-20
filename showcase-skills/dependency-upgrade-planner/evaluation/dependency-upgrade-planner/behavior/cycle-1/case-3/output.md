# Synthetic Profile Service: Pydantic 2.x Migration Plan

## Upgrade Contract

- **Observed project fact:** The scoped workspace is `fixtures/pydantic-service/`. Its manifest requires Python `>=3.10,<3.12` and pins `pydantic==1.10.15` and `fastapi==0.95.2` (`fixtures/pydantic-service/pyproject.toml`).
- **Observed project fact:** The lock records the same direct versions and records FastAPI 0.95.2's Pydantic requirement as `>=1.6.2,<2.0.0` (`fixtures/pydantic-service/poetry.lock`).
- **Requested change:** Move the service to Pydantic 2.x while preserving its profile validation, row-to-model conversion, request handling, response serialization, and schema behavior.
- **Constraint:** This document is planning-only. It does not edit constraints or source, run a solver, install packages, apply `bump-pydantic`, regenerate the lock, or claim successful execution.
- **Unknown:** “2.x” does not identify an exact Pydantic release, and the allowed source routes do not provide a complete release-by-release compatibility matrix or package metadata for selecting an exact FastAPI/Pydantic pair. Exact pins are a blocking decision before dependency mutation.

## Current Graph and Concrete Usage

| Edge or surface | Classification | Evidence and migration relevance |
| --- | --- | --- |
| Service → Pydantic 1.10.15 | Direct, runtime | Exact manifest and lock pins. `service/models.py` imports `BaseModel` and owns all concrete Pydantic calls. |
| Service → FastAPI 0.95.2 | Direct, runtime | Exact manifest and lock pins. `service/api.py` creates `FastAPI`, uses `Profile` as the request annotation and `response_model`, and returns that model. |
| FastAPI 0.95.2 → Pydantic `<2.0.0` | Transitive compatibility edge, runtime | The lock records `pydantic = ">=1.6.2,<2.0.0"`; the current framework cannot coexist with Pydantic 2 under this graph. |
| Python `>=3.10,<3.12` | Platform/runtime prerequisite | Declared in `pyproject.toml`; it is within the Python 3.10+ examples in the live FastAPI migration guide. No deployment image or CI runtime declaration exists in the fixture. |
| Lock/solver workflow | Build-tool | `poetry.lock` exists, but no Poetry version, lock metadata, build script, or CI configuration is supplied. The exact solver command and lock format must be confirmed rather than inferred. |
| Pydantic 2 transitive packages | Transitive | The current synthetic lock lists only Pydantic and FastAPI. The future transitive graph is unknown until an approved solver preview for the selected exact pair. |

Concrete Pydantic-owned behavior in `fixtures/pydantic-service/service/models.py`:

1. `Profile(BaseModel)` validates one required `name: str` field.
2. Inner `class Config` sets `orm_mode = True`.
3. `profile_from_row()` calls `Profile.from_orm(row)`.
4. `parse_profile()` calls `Profile.parse_obj(payload)`.

Concrete framework integration in `fixtures/pydantic-service/service/api.py`:

1. `Profile` validates the POST `/profiles` request body.
2. `response_model=Profile` makes FastAPI validate/serialize the response and generate OpenAPI schema.
3. The endpoint returns the received `Profile`, coupling framework behavior to the selected Pydantic generation.

No other production configuration, platform declaration, generated artifact, adapter, deployment package, or test is present in the supplied fixture.

## Applicable Upstream Requirements

- **Authoritative upstream fact:** Pydantic's live V1-to-V2 migration guide says `parse_obj()` is renamed to `model_validate()`; retained old methods are deprecated and emit warnings. It also says `from_orm()` is deprecated in favor of `model_validate()` when `from_attributes=True` is configured. [Pydantic V2 migration guide](https://docs.pydantic.dev/latest/migration/)
- **Authoritative upstream fact:** The same guide deprecates inner `Config`, introduces `model_config`, and renames `orm_mode` to `from_attributes`. This directly applies to `Profile`. [Pydantic V2 migration guide](https://docs.pydantic.dev/latest/migration/)
- **Authoritative upstream fact:** Pydantic V2 changes validation and serialization semantics beyond method names, including model equality, input-type preservation, standard-type handling, and JSON schema generation. Only concrete `name` validation, attribute extraction, and FastAPI serialization/schema paths need project-specific proof here. [Pydantic V2 migration guide](https://docs.pydantic.dev/latest/migration/)
- **Authoritative upstream fact:** FastAPI first supported either Pydantic generation in 0.100.0. FastAPI 0.119.0 added temporary partial `pydantic.v1` support; 0.126.0 and 0.128.0 progressively removed V1 support, and current versions require V2 models. [FastAPI migration guidance](https://fastapi.tiangolo.com/how-to/migrate-from-pydantic-v1-to-pydantic-v2/)
- **Reasoned implication:** FastAPI must be upgraded before or atomically with the Pydantic pin because the resolved 0.95.2 peer range excludes Pydantic 2. Version 0.100.0 is a documented lower boundary for dual-generation support, not an automatic recommendation for the final framework pin.
- **Authoritative upstream fact:** The Pydantic guide offers `pydantic.v1` as a transition namespace, while FastAPI documents mixed-generation support only in 0.119.0–0.127.x and says current FastAPI releases require V2 models. [Pydantic V2 migration guide](https://docs.pydantic.dev/latest/migration/) [FastAPI migration guidance](https://fastapi.tiangolo.com/how-to/migrate-from-pydantic-v1-to-pydantic-v2/)
- **Reasoned implication:** This four-use fixture is small enough for a direct V2 model migration. Introducing a temporary V1 compatibility layer would add a version window and rollback complexity without an observed need; retain it only as a contingency after selecting a FastAPI version that explicitly supports it.

## Impact Matrix

| Effect | Planned impact | Evidence boundary |
| --- | --- | --- |
| Direct | Change the Pydantic constraint from 1.10.15 to an approved exact 2.x pin and update `Profile` to V2 configuration and methods. | Exact target remains unknown. |
| Direct/framework | Upgrade FastAPI from 0.95.2 to an approved exact release whose declared metadata accepts the selected Pydantic 2.x release. | The live guide establishes support boundaries, not the exact final pair. |
| Transitive | Refresh Pydantic/FastAPI transitive resolutions only after pair selection and source changes are defined. | Future transitive packages and conflicts require solver evidence. |
| Build-tool | Regenerate `poetry.lock` with the repository's confirmed Poetry/toolchain version. | Tool version and canonical command are absent. |
| Runtime | Re-prove dict validation, attribute extraction, request rejection, response serialization, and OpenAPI schema. | No fixture tests or runtime transcript exists. |
| Platform | Validate on declared Python 3.10 and 3.11 bounds. | Deployment/CI platform is unknown; Python 3.12+ is outside the manifest. |

## Dependency-Ordered Migration Stages

### Stage 0 — Freeze Baseline Semantics

**Owner:** service behavior contract.

Before changing dependency state, add or define executable checks that capture:

- `parse_profile({"name": "Ada"})` returns a `Profile` with `name == "Ada"`;
- missing `name`, `None`, and a representative non-string value are rejected or coerced exactly as the accepted baseline contract specifies;
- `profile_from_row()` accepts an object with a `name` attribute and rejects an object without it;
- POST `/profiles` accepts a valid body, returns the expected JSON/status, and rejects invalid bodies with the intended status/error shape;
- generated OpenAPI retains the `/profiles` request and response schema contract.

**Gate:** A product owner or existing external contract must decide whether current V1 coercion and exact 422 error payload/schema text are compatibility requirements. The fixture alone cannot decide this.

**Rollback:** No dependency state changes occur. Discard only newly proposed baseline-test work if its asserted contract is rejected.

**Stop:** Do not proceed if intended coercion, error-shape, or schema compatibility is unresolved.

### Stage 1 — Select a Compatible Version Set

**Owner:** dependency manifest and runtime platform.

1. Select an exact Pydantic 2.x release, not an unbounded floating target.
2. Select an exact FastAPI release and verify its live package metadata accepts that Pydantic release and Python 3.10/3.11.
3. Confirm the repository's canonical Poetry version/lock command and whether any deployment environment imposes narrower Python constraints.
4. Run a read-only solver preview if the package manager supports one; record the full proposed direct and transitive graph.

**Gate:** Solver evidence must show one graph satisfying Python `>=3.10,<3.12`, the selected FastAPI pin, and the selected Pydantic 2.x pin. Review all direct and transitive deltas before approval.

**Rollback:** Since this stage is selection/preview only, retain the current manifest and lock unchanged and reject the candidate pair.

**Stop:** Stop on solver conflict, unsupported Python bound, unavailable metadata, unexplained transitive change, or a requirement to mix V1/V2 models without a selected FastAPI release documented to support that bridge.

### Stage 2 — Migrate the Model Behavior Owner

**Owner:** `fixtures/pydantic-service/service/models.py`.

For the direct V2 path:

1. Replace inner `Config.orm_mode = True` with V2 `model_config` using `from_attributes=True` (normally via `ConfigDict`).
2. Replace `Profile.from_orm(row)` with `Profile.model_validate(row)`; attribute extraction continues only because `from_attributes=True` is set.
3. Replace `Profile.parse_obj(payload)` with `Profile.model_validate(payload)`.
4. Keep the `name: str` field unchanged unless Stage 0 establishes that V2's accepted inputs differ from the required service contract; if they differ, make the coercion/rejection rule explicit at this behavior owner rather than accepting accidental migration behavior.

**Gate:** Run focused model checks on Python 3.10 and 3.11. Require expected valid values, invalid-value failures, attribute extraction, serialized model output, and no Pydantic deprecation warnings from these paths.

**Rollback:** Restore `models.py` to `Config.orm_mode`, `from_orm`, and `parse_obj` while the old dependency pair remains installed. Do not combine reverted V1 source with a final V2-only framework environment.

**Stop:** Stop if semantic probes differ without an approved contract decision or if source requires undocumented compatibility shims.

### Stage 3 — Upgrade the Coupled Dependency Pair and Lock

**Owner:** `fixtures/pydantic-service/pyproject.toml` and `fixtures/pydantic-service/poetry.lock`.

1. Apply the approved exact FastAPI and Pydantic pins together; changing only Pydantic is invalid under the current FastAPI `<2.0.0` edge.
2. Regenerate the lock with the confirmed tool version only after Stage 2 source decisions are complete.
3. Inspect the new lock for the exact direct versions, Python markers, FastAPI's accepted Pydantic range, and every added/removed/changed transitive package.

**Gate:** A clean install from the regenerated lock must succeed on supported Python bounds. Import both libraries and record their runtime versions. Require no unresolved solver warning or unexpected source/registry selection.

**Rollback:** Restore both manifest pins and the entire prior lockfile as one unit, remove the replacement environment/cache created for the trial, and reinstall from the restored lock. Never hand-edit individual lock entries.

**Stop:** Stop if lock metadata does not prove the selected peer edge, the lock changes outside the reviewed graph, or clean recreation differs from the reviewed resolution.

### Stage 4 — Validate Framework Runtime Semantics

**Owner:** FastAPI route integration in `fixtures/pydantic-service/service/api.py`.

No route edit is currently implied, but the route is a mandatory validation surface because it consumes and emits `Profile`.

**Gate:** On both declared Python minors, run the Stage 0 endpoint and OpenAPI checks plus application import/startup. Compare request parsing, response JSON, response-model filtering/validation, status codes, validation-error payloads, and schema against the accepted baseline. Treat any difference as a migration decision, not proof of compatibility.

**Rollback:** Restore the old source, manifest, and lock snapshot together; recreate the old environment and rerun the baseline probes to confirm restoration.

**Stop:** Stop on startup/import failure, route behavior drift, OpenAPI incompatibility, deprecation warnings on exercised paths, or dependency-version mismatch at runtime.

### Stage 5 — Release Gate

**Owner:** service delivery owner (not represented in the fixture).

Run the repository's full static checks, tests, clean build/package step, and deployment smoke check when those commands and environment are identified. Record the exact commands, versions, and results.

**Gate:** All focused and full checks pass from a clean locked environment, and the delivery owner approves any intentional validation/schema change.

**Rollback:** Preserve the last deployable artifact and old lock/source snapshot. Roll back the whole coupled release on production validation or schema regressions.

**Stop:** Absence of CI, packaging, deployment, or observability evidence prevents a release-readiness claim; it does not prevent completing the source migration locally.

## Unresolved Risks and Follow-up Questions

1. **Unknown:** Which exact Pydantic 2.x and FastAPI versions are approved, and what support/security policy governs that choice?
2. **Unknown:** Does any client depend on Pydantic V1 coercion, exact validation-error details, or exact generated OpenAPI?
3. **Unknown:** What concrete row type reaches `profile_from_row()`, and can its attribute access raise or trigger lazy-loading behavior?
4. **Unknown:** Which Poetry version and command own this synthetic lock, and is the abbreviated lock representative of a real resolvable artifact?
5. **Unknown:** What CI, deployment image, ASGI server, test framework, and release checks exist outside the supplied fixture?
6. **Reasoned risk:** Upgrading FastAPI can introduce framework changes beyond Pydantic compatibility. Review the chosen FastAPI release span before implementation; the permitted source ledger does not route a FastAPI release-notes audit.
7. **Reasoned risk:** Pydantic V2 validation and JSON schema changes may alter request acceptance, error payloads, or generated API contracts even when the endpoint source is unchanged.

## Exclusions

- No dependency installation, manifest edit, lock regeneration, source migration, codemod, test execution, build, or deployment.
- No unrelated fixture, React, Vite, repository infrastructure, or documentation work.
- No claim that FastAPI 0.100.0 is the best final target, that an unspecified latest release supports the declared Python range, or that tests/build/runtime currently pass.

## Smallest Safe First Stage

Stage 0 is the smallest safe first stage: define and obtain approval for the concrete model, endpoint, error, and OpenAPI semantics before touching dependency state. Implementation may begin only after those baseline probes are reproducible and Stage 1 supplies authoritative metadata plus a successful read-only resolution for one exact FastAPI/Pydantic pair on Python 3.10 and 3.11.
