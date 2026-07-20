# Behavior Cycle 2 Audit

## Execution

The three isolated cases passed their recorded ground truth after the output-identity refinement.

## Material Failure

The deterministic findings included a truncated unkeyed SHA-256 fingerprint derived from each raw matched value. Low-entropy values such as email addresses or common credentials could be guessed offline and compared with that fingerprint, so the findings were not as disclosure-minimizing as their purpose requires. Ambiguous markers were also suppressed for an entire line whenever any deterministic value on that line was redacted, which could hide unrelated contextual risk.

## Semantic Owner and Refinement

- **Failed meaning:** auditable findings must not retain raw values or unnecessary value-derived material, and deterministic redaction must not conceal unrelated contextual review needs.
- **Owner:** deterministic and ambiguous findings construction in `scripts/redact.py`.
- **Generalized update:** remove value-derived fingerprints and report contextual markers independently of deterministic matches on the same line.
- **Reconciliation:** update the intake criterion, evaluation criterion, focused validator, and evidence wording that previously required or described fingerprints.
- **Preserved meaning:** rule, location, replacement, source hashes, stable placeholders, collision refusal, and no-guarantee limitations remain intact.

## Decision

Cycle 2 failed materially and counts as the second unsuccessful body cycle. Generated outputs containing value-derived fingerprints are transient and must not be retained. Restart all three cases from original fixtures in cycle 3.
