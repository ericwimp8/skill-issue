# Readiness Assessment

**Status:** Not ready to ship.

The frontmatter description already follows the required concise what-it-is and when-to-use form, and the main instruction is brief and behavior-changing. The severity reference also contains material incident-brief guidance. The folder still has three release-blocking issues:

- `README.md` is unrequested installation boilerplate rather than material support for the skill.
- `agents/openai.yaml` defines `interface.default_prompt` without an explicit user request.
- `SKILL.md` tells the agent to read every reference unconditionally instead of indexing the available reference with a concise selection rule. This prevents case-based reference routing and does not identify what `severity.md` contains or when it is relevant.

## Smallest Coherent Correction

1. Remove `README.md`.
2. Remove only `interface.default_prompt` from `agents/openai.yaml`, preserving the display name and short description.
3. Replace the unconditional reference instruction with one concise index entry for `references/severity.md`, identifying it as incident-severity reporting guidance and routing to it only when the brief includes or evaluates a declared severity level.

These changes preserve the sound description, core instruction, useful metadata, and material severity guidance while making every retained file and reference route release-ready.
