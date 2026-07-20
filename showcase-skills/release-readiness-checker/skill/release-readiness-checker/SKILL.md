---
name: release-readiness-checker
description: Evidence-bound evaluation of a repository or release candidate against its authoritative release gates. Use when an agent must decide release readiness or prepare a pre-release review without publishing or deploying.
---

# Release Readiness Checker

## Bound the Candidate

- Identify the candidate, release target, included components, revision or artifact identity, environment, and requested decision scope before evaluating readiness.
- Record access, platform, credential, time, and reproducibility limits. Treat unavailable evidence as a constraint on the decision, never as proof of success or failure.
- Do not publish, deploy, tag, upload, promote, sign, submit, or mutate shared data or infrastructure. Stop before any command whose successful effect crosses that boundary.

## Derive the Gates

1. Read project-owned release instructions, manifests, build and deployment configuration, migration procedures, policy files, and candidate source before choosing checks.
2. Trace each applicable requirement to its semantic owner and concrete effect. Follow referenced scripts or wrappers to the implementation they execute; a command name or configured job is not evidence that the behavior ran.
3. Resolve conflicts using the project's authority hierarchy. Record unresolved contradictions as risk rather than silently selecting the convenient source.
4. Build a gate set specific to the candidate and target. Consider implementation, configuration, migrations and compatibility, version and artifact identity, documentation, verification, security and privacy, observability, rollback or recovery, and known risk only where project evidence makes them applicable.

Do not replace project policy with a generic checklist. When the project has no explicit policy for a material concern, identify the evidence-based risk and proposed decision rule separately from established requirements.

## Evaluate Current Evidence

- Inspect the candidate paths end to end before concluding a gate. Prefer production source and current configuration for implemented behavior; use tests, builds, analysis, and runtime checks only as validation evidence after the intended path is established.
- Capture each command exactly with working scope, relevant inputs, exit status, timestamp when material, and output or retained log path. State what the command proves and what it leaves unproved.
- Verify that evidence applies to the exact candidate and target. Mark historical, different-revision, different-platform, partial, or undocumented evidence stale or limited.
- Never infer execution from a passing document claim, an available script, a configured workflow, or an old result. Run a safe current check when authorized and practical; otherwise classify it `not-run` or `blocked`.
- Preserve material warnings, skipped checks, flaky outcomes, unresolved findings, manual assertions, and evidence disagreement.

## Classify Every Gate

Use exactly one status and explain the evidence boundary:

- `passed`: current applicable evidence satisfies the gate for this candidate and target.
- `failed`: current evidence demonstrates that the gate is unsatisfied.
- `blocked`: the gate cannot be evaluated because a required input, capability, environment, credential, decision, or safe access path is unavailable.
- `not-run`: an applicable check has not been executed against the current candidate, although no external blocker has been established.
- `not-applicable`: authoritative candidate or project evidence shows why the gate does not apply.

Do not combine `blocked` or `not-run` into a pass. Do not use `not-applicable` merely because evidence is missing.

## Decide and Report

Follow an explicit project release policy when one exists. Otherwise:

- decide `ready` only when every release-critical gate passed and remaining limitations do not undermine the scoped target;
- decide `not ready` when a release-critical gate failed or an evidenced unresolved risk exceeds the project's stated tolerance;
- decide `undetermined` when a release-critical gate is blocked or not run and available evidence cannot support either readiness or a definite failure.

Produce these sections:

- **Candidate and Scope:** candidate identity, target, included components, authority sources, and limitations.
- **Decision:** `ready`, `not ready`, or `undetermined`, with the controlling gates and policy or decision rule.
- **Gate Results:** one entry per gate with requirement and source, status, candidate-specific evidence, command evidence where applicable, limitations, and risk.
- **Unresolved Risks:** contradiction, uncertainty, possible impact, and the evidence or decision needed.
- **Prioritized Next Actions:** ordered by release impact; name the exact action, responsible owner when known, evidence it must produce, and the gate it resolves.
- **Evidence Index:** inspected paths, commands and outputs, retained artifacts, dates or candidate hashes, and anything deliberately not run.

Before handing off the decision, verify that every readiness claim traces to current candidate evidence, every incomplete gate remains visibly incomplete, and the next action list contains no release side effect.
