import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';

// Build to ./dist; the Go node embeds this directory via go:embed.
export default defineConfig({
  plugins: [react()],
  build: {
    outDir: 'dist',
    emptyOutDir: true,
    sourcemap: false,
    target: 'es2020',
  },
  server: {
    port: 5173,
  },
});
