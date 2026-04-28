#!/usr/bin/env bash
# Write ui/public/config.json from environment variables. Used by both the
# `make ui-dev` / `make ui-build` targets and (mirrored as) the standalone
# UI container's entrypoint.sh.
#
# All env vars are optional — sensible Sepolia defaults match what we ship
# in the unmodified ui/public/config.json so a bare invocation is a no-op
# in practice.
#
# Usage:
#   RPC_URL=http://127.0.0.1:8545 \
#   MANAGER_ADDRESS=0x... \
#   CHAIN_ID=31337 \
#   CHAIN_NAME=anvil \
#   EXPLORER_URL=https://sepolia.etherscan.io \
#     scripts/render-ui-config.sh [output-path]
set -euo pipefail

OUT="${1:-ui/public/config.json}"

: "${RPC_URL:=https://eth-sepolia.public.blastapi.io}"
: "${MANAGER_ADDRESS:=0x01ee71fdce1705c8823f9f8b2f312100165fdd70}"
: "${CHAIN_ID:=11155111}"
: "${CHAIN_NAME:=sepolia}"
: "${EXPLORER_URL:=https://sepolia.etherscan.io}"
REGISTRY_LINE=""
[ -n "${REGISTRY_ADDRESS:-}" ] && REGISTRY_LINE=$(printf ',\n  "registryAddress": "%s"' "$REGISTRY_ADDRESS")
START_BLOCK_LINE=""
[ -n "${START_BLOCK:-}" ] && START_BLOCK_LINE=$(printf ',\n  "startBlock": %s' "$START_BLOCK")
EXPLORER_LINE=""
[ -n "${EXPLORER_URL:-}" ] && EXPLORER_LINE=$(printf ',\n  "explorerUrl": "%s"' "${EXPLORER_URL%/}")

mkdir -p "$(dirname "$OUT")"
cat > "$OUT" <<EOF
{
  "rpcUrl": "${RPC_URL}",
  "managerAddress": "${MANAGER_ADDRESS}",
  "chainId": ${CHAIN_ID},
  "chainName": "${CHAIN_NAME}"${REGISTRY_LINE}${START_BLOCK_LINE}${EXPLORER_LINE}
}
EOF

echo "[render-ui-config] wrote $OUT:"
cat "$OUT"
