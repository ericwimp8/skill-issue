import assert from 'node:assert/strict';
import { apiNotificationJob, scheduledNotificationJob } from './src/jobs.js';
import { loadPreferences } from './src/preferences.js';

assert.equal(loadPreferences({}).emailEnabled, false);
assert.equal(apiNotificationJob({}).emailEnabled, false);
assert.equal(scheduledNotificationJob({}).emailEnabled, false);
assert.equal(scheduledNotificationJob({ emailEnabled: true }).emailEnabled, true);

