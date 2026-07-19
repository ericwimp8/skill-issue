---
name: documentation-auditor
description: Evidence-backed, report-only auditing of Markdown documentation. Use when reviewing Markdown for contradictions, missing ownership, stale cross-references, or unsupported claims.
---

# Documentation Auditor

## Preserve the Source

- Treat every audited Markdown file as read-only. Do not edit, format, rename, move, or delete audited documentation.
- Produce findings and remediation guidance only. Leave implementation of every remediation to the user or a separate authorized task.

## Audit from Source

1. Read each in-scope document completely before concluding.
2. Follow its cross-references to accessible targets and distinguish authoritative sources from commentary or examples.
3. Check for:
   - contradictions between statements that cannot both govern the same situation;
   - responsibilities, decisions, or required actions without a clear owner;
   - cross-references whose target, anchor, path, or described content is missing or stale;
   - factual or normative claims without support from an identified authoritative source.
4. Base each finding on observable evidence. When evidence is unavailable, record the limitation instead of presenting an assumption as a finding.
5. Keep the absence of private production examples explicit in the report. Do not infer production behavior from unavailable examples.

## Report Findings

- Order findings by severity: `critical`, `high`, `medium`, then `low`.
- For every finding, provide:
  - **Severity:** the impact if the documentation is followed as written;
  - **Concern:** one of `contradiction`, `missing ownership`, `stale cross-reference`, or `unsupported claim`;
  - **Evidence:** each relevant source path and precise location, plus the conflicting, missing, stale, or unsupported content;
  - **Impact:** the concrete ambiguity, error, or governance risk;
  - **Remediation:** a specific documentation change and the semantic owner that should make it.
- Report the audit scope and limitations even when no findings are found. Do not manufacture findings to populate the report.
