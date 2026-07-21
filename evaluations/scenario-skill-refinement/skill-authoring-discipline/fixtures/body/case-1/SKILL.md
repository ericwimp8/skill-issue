---
name: database-migration-review
description: Database migration safety review guidance. Use when reviewing schema or data migrations and their deployment or rollback plans.
---

# Database Migration Review

Establish the migration's intended change, affected data, deployment sequence, and compatibility with every application version that may run during rollout.

Prioritize failure modes that could cause data loss, prolonged blocking, availability impact, or mixed-version incompatibility. Trace the relevant schema changes, data movement, transaction boundaries, and operational controls to concrete evidence.

State unsupported assumptions and unresolved risks explicitly. Require a concrete mitigation or verification step for any risk that prevents a safe rollout decision.

## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/rollback.md`: Rollback safety policy for irreversible migrations. Use when a proposed rollback could encounter destroyed or irreversibly transformed data.
