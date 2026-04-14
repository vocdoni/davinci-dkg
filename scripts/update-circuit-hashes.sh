#!/usr/bin/env bash
# update-circuit-hashes.sh — patch config/circuit_artifacts.go and Solidity verifier
# wrapper contracts with proving/verifying key hashes from a circuit-compile JSON output.
#
# Usage:
#   ./scripts/update-circuit-hashes.sh [artifacts.json] [config/circuit_artifacts.go]
#
# Defaults:
#   artifacts.json          → /tmp/circuit-artifacts.json
#   circuit_artifacts.go    → config/circuit_artifacts.go
#
# Requires: jq, sed
set -euo pipefail

JSON="${1:-/tmp/circuit-artifacts.json}"
CONFIG="${2:-config/circuit_artifacts.go}"
VERIFIERS_DIR="solidity/src/verifiers"

if ! command -v jq &>/dev/null; then
    echo "error: jq is required but not found in PATH" >&2
    exit 1
fi

if [ ! -f "${JSON}" ]; then
    echo "error: artifact JSON not found: ${JSON}" >&2
    echo "Run 'make circuits-compile' first." >&2
    exit 1
fi

# Patch a single hash constant in the Go config file.
update_go_hash() {
    local name="$1"
    local value="$2"
    if [ "${value}" = "null" ] || [ -z "${value}" ]; then
        echo "warning: missing value for ${name}, skipping" >&2
        return
    fi
    sed -i "s|${name}[[:space:]]*= \"[^\"]*\"|${name} = \"${value}\"|" "${CONFIG}"
}

# Patch PROVING_KEY_HASH in a Solidity verifier wrapper contract.
# The hash in Solidity is stored as hex"<64-char-hash>".
update_sol_pk_hash() {
    local sol_file="$1"
    local new_hash="$2"
    if [ "${new_hash}" = "null" ] || [ -z "${new_hash}" ]; then
        echo "warning: missing hash for ${sol_file}, skipping" >&2
        return
    fi
    # Replace the 64-hex-char value inside hex"..." on the PROVING_KEY_HASH line.
    sed -i "s|PROVING_KEY_HASH =\n\?\s*hex\"[0-9a-f]*\"|PROVING_KEY_HASH =\n        hex\"${new_hash}\"|" "${sol_file}" 2>/dev/null || \
    sed -i "/PROVING_KEY_HASH/{ n; s|hex\"[0-9a-f]*\"|hex\"${new_hash}\"|; }" "${sol_file}"
}

# ── Go config ──────────────────────────────────────────────────────────────

update_go_hash ContributionCircuitHash         "$(jq -r '.contribution.circuit_hash'            "${JSON}")"
update_go_hash ContributionProvingKeyHash      "$(jq -r '.contribution.proving_key_hash'         "${JSON}")"
update_go_hash ContributionVerificationKeyHash "$(jq -r '.contribution.verifying_key_hash'       "${JSON}")"

update_go_hash FinalizeCircuitHash         "$(jq -r '.finalize.circuit_hash'            "${JSON}")"
update_go_hash FinalizeProvingKeyHash      "$(jq -r '.finalize.proving_key_hash'         "${JSON}")"
update_go_hash FinalizeVerificationKeyHash "$(jq -r '.finalize.verifying_key_hash'       "${JSON}")"

update_go_hash PartialDecryptCircuitHash         "$(jq -r '.partialdecrypt.circuit_hash'        "${JSON}")"
update_go_hash PartialDecryptProvingKeyHash      "$(jq -r '.partialdecrypt.proving_key_hash'     "${JSON}")"
update_go_hash PartialDecryptVerificationKeyHash "$(jq -r '.partialdecrypt.verifying_key_hash'   "${JSON}")"

update_go_hash DecryptCombineCircuitHash         "$(jq -r '.decryptcombine.circuit_hash'        "${JSON}")"
update_go_hash DecryptCombineProvingKeyHash      "$(jq -r '.decryptcombine.proving_key_hash'     "${JSON}")"
update_go_hash DecryptCombineVerificationKeyHash "$(jq -r '.decryptcombine.verifying_key_hash'   "${JSON}")"

update_go_hash RevealSubmitCircuitHash         "$(jq -r '.revealsubmit.circuit_hash'         "${JSON}")"
update_go_hash RevealSubmitProvingKeyHash      "$(jq -r '.revealsubmit.proving_key_hash'      "${JSON}")"
update_go_hash RevealSubmitVerificationKeyHash "$(jq -r '.revealsubmit.verifying_key_hash'    "${JSON}")"

update_go_hash RevealShareCircuitHash         "$(jq -r '.revealshare.circuit_hash'          "${JSON}")"
update_go_hash RevealShareProvingKeyHash      "$(jq -r '.revealshare.proving_key_hash'       "${JSON}")"
update_go_hash RevealShareVerificationKeyHash "$(jq -r '.revealshare.verifying_key_hash'     "${JSON}")"

echo "Updated ${CONFIG} with hashes from ${JSON}"

# ── Solidity verifier wrapper contracts ────────────────────────────────────

update_sol_pk_hash "${VERIFIERS_DIR}/ContributionVerifier.sol"  "$(jq -r '.contribution.proving_key_hash'   "${JSON}")"
update_sol_pk_hash "${VERIFIERS_DIR}/FinalizeVerifier.sol"      "$(jq -r '.finalize.proving_key_hash'       "${JSON}")"
update_sol_pk_hash "${VERIFIERS_DIR}/PartialDecryptVerifier.sol" "$(jq -r '.partialdecrypt.proving_key_hash' "${JSON}")"
update_sol_pk_hash "${VERIFIERS_DIR}/DecryptCombineVerifier.sol" "$(jq -r '.decryptcombine.proving_key_hash' "${JSON}")"
update_sol_pk_hash "${VERIFIERS_DIR}/RevealShareVerifier.sol"   "$(jq -r '.revealshare.proving_key_hash'    "${JSON}")"
update_sol_pk_hash "${VERIFIERS_DIR}/RevealSubmitVerifier.sol"  "$(jq -r '.revealsubmit.proving_key_hash'   "${JSON}")"

echo "Updated Solidity verifier wrapper contracts in ${VERIFIERS_DIR}/"
