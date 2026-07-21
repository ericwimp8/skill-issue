import {
  DiskProfileRepository,
  MemoryProfileRepository,
} from './repositories.js';
import { greeting } from './profile-service.js';

const memory = await greeting(new MemoryProfileRepository({ name: 'Ada' }));
const disk = await greeting(
  new DiskProfileRepository(new URL('./profile.json', import.meta.url)),
);

if (memory !== 'Welcome, Ada') {
  throw new Error(`memory repository failed: ${memory}`);
}
if (disk !== 'Welcome, Ada') {
  throw new Error(`disk repository failed: ${disk}`);
}
