---
name: repository-onboarding-guide
description: Source-backed onboarding guide creation for unfamiliar repositories. Use when a new contributor needs a reliable map of repository instructions, architecture, runtime behavior, setup, validation, workflows, state, and unknowns.
---

# Repository Onboarding Guide

## Establish Authority

1. Locate repository-level and nested instruction files, including `AGENTS.md`, contribution policies, manifests, build configuration, and tool-owned configuration.
2. Record each instruction's scope and precedence before relying on it. Apply deeper scoped instructions only to files within their tree.
3. Treat current production source as behavioral truth. Use documentation, tests, examples, and history to find intent, conflicts, and validation routes; verify their behavioral claims against concrete production paths.
4. Keep observed facts, source-supported inference, and unresolved questions distinct.

## Map the Repository

- Identify externally reachable entry points first: binaries, server or application bootstraps, package exports, framework registrations, scheduled jobs, migrations, and deployment entry points.
- Follow imports, registrations, dependency wiring, configuration loading, and state access to the concrete implementations that own effects.
- Describe major components by responsibility and dependency direction. Name an ownership boundary only when source layout, wiring, or instructions support it.
- Record generated, cached, ignored, secret-bearing, and machine-local paths separately from tracked source and durable configuration. Do not expose secret values.

## Trace Representative Behavior

Choose a behavior central to the contributor's likely work and trace it end to end:

1. Start at the external trigger or public entry point.
2. Follow every relevant wrapper, interface, registration, and helper to the concrete implementation.
3. Record data transformations, configuration and state reads, side effects, persistence or external calls, and material error handling.
4. Cite repository-relative source paths and identifiers at each transition.
5. Stop the trace only at the concrete effect or a clearly named unresolved boundary.

Add further traces only when they explain a materially different subsystem or ownership boundary.

## Ground Contributor Work

- Derive setup, run, build, test, lint, formatting, generation, and release commands from current manifests, task runners, scripts, CI configuration, and scoped instructions. State prerequisites and working directories supported by those owners.
- Distinguish commands verified by execution from commands discovered statically. Never infer successful runtime behavior from a script's presence or a passing unrelated check.
- Describe common contribution workflows only when supported by authoritative instructions, project-owned automation, configuration, or consistent repository history. Label weaker evidence and conflicts.
- Use tests after establishing behavior from production source to show validation coverage, gaps, or mismatches.

## Write the Guide

Keep the guide concise and navigation-oriented. Include:

1. purpose and repository shape;
2. authoritative instructions and their scopes;
3. entry points, architecture, dependencies, and ownership boundaries;
4. at least one end-to-end production behavior trace;
5. setup and project-native validation commands with evidence and execution status;
6. generated, ignored, local, and secret-bearing state;
7. grounded contribution workflows;
8. documentation or test conflicts; and
9. material unknowns with the next source or runtime check that could resolve each one.

Prefer source links and identifiers over copied code. Remove low-value inventories, repeated facts, unsupported conventions, and conclusions that a new contributor cannot act on.

## Validate the Result

- Recheck every architectural, behavioral, command, and workflow claim against its owning source.
- Confirm the representative trace reaches a concrete production effect and does not stop at an abstraction.
- Confirm paths are repository-relative and no secret, identity, or machine-specific value is exposed.
- Report the evidence level actually established and preserve material unknowns.

