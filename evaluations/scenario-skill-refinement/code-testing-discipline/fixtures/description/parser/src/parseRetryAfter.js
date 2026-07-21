export function parseRetryAfter(value, nowSeconds) {
  const seconds = Number(value);
  if (Number.isFinite(seconds)) {
    return Math.max(0, seconds);
  }

  const retryAtSeconds = Date.parse(value) / 1000;
  return Math.max(0, retryAtSeconds - nowSeconds);
}
