export function orderTotal(lines, percentageDiscount = 0) {
  const subtotal = lines.reduce(
    (sum, line) => sum + line.unitPrice * line.quantity,
    0,
  );
  const discount = subtotal * percentageDiscount;
  return Math.round((subtotal - discount) * 100) / 100;
}
