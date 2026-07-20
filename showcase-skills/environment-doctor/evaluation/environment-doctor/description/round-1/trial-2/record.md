# Trial Record

## Runtime Identity

- Identity: `/root/environment_doctor/env_desc_2`
- Model: inherited runtime model; exact identifier was not exposed
- Reasoning: inherited runtime setting; exact level was not exposed
- Selection: selected naturally from the candidate description because the request concerns an absent executable, environment state, PATH precedence, and a version mismatch

## Target

- Candidate: `showcase-skills/environment-doctor/skill/environment-doctor/SKILL.md`
- SHA-256: `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203`
- The complete selected target was read once after applicability was decided from the request and candidate frontmatter.

## Commands

Initial diagnostic command:

```sh
trial='showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2'; root=$(cd showcase-skills/environment-doctor/fixtures/script/root && pwd -P); PATH="$root/toolchain-primary/bin:$root/toolchain-secondary/bin:$PATH" env -u INTENTIONALLY_UNSET python3 showcase-skills/environment-doctor/skill/environment-doctor/scripts/diagnose.py --root "$root" --output-dir "$trial/output" --tool absent-tool --env INTENTIONALLY_UNSET --expect-path-before toolchain-primary/bin toolchain-secondary/bin --version-file node mismatch.node-version > "$trial/native-evidence.log" 2>&1; status=$?; printf 'exit_code=%s\n' "$status" >> "$trial/native-evidence.log"; exit "$status"
```

The diagnostic created the requested output and returned `1`. The zsh wrapper then rejected reserved variable `status`, also exiting `1`. The diagnostic exit was independently confirmed with this disposable rerun:

```sh
trial='showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2'; tmp=$(mktemp -d /tmp/environment-doctor-trial2.XXXXXX); trap 'rm -rf "$tmp"' EXIT; root=$(cd showcase-skills/environment-doctor/fixtures/script/root && pwd -P); PATH="$root/toolchain-primary/bin:$root/toolchain-secondary/bin:$PATH" env -u INTENTIONALLY_UNSET python3 showcase-skills/environment-doctor/skill/environment-doctor/scripts/diagnose.py --root "$root" --output-dir "$tmp/output" --tool absent-tool --env INTENTIONALLY_UNSET --expect-path-before toolchain-primary/bin toolchain-secondary/bin --version-file node mismatch.node-version > /dev/null 2>&1; rc=$?; printf 'diagnostic_exit_code=%s\n' "$rc" > "$trial/native-evidence.log"; printf 'temporary_output_cleaned=true\n' >> "$trial/native-evidence.log"; exit 0
```

The fixture root was resolved canonically only inside each shell command before constructing the diagnostic child's PATH. No expanded checkout path was retained.

## Exit And Criteria

- Diagnostic exit: `1`, expected because the bounded inspection contains an error finding and warnings
- Absent command: passed; `absent-tool` remained `unavailable`
- Variable state: passed; `INTENTIONALLY_UNSET` remained `unset` with its value omitted
- PATH order: passed; secondary did not precede primary
- Version comparison: passed; actual Node `20.11.1` mismatched declared `22`
- Platform boundary: passed; POSIX behavior recorded separately for macOS `26.2`, Darwin `25.2.0`, `arm64`
- Read-only behavior: passed; the generated report states that no changes were made

## Source Integrity

Before and after SHA-256 hashes matched:

| Source                          | SHA-256                                                            |
| ------------------------------- | ------------------------------------------------------------------ |
| Candidate `SKILL.md`            | `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203` |
| `scripts/diagnose.py`           | `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430` |
| Fixture `.node-version`         | `5378796307535df3ec8d8b15a2e2dc5641419c3d3060cfe32238c0fa973f7aa3` |
| Fixture `mismatch.node-version` | `f14b4987904bcb5814e4459a057ed4d20f58a633152288a761214dcd28780b56` |
| Fixture primary `mystery-tool`  | `626efabe18f06a84cadd6e991752b189b0a2738e72efea358122e3e4b86aeaff` |
| Fixture primary `node`          | `20ccc0498ef0b7960d0fe68b521d2c0da36f1931c8f12b0ebc22421b39b01879` |
| Fixture secondary `node`        | `2e042e73b2d6e214aa4e90c7e4c8300fab299658b2352ee5e27d665364045e61` |

## Privacy And Cleanup

- Reports replace the fixture root with `<root>` and omit the selected environment value.
- No canonical machine checkout path is present in retained artifacts.
- The disposable confirmation directory was removed by its exit trap.
- The trial contains only `request.md`, `output/report.txt`, `output/evidence.json`, `output.md`, `record.md`, and `native-evidence.log`.

## Validation Commands

```sh
jq empty showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2/output/evidence.json
npx prettier --check showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2/output.md showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2/record.md
find showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2 -type f -print | sort
pattern=$(printf '\057\125\163\145\162\163\057|\163\150\141\156\156\157\156\150\141\154\154|\104\145\163\153\164\157\160\057\144\145\166\145\154\157\160\155\145\156\164'); rg -n "$pattern" showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2/output showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2/output.md showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2/record.md showcase-skills/environment-doctor/evaluation/environment-doctor/description/round-1/trial-2/native-evidence.log
```
