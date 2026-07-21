import test from 'node:test';
import assert from 'node:assert/strict';
import { batch } from './batch.js';

test('groups items into requested sizes', () => {
  assert.deepEqual(batch([1, 2, 3], 2), [[1, 2], [3]]);
});

test('rejects invalid batch sizes with the public error', () => {
  for (const size of [0, -1, 1.5, NaN, undefined]) {
    assert.throws(
      () => batch([1, 2, 3], size),
      {
        name: 'RangeError',
        message: 'size must be a positive integer',
      },
      `expected size ${String(size)} to be rejected`,
    );
  }
});
