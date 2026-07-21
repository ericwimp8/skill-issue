# Body Case 1: Regression Red-Green

- Prompt and ground truth: `../cases.md`
- Fresh agent: `/root/eval_code_testing/ctd_body_regression`
- Session: `019f826b-0046-76e3-b48f-553bfdcdc076`
- Model and reasoning: `gpt-5.6-sol`, high
- Isolated workspace: `workspace/`
- Target version: `72277b024cd12d7875792e3a677a73c9a13f294d266231c6fe5876930cfbdd55`

## Native Evidence

The retained native trace records qualified runtime context, a pre-output read of the exact target, and the final response. The workspace preserves the final production correction and regression test. `pre-fix-test-output.txt` preserves the failing assertion before correction; `post-fix-test-output.txt` preserves the focused and broader passing runs.

## Ground-Truth Audit

- Read production contract before assertions: pass; the agent used the production JSDoc rule for non-positive values.
- Smallest owned interface: pass; the test calls exported `pageSize` directly.
- Observable outcomes: pass; assertions compare returned page sizes.
- Proportionate setup: pass; two additional inputs in the existing test file.
- Regression failure before correction: pass; retained output shows `0 !== 20`.
- Root correction: pass; production validation now rejects `parsed <= 0` at the concrete owner.
- Focused then broader validation: pass; focused file failed pre-fix, passed 3/3 post-fix, then `npm test` passed 3/3.
- Material failures: none.

## Result

Pass. No target refinement is indicated by this case.
