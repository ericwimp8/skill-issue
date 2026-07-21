export function cancelOrder(order, reason) {
  order.status = 'cancelled';
  order.events.push({ type: 'order.cancelled', orderId: order.id, reason });
}

export function expireOrder(order) {
  cancelOrder(order, 'expired');
}
