#!/usr/bin/env bash
# Run part of the dkg-node fleet on another host against the local testnet.
#
#   testnet/remote-nodes.sh up   HOST COUNT OFFSET   # e.g. up p4u@10.200.0.25 16 16
#   testnet/remote-nodes.sh down HOST
#   testnet/remote-nodes.sh logs HOST
#
# Needs: the repo at ~/davinci-dkg on HOST (rsync it), circuit artifacts at
# ~/.davinci/artifacts there (or DAVINCI_ARTIFACTS_HOST_DIR), and this host's
# Anvil reachable from HOST at DKG_RPC_URL (default: the deployer's
# advertised address, DKG_ADVERTISE_IP:8545).
set -euo pipefail
cmd=${1:?up|down|logs}; host=${2:?user@host}
dir='~/davinci-dkg/testnet'
case "$cmd" in
  up)
    count=${3:?node count}; offset=${4:?key offset}
    ip=${DKG_ADVERTISE_IP:-$(ip route get 1 | awk '{print $7; exit}')}
    rpc=${DKG_RPC_URL:-http://$ip:8545}
    tmp=$(mktemp); curl -fsS "http://127.0.0.1:${DEPLOYER_PORT:-8888}/addresses.env" > "$tmp"
    ssh "$host" "mkdir -p $dir/remote-addresses"
    scp -q "$tmp" "$host:$dir/remote-addresses/addresses.env"; rm -f "$tmp"
    # Persist the fleet's variables in the remote project's .env so plain
    # `docker compose -f docker-compose.remote.yml ps|logs|down` keep working.
    ssh "$host" "cd $dir && printf 'DKG_RPC_URL=%s\nDKG_KEY_OFFSET=%s\nDKG_THRESHOLD=%s\nDKG_COMMITTEE_SIZE=%s\nDKG_MIN_VALID_CONTRIBUTIONS=%s\nDKG_ALPHA_BPS=%s\n' \
      '$rpc' '$offset' '${DKG_THRESHOLD:-3}' '${DKG_COMMITTEE_SIZE:-4}' '${DKG_MIN_VALID_CONTRIBUTIONS:-3}' '${DKG_ALPHA_BPS:-15000}' > .env && \
      docker compose -f docker-compose.remote.yml up -d --scale dkg-node=$count"
    ;;
  down) ssh "$host" "cd $dir && docker compose -f docker-compose.remote.yml down -v" ;;
  logs) ssh "$host" "cd $dir && docker compose -f docker-compose.remote.yml logs -f --tail=50 dkg-node" ;;
  *) echo "unknown command $cmd" >&2; exit 2 ;;
esac
