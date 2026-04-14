# Solidity Workspace

This workspace follows the same structure used by `davinci-contracts`.

## Main Contracts

- `src/DKGRegistry.sol`: operator registration and long-term public keys
- `src/DKGManager.sol`: round lifecycle and proof-gated submissions for contribution, finalize, partial decrypt, decrypt combine, and secret reconstruction
- `src/libraries/`: shared types, phase helpers, BRLC helpers
- `src/verifiers/`: generated or wrapped verifier contracts

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
