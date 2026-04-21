#!/usr/bin/env sh
# Entrypoint for the dkg-runner container.
# Loads contract addresses from /addresses/addresses.env, then runs the
# full DKG scenario: create round → wait for nodes → encrypt → decrypt → verify.
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

exec /app/dkg-runner \
  --rpc="${DKG_RUNNER_RPC:-http://anvil:8545}" \
  --manager="$MANAGER" \
  --privkey="${DKG_RUNNER_PRIVKEY:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}" \
  --nodes="${DKG_RUNNER_NODES:-3}" \
  --threshold="${DKG_RUNNER_THRESHOLD:-2}" \
  --log-level="${DKG_RUNNER_LOG_LEVEL:-info}" \
  --disclosure-allowed="${DKG_RUNNER_DISCLOSURE_ALLOWED:-false}"
