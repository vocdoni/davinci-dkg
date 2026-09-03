# Battery: load, concurrency and adversarial tests against a live fleet

`tests/battery` drives a **running** davinci-dkg testnet (Anvil + real
`davinci-dkg-node` daemons) from the outside. It never starts Docker: it
connects through the harness' external mode, funds its own throw-away
accounts with `anvil_setBalance`, and measures what the fleet does.

Every test skips unless `DAVINCI_DKG_BATTERY=1`, so `make test` and the
regular integration suite are unaffected.

## Running

```bash
make testnet-up DKG_NODE_COUNT=32 ...            # or whatever brings the fleet up
export DAVINCI_DKG_BATTERY=1
export DAVINCI_DKG_TEST_RPC_URL=http://127.0.0.1:8545
export DAVINCI_DKG_TEST_ADDRESSES=/tmp/testnet-addresses.env   # REGISTRY=…, MANAGER=…
export DAVINCI_ARTIFACTS_DIR=$HOME/.davinci-dkg-artifacts        # same dir the nodes mount
go test ./tests/battery -run TestFleetStatus -v                  # read-only smoke test
go test ./tests/battery -run TestOrganizerSwarm -v -timeout 40m
go test ./tests/battery -run 'TestShareAdversary|TestCiphertextAdversary' -v -timeout 40m
go test ./tests/battery -run 'TestCommitteeAdversary|TestLazyMember' -v -timeout 60m
```

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
| `TestOrganizerSwarm` | `BATTERY_ORGANIZERS` (8) organizers × `BATTERY_CIPHERTEXTS` (6) ciphertexts, concurrently, in one Live epoch. Shares are released immediately, after `BATTERY_SHARE_DELAY_BLOCKS` (6), or withheld (every 4th). Asserts plaintexts, that withheld slots are not combined after `BATTERY_WITHHELD_WAIT_BLOCKS` (40), reports per-ciphertext latency, partial count (expected `t`, never more than `n`), gas and throughput. |
| `TestShareAdversary` | Tampered Δ, share replayed across ciphertexts, share relayed from a stranger, non-existent indexes, unregistered aid, re-submission after the plaintext landed. |
| `TestCiphertextAdversary` | Policy reverts (submitter, aid, cap, window), malformed points, then three poisons the contract accepts by design: cofactor-subgroup C1 (nodes must publish no partial), undecryptable C2 (every combiner burns a BSGS to 2^50), and a ciphertext copied from another application. Each poison is bracketed by honest ciphertexts whose latency is compared with the one before the poison. |
| `TestCommitteeAdversary` | A fresh operator joins the next epoch's real lottery, then: duplicate / late-registered / unregistered claims, non-member and malformed contributions, a genuine contribution the real nodes finalize with, duplicate contribution, early finalize, abort of a healthy epoch, out-of-policy `createEpoch` at the cadence boundary; once Live, a genuine partial decryption (share recovered from the other members' calldata), its duplicate, a broken proof and a combine after the nodes'. |
| `TestLazyMember` | An operator claims a slot and never contributes; the epoch must still go Live and decrypt. |

`TestCommitteeAdversary` and `TestLazyMember` each need an epoch boundary
(with `EPOCH_DURATION_BLOCKS=300` at 2 s blocks, up to 10 minutes of waiting
plus ~2 minutes of Preparation). The `createEpoch` policy probes race the
nodes' own `createEpoch` at the boundary; if a node lands first they revert
with `InvalidPhase` (cadence gate) instead of `InvalidPolicy` and the row
says so.

## Knobs

| Env | Default | Meaning |
|---|---|---|
| `BATTERY_ORGANIZERS` / `BATTERY_CIPHERTEXTS` | 8 / 6 | swarm size |
| `BATTERY_SHARE_DELAY_BLOCKS` | 6 | delay of the "delayed" release mode |
| `BATTERY_WITHHELD_WAIT_BLOCKS` | 40 | blocks a withheld slot is watched before asserting "not combined" |
| `BATTERY_NO_COMBINE_WAIT_BLOCKS` | 40 | same, for bad shares |
| `BATTERY_COMBINE_WAIT_BLOCKS` | 240 | maximum wait for an expected combine (generous on purpose, see `chaos.md`) |
| `BATTERY_POISON_OBSERVE_BLOCKS` | 45 | blocks between the "early" and "late" status of a poisoned slot |
| `BATTERY_MIN_SERVICE_BLOCKS` | 90 | minimum blocks before the cadence boundary for a Live epoch to be picked |
| `BATTERY_TX_TIMEOUT` | 3m | receipt wait per transaction |
| `BATTERY_CONNECT_TIMEOUT` | 10m | RPC / addresses-file wait |
| `BATTERY_LOG_LEVEL` | warn | level of the library logger (`log` package) |

## Notes on encodings and reuse

- Ciphertexts are produced with `crypto/elgamal` under
  `PK_aid = PK_ep + PK_org` exactly like `cmd/dkgapp encrypt`; shares with
  `crypto/dleq.ProveOrganizerShare`; registration with
  `crypto/schnorr.ProveOrganizerRegister`. Points are the reduced
  twisted-Edwards affine coordinates the contracts and `tests/helpers` use.
- Proofs come from `tests/helpers/proofs.go` (`BuildContributionSubmission`,
  `BuildPartialDecryptionSubmissionFromBase`,
  `BuildDecryptCombineOutputFromCiphertext`). `BuildContributionSubmission`
  uses deterministic share-encryption nonces — fine on a throw-away testnet,
  never for a real operator.
- The adversarial member recovers its private share `d_i = Σ_j f_j(i)` from
  the other members' `submitContribution` calldata with the same
  transcript layout and `shareenc.DecryptShareRoundHash` the node uses, and
  checks `d_i·G` against the share commitment hash finalize pinned on chain
  before proving anything.
- Withheld ciphertexts stay pending in every node's scanner forever (the
  set is capped at 1024 per node). That is by design; a long-running fleet
  used for many battery runs accumulates them.
