import test from 'node:test';
import assert from 'node:assert/strict';
import { JobQueue } from './jobQueue.js';

test('drain returns queued jobs', () => {
  const queue = new JobQueue();
  queue.enqueue('first');
  queue.enqueue('second');

  assert.deepEqual(queue.drain(), ['first', 'second']);
});

test('drain removes returned jobs from the queue', () => {
  const queue = new JobQueue();
  queue.enqueue('first');
  queue.enqueue('second');

  assert.deepEqual(queue.drain(), ['first', 'second']);
  assert.deepEqual(queue.drain(), []);
});
