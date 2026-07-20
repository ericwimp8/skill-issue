# Bug Reproduction Kit Evaluation Contract

- **Target:** `showcase-skills/bug-reproduction-kit/skill/bug-reproduction-kit/SKILL.md`
- **Content hash:** `c6679f687223e9da3aceef0040c04410ec9060336f432237f54913ec13e44392`
- **Goal:** turn incomplete bug information into a compact, evidence-backed reproduction package that another person can execute, inspect, and file without hidden context.
- **Intended use:** reproducing, clarifying, filing, or handing off a software defect when the initial report or evidence is incomplete.
- **Expected behavior:** inspect available evidence; separate facts, observations, inferences, and unknowns; establish relevant environment and expected behavior; minimize the reproduction; capture evidence and negative results; ask only material questions; and produce a standalone package even when reproduction fails or is blocked.
- **Expected result:** a ready-to-file issue package with sourced environment facts, prerequisites, minimal steps, expected and actual behavior, reproducibility status, evidence, attempted variations, and explicit open gaps without unsupported claims.
- **Preserved boundaries:** no fabricated evidence or certainty, no unsafe or unauthorized side effects, no requirement that reproduction succeed, and no replacement of source-backed expected behavior with an evaluator-authored answer.
- **Evaluation surface:** generated Markdown artifact from repository and report fixtures.

## Observable Criteria

1. Available project and runtime evidence is inspected before questions are asked.
2. Relevant environment facts have provenance or are marked unknown.
3. Reproduction steps are minimal, ordered, executable from a stated starting state, and include required prerequisites.
4. Expected behavior, actual observations, and inferences remain distinct.
5. Useful evidence is retained or identified with provenance and no invented artifact.
6. Reproduction status and observed frequency do not overstate the evidence.
7. Material gaps state their impact and smallest resolution path.
8. Failed or blocked attempts still produce a useful package containing negative results and next evidence.
9. The final issue title and body stand alone without adding unsupported claims.
