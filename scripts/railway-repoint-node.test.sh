#!/usr/bin/env bash
# Self-check for railway-repoint-node.sh against a fake Railway GraphQL server: --plan sends no
# mutation; --apply sets the image, puts the private RPC first, deletes a preset-overriding
# variable, keeps everything else and deploys once; --rollback restores the saved state; no
# secret (token, key, private RPC, an old RPC value) ever reaches the output.
set -euo pipefail

here=$(cd "$(dirname "$0")" && pwd)
tmp=$(mktemp -d)
trap 'kill "${srv:-0}" 2>/dev/null || true; rm -rf "$tmp"' EXIT
fail() { echo "[railway-repoint-node.test] FAIL: $*"; exit 1; }

printf 'tok-SECRET-123\n' >"$tmp/token"
printf 'https://private.example/KEY-SECRET-456\n' >"$tmp/rpcs"
cat >"$tmp/server.py" <<'PY'
import json, os, sys
from http.server import BaseHTTPRequestHandler, HTTPServer
state = {"image": "ghcr.io/vocdoni/davinci-dkg:v0.10.0",
         "vars": {"DAVINCI_DKG_PRIVKEY": "0xPRIVKEY-SECRET-789", "DAVINCI_DKG_WEB3_RPC": "https://old.example/OLD-SECRET",
                  "DAVINCI_DKG_MANAGER": "0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF", "DAVINCI_DKG_POLL_INTERVAL": "30s",
                  "DAVINCI_DKG_EPOCH_POLICY_COMMITTEE_SIZE": "3"}}
log = open(sys.argv[2], "a")
P, E, S = "142f30d5-0000-4000-8000-000000000001", "52016ac4-0000-4000-8000-000000000002", "1151656f-0000-4000-8000-000000000003"
class H(BaseHTTPRequestHandler):
    def log_message(self, *a): pass
    def do_POST(self):
        assert self.headers["Authorization"] == "Bearer tok-SECRET-123"
        req = json.loads(self.rfile.read(int(self.headers["Content-Length"])))
        q, v = req["query"], req["variables"]
        if q.lstrip().startswith("mutation"):
            op = next(o for o in ("serviceInstanceUpdate", "variableCollectionUpsert", "variableDelete", "serviceInstanceDeployV2") if o in q)
            log.write(json.dumps({"op": op, "v": v}) + "\n"); log.flush()
        if "projects {" in q:
            data = {"projects": {"edges": [{"node": {"id": P, "name": "davinci-dkg-gnosis"}}]}}
        elif "project(id" in q:
            data = {"project": {"name": "davinci-dkg-gnosis",
                    "environments": {"edges": [{"node": {"id": E, "name": "production"}}]},
                    "services": {"edges": [{"node": {"id": S, "name": "dkg-node1"}}]},
                    "volumes": {"edges": [{"node": {"name": "v", "volumeInstances": {"edges": [
                        {"node": {"serviceId": S, "environmentId": E, "mountPath": "/app/run"}}]}}}]}}}
        elif "serviceInstance(" in q:
            data = {"serviceInstance": {"source": {"image": state["image"]}}}
        elif "variables(" in q:
            data = {"variables": state["vars"]}
        elif "serviceInstanceUpdate" in q:
            state["image"] = v["input"]["source"]["image"]; data = {"serviceInstanceUpdate": True}
        elif "variableCollectionUpsert" in q:
            assert v["input"]["skipDeploys"] is True
            state["vars"].update(v["input"]["variables"]); data = {"variableCollectionUpsert": True}
        elif "variableDelete" in q:
            state["vars"].pop(v["input"]["name"]); data = {"variableDelete": True}
        elif "serviceInstanceDeployV2" in q:
            if os.path.exists(os.path.dirname(sys.argv[2]) + "/fail-deploy"):
                body = json.dumps({"errors": [{"message": "deploy refused"}]}).encode()
                self.send_response(200); self.send_header("Content-Type", "application/json"); self.end_headers(); self.wfile.write(body)
                return
            data = {"serviceInstanceDeployV2": "dep-1"}
        else:
            data = None
        body = json.dumps({"data": data}).encode()
        self.send_response(200); self.send_header("Content-Type", "application/json"); self.end_headers(); self.wfile.write(body)
    def do_GET(self):
        body = json.dumps(state).encode()
        self.send_response(200); self.end_headers(); self.wfile.write(body)
HTTPServer(("127.0.0.1", int(sys.argv[1])), H).serve_forever()
PY
port=$((30000 + RANDOM % 20000))
python3 "$tmp/server.py" "$port" "$tmp/mutations" &
srv=$!
for _ in $(seq 50); do curl -s "http://127.0.0.1:$port/" >/dev/null 2>&1 && break; sleep 0.1; done

run() {
	RAILWAY_API=http://127.0.0.1:$port RAILWAY_TOKEN_FILE=$tmp/token RAILWAY_PROJECT=davinci-dkg-gnosis \
		EXPECT_PROJECT_PREFIX=142f30d5 EXPECT_ENVIRONMENT_PREFIX=52016ac4 EXPECT_SERVICE_PREFIX=1151656f \
		IMAGE=ghcr.io/vocdoni/davinci-dkg:v0.11.0 NETWORK=optimism RPC_FILE=$tmp/rpcs STATE_DIR=$tmp/state \
		TMPDIR=$tmp bash "$here/railway-repoint-node.sh" dkg-node1 "$@" >>"$tmp/out" 2>&1
}
state() { curl -s "http://127.0.0.1:$port/"; }
field() { jq -r "$2" <<<"$1"; }

: >"$tmp/mutations"
run --plan || { cat "$tmp/out"; fail "plan failed"; }
[ ! -s "$tmp/mutations" ] || fail "--plan sent a mutation"
grep -q "plan: delete DAVINCI_DKG_EPOCH_POLICY_COMMITTEE_SIZE DAVINCI_DKG_MANAGER" "$tmp/out" ||
	fail "plan does not delete the manager and policy overrides"

run --apply || { cat "$tmp/out"; fail "apply failed"; }
s=$(state)
[ "$(field "$s" '.image')" = ghcr.io/vocdoni/davinci-dkg:v0.11.0 ] || fail "image not set"
[ "$(field "$s" '.vars.DAVINCI_DKG_NETWORK')" = optimism ] || fail "network not set"
[ "$(field "$s" '.vars.DAVINCI_DKG_WEB3_RPC')" = "https://private.example/KEY-SECRET-456,https://mainnet.optimism.io,https://optimism-rpc.publicnode.com" ] || fail "rpc list"
[ "$(field "$s" '.vars | has("DAVINCI_DKG_MANAGER")')" = false ] || fail "manager override kept"
[ "$(field "$s" '.vars | has("DAVINCI_DKG_EPOCH_POLICY_COMMITTEE_SIZE")')" = false ] || fail "fixed policy kept"
[ "$(field "$s" '.vars.DAVINCI_DKG_PRIVKEY')" = 0xPRIVKEY-SECRET-789 ] || fail "operator key touched"
[ "$(field "$s" '.vars.DAVINCI_DKG_POLL_INTERVAL')" = 30s ] || fail "other variable touched"
order=$(jq -r .op "$tmp/mutations" | tr '\n' ' ')
# The image goes first and the one deploy last, so no deployment runs a half-moved service.
[ "$order" = "serviceInstanceUpdate variableCollectionUpsert variableDelete variableDelete serviceInstanceDeployV2 " ] ||
	fail "mutation order: $order"
[ "$(stat -c %a "$tmp"/state/*.json)" = 600 ] || fail "saved state is not 0600"

run --rollback || { cat "$tmp/out"; fail "rollback failed"; }
s=$(state)
[ "$(field "$s" '.image')" = ghcr.io/vocdoni/davinci-dkg:v0.10.0 ] || fail "image not restored"
[ "$(field "$s" '.vars.DAVINCI_DKG_WEB3_RPC')" = https://old.example/OLD-SECRET ] || fail "rpc not restored"
[ "$(field "$s" '.vars.DAVINCI_DKG_MANAGER')" = 0xC6Fb38c42ed3FB35D363a702218746d5C7Da36BF ] || fail "manager not restored"
[ "$(field "$s" '.vars.DAVINCI_DKG_EPOCH_POLICY_COMMITTEE_SIZE')" = 3 ] || fail "policy not restored"
[ "$(field "$s" '.vars | has("DAVINCI_DKG_NETWORK")')" = false ] || fail "network not removed on rollback"

# A second --apply (the natural retry after a partial failure) must not overwrite the saved
# state with the already-moved one: it is refused before any mutation, and --rollback still
# restores the original. A successful rollback sets the saved state aside, so a later
# --apply can run again.
rm -f "$tmp"/state/*.json
run --apply || { cat "$tmp/out"; fail "apply (twice case) failed"; }
: >"$tmp/mutations"
if run --apply; then fail "a second --apply overwrote the saved state"; fi
grep -q "saved state exists" "$tmp/out" || { cat "$tmp/out"; fail "no saved-state refusal message"; }
[ ! -s "$tmp/mutations" ] || fail "the refused second --apply mutated"
run --rollback || { cat "$tmp/out"; fail "rollback after a refused second apply failed"; }
s=$(state)
[ "$(field "$s" '.image')" = ghcr.io/vocdoni/davinci-dkg:v0.10.0 ] || fail "apply twice: image not restored"
[ "$(field "$s" '.vars.DAVINCI_DKG_WEB3_RPC')" = https://old.example/OLD-SECRET ] || fail "apply twice: rpc not restored"
[ "$(field "$s" '.vars | has("DAVINCI_DKG_NETWORK")')" = false ] || fail "apply twice: still on the new network"
ls "$tmp"/state/*.json >/dev/null 2>&1 && fail "the rolled-back state is still in place"
ls "$tmp"/state/*.rolled-back-* >/dev/null 2>&1 || fail "the rolled-back state was not kept aside"
run --apply || { cat "$tmp/out"; fail "--apply after a rollback was refused"; }
# A --rollback whose deploy fails must say so and keep the saved state for another try.
touch "$tmp/fail-deploy"
if run --rollback; then fail "a failed rollback deploy exited 0"; fi
ls "$tmp"/state/*.json >/dev/null 2>&1 || fail "a failed rollback set the saved state aside"
rm "$tmp/fail-deploy"
run --rollback || { cat "$tmp/out"; fail "second rollback failed"; }

# Wrong prefix: refused before any mutation.
: >"$tmp/mutations"
if RAILWAY_API=http://127.0.0.1:$port RAILWAY_TOKEN_FILE=$tmp/token RAILWAY_PROJECT=davinci-dkg-gnosis \
	EXPECT_SERVICE_PREFIX=deadbeef IMAGE=x NETWORK=optimism RPC_FILE=$tmp/rpcs STATE_DIR=$tmp/state TMPDIR=$tmp \
	bash "$here/railway-repoint-node.sh" dkg-node1 --apply >>"$tmp/out" 2>&1; then fail "accepted a wrong service prefix"; fi
[ ! -s "$tmp/mutations" ] || fail "mutated with a wrong prefix"

for secret in tok-SECRET-123 KEY-SECRET-456 PRIVKEY-SECRET-789 OLD-SECRET; do
	! grep -q "$secret" "$tmp/out" || fail "output shows $secret"
done
echo "[railway-repoint-node.test] ok"
