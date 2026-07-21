# Case 3 Output

- `format_invoice_amount` produces `-$12.50` for negative amounts and preserves `$12.50` for positive amounts.
- The formatter owns shared presentation for both email and PDF invoice paths.
- Verification: `python3 -m unittest -v test_invoices.py` passed two tests.

