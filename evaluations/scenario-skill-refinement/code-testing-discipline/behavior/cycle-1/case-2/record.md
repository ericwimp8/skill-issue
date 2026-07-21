# Body Case 2: Replace a Brittle Test

- Prompt and ground truth: `../cases.md`
- Fresh agent: `/root/eval_code_testing/ctd_body_brittle`
- Session: `019f826b-6c40-7373-ae54-f44bc7b5c369`
- Model and reasoning: `gpt-5.6-sol`, high
- Isolated workspace: `workspace/`
- Target version: `72277b024cd12d7875792e3a677a73c9a13f294d266231c6fe5876930cfbdd55`

## Native Evidence

The retained native trace records qualified runtime context, a pre-output read of the exact target, and the final response. The workspace preserves the corrected test, while `test-output.txt` preserves focused and broader command results.

## Ground-Truth Audit

- Established behavior from production source: pass.
- Smallest owned interface: pass; the replacement calls exported `orderTotal`.
- Observable outcome assertion: pass; the test asserts the returned discounted total.
- Removed implementation coupling: pass; source-file reading and expression matching were removed.
- Production restraint: pass; correct production code remained unchanged.
- Focused then broader validation: pass; focused file passed, then `npm test` passed.
- Material failures: none.

## Result

Pass. No target refinement is indicated by this case.
