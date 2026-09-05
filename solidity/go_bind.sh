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
generate_binding "$OUT_DIR/DKGAppManager.sol/DKGAppManager.json" golangtypes DKGAppManager "$BINDINGS_DIR/dkgappmanager.go"

# DKGTypesPoint may be emitted by both the DKGManager and DKGAppManager
# bindings; strip the duplicate from the app-manager binding only when the
# manager binding defines it, so the package compiles either way.
python3 - "$BINDINGS_DIR/dkgappmanager.go" "$BINDINGS_DIR/dkgmanager.go" <<'PY'
import re, sys
path, manager = sys.argv[1], sys.argv[2]
if "type DKGTypesPoint struct" not in open(manager).read():
    sys.exit(0)
src = open(path).read()
pattern = re.compile(
    r"// DKGTypesPoint is an auto generated low-level Go binding around an user-defined struct\.\n"
    r"type DKGTypesPoint struct \{\n[^}]*\}\n\n",
)
new = pattern.sub("", src, count=1)
open(path, "w").write(new)
PY
generate_binding "$OUT_DIR/ContributionVerifier.sol/ContributionVerifier.json" golangtypes ContributionVerifier "$BINDINGS_DIR/contributionverifier.go"
generate_binding "$OUT_DIR/PoolKeyVerifier.sol/PoolKeyVerifier.json" golangtypes PoolKeyVerifier "$BINDINGS_DIR/poolkeyverifier.go"
generate_binding "$OUT_DIR/PartialDecryptVerifier.sol/PartialDecryptVerifier.json" golangtypes PartialDecryptVerifier "$BINDINGS_DIR/partialdecryptverifier.go"
generate_binding "$OUT_DIR/DecryptCombineVerifier.sol/DecryptCombineVerifier.json" golangtypes DecryptCombineVerifier "$BINDINGS_DIR/decryptcombineverifier.go"

echo "go bindings written to $BINDINGS_DIR"
