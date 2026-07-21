import test from 'node:test';
import assert from 'node:assert/strict';
import { profileEndpoint } from './profileEndpoint.js';

test('creates a profile from valid JSON', () => {
  const response = profileEndpoint('{"name":" Ada "}');

  assert.equal(response.status, 201);
  assert.deepEqual(JSON.parse(response.body), { name: 'Ada' });
});

for (const [description, requestBody] of [
  ['malformed JSON', '{"name":'],
  ['a blank profile name', '{"name":"   "}'],
]) {
  test(`returns the stable client error for ${description}`, () => {
    assert.deepEqual(profileEndpoint(requestBody), {
      status: 400,
      headers: { 'content-type': 'application/json' },
      body: JSON.stringify({ error: 'invalid profile' }),
    });
  });
}
