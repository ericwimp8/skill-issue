import test from 'node:test';
import assert from 'node:assert/strict';
import { TokenStore } from './tokenStore.js';

test('redacts the middle of a token', () => {
  assert.equal(new TokenStore('abcdefgh').redact(), 'ab****gh');
});
