import test from 'node:test';
import assert from 'node:assert/strict';
import { pageSize } from './pageSize.js';

test('uses the default for a missing value', () => {
  assert.equal(pageSize(undefined, 20), 20);
});

test('caps large page sizes', () => {
  assert.equal(pageSize('500', 20), 100);
});
