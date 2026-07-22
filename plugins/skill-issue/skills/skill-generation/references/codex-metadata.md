# Codex Metadata

Keep the canonical skill portable. Put Codex-only discovery and invocation configuration in `agents/openai.yaml`.

Generate this file only when Codex is a confirmed target and the intake contract requires Codex metadata. Use only the fields needed by that contract:

```yaml
interface:
  display_name: "Readable Skill Name"
  short_description: "Concise skill summary"

policy:
  allow_implicit_invocation: false
```

- Keep `interface.display_name` readable and aligned with the skill name.
- Keep `interface.short_description` concise and aligned with the `SKILL.md` purpose.
- Set `policy.allow_implicit_invocation` to `false` when the skill must be invoked explicitly.
- Omit `policy` when the intake contract does not require an invocation override.
- Omit `interface.default_prompt` unless the user explicitly requests one.

Do not put Codex metadata in `SKILL.md`, and do not infer equivalent metadata for another harness.
