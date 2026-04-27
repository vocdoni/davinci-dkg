#!/usr/bin/env sh
# Entrypoint for the testnet's standalone UI service. Reads the deployer
# addresses, exports the env knobs the UI image's nginx entrypoint expects,
# and hands off to the standard image start-up.
#
# UI_PUBLIC_RPC is the RPC URL the browser will use. When accessing the
# testnet from another host, override it to http://<host-ip>:8545.
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

PUBLIC_RPC="${UI_PUBLIC_RPC:-http://localhost:8545}"
CHAIN_ID="${UI_CHAIN_ID:-1337}"
CHAIN_NAME="${UI_CHAIN_NAME:-anvil-testnet}"

echo "Starting DKG explorer UI (public RPC: $PUBLIC_RPC, chain: $CHAIN_NAME/$CHAIN_ID)"

# Translate to the env vars the standalone UI image expects. The image's
# /docker-entrypoint.d/40-render-config.sh then writes /config.json from
# them and execs nginx.
export DAVINCI_DKG_RPC_URL="$PUBLIC_RPC"
export DAVINCI_DKG_MANAGER_ADDRESS="$MANAGER"
export DAVINCI_DKG_CHAIN_ID="$CHAIN_ID"
export DAVINCI_DKG_CHAIN_NAME="$CHAIN_NAME"

# Hand off to the image's stock nginx entrypoint, which runs every
# executable in /docker-entrypoint.d/ before starting nginx itself.
exec /docker-entrypoint.sh nginx -g 'daemon off;'
