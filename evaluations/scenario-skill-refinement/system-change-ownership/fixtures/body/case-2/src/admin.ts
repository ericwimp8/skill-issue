import type { PolicyRepository, TenantPolicy } from './policy';

export class RetentionSettingsController {
  constructor(private readonly policies: PolicyRepository) {}

  load(tenantId: string): Promise<TenantPolicy> {
    return this.policies.get(tenantId);
  }
}
