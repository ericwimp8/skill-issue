# Next Diagnostic Action

Capture a failing run before changing notification behavior. Add temporary correlated telemetry around `scheduleRetry`, `subscribe`, and `publish` that records one event identifier, timestamps, retry scheduling and firing, subscriber count and callback identity, and every send invocation. Preserve the resulting trace from the next duplicate report or reproduction attempt.

The current snapshot cannot distinguish a retry-driven second send from duplicate subscription or another caller, and staging has not reproduced the incident. No root cause or source correction is supported until a captured trace identifies the first duplicated control-flow step.
