#!/usr/bin/env bash
# Deploy DKG contracts and seed node keys.
# Writes /addresses/addresses.env on success.
set -euo pipefail

RPC_URL="${RPC_URL:-http://anvil:8545}"
PRIVATE_KEY="${PRIVATE_KEY:-0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80}"
CHAIN_ID="${CHAIN_ID:-1337}"
OUTPUT_DIR="/addresses"

mkdir -p "$OUTPUT_DIR"

printf 'Waiting for Anvil at %s' "$RPC_URL"
until cast block-number --rpc-url "$RPC_URL" >/dev/null 2>&1; do
  printf '.'
  sleep 1
done
echo

cd /workspace/solidity

echo "Building contracts..."
forge build --quiet

echo "Deploying contracts..."
forge script script/DeployAll.s.sol:DeployAllScript \
  --chain "$CHAIN_ID" \
  --rpc-url "$RPC_URL" \
  --private-key "$PRIVATE_KEY" \
  --broadcast \
  --quiet

BROADCAST="/workspace/solidity/broadcast/DeployAll.s.sol/${CHAIN_ID}/run-latest.json"
if [ ! -f "$BROADCAST" ]; then
  echo "ERROR: broadcast output not found at $BROADCAST" >&2
  exit 1
fi

REGISTRY=$(jq -r '.transactions[] | select(.contractName=="DKGRegistry") | .contractAddress' "$BROADCAST")
MANAGER=$(jq -r '.transactions[] | select(.contractName=="DKGManager") | .contractAddress' "$BROADCAST")
CONTRIBUTION_VERIFIER=$(jq -r '.transactions[] | select(.contractName=="ContributionVerifier") | .contractAddress' "$BROADCAST")
FINALIZE_VERIFIER=$(jq -r '.transactions[] | select(.contractName=="FinalizeVerifier") | .contractAddress' "$BROADCAST")
PARTIAL_DECRYPT_VERIFIER=$(jq -r '.transactions[] | select(.contractName=="PartialDecryptVerifier") | .contractAddress' "$BROADCAST")
DECRYPT_COMBINE_VERIFIER=$(jq -r '.transactions[] | select(.contractName=="DecryptCombineVerifier") | .contractAddress' "$BROADCAST")
REVEAL_SUBMIT_VERIFIER=$(jq -r '.transactions[] | select(.contractName=="RevealSubmitVerifier") | .contractAddress' "$BROADCAST")
REVEAL_SHARE_VERIFIER=$(jq -r '.transactions[] | select(.contractName=="RevealShareVerifier") | .contractAddress' "$BROADCAST")

cat > "$OUTPUT_DIR/addresses.env" <<EOF
REGISTRY=$REGISTRY
MANAGER=$MANAGER
CONTRIBUTION_VERIFIER=$CONTRIBUTION_VERIFIER
FINALIZE_VERIFIER=$FINALIZE_VERIFIER
PARTIAL_DECRYPT_VERIFIER=$PARTIAL_DECRYPT_VERIFIER
DECRYPT_COMBINE_VERIFIER=$DECRYPT_COMBINE_VERIFIER
REVEAL_SUBMIT_VERIFIER=$REVEAL_SUBMIT_VERIFIER
REVEAL_SHARE_VERIFIER=$REVEAL_SHARE_VERIFIER
EOF

echo "Contract addresses:"
cat "$OUTPUT_DIR/addresses.env"

# ── Seed node keys ─────────────────────────────────────────────────────────
# Register the BabyJubJub public key for each well-known Anvil account.
# The keys must match what crypto/hash/poseidon.go + group/point.go produce
# from the domain "davinci-dkg:test:registry-key:v1".
# We skip seeding here and let each node register its own key on startup,
# since the key derivation runs in Go (not in bash).
echo "Cleaning up any stale node key locks..."
rm -rf /shared/keylock-*

echo "Node key registration will be done by each node on startup."
echo "Deployer done."
