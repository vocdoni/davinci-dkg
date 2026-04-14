#!/usr/bin/env bash
set -euo pipefail

ROOT_DIR="$(cd "$(dirname "$0")/.." && pwd)"
SOLIDITY_DIR="$ROOT_DIR/solidity"
OUT_DIR="$SOLIDITY_DIR/out"
BINDINGS_DIR="$SOLIDITY_DIR/golang-types"
ABIGEN_CMD=(go run github.com/ethereum/go-ethereum/cmd/abigen@v1.17.1)

mkdir -p "$BINDINGS_DIR"

generate_binding() {
  local artifact_json="$1"
  local package_name="$2"
  local type_name="$3"
  local output_file="$4"

  local abi_file
  local bin_file
  abi_file="$(mktemp)"
  bin_file="$(mktemp)"

  jq '.abi' "$artifact_json" > "$abi_file"
  jq -r '.bytecode.object // .deployedBytecode.object // ""' "$artifact_json" > "$bin_file"

  "${ABIGEN_CMD[@]}" \
    --abi "$abi_file" \
    --bin "$bin_file" \
    --pkg "$package_name" \
    --type "$type_name" \
    --out "$output_file"

  rm -f "$abi_file" "$bin_file"
}

generate_binding "$OUT_DIR/DKGRegistry.sol/DKGRegistry.json" golangtypes DKGRegistry "$BINDINGS_DIR/dkgregistry.go"
generate_binding "$OUT_DIR/DKGManager.sol/DKGManager.json" golangtypes DKGManager "$BINDINGS_DIR/dkgmanager.go"
generate_binding "$OUT_DIR/ContributionVerifier.sol/ContributionVerifier.json" golangtypes ContributionVerifier "$BINDINGS_DIR/contributionverifier.go"
generate_binding "$OUT_DIR/FinalizeVerifier.sol/FinalizeVerifier.json" golangtypes FinalizeVerifier "$BINDINGS_DIR/finalizeverifier.go"
generate_binding "$OUT_DIR/PartialDecryptVerifier.sol/PartialDecryptVerifier.json" golangtypes PartialDecryptVerifier "$BINDINGS_DIR/partialdecryptverifier.go"
generate_binding "$OUT_DIR/DecryptCombineVerifier.sol/DecryptCombineVerifier.json" golangtypes DecryptCombineVerifier "$BINDINGS_DIR/decryptcombineverifier.go"
generate_binding "$OUT_DIR/RevealSubmitVerifier.sol/RevealSubmitVerifier.json" golangtypes RevealSubmitVerifier "$BINDINGS_DIR/revealsubmitverifier.go"
generate_binding "$OUT_DIR/RevealShareVerifier.sol/RevealShareVerifier.json" golangtypes RevealShareVerifier "$BINDINGS_DIR/revealshareverifier.go"

echo "go bindings written to $BINDINGS_DIR"
