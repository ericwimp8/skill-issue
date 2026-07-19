import { defineConfig, loadEnv } from 'vite';
import react from '@vitejs/plugin-react';

export default defineConfig(({ command, mode }) => {
  const env = loadEnv(mode, '.', 'VITE_');

  return {
    base: command === 'serve' ? '/' : env.VITE_BASE_PATH || '/skill-issue/',
    plugins: [react()],
  };
});
