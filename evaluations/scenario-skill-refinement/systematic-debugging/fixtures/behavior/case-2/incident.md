# Duplicate Notification Incident

At 14:07 a customer reported receiving the same invoice-ready notification twice. The current snapshot has one retry timer in `schedule.js` and one subscriber registry in `subscriptions.js`. No event ID, timestamps, retry count, subscriber count, logs, or captured failing run were retained. Staging attempts have not reproduced the report.
