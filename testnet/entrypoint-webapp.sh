#!/usr/bin/env sh
# Entrypoint for the DKG explorer webapp service.
#
# Runs davinci-dkg-node in idle mode — no private key, no on-chain participation.
# The node process only serves the embedded React explorer on :8081 and exposes
# /config.json pointing browsers at the browser-reachable Anvil RPC.
set -e

if [ -f /addresses/addresses.env ]; then
  # shellcheck disable=SC1091
  set -a
  . /addresses/addresses.env
  set +a
else
  echo "ERROR: /addresses/addresses.env not found" >&2
  exit 1
fi

# WEBAPP_PUBLIC_RPC is the RPC URL the browser will use. When accessing the
# testnet from another host, override it to http://<host-ip>:8545.
PUBLIC_RPC="${WEBAPP_PUBLIC_RPC:-http://localhost:8545}"

echo "Starting DKG explorer webapp (public RPC: $PUBLIC_RPC)"

exec /app/davinci-dkg-node \
  --web3.rpc="$DAVINCI_DKG_WEB3_RPC" \
  --web3.network="${WEBAPP_CHAIN_NAME:-anvil-testnet}" \
  --registry="$REGISTRY" \
  --manager="$MANAGER" \
  --log.level="${DAVINCI_DKG_LOG_LEVEL:-info}" \
  --webapp.enabled=true \
  --webapp.listen=0.0.0.0:8081 \
  --webapp.public-rpc="$PUBLIC_RPC"
