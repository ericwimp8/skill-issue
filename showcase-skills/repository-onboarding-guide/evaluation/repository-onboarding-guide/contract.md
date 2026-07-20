# Repository Onboarding Guide Evaluation Contract

- **Target:** `showcase-skills/repository-onboarding-guide/skill/repository-onboarding-guide/`
- **Goal:** create a concise, source-backed guide that gives a new contributor a reliable route into an unfamiliar repository.
- **Intended use:** repository orientation, contributor onboarding, architecture discovery, and source-grounded setup or workflow handoff.
- **Expected behavior:** resolve scoped authority; map entry points, dependencies, and ownership; trace production behavior to concrete effects; ground commands, state, and contribution workflows; identify conflicts and unknowns.
- **Expected result:** an actionable guide with repository-relative source evidence, at least one complete production trace, explicit execution status, and no unsupported conventions.
- **Boundary:** production source owns behavior; tests and documentation are supporting evidence; secrets and machine-specific values remain undisclosed; unresolved claims remain unknowns.

## Observable Criteria

1. The description selects the skill for naturally phrased unfamiliar-repository onboarding requests.
2. Applicable instruction files are discovered with correct scope and precedence.
3. Entry points, architecture, dependencies, and ownership boundaries are grounded in production source and configuration.
4. A representative behavior trace follows abstractions to a concrete production effect with material state and error paths.
5. Commands include source evidence, prerequisites or working directory when material, and static-versus-executed status.
6. Generated, ignored, secret-bearing, and machine-local state is classified without exposing values.
7. Contribution workflows are grounded; documentation and test conflicts are reported rather than silently trusted.
8. The guide remains concise, actionable, repository-relative, and explicit about material unknowns.

## Environment Qualification

- **Harness surface:** Codex CLI custom evaluation route qualified in `plans/harness-setup.md` and `cli/README.md`.
- **Harness version:** `codex-cli 0.144.1`.
- **Model and reasoning:** `gpt-5.6-sol`, medium.
- **Qualification date:** 2026-07-20.
- **Trial method:** one fresh independent agent per isolated task with `fork_turns: "none"`, exact target reads retained before output, and no inherited parent conversation.
- **Direct proactive-invocation evidence:** the qualified CLI route attributes structured skill activation; this campaign additionally retains exact candidate-selection and target-read evidence from each fresh agent.

## Refinement Mode

Automatic refinement is authorized within `showcase-skills/repository-onboarding-guide/`. Retained failures must be corrected at their semantic owner and affected cases rerun from clean fixtures.

