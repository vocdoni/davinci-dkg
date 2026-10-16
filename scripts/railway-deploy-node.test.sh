#!/usr/bin/env bash
# Self-check for railway-deploy-node.sh against a fake Railway GraphQL server: the service gets
# the operator key from the key file (the plain list and Foundry 1.8's `cast wallet new --json`
# wrapper), the network, the RPC list with the private endpoint from RPC_FILE first, the poll
# interval and the volume paths; then a volume, the ALWAYS restart policy and one deploy, in that
# order. No secret (token, key, private RPC) reaches the output, and the request bodies carrying
# them are 0600 even under a permissive umask.
set -euo pipefail

here=$(cd "$(dirname "$0")" && pwd)
tmp=$(mktemp -d)
trap 'kill "${srv:-0}" 2>/dev/null || true; rm -rf "$tmp"' EXIT
fail() { echo "[railway-deploy-node.test] FAIL: $*"; exit 1; }

printf 'tok-SECRET-123\n' >"$tmp/token"
printf 'https://private.example/KEY-SECRET-456\n' >"$tmp/rpcs"
printf '[{"address": "0x00000000000000000000000000000000000000a1", "private_key": "0xPRIVKEY-SECRET-1"}]\n' >"$tmp/plain.json"
printf '{"success": true, "data": [{"address": "0x00000000000000000000000000000000000000a2", "private_key": "0xPRIVKEY-SECRET-2", "public_key": "0x04"}]}\n' >"$tmp/wrapped.json"
cat >"$tmp/server.py" <<'PY'
import glob, json, os, sys
from http.server import BaseHTTPRequestHandler, HTTPServer
log = open(sys.argv[2], "a")
class H(BaseHTTPRequestHandler):
    def log_message(self, *a): pass
    def do_POST(self):
        assert self.headers["Authorization"] == "Bearer tok-SECRET-123"
        req = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
        q, v = req["query"], req["variables"]
        op = next(o for o in ("serviceCreate", "volumeCreate", "serviceInstanceUpdate", "serviceInstanceDeployV2") if o in q)
        if op == "volumeCreate" and os.path.exists(sys.argv[3] + "/fail-volume"):
            log.write(json.dumps({"op": op, "v": v, "modes": []}) + "\n"); log.flush()
            body = json.dumps({"errors": [{"message": "You are creating volumes too quickly"}]}).encode()
            self.send_response(200); self.send_header("Content-Type", "application/json"); self.end_headers(); self.wfile.write(body)
            return
        modes = sorted({oct(os.stat(f).st_mode & 0o777) for f in glob.glob(sys.argv[3] + "/*.body")})
        log.write(json.dumps({"op": op, "v": v, "modes": modes}) + "\n"); log.flush()
        data = {"serviceCreate": {"id": "svc-1", "name": v.get("input", {}).get("name")},
                "volumeCreate": {"id": "vol-1"}, "serviceInstanceUpdate": True, "serviceInstanceDeployV2": "dep-1"}
        body = json.dumps({"data": {op: data[op]}}).encode()
        self.send_response(200); self.send_header("Content-Type", "application/json"); self.end_headers(); self.wfile.write(body)
HTTPServer(("127.0.0.1", int(sys.argv[1])), H).serve_forever()
PY
port=$((30000 + RANDOM % 20000))
python3 "$tmp/server.py" "$port" "$tmp/calls" "$tmp" &
srv=$!
for _ in $(seq 50); do python3 -c "import socket,sys; socket.create_connection(('127.0.0.1', $port), 1)" 2>/dev/null && break; sleep 0.1; done

for key in plain wrapped; do
	: >"$tmp/calls"
	RAILWAY_API=http://127.0.0.1:$port RAILWAY_TOKEN_FILE=$tmp/token RAILWAY_PROJECT_ID=p-1 RAILWAY_ENVIRONMENT_ID=e-1 \
		NODE_NAME=dkg-node4 NODE_KEY_FILE=$tmp/$key.json IMAGE=ghcr.io/vocdoni/davinci-dkg:v0.11.0 NETWORK=optimism \
		RPC_FILE=$tmp/rpcs PUBLIC_RPCS=https://mainnet.optimism.io,https://optimism-rpc.publicnode.com TMPDIR=$tmp \
		bash -c 'umask 002; exec bash "$0"' "$here/railway-deploy-node.sh" >"$tmp/out-$key" 2>&1 || { cat "$tmp/out-$key"; fail "$key: deploy failed"; }
	[ "$(jq -r .op "$tmp/calls" | tr '\n' ' ')" = "serviceCreate volumeCreate serviceInstanceUpdate serviceInstanceDeployV2 " ] ||
		fail "$key: call order $(jq -r .op "$tmp/calls" | tr '\n' ' ')"
	[ "$(jq -c .modes "$tmp/calls" | sort -u)" = '["0o600"]' ] || fail "$key: request body modes $(jq -c .modes "$tmp/calls" | sort -u | tr '\n' ' ')"
	create=$(head -n1 "$tmp/calls")
	n=${key/plain/1}; n=${n/wrapped/2}
	[ "$(jq -r .v.input.variables.DAVINCI_DKG_PRIVKEY <<<"$create")" = "0xPRIVKEY-SECRET-$n" ] || fail "$key: operator key"
	[ "$(jq -r .v.input.variables.DAVINCI_DKG_NETWORK <<<"$create")" = optimism ] || fail "$key: network"
	[ "$(jq -r .v.input.variables.DAVINCI_DKG_WEB3_RPC <<<"$create")" = \
		"https://private.example/KEY-SECRET-456,https://mainnet.optimism.io,https://optimism-rpc.publicnode.com" ] || fail "$key: rpc list"
	[ "$(jq -r .v.input.variables.DAVINCI_DKG_DATADIR <<<"$create")" = /app/run/data ] || fail "$key: datadir"
	[ "$(jq -r .v.input.source.image <<<"$create")" = ghcr.io/vocdoni/davinci-dkg:v0.11.0 ] || fail "$key: image"
	[ "$(sed -n 3p "$tmp/calls" | jq -r .v.input.restartPolicyType)" = ALWAYS ] || fail "$key: restart policy"
	grep -q "operator: 0x00000000000000000000000000000000000000a$n" "$tmp/out-$key" || fail "$key: operator address not shown"
	for secret in tok-SECRET-123 KEY-SECRET-456 PRIVKEY-SECRET; do
		! grep -q "$secret" "$tmp/out-$key" || fail "$key: output shows $secret"
	done
done
# A Railway error on any step stops the script before the deploy (a volume that failed must not
# leave a node deployed without one).
: >"$tmp/calls"; touch "$tmp/fail-volume"
if RAILWAY_API=http://127.0.0.1:$port RAILWAY_TOKEN_FILE=$tmp/token RAILWAY_PROJECT_ID=p-1 RAILWAY_ENVIRONMENT_ID=e-1 \
	NODE_NAME=dkg-node5 NODE_KEY_FILE=$tmp/plain.json NETWORK=optimism RPC_FILE=$tmp/rpcs TMPDIR=$tmp \
	bash "$here/railway-deploy-node.sh" >"$tmp/out-fail" 2>&1; then fail "a volume error did not stop the script"; fi
[ "$(jq -r .op "$tmp/calls" | tr '\n' ' ')" = "serviceCreate volumeCreate " ] || fail "calls after the volume error: $(jq -r .op "$tmp/calls" | tr '\n' ' ')"
echo "[railway-deploy-node.test] ok"
