#!/bin/sh
# Render /usr/share/nginx/html/config.json from environment variables.
# Installed under /docker-entrypoint.d/ — nginx:alpine's stock entrypoint
# runs every executable in that directory before starting nginx itself,
# so this script just renders the file and exits.
#
# Lets a single static-site image target any deployment (sepolia / mainnet
# / local anvil) by setting env at `docker run` time.
#
# All env vars are optional; sensible Sepolia defaults match what we ship
# in public/config.json so a bare `docker run` against the published image
# reaches the production Sepolia deployment.
set -eu

: "${DAVINCI_DKG_RPC_URL:=https://eth-sepolia.public.blastapi.io}"
: "${DAVINCI_DKG_MANAGER_ADDRESS:=0xd3ef727b695b21e108497c36f9dcec52d741298a}"
: "${DAVINCI_DKG_CHAIN_ID:=11155111}"
: "${DAVINCI_DKG_CHAIN_NAME:=sepolia}"

CONFIG_FILE=/usr/share/nginx/html/config.json
TMP=$(mktemp)

cat > "$TMP" <<EOF
{
  "rpcUrl": "${DAVINCI_DKG_RPC_URL}",
  "managerAddress": "${DAVINCI_DKG_MANAGER_ADDRESS}",
  "chainId": ${DAVINCI_DKG_CHAIN_ID},
  "chainName": "${DAVINCI_DKG_CHAIN_NAME}"$([ -n "${DAVINCI_DKG_REGISTRY_ADDRESS:-}" ] && printf ',\n  "registryAddress": "%s"' "${DAVINCI_DKG_REGISTRY_ADDRESS}")$([ -n "${DAVINCI_DKG_START_BLOCK:-}" ] && printf ',\n  "startBlock": %s' "${DAVINCI_DKG_START_BLOCK}")
}
EOF

mv "$TMP" "$CONFIG_FILE"
echo "[entrypoint] wrote $CONFIG_FILE:"
cat "$CONFIG_FILE"
echo
