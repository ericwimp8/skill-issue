# Profile Save Can Leave the Mobile App on a Permanent Spinner

## Summary

The incident report states that the production mobile app sometimes continues showing a spinner after a profile save. The only retained runtime evidence is a server excerpt showing one `/profile/save` request that returned HTTP 200 in 184 ms and emitted a response. The client-side behavior has not been observed or reproduced within the permitted investigation boundary.

## Evidence Status

**Blocked.** No reproduction attempt was performed because there is no approved test account, no identified non-production environment, and no authorized device or build context. The reported frequency is only “sometimes”; the attempt count, failure count, and trigger rate are unknown. The server excerpt contains one successful response, but it cannot establish the corresponding client state or whether that request belonged to an affected occurrence.

## Environment

- **Reported environment:** production mobile app, from `fixtures/behavior/case-3/report.md`.
- **Retained server observation:** at `2026-07-17T11:02:04Z`, correlation ID `prof-19`, route `/profile/save`, HTTP status 200, duration 184 ms; a response of 42 bytes was then sent, from `fixtures/behavior/case-3/log.txt`.
- **Required safe environment:** non-production, using an approved test account and non-customer data, from `fixtures/behavior/case-3/contract.md`.
- **Unknown:** app revision and build, mobile OS and version, device model, backend revision and environment corresponding to the excerpt, account role, profile payload, network conditions, configuration, feature state, and whether correlation ID `prof-19` corresponds to a reported spinner occurrence.

These unknowns prevent an executable current reproduction and limit the server evidence to confirming only that one response was sent successfully.

## Prerequisites

The following must be provided before attempting reproduction:

- An approved test account with permission to edit its profile.
- A non-production mobile build connected to a non-production backend.
- Non-customer profile data that can be safely changed and saved.
- The app build identifier, device model, mobile OS version, and relevant feature/configuration state.
- Permission to capture client logs or a screen recording in the test environment.

Production access and customer data are outside the permitted reproduction boundary.

## Minimal Reproduction

**Starting state:** Not currently available. Establish an approved test account in a non-production environment, open its saved profile in the identified mobile build, and enable permitted client evidence capture.

Once the prerequisites exist, execute this bounded path:

1. Record the app build, device, OS, backend environment, account role, configuration, and starting profile values.
2. Change one non-sensitive profile field to a distinct non-customer test value.
3. Start a timestamped screen recording or equivalent UI capture and permitted client logging.
4. Select the profile save action once.
5. Observe whether progress is dismissed and the saved profile appears within 2 seconds.
6. If progress remains after 2 seconds, retain the elapsed time, client logs, capture, request timestamp, and correlation identifier; stop without retrying or changing another condition.
7. Restore the test value if restoration can be performed safely in the same non-production account.

No executable reproduction command or fixture is available in the supplied material.

## Expected Behavior

After a successful profile save, the app dismisses progress and displays the saved profile within 2 seconds. Source: `fixtures/behavior/case-3/contract.md`.

## Actual Behavior

The incident report states that the production app sometimes shows a permanent spinner after profile save. There is no retained device observation, recording, client log, or test attempt confirming the UI divergence. The only direct retained runtime observation is that one server request returned HTTP 200 in 184 ms and a response was sent; the evidence does not establish what the app displayed afterward.

The precise divergence point to test is the client failing to dismiss progress and display the saved profile within 2 seconds after a successful save response.

## Evidence

- `fixtures/behavior/case-3/report.md`: incident statement; no test account, device recording, or production access is available.
- `fixtures/behavior/case-3/contract.md`: expected two-second UI outcome and the approved-test-account, non-customer-data, non-production access boundary.
- `fixtures/behavior/case-3/log.txt`: server excerpt for correlation ID `prof-19` at `2026-07-17T11:02:04Z`, showing HTTP 200 in 184 ms and `response_sent` with 42 bytes.
- No screenshot, recording, client log, app state snapshot, profile payload, or reproduction fixture was supplied.
- No secrets or personal data appear in the supplied artifacts; no redaction was required for this package.

## Attempts and Variations

- No live attempt was made because the safe prerequisites are unavailable.
- No condition was varied, so no trigger or exclusion condition is established.
- The single server excerpt narrows the investigation only by showing that a server response was sent for one request; it does not rule out a client-side, transport-after-send, response-content, timing, or occurrence-correlation issue.

## Open Gaps

- **Approved non-production test context:** Without an approved account, non-production build/backend, and non-customer data, the path cannot be attempted safely. Provide those prerequisites.
- **Affected occurrence correlation:** It is unknown whether `prof-19` belongs to an occurrence with a permanent spinner. Supply a reporter timestamp or matching client/server correlation evidence from an affected occurrence.
- **Client failure evidence:** The spinner and its duration are not retained. Capture a timestamped screen recording and permitted client logs during an authorized attempt.
- **Reproduction environment:** Build, device, OS, backend revision, configuration, and account role are unknown, preventing equivalent setup. Record them from the affected context or define them for the non-production attempt.
- **Trigger and frequency:** The report says “sometimes” without counts or variations. After prerequisites are available, repeat the unchanged minimal path enough times to report attempts and failures, then vary only one material condition at a time.
- **Request and response details:** The profile input and 42-byte response content are unavailable. Capture a redacted non-production request/response pair if permitted, excluding tokens and personal data.

## Ready-to-File Issue

### Title

Profile save can leave the mobile app showing a permanent spinner

### Issue Body

The production incident report states that the mobile app sometimes shows a permanent spinner after profile save. This has not been reproduced because no approved test account, non-production environment, device recording, or production access is available.

The profile-save contract requires the app to dismiss progress and display the saved profile within 2 seconds after a successful save. Reproduction must use an approved test account, non-customer data, and a non-production environment.

The only retained runtime evidence is this server activity:

```text
2026-07-17T11:02:04Z correlation_id=prof-19 route=/profile/save status=200 duration_ms=184
2026-07-17T11:02:04Z correlation_id=prof-19 event=response_sent bytes=42
```

It is unknown whether correlation ID `prof-19` corresponds to an affected UI occurrence, and the excerpt does not show the client state after the response.

Authorized reproduction path:

1. Use an approved test account with non-customer data in a non-production mobile build and backend.
2. Record the build, device, OS, backend environment, account role, configuration, and starting profile values.
3. Change one non-sensitive profile field and begin permitted timestamped UI and client-log capture.
4. Save once and observe whether progress is dismissed and the saved profile appears within 2 seconds.
5. If it does not, retain the elapsed time, recording, client logs, request timestamp, and correlation identifier without changing another condition.

Expected: progress is dismissed and the saved profile appears within 2 seconds after a successful save.

Reported actual: the app sometimes remains on a permanent spinner after profile save. No direct client-side observation is retained.

Next evidence needed: an approved non-production test context, an affected-occurrence correlation, client logs or a screen recording, full environment identifiers, and a redacted non-production request/response pair if capture is permitted.
