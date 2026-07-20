# Incident Timeline Builder Evaluation Contract

- **Target:** `showcase-skills/incident-timeline-builder/skill/incident-timeline-builder/SKILL.md`
- **Content hash:** `34813535ad5650140c50836528d9a62f767d9c26e0e389ab9ae7a890586a9bf1`
- **Goal:** turn heterogeneous operational evidence into a deterministic, source-preserving incident chronology without invented temporal or causal claims.
- **Intended use:** reconciling supplied logs, alerts, deployment records, and operator notes for incident analysis, review, or handoff.
- **Expected behavior:** inventory sources; extract provenance-bearing records; normalize only complete explicit instants; preserve raw time and uncertainty; stably order resolved events; separate unplaced evidence; distinguish epistemic classes; and connect gaps or contradictions to follow-up actions.
- **Expected result:** a reproducible incident record with scope, time-normalization notes, chronological events, unplaced evidence, bounded inferences, contradictions, unresolved evidence, and prioritized follow-up actions.
- **Boundary:** no source mutation, invented temporal facts, unsupported total ordering, collapsed contradictions, private-data disclosure, or causal claims from chronology alone.
- **Evaluation surface:** generated Markdown incident record plus any helper-produced JSON derived from synthetic repository fixtures.

## Observable Criteria

1. Every supplied source is inventoried and remains byte-identical.
2. Every material timeline statement retains source identity and locator.
3. Explicit offsets and named zones normalize correctly while raw timestamps remain visible.
4. Missing, ambiguous, invalid, or partial times remain unresolved without invention.
5. Resolved exact ties use stable input order and communicate the tie without false precision.
6. Observations, reports, inferences, contradictions, and gaps remain distinguishable.
7. Inferences cite supporting evidence and avoid unsupported causal language.
8. Contradictions retain competing accounts and identify resolution evidence.
9. The final output includes all seven contract sections and useful follow-up actions tied to unresolved evidence.
10. The helper is used when repeated timestamp parsing or ordering is present, and its JSON remains traceable to the extracted records.
