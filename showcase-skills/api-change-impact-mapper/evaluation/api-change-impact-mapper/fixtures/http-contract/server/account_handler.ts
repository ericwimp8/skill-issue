import { loadAccount } from '../storage/account_store';

export async function getAccount(id: string, apiVersion: 'v1' | 'v2') {
  const stored = await loadAccount(id);
  if (!stored)
    return apiVersion === 'v2' ? { status: 200, body: null } : { status: 404 };
  if (apiVersion === 'v2')
    return {
      status: 200,
      body: {
        id,
        profile: { label: stored.display_name },
        status: stored.status,
      },
    };
  return { status: 200, body: stored };
}
