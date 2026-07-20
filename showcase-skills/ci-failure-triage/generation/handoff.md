# CI Failure Triage Generation Handoff

- **Generated skill:** `showcase-skills/ci-failure-triage/skill/ci-failure-triage/`
- **Canonical target:** `showcase-skills/ci-failure-triage/skill/ci-failure-triage/SKILL.md`
- **Intake contract:** `showcase-skills/ci-failure-triage/plans/ci-failure-triage/ci-failure-triage-a-to-b-plan.md`
- **Supported surfaces:** portable Agent Skills content and OpenAI Codex project or user delivery.
- **Goal:** turn noisy or incomplete CI evidence into a causal, evidence-backed diagnosis with bounded remediation and exact verification.
- **Intended use:** investigating failed or suspicious CI jobs from supplied logs, workflow configuration, repository source, and available local tooling.
- **Expected behavior:** inventory evidence, reconstruct workflow order, trace relevant production paths, classify primary and cascading failures, preserve unresolved alternatives, identify the smallest responsible owner, and define exact verification without unauthorized remote mutation.
- **Expected result:** a standalone triage report containing run context, evidence inventory, failure sequence, primary diagnosis or unresolved hypotheses, cascade classification, remediation direction, verification plan, uncertainties, and authorization boundary.
- **Preserved boundaries:** source is authority over tests; unavailable evidence remains unavailable; diagnosis does not imply implementation authority; remote CI, secrets, branches, and releases remain unchanged without explicit authorization.
- **Runtime criteria:** all intake completion criteria require representative execution evidence across causally linked, independent, and evidence-limited cases.
- **Known limitations:** a future diagnosis may remain unresolved when logs, workflow revisions, platform access, secrets, or local runtimes are unavailable.
- **Refinement mode:** automatic semantic refinement only when retained evaluation evidence establishes a material contract failure.
- **Evaluation route:** `showcase-skills/ci-failure-triage/evaluation/ci-failure-triage/`.
- **Generation decision:** continue directly into the qualified fresh-agent campaign; no user-owned stop is active.
