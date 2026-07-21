import { auditConfig, dashboardConfig, loadWebhookConfig } from './config.js';

export function buildOutputs(environment) {
  const config = loadWebhookConfig(environment);
  return {
    dashboard: dashboardConfig(config),
    audit: auditConfig(config),
  };
}

