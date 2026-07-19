---
name: repository-owner-finder
description: Read-only repository owner lookup from the workspace's authoritative CODEOWNERS file. Use when Codex must identify the owner of a supplied path in the local project.
---

# Repository Owner Finder

## Resolve the Owner

1. Pass the supplied repository path to `scripts/find_owner.py`:

   ```bash
   python3 <skill-directory>/scripts/find_owner.py '<repository-path>'
   ```

2. Treat the script's `Owner`, `Pattern`, and `Source` fields as the complete evidence for a successful result. Report them concisely.
3. If the script exits unsuccessfully, report that the owner cannot be determined and include its error. Do not infer, substitute, or invent an owner.
4. Keep the lookup read-only. Do not edit `CODEOWNERS`, `OWNER_POLICY.md`, or any other project file.

The script resolves the workspace's live `project/CODEOWNERS` from the skill location. Do not replace that source with copied rules or remembered values.
