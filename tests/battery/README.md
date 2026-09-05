# Battery: load, concurrency and adversarial tests against a live fleet

`tests/battery` drives a **running** davinci-dkg testnet (Anvil + real
`davinci-dkg-node` daemons) from the outside. It never starts Docker: it
connects through the harness' external mode, funds its own throw-away
accounts with `anvil_setBalance`, and measures what the fleet does.

Every test skips unless `DAVINCI_DKG_BATTERY=1`, so `make test` and the
regular integration suite are unaffected.

## Running

CI runs the swarm and the reveal adversary against a 4-node testnet on pushes
to `main`, on pull requests and nightly (`.github/workflows/battery.yml`).
Locally, `make battery` runs `$(BATTERY_RUN)` against whatever
`DAVINCI_DKG_TEST_RPC_URL` points at.

```bash
make battery-testnet-up DKG_NODE_COUNT=32 ...    # testnet-up plus the battery override, see below
export DAVINCI_DKG_BATTERY=1
export DAVINCI_DKG_TEST_RPC_URL=http://127.0.0.1:8545
export DAVINCI_DKG_TEST_ADDRESSES=/tmp/testnet-addresses.env   # REGISTRY=…, MANAGER=…
export DAVINCI_ARTIFACTS_DIR=$HOME/.davinci-dkg-artifacts        # same dir the nodes mount
go test ./tests/battery -run TestFleetStatus -v                  # read-only smoke test
go test ./tests/battery -run TestOrganizerSwarm -v -timeout 40m
go test ./tests/battery -run 'TestRevealAdversary|TestCrossApplicationAdversary' -v -timeout 40m
go test ./tests/battery -run TestCommitteeAdversary -v -timeout 60m
```

`make battery-testnet-up` layers [`compose.battery.yml`](compose.battery.yml)
over the testnet stack, which runs the nodes with
`DAVINCI_DKG_ACTIVATE_AHEAD=8`. Every epoch deals `MAX_K = 8` pool keys and
a registration claims the next *activated* one; a stock node keeps only two
activated past the claim cursor, so the swarm's burst of registrations would
stall on the nodes' activation rotation. Against a stock fleet the battery
still works — a registration waits for its key's activation
(`BATTERY_ACTIVATION_WAIT_BLOCKS`) and the swarm shrinks its waves to what
is activated — but slower.

The battery polls the RPC and the addresses file for up to
`BATTERY_CONNECT_TIMEOUT` (default 10m) before giving up, so it can be
started while the deployment is still coming up.

Accounts: Anvil's default accounts 0–31 belong to the deployer and the
nodes (their tx managers allocate nonces locally, so sharing a key would
corrupt a node). The battery only ever uses fresh random keys.

## Report

Every observation — each transaction with its gas and inclusion latency,
each expected revert with the decoded custom-error name, each measurement —
is appended to `DAVINCI_DKG_BATTERY_REPORT` (default
`/tmp/battery-report.json`) as it happens; `TestMain` writes a Markdown
summary next to it (`/tmp/battery-report.md`) with per-scenario tables and a
gas digest per transaction kind. `t.Logf` mirrors every row, so `-v` output
is readable on its own.

## Scenarios

| Test | What it does |
|---|---|
| `TestFleetStatus` | Prints immutables, registry size and the newest epochs; checks the torsion-point construction. |
| `TestOrganizerSwarm` | `BATTERY_ORGANIZERS` (6) organizers × `BATTERY_CIPHERTEXTS` (6) ciphertexts, concurrently, in waves of at most `MAX_K` per Live epoch: a wave takes as many organizers as the newest Live epoch has pool keys left, the rest run in the next epoch (the nodes create it early once the pool drains) while the first wave is still being judged. Each organizer registers an automatic application, a locked one revealed after `BATTERY_REVEAL_DELAY_BLOCKS` (6), or a locked one whose secret is withheld (every 4th). Asserts plaintexts, that withheld applications are not combined after `BATTERY_WITHHELD_WAIT_BLOCKS` (40), reports per-ciphertext latency, partial count (expected `t`, never more than `n`), gas and throughput. |
| `TestRevealAdversary` | A wrong / zero organizer secret, the sealed window before the right one lands, a stranger relaying the reveal (it is permissionless), a second reveal, a ciphertext submitted after the reveal, and a reveal aimed at an automatic application. |
| `TestCrossApplicationAdversary` | Two applications of one epoch hold different committee keys; A's ciphertext copied into B must never combine and must never yield A's plaintext. |
| `TestCiphertextAdversary` | Policy reverts (submitter, aid, cap, window), malformed points, then three poisons the contract accepts by design: cofactor-subgroup C1 (nodes must publish no partial), undecryptable C2 (every combiner burns a BSGS to 2^50), and a ciphertext copied from another application. Each poison is bracketed by honest ciphertexts in one shared, never-poisoned neighbour application whose latency is compared with the one before the poison; the scenario claims six of its epoch's eight pool keys; an undecryptable ciphertext taints its own application for the epoch (the nodes stop serving it after the first failed search), which is reported for the same-application probe rather than judged. |
| `TestCommitteeAdversary` | A fresh operator joins the next epoch's real lottery, then: duplicate / late-registered / unregistered claims, non-member and malformed contributions, a genuine contribution the real nodes finalize with, duplicate contribution, early finalize, abort of a healthy epoch, out-of-policy `createEpoch` at the cadence boundary; once Live, a genuine partial decryption (share recovered from the other members' calldata, path rebuilt from the activation transcript), its duplicate, a broken proof, a broken Merkle path and a combine after the nodes'. |

`TestCommitteeAdversary` needs an epoch boundary
(with `EPOCH_DURATION_BLOCKS=300` at 2 s blocks, up to 10 minutes of waiting
plus ~2 minutes of Preparation). The `createEpoch` policy probes race the
nodes' own `createEpoch` at the boundary; if a node lands first they revert
with `InvalidPhase` (cadence gate) instead of `InvalidPolicy` and the row
says so.

## Knobs

| Env | Default | Meaning |
|---|---|---|
| `BATTERY_ORGANIZERS` / `BATTERY_CIPHERTEXTS` | 6 / 6 | swarm size (six organizers plus the reveal adversary's two applications fill one epoch's pool of `MAX_K = 8` keys) |
| `BATTERY_ACTIVATION_WAIT_BLOCKS` | 90 | blocks a registration (or a swarm wave) waits for the nodes to activate its pool key(s) |
| `BATTERY_REVEAL_DELAY_BLOCKS` | 6 | delay of the "delayed" reveal mode |
| `BATTERY_WITHHELD_WAIT_BLOCKS` | 40 | blocks a withheld application is watched before asserting "not combined" |
| `BATTERY_NO_COMBINE_WAIT_BLOCKS` | 40 | same, for a locked application before its reveal |
| `BATTERY_COMBINE_WAIT_BLOCKS` | 240 | maximum wait for an expected combine (generous on purpose, see [`chaos.md`](chaos.md)) |
| `BATTERY_POISON_OBSERVE_BLOCKS` | 45 | blocks between the "early" and "late" status of a poisoned slot |
| `BATTERY_MIN_SERVICE_BLOCKS` | 90 | minimum blocks before the cadence boundary for a Live epoch to be picked |
| `BATTERY_TX_TIMEOUT` | 3m | receipt wait per transaction |
| `BATTERY_CONNECT_TIMEOUT` | 10m | RPC / addresses-file wait |
| `BATTERY_LOG_LEVEL` | warn | level of the library logger (`log` package) |

## Notes on encodings and reuse

- Every scenario asks `waitLiveEpoch` for the number of pool keys it will
  claim, so it never hits `PoolExhausted` half-way; an epoch with too few
  unclaimed keys is waited out until the nodes create the next one. A
  registration waits for the key at the claim cursor to be activated rather
  than reporting `PoolKeyNotActive` as a failure.
- Ciphertexts are produced with `crypto/elgamal` under
  `PK_aid = P_j (+ PK_org)` exactly like `cmd/dkgapp encrypt`, with `P_j` the
  pool key the registration claimed; registration with
  `crypto/schnorr.ProveOrganizerRegister`. Points are the reduced
  twisted-Edwards affine coordinates the contracts and `tests/helpers` use.
- Proofs come from `tests/helpers/proofs.go` (`BuildContributionSubmission`,
  `BuildPartialDecryptionSubmissionFromBase`,
  `BuildDecryptCombineOutputFromCiphertext`). `BuildContributionSubmission`
  uses deterministic share-encryption nonces — fine on a throw-away testnet,
  never for a real operator.
- The adversarial member recovers its private share `d_i = Σ_c f_{c,j}(i)` of
  the application's pool key `j` from the other members' `submitContribution`
  calldata with the same transcript layout and
  `shareenc.DecryptShareRoundHash` the node uses, and checks `d_i·G` against
  the share commitment the key's `activatePoolKey` transcript published
  before proving anything. The Merkle path `submitPartialDecryption` takes is
  rebuilt from that same transcript.
- Ciphertexts of an application whose secret is withheld stay pending in every node's scanner forever (the
  set is capped at 1024 per node). That is by design; a long-running fleet
  used for many battery runs accumulates them.

To disrupt the fleet by hand while the swarm runs, see [`chaos.md`](chaos.md).
