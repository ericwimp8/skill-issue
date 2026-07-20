export function forwardToPartner(event: unknown) {
  return partner.post('/order-events', event);
}
