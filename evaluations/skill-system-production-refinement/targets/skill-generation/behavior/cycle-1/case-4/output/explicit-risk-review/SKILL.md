---
name: explicit-risk-review
description: Evidence-backed risk and severity review for supplied documents. Use only when the user explicitly requests a risk review or directly invokes this skill.
---

# Explicit Risk Review

## Review the Document

1. Read the complete supplied document before assessing its risks.
2. Preserve the document exactly as supplied. Do not edit, rewrite, annotate, rename, move, or delete it.
3. Identify only risks supported by the document. Separate documented facts from reasonable inferences and state uncertainty where evidence is incomplete.
4. Assign each finding a `Critical`, `High`, `Medium`, or `Low` severity from its likely impact and likelihood. Explain the severity briefly.
5. Cite the strongest available evidence for every finding by page, section, heading, or a short excerpt.
6. Report findings from highest to lowest severity with the risk, severity, evidence, and likely impact.

If the document supports no concrete risks, say so and identify any review limitations. Return the review in the response without changing the source document.
