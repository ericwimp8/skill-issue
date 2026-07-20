export function charge(event: { order_id: string; amount_cents: number }) {
  return gateway.charge(event.amount_cents);
}
