# Greeter Reproduction Package

## Summary

The greeter prints `hello world` instead of using the `name` value from the supplied JSON configuration. The documented command reproduces the failure with the supplied `example.json` fixture.

## Evidence Status

**Reproduced.** The documented path produced `hello world` in 3 of 3 consecutive local runs on 2026-07-20. This establishes consistent reproduction in the environment below; it does not establish behavior on other revisions or operating systems.

## Environment

- Repository revision: `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`, obtained with `git rev-parse HEAD`.
- Operating system: Darwin `25.2.0` on `arm64`, obtained with `uname -srm`.
- Shell entry point: the executable fixture declares `#!/bin/sh`.
- Fixture paths:
  - `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/greeter.sh`
  - `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/example.json`
- Material unknown: behavior on other repository revisions and operating systems was not tested.

## Prerequisites

- A checkout containing the fixture files at the paths above.
- A POSIX-compatible `/bin/sh` environment with `sed` available.
- Execute permission on `greeter.sh`; the inspected fixture was executable.
- No network access, account, privileged permission, or external service is required.

## Minimal Reproduction

Starting from the repository root at revision `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`:

1. Confirm that `example.json` contains `{"name":"Ada"}`.
2. Run:

   ```sh
   showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/greeter.sh \
     --config \
     showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/example.json
   ```

3. Observe the command output.

## Expected Behavior

The command should print `hello Ada`. The fixture README documents that `./greeter.sh --config <path>` prints `hello <name>` using the JSON file's `name` field, and the supplied configuration sets that field to `Ada`.

Source: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/README.md` and `example.json`.

## Actual Behavior

The command prints:

```text
hello world
```

The divergence occurs when the script reads the configuration: `greeter.sh` searches for a JSON key named `user`, while the documented contract and supplied example use `name`. Because the search returns no value, the final expansion falls back to `world`.

## Evidence

- Original report: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/report.md`.
- Documented contract: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/README.md`.
- Supplied input: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/example.json` contains `{"name":"Ada"}`.
- Implementation: `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/greeter.sh` extracts the `user` key and falls back to `world` when extraction is empty.
- Local observation on 2026-07-20: three consecutive executions of the documented path each printed `hello world`.
- Controlled variation: a temporary configuration containing `{"user":"Ada"}` printed `hello Ada`, supporting the key-mismatch explanation. The temporary file was removed after the check.
- Redaction: none required; the fixtures and captured outputs contain no secrets or personal account data.

## Attempts and Variations

- Repeated the documented command three times with the supplied `example.json`; all three runs printed `hello world`.
- Changed only the configuration key from `name` to `user` in a temporary file; the script printed `hello Ada`.
- No other code, configuration, runtime, or fixture variation was tested.

## Open Gaps

- Cross-platform and cross-revision frequency are unknown. Test the minimal reproduction on any additional supported environments if broader impact needs to be established.
- The intended remediation has not been executed or validated. Align the implementation with the README's `name`-field contract, then rerun the minimal reproduction and relevant project checks.

## Ready-to-File Issue

### Title

Greeter ignores the documented `name` configuration field

### Issue Body

The greeter prints `hello world` instead of using the supplied JSON configuration's `name` value.

**Environment**

- Revision: `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`
- OS: Darwin `25.2.0` (`arm64`)

**Steps to reproduce**

From the repository root:

1. Confirm `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/example.json` contains `{"name":"Ada"}`.
2. Run:

   ```sh
   showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/greeter.sh \
     --config \
     showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/fixtures/description/trial-1/example.json
   ```

**Expected**

`hello Ada`, as specified by the fixture README's `name`-field contract.

**Actual**

`hello world`

This reproduced in 3 of 3 consecutive local runs. The script currently searches the configuration for a `user` key, while the README and supplied example use `name`. A controlled temporary input using `{"user":"Ada"}` printed `hello Ada`.
