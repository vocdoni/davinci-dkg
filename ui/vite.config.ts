import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tsconfigPaths from 'vite-tsconfig-paths'

// Build to ./dist; the static-site Docker image serves it via nginx.
// Eventually the Go node will embed this directory in place of webapp/dist.
export default defineConfig({
  plugins: [react(), tsconfigPaths()],
  resolve: {
    // We deliberately do NOT add the `source` condition here. The old
    // webapp's vite.config used ['source', 'import', ...] to pick up the
    // SDK's untranspiled TypeScript directly, but that same condition
    // collides with @pandacss/is-valid-prop (a Chakra v3 transitive dep)
    // whose published package.json declares `"source": "./src/index.ts"`
    // pointing at a file that isn't shipped — Vite then fails to resolve.
    //
    // Tradeoff: the SDK must be pre-built (sdk/dist must exist) before the
    // UI is installed/built. The Dockerfile's sdk-builder stage and the
    // job_ui CI step both do that, and it's documented in ui/README.md.
    //
    // viem and @zk-kit/baby-jubjub are declared as direct UI deps (see
    // package.json) so they hoist into ui/node_modules where the SDK's
    // built `dist/` can find them via standard pnpm resolution.
  },
  build: {
    outDir: 'dist',
    emptyOutDir: true,
    sourcemap: false,
    target: 'es2020',
  },
  server: {
    // Bind on all interfaces so the dev server is reachable from
    // containers, VMs, and other hosts on the LAN — useful when running
    // a dev build alongside a remote anvil or testing on a tablet/phone
    // pointed at the workstation. Override per-developer with
    // `pnpm dev --host 127.0.0.1` if you need a tighter bind.
    host: '0.0.0.0',
    port: 5174,
  },
})
