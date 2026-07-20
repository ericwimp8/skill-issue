# Validation Audit

## Structural and Contract Checks

- `quick_validate.py showcase-skills/api-change-impact-mapper/skill/api-change-impact-mapper` returned `Skill is valid!` before and after the campaign.
- The skill name, folder, frontmatter, description form, portable body, and OpenAI metadata satisfy the generation contracts.
- No reference, script, or asset exists because generation established no separate conditional detail, deterministic repeated operation, or consumed output material.
- Four fresh `gpt-5.6-sol` medium-reasoning description trials selected the exact target and returned its frozen hash.
- Three fresh body trials opened the exact target, inspected isolated connected fixtures, produced the required reports, and retained matching before/after fixture hashes.
- Every evaluation criterion passed; no material failure supported a refinement.

## Formatting and Diff Checks

- Repository-required `npm run format:check` passed after all artifacts were created.
- `git diff --check -- showcase-skills/api-change-impact-mapper` passed.
- A broader optional Prettier check over the showcase workspace identified style changes in the three retained agent outputs and one already-evaluated fixture file. Those files remain byte-for-byte unchanged to preserve the direct output hashes and fixture tree evidence. This optional surface is outside the repository's configured `format:check` paths.

## Hash Evidence

- Final skill SHA-256: `7b5dfd09fae4349f467f13aa4d7a85ddae21b831eef6dee0f93421d03cd4876e`.
- HTTP fixture tree SHA-256: `3354c05dffb6aa15259fe0aa276283c0cb152b131641857dd1e9b4f6d2e21c4c`.
- Event fixture tree SHA-256: `2252998ec9ad3f53799df7b5ee4737919f17fe77a2b4b079a80a2be088bc1309`.
- Library fixture tree SHA-256: `91b7e73f05a0062d2e91dadba11097a51fda8f7296d3036b973070b40922a268`.
- Case output hashes are recorded in their owning audit records and native evidence.

## Scope Check

All 55 substantive files present before this audit were untracked additions under `showcase-skills/api-change-impact-mapper/`. Concurrent changes outside that directory were observed but not modified, reformatted, reverted, or included in this validation.
