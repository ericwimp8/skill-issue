export async function replayArchive() {
  for await (const event of archive.readAll('orders.placed'))
    await bus.publish('orders.placed', event);
}
