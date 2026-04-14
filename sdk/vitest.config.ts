import { defineConfig } from 'vitest/config';

export default defineConfig({
  test: {
    // Single forked process so all test files share the same global state
    // injected by globalSetup (rpcUrl, managerAddress, etc.).
    pool: 'forks',
    poolOptions: {
      forks: { singleFork: true },
    },

    globalSetup: ['./tests/globalSetup.ts'],

    // Per-test timeout: chain interactions can take several seconds.
    testTimeout: 60_000,
    // Hook timeout: beforeAll can start Docker Compose and wait for
    // the deployer, which may take up to a few minutes.
    hookTimeout: 360_000,

    include: ['tests/**/*.test.ts'],

    // Produce a clean, readable summary in CI logs.
    reporters: ['verbose'],
  },
});
