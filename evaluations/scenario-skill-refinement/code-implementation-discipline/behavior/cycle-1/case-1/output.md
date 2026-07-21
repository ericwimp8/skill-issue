# Case 1 Output

- `loadPreferences` owns canonical preference normalization; missing or null `emailEnabled` becomes `false`.
- API-only fallback was removed from `apiNotificationJob`.
- API and scheduler paths consume the same canonical value.
- Verification: `npm test` plus focused missing, undefined, null, false, and true checks passed.

