// Testnet harness for SDK integration tests.
//
// Two modes:
//   External  — DAVINCI_DKG_TEST_RPC_URL + DAVINCI_DKG_TEST_ADDRESSES are set.
//               The harness connects to an already-running testnet and skips
//               Docker Compose lifecycle management.
//   Local     — Neither env var is set. The harness reserves two random TCP
//               ports, starts the Anvil + deployer Docker Compose stack, waits
//               for the deployer health check, fetches addresses.env, and
//               returns a teardown callback that calls `docker compose down`.

import { spawnSync, spawn } from 'node:child_process';
import * as fs from 'node:fs';
import * as net from 'node:net';
import * as os from 'node:os';
import * as path from 'node:path';
import { fileURLToPath } from 'node:url';

const __filename = fileURLToPath(import.meta.url);
const __dirname  = path.dirname(__filename);

export interface TestnetConfig {
  rpcUrl: string;
  managerAddress: `0x${string}`;
  registryAddress: `0x${string}`;
  /** Path to the addresses.env file on disk (for passing to Go fixture). */
  addressesFile: string;
}

// ── Env var names (mirrors Go constants) ──────────────────────────────────────
export const ENV_RPC_URL   = 'DAVINCI_DKG_TEST_RPC_URL';
export const ENV_ADDRESSES = 'DAVINCI_DKG_TEST_ADDRESSES';
export const ENV_RUN_TESTS = 'RUN_INTEGRATION_TESTS';

// ── Anvil default accounts (matches Go helpers/constants.go) ─────────────────
export const ANVIL_PRIVATE_KEYS = [
  '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
  '0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d',
  '0x5de4111afa1a4b94908f83103eb1f1706367c2e68ca870fc3fb9a804cdab365a',
  '0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6',
  '0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a',
] as const;

// ── Port helpers ──────────────────────────────────────────────────────────────

function reservePort(): Promise<number> {
  return new Promise((resolve, reject) => {
    const server = net.createServer();
    server.listen(0, '127.0.0.1', () => {
      const { port } = server.address() as net.AddressInfo;
      server.close(() => resolve(port));
    });
    server.on('error', reject);
  });
}

// ── Addresses parsing ─────────────────────────────────────────────────────────

function parseAddressesEnv(text: string): Pick<TestnetConfig, 'managerAddress' | 'registryAddress'> {
  const map: Record<string, string> = {};
  for (const line of text.split('\n')) {
    const trimmed = line.trim();
    if (!trimmed || trimmed.startsWith('#')) continue;
    const eqIdx = trimmed.indexOf('=');
    if (eqIdx < 0) continue;
    map[trimmed.slice(0, eqIdx).trim()] = trimmed.slice(eqIdx + 1).trim();
  }
  if (!map['MANAGER'] || !map['REGISTRY']) {
    throw new Error(`addresses.env is missing MANAGER or REGISTRY: ${text}`);
  }
  return {
    managerAddress:  map['MANAGER']  as `0x${string}`,
    registryAddress: map['REGISTRY'] as `0x${string}`,
  };
}

// ── Deployer polling ──────────────────────────────────────────────────────────

async function waitForDeployer(
  deployerUrl: string,
  timeoutMs = 180_000,
): Promise<string> {
  const deadline = Date.now() + timeoutMs;
  let lastError = '';
  while (Date.now() < deadline) {
    try {
      const res = await fetch(`${deployerUrl}/addresses.env`);
      if (res.ok) return res.text();
    } catch (e) {
      lastError = String(e);
    }
    await new Promise((r) => setTimeout(r, 1_000));
  }
  throw new Error(`Deployer at ${deployerUrl} did not become ready. Last error: ${lastError}`);
}

// ── Repo root discovery ───────────────────────────────────────────────────────

function findRepoRoot(start: string): string {
  let dir = start;
  while (true) {
    if (fs.existsSync(path.join(dir, 'go.mod'))) return dir;
    const parent = path.dirname(dir);
    if (parent === dir) throw new Error('Could not find repo root (no go.mod found)');
    dir = parent;
  }
}

// ── Main harness ──────────────────────────────────────────────────────────────

export interface Harness extends TestnetConfig {
  teardown: () => void;
}

export async function startHarness(): Promise<Harness> {
  const rpcUrlEnv   = process.env[ENV_RPC_URL]   ?? '';
  const addressesEnv = process.env[ENV_ADDRESSES] ?? '';

  // ── External mode: testnet already running ───────────────────────────────
  if (rpcUrlEnv && addressesEnv) {
    const content = fs.readFileSync(addressesEnv, 'utf8');
    const addrs   = parseAddressesEnv(content);
    return {
      rpcUrl:          rpcUrlEnv,
      ...addrs,
      addressesFile:   addressesEnv,
      teardown:        () => {},
    };
  }

  // ── Local mode: start Docker Compose ─────────────────────────────────────
  const repoRoot   = findRepoRoot(__dirname);
  const composePath = path.join(repoRoot, 'tests', 'docker', 'docker-compose.yml');

  const [anvilPort, deployerPort] = await Promise.all([reservePort(), reservePort()]);
  const projectName = `sdk-test-${process.pid}`;

  const composeEnv = {
    ...process.env,
    ANVIL_PORT_RPC_HTTP:          String(anvilPort),
    DEPLOYER_SERVER:              String(deployerPort),
    CONTRIBUTION_VERIFIER:        '0x3000000000000000000000000000000000000003',
    PARTIAL_DECRYPT_VERIFIER:     '0x4000000000000000000000000000000000000004',
  };

  const up = spawnSync(
    'docker',
    ['compose', '-p', projectName, '-f', composePath, 'up', '-d', '--build', '--wait'],
    { env: composeEnv, stdio: 'pipe', encoding: 'utf8' },
  );

  if (up.status !== 0) {
    throw new Error(
      `docker compose up failed (exit ${up.status}):\n${up.stderr}\n${up.stdout}`,
    );
  }

  const teardown = () => {
    spawnSync(
      'docker',
      ['compose', '-p', projectName, '-f', composePath, 'down', '-v'],
      { env: composeEnv, stdio: 'pipe' },
    );
  };

  let addressesContent: string;
  try {
    addressesContent = await waitForDeployer(`http://127.0.0.1:${deployerPort}`);
  } catch (e) {
    teardown();
    throw e;
  }

  const addrs = parseAddressesEnv(addressesContent);

  // Write addresses to a temp file so the Go fixture binary can read it.
  const addressesFile = path.join(os.tmpdir(), `dkg-addresses-${process.pid}.env`);
  fs.writeFileSync(addressesFile, addressesContent, 'utf8');

  return {
    rpcUrl:          `http://127.0.0.1:${anvilPort}`,
    ...addrs,
    addressesFile,
    teardown: () => {
      teardown();
      try { fs.unlinkSync(addressesFile); } catch {}
    },
  };
}
