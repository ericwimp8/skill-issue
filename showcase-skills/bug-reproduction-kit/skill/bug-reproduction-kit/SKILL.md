---
name: bug-reproduction-kit
description: Evidence-first workflow for turning incomplete bug reports into reproducible investigation packages. Use when an agent must reproduce, clarify, or prepare a software defect for filing or handoff.
---

# Bug Reproduction Kit

## Establish the Evidence Boundary

- Inspect the supplied report, repository, runtime state, documentation, and existing artifacts before asking questions.
- Keep confirmed facts, direct observations, evidence-supported inferences, and unknowns distinct. Never turn a likely explanation into a reported fact.
- Define the affected behavior and the smallest safe investigation scope. Do not use production data, destructive commands, privileged access, or external side effects without appropriate authority.
- Preserve original evidence. Record where each material fact or artifact came from, and redact secrets or personal data in shared copies.

## Establish the Reproduction Context

- Record only environment facts that can affect the behavior, such as revision, build, operating system, runtime, dependency or service versions, configuration, feature state, account role, device, or data prerequisites.
- Derive values from available sources or commands when possible. Mark unresolved values as unknown and state how they limit the result.
- Identify the expected behavior from the strongest available contract: current source behavior, product specification, documentation, acceptance criteria, or an explicit reporter statement. Cite that source in the package.
- Describe actual behavior only from retained observations. If the original report and the current attempt differ, preserve both rather than reconciling them speculatively.

## Produce the Minimal Reproduction

1. Attempt the reported path without silently filling gaps. Record the exact setup, inputs, actions, and result.
2. Reduce the path while preserving the failure: remove irrelevant actions and vary one material condition at a time.
3. Re-run the smallest credible path enough to report the observed frequency. Do not claim determinism from one success or failure.
4. Capture useful evidence at the failure boundary, such as logs, stack traces, timestamps, request or correlation identifiers, screenshots, recordings, state snapshots, or a minimal fixture.
5. Record negative results and attempted variations when they narrow the trigger or rule out a plausible condition.

Prefer the fewest ordered steps another person can execute from a stated starting state. Include a command or fixture only when it is required to reproduce the observation.

## Handle Missing Information

- Continue with the evidence available when a useful bounded attempt is possible.
- Ask a focused question only when its answer would materially change the reproduction path, expected contract, safety boundary, or interpretation of the result.
- For every blocking gap, name the missing fact, why it matters, and the smallest evidence or action that would resolve it.
- If reproduction is unavailable or unsuccessful, produce the package with the attempted paths, limitations, and next evidence needed. Do not manufacture clean steps or diagnostic certainty.

## Deliver the Reproduction Package

Use the following headings, omitting none. Keep unknowns and unavailable artifacts explicit.

- **Summary:** concise failure statement and affected behavior.
- **Evidence Status:** reproduced, partially reproduced, not reproduced, or blocked; include observed frequency and confidence limits.
- **Environment:** relevant facts with sources, plus material unknowns.
- **Prerequisites:** required state, data, permissions, configuration, and setup.
- **Minimal Reproduction:** numbered actions from a declared starting state, including exact required inputs or commands.
- **Expected Behavior:** expected result and its source.
- **Actual Behavior:** direct observations and the precise divergence point.
- **Evidence:** artifact paths or links, capture context, timestamps or identifiers when relevant, and redaction notes.
- **Attempts and Variations:** negative results or conditions that changed the outcome.
- **Open Gaps:** material missing information, impact, and the next action needed.
- **Ready-to-File Issue:** a standalone title and issue body assembled from the evidence above without adding unsupported claims.

Before filing or handoff, verify that another person can execute the steps using only the package and referenced artifacts, and that every certainty claim is supported by retained evidence.
