import assert from 'node:assert/strict';
import { cancelOrder, expireOrder } from './src/orders.js';

const manual = { id: 'a', status: 'active', events: [] };
cancelOrder(manual, 'customer-request');
assert.equal(manual.events[0].reason, 'customer-request');

const expired = { id: 'b', status: 'active', events: [] };
expireOrder(expired);
assert.equal(expired.events[0].reason, 'expired');

