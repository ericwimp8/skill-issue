export function routeFor(request) {
  const region = request.headers['x-region'];
  return region === 'EU' ? '/eu' : '/default';
}
