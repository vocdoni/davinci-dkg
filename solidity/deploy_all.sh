#!/usr/bin/env bash
#
# Deploy the full davinci-dkg smart-contract suite: 6 Groth16 verifier
# wrappers + DKGRegistry + DKGManager. The addresses of the deployed
# contracts are extracted from the Foundry script output and written to
# <repo-root>/solidity/.last_deployed_addresses.env so that node and
# explorer configurations can pick them up without further parsing.
#
# Required environment (loaded from .env at the repo root or solidity/.env):
#   RPC_URL       JSON-RPC endpoint of the target chain
#   CHAIN_ID      numeric chain ID (matches RPC_URL)
#   PRIVATE_KEY   deployer private key (hex, with or without 0x prefix)
#
# Optional:
#   ETHERSCAN_API_KEY   if set, each contract is submitted to Etherscan
#                       for source verification after deployment.
#   SKIP_TESTS=1        skip the `forge test` gate before broadcasting.
#   VERIFIER_URL        custom Etherscan-compatible API URL (e.g. for L2s
#                       or block explorers like Blockscout).
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
ROOT_DIR="$(dirname "$SCRIPT_DIR")"

# ── colors ────────────────────────────────────────────────────────────────
if [ -t 1 ]; then
    GREEN=$'\033[0;32m'
    YELLOW=$'\033[1;33m'
    RED=$'\033[0;31m'
    BLUE=$'\033[0;34m'
    BOLD=$'\033[1m'
    NC=$'\033[0m'
else
    GREEN="" YELLOW="" RED="" BLUE="" BOLD="" NC=""
fi

info()  { printf '%b[info]%b  %s\n'  "$BLUE"   "$NC" "$*"; }
warn()  { printf '%b[warn]%b  %s\n'  "$YELLOW" "$NC" "$*" >&2; }
error() { printf '%b[error]%b %s\n'  "$RED"    "$NC" "$*" >&2; }
ok()    { printf '%b[ok]%b    %s\n'  "$GREEN"  "$NC" "$*"; }

# ── load .env ─────────────────────────────────────────────────────────────
# Precedence: solidity/.env (closest to the script) overrides the repo-root
# .env, which overrides environment variables already set in the shell.
for candidate in "$ROOT_DIR/.env" "$SCRIPT_DIR/.env"; do
    if [ -f "$candidate" ]; then
        # shellcheck disable=SC1090
        set -a; . "$candidate"; set +a
        ok "loaded environment from $candidate"
    fi
done

# ── required variables ───────────────────────────────────────────────────
missing=0
for var in RPC_URL CHAIN_ID PRIVATE_KEY; do
    if [ -z "${!var:-}" ]; then
        error "$var is not set"
        missing=1
    fi
done
if [ "$missing" -ne 0 ]; then
    error "required variables missing — aborting"
    exit 1
fi

# Normalise PRIVATE_KEY: cast needs the 0x prefix.
case "$PRIVATE_KEY" in
    0x*) ;;
    *)   PRIVATE_KEY="0x$PRIVATE_KEY" ;;
esac

# ── optional etherscan verification ──────────────────────────────────────
VERIFY_ARGS=()
if [ -n "${ETHERSCAN_API_KEY:-}" ]; then
    VERIFY_ARGS+=(--verify --etherscan-api-key "$ETHERSCAN_API_KEY")
    if [ -n "${VERIFIER_URL:-}" ]; then
        VERIFY_ARGS+=(--verifier-url "$VERIFIER_URL")
    fi
    ok "etherscan verification enabled"
else
    warn "ETHERSCAN_API_KEY is not set — contract verification will be skipped"
fi

# ── preflight ────────────────────────────────────────────────────────────
DEPLOYER_ADDR="$(cast wallet address "$PRIVATE_KEY")"

printf '\n'
printf '%b===========================================%b\n' "$BOLD" "$NC"
printf '%b  davinci-dkg contract deployment%b\n'           "$BOLD" "$NC"
printf '%b===========================================%b\n' "$BOLD" "$NC"
printf '  RPC URL      : %s\n' "$RPC_URL"
printf '  Chain ID     : %s\n' "$CHAIN_ID"
printf '  Deployer     : %s\n' "$DEPLOYER_ADDR"
printf '  Verification : %s\n' "$([ -n "${ETHERSCAN_API_KEY:-}" ] && echo "ON" || echo "OFF")"
printf '\n'

# ── run test gate ────────────────────────────────────────────────────────
if [ "${SKIP_TESTS:-0}" != "1" ]; then
    info "running 'forge test' before broadcasting (SKIP_TESTS=1 to bypass) ..."
    (cd "$SCRIPT_DIR" && forge test >/dev/null)
    ok "forge test passed"
else
    warn "skipping forge test (SKIP_TESTS=1)"
fi

# ── broadcast ────────────────────────────────────────────────────────────
DEPLOY_LOG="$SCRIPT_DIR/deploy.log"
info "broadcasting DeployAllScript ..."
(
    cd "$SCRIPT_DIR"
    forge script script/DeployAll.s.sol:DeployAllScript \
        --chain "$CHAIN_ID" \
        --rpc-url "$RPC_URL" \
        --private-key "$PRIVATE_KEY" \
        --broadcast \
        "${VERIFY_ARGS[@]}" \
        -vvv \
        2>&1 | tee "$DEPLOY_LOG"
)

# ── parse deployed addresses ─────────────────────────────────────────────
# The Solidity script logs one line per contract in the form:
#   "<Contract> deployed at: 0x…"
extract() {
    grep -E "^ *$1 deployed at:" "$DEPLOY_LOG" \
        | grep -oE '0x[a-fA-F0-9]{40}' \
        | tail -n 1 \
        || true
}

CONTRIBUTION_VERIFIER=$(extract ContributionVerifier)
FINALIZE_VERIFIER=$(extract FinalizeVerifier)
PARTIAL_DECRYPT_VERIFIER=$(extract PartialDecryptVerifier)
DECRYPT_COMBINE_VERIFIER=$(extract DecryptCombineVerifier)
REVEAL_SUBMIT_VERIFIER=$(extract RevealSubmitVerifier)
REVEAL_SHARE_VERIFIER=$(extract RevealShareVerifier)
REGISTRY=$(extract DKGRegistry)
MANAGER=$(extract DKGManager)

missing=0
for pair in \
    "ContributionVerifier=$CONTRIBUTION_VERIFIER" \
    "FinalizeVerifier=$FINALIZE_VERIFIER" \
    "PartialDecryptVerifier=$PARTIAL_DECRYPT_VERIFIER" \
    "DecryptCombineVerifier=$DECRYPT_COMBINE_VERIFIER" \
    "RevealSubmitVerifier=$REVEAL_SUBMIT_VERIFIER" \
    "RevealShareVerifier=$REVEAL_SHARE_VERIFIER" \
    "DKGRegistry=$REGISTRY" \
    "DKGManager=$MANAGER"; do
    name=${pair%%=*}
    value=${pair#*=}
    if [ -z "$value" ]; then
        error "failed to extract address for $name from $DEPLOY_LOG"
        missing=1
    fi
done
if [ "$missing" -ne 0 ]; then
    exit 1
fi

# ── persist addresses ────────────────────────────────────────────────────
OUT_FILE="$SCRIPT_DIR/.last_deployed_addresses.env"
cat >"$OUT_FILE" <<EOF
# Generated by solidity/deploy_all.sh at $(date -u +%Y-%m-%dT%H:%M:%SZ)
# Deployer: $DEPLOYER_ADDR
# Chain:    $CHAIN_ID ($RPC_URL)
REGISTRY=$REGISTRY
MANAGER=$MANAGER
CONTRIBUTION_VERIFIER=$CONTRIBUTION_VERIFIER
FINALIZE_VERIFIER=$FINALIZE_VERIFIER
PARTIAL_DECRYPT_VERIFIER=$PARTIAL_DECRYPT_VERIFIER
DECRYPT_COMBINE_VERIFIER=$DECRYPT_COMBINE_VERIFIER
REVEAL_SUBMIT_VERIFIER=$REVEAL_SUBMIT_VERIFIER
REVEAL_SHARE_VERIFIER=$REVEAL_SHARE_VERIFIER
EOF

# ── summary ──────────────────────────────────────────────────────────────
printf '\n'
printf '%b===========================================%b\n' "$BOLD" "$NC"
printf '%b  deployment complete%b\n'                        "$GREEN" "$NC"
printf '%b===========================================%b\n' "$BOLD" "$NC"
printf '  DKGRegistry               : %s\n' "$REGISTRY"
printf '  DKGManager                : %s\n' "$MANAGER"
printf '  ContributionVerifier      : %s\n' "$CONTRIBUTION_VERIFIER"
printf '  FinalizeVerifier          : %s\n' "$FINALIZE_VERIFIER"
printf '  PartialDecryptVerifier    : %s\n' "$PARTIAL_DECRYPT_VERIFIER"
printf '  DecryptCombineVerifier    : %s\n' "$DECRYPT_COMBINE_VERIFIER"
printf '  RevealSubmitVerifier      : %s\n' "$REVEAL_SUBMIT_VERIFIER"
printf '  RevealShareVerifier       : %s\n' "$REVEAL_SHARE_VERIFIER"
printf '\n'
printf '  addresses saved to : %s\n'  "$OUT_FILE"
printf '  deploy log         : %s\n'  "$DEPLOY_LOG"
printf '\n'
printf 'Next steps:\n'
printf '  1. Export DAVINCI_DKG_REGISTRY and DAVINCI_DKG_MANAGER in your\n'
printf '     node environment (see .env.example) or source the file above:\n'
printf '         set -a && . %s && set +a\n' "$OUT_FILE"
printf '  2. Restart davinci-dkg-node so it picks up the new contracts.\n'
if [ -z "${ETHERSCAN_API_KEY:-}" ]; then
    printf '  3. (optional) Re-run with ETHERSCAN_API_KEY set to publish the\n'
    printf '     contract source for block-explorer verification.\n'
fi
printf '\n'
