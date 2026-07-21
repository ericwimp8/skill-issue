import { readSession, writeSession } from './session.js';

const store = new Map();
writeSession(store, 'User-A', { active: true });
const actual = readSession(store, 'User-A');
if (actual?.active !== true) {
  throw new Error('written session could not be read for the same user');
}
