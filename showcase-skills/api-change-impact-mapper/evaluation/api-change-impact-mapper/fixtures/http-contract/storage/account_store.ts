export type StoredAccount = {
  id: string;
  display_name: string;
  status: 'active' | 'suspended' | 'closed';
};
export async function loadAccount(
  id: string,
): Promise<StoredAccount | undefined> {
  return undefined;
}
export function cacheKey(account: StoredAccount) { return `${account.status}:${account.display_name}`; }
