# Architecture Placement Decision

## Required outcome

Every archive export must apply the deployment-wide `ExportPolicy.encryptionRequired` value, regardless of whether the export is initiated through the CLI or HTTP.

## Observation point

The gap was observed in `ExportCommand`, but the command only delegates to `ExportService.export`. It neither owns deployment policy nor writes archives, so a CLI-only `--encrypt` default would fix one entry point while leaving HTTP exports unchanged and would create a second, channel-specific policy source.

## Current flow and owners

- `AppConfig` owns the deployment's configured `ExportPolicy` value.
- `ExportCommand.run` and `ExportRoute.post` are adapters that both delegate unchanged data to `ExportService.export`.
- `ExportService` is the shared export workflow owner, but it currently hard-codes `false` when calling `ArchiveWriter.write`.
- `ArchiveWriter` owns the concrete archive-writing effect and already accepts the resolved encryption decision; it does not own choosing deployment policy.

The current flows are `CLI -> ExportService -> ArchiveWriter` and `HTTP -> ExportService -> ArchiveWriter`. The policy is defined in configuration but is not connected to the shared export workflow.

## Selected owner versus nearby hooks

`ExportService` must own enforcement of the configured export policy because it is the common workflow through which both connected export entry points pass. `AppConfig` remains the owner of the policy value, while `ExportService` consumes that value and ensures every write uses it.

Neither `ExportCommand` nor `ExportRoute` should resolve or override encryption: they are nearby invocation hooks and placing policy there would duplicate enforcement across adapters. `ArchiveWriter` should continue receiving the resolved boolean rather than acquiring configuration responsibility, because its role is performing the write selected by the export workflow.

## Dependent changes

- Supply the deployment `ExportPolicy` (or its `encryptionRequired` value) to `ExportService` when the service is composed from `AppConfig`.
- Replace the hard-coded `false` in `ExportService.export` with the configured `encryptionRequired` value passed to `ArchiveWriter.write`.
- Update every `ExportService` construction site to provide the configured policy.
- Leave `ExportCommand` and `ExportRoute` as policy-free delegates so both automatically receive the shared behavior.
- Keep the `ArchiveWriter` interface unless a concrete writer cannot fulfill its existing `encrypted` contract; verify each implementation honors the supplied value.

## Smallest complete placement

Strengthen the existing shared `ExportService` rather than adding a CLI flag or parallel adapter-level enforcement. Inject the deployment policy into the service and use it for the service's single archive-write handoff. This connects the existing configuration owner to the existing workflow owner and preserves one enforcement point for all connected export paths.

## Ownership-level verification

- With `encryptionRequired: true`, verify `ExportService.export` calls the writer with `encrypted: true`; with `false`, verify it calls the writer with `false`.
- Verify both `ExportCommand.run` and `ExportRoute.post` reach that same policy-aware service path and cannot supply independent encryption defaults.
- Verify all service composition sites source the value from the deployment `AppConfig`, with no remaining hard-coded or adapter-local encryption decision.
- Verify concrete `ArchiveWriter` implementations produce encrypted archives when passed `true` and unencrypted archives only when policy permits `false`.
- Search all archive-export entry points to confirm none bypass `ExportService`; any bypass must be routed through the same enforcement owner or apply the same policy at the actual shared export owner.
