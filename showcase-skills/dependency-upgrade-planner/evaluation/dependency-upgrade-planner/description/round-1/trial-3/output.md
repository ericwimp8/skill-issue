# Pydantic 1.10 to 2.x Migration Plan

## Upgrade Contract

- **Observed project fact:** The service directly pins `pydantic==1.10.15` and `fastapi==0.95.2` in `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/pydantic-service/pyproject.toml`. The declared runtime range is Python `>=3.10,<3.12`.
- **Observed project fact:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/pydantic-service/poetry.lock` resolves those same versions and records FastAPI's Pydantic constraint as `>=1.6.2,<2.0.0`.
- **Requested change:** Move the service to Pydantic 2.x through a reviewable, dependency-ordered migration. Planning authority is read-only: no manifest, lockfile, source, generated artifact, or environment may be changed by this plan.
- **Unknown:** “2.x” does not select an exact Pydantic release or constraint policy. The implementation owner must choose an exact target before resolution. The live Pydantic documentation identifies 2.12 as the current stable documentation line and says V2 minor releases avoid intentional breaking changes, while allowing some behavior-affecting fixes and schema/error changes ([Pydantic version policy](https://docs.pydantic.dev/latest/version-policy/)).
- **Unknown:** The fixture contains a file named `poetry.lock`, but no `[tool.poetry]` declaration or package-manager instructions. Confirm the authoritative package manager and lock-generation command before changing dependency state.

## Current Graph and Concrete Usage

| Relationship | Evidence | Classification | Migration consequence |
| --- | --- | --- | --- |
| Service → Pydantic 1.10.15 | `pyproject.toml` direct pin and `poetry.lock` resolution | Direct, runtime | The requested major upgrade changes the validation model used by both helper functions and the HTTP contract. |
| Service → FastAPI 0.95.2 | `pyproject.toml` direct pin and `service/api.py` import/startup | Direct, runtime/framework | The framework owns request parsing, response-model serialization, and OpenAPI generation around `Profile`. |
| FastAPI 0.95.2 → Pydantic `<2.0.0` | `poetry.lock` dependency constraint | Transitive/peer compatibility edge, runtime | Pydantic 2 cannot resolve while the locked FastAPI constraint remains. FastAPI compatibility is a prerequisite, not follow-up cleanup. |
| `Profile` → V1 model configuration | `service/models.py` uses inner `class Config` with `orm_mode = True` | Direct API/configuration | Replace the model-owned configuration with the V2 `model_config`/`from_attributes` equivalent. |
| Helper functions → V1 loading APIs | `service/models.py` calls `Profile.from_orm(row)` and `Profile.parse_obj(payload)` | Direct API/runtime semantics | Move both paths to V2 validation APIs and preserve their distinct mapping-versus-attribute behavior. |
| `/profiles` → `Profile` | `service/api.py` uses `Profile` as request annotation, return annotation, and `response_model` | Framework/runtime contract | Validate body acceptance, response serialization, validation errors, and generated OpenAPI after the coordinated framework/model migration. |
| Python runtime → dependencies | `pyproject.toml` permits Python 3.10 and 3.11 | Platform/runtime | Pydantic's current installation guide requires Python 3.9+, so the declared range is eligible; the selected FastAPI/Pydantic pair still needs resolver proof on both declared Python minors. |

No build-tool dependency, deployment declaration, generated source, override, patch, registry configuration, or test suite is present in the fixture. The lockfile is sparse and does not enumerate Pydantic 2's runtime dependencies, so it cannot establish the final transitive graph.

## Applicable Upstream Requirements

- **Authoritative upstream fact:** FastAPI 0.100.0 introduced support for both Pydantic v1 and v2. Current FastAPI guidance also says recent FastAPI versions require Pydantic v2, so the final pair must be selected together rather than assuming every FastAPI version is a bridge ([FastAPI migration guidance](https://fastapi.tiangolo.com/how-to/migrate-from-pydantic-v1-to-pydantic-v2/), [FastAPI 0.100.0 release notes](https://fastapi.tiangolo.com/release-notes/#01000)).
- **Authoritative upstream fact:** Pydantic V2 maps `parse_obj()` to `model_validate()`. It deprecates `from_orm()` in favor of `model_validate()` when `from_attributes=True` is configured ([Pydantic V2 migration guide](https://docs.pydantic.dev/latest/migration/)).
- **Authoritative upstream fact:** V2 prefers the `model_config` class attribute over the deprecated inner `Config` class. `ConfigDict.from_attributes` controls construction from object attributes ([Pydantic migration guide](https://docs.pydantic.dev/latest/migration/), [Pydantic configuration API](https://docs.pydantic.dev/latest/api/config/#pydantic.config.ConfigDict.from_attributes)).
- **Authoritative upstream fact:** Pydantic V2 changes validation behavior beyond renamed APIs, including input-type preservation, unions, optional/nullable fields, equality, JSON Schema, and validation-error details. Only the simple `name: str` field is observed here, but request, response, and error behavior still require executed comparison rather than a compatibility assumption ([Pydantic V2 migration guide](https://docs.pydantic.dev/latest/migration/)).
- **Authoritative upstream fact:** Current Pydantic installation documentation requires Python 3.9+ and identifies `pydantic-core`, `typing-extensions`, and `annotated-types` as dependencies. These will become transitive lockfile entries after authorized resolution ([Pydantic installation](https://docs.pydantic.dev/latest/install/)).

## Impact Matrix

| Effect | Confirmed impact | Evidence boundary |
| --- | --- | --- |
| Direct dependency | Pydantic pin and FastAPI pin must become a mutually compatible pair. | Exact target releases remain unknown until owner selection and resolver proof. |
| Transitive/peer | FastAPI 0.95.2 excludes Pydantic 2; Pydantic 2 adds core/runtime dependencies absent from the sparse lock. | Do not edit transitive packages directly; regenerate the authoritative lock after direct constraints are approved. |
| Runtime/model | `Profile` configuration and both loading helpers use V1 APIs. | Deprecation shims are migration aids, not evidence of completed V2 behavior. |
| Runtime/framework | FastAPI parses and serializes `Profile` at `POST /profiles` and derives OpenAPI from it. | No server, request, response, or schema behavior was executed during planning. |
| Platform | Python 3.10–3.11 falls within current Pydantic's documented Python floor. | The chosen FastAPI/Pydantic releases and full lock must resolve and run on both declared minors. |
| Build tool | None observed. | No build or packaging configuration beyond `pyproject.toml` is present. |

## Dependency-Ordered Migration Stages

### Stage 0 — Freeze the Target and Baseline Contract

**Owner:** dependency/runtime maintainer, with API behavior owner approval.

1. Choose an exact supported Pydantic 2 release and constraint policy, then choose a FastAPI release whose published metadata accepts it and whose Python range includes 3.10–3.11. Record whether a temporary dual-support FastAPI bridge is desired; FastAPI 0.100.x is an upstream-documented bridge, while current releases require V2 models.
2. Confirm which tool owns `poetry.lock` and its clean, reproducible lock command.
3. Define baseline cases before mutation: mapping input `{"name": "Ada"}`, an attribute-backed row with `.name`, representative invalid/missing `name` values, `POST /profiles` success/error responses, response serialization, and OpenAPI for the route.

**Gate:** Proceed only when the exact version pair resolves in a disposable environment for Python 3.10 and 3.11, the lock owner is known, and baseline expectations are approved. A resolver preview is evidence, not authorization to alter the fixture.

**Stop conditions:** No mutually compatible pair; selected FastAPI drops a required service behavior; package-manager ownership remains unknown; or baseline API/validation semantics lack an owner decision.

**Rollback:** None required because this stage is evidence collection and decision recording only.

### Stage 1 — Establish the FastAPI Compatibility Prerequisite

**Owner:** framework/dependency owner; affected surface: `pyproject.toml`, then the authoritative lockfile.

1. If the team chooses a dual-support bridge, update FastAPI first to the approved bridge release while retaining Pydantic 1.10.15; regenerate the lock and validate the existing service contract. This isolates the framework edge before V2 source work.
2. If the team chooses a current V2-only FastAPI release, treat FastAPI, Pydantic, model source, and lockfile changes as one coordinated stage; do not attempt an independently deployable FastAPI-only step.

**Gate:** Clean dependency resolution on Python 3.10 and 3.11; import/startup succeeds; existing model-helper and `/profiles` baseline checks pass; OpenAPI is captured and reviewed.

**Rollback:** Restore the prior FastAPI constraint and complete prior lockfile. Discard any environment generated from the failed candidate.

### Stage 2 — Migrate the Model at Its Behavior Owner

**Owner:** `Profile` in `service/models.py`; dependent surface: `service/api.py`.

1. Replace the inner V1 `Config` owner with V2 model configuration using `from_attributes=True`.
2. Replace `Profile.from_orm(row)` with `Profile.model_validate(row)` under that model configuration.
3. Replace `Profile.parse_obj(payload)` with `Profile.model_validate(payload)`.
4. Review `service/api.py` without adding parallel compatibility wrappers: the route should continue to consume the single migrated `Profile` owner for request parsing, response validation, serialization, and schema generation.

**Reasoned implication:** Keeping configuration and loading semantics on `Profile` avoids splitting V1/V2 behavior across route-local adapters. The apparent failures may surface at the route, but model validation belongs to the model.

**Gate:** Focused checks demonstrate mapping parsing, attribute-backed validation, invalid/missing-field behavior, and serialization. Treat warnings from legacy APIs/configuration as failures so deprecated shims cannot masquerade as completion.

**Rollback:** Restore `service/models.py` and any coordinated `service/api.py` change to the pre-stage revision. If this stage included the V2 dependency transition, restore the prior manifest and complete lockfile together.

### Stage 3 — Resolve Pydantic 2 and Reconcile the Runtime Graph

**Owner:** dependency/runtime maintainer; affected surfaces: `pyproject.toml`, `poetry.lock`, isolated environments.

1. Apply the approved Pydantic 2 constraint and final FastAPI constraint.
2. Regenerate the complete lock with the confirmed tool rather than hand-editing Pydantic or its transitive packages.
3. Inspect the resolved graph for the chosen Pydantic, FastAPI, `pydantic-core`, `typing-extensions`, `annotated-types`, Starlette, and any other newly resolved dependencies; retain exact versions as review evidence.

**Gate:** Reproducible clean resolution and installation on Python 3.10 and 3.11, with no V1-only FastAPI constraint or unexpected compatibility override.

**Rollback:** Restore `pyproject.toml` and the entire prior `poetry.lock`; discard V2 environments and caches used for the candidate.

### Stage 4 — Prove Service and Release Behavior

**Owner:** API/service owner, then release owner if a deployment surface exists outside this fixture.

1. Run focused model checks from Stage 2.
2. Exercise `POST /profiles` with valid, invalid, and missing-field payloads; compare status, response body, validation-error structure used by callers, and response serialization to the approved baseline.
3. Generate and compare OpenAPI, reviewing the `Profile` schema and route request/response declarations.
4. Run the repository's static checks and full service suite once their commands are identified. Run a startup/runtime probe on both declared Python minors.
5. Before release, discover deployment packaging and downstream consumers outside this sparse fixture; validate them if they exist.

**Gate:** All approved model, HTTP, schema, static, suite, startup, Python-version, and downstream checks provide recorded passing evidence. Any semantic difference requires explicit API-owner acceptance.

**Rollback:** Revert the complete migration change set—model/API source, manifest, and lockfile—to the last passing stage. Restore the prior deployment artifact if a release probe occurred.

## Unresolved Risks and Follow-Up Questions

1. **Unknown:** Which exact Pydantic 2 and FastAPI releases should be the maintained endpoint rather than a temporary bridge?
2. **Unknown:** Is `poetry.lock` authoritative, and what command generated its intentionally sparse content?
3. **Unknown:** Are tests, CI, deployment files, database row types, API consumers, or validation-error parsers present outside the fixture boundary?
4. **Reasoned implication:** The annotated `name: str` appears simple, but V2 conversion and error semantics can still alter accepted input or caller-visible errors; baseline and runtime evidence are mandatory.
5. **Unknown:** Does the team require byte-for-byte OpenAPI stability, or only reviewed semantic compatibility?

## Explicit Exclusions

- No dependency installation, manifest edit, lock regeneration, codemod, source migration, test creation, server run, build, deployment, or release is performed.
- No claim is made that the fixture currently installs, starts, builds, tests, or runs successfully.
- No compatibility is claimed for dependencies or consumers absent from the fixture.

## Smallest Safe First Stage

Perform **Stage 0 only**: approve the exact Pydantic/FastAPI pair, confirm lockfile ownership, and define the current model/API/OpenAPI baseline. Implementation should not begin until disposable resolver evidence covers Python 3.10 and 3.11 and the service owner has approved the behavioral comparison cases.
