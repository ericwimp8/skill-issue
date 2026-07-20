# Behavior Case 1 Record

- **Unmodified prompt:** `request.md`
- **Fresh agent:** `/root/ci_failure_triage/ci_body_1`, session `019f80c1-472c-7ba1-92a4-2ed65222b5e2`
- **Requested environment:** `gpt-5.6-sol`, medium reasoning, `fork_turns: "none"`
- **Fixture:** `fixtures/behavior/case-1/`
- **Direct target-load evidence:** native trace read the exact target before the request in the pre-output command at `2026-07-20T18:19:47Z`.
- **Output:** `report.md`
- **Ground-truth comparison:** correctly traces the prefixed version parsing, classifies skipped build/archive and absent artifact as cascade, preserves the unknown version-format contract, and reports executed versus planned verification separately.
- **Result:** pass.
- **Cleanup owner:** this campaign; fixture stayed unchanged and no remote state was accessed.
