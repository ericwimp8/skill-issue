# Description Trial 1

- Stage: representative
- Prompt: retained verbatim in `../prompts.md`
- Fresh agent: `/root/eval_code_testing/ctd_desc_high_parser`
- Session: `019f8265-0d3a-7ed2-8e1e-4047e8bc251d`
- Harness: Codex Desktop backed by Codex CLI `0.145.0-alpha.18`
- Model and reasoning: `gpt-5.6-sol`, high
- Fixture: `fixtures/description/parser`
- Target version: `72277b024cd12d7875792e3a677a73c9a13f294d266231c6fe5876930cfbdd55`

## Native Selection Evidence

Before producing or editing the test, the native trace records an `exec` tool call that reads `supporting-skills/code-testing-discipline/SKILL.md`. The retained JSONL also contains the session metadata, high-reasoning turn context, and final output. This pre-output read is the controlling selection evidence.

## Observable Output

The agent added a focused exported-interface regression test for a fractional HTTP-date delay, ran it successfully, and declined a production change because the source and executable reproduction showed no rounding defect. It stated that a production correction would require a failing exported-interface reproduction, governing precision contract, and traced caller path.

## Audit

- Native target load: pass
- Fresh qualified agent: pass
- Source-before-assertion behavior: pass
- Owned-interface test surface: pass
- Focused executable result: pass, 1 selected test passed
- Task quality: pass; the agent did not convert an unsupported report into a production change
- Description selection result: pass

## Cleanup

The generated test is retained as `output.test.js`. The shared parser input fixture was restored to its recorded clean baseline before another agent could use it.
