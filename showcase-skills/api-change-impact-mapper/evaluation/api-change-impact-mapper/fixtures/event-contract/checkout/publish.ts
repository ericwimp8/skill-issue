export function publishOrder(order: Order, useV2: boolean) {
  return bus.publish(
    'orders.placed',
    useV2
      ? {
          order_id: order.id,
          amount_micros: order.total * 1000000,
          status: order.status,
          warehouse_id: order.warehouse,
        }
      : {
          order_id: order.id,
          amount_cents: order.total * 100,
          status: order.status,
        },
  );
}
