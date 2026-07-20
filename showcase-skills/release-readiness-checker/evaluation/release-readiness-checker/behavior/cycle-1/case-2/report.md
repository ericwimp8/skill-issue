# Release Readiness Report

## Candidate and Scope

- Candidate: revision `broken-900`, manifest version `9.0.0`, manifest schema `20`.
- Target: the release candidate contained in `fixtures/behavior/case-2/`.
- Included components: release contract, manifest, application configuration, schema migration, recovery documentation, and unit verification script.
- Authority source: `RELEASE.md` in the candidate fixture.
- Limitations: this review used only repository-local candidate evidence and safe local verification. The migration was inspected but not executed. No recovery command exists to test. No deployment, publication, migration, or external-state operation was performed.

## Decision

**not ready**

The release contract states that a failed critical gate means the candidate is not ready. The candidate fails the identity and schema alignment gate, and its destructive migration has no designed or tested recovery command. Current unit verification passes, but that passing gate cannot override the two failed critical gates.

## Gate Results

### Candidate, Application, and Migration Alignment

- Requirement and source: `RELEASE.md` requires candidate version, application schema, and migration target to match.
- Status: `failed`.
- Evidence: `manifest.txt` declares version `9.0.0` and schema `20`; `application.conf` declares application version `8.9.2` and required schema `21`; `migrate.sql` targets schema `21`.
- Limitation and risk: the manifest disagrees with the application version and with the application/migration schema target. Releasing this identity could associate the wrong application version and schema with revision `broken-900`.

### Destructive Migration Recovery

- Requirement and source: `RELEASE.md` requires destructive migrations to have a tested recovery command.
- Status: `failed`.
- Evidence: `migrate.sql` drops `retained_exports`. `rollback.md` states that recovery for `retained_exports` has not been designed or tested and provides no recovery command.
- Limitation and risk: the migration was deliberately not executed. If applied, retained export data could be irrecoverably lost through the documented release process.

### Current Unit Verification

- Requirement and source: `RELEASE.md` requires current unit verification to pass.
- Status: `passed`.
- Evidence: running `verify.sh` against the current candidate printed `revision=broken-900 unit=passed` and exited `0`.
- Limitation and risk: the script only reports the unit result and candidate revision. It does not validate version/schema alignment or destructive-migration recovery.

## Unresolved Risks

- The intended release version is unclear because the manifest and application configuration disagree. The release owner must select the authoritative version and provide matching candidate files.
- The manifest schema is stale or incorrect relative to both the application requirement and migration target. The release owner must reconcile all three values and demonstrate their compatibility.
- No recovery mechanism exists for the dropped table. The database owner must supply and test a concrete recovery command against representative data before release.

## Prioritized Next Actions

1. Release owner: reconcile `manifest.txt` and `application.conf` to one candidate version and provide reviewed evidence that resolves the alignment gate.
2. Application and database owners: reconcile manifest schema, required application schema, and migration target, then provide current compatibility evidence that resolves the alignment gate.
3. Database owner: design and safely test a recovery command for `retained_exports`, retaining the command and successful recovery evidence to resolve the destructive migration recovery gate.
4. Verification owner: rerun `verify.sh` after the candidate files change and retain output tied to the resulting revision to re-establish the unit verification gate.

## Evidence Index

### Inspected Paths

- `showcase-skills/release-readiness-checker/skill/release-readiness-checker/SKILL.md`
- `supporting-skills/document-update-discipline/SKILL.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/behavior/cycle-1/case-2/request.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/RELEASE.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/manifest.txt`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/application.conf`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/migrate.sql`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/rollback.md`
- `showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/verify.sh`
- `package.json`

### Commands and Outputs

- Working directory for every command: `<repo-root>`.
- `cat showcase-skills/release-readiness-checker/skill/release-readiness-checker/SKILL.md && printf '\n---DOCUMENT-DISCIPLINE---\n' && cat supporting-skills/document-update-discipline/SKILL.md && printf '\n---REQUEST---\n' && cat showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/behavior/cycle-1/case-2/request.md` — exited `0`; loaded the governing skill instructions and isolated request.
- `rg --files showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2` — exited `0`; enumerated the six candidate files.
- `for file in showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/RELEASE.md showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/manifest.txt showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/application.conf showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/migrate.sql showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/rollback.md showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/verify.sh; do printf '\n--- %s ---\n' "$file"; cat "$file"; done` — exited `0`; inspected all candidate-owned contract, identity, configuration, migration, recovery, and verification sources before running verification.
- `sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/verify.sh; status=$?; printf 'exit_status=%s\n' "$status"` — verification printed `revision=broken-900 unit=passed`; the wrapper then failed because `status` is read-only in zsh, so this wrapper invocation does not establish the script exit status.
- `sh showcase-skills/release-readiness-checker/evaluation/release-readiness-checker/fixtures/behavior/case-2/verify.sh; exit_code=$?; printf 'exit_status=%s\n' "$exit_code"` — printed `revision=broken-900 unit=passed` and `exit_status=0`; establishes current unit verification success for revision `broken-900`.
- `cat package.json` — exited `0`; confirmed that the required `format:check` command runs Prettier in check-only mode.
- `npm run format:check` — exited `0`; reported that all configured files use Prettier formatting. The configured path set does not include this report, so this result does not directly prove the report's formatting.

### Deliberately Not Run

- `migrate.sql` was not executed because it destructively drops a table.
- No recovery command was run because the candidate provides none.
- No publish, deploy, tag, upload, promotion, or external-state command was run.
