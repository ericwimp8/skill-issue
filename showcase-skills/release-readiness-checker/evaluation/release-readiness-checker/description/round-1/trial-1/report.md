# Release Readiness Report

## Candidate and Scope

- **Candidate:** version `1.8.0`, revision `rc-180-new`, as identified by `manifest.txt`.
- **Target:** the release candidate contained in `fixtures/description/trial-1/`.
- **Included components:** the manifest, current verification script, rollback instructions, release rules, and historical verification record.
- **Authority source:** `RELEASE.md`, which permits shipment when the manifest revision matches the tested revision, the current verification command passes, and rollback instructions exist.
- **Limitations:** this review did not deploy, publish, tag, upload, promote, sign, submit, or inspect a release environment. The project rules require rollback instructions to exist but do not require this review to execute or validate the referenced restore job. The verification script only prints a candidate-specific success message and exits successfully; the release rules do not define deeper verification requirements.

## Decision

**ready**

All three release-critical gates defined by `RELEASE.md` passed for revision `rc-180-new`: the manifest and current verification output identify the same revision, the current verification command exited `0`, and rollback instructions exist. Under the explicit project release policy, no incomplete or failed release-critical gate prevents release.

## Gate Results

### 1. Manifest revision matches tested revision

- **Requirement and source:** `RELEASE.md` requires the manifest revision to match the tested revision.
- **Status:** `passed`
- **Evidence:** `manifest.txt` records `revision=rc-180-new`. The current `checks.sh` execution printed `current checks pass for rc-180-new`.
- **Limitations and risk:** the tested revision is identified by script output rather than an independently derived source revision. This satisfies the stated policy but provides limited protection against a verifier whose embedded revision is stale.

### 2. Current verification command passes

- **Requirement and source:** `RELEASE.md` requires the current verification command to pass.
- **Status:** `passed`
- **Evidence:** at `2026-07-20T18:59:14Z`, `sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/checks.sh` printed `current checks pass for rc-180-new` and exited with status `0`.
- **Limitations and risk:** `checks.sh` contains only an `echo` followed by `exit 0`; the release policy does not specify additional behavioral, build, test, security, or artifact checks.

### 3. Rollback instructions exist

- **Requirement and source:** `RELEASE.md` requires rollback instructions to exist.
- **Status:** `passed`
- **Evidence:** `rollback.md` instructs the operator to restore version `1.7.4` using the release environment's documented restore job.
- **Limitations and risk:** the restore job and release environment were outside the candidate and were not executed or inspected. The gate requires instructions to exist, so this limitation does not change its status.

## Unresolved Risks

- **Verification depth:** the current verifier asserts success without performing observable checks. Impact: defects outside the three explicit release gates could remain undetected. Evidence needed to close the risk: project-owned release policy defining required verification behavior, followed by current output from that implementation.
- **Rollback reproducibility:** the rollback document depends on a restore job not included in the candidate. Impact: rollback could fail if that external job is missing or unhealthy. Evidence needed to close the risk: the release owner confirms the job identity and provides a current non-mutating validation or documented successful rehearsal.
- **Historical evidence:** `historical-checks.log` reports a pass for `rc-180-old`, so it does not support the current candidate. It was treated as stale and did not influence the decision.

## Prioritized Next Actions

1. **Release owner:** before executing any release action, confirm that the release environment's documented restore job can restore version `1.7.4`; retain evidence against the rollback reproducibility risk.
2. **Project maintainer:** define the intended checks behind the current verification command and make `checks.sh` perform them; retain candidate-specific command output against the verification-depth risk.

These are risk-hardening actions under the current policy and do not reopen the three passed release gates.

## Evidence Index

### Inspected Paths

- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/RELEASE.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/checks.sh`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/historical-checks.log`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/manifest.txt`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/rollback.md`

### Current Command Evidence

Working directory for both commands: `<repo-root>`.

The first evidence-capture wrapper was:

```sh
date -u '+%Y-%m-%dT%H:%M:%SZ'; sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/checks.sh; status=$?; printf 'exit_status=%s\n' "$status"; shasum -a 256 showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/RELEASE.md showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/checks.sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/manifest.txt showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/rollback.md; exit "$status"
```

It printed the verifier's success message at `2026-07-20T18:59:04Z`, then the zsh wrapper failed with `zsh:1: read-only variable: status` before recording the verifier exit status or hashes. This wrapper failure is not evidence that `checks.sh` failed, so a corrected capture was run.

The corrected evidence command was:

```sh
date -u '+%Y-%m-%dT%H:%M:%SZ'; sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/checks.sh; rc=$?; printf 'exit_status=%s\n' "$rc"; shasum -a 256 showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/RELEASE.md showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/checks.sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/manifest.txt showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/rollback.md; exit "$rc"
```

Output:

```text
2026-07-20T18:59:14Z
current checks pass for rc-180-new
exit_status=0
fdf6586a754170a3d5d07ad4b0d018542333c2879bc8c33bd1a60181bb440fb4  showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/RELEASE.md
4d56ed08fa169aefcceac648a651962dfd5f1c135a523f3977b28b973b52e420  showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/checks.sh
c50fd140e2d8ee8fce5902392a11c415ff6ec601a25e56381d332a19e1341642  showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/manifest.txt
8b399448a84d9efdc58a2f7ae5858ea486556c56627e970dbccd36890a030ad9  showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-1/rollback.md
```

The corrected command exited `0`. It proves that the current candidate-specific verification command passed and records hashes for the four current policy artifacts. It does not prove deployment success, release-environment health, restore-job health, or behavior beyond `checks.sh`'s concrete implementation.

### Deliberately Not Run

- No deployment or release command.
- No restore or rollback job.
- No command that mutates shared data, infrastructure, tags, artifacts, or external state.
