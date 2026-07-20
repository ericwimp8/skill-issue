# Behavior Cycle 1 Ground Truth

## Case 1 — Supported Deterministic Matches

- Execute the bundled script against `fixtures/behavior/case-1/input.log` in an isolated empty output directory.
- Preserve the source hash and surrounding log structure.
- Replace the synthetic API key, authorization value, email address, IP address, and user-home segment with their stable placeholders.
- Record five deterministic findings without retaining any matched value.

## Case 2 — Unchanged Material

- Execute the bundled script against `fixtures/behavior/case-2/input.txt` in a distinct empty output directory.
- Produce a byte-identical sanitized copy with zero deterministic and ambiguous findings.
- Report that review is not required by the supported rules while preserving the global limitation statement.

## Case 3 — Ambiguous Contextual Risk

- Execute the bundled script against `fixtures/behavior/case-3/input.txt` in a distinct empty output directory.
- Leave the permitted synthetic identity unchanged because free-form names are outside deterministic support.
- Report the contextual markers as ambiguous risk and require source-and-output review.
- Preserve the explicit statement that automated redaction cannot guarantee complete privacy or secrecy.

## Shared Result Characteristics

Each case is isolated from every other case, writes new artifacts only, leaves its source byte-identical, and satisfies the observable criteria in `contract.md` without using real credentials or non-permitted identities.
