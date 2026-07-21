export function loadWebhookConfig(environment) {
  return {
    endpoint: environment.WEBHOOK_URL,
    credential: environment.WEBHOOK_TOKEN,
  };
}

export function dashboardConfig(config) {
  return { ...config, credential: redactCredential(config.credential) };
}

export function auditConfig(config) {
  return { ...config };
}

function redactCredential(credential) {
  return `${credential.slice(0, 3)}...${credential.slice(-2)}`;
}
