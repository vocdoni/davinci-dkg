// Vitest global setup — runs once before any test file, in the main thread.
//
// Responsibilities:
//   1. Check that integration tests are enabled (RUN_INTEGRATION_TESTS=true).
//   2. Start the Anvil + deployer Docker Compose stack (or connect to an
//      external testnet when env vars are set).
//   3. Provide { rpcUrl, managerAddress, registryAddress, addressesFile } to
//      all test files via Vitest's provide/inject mechanism.
//   4. Return a teardown callback that stops the stack on exit.
//
// Test files access the values with:
//   import { inject } from 'vitest';
//   const rpcUrl = inject('rpcUrl');

import type { GlobalSetupContext } from 'vitest/node';
import { startHarness, ENV_RUN_TESTS } from './helpers/harness.js';

export async function setup({ provide }: GlobalSetupContext) {
  if (process.env[ENV_RUN_TESTS] !== 'true') {
    // Signal to tests that the harness is unavailable.
    provide('integrationEnabled', false);
    provide('rpcUrl',            '');
    provide('managerAddress',    '');
    provide('registryAddress',   '');
    provide('addressesFile',     '');
    return;
  }

  console.log('[sdk-test] Starting testnet harness…');
  const harness = await startHarness();
  console.log(`[sdk-test] Testnet ready — RPC: ${harness.rpcUrl}`);
  console.log(`[sdk-test] Manager:  ${harness.managerAddress}`);
  console.log(`[sdk-test] Registry: ${harness.registryAddress}`);

  provide('integrationEnabled', true);
  provide('rpcUrl',             harness.rpcUrl);
  provide('managerAddress',     harness.managerAddress);
  provide('registryAddress',    harness.registryAddress);
  provide('addressesFile',      harness.addressesFile);

  return harness.teardown;
}
