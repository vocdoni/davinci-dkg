#!/usr/bin/env bash
# Write ui/public/config.json from environment variables. Used by the
# `make ui-dev` / `make ui-build` targets and by the UI image build.
#
# Every key falls back to whatever the target file already holds, so a
# partial override (RPC_URL alone, say) keeps the rest of the committed
# config instead of resetting it to a snapshot baked into this script.
# The literals below are only a last resort for a missing file; the
# committed ui/public/config.json is the source of truth for the chain
# we ship.
#
# Usage:
#   RPC_URL=http://127.0.0.1:8545 \
#   MANAGER_ADDRESS=0x... \
#   CHAIN_ID=31337 \
#   CHAIN_NAME=anvil \
#   DEPLOY_BLOCK=11639686 \
#   EXPLORER_URL=https://sepolia.etherscan.io \
#     scripts/render-ui-config.sh [output-path]
#
# DEPLOY_BLOCK is the block the DKGManager was deployed at. The explorer starts
# every historical log scan (operator statistics, per-epoch activity) there. A
# value of 0 means "unknown": the explorer then bisects for the deploy block
# rather than walking the chain from genesis.
set -euo pipefail

OUT="${1:-ui/public/config.json}"

# Read one key out of the existing config. The file is flat and generated
# by this script, so a line-oriented match is enough and keeps the build
# stage free of jq or python.
current() {
	[ -f "$OUT" ] || return 0
	sed -n "s/^[[:space:]]*\"$1\"[[:space:]]*:[[:space:]]*\"\{0,1\}\([^\",]*\)\"\{0,1\},\{0,1\}[[:space:]]*$/\1/p" "$OUT" | head -1
}

# resolve VAR JSON_KEY FALLBACK — environment wins, then the existing
# file, then the fallback.
resolve() {
	local name=$1 value=${!1:-}
	[ -n "$value" ] || value=$(current "$2")
	[ -n "$value" ] || value=$3
	printf -v "$name" '%s' "$value"
}

resolve RPC_URL rpcUrl https://w3.ch4in.net/sepolia
resolve MANAGER_ADDRESS managerAddress 0x6dd442e96cd0b5d8408c2e461a6504be8893229c
resolve CHAIN_ID chainId 11155111
resolve CHAIN_NAME chainName sepolia
resolve DEPLOY_BLOCK deployBlock 11639686
resolve EXPLORER_URL explorerUrl https://sepolia.etherscan.io
resolve REGISTRY_ADDRESS registryAddress ""
resolve START_BLOCK startBlock ""

REGISTRY_LINE=""
[ -n "$REGISTRY_ADDRESS" ] && REGISTRY_LINE=$(printf ',\n  "registryAddress": "%s"' "$REGISTRY_ADDRESS")
START_BLOCK_LINE=""
[ -n "$START_BLOCK" ] && START_BLOCK_LINE=$(printf ',\n  "startBlock": %s' "$START_BLOCK")
EXPLORER_LINE=""
[ -n "$EXPLORER_URL" ] && EXPLORER_LINE=$(printf ',\n  "explorerUrl": "%s"' "${EXPLORER_URL%/}")

mkdir -p "$(dirname "$OUT")"
cat > "$OUT" <<EOF
{
  "rpcUrl": "${RPC_URL}",
  "managerAddress": "${MANAGER_ADDRESS}",
  "chainId": ${CHAIN_ID},
  "chainName": "${CHAIN_NAME}",
  "deployBlock": ${DEPLOY_BLOCK}${REGISTRY_LINE}${START_BLOCK_LINE}${EXPLORER_LINE}
}
EOF

echo "[render-ui-config] wrote $OUT:"
cat "$OUT"
