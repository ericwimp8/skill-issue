import { nextJob } from './queue.js';

const actual = nextJob([
  { id: 'low', priority: 2 },
  { id: 'high', priority: 10 },
  { id: 'medium', priority: 5 },
]);
if (actual.id !== 'high') {
  throw new Error(`expected high, received ${actual.id}`);
}
