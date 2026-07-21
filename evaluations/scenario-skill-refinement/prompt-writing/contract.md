# Prompt Writing Evaluation Contract

## Target

- Canonical skill: `evaluations/scenario-skill-refinement/prompt-writing/skill/`
- `SKILL.md` SHA-256: `c6b7ff268497e3174cdf572ad6b1902d447c00ca4d5f02e43157a89fabe3e05f`
- `agents/openai.yaml` SHA-256: `44e0103c0062d1026d7a2ba66e7401d220d59418097b6d06848b87b616a7caf0`
- Environment qualification: `evaluations/skill-system-production-refinement/environment-qualification.md`

## Interpreted Contract

1. **Goal:** Produce prompts that communicate the task-specific delta an agent needs while relying on its existing context and named authoritative sources.
2. **Intended use:** Writing or revising prompts for agents, sub-agents, invoked skills, and authoritative plans.
3. **Expected behavior:** Include task-owned goal, completion criteria, and autonomy boundary when needed; read and point to named authority rather than restating it; request the smallest useful deliverable; frame discovery by category; remove redundant context; and repair the failed decision rather than accumulating broad instructions.
4. **Expected result:** A concise, actionable prompt whose recipient can complete the requested task without duplicated authority, invented structure, unnecessary prescription, or an ambiguous action boundary.
5. **Boundary:** Preserve the task's actual scope and source authority. Do not make every prompt follow a fixed template, omit task-critical context, or turn prompt writing into execution of the prompted task.

## Observable Completion Criteria

- The prompt states the task-specific goal or requested outcome.
- The prompt includes only completion criteria and autonomy boundaries needed for the task.
- Named source material is referenced as authority and is not redundantly retaught.
- The deliverable is the smallest artifact that enables the next action.
- Open-ended investigation asks for a category, not a seeded inventory of expected findings.
- Wording avoids unnecessary workflow prescription and reusable-agent boilerplate.
- The produced prompt is internally coherent and directly usable.

## Evaluation Surfaces

- Description loop: native proactive selection on four fresh-agent prompt-writing tasks.
- Reference loop: not applicable because the target contains no `references/` directory or files.
- Body loop: isolated generated-artifact cases, audited as prompt text against task fixtures and this contract.
