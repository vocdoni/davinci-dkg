#!/usr/bin/env bash
# Self-check for deploy_all.sh with stub cast and forge: the RPC URL's API key
# (path, query, userinfo) reaches neither the script's output nor the
# generated address file; its scheme://host still identifies the chain.
set -euo pipefail

here=$(cd "$(dirname "$0")" && pwd)
tmp=$(mktemp -d)
trap 'rm -rf "$tmp"' EXIT
fail() { echo "[deploy_all.test] FAIL: $*"; exit 1; }

# A copy in its own tree, so no .env next to the real script is loaded.
mkdir -p "$tmp/solidity" "$tmp/bin"
cp "$here/deploy_all.sh" "$tmp/solidity/"
cat >"$tmp/bin/cast" <<'STUB'
#!/usr/bin/env bash
echo 0x00000000000000000000000000000000000000d1
STUB
cat >"$tmp/bin/forge" <<'STUB'
#!/usr/bin/env bash
n=1
for name in ContributionVerifier FinalizeVerifier PartialDecryptVerifier DecryptCombineVerifier DKGRegistry DKGManager DKGAppManager; do
	printf '  %s deployed at: 0x%040d\n' "$name" "$n"
	n=$((n + 1))
done
STUB
chmod +x "$tmp/bin/cast" "$tmp/bin/forge"

RPC_URL='https://deployer:DEPLOYPASS@rpc.example/v2/DEPLOYPATHKEY?apikey=DEPLOYQUERYKEY' \
	CHAIN_ID=31337 PRIVATE_KEY=0x01 MIN_THRESHOLD=2 MIN_COMMITTEE_SIZE=3 MAX_LOTTERY_ALPHA_BPS=20000 SKIP_TESTS=1 \
	PATH="$tmp/bin:$PATH" bash "$tmp/solidity/deploy_all.sh" >"$tmp/out" 2>&1 || { cat "$tmp/out"; fail "deploy_all.sh failed"; }

for secret in deployer DEPLOYPASS DEPLOYPATHKEY DEPLOYQUERYKEY; do
	if grep -q "$secret" "$tmp/out" "$tmp/solidity/.last_deployed_addresses.env"; then
		fail "$secret leaked"
	fi
done
grep -q 'https://rpc.example' "$tmp/out" || fail "the output does not name the RPC host"
grep -q 'https://rpc.example' "$tmp/solidity/.last_deployed_addresses.env" || fail "the address file does not name the RPC host"
grep -q '^MANAGER=0x' "$tmp/solidity/.last_deployed_addresses.env" || fail "no manager address written"
echo "[deploy_all.test] ok"
