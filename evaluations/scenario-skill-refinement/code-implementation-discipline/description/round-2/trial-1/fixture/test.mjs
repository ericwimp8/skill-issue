import assert from 'node:assert/strict';
import { buildOutputs } from './src/index.js';

const outputs = buildOutputs({ WEBHOOK_URL: 'https://example.test', WEBHOOK_TOKEN: 'secret-token' });
assert.equal(outputs.dashboard.credential, 'sec...en');
assert.equal(outputs.audit.credential, 'sec...en');

