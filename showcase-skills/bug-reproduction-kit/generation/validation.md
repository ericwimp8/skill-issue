# Generation Validation Evidence

- Target hash at generation and campaign conclusion: `c6679f687223e9da3aceef0040c04410ec9060336f432237f54913ec13e44392`.
- Structural validator: `quick_validate.py showcase-skills/bug-reproduction-kit/skill/bug-reproduction-kit` returned `Skill is valid!` before and after evaluation.
- Frontmatter: required `name` and `description` are present; the lowercase hyphenated name matches the folder.
- Description: one concise purpose sentence followed by one concise use-boundary sentence.
- Body: behavior-changing instructions own evidence boundaries, reproduction context, minimal reproduction, missing-information handling, and the output package.
- Resources: no script, reference, or asset was generated because the behavior requires contextual investigation rather than one repeated deterministic operation or bundled output material.
- OpenAI metadata: `agents/openai.yaml` contains only `display_name` and `short_description`; no unrequested default prompt or capability claim is present.
- Intake completion criteria: each criterion is represented in the written target and was handed to runtime evaluation at `../evaluation/bug-reproduction-kit/contract.md`.
