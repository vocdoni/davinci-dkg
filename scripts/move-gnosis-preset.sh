#!/usr/bin/env bash
# Point the built-in `gnosis` preset at a new deployment: the node and dkgapp
# (config/networks.go), the SDK (sdk/src/networks.ts), the explorer config and
# its fallbacks (ui/public/config.json, scripts/render-ui-config.sh, the DO App
# Platform spec), their tests and the docs. Run it once, right after
# solidity/deploy_all.sh, then review the diff and run the checks it prints.
#
# Usage:
#   scripts/move-gnosis-preset.sh <addresses.env> <deploy-block> [YYYY-MM-DD]
#
#   addresses.env  the file deploy_all.sh writes (solidity/.last_deployed_addresses.env)
#   deploy-block   block number of the DKGManager deployment transaction
#   YYYY-MM-DD     deployment date for the networks.go comment (default: today, UTC)
#
# The current preset (manager and start block) is read from config/networks.go
# and replaced wherever it appears, in every spelling the repo uses.
set -euo pipefail

ROOT=$(cd "$(dirname "$0")/.." && pwd)
cd "$ROOT"

[ $# -ge 2 ] || { sed -n '9,16p' "$0"; exit 2; }
ADDRS=$1
NEW_BLOCK=$2
DATE=${3:-$(date -u +%Y-%m-%d)}
[[ $NEW_BLOCK =~ ^[0-9]+$ ]] || { echo "deploy-block must be a number" >&2; exit 2; }

# shellcheck disable=SC1090
. "$ADDRS"
for v in MANAGER REGISTRY APP_MANAGER CONTRIBUTION_VERIFIER FINALIZE_VERIFIER PARTIAL_DECRYPT_VERIFIER \
	DECRYPT_COMBINE_VERIFIER; do
	[[ ${!v:-} =~ ^0x[0-9a-fA-F]{40}$ ]] || { echo "$v missing or malformed in $ADDRS" >&2; exit 2; }
done

command -v cast >/dev/null || { echo "cast (Foundry) is required: export PATH=$HOME/.foundry/bin:$PATH" >&2; exit 1; }
checksum() { cast to-check-sum-address "$1"; }
lower() { printf '%s' "$1" | tr 'A-F' 'a-f'; }
underscored() { printf '%s' "$1" | rev | sed -E 's/([0-9]{3})/\1_/g; s/_$//' | rev; }
commas() { printf '%s' "$1" | rev | sed -E 's/([0-9]{3})/\1,/g; s/,$//' | rev; }

OLD_MANAGER=$(sed -n '/"gnosis": {/,/},/s/.*HexToAddress("\(0x[0-9a-fA-F]\{40\}\)").*/\1/p' config/networks.go)
OLD_BLOCK=$(sed -n '/"gnosis": {/,/},/s/.*StartBlock: *\([0-9_]*\),.*/\1/p' config/networks.go | tr -d _)
[ -n "$OLD_MANAGER" ] && [ -n "$OLD_BLOCK" ] || { echo "cannot read the gnosis preset from config/networks.go" >&2; exit 1; }

NEW_MANAGER=$(checksum "$MANAGER")
echo "gnosis preset: $OLD_MANAGER @ $OLD_BLOCK  ->  $NEW_MANAGER @ $NEW_BLOCK"

FILES=(
	config/networks.go config/networks_test.go node/datadir_test.go
	sdk/src/networks.ts sdk/tests/networks.test.ts
	ui/public/config.json ui/README.md ui/.do/davinci-dkg-ui.yaml
	ui/src/pages/docs/docs.test.tsx ui/src/app/wagmi.test.ts
	scripts/render-ui-config.sh scripts/render-ui-config.test.sh Makefile
	README.md docs/deployments.md
)
for f in "${FILES[@]}"; do
	sed -i \
		-e "s/$OLD_MANAGER/$NEW_MANAGER/g" \
		-e "s/$(lower "$OLD_MANAGER")/$(lower "$NEW_MANAGER")/g" \
		-e "s/$(underscored "$OLD_BLOCK")/$(underscored "$NEW_BLOCK")/g" \
		-e "s/$(commas "$OLD_BLOCK")/$(commas "$NEW_BLOCK")/g" \
		-e "s/\b$OLD_BLOCK\b/$NEW_BLOCK/g" \
		"$f"
done

# networks.go dates the deployment in a comment.
sed -i "/\"gnosis\": {/,/},/s/(20[0-9-]*)/($DATE)/" config/networks.go

# docs/deployments.md: the Gnosis table rows of the contracts read from the manager.
row() { sed -i "/^## Gnosis Chain/,/^## /s/^| $1 | \`0x[0-9a-fA-F]*\` |/| $1 | \`$(lower "$2")\` |/" docs/deployments.md; }
row DKGAppManager "$APP_MANAGER"
row DKGRegistry "$REGISTRY"
row ContributionVerifier "$CONTRIBUTION_VERIFIER"
row FinalizeVerifier "$FINALIZE_VERIFIER"
row PartialDecryptVerifier "$PARTIAL_DECRYPT_VERIFIER"
row DecryptCombineVerifier "$DECRYPT_COMBINE_VERIFIER"

left=$(grep -rIl -e "$OLD_MANAGER" -e "$(lower "$OLD_MANAGER")" --exclude-dir=node_modules --exclude-dir=.git \
	--exclude-dir=dist --exclude-dir=new_dkg_paper.git . || true)
[ -z "$left" ] || { echo "still mentioning the old manager (review by hand):"; echo "$left"; }

cat <<EOF

Now:
  git diff
  go test ./config ./node -run 'Network|DataDir' && bash scripts/render-ui-config.test.sh
  (cd sdk && npx -y pnpm@10 vitest run tests/networks.test.ts)
  make ui-test    # rebuilds the SDK first: the explorer compares its config with the SDK preset
and record $OLD_MANAGER as retired in docs/deployments.md.
EOF
