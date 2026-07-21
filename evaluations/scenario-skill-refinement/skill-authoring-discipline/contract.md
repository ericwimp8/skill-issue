# Skill Authoring Discipline Evaluation Contract

## Target

- Canonical path: `evaluations/scenario-skill-refinement/skill-authoring-discipline/skill/`
- Initial `SKILL.md` SHA-256: `7fcb2f074490433c6eb6dc3f288d3af1caf353f1cd86b8c65be91d6f017d8bb4`
- Initial `agents/openai.yaml` SHA-256: `0262113579c7fb0884333d09c47f27a9df9c7426c47688d34521457ae7f57fdf`
- Advertised path: `.codex/skills/skill-authoring-discipline`

## Interpreted Outcome

1. **Goal:** Make Codex skill creation, updates, and reviews produce concise, decision-useful skills whose loaded instructions change agent behavior without becoming full operating manuals.
2. **Intended use:** Creating, updating, or reviewing Codex skills, including their frontmatter, bodies, metadata, references, and folder contents.
3. **Expected behavior:** Prefer governing constraints over exhaustive detail; keep descriptions in concise what-it-is and when-to-use form; keep bodies behavior-changing and free of activation guidance; avoid unrequested auxiliary files and default prompts; index only useful references with concise selection rules; verify the complete folder before finishing.
4. **Expected result:** A skill folder whose description routes appropriately, whose body gives concise post-load guidance, whose metadata and reference structure respect the stated boundaries, and whose contents contain only material support for the skill.
5. **Boundary:** Preserve the discipline as authoring guidance rather than a general skill generator, packaging manual, invocation workflow, or exhaustive template. Preserve the explicit exception allowing `interface.default_prompt` when the user requests it and the conditional use of references when extra material is genuinely needed.

## Evaluation Surfaces

- Description loop: proactive native selection by fresh independent Codex agents for representative skill creation, update, and review tasks.
- Body loop: isolated document and generated-artifact cases covering description/body discipline, metadata/folder boundaries, and reference routing decisions.
- References: not applicable because the target has no `references/` directory.

## Completion Criteria

- Four fresh description trials retain direct native evidence that the exact project-local target loaded before output.
- Every isolated body case satisfies the contract without a material retained failure.
- `agents/openai.yaml` remains free of `interface.default_prompt` unless a case explicitly requests one.
- No refinement adds fixture wording, exhaustive case lists, duplicate semantic owners, or unrelated scope.
- Any target edit is followed by clean reruns of affected evidence.

## Environment

- Qualification: `evaluations/skill-system-production-refinement/environment-qualification.md`
- Qualified surface: Codex Desktop fresh sub-agents backed by Codex CLI `0.145.0-alpha.18`, `gpt-5.6-sol`, high reasoning, `fork_turns: "none"`.
- Attribution risk: another installed plugin advertises a skill with the same suffix. Only an exact project-local injection or read of `supporting-skills/skill-authoring-discipline/SKILL.md` counts.

## Refinement Mode

Automatic semantic refinement is authorized by the user when retained evidence identifies a generalized failure at the target's semantic owner.
