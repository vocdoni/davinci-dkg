import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react';
import path from 'path';

// Build to ./dist; the Go node embeds this directory via go:embed.
export default defineConfig({
  plugins: [react()],
  resolve: {
    // Prefer the "source" export condition so vite compiles the SDK's TypeScript
    // source directly rather than requiring a pre-built dist/.  This makes the
    // webapp build independent of the SDK build step, which avoids a CI race
    // where pnpm's content-store cache has the SDK packed without dist/ and
    // --frozen-lockfile prevents a re-pack after the SDK build step runs.
    conditions: ['source', 'import', 'module', 'browser', 'default'],

    // Vite resolves imports from their real on-disk path (sdk/src/…), not the
    // symlink path (node_modules/@vocdoni/…).  SDK source files import 'viem'
    // and '@zk-kit/baby-jubjub' which live in webapp's node_modules, not in
    // an ancestor of sdk/.  Pinning them here ensures they resolve correctly
    // in all environments (local, Docker, CI) regardless of whether
    // sdk/node_modules/ is populated.
    alias: {
      viem: path.resolve(__dirname, 'node_modules/viem'),
      '@zk-kit/baby-jubjub': path.resolve(
        __dirname,
        'node_modules/@zk-kit/baby-jubjub',
      ),
    },
  },
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
