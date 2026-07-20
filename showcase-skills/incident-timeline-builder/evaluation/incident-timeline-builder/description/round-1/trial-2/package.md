# Incident Chronology

## Scope and Evidence

This chronology covers the supplied deployment record, monitoring records, and operator note in `fixtures/description/trial-2/`. No other incident records were supplied.

| Source ID | Source | Type | Locator convention | Stated time-zone context |
| --- | --- | --- | --- | --- |
| SRC-01 | `fixtures/description/trial-2/deploy.md` | Deployment record | Markdown line | `Asia/Tokyo` |
| SRC-02 | `fixtures/description/trial-2/monitor.jsonl` | Automated monitor records | JSONL line | UTC (`Z`) |
| SRC-03 | `fixtures/description/trial-2/operator.md` | Operator report | Markdown line | Explicit `+09:00` offset |

The source files were read without modification. No redactions were needed. `records.jsonl` is the extracted evidence ledger, and `timeline.json` is its deterministic normalized representation.

## Time Normalization

- Display zone: UTC.
- `2026-08-12 17:10:00` in `Asia/Tokyo` resolves to `2026-08-12T08:10:00Z` because Tokyo is UTC+09:00 on that date.
- Monitor timestamps already use UTC (`Z`).
- `2026-08-12T17:20:00+09:00` resolves to `2026-08-12T08:20:00Z`.
- All supplied timestamps have one-second precision. The two monitor records at `08:15:00Z` are simultaneous at the available precision and retain source/input order only as a deterministic display rule, not as evidence of sequence.
- No supplied timestamp is unresolved.

## Chronological Timeline

| Normalized UTC | Raw timestamp | Classification | Event | Provenance | Limitation |
| --- | --- | --- | --- | --- | --- |
| 2026-08-12T08:10:00Z | `2026-08-12 17:10:00` (`Asia/Tokyo`) | Observed | Deployment record states worker revision `w42` completed. | SRC-01, line 3 | Records completion only; it does not establish incident impact or causation. |
| 2026-08-12T08:15:00Z | `2026-08-12T08:15:00Z` | Observed | Monitor recorded `queue_depth_high`. | SRC-02, line 1 | Simultaneous with the next observation at available precision. |
| 2026-08-12T08:15:00Z | `2026-08-12T08:15:00Z` | Observed | Monitor recorded `worker_errors_high`. | SRC-02, line 2 | Simultaneous with the preceding observation at available precision. |
| 2026-08-12T08:20:00Z | `2026-08-12T17:20:00+09:00` | Reported | Operator reported that queue depth looked normal again. | SRC-03, line 3 | Human report; the note explicitly does not establish what caused recovery. |

## Unplaced Evidence

None. Every supplied record contains a complete, unambiguous instant.

## Inferences and Contradictions

### Bounded hypotheses

- The deployment completion preceded both high-condition monitor observations by five minutes. This temporal proximity makes a deployment relationship reasonable to investigate, but the records do not show that revision `w42` caused either condition.
- The queue-depth and worker-error observations occurred at the same recorded instant. Their simultaneity supports investigating a shared condition or interaction, but it establishes neither direction nor common cause.
- The operator's recovery report followed the alerts by five minutes. It shows a reported return to normal queue depth by `08:20:00Z`; it does not identify a recovery action or explain whether worker errors also cleared.

### Contradictions

No direct contradiction appears in the supplied records. The monitor's earlier `queue_depth_high` observation and the operator's later report that queue depth looked normal describe different instants and can both be true.

## Unresolved Evidence

- No deployment start time, change contents, rollout scope, or revision comparison is supplied, limiting assessment of whether `w42` affected the monitored conditions.
- No alert thresholds, measurements, affected worker identifiers, or alert-clear records are supplied, limiting impact and duration estimates.
- No worker-error recovery observation is supplied, so its state after `08:15:00Z` is unknown.
- No operator action or automated remediation record is supplied, so the recovery mechanism cannot be determined.
- No service-impact or user-impact record is supplied.

## Follow-up Actions

1. Retrieve the `w42` deployment manifest, rollout logs, and previous-revision diff to test whether the changed components could affect queue depth or worker errors.
2. Retrieve underlying queue-depth values, worker-error details, alert thresholds, and alert-clear events from `08:05` through `08:30` UTC to determine onset, duration, scope, and whether both signals recovered.
3. Retrieve worker and queue service logs keyed to affected instances for the same window to test shared-cause and directional hypotheses.
4. Ask the operator what observations and actions informed the `08:20` report to distinguish passive recovery from intervention.
5. Retrieve service-level and user-impact telemetry for the incident window to establish operational impact.
