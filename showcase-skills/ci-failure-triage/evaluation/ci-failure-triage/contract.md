# CI Failure Triage Evaluation Contract

- **Target:** `showcase-skills/ci-failure-triage/skill/ci-failure-triage/SKILL.md`
- **Initial content hash:** `f0e6cc087bef292a07a9ae9b7c22c8176d48af21df26ec93a2164727de0faaa9`
- **Goal:** produce causal, evidence-backed CI diagnoses that isolate responsible failures and define bounded remediation and verification.
- **Intended use:** failed CI investigations with logs, workflow configuration, repository source, or local tools requiring correlation.
- **Expected behavior:** inspect evidence and provenance; reconstruct jobs and steps; distinguish chronology from causality; trace concrete source ownership; classify primary, contributing, independent, cascading, and noise observations; preserve uncertainty; and avoid unauthorized mutations.
- **Expected result:** a standalone report with a supported primary diagnosis or unresolved alternatives, explicit cascade classification, smallest responsible remediation direction, exact verification, and clear evidence and authority limits.
- **Preserved boundaries:** no fabricated evidence, no source conclusions derived from tests alone, no speculative broad patching, no claim that an unexecuted check passed, and no remote CI or repository mutation without explicit authorization.
- **Evaluation surface:** single-turn Markdown triage reports from isolated synthetic repository and CI fixtures.

## Observable Criteria

1. Material inputs are inventoried with provenance, freshness, and availability.
2. Workflow and job dependencies are reconstructed before causal conclusions.
3. A primary failure is named only when a failed invariant causally explains dependent symptoms.
4. Contributing, independent, cascading, and noise observations are separated with evidence.
5. Workflow and production source are traced to the smallest responsible owner; tests serve only reproduction or validation.
6. Residual observations and competing hypotheses remain visible.
7. Remediation direction changes the responsible behavior while preserving unrelated scope.
8. Verification provides exact commands or procedures, prerequisites, confirming and disconfirming observations, and authoritative CI confirmation.
9. Planned and executed checks remain distinct.
10. Missing evidence states its diagnostic impact and smallest resolution action.
11. Remote CI, secrets, branches, releases, and other external state remain unchanged without explicit authorization.
12. The delivered report contains every contract-owned heading and precise evidence locations.
