import test from 'node:test';
import assert from 'node:assert/strict';
import { parseRetryAfter } from './parseRetryAfter.js';

test('parses delay seconds', () => {
  assert.equal(parseRetryAfter('12', 100), 12);
});
