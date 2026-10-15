#!/usr/bin/env bash
# Fill the built-in `optimism` preset with a deployment: the node and dkgapp
# (config/networks.go), the SDK (sdk/src/networks.ts), the OP Mainnet tables
# of docs/deployments.md and the README networks row. Run it once, right after
# solidity/deploy_all.sh, then review the diff and run the checks it prints.
# It edits only inside those blocks and sections (the zero address and the
# Gnosis rows appear elsewhere), and works whether the preset still holds the
# placeholder or an earlier deployment. The explorer config is not touched.
#
# Usage:
#   scripts/move-optimism-preset.sh <addresses.env> <deploy-block> [YYYY-MM-DD]
#
#   addresses.env  the file deploy_all.sh writes (solidity/.last_deployed_addresses.env)
#   deploy-block   block number of the DKGManager deployment transaction
#   YYYY-MM-DD     deployment date for the comments (default: today, UTC)
set -euo pipefail

ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$ROOT"

[ $# -ge 2 ] || { sed -n '12,17p' "$0"; exit 2; }
ADDRS=$1
NEW_BLOCK=$2
DATE=${3:-$(date -u +%Y-%m-%d)}
[[ $NEW_BLOCK =~ ^[0-9]+$ ]] || { echo "deploy-block must be a number" >&2; exit 2; }
[[ $DATE =~ ^[0-9]{4}-[0-9]{2}-[0-9]{2}$ ]] || { echo "date must be YYYY-MM-DD" >&2; exit 2; }

# shellcheck disable=SC1090
. "$ADDRS"
for v in MANAGER REGISTRY APP_MANAGER CONTRIBUTION_VERIFIER FINALIZE_VERIFIER PARTIAL_DECRYPT_VERIFIER \
	DECRYPT_COMBINE_VERIFIER; do
	[[ ${!v:-} =~ ^0x[0-9a-fA-F]{40}$ ]] || { echo "$v missing or malformed in $ADDRS" >&2; exit 2; }
done

command -v cast >/dev/null || { echo "cast (Foundry) is required: export PATH=$HOME/.foundry/bin:$PATH" >&2; exit 1; }
lower() { printf '%s' "$1" | tr 'A-F' 'a-f'; }
underscored() { printf '%s' "$1" | rev | sed -E 's/([0-9]{3})/\1_/g; s/_$//' | rev; }
commas() { printf '%s' "$1" | rev | sed -E 's/([0-9]{3})/\1,/g; s/,$//' | rev; }

NEW_MANAGER=$(cast to-check-sum-address "$MANAGER")
OLD_MANAGER=$(sed -n '/^\t"optimism": {/,/^\t},/s/.*HexToAddress("\(0x[0-9a-fA-F]*\)").*/\1/p' config/networks.go)
grep -q '^	"optimism": {' config/networks.go || { echo "no optimism preset in config/networks.go" >&2; exit 1; }
grep -q '^  optimism: {' sdk/src/networks.ts || { echo "no optimism preset in sdk/src/networks.ts" >&2; exit 1; }
grep -q '^## OP Mainnet' docs/deployments.md || { echo "no OP Mainnet section in docs/deployments.md" >&2; exit 1; }
echo "optimism preset: ${OLD_MANAGER:-placeholder}  ->  $NEW_MANAGER @ $NEW_BLOCK"

# inplace <file> <awk program>: rewrite a file through awk (variables via -v).
inplace() {
	local f=$1 tmp
	shift
	tmp=$(mktemp)
	awk "$@" "$f" >"$tmp"
	cat "$tmp" >"$f"
	rm -f "$tmp"
}

# config/networks.go: Manager and StartBlock inside the "optimism" block; the
# placeholder comment goes, a dated one replaces it.
inplace config/networks.go -v mgr="$NEW_MANAGER" -v blk="$(underscored "$NEW_BLOCK")" -v date="$DATE" '
	/^\t"optimism": \{/ { inb = 1 }
	inb && /^\t\},/ { inb = 0 }
	inb && /PLACEHOLDER/ { next }
	inb && /^\t\tManager:/ { print "\t\tManager: common.HexToAddress(\"" mgr "\"),"; next }
	inb && /^\t\tStartBlock:/ { print "\t\tStartBlock: " blk ", // DKGManager deployment block on OP Mainnet (" date ")"; next }
	{ print }'
command -v gofmt >/dev/null && gofmt -w config/networks.go

# sdk/src/networks.ts: managerAddress and startBlock inside the optimism block.
inplace sdk/src/networks.ts -v mgr="$NEW_MANAGER" -v blk="$(underscored "$NEW_BLOCK")" '
	/^  optimism: \{/ { inb = 1 }
	inb && /^  \},/ { inb = 0 }
	inb && /PLACEHOLDER/ { next }
	inb && /^    managerAddress:/ { print "    managerAddress: \x27" mgr "\x27,"; next }
	inb && /^    startBlock:/ { print "    startBlock: " blk "n,"; next }
	{ print }'

# docs/deployments.md: the table rows, the block and the status of the OP section.
inplace docs/deployments.md \
	-v mgr="$NEW_MANAGER" -v app="$(lower "$APP_MANAGER")" -v reg="$(lower "$REGISTRY")" \
	-v cv="$(lower "$CONTRIBUTION_VERIFIER")" -v fv="$(lower "$FINALIZE_VERIFIER")" \
	-v pv="$(lower "$PARTIAL_DECRYPT_VERIFIER")" -v dv="$(lower "$DECRYPT_COMBINE_VERIFIER")" \
	-v blk="$(commas "$NEW_BLOCK")" -v date="$DATE" '
	function row(name, addr) { print "| " name " | `" addr "` |" }
	/^## / { inb = ($0 == "## OP Mainnet") }
	inb && /^\| DKGManager \|/ { row("DKGManager", mgr); next }
	inb && /^\| DKGAppManager \|/ { row("DKGAppManager", app); next }
	inb && /^\| DKGRegistry \|/ { row("DKGRegistry", reg); next }
	inb && /^\| ContributionVerifier \|/ { row("ContributionVerifier", cv); next }
	inb && /^\| FinalizeVerifier \|/ { row("FinalizeVerifier", fv); next }
	inb && /^\| PartialDecryptVerifier \|/ { row("PartialDecryptVerifier", pv); next }
	inb && /^\| DecryptCombineVerifier \|/ { row("DecryptCombineVerifier", dv); next }
	inb && /^Deployment block:/ { print "Deployment block: " blk " (" date ")."; next }
	inb && /^Status:/ { print "Status: **live**."; next }
	{ print }'

# README.md: the OP Mainnet row of the networks table.
inplace README.md -v mgr="$NEW_MANAGER" '
	/^\| OP Mainnet / {
		n = split($0, c, "|")
		c[2] = " OP Mainnet (default) "
		c[4] = " `" mgr "` "
		out = ""
		for (i = 2; i < n; i++) out = out "|" c[i]
		print out "|"
		next
	}
	{ print }'

cat <<EOF

Now:
  git diff
  go test ./config ./node
  (cd sdk && npx -y pnpm@10 vitest run tests/networks.test.ts)
  make ui-test    # rebuilds the SDK first
  bash scripts/move-optimism-preset.test.sh
Then drop the "pending deployment" / "no DKGManager yet" notes that remain:
EOF
grep -rIn -e 'pending deployment' -e 'no DKGManager yet' --include='*.md' --include='*.example' --include='*.yml' \
	--exclude-dir=node_modules --exclude-dir=.git . || true
if [ -n "$OLD_MANAGER" ] && [ "$(lower "$OLD_MANAGER")" != "$(lower "$NEW_MANAGER")" ]; then
	echo "and record the replaced manager $OLD_MANAGER as retired in docs/deployments.md."
fi
