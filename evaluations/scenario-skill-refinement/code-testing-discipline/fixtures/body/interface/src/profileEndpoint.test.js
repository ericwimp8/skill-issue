import test from 'node:test';
import assert from 'node:assert/strict';
import { profileEndpoint } from './profileEndpoint.js';

test('creates a profile from valid JSON', () => {
  const response = profileEndpoint('{"name":" Ada "}');

  assert.equal(response.status, 201);
  assert.deepEqual(JSON.parse(response.body), { name: 'Ada' });
});
