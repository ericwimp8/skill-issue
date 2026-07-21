export function loadPreferences(record) {
  return {
    emailEnabled: record.emailEnabled ?? false,
    locale: record.locale ?? 'en',
  };
}
