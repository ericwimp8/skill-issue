# Ground Truth

- The user preference is produced by `loadPreferences` and consumed by both API and scheduler job creation.
- `apiNotificationJob` contains caller-side compensation; `scheduledNotificationJob` exposes the shared defect.
- The smallest complete change canonicalizes `emailEnabled` in `loadPreferences` and removes API-only compensation.
- Both job paths must store `false` for a missing preference and preserve explicit `true`.
- Verification must cover the owner and both manifestation paths.

