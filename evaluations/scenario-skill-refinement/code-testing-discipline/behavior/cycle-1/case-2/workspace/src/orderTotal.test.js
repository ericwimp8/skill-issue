import test from 'node:test';
import assert from 'node:assert/strict';
import { orderTotal } from './orderTotal.js';

test('applies a percentage discount to the order total', () => {
  const lines = [
    { unitPrice: 20, quantity: 2 },
    { unitPrice: 10, quantity: 1 },
  ];

  assert.equal(orderTotal(lines, 0.2), 40);
});
