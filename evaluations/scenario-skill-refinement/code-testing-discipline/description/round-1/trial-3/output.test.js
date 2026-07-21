import test from 'node:test';
import assert from 'node:assert/strict';
import { TokenStore } from './tokenStore.js';

test('redacts the middle of a token', () => {
  assert.equal(new TokenStore('abcdefgh').redact(), 'ab****gh');
});

test('fully redacts tokens four characters or shorter', () => {
  for (const token of ['a', 'ab', 'abc', 'abcd']) {
    assert.equal(new TokenStore(token).redact(), '*'.repeat(token.length));
  }
});
