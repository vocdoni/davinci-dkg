import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'
import tsconfigPaths from 'vite-tsconfig-paths'

// Build to ./dist; the static-site Docker image serves it via nginx.
export default defineConfig({
  plugins: [react(), tailwindcss(), tsconfigPaths()],
  resolve: {
    // We deliberately do NOT add the `source` condition here: the SDK must be
    // pre-built (sdk/dist must exist) before the UI is installed/built. The
    // Dockerfile's sdk-builder stage, `make ui-build` and the `sdk-build`
    // package script all do that.
    //
    // viem and @zk-kit/baby-jubjub are declared as direct UI deps so they
    // hoist into ui/node_modules where the SDK's built `dist/` finds them.
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true,
    sourcemap: false,
    target: 'es2020',
  },
  server: {
    // Bind on all interfaces so the dev server is reachable from containers,
    // VMs and other hosts on the LAN. Override with `pnpm dev --host 127.0.0.1`.
    host: '0.0.0.0',
    port: 5174,
  },
})
