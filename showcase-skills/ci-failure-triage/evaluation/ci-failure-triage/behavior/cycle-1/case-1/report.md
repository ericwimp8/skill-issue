# CI Failure Triage Report

## Run Context

- Workflow: `tool-check`; supplied configuration defines one `verify` job with three sequential steps (`workflow.yml:1-7`).
- Observed step order: tool validation, build, then archive (`workflow.yml:5-7`).
- The validation step supplies `TOOL_VERSION=v2.4.0` to `./scripts/check-tool.sh` (`workflow.yml:5`).
- The run log covers `2026-07-21T05:00:00Z` through `2026-07-21T05:00:01Z` (`run.log:1-6`).
- Revision, event, trigger, runner, matrix, operating system, dependency state, cache state, and environment details beyond `TOOL_VERSION` are absent from the fixture.

## Evidence Inventory

- `workflow.yml`: supplied fixture configuration for the failing workflow; no revision or generation timestamp is recorded.
- `check-tool.sh`: supplied fixture implementation of the invoked validation boundary; no revision or generation timestamp is recorded.
- `run.log`: supplied fixture log dated `2026-07-21`; it contains the complete visible sequence for the single `verify` job, with no explicit truncation marker.
- Local reproduction executed against the fixture script: `TOOL_VERSION=v2.4.0 sh showcase-skills/ci-failure-triage/evaluation/ci-failure-triage/fixtures/behavior/case-1/check-tool.sh` printed `unsupported tool major: v2` and exited `42`.
- Local counterexample executed against the same script: `TOOL_VERSION=2.4.0 sh showcase-skills/ci-failure-triage/evaluation/ci-failure-triage/fixtures/behavior/case-1/check-tool.sh` printed `tool accepted` and exited `0`.
- Build and archive scripts, remote CI metadata, revision contents, and artifact-upload configuration were not supplied, so their independent behavior cannot be inspected.

## Failure Sequence

1. The `verify` job starts its first step with `TOOL_VERSION=v2.4.0` (`workflow.yml:3-5`; `run.log:1`).
2. `check-tool.sh` extracts the substring before the first dot, producing `v2`, then requires that value to equal the literal `2` (`check-tool.sh:3-4`).
3. The comparison fails, the script reports `unsupported tool major: v2`, and exits `42` (`check-tool.sh:5-6`; `run.log:2-3`). The local reproduction produced the same message and status.
4. Because the first sequential step failed, the build and archive steps do not run (`workflow.yml:6-7`; `run.log:4-5`).
5. With archive production skipped, the expected `dist/app.tar.gz` artifact is unavailable for upload, producing the final warning (`run.log:6`).

## Primary Diagnosis

- **Failed invariant:** the configured major-2 tool version must be recognized by the validation step before build work can proceed.
- **Observation:** the workflow supplies `v2.4.0`, while the validator compares the unnormalized first component `v2` with `2` (`workflow.yml:5`; `check-tool.sh:3-4`).
- **Responsible owner:** the version parsing and acceptance behavior in `check-tool.sh` at the workflow-script boundary.
- **Confidence:** high. The source paths establish the mismatch, the run log records its exact result, and local execution reproduces exit `42`; changing only the input to `2.4.0` reaches the accepted path.
- **Causal explanation:** the parser treats a conventional leading `v` as part of the major component. That formatting mismatch fails the first sequential step and sufficiently explains every later skip and the absent-artifact warning.

## Cascade Classification

- **Primary:** `check-tool` exit `42` caused by parsing `v2.4.0` as major `v2` (`workflow.yml:5`; `check-tool.sh:3-6`; `run.log:1-3`).
- **Cascading:** the build step is skipped after the validator failure (`workflow.yml:6`; `run.log:4`).
- **Cascading:** the archive step is skipped after the same upstream failure (`workflow.yml:7`; `run.log:5`).
- **Cascading:** the missing `dist/app.tar.gz` upload follows from the archive never running (`run.log:5-6`).
- **Contributing:** none established by the supplied evidence.
- **Independent:** none established; the skipped build and archive paths were not executed and therefore reveal no independent defect.
- **Noise:** none established. The artifact warning is consequential diagnostic output, although it is not a separate cause.

## Remediation Direction

- At the validator's version-parsing boundary, normalize one supported leading `v` before extracting and comparing the major number, so the configured `v2.4.0` is evaluated as major `2`.
- Preserve rejection of versions whose normalized major is not `2`, preserve exit `42` for unsupported majors, and leave build, archive, and artifact behavior unchanged.
- If an external authoritative contract requires numeric-only `TOOL_VERSION` values, the alternative smallest change is to supply `2.4.0` in `workflow.yml`; that contract is not present in the fixture, so this alternative requires confirmation before implementation.

## Verification Plan

1. **Responsible unit:** after the parser change, run `TOOL_VERSION=v2.4.0 sh showcase-skills/ci-failure-triage/evaluation/ci-failure-triage/fixtures/behavior/case-1/check-tool.sh`. Expect `tool accepted` and status `0`; `unsupported tool major: v2` or any nonzero status would disconfirm the fix. Planned, because no implementation was authorized.
2. **Preserved rejection:** run the same script with `TOOL_VERSION=v3.0.0`. Expect `unsupported tool major: 3` and status `42`; acceptance would show that unsupported-major behavior regressed. Planned.
3. **Preserved numeric form:** run the same script with `TOOL_VERSION=2.4.0`. Expect `tool accepted` and status `0`. This input was executed before any change and passed; post-change verification remains planned.
4. **Job-equivalent path:** run the `verify` job's three commands in workflow order in a workspace containing the referenced scripts. Expect validation, build, and archive to complete and `dist/app.tar.gz` to exist. The fixture lacks the build and archive scripts, so this is planned and unavailable locally.
5. **Authoritative confirmation:** rerun the `verify` job for the same revision and environment. Expect all three steps to run successfully and the artifact upload warning to disappear. This remote rerun was not authorized or executed.

## Uncertainties

- The fixture does not define the authoritative `TOOL_VERSION` format. The configured `v2.4.0` makes optional-leading-`v` normalization the best-supported remediation, but a numeric-only external contract would instead place the correction in workflow input. The smallest resolution action is to inspect the version producer or documented environment contract.
- The absent build and archive scripts prevent verification that those steps are otherwise healthy. This does not weaken the established cause of their skips; it limits claims about what happens after validation succeeds.
- The artifact uploader configuration is absent. The log supports a cascade from the skipped archive to the missing artifact, but the exact upload step and its conditions cannot be inspected.
- Revision and runner metadata are absent, so the diagnosis cannot be tied to a commit or platform-specific context.

## Authorization Boundary

- Investigation remained local and read-only except for creation of this report.
- No workflow rerun, retry, dispatch, cancellation, approval, secret access, branch update, push, tag, artifact publication, fixture edit, or other remote mutation was performed.
- Implementation and authoritative CI confirmation require separate explicit authorization.
