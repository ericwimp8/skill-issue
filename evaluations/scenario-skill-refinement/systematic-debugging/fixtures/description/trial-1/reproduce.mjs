import { formatAttempt } from './formatter.js';

const actual = formatAttempt(0);
if (actual !== 'attempt-0') {
  throw new Error(`expected attempt-0, received ${actual}`);
}
