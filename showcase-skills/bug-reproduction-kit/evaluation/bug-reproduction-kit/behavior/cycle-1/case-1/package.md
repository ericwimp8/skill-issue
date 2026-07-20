# Reproduction Package

## Summary

`total.py` joins the two price arguments as text instead of printing their numeric sum. With inputs `12.50` and `4.25`, it prints `12.504.25`, matching the malformed output recorded in the supplied report.

## Evidence Status

**Reproduced.** The minimal command produced `12.504.25` in 3 of 3 consecutive runs on the environment below. This frequency applies only to the stated local revision, runtime, and inputs; the reporter's original command and environment remain unknown.

## Environment

- Project revision: Git commit `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`, obtained with `git rev-parse HEAD`.
- Operating system: Darwin `25.2.0`, ARM64, obtained with `uname -srm`.
- Runtime: Python `3.14.2`, obtained with `python3 --version`.
- Fixture source: `fixtures/behavior/case-1/`.
- Observation time: `2026-07-20T00:23:54Z` (UTC).
- Material unknowns: the reporter's Python version, operating system, exact command, exact argument values, revision, and observed frequency were not supplied.

## Prerequisites

- A local checkout containing `fixtures/behavior/case-1/README.md` and `fixtures/behavior/case-1/total.py` at the revision above.
- A working `python3` command.
- No account, network access, elevated permission, service, configuration, or production data is required.

## Minimal Reproduction

Starting from `showcase-skills/bug-reproduction-kit/evaluation/bug-reproduction-kit/`:

1. Run:

   ```sh
   cd fixtures/behavior/case-1
   python3 total.py 12.50 4.25
   ```

2. Observe the single output line:

   ```text
   12.504.25
   ```

## Expected Behavior

The command should print the numeric sum of the two decimal prices, so these inputs should produce `16.75`. The contract is stated in `fixtures/behavior/case-1/README.md`: `python3 total.py <first-price> <second-price>` must print the numeric sum.

## Actual Behavior

The command prints `12.504.25`. The divergence occurs when `total.py` evaluates `sys.argv[1] + sys.argv[2]`: both values are command-line strings, so `+` concatenates them rather than adding numeric prices.

## Evidence

- `fixtures/behavior/case-1/report.md`: original report states that decimal invoice totals look wrong and records one observation of `12.504.25`; it explicitly lacks the Python version and exact command.
- `fixtures/behavior/case-1/README.md`: expected-behavior contract and invocation shape.
- `fixtures/behavior/case-1/total.py`: retained implementation at the failure boundary.
- Local command output captured at `2026-07-20T00:23:54Z`: three consecutive runs of `python3 total.py 12.50 4.25` each printed `12.504.25`.
- Fixture integrity check: SHA-256 hashes before and after the checks were identical (`README.md`: `94bb5e656511b0ddabc9d42993b211e623fc03a86b26c5f110833e5ccb7664e5`; `report.md`: `68eb36674c7646993b2d2a061ca73b5ffe939019d6392dd49655d54b720beaf2`; `total.py`: `c57460e3e5f5061424358c4a0d0de1e24b54706ed0313bf69fb7c01fa60e123a`).
- No secrets or personal data were encountered, so no redaction was required.

## Attempts and Variations

- Repeated `python3 total.py 12.50 4.25` three times; all three runs printed `12.504.25`.
- Ran `python3 total.py 0.50 0.25`; it printed `0.500.25`, showing the same divergence with different decimal values.
- Ran `python3 total.py 12 4`; it printed `124`, showing that the behavior is argument-string concatenation and is not limited to decimal punctuation.
- No tested variation produced a numeric sum.

## Open Gaps

- The reporter's exact command and argument values are unknown. This prevents confirmation that the original setup was identical; obtain the shell history or original invocation to resolve it.
- The reporter's runtime, operating system, revision, and frequency are unknown. This limits comparison between the original observation and this local reproduction; obtain `python3 --version`, `uname -srm`, the project revision, and retry counts from the reporter.
- Behavior with malformed prices or missing arguments was outside the reported path and was not investigated.

## Ready-to-File Issue

### Title

Invoice total command concatenates price arguments instead of adding them

### Issue Body

The invoice-total command prints joined argument text instead of the numeric sum.

**Environment**

- Revision: `9ccef67e5f40c8954d1fe4bcc0f8ba009d4d6823`
- OS: Darwin `25.2.0` ARM64
- Python: `3.14.2`

**Steps to reproduce**

From `fixtures/behavior/case-1/`, run:

```sh
python3 total.py 12.50 4.25
```

**Expected**

The command prints the numeric sum, `16.75`, as required by `README.md`.

**Actual**

The command prints `12.504.25`. This occurred in 3 of 3 consecutive local runs. `total.py` evaluates `sys.argv[1] + sys.argv[2]`, which concatenates the two command-line strings.

**Additional observations**

- `python3 total.py 0.50 0.25` printed `0.500.25`.
- `python3 total.py 12 4` printed `124`.
- The supplied reporter environment and exact original command are unavailable.
