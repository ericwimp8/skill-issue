# Bug Reproduction Package

## Summary

The local export command crashes when given a valid record that has an `id` but omits `nickname`. The supplied contract permits an omitted `nickname` and requires the command to complete and print the record's `id`.

## Evidence Status

**Reproduced.** The missing-`nickname` case failed in 3 of 3 local runs. This establishes repeatability only in the environment recorded below; the sparse incident note does not identify the reporter's version, operating system, exact record, or observed frequency.

## Environment

- Checkout revision at inspection time: `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`, from `git rev-parse HEAD`.
- Fixture status: the `fixtures/description/trial-2/` directory was untracked in this checkout, from `git status --short -- fixtures/description/trial-2/`; the revision therefore does not identify the fixture contents.
- Operating system: macOS 26.2, from `sw_vers -productName` and `sw_vers -productVersion`.
- Node.js: v23.6.1, from `node --version`.
- Reporter environment: unknown. `fixtures/description/trial-2/incident.md` explicitly states that version and operating system were not included.

## Prerequisites

- A local checkout containing the unchanged files under `fixtures/description/trial-2/`.
- Node.js available on `PATH`.
- No network access, account, permissions, service, or production data is required.
- A JSON record containing an `id` and no `nickname`; the synthetic record used here is `{"id":"record-1"}`.

## Minimal Reproduction

Starting state: open a shell in `fixtures/description/trial-2/` with the supplied `export.js` unchanged.

1. Run:

   ```sh
   node export.js '{"id":"record-1"}'
   ```

2. Observe the process output and exit status.

The command is sufficient to reproduce the observed failure; no setup beyond the prerequisites is required.

## Expected Behavior

The command completes and prints `record-1`. This expectation comes from `fixtures/description/trial-2/contract.md`, which states that records may omit `nickname` and that export must still complete and print the record's `id`.

## Actual Behavior

The command prints no record ID, throws `TypeError: Cannot read properties of undefined (reading 'toUpperCase')` at `export.js:2:29`, and exits with status 1. The divergence occurs when `export.js` evaluates `record.nickname.toUpperCase()` while `nickname` is absent.

## Evidence

- `fixtures/description/trial-2/incident.md`: original sparse report; says one record crashes and proposes missing `nickname` only as a possibility.
- `fixtures/description/trial-2/contract.md`: expected behavior and supported input shape.
- `fixtures/description/trial-2/export.js`: inspected implementation; line 2 accesses `record.nickname.toUpperCase()` and does not print `record.id`.
- Local command observation on 2026-07-20: three executions with `{"id":"record-1"}` each produced the same `TypeError` at `export.js:2:29` and exit status 1.
- Local comparison on 2026-07-20: `node export.js '{"id":"record-1","nickname":"ace"}'` printed `ACE` and exited 0.
- No secrets or personal data were used. The record values in this package are synthetic, so no redaction was required.

## Attempts and Variations

- Missing `nickname`: failed 3/3 times with the same exception and exit status 1.
- Present string `nickname` (`"ace"`): completed 1/1 time and printed `ACE`, narrowing the crash trigger to the absent value in this implementation.
- The present-`nickname` comparison also did not satisfy the contract's output requirement because it printed the nickname rather than the record ID.
- Other Node.js versions, operating systems, record shapes, and `nickname` values were not tested.

## Open Gaps

- **Reporter environment:** its absence prevents confirming that the local result matches the originally affected version and operating system. Resolve by obtaining the reporter's checkout/build revision, Node.js version, and operating system.
- **Original input and error artifact:** their absence prevents confirming that the reporter saw this exact exception rather than another crash. Resolve by obtaining a redacted failing record and retained stderr/exit status from the original run.
- **Affected population and frequency:** the note identifies only “one” record and provides no denominator or retry history. Resolve by rerunning the original record repeatedly and checking a bounded sample of records with and without `nickname`.
- **Cross-environment behavior:** only the environment above was exercised. Resolve by running the minimal command in the reporter's environment once it is known.

## Ready-to-File Issue

### Title

Export crashes when a record omits `nickname`

### Issue Body

The export contract permits records to omit `nickname` and requires export to complete and print the record's `id`. In the inspected fixture, a record without `nickname` instead causes the command to throw before printing the ID.

**Environment reproduced**

- macOS 26.2
- Node.js v23.6.1
- Checkout HEAD `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`; the fixture directory was untracked, so this revision does not identify its contents

**Steps**

From `fixtures/description/trial-2/`, run:

```sh
node export.js '{"id":"record-1"}'
```

**Expected**

The command completes and prints `record-1`, as specified in `fixtures/description/trial-2/contract.md`.

**Actual**

The command throws `TypeError: Cannot read properties of undefined (reading 'toUpperCase')` at `export.js:2:29`, prints no ID, and exits 1. The result occurred in 3 of 3 local runs.

With `{"id":"record-1","nickname":"ace"}`, the command completed 1 of 1 times but printed `ACE` rather than the required record ID.

The reporter's environment, original record, original error output, and observed frequency remain unknown. A redacted original input plus the reporter's revision, Node.js version, operating system, stderr, and exit status are needed to confirm that this is the same incident.
