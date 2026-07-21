import test from 'node:test';
import assert from 'node:assert/strict';
import fs from 'node:fs';

test('discount calculation multiplies before subtracting', () => {
  const source = fs.readFileSync(
    new URL('./orderTotal.js', import.meta.url),
    'utf8',
  );
  assert.match(source, /subtotal \* percentageDiscount/);
  assert.match(source, /subtotal - discount/);
});
