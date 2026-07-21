const subscribers = [];

export function subscribe(send) {
  subscribers.push(send);
}

export function publish(message) {
  subscribers.forEach((send) => send(message));
}
