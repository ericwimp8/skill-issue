import { resolveRoute } from '../src';
const route = resolveRoute(process.argv[2], true);
if (!route) process.exit(2);
