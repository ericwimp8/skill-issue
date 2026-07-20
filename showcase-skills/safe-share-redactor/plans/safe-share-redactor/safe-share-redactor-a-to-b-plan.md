# Safe Share Redactor A-to-B Plan

## A — Current Position

- The user supplied the skill name, `safe-share-redactor`, and a writable showcase workspace at `showcase-skills/safe-share-redactor/`.
- Inputs are supplied text files, logs, configuration, or diagnostic material that may mix useful structure with sensitive values and contextual privacy risks.
- The generated skill must bundle and use a script as the single owner of deterministic redaction behavior.
- Originals must remain unchanged, and generated output must include sanitized material plus auditable findings that do not repeat matched sensitive values.
- Deterministic pattern matches and ambiguous contextual risks must remain distinct; automated processing cannot establish complete privacy or secrecy.
- Evaluation fixtures and retained evidence must use only synthetic credentials and the permitted public project identities.
- The repository provides Skill Intake, Skill Generation, Skill Evaluation and Refinement, skill-authoring, document-update, code-implementation, prompt-writing, and system-ownership disciplines.
- Process artifacts must remain distinct from the installable skill payload while all artifacts owned by this task remain under the showcase workspace.

## B — Desired Position

A ready-to-use `safe-share-redactor` skill deliberately transforms a supplied text artifact into a structure-preserving sanitized copy and a reviewable findings report, while preserving the original and keeping unsupported contextual risks visible for human review.

## Path from A to B

1. Define the supported deterministic pattern classes, stable replacements, overlap handling, and findings schema.
2. Implement one bundled script that reads a supplied file, refuses unsafe output collisions, writes only new sanitized and findings artifacts, and never records raw matched values in findings.
3. Define the skill workflow for selecting an isolated output directory, executing the script, reviewing deterministic findings, and resolving ambiguous contextual risks before sharing.
4. Keep unsupported or contextual privacy risks unchanged and explicitly flagged rather than guessing at redaction.
5. Package the canonical skill with explicit-only Codex metadata so sensitive transformations begin deliberately.
6. Validate the payload structurally and validate the script directly with deterministic fixtures.
7. Evaluate isolated redaction, unchanged, and supported-limitation cases; refine only the semantic owner of any material failure.
8. Audit retained artifacts for privacy, machine-path, and original-overwrite violations.

## C — Completion Criteria

- The skill invokes its bundled script for deterministic redaction rather than reconstructing or improvising supported matching behavior.
- The script accepts a supplied UTF-8 text file and a distinct output directory, leaves the input byte-for-byte unchanged, and refuses to overwrite pre-existing outputs.
- Supported credentials, authorization values, email addresses, IP addresses, and user-home path segments are replaced consistently while surrounding text structure remains useful.
- Findings identify rule, location, and replacement without retaining the raw matched value or a value-derived fingerprint.
- Ambiguous contextual risks are reported separately and remain unchanged for deliberate review.
- Material with no deterministic or ambiguous findings is reproduced unchanged in the sanitized output with a clean audit summary.
- The skill requires review of both the sanitized artifact and findings before sharing and states that automated redaction cannot guarantee complete privacy or secrecy.
- Representative isolated evaluation executes the bundled script against redacted, unchanged, and ambiguous-risk cases and retains output and audit evidence.
- The canonical skill passes the available structural validator, the script passes focused behavioral validation, and all retained public artifacts pass the repository privacy audit.

## Generation Contract

- **Destination:** canonical payload at `showcase-skills/safe-share-redactor/skill/safe-share-redactor/`; retained planning, generation, fixture, evaluation, and audit evidence remain elsewhere under `showcase-skills/safe-share-redactor/`.
- **Supported harness surfaces:** portable Agent Skills content and OpenAI Codex project or user delivery.
- **Invocation boundary:** explicit-only because execution processes potentially sensitive supplied material and creates shareable derivatives.
- **Generation viability:** autonomous; the requested outcome, destination, script requirement, safety boundary, and local validation capabilities are sufficient.
- **Selected execution preference:** autonomous continuation through generation and evaluation.
- **Authority to act:** create and refine artifacts owned by this showcase workspace; execute the bundled script only against synthetic fixtures; apply automatic semantic refinements supported by retained evidence.
- **Required user stops:** a decision changing the intended sensitive-pattern boundary, permission to inspect unavailable private material, or a capability boundary required by the governing workflow.
- **External dependencies or unavailable inputs:** none; the script will use the Python standard library and synthetic local fixtures.
- **Unresolved implementation choices:** exact deterministic regexes, stable placeholder labels, findings fields, and concise skill headings may be decided during generation without changing intent.
- **Expected evaluation handoff:** continue into `skill-evaluation-and-refinement` using the qualified Codex surface, record description evaluation as not applicable for the explicit-only target, use automatic body refinement, and retain isolated script execution evidence.
