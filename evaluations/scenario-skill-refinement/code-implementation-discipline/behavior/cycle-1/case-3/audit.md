# Case 3 Audit

- Fresh agent: `/root/eval_code_implementation/code_impl_body_c1_case3`.
- Session: `019f8274-491c-7731-bede-db454b930cdd`.
- Runtime: Codex Desktop / CLI `0.145.0-alpha.18`, `gpt-5.6-sol`, high reasoning.
- Target load: `native-trace.jsonl` records the exact canonical target read before fixture inspection and editing.
- Outcome: negative sign precedes the currency symbol across every invoice presentation path.
- Observation and path: email and PDF callers both consume `format_invoice_amount`; numeric signed amounts are created upstream and remain unchanged.
- Owner vs callers: the visible formatter genuinely owns presentation, so the correct change is local rather than moved away reflexively.
- Approach and fit: format the sign at the shared formatter and preserve positive formatting without caller changes.
- Reconciliation: both callers remain thin consumers and require no compensation.
- Verification: evaluator reran `python3 -m unittest -v test_invoices.py`; both manifestation and preservation checks passed.
- Ground-truth result: all recorded conditions satisfied.
- Case result: **pass**.
- Cleanup ownership: evaluator; retain completed fixture, output, trace, and audit.

