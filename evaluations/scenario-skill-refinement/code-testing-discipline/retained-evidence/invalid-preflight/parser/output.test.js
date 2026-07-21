import test from 'node:test';
import assert from 'node:assert/strict';
import { parseRetryAfter } from './parseRetryAfter.js';

test('parses delay seconds', () => {
  assert.equal(parseRetryAfter('12', 100), 12);
});

test('preserves fractional delay from an HTTP-date', () => {
  assert.equal(
    parseRetryAfter('Wed, 21 Oct 2015 07:28:00 GMT', 1445412478.5),
    1.5,
  );
});
