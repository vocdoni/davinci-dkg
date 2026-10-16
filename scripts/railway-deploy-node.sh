#!/usr/bin/env bash
# Deploy one davinci-dkg node as a Railway service (public GHCR image, one
# volume, always-restart) through Railway's GraphQL API. See docs/node.md.
#
#   RAILWAY_TOKEN_FILE=railway-api-key \
#   RAILWAY_PROJECT_ID=<uuid> RAILWAY_ENVIRONMENT_ID=<uuid> \
#   NODE_NAME=dkg-node1 NODE_KEY_FILE=~/keys/node1.json \
#   scripts/railway-deploy-node.sh
#
# With NETWORK and RPC unset the node runs on its default network, OP
# Mainnet, through the public endpoints built into the binary. The older
# Sepolia testnet needs both, e.g. NETWORK=sepolia
# RPC=https://ethereum-sepolia-rpc.publicnode.com,https://1rpc.io/sepolia.
# RPC_FILE (its first line, a private endpoint) goes first, then PUBLIC_RPCS:
#   RPC_FILE=~/optimism-rpcs.txt PUBLIC_RPCS=https://mainnet.optimism.io,https://optimism-rpc.publicnode.com
#
# NODE_KEY_FILE holds `[{"address": "0x…", "private_key": "0x…"}]` or the
# `{"data": [...]}` that `cast wallet new --json` prints. The key and the
# private RPC travel only inside request bodies in 0600 temp files, and the
# token through a curl config file: none is an argument or printed.
set -euo pipefail
umask 077

: "${RAILWAY_TOKEN_FILE:=railway-api-key}"
: "${RAILWAY_PROJECT_ID:?project id (projectCreate output)}"
: "${RAILWAY_ENVIRONMENT_ID:?environment id (usually the production environment of the project)}"
: "${NODE_NAME:?service name, e.g. dkg-node9}"
: "${NODE_KEY_FILE:?operator key file}"
: "${IMAGE:=ghcr.io/vocdoni/davinci-dkg:latest}"
: "${NETWORK:=}"
: "${RPC:=}"
: "${RPC_FILE:=}"
: "${PUBLIC_RPCS:=}"
: "${POLL_INTERVAL:=30s}"
: "${MOUNT_PATH:=/app/run}"
API=${RAILWAY_API:-https://backboard.railway.com/graphql/v2}

vars=$(mktemp); chmod 600 "$vars"
auth=$(mktemp); chmod 600 "$auth"
trap 'rm -f "$vars" "$vars.body" "$auth"' EXIT
printf 'header = "Authorization: Bearer %s"\n' "$(tr -d '[:space:]' < "$RAILWAY_TOKEN_FILE")" > "$auth"

# gql QUERY [VARS_FILE]: post one GraphQL document, print the JSON response.
gql() {
	local q=$1 vf=${2:-}
	python3 - "$q" "$vf" <<'PY' > "$vars.body"
import sys, json
q, vf = sys.argv[1], sys.argv[2]
v = json.load(open(vf)) if vf else {}
print(json.dumps({"query": q, "variables": v}))
PY
	curl -sS --max-time 90 "$API" -K "$auth" -H 'Content-Type: application/json' --data @"$vars.body"
	rm -f "$vars.body"
}
# field RESPONSE PATH: extract a JSON path (dot separated) or fail loudly.
field() {
	python3 -c 'import sys,json; d=json.loads(sys.argv[1]);
[sys.exit("railway error: "+json.dumps(d["errors"])) for _ in [0] if d.get("errors")]
for k in sys.argv[2].split("."): d=d[k]
print(d)' "$1" "$2"
}

# 1. Service with its variables (deployed with them from the first start).
# The RPC list reaches python through its environment, never its arguments.
RW_RPC=$RPC RW_RPC_FILE=$RPC_FILE RW_PUBLIC_RPCS=$PUBLIC_RPCS \
python3 - "$NODE_KEY_FILE" "$RAILWAY_PROJECT_ID" "$RAILWAY_ENVIRONMENT_ID" "$NODE_NAME" "$IMAGE" "$NETWORK" "$POLL_INTERVAL" "$MOUNT_PATH" <<'PY' > "$vars"
import os, sys, json
keyfile, project, env, name, image, network, poll, mount = sys.argv[1:]
k = json.load(open(keyfile))
k = k.get("data", k) if isinstance(k, dict) else k  # cast wallet new --json wraps the list
k = k[0] if isinstance(k, list) else k
rpc = os.environ.get("RW_RPC", "")
if os.environ.get("RW_RPC_FILE"):
    private = open(os.environ["RW_RPC_FILE"]).readline().strip()
    public = [u.strip() for u in os.environ.get("RW_PUBLIC_RPCS", "").split(",") if u.strip()]
    rpc = ",".join([private] + public)
variables = {
    "DAVINCI_DKG_NETWORK": network,
    "DAVINCI_DKG_PRIVKEY": k["private_key"],
    "DAVINCI_DKG_WEB3_RPC": rpc,
    "DAVINCI_DKG_POLL_INTERVAL": poll,
    "DAVINCI_DKG_DATADIR": mount + "/data",
    "DAVINCI_DKG_ARTIFACTS_DIR": mount + "/artifacts",
}
print(json.dumps({"input": {
    "projectId": project, "environmentId": env, "name": name,
    "source": {"image": image},
    # Unset network and endpoints leave the node on its default network.
    "variables": {key: value for key, value in variables.items() if value},
}}))
PY
resp=$(gql 'mutation($input: ServiceCreateInput!) { serviceCreate(input: $input) { id name } }' "$vars")
service=$(field "$resp" data.serviceCreate.id)
echo "service $NODE_NAME: $service"

# 2. One volume for the pinned artifacts (1.1 GB, downloaded once) and the
#    node's caches; attaching it redeploys the service.
printf '{"input":{"projectId":"%s","environmentId":"%s","serviceId":"%s","mountPath":"%s"}}\n' \
	"$RAILWAY_PROJECT_ID" "$RAILWAY_ENVIRONMENT_ID" "$service" "$MOUNT_PATH" > "$vars"
resp=$(gql 'mutation($input: VolumeCreateInput!) { volumeCreate(input: $input) { id } }' "$vars")
vol=$(field "$resp" data.volumeCreate.id)
echo "volume: $vol at $MOUNT_PATH"

# 3. Restart whatever happens (an RPC outage must not leave the node down).
printf '{"serviceId":"%s","environmentId":"%s","input":{"restartPolicyType":"ALWAYS"}}\n' "$service" "$RAILWAY_ENVIRONMENT_ID" > "$vars"
resp=$(gql 'mutation($serviceId: String!, $environmentId: String!, $input: ServiceInstanceUpdateInput!) { serviceInstanceUpdate(serviceId: $serviceId, environmentId: $environmentId, input: $input) }' "$vars")
policy=$(field "$resp" data.serviceInstanceUpdate)
echo "restart policy: $policy"

# 4. Deploy with everything above in place.
printf '{"serviceId":"%s","environmentId":"%s"}\n' "$service" "$RAILWAY_ENVIRONMENT_ID" > "$vars"
resp=$(gql 'mutation($serviceId: String!, $environmentId: String!) { serviceInstanceDeployV2(serviceId: $serviceId, environmentId: $environmentId) }' "$vars")
dep=$(field "$resp" data.serviceInstanceDeployV2)
echo "deploy: $dep"
echo "operator: $(python3 -c 'import sys,json; k=json.load(open(sys.argv[1])); k=k.get("data",k) if isinstance(k,dict) else k; print((k[0] if isinstance(k,list) else k)["address"])' "$NODE_KEY_FILE")"
