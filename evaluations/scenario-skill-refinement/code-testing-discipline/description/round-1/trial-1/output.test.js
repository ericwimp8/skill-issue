import test from 'node:test';
import assert from 'node:assert/strict';
import { parseRetryAfter } from './parseRetryAfter.js';

test('parses delay seconds', () => {
  assert.equal(parseRetryAfter('12', 100), 12);
});

test('preserves a fractional HTTP-date delay', () => {
  const retryAt = 'Wed, 21 Oct 2015 07:28:00 GMT';
  const nowSeconds = Date.parse(retryAt) / 1000 - 1.5;

  assert.equal(parseRetryAfter(retryAt, nowSeconds), 1.5);
});
