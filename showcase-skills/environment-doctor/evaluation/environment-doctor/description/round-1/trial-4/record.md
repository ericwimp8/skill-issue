# Trial Record

## Identity

- Trial identity: `/root/environment_doctor/env_desc_4`
- Model: inherited Codex agent model; exact model identifier was not exposed to the trial
- Reasoning: inherited/default reasoning configuration; exact effort label was not exposed to the trial
- Selection: selected naturally from candidate frontmatter because the request explicitly concerns executable resolution and PATH precedence

## Target

- Candidate target: `showcase-skills/environment-doctor/skill/environment-doctor/SKILL.md`
- Target SHA-256: `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203`
- Complete target reads after selection: `1`

## Exact Diagnostic Commands

```sh
fixture_root="$(cd showcase-skills/environment-doctor/fixtures/script/root && pwd -P)"
DIAGNOSTIC_SECRET=trial-only-value PATH="$fixture_root/toolchain-primary/bin:$fixture_root/toolchain-secondary/bin:$fixture_root/toolchain-primary/bin:$PATH" python3 showcase-skills/environment-doctor/skill/environment-doctor/scripts/diagnose.py --root "$fixture_root" --output-dir showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-4/output --tool node --env DIAGNOSTIC_SECRET --expect-path-before toolchain-primary/bin toolchain-secondary/bin
```

- Diagnostic exit: `1`
- Exit interpretation: expected because the duplicate PATH state produces a warning
- Canonical root handling: the fixture root was resolved inside the shell command and retained only in the shell variable used to construct the diagnostic child environment
- Environment isolation: the repeated PATH prefix and synthetic selected variable were applied only to the diagnostic process invocation

## Source Hashes

| Source                          | Before                                                             | After                                                              |
| ------------------------------- | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| Candidate `SKILL.md`            | `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203` | `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203` |
| `scripts/diagnose.py`           | `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430` | `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430` |
| Fixture `.node-version`         | `5378796307535df3ec8d8b15a2e2dc5641419c3d3060cfe32238c0fa973f7aa3` | `5378796307535df3ec8d8b15a2e2dc5641419c3d3060cfe32238c0fa973f7aa3` |
| Fixture `mismatch.node-version` | `f14b4987904bcb5814e4459a057ed4d20f58a633152288a761214dcd28780b56` | `f14b4987904bcb5814e4459a057ed4d20f58a633152288a761214dcd28780b56` |
| Fixture primary `mystery-tool`  | `626efabe18f06a84cadd6e991752b189b0a2738e72efea358122e3e4b86aeaff` | `626efabe18f06a84cadd6e991752b189b0a2738e72efea358122e3e4b86aeaff` |
| Fixture primary `node`          | `20ccc0498ef0b7960d0fe68b521d2c0da36f1931c8f12b0ebc22421b39b01879` | `20ccc0498ef0b7960d0fe68b521d2c0da36f1931c8f12b0ebc22421b39b01879` |
| Fixture secondary `node`        | `2e042e73b2d6e214aa4e90c7e4c8300fab299658b2352ee5e27d665364045e61` | `2e042e73b2d6e214aa4e90c7e4c8300fab299658b2352ee5e27d665364045e61` |
| Trial `request.md`              | `5127945e9d05601cf5715629376064a90d53471aa5001eba4c95aa347af940c5` | `5127945e9d05601cf5715629376064a90d53471aa5001eba4c95aa347af940c5` |

## Criteria

- Pass: candidate selected from frontmatter before the complete target was read.
- Pass: only `node`, the primary-before-secondary expectation, and `DIAGNOSTIC_SECRET` were selected.
- Pass: primary `node` was selected at PATH index `0` and reported `20.11.1`.
- Pass: the selected variable was reported only as `set`, with its value omitted.
- Pass: duplicate precedence remained unresolved rather than being converted into a satisfied conclusion.
- Pass: the report requires consent before any actual setup change.
- Pass: candidate, script, request, and fixture source hashes were unchanged.

## Privacy

- Structured and human-readable diagnostic outputs use `<root>` rather than retaining the canonical checkout path.
- The selected variable value is absent from `output/report.txt` and `output/evidence.json`.
- The synthetic input value appears only in the reproducibility command evidence requested for this trial.
- Home paths in structured evidence are normalized to `~`; the exact tested host is identified only as POSIX.

## Cleanup

- A failed first wrapper attempt was removed before the final run; it had completed the diagnostic but collided with zsh's reserved `status` variable while capturing the exit.
- Final retained files are limited to `output/report.txt`, `output/evidence.json`, `output.md`, `record.md`, and `native-evidence.log`.
- No temporary evaluation directories remain.
