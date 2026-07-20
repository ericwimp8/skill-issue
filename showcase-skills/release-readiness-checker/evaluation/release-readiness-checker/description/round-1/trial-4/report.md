# Release Readiness Report

## Candidate and Scope

- **Candidate:** Service release `4.0.0`, revision `service-400`.
- **Target:** The release candidate represented by the complete contents of `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4/`.
- **Included components:** Manifest metadata, current source schema requirement, destructive schema migration, rollback documentation, and verification log.
- **Authority source:** `RELEASE.md` requires agreement among the manifest, migration, and current source schema versions; a tested recovery procedure for every destructive migration; and current verification evidence.
- **Limitations:** No environment, release platform, release owner, verification command, verification timestamp, or retained test output is identified. This audit used only the supplied candidate files and read-only inspection. No migration or release operation was run.

## Decision

**not ready**

The candidate fails two release-critical project gates. The manifest declares schema `12`, while both the migration and current source require schema `13`. The migration destructively drops `legacy_sessions`, and the supplied rollback document explicitly says that no recovery procedure for this removal has been validated. Either failure is sufficient to prevent readiness under `RELEASE.md`.

## Gate Results

### 1. Schema Version Agreement

- **Requirement and source:** `RELEASE.md` requires the manifest, migration, and current source to agree on schema version.
- **Status:** `failed`
- **Evidence:** `manifest.txt` declares `schema=12`; `migrate.sql` declares target schema `13`; `source.txt` declares `required_schema=13`.
- **Limitations:** None affecting the observed contradiction; all three authoritative candidate values are directly readable.
- **Risk:** Releasing with the manifest at schema `12` can misidentify the database compatibility expected by source and migration behavior.

### 2. Destructive Migration Recovery

- **Requirement and source:** `RELEASE.md` requires every destructive migration to have a tested recovery procedure.
- **Status:** `failed`
- **Evidence:** `migrate.sql` runs `DROP TABLE legacy_sessions;`. `rollback.md` states that no recovery procedure has been validated for the removal of `legacy_sessions`.
- **Limitations:** The migration was deliberately not executed. Execution is unnecessary to establish that the required validated recovery evidence is absent because the candidate states that fact explicitly.
- **Risk:** If the destructive migration causes data loss or service failure, the candidate has no validated path to restore the removed table or its data.

### 3. Current Verification Evidence

- **Requirement and source:** `RELEASE.md` requires current verification evidence.
- **Status:** `passed`
- **Evidence:** `checks.log` records `revision=service-400` and `result=passed`; `manifest.txt` identifies the candidate as revision `service-400`.
- **Limitations:** The log does not identify the verification command, checks performed, environment, timestamp, or retained detailed output. It proves only that the supplied evidence records a passing result for the exact candidate revision.
- **Risk:** The opaque log provides weak diagnostic and reproducibility value, although its revision matches the candidate.

## Unresolved Risks

- **Verification coverage uncertainty:** The passing log does not identify which checks ran or whether they covered the schema mismatch and recovery requirement. Impact: stakeholders may overestimate what the verification result establishes. Needed evidence: the exact verification command, environment, timestamp, check list, exit status, and retained output for revision `service-400`.
- **Recovery feasibility:** No evidence establishes that `legacy_sessions` data can be recreated or restored after the drop. Impact: irreversible data loss may occur. Needed evidence: a recovery procedure and a successful recovery test using representative data in an isolated environment.

## Prioritized Next Actions

1. **Release owner:** Update `manifest.txt` so its schema version agrees with the intended target schema `13`, then obtain a review that confirms `manifest.txt`, `migrate.sql`, and `source.txt` all declare schema `13`. This resolves the schema version agreement gate.
2. **Database or service owner:** Define a recovery procedure for `legacy_sessions` that restores both table structure and required data, then test it after applying the migration in an isolated disposable environment. Retain the commands, inputs, outputs, exit statuses, and recovery verification results. This resolves the destructive migration recovery gate.
3. **Verification owner:** Re-run the project verification checks against the corrected candidate revision and retain an evidence record containing the exact revision, command, environment, timestamp, checks performed, exit status, and detailed output. This strengthens and reconfirms the current verification evidence gate after candidate changes.
4. **Release owner:** Repeat this readiness audit against the corrected, revision-identified candidate. Proceed to release operations only after the schema agreement and recovery gates pass and current verification evidence applies to that same revision.

## Evidence Index

### Inspected Paths

- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/description/round-1/trial-4/request.md`
- `showcase-skills/release-readiness-checker/skill/release-readiness-checker/SKILL.md`
- `supporting-skills/document-update-discipline/SKILL.md`
- `.repository-privacy.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4/RELEASE.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4/checks.log`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4/manifest.txt`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4/migrate.sql`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4/rollback.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4/source.txt`

### Commands

- `find showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/description/trial-4 -type f -print | sort` — run from repository root; exit status `0`; enumerated the candidate files.
- `cat` on each inspected file listed above — run from repository root; exit status `0`; supplied the candidate requirements and evidence used in this report.

### Deliberately Not Run

- The migration was not executed.
- No release, publish, deploy, tag, upload, promotion, signing, submission, or shared-state mutation was performed.
- No unavailable verification command was invented or run.
