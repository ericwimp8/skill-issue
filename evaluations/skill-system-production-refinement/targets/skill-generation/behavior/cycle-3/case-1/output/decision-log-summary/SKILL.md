---
name: decision-log-summary
description: Read-only summaries of decisions, statuses, and consequences from selected Markdown architecture decision records. Use when a user asks to summarize specific Markdown ADRs or decision logs.
---

# Decision Log Summary

## Summarize Selected ADRs

1. Identify the Markdown ADRs the user selected. Ask for the selection when it is missing or ambiguous.
2. Read each selected source without modifying it.
3. Process the ADRs in the user's selection order, or their source order when the selection comes from one document.
4. Report each ADR separately with its title or source identifier, followed by `Decision`, `Status`, and `Consequences`.
5. Summarize only explicitly stated source content. Write `Not stated` for any missing or empty field.
6. Preserve the source order of consequence items and distinguish favorable, unfavorable, and neutral consequences only when the source does.

Keep the result concise. Never edit, reformat, rename, or otherwise write to the source ADRs.
