import { cacheKey } from './cache.js';

if (cacheKey('/users', { page: 1 }) === cacheKey('/users', { page: 2 })) {
  throw new Error('distinct query values collapsed');
}
if (cacheKey('/users', { page: 1 }) === cacheKey('/teams', { page: 1 })) {
  throw new Error('distinct paths collapsed');
}
