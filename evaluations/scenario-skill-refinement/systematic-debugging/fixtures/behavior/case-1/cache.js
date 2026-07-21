/** Query property insertion order does not affect cache identity. */
export function cacheKey(path, query) {
  const canonicalQuery = Object.fromEntries(
    Object.entries(query).sort(([leftKey], [rightKey]) =>
      leftKey.localeCompare(rightKey),
    ),
  );
  return `${path}?${JSON.stringify(canonicalQuery)}`;
}

export function getCached(cache, path, query) {
  return cache.get(cacheKey(path, query));
}

export function invalidate(cache, path, query) {
  return cache.delete(cacheKey(path, query));
}
