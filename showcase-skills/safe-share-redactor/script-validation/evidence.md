# Script Validation Evidence

- **Validator:** `showcase-skills/safe-share-redactor/script-validation/validate_redactor.py`.
- **Command:** `python3 showcase-skills/safe-share-redactor/script-validation/validate_redactor.py`.
- **Result:** `PASS: deterministic rules, ambiguity, source preservation, collisions, encoding`.
- **Rule coverage:** quoted secret, authorization, URL credentials, known token, JWT, email, documentation-range IP address, Unix and Windows user-home segments, and private-key material.
- **Safety coverage:** source hash preservation; omission of matched values, value-derived fingerprints, input names, and source paths from findings; constant-safe output naming; output-collision refusal without mutation; and invalid UTF-8 refusal.
- **Context coverage:** the permitted fixture identity remains unchanged and produces ambiguous-context review evidence.
- **Dependency boundary:** validation and production execution use only the Python standard library.
