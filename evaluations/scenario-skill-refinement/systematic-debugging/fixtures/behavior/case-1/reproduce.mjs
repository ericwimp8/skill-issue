import { cacheKey, getCached, invalidate } from './cache.js';

const cache = new Map();
const first = { page: 2, sort: 'name' };
const reordered = { sort: 'name', page: 2 };
cache.set(cacheKey('/users', first), 'result');

if (getCached(cache, '/users', reordered) !== 'result') {
  throw new Error('equivalent reordered query missed cache');
}
invalidate(cache, '/users', reordered);
if (cache.size !== 0) {
  throw new Error('equivalent reordered query did not invalidate cache');
}
