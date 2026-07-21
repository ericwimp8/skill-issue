import type { RecordRepository } from './records';

export class CleanupCron {
  constructor(private readonly records: RecordRepository) {}

  async run(now: Date): Promise<void> {
    for (const tenantId of await this.records.tenantsWithRecords()) {
      const cutoff = new Date(now.getTime() - 30 * 86_400_000);
      await this.records.deleteCreatedBefore(tenantId, cutoff);
    }
  }
}
