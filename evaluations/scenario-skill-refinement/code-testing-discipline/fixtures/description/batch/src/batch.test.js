import test from 'node:test';
import assert from 'node:assert/strict';
import { batch } from './batch.js';

test('groups items into requested sizes', () => {
  assert.deepEqual(batch([1, 2, 3], 2), [[1, 2], [3]]);
});
