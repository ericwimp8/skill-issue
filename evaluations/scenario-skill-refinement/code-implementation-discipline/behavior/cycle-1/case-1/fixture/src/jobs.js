import { loadPreferences } from './preferences.js';

export function apiNotificationJob(record) {
  const preferences = loadPreferences(record);
  return { emailEnabled: preferences.emailEnabled };
}

export function scheduledNotificationJob(record) {
  const preferences = loadPreferences(record);
  return { emailEnabled: preferences.emailEnabled };
}
