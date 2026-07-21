export type TenantPolicy = {
  tenantId: string;
  retentionDays: number;
};

export interface PolicyRepository {
  get(tenantId: string): Promise<TenantPolicy>;
}
