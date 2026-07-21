export async function greeting(repository) {
  const profile = await repository.load();
  return `Welcome, ${profile.name}`;
}
