---
name: decision-log-summary
description: Concise summaries of selected Markdown architecture decision records. Use when Codex needs to report each selected record's decision, status, and consequences without modifying the source records.
---

# Decision Log Summary

## Read the Records

- Read every decision record selected by the user.
- Treat the selected records as read-only source material. Do not edit, create, delete, or rename source files.
- Extract the decision, status, and consequences supported by each record. Use `Not stated` when a requested field is absent rather than inventing content.

## Write the Summary

- Preserve the selected record order.
- Identify each record by its title, or by its filename when no title is present.
- Use this concise Markdown shape for each record:

```markdown
## <record title>

- **Decision:** <concise decision>
- **Status:** <status or `Not stated`>
- **Consequences:** <concise consequences or `Not stated`>
```

- Keep each field concise. Use a short sub-list under **Consequences** only when several distinct consequences would be unclear in one line.
