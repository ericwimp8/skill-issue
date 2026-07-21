import { routeFor } from './route.js';

const request = { headers: { 'x-region': 'eu' } };
const actual = routeFor(request);
if (actual !== '/eu') {
  throw new Error(`expected /eu, received ${actual}`);
}
