export function loadWebhookConfig(environment) {
  return {
    endpoint: environment.WEBHOOK_URL,
    credential: redactCredential(environment.WEBHOOK_TOKEN),
  };
}

export function dashboardConfig(config) {
  return { ...config };
}

export function auditConfig(config) {
  return { ...config };
}

function redactCredential(credential) {
  return `${credential.slice(0, 3)}...${credential.slice(-2)}`;
}
