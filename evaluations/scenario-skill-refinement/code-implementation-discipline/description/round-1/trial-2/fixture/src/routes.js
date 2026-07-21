import { cancelOrder } from './orders.js';

export function cancelRoute(order, request) {
  cancelOrder(order, request.reason ?? 'customer-request');
}

