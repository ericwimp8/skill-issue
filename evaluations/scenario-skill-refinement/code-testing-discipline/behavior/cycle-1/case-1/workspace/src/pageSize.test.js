import test from 'node:test';
import assert from 'node:assert/strict';
import { pageSize } from './pageSize.js';

test('uses the default for a missing value', () => {
  assert.equal(pageSize(undefined, 20), 20);
});

test('uses the default for non-positive values', () => {
  assert.equal(pageSize('0', 20), 20);
  assert.equal(pageSize('-5', 20), 20);
});

test('caps large page sizes', () => {
  assert.equal(pageSize('500', 20), 100);
});
