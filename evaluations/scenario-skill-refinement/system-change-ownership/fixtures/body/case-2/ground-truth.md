# Ground Truth

- Required outcome: record expiry is derived consistently from each tenant's retention policy.
- Observation point: the cron contains the current fixed cutoff.
- Current owners: `PolicyRepository` owns stored tenant settings, `RecordRepository` owns deletion effects, and the cron owns scheduling only.
- Smallest complete placement: introduce or strengthen a retention-policy/domain owner that converts a tenant policy and current time into the cutoff, then have the cron orchestrate policy lookup, cutoff evaluation, and repository deletion.
- Reconciliation: the settings surface continues through the policy repository; any future expiry preview should reuse the same retention evaluator; the cron does not own policy interpretation.
- Verification: distinct tenant policies, boundary timestamps, repeated cron execution, and absence of a fixed or independently interpreted retention rule in the cron.
