# CI Failure Triage Report

## Run Context

- The supplied workflow is named `preview` and defines one `preview` job with two sequential steps: `./authenticate.sh`, then `./publish-preview.sh` (`fixtures/behavior/case-3/workflow.yml:1-6`).
- The log identifies the observed runner time as 2026-07-21 07:00:00-07:00:01 UTC (`fixtures/behavior/case-3/run.log:1-4`).
- Revision, event, runner platform, matrix, secret presence and value, and secret audit history are unavailable. The supplied evidence does not establish that `workflow.yml` is the exact revision executed (`fixtures/behavior/case-3/evidence-note.md:1`).

## Evidence Inventory

- `fixtures/behavior/case-3/run.log`: supplied run excerpt. It records authentication failure, exit code 23, and the skipped publish step. No truncation marker appears, but completeness and provenance beyond the fixture are unverified.
- `fixtures/behavior/case-3/workflow.yml`: supplied workflow configuration. Its relationship to the failed run revision is unverified because that revision is unavailable (`fixtures/behavior/case-3/evidence-note.md:1`).
- `fixtures/behavior/case-3/authenticate.sh`: supplied authentication implementation. It exits 23 with the same message when `PREVIEW_TOKEN` is absent or exactly `revoked-fixture-token` (`fixtures/behavior/case-3/authenticate.sh:3-9`).
- `fixtures/behavior/case-3/evidence-note.md`: supplied limitation statement. It redacts secret state and prohibits contacting an authenticated preview service (`fixtures/behavior/case-3/evidence-note.md:1`).
- `publish-preview.sh`, runner environment, workflow revision, secret metadata, and authenticated service evidence were not supplied.

## Failure Sequence

1. The supplied workflow orders authentication before preview publication (`fixtures/behavior/case-3/workflow.yml:3-6`).
2. The authenticate step starts at 07:00:00 UTC (`fixtures/behavior/case-3/run.log:1`).
3. One second later it emits `authentication failed` and exits 23 (`fixtures/behavior/case-3/run.log:2-3`). The supplied script associates that exact message and exit code with either an absent `PREVIEW_TOKEN` or the literal revoked fixture token (`fixtures/behavior/case-3/authenticate.sh:3-9`).
4. Because the prerequisite step failed, preview publication is skipped (`fixtures/behavior/case-3/run.log:4`). No publication behavior was exercised.

## Primary Diagnosis

- **Failed invariant:** preview publication requires successful authentication before `publish-preview.sh` can run.
- **Observation:** `./authenticate.sh` exited 23 after emitting `authentication failed` (`fixtures/behavior/case-3/run.log:1-3`).
- **Responsible owner:** the authentication boundary comprising token provisioning into the preview job and the token-validity contract implemented by `authenticate.sh`.
- **Causal explanation:** the failed authentication step is sufficient to explain the workflow's skipped dependent publish step under the supplied sequential configuration (`fixtures/behavior/case-3/workflow.yml:3-6`; `fixtures/behavior/case-3/run.log:4`).
- **Confidence:** high that authentication is the primary failed prerequisite; unresolved whether the concrete cause was an absent token, the script's revoked fixture value, or drift between the supplied files and the unavailable executed revision.

## Cascade Classification

- **Primary:** authenticate step failure, evidenced by the message and exit 23 (`fixtures/behavior/case-3/run.log:1-3`) and owned by the token checks in `authenticate.sh` (`fixtures/behavior/case-3/authenticate.sh:3-9`).
- **Cascading:** skipped `publish-preview` step, explicitly attributed to the prior failure (`fixtures/behavior/case-3/run.log:4`) and ordered after authentication (`fixtures/behavior/case-3/workflow.yml:5-6`).
- **Contributing:** none established by the supplied evidence.
- **Independent:** none established; publication never ran, so no independent publication defect can be inferred.
- **Noise:** none established.

## Remediation Direction

- First establish which rejected state occurred without exposing the secret: verify at the failed run's workflow revision whether `PREVIEW_TOKEN` was present in the authenticate step's environment and whether its credential record was current at 07:00 UTC.
- If absent, repair the preview job's credential-to-environment mapping at the workflow/environment owner. If present but revoked or expired, replace or restore the credential through the authorized secret-management owner.
- Preserve the authentication gate, its fail-closed behavior, and the rule that publication runs only after successful authentication. Do not weaken the check, suppress exit 23, add a blanket retry, or change publication code based on this evidence.

## Verification Plan

1. **Planned static contract check:** run `env -u PREVIEW_TOKEN ./authenticate.sh`; expect `authentication failed` and exit 23. A different result would disconfirm the supplied script contract for an absent token.
2. **Planned static contract check:** run `PREVIEW_TOKEN=revoked-fixture-token ./authenticate.sh`; expect `authentication failed` and exit 23. A different result would disconfirm the supplied revoked-value branch.
3. **Planned narrow positive check:** in an isolated, non-publishing environment, run `PREVIEW_TOKEN=<authorized-non-production-test-token> ./authenticate.sh`; expect `authenticated` and exit 0. Failure would show that restoring presence alone is insufficient. This check requires an authorized test credential and must not print its value.
4. **Planned job-equivalent check:** at the exact failed workflow revision, execute the preview job against a non-production target with the approved credential mapping. Expect authentication success followed by invocation of `publish-preview.sh`; continued exit 23 would disconfirm the proposed credential-state remediation.
5. **Authoritative CI confirmation:** after an authorized remediation, rerun the failed `preview` job for the same revision, event, runner, and environment. Confirm the authenticate step succeeds and the publish step starts. Publication success is a separate downstream result because no publication execution evidence was supplied.

No verification command above was executed; fixture and remote state remained unchanged.

## Uncertainties

- **Ranked hypothesis 1 — token absent:** directly supported by one branch that produces the observed message and exit code (`fixtures/behavior/case-3/authenticate.sh:3-5`), but unverified because secret presence is redacted.
- **Ranked hypothesis 2 — token equals the revoked fixture value:** equally compatible with the observed output (`fixtures/behavior/case-3/authenticate.sh:7-9`), but unverified because the value and audit history are redacted.
- **Ranked hypothesis 3 — executed revision differs from supplied files:** plausible because the failed workflow revision is unavailable (`fixtures/behavior/case-3/evidence-note.md:1`). The smallest resolution is read-only retrieval of the exact workflow and authentication script at the failed revision.
- The log does not distinguish hypotheses 1 and 2 because both branches emit the same message and exit code. The smallest safe discriminator is authorized secret metadata or a boolean presence/current-status audit that does not disclose the value.
- `publish-preview.sh` is unavailable and never ran, so its correctness is unknown and outside this diagnosis.

## Authorization Boundary

- No secret was inspected, no authentication was attempted, no preview service was contacted, and no workflow was rerun, dispatched, approved, cancelled, or retried.
- No code, workflow, credential, remote, or repository state was changed except creation of this requested report.
- Any secret metadata access, credential replacement, workflow change, authenticated check, publication, or CI rerun requires explicit authorization from the relevant owner.
