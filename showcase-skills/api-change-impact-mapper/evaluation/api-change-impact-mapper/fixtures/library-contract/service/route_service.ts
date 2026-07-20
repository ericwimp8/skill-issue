import { resolveRoute } from '../src';
export function routeRequest(path: string) {
  return resolveRoute(path, false)?.target ?? '/404';
}
