The canonical artifact is `explicit-risk-review/SKILL.md` at the contracted destination. It owns the portable skill description and the complete read-only review behavior: report evidence-backed risks with severity while leaving the supplied document unchanged.

Platform metadata remains outside the canonical body. Codex owns its explicit-only policy in `agents/openai.yaml`; Claude Code and Pi own their supported explicit-only overlays in the supporting platform-metadata area; OpenCode relies on the canonical explicit-only description and documented standard frontmatter.

Authority is limited to the contracted destination and its supporting platform-metadata subdirectory. Source documents, external installation surfaces, and runtime behavior are outside this generation write boundary.

The first generation action is to create the canonical skill directory and draft `SKILL.md`, establishing the read-only behavior and portable output contract before adding host-specific metadata.
