# Request

Tenants can now configure record-retention days. The cleanup cron currently deletes records older than 30 days, and the operations team proposes reading tenant settings directly inside the cron loop. Inspect the connected source files and decide where configurable retention should be owned and how the paths should change.
