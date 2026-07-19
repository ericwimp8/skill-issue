---
name: decision-log-summary
description: Read-only summaries of selected Markdown architecture decision records (ADRs), reporting each decision, status, and consequences in source order and marking absent fields as Not stated. Use when Codex needs to review or summarize a specified set of Markdown ADRs or decision logs without changing the sources.
---

# Decision Log Summary

## Establish the Selection

- Use only the Markdown ADR files explicitly selected by the user.
- Ask for the selection when no files are identified or the requested set is ambiguous.
- Preserve the selected source order. Do not reorder files by title, date, status, or content.

## Extract the Record

1. Read each selected file without writing to it.
2. Identify content explicitly labeled as the decision, status, and consequences. Accept headings, frontmatter keys, or labeled fields only when their meaning is unambiguous.
3. Summarize the stated content faithfully. Preserve the order of consequence paragraphs and list items.
4. Write exactly `Not stated` for every requested field that is absent or empty. Do not infer missing values from context.
5. Report a read failure as an error for that source rather than treating its fields as missing.

## Report the Summary

- Produce one clearly identified entry per ADR in the selected source order.
- Report `Decision`, `Status`, and `Consequences` for every readable ADR.
- Keep conclusions traceable to the selected source and distinguish source statements from any necessary qualification.

## Preserve Sources

- Do not edit, rename, move, reformat, or create replacements for the selected ADRs.
- Use read-only inspection commands and return the summary in the response.
