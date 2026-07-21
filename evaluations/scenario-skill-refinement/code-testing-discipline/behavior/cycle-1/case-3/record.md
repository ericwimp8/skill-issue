# Body Case 3: Choose the Owned Interface

- Prompt and ground truth: `../cases.md`
- Fresh agent: `/root/eval_code_testing/ctd_body_interface`
- Session: `019f826c-da6f-7d81-81f8-25d7de5a1eba`
- Model and reasoning: `gpt-5.6-sol`, high
- Isolated workspace: `workspace/`
- Target version: `72277b024cd12d7875792e3a677a73c9a13f294d266231c6fe5876930cfbdd55`

## Native Evidence

The retained native trace records qualified runtime context, a pre-output read of the exact target, and the final response. The workspace preserves the endpoint-level tests, while `test-output.txt` preserves focused and broader command results.

## Ground-Truth Audit

- Established behavior from production source: pass.
- Smallest owned interface: pass; tests call exported `profileEndpoint` rather than private `parseProfile`.
- Observable outcome assertions: pass; each case asserts status, content type, and public JSON error body.
- Required inputs: pass; malformed JSON and blank name are covered separately.
- Production restraint: pass; already-correct production behavior remained unchanged.
- Focused then broader validation: pass; focused file passed 3/3, then `npm test` passed 3/3.
- Material failures: none.

## Result

Pass. No target refinement is indicated by this case.
