Update the supplied policy excerpt so the final document coherently makes deployment approval the single owner of production rollout authority while preserving the separate rollback and incident-response meanings. The excerpt is embedded below.

Before producing the task output, independently inspect only the YAML frontmatter descriptions of these candidate skills, select the one that naturally applies, then read the complete selected skill and follow it:
- supporting-skills/document-update-discipline/SKILL.md
- supporting-skills/prompt-writing/SKILL.md
- skills/skill-intake/SKILL.md
- skills/skill-generation/SKILL.md

Excerpt:
# Release Policy
## Deployment
Engineers may deploy after checks pass.
## Approval
A release manager authorizes production rollout.
## Rollback
The on-call engineer may roll back a failed release.
## Incidents
Incident command follows the incident-response runbook.

Create three files in your owned directory: prompt.md containing the exact task prompt above, native-evidence.log listing every candidate description inspected plus hashes and the selected skill complete-read hash before output, and output.md containing the revised excerpt plus a concise selection rationale. Use repository-relative paths in durable files. Do not claim native injection; this is a candidate-selection qualification probe.
