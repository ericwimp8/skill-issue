# Environment Doctor A-to-B Plan

## A — Current Position

- The user supplied the stable skill name `environment-doctor` and the writable showcase workspace `showcase-skills/environment-doctor/`.
- The diagnostic target is a supplied development root observed through the script process environment, an explicitly selected tool set, selected environment-variable names, requested path-order relationships, and selected version declaration files.
- The skill must bundle a script as the single owner of deterministic inspection and findings generation.
- Diagnostics must be read-only: no installation, configuration edit, environment mutation, shell startup, or source-tree write is permitted.
- Retained output must be human-readable and machine-readable, preserve unavailable and unparseable states, and avoid exposing selected environment-variable values or unredacted home and supplied-root paths.
- Version probing can safely support only a bounded registry of tools with fixed non-mutating version arguments; unknown tools can be resolved without execution.
- PATH resolution behavior is platform-dependent. The initial supported execution surface is Python 3 on POSIX systems using colon-separated PATH semantics; claims beyond directly tested behavior remain unsupported.
- Evaluation fixtures and retained evidence must use synthetic tools, configurations, paths, and environment values.
- The repository supplies current intake, generation, evaluation, skill-authoring, document-update, code-implementation, prompt-writing, and system-ownership authorities.
- Process artifacts must remain separate from the installable payload while every task-owned artifact remains under this showcase workspace.

## B — Desired Position

A ready-to-use `environment-doctor` skill runs a bounded read-only diagnostic against a supplied development root and selected environment surfaces, then returns actionable findings and privacy-safe structured evidence without concealing unknowns or changing the inspected environment.

## Path from A to B

1. Define a narrow command interface for selecting the root, output location, tools, environment-variable names, PATH ordering expectations, and tool version declaration files.
2. Implement a standard-library Python script that inventories selected executables, probes only fixed registered version commands, classifies selected environment state, evaluates requested PATH order, and compares parseable tool and declaration versions.
3. Normalize supplied-root and home paths, omit selected non-PATH environment values, bound subprocess output, and preserve unavailable, unsupported, timeout, and unparseable states.
4. Generate stable human-readable and JSON evidence with finding identifiers, severity, evidence, safe remediation, and verification guidance.
5. Define the concise skill workflow for selecting inspection scope, executing the owner, interpreting evidence, preserving uncertainty, and obtaining consent before any later remediation.
6. Package OpenAI metadata for ordinary proactive invocation and avoid unnecessary references or assets.
7. Validate structure and directly test determinism, read-only behavior, source preservation, collisions, failures, path privacy, version mismatch handling, and unsupported-tool behavior with synthetic fixtures.
8. Qualify the requested fresh-agent surface, run the four-trial description protocol and isolated behavior cases, and refine only the semantic owner of retained material failures.
9. Audit retained artifacts for privacy, machine paths, scope, hashes, diffs, and unsupported platform claims.

## C — Completion Criteria

- The skill directs the agent to execute the bundled script rather than reconstructing diagnostic behavior.
- The script accepts an existing development root and a new output directory, writes only `report.txt` and `evidence.json` under that output directory, and leaves the root and process environment unchanged.
- Selected tools report all PATH candidates, selected resolution, executable state, and a version state; only registered tools execute fixed version arguments, while unsupported tools preserve an `unsupported` version state.
- Selected environment variables report `unset`, `empty`, or `set` without exposing non-PATH values; PATH and path-like evidence normalizes the supplied root and home directory.
- Requested PATH-order relationships report satisfied, missing, duplicate, or reversed states without silently choosing an ambiguous result.
- Selected version declaration files report missing, unavailable-tool, unparseable, match, or mismatch states and never read outside the supplied root.
- Human and JSON outputs use stable ordering and contain equivalent finding identifiers, severities, evidence, remediation, and verification guidance.
- Failures, timeouts, unsupported platforms, invalid selections, and output collisions stop safely or remain explicit; the script never installs, edits, sources, or repairs anything.
- Representative direct validation proves deterministic repeated output, unchanged fixture hashes and environment, privacy normalization, collision refusal, and correct success and finding exit codes.
- Four fresh medium-reasoning description trials proactively load the skill, and isolated body trials satisfy every observable criterion with retained direct evidence.
- The canonical skill passes structural validation and all retained public artifacts pass format, diff, hash, and repository-privacy checks.

## Generation Contract

- **Destination:** canonical payload at `showcase-skills/environment-doctor/skill/environment-doctor/`; governed planning, generation, validation, evaluation, refinement, fixture, evidence, and audit artifacts remain elsewhere under `showcase-skills/environment-doctor/`.
- **Supported harness surfaces:** portable Agent Skills content and OpenAI Codex project or user delivery; executable behavior requires Python 3 on a POSIX host.
- **Invocation boundary:** ordinary implicit or explicit use for diagnosing a supplied local development environment; any later environment modification requires a separate user decision.
- **Generation viability:** autonomous; the outcome, destination, deterministic-script requirement, safety boundary, and local synthetic validation capabilities are sufficient.
- **Selected execution preference:** autonomous continuation through generation and evaluation.
- **Authority to act:** create and refine only task-owned showcase artifacts; execute the script only against synthetic or repository-owned fixtures; run fresh independent evaluation trials; apply evidence-supported automatic semantic refinements.
- **Required user stops:** a decision that changes the diagnostic boundary, permission to inspect unavailable private material, permission to alter an environment, or a capability boundary required by the governing workflow.
- **External dependencies or unavailable inputs:** none for generation; Python 3 and fresh Codex agents are locally available. No external authenticated runner is required.
- **Unresolved implementation choices:** exact CLI spellings, fixed version-probe registry, evidence schema details, normalized external-path labels, and concise skill headings may be decided during generation without changing intent.
- **Expected evaluation handoff:** continue into `skill-evaluation-and-refinement`, qualify the exact `gpt-5.6-sol` medium-reasoning fresh-agent surface, use automatic body refinement, and retain isolated direct script and agent evidence.
