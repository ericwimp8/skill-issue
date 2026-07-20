# Accessibility First-Pass Evaluation Contract

- **Target:** `showcase-skills/accessibility-first-pass/skill/accessibility-first-pass/`
- **Goal:** produce a bounded, evidence-labeled first-pass accessibility review that prioritizes actionable user impact without overstating completeness or conformance.
- **Intended use:** preliminary review of supplied web pages, features, implementations, or representative fixtures.
- **Expected behavior:** inspect available production source, rendered evidence, native tooling, and authoritative guidance; combine applicable automated and manual checks; classify evidence; identify affected users; prioritize findings; and route unresolved checks to human or assistive-technology testing.
- **Expected result:** a report with scope, methods, limitations, prioritized findings, passed checks within tested scope, unresolved checks, next actions, and an explicit first-pass boundary.
- **Boundary:** preserve evidence uncertainty; do not claim overall accessibility, certification, conformance, or absence of barriers; do not install tools or broaden scope without authority.

## Observable Criteria

1. The description selects the skill for naturally phrased first-pass web accessibility review requests.
2. The packaged evidence reference is opened and used for evidence classes, authority, affected-user reasoning, priority, and follow-up boundaries.
3. Reports distinguish observed, inferred, and unverified behavior.
4. Reports assign review priority from user impact and evidence rather than copying scanner severity.
5. Reports correlate source or rendered evidence with affected users and reproducible inspection steps.
6. Reports identify manual, human, and assistive-technology follow-up precisely.
7. Reports preserve scope and material limitations while giving remediation direction at the likely behavior owner.
8. Reports refuse unsupported accessibility or conformance conclusions, including after a clean automated scan.

## Environment Qualification

- **Harness surface:** OpenAI Codex CLI custom evaluation route.
- **Harness version:** `codex-cli 0.144.1`.
- **Model and reasoning:** `gpt-5.6-sol`, medium.
- **Qualification date:** 2026-07-20 retained in `cli/README.md`; current live version and authentication rechecked before this campaign.
- **Trial method:** one fresh isolated Codex session per one-turn custom scenario, installed project-only with ambient skills and project instructions excluded.
- **Direct proactive-invocation evidence:** the development CLI injects an opaque neutral signal instruction and records the attributed native `command_execution` attempt or completed signal in `result.json` and `events.jsonl`; `plans/harness-setup.md` defines that structured attempt as Codex activation evidence.

## Refinement Mode

Automatic refinement is authorized within `showcase-skills/accessibility-first-pass/`. Any retained failure must be corrected at its semantic owner and the affected evaluation repeated from clean fixtures.
