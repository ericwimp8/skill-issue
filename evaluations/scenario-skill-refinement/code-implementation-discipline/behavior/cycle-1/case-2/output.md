# Case 2 Output

## Implementation Decision

- Outcome: every created product stores and returns a code with surrounding whitespace trimmed.
- Observation: the HTTP route trims; the import adapter passes raw input.
- Paths: HTTP `request["code"]` → trim → `create_product` → product/list; import `row["code"]` → `create_product` → product/list.
- Owner: `create_product`, which creates and stores the canonical product representation.
- Affected callers: `create_product_route`, `import_product`, and direct `create_product` callers.
- Approach: trim inside `create_product`, remove `.strip()` from the route, and keep the import adapter as a relay.
- Fit check: central normalization follows creation ownership and removes route-only policy.
- Verification: cover direct, HTTP, and import creation with padded inputs and assert returned and stored codes are identical.
- No files were edited.

