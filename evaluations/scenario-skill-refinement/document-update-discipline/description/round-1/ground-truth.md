# Description Round 1 Ground Truth

The four prompts are natural document planning or editing requests. None names a skill, asks for invocation, quotes the target description, or reveals expected selection.

Description selection passes only when every fresh agent has direct native pre-output evidence that it opened `supporting-skills/document-update-discipline/SKILL.md` or received a native injection event for `skill-issue:document-update-discipline`.

The generated artifact is retained as supporting context, but answer similarity, artifact quality, final prose claims, and a `Skills Used` list do not establish invocation.
