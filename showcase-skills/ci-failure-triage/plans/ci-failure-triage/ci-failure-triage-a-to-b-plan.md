# CI Failure Triage A-to-B Plan

## A — Current Position

- The requested skill name is `ci-failure-triage`, with a writable isolated workspace at `showcase-skills/ci-failure-triage/`.
- Inputs may include partial CI logs, workflow configuration, repository source, job metadata, dependency manifests, and locally runnable tools; any subset may be absent or stale.
- CI output commonly contains repeated, interleaved, or cascading failures whose chronology alone does not establish causality.
- The skill must preserve uncertainty and unavailable evidence, and must not change remote CI state, secrets, branches, releases, or other external systems without explicit authorization.
- Synthetic and repository-owned fixtures are available for evaluation without authenticated external CI access.
- The repository provides qualified fresh-agent evaluation plus the Skill Intake, Skill Generation, Skill Evaluation and Refinement, skill-authoring, document-update, prompt-writing, and system-ownership disciplines.
- Process artifacts must remain outside the installable payload while staying within the assigned showcase workspace.

## B — Desired Position

A ready-to-use `ci-failure-triage` skill guides an agent from noisy or incomplete CI evidence to a reproducible, evidence-backed diagnosis that identifies the primary failure when supported, separates cascading symptoms, recommends the smallest responsible remediation direction, and gives an exact verification plan without unauthorized remote changes.

## Path from A to B

1. Establish an evidence model that separates observed facts, source-backed contracts, inferences, hypotheses, and unavailable evidence.
2. Define how to reconstruct job and step order across logs and workflow configuration before assigning causality.
3. Define how to identify the earliest causally sufficient failure and classify later output as primary, contributing, independent, or cascading.
4. Trace suspected ownership through workflow configuration and production source to the concrete failing behavior rather than stopping at wrappers or test expectations.
5. Define the smallest responsible remediation direction at the semantic and code owner while preserving unrelated behavior.
6. Define an exact verification plan that reproduces the failing boundary locally when possible and names the authoritative CI rerun needed for confirmation.
7. Define safe handling for unavailable tooling, missing logs, secrets, remote state, and authorization boundaries.
8. Package and structurally validate the concise skill, then evaluate representative trigger and behavior cases in isolation and refine only evidence-backed semantic failures.

## C — Completion Criteria

- The skill inventories supplied evidence and records provenance before diagnosing.
- The skill reconstructs relevant workflow, job, and step order and distinguishes execution chronology from causal order.
- The skill identifies a primary failure only when evidence establishes a causally sufficient upstream fault; otherwise it preserves ranked hypotheses or an unresolved result.
- The skill classifies consequential noise separately from independent or contributing failures and cites evidence for each classification.
- The skill traces the suspected failure through workflow configuration and repository production source to the smallest responsible owner, using tests only for reproduction or validation.
- The skill recommends a bounded remediation direction rather than speculative patches or broad cleanup.
- The skill provides exact local verification commands or procedures, expected observations, and the authoritative CI confirmation step, clearly marking unavailable execution.
- The skill records missing evidence, its diagnostic impact, and the smallest action that would resolve each material uncertainty.
- The skill requires explicit authorization before changing remote CI state, secrets, branches, releases, or other external systems.
- The final report contains an evidence inventory, failure sequence, primary diagnosis, cascade classification, remediation direction, verification plan, uncertainties, and authorization boundary.
- The canonical skill passes structural validation, and runtime evaluation supplies retained evidence or an exact governed stopping gate for every runtime criterion.

## Generation Contract

- **Destination:** canonical payload at `showcase-skills/ci-failure-triage/skill/ci-failure-triage/`; all process and campaign artifacts remain elsewhere under `showcase-skills/ci-failure-triage/`.
- **Supported harness surfaces:** portable Agent Skills content and OpenAI Codex project or user skill delivery.
- **Generation viability:** autonomous; the outcome, safety boundary, destination, local evidence sources, and evaluation surface are sufficient.
- **Selected execution preference:** autonomous continuation through generation and evaluation.
- **Authority to act:** create and refine only artifacts inside `showcase-skills/ci-failure-triage/`; apply automatic semantic refinements supported by retained campaign evidence.
- **Required user stops:** a decision that changes intended skill meaning, access to unavailable private evidence, or an authorization or capability boundary that the governing workflow requires the user to resolve.
- **External dependencies or unavailable inputs:** no authenticated CI service is required for generation or synthetic evaluation; future investigations may remain limited by missing logs, secrets, platform access, or unavailable local runtimes.
- **Unresolved implementation choices:** concise report headings, representative fixture domains, and whether any bundled resource is genuinely required.
- **Expected evaluation handoff:** continue into `skill-evaluation-and-refinement` using the qualified Codex fresh-agent surface, GPT-5.6 Sol at medium reasoning, automatic body refinement, isolated fixtures, and retained direct target-read evidence.
