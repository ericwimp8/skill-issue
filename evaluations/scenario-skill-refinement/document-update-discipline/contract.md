# Document Update Discipline Evaluation Contract

## Target

- Canonical skill: `evaluations/scenario-skill-refinement/document-update-discipline/skill/`

## Goal

Prevent document changes from following the nearest visible mention when another location owns the meaning, while leaving the whole document coherent.

## Intended Use

Planning or performing updates to documents whose requested meaning may appear at an observation point, a semantic owner, related manifestations, or conflicting text.

## Expected Behavior

The agent maps the complete document before editing, distinguishes the observation point from the semantic home, identifies related manifestations and preserved meaning, chooses the required update operation, edits the semantic owner, reconciles related text, and verifies the document as a whole.

## Expected Result

The updated document or update plan places the complete meaning at its semantic owner, removes conflicts and accidental duplicate ownership, preserves unrelated meaning, and reads as one intentional document.

## Boundary

The evaluation must preserve the target's applicability to document planning and editing, its whole-document reading requirement, its allowance for a local edit when the local site genuinely owns the idea, its semantic-preservation requirements, and its user-controlled limited-scope exception.

## Observable Completion Criteria

1. The result identifies or demonstrates the required meaning, document purpose, observation point, semantic home, related manifestations, operation, and preserved unrelated meaning.
2. The complete normative meaning lives at the semantic owner rather than being appended only at the visible mention.
3. Related manifestations agree with the owner without becoming accidental competing owners.
4. Conflicts are resolved and unrelated meaning remains materially intact.
5. Planning-only cases do not edit their fixtures.
6. Edited documents match their existing tone and structure and state current truth directly.

## Evaluation Surface

- Description loop: natural document-update planning and editing tasks with direct native target-load evidence.
- Reference qualification: not applicable because the canonical target has no `references/` directory.
- Body loop: isolated Markdown document edits and a planning-only response, audited against condition-based ground truth.
