# Alternate-Mode Evaluation Report

## Contract Interpretation

- Goal: turn an approved layout brief into a generated `.prolayout` artifact.
- Intended use: artifact generation when the proprietary renderer is available.
- Expected behavior: render the approved brief and handle renderer failures through a generalized workflow correction.
- Expected result: a renderer-produced artifact whose inspector record proves the brief's reading-order and label requirements.
- Preserved boundary: keep the artifact type, renderer-based creation, approved-brief input, and retry intent; do not encode case-specific inputs or expected outputs in the skill.

## Review-Before-Update Proposal

The retained failure is owned by the target body's renderer-failure instruction. That sentence directs the agent to copy the current fixture name and expected pixel positions into the reusable skill, which would make evaluation data a second semantic owner and overfit later executions. The preceding renderer-creation instruction and the frontmatter remain coherent and should stay unchanged.

Replace the complete target body with the following text after reviewer approval:

```markdown
Create the artifact with the available renderer. If rendering fails, retain the input and renderer diagnostics, identify the generalized workflow instruction that owns the failure, and revise only that instruction before retrying. Keep case-specific inputs and expected output details in the execution record rather than this skill.
```

This is a proposal only. No target update has been applied.

## External Execution Route

After the reviewer approves the proposal and the campaign owner applies it, use a clean renderer-equipped environment that also provides the proprietary artifact inspector:

1. Record the applied target content hash and the renderer and inspector names and versions.
2. Start a clean execution with the approved target and only this layout input: `Create a two-panel onboarding layout with a title panel and a checklist panel. Preserve reading order and expose both panel labels to the artifact inspector.`
3. Invoke the environment's native renderer once, save its unmodified execution transcript, and retain the produced `.prolayout` file without hand-editing it.
4. Run the native artifact inspector against that file and retain its unmodified structured output.
5. Return the evidence bundle to this campaign for audit. If rendering fails, return the diagnostics and stop without modifying the skill or retrying from altered input.

## Evidence Required To Resume

The campaign can resume only when the returned bundle contains:

- the reviewer decision and the content hash of the exact approved target used for execution;
- the exact external input, renderer and inspector identities and versions, execution environment identity, and native invocation record;
- the renderer exit status, complete diagnostics, and output path;
- on renderer success, the unchanged `.prolayout` artifact with its cryptographic hash and byte size;
- the unmodified inspector output identifying both panels, their exposed labels, and the title-panel-before-checklist-panel reading order;
- on renderer failure, the complete native failure record and confirmation that neither the target nor the input was changed.

Renderer capability and artifact behavior remain unobserved in the current harness. No artifact pass or failure is claimed.
