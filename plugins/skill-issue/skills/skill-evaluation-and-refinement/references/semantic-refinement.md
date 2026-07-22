# Semantic Refinement Constraints

## Governing Rule

Change the meaning that owns the failure. Do not append a nearby instruction merely because the failed case made that location visible.

## Required Constraints

- **Documents:** map the complete skill, locate the semantic owner, reconcile related meanings, and preserve unrelated meaning.
- **Code and executable fixtures:** trace behavior to its concrete owner, change the smallest complete implementation there, and reconcile affected dependants.
- **System structure:** place responsibilities with their normal successful owner rather than a convenient adapter, prompt, or cleanup step.
- **Skill authoring:** keep descriptions concise, keep body instructions behavior-changing, and move conditional platform detail into indexed references.

Apply the installed `document-update-discipline`, `code-implementation-discipline`, `system-change-ownership`, and `skill-authoring-discipline` skills when available. These inline constraints remain the minimum contract when one is unavailable.

## Reject

- Evaluation fixture names, literal expected answers, or instructions tailored to one failed example.
- A growing sequence of exceptions that leaves the original rule unchanged.
- Duplicate or conflicting semantic owners.
- New detail whose only purpose is to make the current evaluation pass.
- A chat proposal produced without the same constraints that would govern an automatic edit.

## Form the Update

State the failed contract meaning, the evidence, the semantic owner, the generalized change, related meanings requiring reconciliation, and meanings that must remain intact. Rewrite the smallest coherent section that makes the complete skill correct.
