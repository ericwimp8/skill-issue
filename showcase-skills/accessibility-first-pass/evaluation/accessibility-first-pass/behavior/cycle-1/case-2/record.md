# Behavior Cycle 1 Case 2 Record

- **Publication-safe prompt:** Read and follow the complete candidate and its resources, then review the dialog fixture through source inspection, rendered interaction, available project tooling, and current authoritative guidance, including opening, traversal, dismissal, focus restoration, and destructive confirmation.
- **Prompt preservation:** Native session `019f7eef-b470-79f2-85af-17bce081c021` retains the unmodified zero-context prompt; this record omits its machine-specific checkout path.
- **Fresh agent:** `a11y_body_dialog`; Codex CLI `0.145.0-alpha.18`; `gpt-5.6-sol` with high reasoning.
- **Target evidence:** The native pre-output trace read the exact candidate at `2026-07-20T09:51:24.959Z`.
- **Output:** `output.md`.
- **Ground-truth comparison:** Pass. The report directly exercises focus placement, forward and reverse traversal, Escape, Cancel, restoration, destructive activation, and semantic exposure; identifies the missing operation, modal lifecycle, focus visibility, and accessible name; and separates browser, source, scanner, inference, and unverified assistive-technology evidence.
- **Criteria:** All nine observable criteria pass.
- **Cleanup:** The agent closed its local browser tab and server; no fixture change was retained.
