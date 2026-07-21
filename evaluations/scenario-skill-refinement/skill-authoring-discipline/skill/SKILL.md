---
name: skill-authoring-discipline
description: Codex skill authoring constraints and instructions. Use any time creating, updating, or reviewing Codex skills.
---

# Skill Authoring Discipline

Create skills that help an agent decide and act. Keep them concise, direct, and useful after the triggering decision has already been made.

## Core Rules

- Treat a skill as a gap filler, not a full operating manual.
- Include only instructions, tools, and references the agent genuinely needs.
- Prefer strong general constraints over long lists of tiny prescribed details.
- Prescribe a small detail only when missing that detail would likely break the work.
- Remove boilerplate, rationale history, and generic explanations.
- Do not add auxiliary files such as README, changelog, quick-reference, or installation notes unless the user explicitly asks.

Long prescriptive lists are risky: once a skill lists many small details, the agent may optimize around that list and miss unlisted details. State the governing rule clearly enough that the agent can reason through the particulars.

## Description Format

Write every skill `description` as one concise what-it-is sentence followed by one concise when-to-use sentence. Do not add examples, edge cases, implementation details, or long trigger lists.

Good:

```yaml
description: Flutter UI layout constraints and instructions. Use any time working with, editing, or understanding Flutter layouts.
```

Bad:

```yaml
description: A Flutter skill that helps with layouts, constraints, widths, heights, expanded values, layouts that need constraining or loose constraints. Use when dealing with columns, rows, expanded flex values, constraints in the UI where the UI needs to constrain or have loose constraints on layouts.
```

## Skill Body

The body must tell the agent what changes its behavior after the skill is loaded.

Use:

- Short imperative rules.
- Concrete workflow gates only when order matters.
- Script paths and exact commands when deterministic execution matters.
- Reference indexes when extra material should be loaded only for matching cases.

Mandatory body gates:

- Do not repeat or restate the description in the skill body.
- Do not include trigger conditions, "Use this skill...", "When to use...", or similar activation guidance in the skill body. If trigger wording is needed, keep it in the frontmatter `description` or external discovery metadata.
- Do not explain obvious model capabilities.
- Do not add exhaustive checklists that try to enumerate every possible case.
- Do not add showcase examples that are longer than the rule they illustrate.
- If any body text violates these gates, revise the skill before treating the update as complete.

## OpenAI Metadata

Do not add `interface.default_prompt` to `agents/openai.yaml` unless the user explicitly requests one.

## Reference Documents

When a skill has reference documents, index them in `SKILL.md` so the agent can choose the right one quickly.

Use this structure:

```markdown
## Reference Documents

Use the relevant reference document when needed from this skill.

- `references/flutter-layout.md`: Flutter layout constraints and examples. Use when working with Flutter layout behavior, sizing, flex, scrolling, or constraint errors.
- `references/material-components.md`: Material component usage rules. Use when choosing, adapting, or replacing Material components.
- `references/test-patterns.md`: Flutter widget test patterns. Use when writing or reviewing widget tests for this component family.
```

Reference entry rules:

- Use the same what-it-is and when-to-use format as skill descriptions.
- Keep each entry to one concise line unless a longer note is truly required.
- Prefer scan-friendly entries over paragraphs.
- Do not summarize the whole reference document in the index.
- Do not make the agent read every reference by default.

## Final Check

Before finishing a skill, verify:

- The description is short and follows what-it-is plus when-to-use.
- The body contains only behavior-changing guidance and passes the mandatory body gates.
- `agents/openai.yaml` does not include `interface.default_prompt` unless requested.
- Reference documents, if present, are indexed with concise selection rules.
- The skill does not over-prescribe details the agent can reason through.
- The skill folder contains only files that support the skill.
