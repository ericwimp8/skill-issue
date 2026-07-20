# Generation Validation

## Structural Result

- **Validator:** `skill-creator/scripts/quick_validate.py`
- **Target:** `showcase-skills/dependency-upgrade-planner/skill/dependency-upgrade-planner/`
- **Result:** PASS — `Skill is valid!`
- **Folder and frontmatter name:** `dependency-upgrade-planner`
- **Resources:** OpenAI interface metadata only; no project-specific reference, script, or asset is required.
- **Metadata:** contains display and short-description interface copy; no unrequested default prompt.

## Generated Content Hashes

| File | SHA-256 |
| --- | --- |
| `SKILL.md` | `42b99ec21c5dee1e081b7cab161fbad01f91d9b55b4d0513aa298bb0e8588096` |
| `agents/openai.yaml` | `419b439853e628d185ff5be9796d853ae540e72e71a996c42fad0dd78393bde1` |

## Intake Criteria Walkthrough

- PASS: confirms package identity, versions, ecosystem, workspace, constraints, and authority.
- PASS: starts from production manifests, lockfiles, configuration, platform declarations, and concrete usages.
- PASS: traces wrappers to concrete effects and treats tests as later validation surfaces.
- PASS: distinguishes direct, transitive, build-tool, runtime, and platform effects while preserving overlaps.
- PASS: requires authoritative, version-applicable upstream evidence and both-sided compatibility checks.
- PASS: orders prerequisites and migration work by dependency edges and behavior owners.
- PASS: attaches validation evidence, rollback points, stops, exclusions, and unresolved risks.
- PASS: labels project facts, upstream facts, implications, and unknowns.
- PASS: preserves the planning-only boundary and prohibits unsupported compatibility claims.
- RUNTIME PROOF REQUIRED: natural description selection and generalized execution across varied ecosystems and effect classes.

Structural validation is generation evidence only. Runtime evidence is owned by `showcase-skills/dependency-upgrade-planner/evaluation/dependency-upgrade-planner/`.
