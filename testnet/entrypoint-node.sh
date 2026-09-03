#!/usr/bin/env sh
# Entrypoint for a DKG node container.
#
# Determines which Anvil private key to use: the first free mkdir lock on the
# shared volume, remembered per container so restarts keep their identity,
# shifted by KEY_OFFSET on worker hosts (keys in /testnet/anvil-keys.txt).
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
# A restarted container must get its previous key back (the operator identity
# is derived from it), so each lock records its owner: the container's
# hostname, which Docker keeps stable across restarts.
ME="$(hostname)"
INDEX=""
for d in /shared/keylock-*; do
  if [ -f "$d/owner" ] && [ "$(cat "$d/owner")" = "$ME" ]; then
    INDEX="${d##*-}"
    echo "Reclaimed key index $INDEX"
    break
  fi
done
if [ -z "$INDEX" ]; then
  INDEX=1
  while [ "$INDEX" -le 32 ]; do
    if mkdir "/shared/keylock-$INDEX" 2>/dev/null; then
      echo "$ME" > "/shared/keylock-$INDEX/owner"
      break
    fi
    INDEX=$((INDEX+1))
  done
fi

INDEX=$((INDEX + ${KEY_OFFSET:-0}))
KEY="$(sed -n "${INDEX}p" /testnet/anvil-keys.txt)"
if [ -z "$KEY" ]; then
  echo "ERROR: no key available for index $INDEX" >&2
  exit 1
fi

echo "Claimed key index $INDEX for this node"
echo "Starting davinci-dkg-node..."

exec /app/davinci-dkg-node \
  --web3.rpc="$DAVINCI_DKG_WEB3_RPC" \
  --privkey="$KEY" \
  --manager="$MANAGER" \
  --log.level="${DAVINCI_DKG_LOG_LEVEL:-info}" \
  --poll-interval=5s
