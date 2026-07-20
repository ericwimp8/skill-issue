# Incident Record: Payments Latency and Deployment `p-88`

## Scope and Evidence

This record covers the supplied evidence concerning deployment `p-88`, payments-service timeouts, latency alerts, and the reported rollback. It includes every file under `fixtures/behavior/case-1/` and no external evidence.

| Source ID | Path | Type | Locator convention | Stated time-zone context |
| --- | --- | --- | --- | --- |
| `SRC-ALERTS` | `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/alerts.json` | Automated alert records | JSON array index | `Z` (UTC) in each timestamp |
| `SRC-DEPLOYMENTS` | `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/deployments.csv` | Deployment-system export | CSV row number, including header as row 1 | Explicit `+10:30` offset |
| `SRC-NOTES` | `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/notes.md` | Operator notes | Line number | “local Adelaide time” for the rollback report; no date supplied |
| `SRC-SERVICE` | `showcase-skills/incident-timeline-builder/evaluation/incident-timeline-builder/fixtures/behavior/case-1/service.log` | Service log | Line number | Explicit `+10:30` offset |

The source fixtures remain unchanged. The derived artifacts redact nothing because the supplied evidence contains no secret or personal identifier. No comparison telemetry, traces, deployment diff, rollback-completion event, or additional service-health data was supplied.

## Time Normalization

The display zone is UTC. Complete instants carrying `Z` or an explicit numeric offset are converted directly to UTC while retaining their raw timestamps in `records.jsonl` and `timeline.json`. Exact ties would retain input order; none occur. The rollback report's raw time, `09:16 local Adelaide time`, has a named zone but no date, so it cannot identify a UTC instant and remains unresolved. The two other note statements have no event timestamp. No date, offset, seconds, or ordering was inferred for unresolved evidence.

## Chronological Timeline

| UTC time | Raw timestamp | Classification | Event | Provenance | Confidence or limitation |
| --- | --- | --- | --- | --- | --- |
| `2026-03-18T22:38:00Z` | `2026-03-19T09:08:00+10:30` | Observed | The deployment system recorded completion of `p-88` for `payments`. | `SRC-DEPLOYMENTS`, row 2 | Direct observation of the deployment export; completion does not establish causation. |
| `2026-03-18T22:40:30Z` | `2026-03-19T09:10:30+10:30` | Observed | The service logged an upstream `ledger` timeout with `correlation_id=c-4`. | `SRC-SERVICE`, line 1 | Direct log observation. |
| `2026-03-18T22:42:00Z` | `2026-03-19T09:12:00+10:30` | Observed | The service logged an upstream `ledger` timeout with `correlation_id=c-5`. | `SRC-SERVICE`, line 2 | Direct log observation. |
| `2026-03-18T23:12:00Z` | `2026-03-18T23:12:00Z` | Observed | The alert system recorded `payments_latency_high`. | `SRC-ALERTS`, `$[0]` | Direct alert-system observation; the threshold and onset before this record are unknown. |
| `2026-03-18T23:25:00Z` | `2026-03-18T23:25:00Z` | Observed | The alert system recorded `payments_latency_recovered`. | `SRC-ALERTS`, `$[1]` | Direct alert-system observation; it does not by itself prove service recovery or a rollback effect. |

The first recorded timeout follows deployment completion by 2 minutes 30 seconds. The latency-high alert follows deployment completion by 34 minutes. These are reproducible intervals between recorded events, not proof that the deployment caused either condition.

## Unplaced Evidence

| Unresolved time | Classification | Evidence | Provenance | Limitation |
| --- | --- | --- | --- | --- |
| Raw `09:16`, stated as local Adelaide time | Reported | An operator reported rolling back `p-88`. | `SRC-NOTES`, line 3 | No date is supplied, so the report cannot be normalized or placed relative to the dated records. It also reports an action rather than an automated completion event. |
| No timestamp | Gap | The deployment system has no rollback-completion record. | `SRC-NOTES`, line 4 | The absence is reported in a note and supplies neither completion time nor confirmation of outcome. |
| No timestamp | Reported | The operator suspected `p-88` caused the latency, while stating that no comparison or trace proves it. | `SRC-NOTES`, line 5 | This is an attributed hypothesis with no supplied proving evidence. |

## Inferences and Contradictions

**Bounded hypothesis:** `p-88` may be related to the latency because its recorded completion precedes the two supplied timeout logs and the latency-high alert. The evidence also supports alternatives: the upstream `ledger` may have been independently impaired, the latency may have started before the deployment or alert threshold crossing, or another unsupplied change may be responsible. Sequence and proximity alone do not establish causation.

**Reported rollback relationship:** If the rollback report refers to 2026-03-19 in Adelaide, it could bear on the later recovery alert, but the evidence omits the date and an automated rollback-completion record. That conditional relationship is therefore not placed in the chronology and cannot support a recovery claim.

There is no direct field-level contradiction among the supplied automated records. There is an unresolved evidentiary mismatch between the operator's report of rolling back `p-88` (`SRC-NOTES`, line 3) and the stated absence of a rollback-completion record (`SRC-NOTES`, line 4). The two claims can both be true: initiation or attempted rollback does not prove completion.

## Unresolved Evidence

- The deployment contents, changed code or configuration, and expected behavioral impact are unavailable.
- Traces or request-level comparison evidence linking `p-88`, `correlation_id` values `c-4` and `c-5`, upstream `ledger`, and latency are unavailable.
- The incident's pre-deployment latency baseline and alert threshold history are unavailable, so onset cannot be bounded by the alert timestamp alone.
- The rollback report lacks a date, and no automated rollback-completion status or timestamp is available.
- The alert records do not state whether `payments_latency_recovered` corresponds to user-visible recovery or what intervention preceded it.

## Follow-up Actions

1. Obtain the deployment diff and configuration manifest for `p-88`, plus the immediately preceding version. This determines whether the change could affect calls to `ledger` or the measured payments-latency path.
2. Retrieve distributed traces and latency/error metrics spanning before `2026-03-18T22:38:00Z` through after `2026-03-18T23:25:00Z`, keyed where possible by `c-4`, `c-5`, deployment version, and upstream `ledger`. This tests whether failures are version-correlated and whether latency began before or after deployment completion.
3. Obtain deployment-controller and audit logs for rollback initiation, target version, completion or failure, and exact timestamp; clarify the date and zone of the operator's `09:16` report. This determines whether a completed rollback can be placed before recovery.
4. Retrieve equivalent metrics and traces from unaffected instances, regions, or the prior version during the same window. This supplies the comparison explicitly missing from the operator note and helps distinguish deployment impact from an upstream-wide issue.
5. Obtain the alert rule definition and evaluation history. This clarifies threshold lag, actual onset, and what the recovery event demonstrates.

The highest-value next evidence is version-tagged request traces and comparative latency/error metrics across `p-88` and the prior version, paired with the deployment diff. Together they can test the suspected deployment relationship rather than relying on chronology alone.
