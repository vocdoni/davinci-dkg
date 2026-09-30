# Battery

Load, concurrency and adversarial tests against a running testnet: Anvil plus real
`davinci-dkg-node` daemons. The battery never starts Docker itself. It connects to the chain,
funds throw-away accounts with `anvil_setBalance` and measures what the fleet does. Every test
skips unless `DAVINCI_DKG_BATTERY=1`, so `make test` and the integration suite are unaffected.

CI runs the swarm and the reveal adversary against a four-node testnet on pushes to `main`, on
pull requests and nightly (`.github/workflows/battery.yml`).

## Running

```bash
make battery-testnet-up DKG_NODE_COUNT=8 DKG_THRESHOLD=5
curl -fsS http://127.0.0.1:8888/addresses.env > /tmp/testnet-addresses.env
export DAVINCI_DKG_BATTERY=1
export DAVINCI_DKG_TEST_RPC_URL=http://127.0.0.1:8545
export DAVINCI_DKG_TEST_ADDRESSES=/tmp/testnet-addresses.env   # REGISTRY=..., MANAGER=...
export DAVINCI_ARTIFACTS_DIR=$HOME/.davinci/artifacts           # the directory the nodes mount
go test ./tests/battery -run TestFleetStatus -v                 # read-only smoke test
go test ./tests/battery -run TestOrganizerSwarm -v -timeout 40m
go test ./tests/battery -run 'TestRevealAdversary|TestCrossApplicationAdversary' -v -timeout 40m
go test ./tests/battery -run TestCommitteeAdversary -v -timeout 60m
```

`make battery` runs `BATTERY_RUN` (default `TestOrganizerSwarm|TestRevealAdversary`) against
`DAVINCI_DKG_TEST_RPC_URL`. The battery waits up to `BATTERY_CONNECT_TIMEOUT` for the RPC and the
addresses file, so it can start while the deployment is still coming up.

Anvil's default accounts 0–31 belong to the deployer and the nodes, whose transaction managers
allocate nonces locally; sharing one of those keys would stall a node. The battery only uses
fresh random keys.

## Report

Every observation (each transaction with its gas and inclusion latency, each expected revert with
its decoded error, each measurement) is appended to `DAVINCI_DKG_BATTERY_REPORT` (default
`/tmp/battery-report.json`) as it happens, and `TestMain` writes a Markdown summary next to it
with per-scenario tables and gas per transaction kind. With `-v` every row is also logged.

## Scenarios

| Test | What it does |
|---|---|
| `TestFleetStatus` | Prints the contract parameters, registry size and newest epochs. |
| `TestOrganizerSwarm` | `BATTERY_ORGANIZERS` organizers × `BATTERY_CIPHERTEXTS` ciphertexts, concurrently, in waves that fit the free pool keys of the newest `Live` epoch; the rest wait for the next epoch. Organizers register automatic applications, locked ones revealed after a delay, or locked ones whose secret is withheld (every fourth). Checks plaintexts, that withheld applications are never combined, that no ciphertext gets more than `n` partials, and reports latency, partial count, gas and throughput. |
| `TestRevealAdversary` | Wrong and zero organizer secrets, the sealed window before the right reveal, a stranger relaying the reveal, a second reveal, a ciphertext submitted after the reveal, a reveal aimed at an automatic application. |
| `TestCrossApplicationAdversary` | Two applications of one epoch; a ciphertext copied from one into the other must never combine to the original plaintext. |
| `TestCiphertextAdversary` | Policy reverts (submitter, aid, cap, window) and malformed points, then three ciphertexts the contract accepts by design: a small-subgroup `C1` (nodes must publish no partial), an undecryptable `C2`, and a copy from another application. Each is bracketed by honest ciphertexts in a neighbour application whose latency is compared before and after. |
| `TestCommitteeAdversary` | A fresh operator joins the next epoch's lottery, then tries duplicate, late and unregistered claims, non-member and malformed contributions, a duplicate contribution, early finalization, aborting a healthy epoch and out-of-policy `createEpoch`; once `Live`, a genuine partial decryption, its duplicate, a broken proof, a broken Merkle path and a late combine. |

`TestCommitteeAdversary` needs an epoch boundary: with `EPOCH_DURATION_BLOCKS=300` at 2 s blocks
that is up to ten minutes of waiting plus the preparation window. Its `createEpoch` probes race
the nodes at the boundary; when a node lands first they revert `InvalidPhase` instead of
`InvalidPolicy`, and the report says so.

## Settings

| Variable | Default | Meaning |
|---|---|---|
| `BATTERY_ORGANIZERS` / `BATTERY_CIPHERTEXTS` | 6 / 6 | Swarm size |
| `BATTERY_REVEAL_DELAY_BLOCKS` | 6 | Delay before a delayed reveal |
| `BATTERY_WITHHELD_WAIT_BLOCKS` | 40 | Blocks a withheld application is watched before asserting it was not combined |
| `BATTERY_NO_COMBINE_WAIT_BLOCKS` | 40 | The same, for a locked application before its reveal |
| `BATTERY_COMBINE_WAIT_BLOCKS` | 240 | Maximum wait for an expected combine |
| `BATTERY_POISON_OBSERVE_BLOCKS` | 45 | Blocks between the early and late status of an adversarial ciphertext |
| `BATTERY_MIN_SERVICE_BLOCKS` | 90 | Minimum blocks left before the cadence boundary for a `Live` epoch to be used |
| `BATTERY_TX_TIMEOUT` | 3m | Receipt wait per transaction |
| `BATTERY_CONNECT_TIMEOUT` | 10m | Wait for the RPC and the addresses file |
| `BATTERY_LOG_LEVEL` | warn | Library log level |

Proofs come from `tests/helpers/proofs.go`, which uses deterministic share-encryption nonces: fine
on a throw-away testnet, never for a real operator. Ciphertexts of a withheld application stay
pending in every node forever (capped at 1024 per node), so a fleet used for many runs
accumulates them.

## Disrupting the fleet

The swarm reports what happened rather than failing hard, so it can run while you break things
from another terminal. Each ciphertext gets its own report row, every wait is bounded in blocks
and generous, and a slowdown shows up as higher `latencyBlocks` until a bound is exceeded.

- **Restart a node** (`docker restart testnet-dkg-node-<k>`) while ciphertexts are served. Its
  partials come from the next wave, a few blocks later; on boot it rescans
  `--decrypt-lookback-blocks` and must not submit twice (an `AlreadyPartiallyDecrypted` revert in
  its log is harmless). Restarting the node holding combine slot 0 (named in earlier rows'
  `combiner=`) moves the combine to the next node in the rotation.
- **Pause the chain** (`docker pause testnet-anvil-1`, then `docker unpause`) for 30 to 60
  seconds. Nodes log RPC errors and back off; keep the pause under `BATTERY_TX_TIMEOUT` or receipt
  waits expire and rows fail.
- **Stop up to `n − t` committee members** for good. Later waves fill in and combines still land;
  from `n − t + 1` stopped nodes on, rows report `not combined within N blocks`.
- **Cross an epoch boundary** by starting the swarm late in an epoch. The `Live` epoch keeps
  serving while the nodes prepare the next one, so combines slow down; the summary counts
  `combinedAfterEpochBoundary`.

Restarting the deployer breaks the run, since it rewrites the addresses file.
