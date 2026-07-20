# Incident Timeline Builder Generation Handoff

- **Generated skill:** `showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/`
- **Canonical target:** `showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- **Intake contract:** `showcase-skills/incident-timeline-builder/plans/incident-timeline-builder/incident-timeline-builder-a-to-b-plan.md`
- **Supported surfaces:** portable Agent Skills content and OpenAI Codex project or user delivery.
- **Goal:** turn heterogeneous operational evidence into a deterministic, source-preserving incident chronology without invented temporal or causal claims.
- **Intended use:** reconciling supplied logs, alerts, deployment records, and operator notes for incident analysis, review, or handoff.
- **Expected behavior:** inventory immutable sources; extract provenance-bearing records; normalize only complete explicit instants; preserve raw times and uncertainty; stably order resolved events; separate unplaced evidence; distinguish epistemic classes; and connect gaps to follow-up actions.
- **Expected result:** a reproducible incident record with scoped evidence, normalization notes, chronological events, unplaced evidence, bounded inferences, contradictions, unresolved evidence, and prioritized follow-up actions.
- **Preserved boundaries:** no source-file mutation, invented times or time zones, unsupported ordering, collapse of conflicting accounts, private-data disclosure, or promotion of correlation to causation.
- **Runtime criteria:** all intake completion criteria require representative behavior evidence; the helper additionally requires direct determinism, time-zone, preservation, error, and privacy checks.
- **Known limitations:** the helper accepts already extracted JSONL records and does not infer timestamps from arbitrary source syntax; contextual extraction and evidence classification remain agent responsibilities.
- **Refinement mode:** automatic semantic refinement only for material failures supported by retained evidence.
- **Evaluation route:** `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/`.
- **Generation decision:** continue through script validation and the fresh-agent campaign; no user-owned stop is active.
