#!/usr/bin/env bash
# Self-check for move-optimism-preset.sh: runs it on a temporary copy of the
# files it edits, from the placeholder and again from an earlier address, and
# checks that only the optimism blocks and the OP rows change. The working
# tree is never touched.
set -euo pipefail

here=$(cd "$(dirname "$0")" && pwd)
root=$(cd "$here/.." && pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
export PATH="$HOME/.foundry/bin:$PATH"
command -v cast >/dev/null || { echo "[move-optimism-preset.test] skipped: cast not found"; exit 0; }

mkdir -p "$tmp/scripts" "$tmp/config" "$tmp/sdk/src" "$tmp/docs"
cp "$here/move-optimism-preset.sh" "$tmp/scripts/"
cp "$root/config/networks.go" "$tmp/config/"
cp "$root/sdk/src/networks.ts" "$tmp/sdk/src/"
cp "$root/docs/deployments.md" "$tmp/docs/"
cp "$root/README.md" "$tmp/"
before=$(cd "$root" && git status --porcelain 2>/dev/null | md5sum)

addrs() { # <file> <hex digit>
	local d=$2 f=$1 n
	: >"$f"
	for n in MANAGER REGISTRY APP_MANAGER CONTRIBUTION_VERIFIER FINALIZE_VERIFIER PARTIAL_DECRYPT_VERIFIER \
		DECRYPT_COMBINE_VERIFIER; do
		printf '%s=0x%s\n' "$n" "$(yes "$d" | head -n 40 | tr -d "\n")" >>"$f"
		d=$(printf '%x' $(((0x$d + 1) % 16)))
	done
}
fail() { echo "FAIL: $*"; exit 1; }
has() { grep -qF -- "$2" "$1" || fail "$1 lacks: $2"; }
hasnt() { ! grep -qF -- "$2" "$1" || fail "$1 still has: $2"; }

run() { (cd "$tmp" && bash scripts/move-optimism-preset.sh "$@" >"$tmp/out.log" 2>&1) || { cat "$tmp/out.log"; fail "script exited non-zero"; }; }

# 1. From the placeholder.
addrs "$tmp/a.env" a
run "$tmp/a.env" 158000000 2026-11-02
mgr=$(cast to-check-sum-address 0xaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa)
grep -Eq "Manager: +common.HexToAddress\(\"$mgr\"\)," "$tmp/config/networks.go" || fail "networks.go manager"
has "$tmp/config/networks.go" "158_000_000, // DKGManager deployment block on OP Mainnet (2026-11-02)"
hasnt "$tmp/config/networks.go" PLACEHOLDER
has "$tmp/config/networks.go" 'common.HexToAddress("0xc73b7a868eca6ac7e3e647e2665aa16a793cf551")'
has "$tmp/config/networks.go" "StartBlock: 11_668_198,"
has "$tmp/sdk/src/networks.ts" "managerAddress: '$mgr',"
has "$tmp/sdk/src/networks.ts" "startBlock: 158_000_000n,"
hasnt "$tmp/sdk/src/networks.ts" PLACEHOLDER
has "$tmp/sdk/src/networks.ts" "managerAddress: '0xc73b7a868eca6ac7e3e647e2665aa16a793cf551',"
has "$tmp/docs/deployments.md" '| DKGManager | `'"$mgr"'` |'
has "$tmp/docs/deployments.md" '| DKGAppManager | `0xcccccccccccccccccccccccccccccccccccccccc` |'
has "$tmp/docs/deployments.md" '| DKGRegistry | `0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb` |'
has "$tmp/docs/deployments.md" '| ContributionVerifier | `0xdddddddddddddddddddddddddddddddddddddddd` |'
has "$tmp/docs/deployments.md" '| FinalizeVerifier | `0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee` |'
has "$tmp/docs/deployments.md" '| PartialDecryptVerifier | `0xffffffffffffffffffffffffffffffffffffffff` |'
has "$tmp/docs/deployments.md" '| DecryptCombineVerifier | `0x0000000000000000000000000000000000000000` |'
has "$tmp/docs/deployments.md" "Deployment block: 158,000,000 (2026-11-02)."
hasnt "$tmp/docs/deployments.md" 'pending deployment'
# Other sections keep their rows.
has "$tmp/docs/deployments.md" '| DKGManager | `0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF` |'
has "$tmp/docs/deployments.md" '| DKGManager | `0xc73b7a868eca6ac7e3e647e2665aa16a793cf551` |'
has "$tmp/README.md" "| OP Mainnet (default) | 10 | \`$mgr\` |"
has "$tmp/README.md" '| Gnosis Chain (retired) | 100 | `0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF` |'
has "$tmp/README.md" '`0xc73b7a868eca6ac7e3e647e2665aa16a793cf551`'

# 2. From an earlier address: replaced, with a retirement hint.
addrs "$tmp/b.env" 1
run "$tmp/b.env" 159000000 2026-12-01
new=$(cast to-check-sum-address 0x1111111111111111111111111111111111111111)
grep -Eq "Manager: +common.HexToAddress\(\"$new\"\)," "$tmp/config/networks.go" || fail "second networks.go manager"
hasnt "$tmp/config/networks.go" "$mgr"
has "$tmp/config/networks.go" "159_000_000, // DKGManager deployment block on OP Mainnet (2026-12-01)"
has "$tmp/sdk/src/networks.ts" "managerAddress: '$new',"
has "$tmp/sdk/src/networks.ts" "startBlock: 159_000_000n,"
hasnt "$tmp/sdk/src/networks.ts" "$mgr"
has "$tmp/docs/deployments.md" "Deployment block: 159,000,000 (2026-12-01)."
hasnt "$tmp/docs/deployments.md" "$mgr"
has "$tmp/README.md" "| OP Mainnet (default) | 10 | \`$new\` |"
has "$tmp/out.log" "record the replaced manager $mgr as retired"

# 3. Bad input is refused.
if (cd "$tmp" && bash scripts/move-optimism-preset.sh "$tmp/a.env" notanumber >/dev/null 2>&1); then fail "accepted a bad block"; fi

# The working tree is untouched.
[ "$before" = "$(cd "$root" && git status --porcelain 2>/dev/null | md5sum)" ] || fail "working tree changed"

echo "[move-optimism-preset.test] ok"
