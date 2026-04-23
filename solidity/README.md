# Solidity Workspace

This workspace follows the same structure used by `davinci-contracts`.

## Main Contracts

- `src/DKGRegistry.sol`: operator registration and long-term public keys
- `src/DKGManager.sol`: round lifecycle and proof-gated submissions for contribution, finalize, partial decrypt, decrypt combine, and secret reconstruction
- `src/libraries/`: shared types, phase helpers, BRLC helpers
- `src/verifiers/`: generated or wrapped verifier contracts

## Round policy notes

The `RoundPolicy` struct has a `finalizeNotBeforeBlock` field (uint64) that
must satisfy `finalizeNotBeforeBlock > contributionDeadlineBlock`. The
`finalizeRound` function reverts with `FinalizeTooEarly` when called before
that block — this gives every selected participant a window to submit before
the contribution set is frozen at finalize time. In production, `davinci-dkg-node`
instances finalize automatically once the gate opens, using a deterministic
per-round stagger derived from the lottery seed so only one node submits at
a time.

## Common Commands

```bash
forge build
forge test
forge test --gas-report
./go_bind.sh
```

## Deployment

- `script/DeployAll.s.sol` deploys verifier wrappers and the core contracts.
- `script/Deploy.s.sol` is a thin single-entry wrapper kept for parity with typical Foundry workflows.
- `../testnet/deploy.sh` is the Docker-friendly local deploy path reused by both the multi-node testnet (`testnet/docker-compose.yml`) and the Go integration-test harness (`tests/docker/docker-compose.yml`).
