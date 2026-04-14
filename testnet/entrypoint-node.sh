#!/usr/bin/env sh
# Entrypoint for a DKG node container.
#
# Determines which Anvil private key to use from the container's hostname
# (e.g., testnet-dkg-node-3 → index 2 → third key in ANVIL_KEYS).
#
# Then loads contract addresses from /addresses/addresses.env and starts
# the davinci-dkg-node daemon.
set -e

# ── load contract addresses ────────────────────────────────────────────────
if [ -f /addresses/addresses.env ]; then
  # shellcheck disable=SC1091
  set -a
  . /addresses/addresses.env
  set +a
else
  echo "ERROR: /addresses/addresses.env not found" >&2
  exit 1
fi

# ── pick private key by atomic locking ────────────────────────────────────
INDEX=1
while [ "$INDEX" -le 32 ]; do
  if mkdir "/shared/keylock-$INDEX" 2>/dev/null; then
    break
  fi
  INDEX=$((INDEX+1))
done

KEY="$(echo "$ANVIL_KEYS" | tr ' ' '\n' | sed -n "${INDEX}p")"
if [ -z "$KEY" ]; then
  echo "ERROR: no key available for index $INDEX" >&2
  exit 1
fi

echo "Claimed key index $INDEX for this node"
echo "Starting davinci-dkg-node..."

exec /app/davinci-dkg-node \
  --web3.rpc="$DAVINCI_DKG_WEB3_RPC" \
  --privkey="$KEY" \
  --registry="$REGISTRY" \
  --manager="$MANAGER" \
  --log.level="${DAVINCI_DKG_LOG_LEVEL:-info}" \
  --shared-dir="${DAVINCI_DKG_SHARED_DIR:-/shared}" \
  --poll-interval=5s
