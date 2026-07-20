# Bug Reproduction Package

## Summary

The historical report says the label formatter returned a blank label for an ordinary two-word project name. The reported failure could not be reproduced with the current fixture: representative two-word input produced the expected uppercase label in every attempt. The original input, invocation, revision, environment, and output evidence are unavailable, so the historical observation cannot be confirmed or ruled out.

## Evidence Status

**Not reproduced.** The current utility produced a nonblank label in 4 of 4 representative two-word attempts: three runs with one quoted argument and one run with two unquoted arguments. This result applies only to the current fixture and environment. Confidence about the historical failure is low because none of its material reproduction details were retained.

## Environment

- Repository `HEAD`: `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`, from `git rev-parse HEAD` on 2026-07-20.
- Fixture state: the case-2 fixture directory is untracked according to `git status --short -- <fixture>`, so the repository revision does not identify these exact fixture contents.
- Fixture content hashes (SHA-256): `README.md` `6cf8c6f08eea22b9466163c653503726acf9cc3030ff9bfaab23b3f83d5a9d7d`; `label.sh` `28189429eaf8823f476dae97a5c81d9da1a7f88c483c8291c5b57d7eb0b85624`; `report.md` `ce53fd1380256a67cfc4a806b8305778c16e5c957f4d07c8cb57773d079ae86b`.
- Operating system: macOS 26.2 build 25C56, Darwin 25.2.0, arm64, from `sw_vers` and `uname -a`.
- Interpreter and utility: `/bin/sh` and `/usr/bin/tr`, resolved in the current shell.
- Historical revision, operating system, shell, `tr` implementation, locale, working directory, invocation method, and environment are unknown.

## Prerequisites

- Start in the repository root.
- Use the executable fixture at `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/behavior/case-2/label.sh`.
- No account, network access, service, configuration, or data setup is required by the current utility.
- To reproduce the historical observation exactly, the original project name and invocation context are required; both are unavailable.

## Minimal Reproduction

Starting from the repository root, run:

1. Execute `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/behavior/case-2/label.sh 'ordinary project'`.
2. Observe the single output line.
3. Repeat twice to measure the current result frequency.

Observed on 2026-07-20: all 3 runs printed `ORDINARY PROJECT` followed by a newline.

This is a bounded attempt using representative input. It is not the historical command because the original name and invocation are unavailable.

## Expected Behavior

For a supplied project name, the utility should print the name in uppercase while preserving spaces. Empty input should print an empty line. The source is the fixture contract in `fixtures/behavior/case-2/README.md`; the current `label.sh` implementation also passes all arguments through `printf '%s\n' "$*"` and uppercases lowercase characters with `tr`.

For the representative input `ordinary project`, the expected output is `ORDINARY PROJECT` followed by a newline.

## Actual Behavior

The current utility matched the documented behavior. Each of three quoted-input runs printed `ORDINARY PROJECT` followed by a newline; the captured bytes were `4f5244494e4152592050524f4a4543540a` each time. No divergence occurred in the current attempts.

The only retained description of the historical actual behavior is `report.md`: it says a blank label was returned for an ordinary two-word project name last week. There is no retained output capture or precise failure boundary.

## Evidence

- Contract: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/behavior/case-2/README.md`.
- Implementation: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/behavior/case-2/label.sh`.
- Historical report: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/behavior/case-2/report.md`.
- Current command output, byte output, environment commands, file metadata, hashes, and Git state were observed in the investigation shell on 2026-07-20. No separate log artifact was created.
- No secrets or personal data were present in the shared command inputs or outputs; no redaction was required.

## Attempts and Variations

- Quoted representative input, `label.sh 'ordinary project'`: nonblank expected output in 3 of 3 runs.
- Unquoted representative input, `label.sh ordinary project`: nonblank expected output in 1 of 1 run, showing that the current use of `$*` joins these two arguments with a space.
- Empty-input control, `label.sh`: printed only newline byte `0a`, matching the documented empty-input behavior and confirming that blank output is reachable when no input is supplied.
- No historical input or environment variations could be tested because those values were not retained.

## Open Gaps

- **Original project name and exact command:** needed to determine whether the formatter received no input, different quoting, shell expansion, or unexpected characters. Next action: obtain the reporter's shell history, script invocation, or surrounding caller logs.
- **Historical output capture:** needed to distinguish an empty line from invisible characters, display-layer suppression, or downstream loss. Next action: obtain raw stdout bytes or a terminal/application capture from the failure.
- **Historical revision and file contents:** needed to compare the failing implementation with the current 59-byte script. Next action: identify the checkout, commit, release, or archived script used at the time.
- **Historical environment:** shell, operating system, locale, `tr` implementation, working directory, and caller are needed to test environment-specific behavior. Next action: obtain machine diagnostics or CI/job metadata.
- **Occurrence count:** the report does not state how many attempts failed or whether retries succeeded. Next action: ask the reporter for attempt and failure counts.

These gaps prevent an exact reproduction and prevent attributing the historical blank label to the current utility.

## Ready-to-File Issue

### Title

Historical blank output from label formatter with two-word project name

### Issue Body

The label formatter was reported to return a blank label for an ordinary two-word project name. The original project name, exact command, checkout revision, machine details, occurrence count, and output capture were not retained.

The current fixture documents that `./label.sh <project name>` prints the supplied name in uppercase with spaces preserved, while empty input prints an empty line. On macOS 26.2 build 25C56 (Darwin 25.2.0 arm64), using `/bin/sh` and `/usr/bin/tr`, the current script produced `ORDINARY PROJECT` followed by a newline in 3 of 3 runs of:

```sh
showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/behavior/case-2/label.sh 'ordinary project'
```

It also produced the same output in 1 of 1 run with two unquoted arguments. An empty-input control printed only a newline, as documented. The blank-output failure occurred in 0 of 4 representative two-word attempts against the current fixture.

The fixture is untracked in the current checkout, so repository `HEAD` `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823` does not identify its exact contents. The tested `label.sh` SHA-256 is `28189429eaf8823f476dae97a5c81d9da1a7f88c483c8291c5b57d7eb0b85624`.

To continue the investigation, obtain the original name and command, raw output or caller logs, the historical script or revision, machine and locale details, and the number of failed versus successful attempts. These are required to execute the reported path and locate the divergence without guessing.
