import {
  DiskProfileRepository,
  MemoryProfileRepository,
} from './repositories.js';

for (const repository of [
  new MemoryProfileRepository({ name: 'Grace' }),
  new DiskProfileRepository(new URL('./profile.json', import.meta.url)),
]) {
  const profile = await repository.load();
  if (typeof profile !== 'object' || typeof profile.name !== 'string') {
    throw new Error('repository contract violated');
  }
}
