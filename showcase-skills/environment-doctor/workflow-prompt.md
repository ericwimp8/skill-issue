# Environment Doctor Workflow Prompt

Create a reusable skill named `environment-doctor` under `<repo-root>/showcase-skills/environment-doctor/` by running the complete current project-local Skill Issue workflow from intake through generation, evaluation, and refinement.

The skill must provide a genuinely useful deterministic, non-destructive diagnostic script that inspects a supplied development environment for selected tool availability and versions, executable resolution, relevant path ordering, explicitly selected environment variables, and common version-configuration mismatches. It must produce clear human-readable findings plus structured evidence for follow-up, preserve unknowns, redact or omit sensitive values, avoid changing the environment, and provide safe remediation and verification guidance. Keep platform claims bounded to behavior actually supported and tested.

Retain the governed intake, generation, skill, resource, script-validation, evaluation, refinement, fixture, evidence, and audit artifacts in the showcase workspace. Use synthetic or repository-owned fixtures and public authoritative sources. Test the script directly for determinism, safety, source preservation, error handling, and privacy. Run the required structural, formatting, diff, hash, and privacy checks, and report exact paths, evidence, refinements, blockers, and workflow weaknesses.
