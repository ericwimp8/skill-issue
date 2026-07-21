---
name: database-migration-review
description: Database migration safety review constraints and workflow. Use when reviewing schema or data migrations, deployment sequencing, or rollback plans.
---

# Database Migration Review

- Trace the migration, application compatibility, and deployment sequence end to end.
- Evaluate operational risk from the database's concrete behavior, affected data volume, and production topology.
- Require bounded execution, recovery, and observability for long-running or retryable work.
- Classify reversibility before approving a rollback strategy.

## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/rollback.md`: Irreversible migration rollback policy. Use when a migration destroys or irreversibly transforms data.
