# Workflow Observations

## Unmodified Prompt Retention and Repository Privacy

The evaluation campaign contract requires each trial record to retain the unmodified prompt. Fresh zero-context agents require a concrete checkout path, while the repository privacy contract prohibits durable machine-specific checkout paths.

This campaign preserves each unmodified prompt in its native Codex session and retains a publication-safe prompt copy plus the native session identifier in the repository. The evaluation workflow does not currently define this privacy-normalized evidence shape or state whether the private native session alone satisfies durable prompt retention. A later workflow review should establish one explicit owner for publication-safe prompt evidence without weakening reproducibility or privacy.
