export type RecordRow = {
  tenantId: string;
  createdAt: Date;
};

export interface RecordRepository {
  tenantsWithRecords(): Promise<string[]>;
  deleteCreatedBefore(tenantId: string, cutoff: Date): Promise<number>;
}
