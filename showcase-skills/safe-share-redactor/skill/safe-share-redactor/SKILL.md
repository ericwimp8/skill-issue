---
name: safe-share-redactor
description: Deterministic redaction workflow for preparing supplied text artifacts for safer sharing. Use only when the user explicitly asks to sanitize material with this skill.
---

# Safe Share Redactor

## Preserve the Source

- Work from a supplied UTF-8 text file. Do not edit, replace, or normalize the original.
- Create a new, empty output directory outside the supplied material. Keep its access no broader than the source while reviewing it.
- Do not paste suspected sensitive values into chat, commands, filenames, or explanatory prose.

## Run the Deterministic Redactor

Execute the bundled owner rather than recreating its matching rules:

```sh
python3 scripts/redact.py <input-file> --output-dir <new-output-directory>
```

Run the command from this skill directory, or use the corresponding absolute script path when the harness requires it. The script refuses existing output artifacts and writes:

- `sanitized.<type>`: the sanitized copy, using a supported non-sensitive text extension or `.txt`.
- `findings.json`: deterministic matches and ambiguous contextual risks without raw matched values.

If the command fails, preserve its error and stop. Do not bypass collision, encoding, or source-preservation checks with ad hoc copying or search-and-replace.

## Review Before Sharing

1. Verify the original file hash is unchanged from the pre-run hash recorded by the script.
2. Inspect every deterministic finding by rule and location in the sanitized copy. Confirm useful syntax and surrounding structure remain intact.
3. Inspect every ambiguous contextual risk in the source and sanitized copy. Decide whether it needs manual removal, generalization, or approval; the script deliberately leaves these cases unchanged.
4. Scan the complete sanitized copy for unsupported secrets, identifiers, personal information, proprietary context, encoded values, attachments, and cross-file clues.
5. Re-run the script from the original into a new empty directory after any source-side correction. Do not edit the generated findings into a false clean result.

Treat the sanitized copy as a review candidate. Automated redaction reduces supported risks but does not guarantee complete privacy, secrecy, anonymity, or safe disclosure.

## Report the Result

Provide the sanitized path, findings path, deterministic and ambiguous counts, unchanged-source verification, and unresolved review risks. Share the sanitized artifact only after the user or governing policy accepts the remaining contextual risk.
