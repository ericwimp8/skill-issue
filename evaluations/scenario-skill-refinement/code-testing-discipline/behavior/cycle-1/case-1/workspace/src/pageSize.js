/**
 * Converts an optional query value to a page size.
 * Missing, non-numeric, non-integer, and non-positive values use the default.
 * Valid values are capped at 100.
 */
export function pageSize(value, defaultSize = 25) {
  if (value === undefined) {
    return defaultSize;
  }

  const parsed = Number(value);
  if (!Number.isInteger(parsed) || parsed <= 0) {
    return defaultSize;
  }
  return Math.min(parsed, 100);
}
