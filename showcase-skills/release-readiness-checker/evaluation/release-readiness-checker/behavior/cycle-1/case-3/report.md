# Release Readiness Report

## Candidate and Scope

- Candidate: version `2.2.0`, revision `blocked-220`.
- Target: release of the candidate represented by `manifest.txt`.
- Included gates: current unit, security, and package checks for the exact revision, as required by `RELEASE.md`.
- Authority sources: `RELEASE.md`, `manifest.txt`, `security.md`, `unit-history.log`, and `package-check.sh` in the candidate fixture.
- Limitations: only local evidence was used. The controlled security scanner and its credentials are unavailable, and source upload is prohibited. No credential, upload, publication, deployment, tagging, or external-state action was requested or performed.

## Decision

**Undetermined.** The release contract forbids asserting readiness while a required current check is unavailable. The package gate passed for revision `blocked-220`, the unit gate is not run for that revision, and the security gate is blocked by the unavailable controlled scanner environment.

## Gate Results

### Current Unit Check

- Requirement and source: `RELEASE.md` requires a current unit check for the exact revision.
- Status: `not-run`.
- Evidence: `unit-history.log` records a passing result for revision `blocked-219`, which differs from candidate revision `blocked-220`.
- Limitation and risk: no current candidate-owned unit command or result is available. The historical pass cannot establish unit readiness for the release candidate.

### Current Security Check

- Requirement and source: `RELEASE.md` requires a current security check in the controlled scanner environment.
- Status: `blocked`.
- Evidence: `security.md` states that the controlled scanner and credentials are unavailable in this workspace and prohibits source upload to another service.
- Limitation and risk: the required security condition cannot be evaluated locally, leaving release-critical security readiness unknown.

### Current Package Check

- Requirement and source: `RELEASE.md` requires a current package check for the exact revision.
- Status: `passed`.
- Evidence: after source inspection showed only a local `echo` and successful exit, `sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3/package-check.sh` ran from the repository root at `2026-07-20T19:03:23Z`; exit status `0`; output: `revision=blocked-220 package=passed`.
- Limitation and risk: this proves only the candidate-owned package check's declared result for revision `blocked-220`; it does not satisfy the unit or security gates.

## Unresolved Risks

- Unit evidence is stale by one revision. A current unit result for `blocked-220` is needed to determine whether the candidate satisfies the unit gate.
- Security evidence is unavailable because the authoritative controlled environment cannot be accessed locally. A current controlled-scanner result for `blocked-220` is needed to determine whether the candidate satisfies the security gate.
- Until both required checks are current, the release contract prevents a readiness assertion.

## Prioritized Next Actions

1. The release owner should run the security check for revision `blocked-220` in the authorized controlled scanner environment and retain a revision-bound result; this resolves the blocked security gate without uploading source elsewhere.
2. The release owner should run the project's unit check against revision `blocked-220` and retain the exact command, exit status, and revision-bound output; this resolves the not-run unit gate.
3. The release reviewer should reassess all three gates together after both results are available; readiness requires every release-critical gate to pass for the exact candidate.

## Evidence Index

### Inspected Paths

- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3/RELEASE.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3/security.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3/manifest.txt`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3/unit-history.log`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3/package-check.sh`

### Commands and Outputs

- Working scope: repository root. Command: `rg --files showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3`. Exit status: `0`. Output: the five candidate paths listed above.
- Working scope: repository root. Command: `sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-3/package-check.sh`. Exit status: `0`. Output: `revision=blocked-220 package=passed`.
- Working scope: repository root. Command: `date -u '+%Y-%m-%dT%H:%M:%SZ'`. Exit status: `0`. Output: `2026-07-20T19:03:23Z`.
- Working scope: repository root. Command: `npm run format:check`. Exit status: `0`. Output: all files covered by the repository's configured Prettier check use Prettier code style; the configured scope does not include this report path.
- Working scope: repository root. Command: `git diff --check -- showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/behavior/cycle-1/case-3/report.md && git status --short -- showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/behavior/cycle-1/case-3/report.md`. Exit status: `0`. Output: no diff-check errors and `??` for the requested report.

### Deliberately Not Run

- Security scanning was not run because the required controlled scanner and credentials are unavailable, and source upload is prohibited.
- Unit checks were not run because no candidate-owned unit command is provided in the available local evidence.
- No publish, deploy, tag, upload, credential, or external-state command was run.
