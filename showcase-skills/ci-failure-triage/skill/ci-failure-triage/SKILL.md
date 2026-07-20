---
name: ci-failure-triage
description: Evidence-first diagnosis of failed continuous-integration runs that separates primary faults from cascading noise. Use when CI logs, workflow configuration, repository source, or local tooling must be correlated to explain a failure and plan its smallest responsible fix.
---

# CI Failure Triage

## Establish the Investigation Boundary

- Inventory the supplied logs, job metadata, workflow configuration, revision, repository source, dependency state, and available local tools. Record the provenance and freshness of each material input.
- Keep direct observations, source-backed contracts, evidence-supported inferences, hypotheses, and unknowns distinct. Preserve truncated, unavailable, redacted, or stale evidence as explicit limitations.
- Treat production source and the concrete workflow definition as behavioral authority. Use tests to reproduce or validate after tracing the relevant production path.
- Work read-only against remote CI unless the user explicitly authorizes a specific mutation. Do not rerun, cancel, dispatch, approve, retry, edit secrets, push branches, tag releases, or publish artifacts on implied authority.

## Reconstruct the Failure Sequence

1. Map workflow triggers, job dependencies, conditions, matrices, services, caches, artifacts, and step order from configuration before interpreting log order.
2. Normalize timestamps and runner boundaries when available. Separate concurrent job output and repeated summaries from the first occurrence of each distinct failure.
3. Build a short causal sequence from prerequisite setup through the first broken invariant to downstream symptoms. Include skipped, cancelled, timed-out, and never-started work when it changes interpretation.
4. Verify log assertions against the invoked command, script, configuration, and concrete production implementation. Follow wrappers until reaching the code or external boundary that owns the observed effect.

Do not equate the earliest timestamp, loudest stack trace, final exit code, or largest group of failed tests with the primary cause without a causal link.

## Classify the Failures

- Mark a failure **primary** only when evidence shows it broke a required invariant and sufficiently explains the dependent symptoms.
- Mark a failure **contributing** when it materially worsened or exposed the run but does not alone explain the outcome.
- Mark a failure **independent** when it requires separate ownership or remediation.
- Mark output **cascading** when it follows from an established upstream failure, such as missing generated files, unavailable services, aborted setup, or a shared invalid state.
- Mark output **noise** only when evidence establishes that it neither changes the run result nor indicates a separate defect.

For every classification, cite the supporting log location and owning configuration or source path. When evidence cannot distinguish candidates, retain ranked hypotheses with supporting and disconfirming evidence instead of choosing a primary failure.

## Diagnose at the Responsible Owner

- State the failed invariant, the concrete observation, and the causal link separately.
- Trace the invariant to its semantic owner: workflow orchestration for ordering and environment, build or dependency configuration for toolchain resolution, and production source for application behavior.
- Check revision, platform, environment, dependency, cache, permission, service, timing, and nondeterminism explanations only when the available evidence makes them plausible.
- Prefer the smallest diagnosis that explains all classified dependent symptoms. Record residual observations that it does not explain.
- If access or tooling is unavailable, state which conclusion remains unverified and the smallest evidence needed to resolve it.

## Direct the Remediation

- Recommend the smallest responsible change direction at the owner of the failed invariant. Name the behavior to change and the unrelated behavior to preserve.
- Avoid proposing suppression of symptoms, blanket retries, cache deletion, dependency upgrades, or broad refactors unless evidence identifies that action as the responsible remediation.
- Distinguish diagnosis from implementation. Do not edit code, workflow configuration, or external state unless the user also authorized implementation.

## Define Exact Verification

1. Reproduce the failing boundary with the narrowest available local command, fixture, environment, or static check. Give exact commands and prerequisites when known.
2. State the expected observation that would confirm the diagnosis and the result that would disconfirm it.
3. Verify the responsible unit first, then the affected job-equivalent path, then the smallest broader regression surface justified by the dependency path.
4. Name the authoritative CI rerun or job set required for final confirmation, including matrix entries or environment conditions that matter.
5. Mark every unexecuted step as planned rather than passed, and record unavailable credentials, services, platforms, or remote permissions.

## Deliver the Triage Report

Use these headings:

- **Run Context:** revision, workflow, jobs, matrix, runner, event, and material environment facts.
- **Evidence Inventory:** sources, provenance, freshness, and unavailable or truncated evidence.
- **Failure Sequence:** ordered steps from the last known-good invariant through the observed run outcome.
- **Primary Diagnosis:** failed invariant, evidence, responsible owner, confidence, and causal explanation; use `unresolved` when no primary failure is established.
- **Cascade Classification:** primary, contributing, independent, cascading, and noise items with evidence for each label.
- **Remediation Direction:** smallest responsible behavioral change and preserved scope.
- **Verification Plan:** exact local checks, expected and disconfirming observations, broader regression checks, and authoritative CI confirmation.
- **Uncertainties:** competing hypotheses, missing evidence, diagnostic impact, and smallest resolution action.
- **Authorization Boundary:** remote or mutating actions that were excluded and any explicit authorization still required.

Keep citations precise enough that another investigator can locate the evidence. Present commands as executed only when retained output proves execution.
