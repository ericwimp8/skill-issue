# Ground Truth

- `format_invoice_amount` owns invoice amount presentation for both email and PDF paths.
- Upstream invoice amounts must remain numeric and signed; neither caller should compensate.
- The smallest complete change is local to `format_invoice_amount` and preserves positive formatting.
- Both email and PDF manifestations must render `-$12.50` for `-12.5`, while positive `12.5` remains `$12.50`.

