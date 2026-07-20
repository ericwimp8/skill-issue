# Generation Validation Evidence

- Initial target content hash: `31170d579a9c825af4f40d072cbb2e8e1f6d072523818c8957016d224b11d71a`.
- Initial script content hash: `0fa4b59808edb66648bfd04ab561e2cf79824dfb9868f99084ff51a5958abe10`.
- Structural validator: `quick_validate.py showcase-skills/environment-doctor/skill/environment-doctor` returned `Skill is valid!`.
- Python construction: `python3 -m py_compile` passed for the bundled script and direct-validation harness.
- Direct script validation: `python3 showcase-skills/environment-doctor/script-validation/validate_diagnose.py` returned `environment-doctor script validation passed`.
- Frontmatter: required `name` and two-sentence what/when description are present; the lowercase hyphenated name matches the folder.
- Body: inspection selection, deterministic execution, interpretation, consent, and platform boundary each have one concise semantic owner.
- Resources: one standard-library Python script is required; no reference or asset has a distinct necessary meaning.
- OpenAI metadata: supported interface fields and `policy.allow_implicit_invocation: true` are present; no default prompt was added.
- Runtime criteria: handed to `showcase-skills/environment-doctor/evaluation/environment-doctor/contract.md`.
- Final target content hash: `502a690ae603b8f0399fb6e98d66753acc0813f83dc1b769ce85732df261a203`.
- Final script content hash: `d3f235daeec5c1a90b3696619e4249a8018583a6fa3a6f0761c3c7c26fcab430`.
- Final direct-harness hash: `89108ef5cda87a68a1050d742a2f51d6f4287b05db1bb3d3964cfc81bd545a72`.
- Evaluation result: description passed 4/4; body passed after one evaluator-coverage refinement and fresh verification.
