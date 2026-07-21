# Architecture Placement Decision

## Required outcome

Each tenant's records must be deleted using that tenant's configured `retentionDays`, with one authoritative component responsible for turning the stored policy into a deletion cutoff.

## Observation point

The gap is visible in `CleanupCron.run`, where every tenant currently receives the same hard-coded 30-day cutoff. That loop is the consumer of the retention decision, not its owner.

## Current flow and owners

- `TenantPolicy` and `PolicyRepository` own the tenant-scoped retention setting and its retrieval.
- `RetentionSettingsController` exposes the same policy for administration; it does not define or apply retention semantics.
- `CleanupCron` owns scheduled cleanup orchestration: enumerate tenants and request deletion.
- `RecordRepository` owns record enumeration and deletion persistence; it does not decide how long records should be retained.

The current cleanup flow is `CleanupCron -> RecordRepository.tenantsWithRecords() -> hard-coded cutoff -> RecordRepository.deleteCreatedBefore(...)`. The administrative flow is `RetentionSettingsController -> PolicyRepository.get(...)`.

## Selected owner versus nearby hooks

Retention interpretation belongs in the policy layer, next to `TenantPolicy` and `PolicyRepository`. Add a policy-layer retention service that depends on `PolicyRepository` and returns the cutoff for a tenant and reference time, for example `cutoffFor(tenantId, now)`.

Do not make `CleanupCron` read `retentionDays` and perform the date arithmetic itself. That would make a scheduling adapter a second owner of policy semantics. `RetentionSettingsController` is only an administrative hook, and `RecordRepository` is only a persistence hook; placing cutoff logic in either would likewise mix responsibilities.

## Dependent changes

- Inject the policy-layer retention service into `CleanupCron`.
- In the tenant loop, ask that service for the tenant-specific cutoff, then pass the result to `RecordRepository.deleteCreatedBefore`.
- Keep `RetentionSettingsController.load` backed by `PolicyRepository`, so administration and cleanup consume the same authoritative tenant policy.
- Update the composition root to construct the retention service with `PolicyRepository` and provide it to the cron.
- Leave the `RecordRepository` contract unchanged; it should continue accepting an already-decided cutoff.

## Smallest complete placement

Create one retention-policy owner in the policy module, move the 30-day cutoff calculation out of `CleanupCron`, and route cleanup through that owner. No parallel setting lookup or cutoff calculation should remain in the cron, controller, or record repository.

## Ownership-level verification

Verify two tenants with different `retentionDays` produce different cutoffs from the policy service for the same `now`, and that `CleanupCron` passes those exact cutoffs to deletion. Also verify the cron has no direct dependency on `PolicyRepository` and does not access `retentionDays` or calculate days-to-milliseconds; the controller and cleanup must both derive their behavior from the same repository-backed `TenantPolicy`, while `RecordRepository` remains policy-agnostic.
