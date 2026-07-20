# Shareable Incident Timeline

## Scope and Evidence

This package covers the incident evidence supplied in `fixtures/behavior/case-3/`.

| Source ID | Source | Type | Locator convention | Stated time-zone context |
| --- | --- | --- | --- | --- |
| `gateway-log` | `fixtures/behavior/case-3/gateway.log` | Automated gateway log | Line number | Explicit `Z` on each timestamp |
| `status-page-draft` | `fixtures/behavior/case-3/status-page.md` | Draft public status report | Line number | `UTC` beside each clock time; no date |
| `internal-note` | `fixtures/behavior/case-3/private-note.md` | Internal operator report | Line number | No action time or zone recorded |

All supplied sources are included and preserved unchanged. No additional sources were supplied. The derived artifacts omit credential values and do not attempt to reconstruct sensitive data. The credential-rotation report is retained only as a provenance-linked, untimed report.

## Time Normalization

The display zone is UTC. Complete timestamps carrying `Z` are normalized directly to UTC and retain second precision. The status-page clock values retain their raw minute precision but remain unresolved because the source supplies no date. The internal note supplies no timestamp. Neither the incident context nor neighboring records is used to add a missing date or time.

## Chronological Timeline

| Time (UTC) | Raw timestamp | Classification | Event | Provenance | Confidence or limitation |
| --- | --- | --- | --- | --- | --- |
| `2026-02-02T04:00:00Z` | `2026-02-02T04:00:00Z` | Observed | Gateway recorded trace `t-1` with HTTP status `503`. | `gateway-log`, line 1 | Direct automated record. |
| `2026-02-02T04:06:00Z` | `2026-02-02T04:06:00Z` | Observed | Gateway recorded trace `t-2` with HTTP status `200`. | `gateway-log`, line 2 | Direct automated record. One successful request does not establish full recovery. |

## Unplaced Evidence

| Raw time | Timestamp state | Classification | Event | Provenance | Supported ordering or limitation |
| --- | --- | --- | --- | --- | --- |
| `04:01 UTC` | Unresolved: date missing | Reported | Status page draft states that elevated errors were under investigation. | `status-page-draft`, line 3 | The draft presents this before its `04:08 UTC` recovery statement; neither statement can be assigned to a calendar date from this source. |
| `04:08 UTC` | Unresolved: date missing | Reported | Status page draft states that the service was recovering. | `status-page-draft`, line 4 | The draft presents this after its `04:01 UTC` investigation statement; it cannot be safely interleaved with the dated gateway events. |
| Not recorded | Unresolved: timestamp missing | Reported | An engineer reported rotating a credential during the incident. | `internal-note`, line 3 | The note bounds the action only to “during the incident”; it does not establish sequence relative to any supplied event. |

## Inferences and Contradictions

- The change from a recorded `503` at `04:00:00Z` to a recorded `200` at `04:06:00Z` is consistent with improvement for those two requests. It does not prove service-wide recovery or identify a cause.
- The status draft's investigation-to-recovery sequence is consistent with improvement, but its missing date prevents evidence-based placement beside the gateway events.
- The credential rotation could be relevant to recovery, unrelated, or consequential in another way. With no action time, audit identifier, affected credential identifier, or system linkage, the supplied evidence cannot support a causal relationship.
- No direct contradiction appears in the supplied evidence. The main limitation is unresolved temporal alignment across sources.

## Unresolved Evidence

- The calendar date associated with the status-page draft times is absent, preventing cross-source alignment.
- The credential rotation lacks non-secret audit metadata, including an action timestamp, actor or automation identity, target resource identifier, and audit event identifier.
- The gateway records contain only two requests and do not establish incident start, duration, scope, or service-wide recovery.
- No deployment, alert, metric, or credential-audit source was supplied to test a causal explanation.

## Follow-up Actions

1. Retrieve the status system's non-secret publication or revision metadata to establish the calendar date and authoritative timestamps; this determines whether the two status statements can be aligned with the gateway events.
2. Query the credential-management or identity-provider audit log for non-secret metadata covering the incident window: action timestamp, actor or automation identity, target resource identifier, operation result, and audit event identifier. Do not retrieve or expose the credential value. This determines when the reported rotation occurred and whether it can be temporally compared with recovery evidence.
3. Retrieve gateway aggregate error-rate and request-volume records spanning before `04:00:00Z` through after `04:06:00Z` on `2026-02-02`; this determines whether the two traces reflect a broader outage and recovery pattern.
4. Retrieve deployment, configuration-change, and alert records for the same validated incident window; this tests alternative explanations without treating temporal proximity as causation.
