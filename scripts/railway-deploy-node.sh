#!/usr/bin/env bash
# Deploy one davinci-dkg node as a Railway service (public GHCR image, one
# volume, always-restart) through Railway's GraphQL API. See railway-deploy.md.
#
#   RAILWAY_TOKEN_FILE=railway-api-key \
#   RAILWAY_PROJECT_ID=<uuid> RAILWAY_ENVIRONMENT_ID=<uuid> \
#   NODE_NAME=dkg-node9 NODE_KEY_FILE=~/.davinci-dkg-sepolia/node9.json \
#   scripts/railway-deploy-node.sh
#
# NODE_KEY_FILE holds `[{"address": "0x…", "private_key": "0x…"}]` (the fleet
# format). The key travels only inside the request body read from a 0600
# temp file; it is never an argument and never printed.
set -euo pipefail

: "${RAILWAY_TOKEN_FILE:=railway-api-key}"
: "${RAILWAY_PROJECT_ID:?project id (projectCreate output)}"
: "${RAILWAY_ENVIRONMENT_ID:?environment id (usually the production environment of the project)}"
: "${NODE_NAME:?service name, e.g. dkg-node9}"
: "${NODE_KEY_FILE:?operator key file}"
: "${IMAGE:=ghcr.io/vocdoni/davinci-dkg:latest}"
: "${NETWORK:=sepolia}"
: "${RPC:=https://ethereum-sepolia-rpc.publicnode.com,https://1rpc.io/sepolia,https://sepolia.gateway.tenderly.co}"
: "${POLL_INTERVAL:=30s}"
: "${MOUNT_PATH:=/app/run}"
API=https://backboard.railway.com/graphql/v2

tok=$(tr -d '[:space:]' < "$RAILWAY_TOKEN_FILE")
vars=$(mktemp); chmod 600 "$vars"; trap 'rm -f "$vars"' EXIT

# gql QUERY [VARS_FILE]: post one GraphQL document, print the JSON response.
gql() {
	local q=$1 vf=${2:-}
	python3 - "$q" "$vf" <<'PY' > "$vars.body"
import sys, json
q, vf = sys.argv[1], sys.argv[2]
v = json.load(open(vf)) if vf else {}
print(json.dumps({"query": q, "variables": v}))
PY
	curl -sS --max-time 90 "$API" -H "Authorization: Bearer $tok" -H 'Content-Type: application/json' --data @"$vars.body"
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
python3 - "$NODE_KEY_FILE" "$RAILWAY_PROJECT_ID" "$RAILWAY_ENVIRONMENT_ID" "$NODE_NAME" "$IMAGE" "$NETWORK" "$RPC" "$POLL_INTERVAL" "$MOUNT_PATH" <<'PY' > "$vars"
import sys, json
keyfile, project, env, name, image, network, rpc, poll, mount = sys.argv[1:]
k = json.load(open(keyfile)); k = k[0] if isinstance(k, list) else k
print(json.dumps({"input": {
    "projectId": project, "environmentId": env, "name": name,
    "source": {"image": image},
    "variables": {
        "DAVINCI_DKG_NETWORK": network,
        "DAVINCI_DKG_PRIVKEY": k["private_key"],
        "DAVINCI_DKG_WEB3_RPC": rpc,
        "DAVINCI_DKG_POLL_INTERVAL": poll,
        "DAVINCI_DKG_DATADIR": mount + "/data",
        "DAVINCI_DKG_ARTIFACTS_DIR": mount + "/artifacts",
    }}}))
PY
resp=$(gql 'mutation($input: ServiceCreateInput!) { serviceCreate(input: $input) { id name } }' "$vars")
service=$(field "$resp" data.serviceCreate.id)
echo "service $NODE_NAME: $service"

# 2. One volume for the pinned artifacts (1.1 GB, downloaded once) and the
#    node's caches; attaching it redeploys the service.
printf '{"input":{"projectId":"%s","environmentId":"%s","serviceId":"%s","mountPath":"%s"}}\n' \
	"$RAILWAY_PROJECT_ID" "$RAILWAY_ENVIRONMENT_ID" "$service" "$MOUNT_PATH" > "$vars"
resp=$(gql 'mutation($input: VolumeCreateInput!) { volumeCreate(input: $input) { id } }' "$vars")
echo "volume: $(field "$resp" data.volumeCreate.id) at $MOUNT_PATH"

# 3. Restart whatever happens (an RPC outage must not leave the node down).
printf '{"serviceId":"%s","environmentId":"%s","input":{"restartPolicyType":"ALWAYS"}}\n' "$service" "$RAILWAY_ENVIRONMENT_ID" > "$vars"
resp=$(gql 'mutation($serviceId: String!, $environmentId: String!, $input: ServiceInstanceUpdateInput!) { serviceInstanceUpdate(serviceId: $serviceId, environmentId: $environmentId, input: $input) }' "$vars")
echo "restart policy: $(field "$resp" data.serviceInstanceUpdate)"

# 4. Deploy with everything above in place.
printf '{"serviceId":"%s","environmentId":"%s"}\n' "$service" "$RAILWAY_ENVIRONMENT_ID" > "$vars"
resp=$(gql 'mutation($serviceId: String!, $environmentId: String!) { serviceInstanceDeployV2(serviceId: $serviceId, environmentId: $environmentId) }' "$vars")
echo "deploy: $(field "$resp" data.serviceInstanceDeployV2)"
echo "operator: $(python3 -c 'import sys,json; k=json.load(open(sys.argv[1])); print((k[0] if isinstance(k,list) else k)["address"])' "$NODE_KEY_FILE")"
