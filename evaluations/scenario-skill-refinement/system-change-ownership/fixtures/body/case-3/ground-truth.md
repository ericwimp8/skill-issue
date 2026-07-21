# Ground Truth

- Required outcome: every export entry point applies the deployment-owned encryption policy.
- Observation point: the CLI exposes an unencrypted export.
- Current owners: `AppConfig` owns the deployment policy, `ExportService` owns export behavior shared by CLI and HTTP, and `ArchiveWriter` owns the concrete write effect.
- Smallest complete placement: inject `ExportPolicy` into `ExportService` and pass the policy decision to the writer for every export; keep entry points policy-free.
- Reconciliation: composition constructs the service from `AppConfig`; CLI and HTTP continue through the same service; avoid caller flags that can contradict a required deployment policy.
- Verification: required and optional policy configurations through both entry points, writer arguments, and absence of independent encryption defaults in callers.
