# CI Failure Triage Report

## Run Context

- The retained run is labeled `integration`, with a `wait-for-api` boundary logged at `2026-07-21T04:00:00Z` (`run.log:1`).
- The supplied workflow defines one `integration` job, an `api` service using `fixture/api:latest`, then `./wait-for-api.sh` followed by `./integration-tests.sh` (`workflow.yml:1-9`).
- The event, revision, runner, platform, matrix, concrete environment, and actual workflow revision are unavailable. The supplied workflow may differ from the failed run (`evidence-note.md:1`).

## Evidence Inventory

- `run.log` is retained but incomplete: it contains the wait-step label, one connection error after 30 seconds, and a notice that all remaining step output expired (`run.log:1-3`).
- `workflow.yml` provides a possible job shape, service image tag, and step order (`workflow.yml:1-9`), but its provenance is uncertain because it may not be the run's workflow revision (`evidence-note.md:1`).
- `evidence-note.md` records that the service-container log, image digest, workflow revision, and retry history were not retained (`evidence-note.md:1`).
- No service status, container exit details, port publication, runner metadata, script contents, integration-test output, or complete job summary is available.

## Failure Sequence

1. The available log marks the start of `integration / wait-for-api` at `04:00:00Z` (`run.log:1`).
2. Thirty seconds later, `curl` received `connection refused` from `127.0.0.1:8080` (`run.log:2`). This establishes that no process accepted that connection at that observation time; it does not establish why.
3. Later output is unavailable because retention expired (`run.log:3`). Whether `./integration-tests.sh` ran, skipped, or failed independently cannot be confirmed.
4. If the supplied workflow matches the executed revision, the wait script preceded the integration-test script and the job declared an `api` service (`workflow.yml:3-9`). That condition is unverified.

## Primary Diagnosis

**Unresolved.** The failed invariant is that the API expected by the wait boundary must accept a connection at `127.0.0.1:8080` before integration tests proceed. The direct observation is a refused connection after 30 seconds (`run.log:1-2`). The responsible owner cannot be assigned from retained evidence because the same symptom is consistent with several untested causes: service startup failure, startup exceeding the wait budget, incorrect host or port, missing service-port publication, or wait-script behavior.

Confidence is high that the readiness boundary failed at the recorded instant and low for every root-cause candidate. The strongest configuration hypothesis, conditional on `workflow.yml` being the executed definition and on runner networking requiring explicit host publication, is that the declared service has no visible port mapping (`workflow.yml:4-8`). Missing runner identity, the actual workflow revision, and service evidence prevent promoting that hypothesis to a primary diagnosis.

## Cascade Classification

- **Primary:** none established.
- **Contributing:** none established.
- **Independent:** none established; later output is unavailable (`run.log:3`).
- **Cascading:** any skipped or aborted integration-test work may have followed the readiness failure, but its status is unconfirmed because no later step output or job summary remains.
- **Noise:** none established. The connection error is material evidence of the failed readiness boundary (`run.log:2`).

## Remediation Direction

Do not change application, workflow, retry, or cache behavior on the retained evidence. First identify the owner of the unavailable listener. If exact-run evidence shows that the service was healthy but unreachable from the runner, correct the service networking or the wait target while preserving the service image and test behavior. If the service failed or remained unready, repair that startup invariant or align the bounded readiness budget with observed startup behavior. Avoid blanket retries because retry history and failure stability are unknown (`evidence-note.md:1`).

## Verification Plan

1. Recover the failed run's immutable revision, exact expanded workflow, runner type, job summary, step conclusions, and full timestamps. Confirm whether the executed workflow contains the service and step order shown in `workflow.yml:3-9`. A mismatch disconfirms configuration conclusions drawn from the supplied file.
2. Recover the service image digest and service-container creation, health, stdout, stderr, exit, and port-publication records for that same job. A running healthy service with the expected port published shifts ownership toward runner networking or the wait target; a stopped, unhealthy, or non-listening service shifts ownership to service startup.
3. At the exact revision, inspect the concrete `./wait-for-api.sh` invocation and the service's configured listen address and port. Confirm the script targets `127.0.0.1:8080`, records attempts and elapsed time, and uses the runner-appropriate service address. A different configured endpoint or script target disconfirms the current endpoint assumption.
4. Reproduce only the readiness boundary in the same runner topology with the immutable service digest. Capture container status and port state before running the exact wait command. Confirmation is a refused connection paired with evidence that the expected listener is absent or unreachable; successful readiness under the same topology disconfirms a deterministic configuration or image failure and raises timing or nondeterminism as a hypothesis.
5. After the owner-specific correction, run the readiness boundary first, then `./integration-tests.sh`, then the complete `integration` job under the same runner conditions. The authoritative confirmation is a fresh run of the affected job at the corrected revision with retained service logs, digest, workflow provenance, and step conclusions. These checks are planned and were not executed.

## Uncertainties

- The supplied workflow may be stale, so its service declaration and missing visible port mapping cannot be treated as executed-run facts (`evidence-note.md:1`; `workflow.yml:4-8`).
- The absent service log and image digest prevent distinguishing service failure from runner-to-service connectivity (`evidence-note.md:1`).
- The absent retry history prevents assessing reproducibility or timing sensitivity (`evidence-note.md:1`).
- The expired remainder of the log prevents classifying downstream tests or other failures (`run.log:3`).
- The smallest resolving evidence is the exact workflow revision plus runner metadata, service digest/status/logs/port publication, wait-script contents, and complete step conclusions from the same run.

## Authorization Boundary

No remote CI action was requested or triggered. No workflow, source, dependency, cache, secret, branch, artifact, or service state was changed. Recovering remote run evidence or dispatching an authoritative confirmation run requires separate explicit authorization.
