# Incident Timeline Builder A-to-B Plan

## A — Current Position

- The user supplied the stable name `incident-timeline-builder` and the writable showcase workspace `showcase-skills/incident-timeline-builder/`.
- Inputs may combine logs, alerts, deployment records, and operator notes with different timestamp formats, explicit offsets, named time zones, missing times, and conflicting accounts.
- Original sources must remain unchanged, and every material timeline claim must retain provenance back to a supplied source and locator.
- The repository provides the complete Skill Intake, Skill Generation, Skill Evaluation and Refinement, skill-authoring, document-update, prompt-writing, code-implementation, and system-ownership disciplines.
- The portable Agent Skills payload and OpenAI Codex metadata require no external service, secret, network call, or private input.
- Public artifacts must use synthetic or repository-owned evidence, repository-relative paths, and the repository privacy contract.

## B — Desired Position

A ready-to-use `incident-timeline-builder` skill turns heterogeneous incident evidence into a source-preserving chronological timeline that normalizes only explicit temporal facts, maintains epistemic distinctions, and exposes unresolved evidence and follow-up work.

## Path from A to B

1. Define the evidence inventory, provenance model, and immutable-source boundary.
2. Define timestamp extraction and normalization rules that preserve raw values, explicit time-zone context, ambiguity, precision, and missing-time states.
3. Define deterministic ordering for resolved timestamps and explicit treatment for unplaced or partially ordered evidence.
4. Define classifications for observations, reported statements, inferences, contradictions, and gaps without turning sequence or correlation into causation.
5. Define a concise incident deliverable containing scope, chronology, unresolved evidence, contradictions, and prioritized follow-up actions.
6. Package deterministic normalization and ordering as a local script because repeated timestamp parsing, daylight-saving ambiguity checks, stable ordering, and provenance reconstruction are error-prone.
7. Validate structure and script behavior, then evaluate representative multi-source incident cases in isolated fresh-agent trials and refine only the semantic owner of retained failures.

## C — Completion Criteria

- The skill inventories supplied evidence and preserves every original source without in-place edits.
- Each timeline item retains source identity, source path or equivalent identifier, locator, raw timestamp, and classification.
- Explicit offsets and named time zones normalize to UTC while preserving the raw timestamp and stated zone.
- Missing, invalid, ambiguous daylight-saving, and insufficient-precision timestamps remain explicit unresolved states and are never assigned invented times.
- Resolved events use deterministic chronological ordering with a documented stable tie rule; unresolved items remain useful without being falsely interleaved.
- Observed events, reported statements, inferred relationships, contradictions, and evidence gaps remain distinguishable in the output.
- Inferences identify their supporting events and confidence boundary; chronology or correlation alone is never stated as causation.
- Contradictions preserve all relevant accounts and identify the exact disagreement rather than selecting an unsupported winner.
- The final output contains scope and normalization notes, a chronological timeline, unplaced evidence, contradictions and gaps, unresolved evidence, and prioritized follow-up actions.
- The deterministic helper is directly tested for repeatability, offset and named-zone handling, daylight-saving ambiguity, missing timestamps, stable ties, input preservation, invalid input, output collision, and local privacy behavior.
- Structural validation and independent runtime evaluation either provide retained evidence for every criterion or record the exact governing stop.

## Generation Contract

- **Destination:** canonical payload at `showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/`; governed process and campaign artifacts remain elsewhere in this showcase workspace.
- **Supported harness surfaces:** portable Agent Skills content and OpenAI Codex project or user skill delivery.
- **Codex metadata:** required for the confirmed Codex target; implicit invocation remains enabled because incident chronology requests should select the skill naturally.
- **Generation viability:** autonomous; the requested outcome, evidence boundary, destination, local execution capability, and synthetic-fixture authority are sufficient.
- **Selected execution preference:** autonomous continuation through generation and evaluation with automatic semantic refinement.
- **Authority to act:** create and refine artifacts only inside this showcase workspace; run local deterministic checks; use fresh clean-context GPT-5.6 Sol medium agents for trials.
- **Required user stops:** a new decision that changes the intended skill, permission to access unavailable private evidence, or a governing evaluation capability boundary that cannot be satisfied locally.
- **External dependencies or unavailable inputs:** none for generation; real incident work remains bounded by the evidence and access supplied in that future task.
- **Unresolved implementation choices:** exact JSONL record field names, output headings, fixture incident domains, and concise helper error text.
- **Expected evaluation handoff:** qualify the specified medium-reasoning fresh-agent surface, complete the four-trial description pass, validate the script directly, and execute representative behavior cases with retained native evidence.
