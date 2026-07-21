export function readSession(store, userId) {
  return store.get(`session:${userId.trim()}`);
}

export function writeSession(store, userId, value) {
  store.set(`session:${userId.toLowerCase()}`, value);
}
