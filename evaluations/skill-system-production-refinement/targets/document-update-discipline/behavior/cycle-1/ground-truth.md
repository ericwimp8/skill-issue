# Document Update Cycle 1 Ground Truth

## Case 1 — Release Approval

- Required meaning: every production release, including emergency releases, requires approval from two maintainers other than the author.
- Semantic owner: `Approval Policy`.
- Observation point: release checklist item 2.
- Related manifestation: emergency self-approval exception.
- Pass shape: complete rule lives once in policy; checklist defers to it; emergency exception is reconciled; unrelated verification and versioning meaning remains.

## Case 2 — Export Retention

- Required meaning: all customer-data exports are retained for exactly 30 days, including diagnostic exports.
- Semantic owner: `Retention Policy`.
- Observation point: export procedure.
- Related manifestation: diagnostic exception and archive wording.
- Pass shape: policy owns the complete rule; procedure and exception agree without duplicate ownership; unrelated encryption and access meaning remains.

## Case 3 — Planning Only

- Required meaning: privileged service accounts must rotate credentials every 60 days, including break-glass accounts.
- Semantic owner: `Credential Rotation`.
- Observation point: onboarding checklist.
- Related manifestation: break-glass exception.
- Pass shape: update plan names owner, observation point, manifestations, operation, and preserved meaning without editing the fixture.
