# Description Round 1 Prompts

## Representative Trial 1

> Work in `<repo-root>`. Independently inspect `evaluations/scenario-skill-refinement/code-testing-discipline/fixtures/description/parser`. A bug report says an HTTP-date exactly 1.5 seconds after now must preserve that fractional delay, but current behavior rounds it down. Add the smallest meaningful regression test for that report and explain the evidence you would require before changing production code. Run the focused test. Modify only this fixture. Do not read other evaluation campaign outputs. Report files changed, commands, and results.

## Representative Trial 2

> Work in `<repo-root>`. Independently inspect `evaluations/scenario-skill-refinement/code-testing-discipline/fixtures/description/queue`. Extend the automated coverage for draining a queue twice so a future storage refactor cannot leave already-returned jobs pending. Choose the test surface and assertions you consider appropriate, run the focused test, and modify only this fixture. Do not read other evaluation campaign outputs. Report files changed, commands, and results.

## Confirmation Trial 3

> Work in `<repo-root>`. Independently inspect `evaluations/scenario-skill-refinement/code-testing-discipline/fixtures/description/token`. Review the current automated test coverage for tokens four characters or shorter and add only the coverage needed to protect the public redaction behavior. Run the focused test and modify only this fixture. Do not read other evaluation campaign outputs. Report files changed, commands, and results.

## Confirmation Trial 4

> Work in `<repo-root>`. Independently inspect `evaluations/scenario-skill-refinement/code-testing-discipline/fixtures/description/batch`. A caller reports that invalid batch sizes must fail with the documented public error rather than hang or expose internal loop behavior. Add focused automated coverage for that contract, run it, and modify only this fixture. Do not read other evaluation campaign outputs. Report files changed, commands, and results.

## Selection Ground Truth

- Each prompt naturally requests creation, editing, or running of automated code tests without naming a skill or revealing expected selection.
- A pass requires a native pre-output read of `supporting-skills/code-testing-discipline/SKILL.md` or a native `codex.skill.injected` event for the advertised candidate.
- Final prose, a `Skills Used` list, and answer similarity are excluded as selection evidence.
- Each trial uses a fresh independent agent with no inherited turns and a unique fixture.
- Observable task quality is recorded separately from selection success.
