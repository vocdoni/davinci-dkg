#!/usr/bin/env bash
# Self-check for render-ui-config.sh: the environment wins, an existing
# config is the default, and a partial override leaves every other key
# alone (the failure mode that used to reset the manager address to a
# stale literal).
set -euo pipefail

here=$(cd "$(dirname "$0")" && pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
out=$tmp/config.json

get() { sed -n "s/^[[:space:]]*\"$1\"[[:space:]]*:[[:space:]]*\"\{0,1\}\([^\",]*\)\"\{0,1\},\{0,1\}[[:space:]]*$/\1/p" "$out" | head -1; }
want() { [ "$(get "$1")" = "$2" ] || { echo "FAIL: $1 = $(get "$1"), want $2"; exit 1; }; }

# No file, no environment: the built-in snapshot.
bash "$here/render-ui-config.sh" "$out" >/dev/null
want chainName sepolia
want managerAddress 0xd38af14cd3b550e268693b459c08ef7331cb23b0

# The environment wins over both the file and the snapshot.
MANAGER_ADDRESS=0xabc CHAIN_ID=31337 CHAIN_NAME=anvil bash "$here/render-ui-config.sh" "$out" >/dev/null
want managerAddress 0xabc
want chainId 31337

# A partial override keeps the other keys as the file had them.
RPC_URL=http://127.0.0.1:8545 bash "$here/render-ui-config.sh" "$out" >/dev/null
want rpcUrl http://127.0.0.1:8545
want managerAddress 0xabc
want chainName anvil

# Rendering twice with no input is a no-op.
cp "$out" "$tmp/before"
bash "$here/render-ui-config.sh" "$out" >/dev/null
diff -q "$tmp/before" "$out" >/dev/null || { echo "FAIL: not idempotent"; exit 1; }

echo "[render-ui-config.test] ok"
