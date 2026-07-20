# Script Validation Evidence

- **Validator:** `showcase-skills/incident-timeline-builder/script-validation/validate_timeline.py`.
- **Command:** `python3 showcase-skills/incident-timeline-builder/script-validation/validate_timeline.py`.
- **Result:** `PASS: determinism, zones, ambiguity, missing time, provenance, errors, preservation, privacy`.
- **Determinism:** two independent stdout runs are byte-identical; file output is byte-identical to stdout; exact timestamp ties retain input order.
- **Time zones:** explicit numeric offsets, `Z`, and `Australia/Adelaide` normalize to UTC; the repeated `America/New_York` local hour remains `ambiguous`; a naive timestamp without a named zone remains `missing-timezone`.
- **Missing evidence:** a null raw timestamp remains `missing` and no normalized instant is emitted.
- **Source preservation:** the input SHA-256 is retained; all source identifiers, repository-relative paths, locators, raw timestamps, classifications, and summaries survive transformation; the input fixture remains byte-identical.
- **Error behavior:** invalid JSON, input/output identity, and existing-output collision return status 2; collision refusal leaves the existing output unchanged.
- **Privacy:** the helper is local and standard-library-only, imports no network-capable client, does not add host paths or identities, and emits only supplied record data plus deterministic derived fields. The validator rejects host-identity leakage in the synthetic result.
- **Fixture correction:** the first validator run correctly exposed a mistaken expected Adelaide UTC value in the validation ground truth. The expectation changed from `15:55Z` to `14:55Z`; production script behavior did not change.
