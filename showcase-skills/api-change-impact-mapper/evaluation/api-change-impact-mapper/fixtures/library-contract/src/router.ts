export function resolveRoute(path: string, strict: boolean) {
  const backend = process.env.ROUTER_BACKEND ?? 'memory';
  return lookup(backend, path, strict);
}
