# Description Trial 3 Record

## User Task

The exact task is retained in `prompt.md`.

## Run Identity

- **Fresh identity:** `/root/dependency_upgrade_planner/dep_desc_3`
- **Model and reasoning:** `gpt-5.6-sol`, medium
- **Fixture:** `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/fixtures/pydantic-service/`
- **Selected skill:** `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/SKILL.md`
- **Selection:** The target description naturally matches source-backed dependency-upgrade planning without dependency changes. The other candidates concern explicit skill intake, structural change placement, or prompt authoring.

## Sources

- Fixture production sources: `pyproject.toml`, `poetry.lock`, `service/models.py`, and `service/api.py` under the fixture path.
- Live authoritative source: `https://docs.pydantic.dev/latest/migration/`.
- Live authoritative source: `https://docs.pydantic.dev/latest/api/config/#pydantic.config.ConfigDict.from_attributes`.
- Live authoritative source: `https://docs.pydantic.dev/latest/version-policy/`.
- Live authoritative source: `https://docs.pydantic.dev/latest/install/`.
- Live authoritative source: `https://fastapi.tiangolo.com/how-to/migrate-from-pydantic-v1-to-pydantic-v2/`.
- Live authoritative source: `https://fastapi.tiangolo.com/release-notes/#01000`.

## Eight-Criterion Contract Audit

1. **Natural selection — PASS:** The unmodified migration-planning prompt selected and loaded `dependency-upgrade-planner`; `native-evidence.log` records description hashes and the complete selected-file hash before output.
2. **Project inspection — PASS:** The plan cites the manifest, lockfile, Python range, both production modules, every concrete Pydantic V1 use, and the absence of further configuration/build/deployment/test surfaces in the fixture.
3. **Effect classification — PASS:** Direct, transitive/peer, runtime, framework, platform, and absent build-tool effects are separated with evidence boundaries.
4. **Authoritative upstream evidence — PASS:** Live Pydantic and FastAPI owner documentation is linked near claims, with V2/current-stable, 0.100.0 bridge, and current FastAPI applicability distinguished.
5. **Breaks and unknowns — PASS:** The FastAPI `<2.0.0` blocker, V1 configuration/loading APIs, validation semantics, exact-target ambiguity, lock ownership, and missing consumer evidence remain explicit.
6. **Dependency and ownership order — PASS:** Exact pair and lock ownership precede framework compatibility, model-owner migration, resolution, and integrated service validation.
7. **Gates and recovery — PASS:** Every state-changing stage has validation evidence, rollback scope, and stop conditions; exclusions and the smallest safe first stage are explicit.
8. **Evidence labels and planning boundary — PASS:** Observed project facts, authoritative upstream facts, reasoned implications, and unknowns are labeled; the fixture remains unchanged and no successful runtime claim is made.

## Cleanup Ownership

This trial owns only `prompt.md`, `native-evidence.log`, `output.md`, and `record.md` in `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/description/round-1/trial-3/`. The evaluation coordinator may remove them after retaining required evidence. No cleanup is required in the fixture or elsewhere.

## Result

**PASS**
