#!/usr/bin/env bash
# Move an existing Railway node service to another network: image, DAVINCI_DKG_NETWORK and the
# RPC list, keeping its volume (the node keeps each deployment's state in its own
# <datadir>/<chainid>-<manager>, and the circuit artifacts are reused) and its operator key.
#
#   RAILWAY_TOKEN_FILE=/home/p4u/davinci-dkg/railway-api-key \
#   RAILWAY_PROJECT=davinci-dkg-gnosis RAILWAY_ENVIRONMENT=production \
#   IMAGE=ghcr.io/vocdoni/davinci-dkg:v0.11.0 NETWORK=optimism \
#   RPC_FILE=~/optimism-rpcs.txt PUBLIC_RPCS=https://mainnet.optimism.io,https://optimism-rpc.publicnode.com \
#   scripts/railway-repoint-node.sh dkg-node1 [--plan|--apply|--rollback]
#
# --plan (default) only reads: it resolves the service, prints its current image, volume and
# variable names, and what --apply would change. --apply saves the values it is about to replace
# to a 0600 file under ~/.cache/railway-repoint/, sets the image, upserts DAVINCI_DKG_NETWORK and
# DAVINCI_DKG_WEB3_RPC (the first line of RPC_FILE first, then PUBLIC_RPCS), deletes
# DAVINCI_DKG_MANAGER and DAVINCI_DKG_DECRYPT_LOOKBACK_BLOCKS when present (either would override
# the network's preset), and deploys; it refuses to run while a saved state exists. --rollback
# restores that saved state, deploys, and sets the file aside (*.rolled-back-<time>).
#
# RAILWAY_PROJECT and RAILWAY_ENVIRONMENT are names or full ids; EXPECT_PROJECT_PREFIX,
# EXPECT_ENVIRONMENT_PREFIX and EXPECT_SERVICE_PREFIX (id prefixes from a deployment record) make
# it refuse anything else. The token, the operator key, the private RPC and every variable value
# stay out of arguments and output: request bodies go through 0600 files.
set -euo pipefail

service_name=${1:?service name, e.g. dkg-node1}
mode=${2:---plan}
case $mode in --plan | --apply | --rollback) ;; *) echo "mode must be --plan, --apply or --rollback" >&2; exit 2 ;; esac
: "${RAILWAY_TOKEN_FILE:?file holding the Railway API token}"
: "${RAILWAY_PROJECT:?project name or id}"
: "${RAILWAY_ENVIRONMENT:=production}"
: "${IMAGE:=}"
: "${NETWORK:=optimism}"
: "${RPC_FILE:=}"
: "${PUBLIC_RPCS:=https://mainnet.optimism.io,https://optimism-rpc.publicnode.com}"
: "${STATE_DIR:=$HOME/.cache/railway-repoint}"
API=${RAILWAY_API:-https://backboard.railway.com/graphql/v2}
OVERRIDES=(DAVINCI_DKG_MANAGER DAVINCI_DKG_DECRYPT_LOOKBACK_BLOCKS)

tmp=$(mktemp -d "${TMPDIR:-$HOME/.cache}/repoint.XXXXXX")
chmod 700 "$tmp"
trap 'rm -rf "$tmp"' EXIT
umask 077
[[ -r $RAILWAY_TOKEN_FILE ]] || { echo "token file $RAILWAY_TOKEN_FILE is not readable" >&2; exit 1; }
printf 'header = "Authorization: Bearer %s"\n' "$(tr -d '[:space:]' <"$RAILWAY_TOKEN_FILE")" >"$tmp/auth"

# gql QUERY VARS_FILE: post one document (the token through a curl config file, never an
# argument); print the data object or fail with the GraphQL error.
gql() {
	python3 -c 'import json, sys; print(json.dumps({"query": sys.argv[1], "variables": json.load(open(sys.argv[2]))}))' \
		"$1" "$2" >"$tmp/body"
	curl -sS --max-time 90 "$API" -K "$tmp/auth" -H 'Content-Type: application/json' --data @"$tmp/body" |
		python3 -c 'import json, sys; d = json.load(sys.stdin)
sys.exit("railway error: " + json.dumps(d["errors"])) if d.get("errors") else print(json.dumps(d["data"]))'
}
# vars FILE k=v...: write a JSON object of string values.
vars() {
	local f=$1
	shift
	python3 -c 'import json, sys; print(json.dumps(dict(a.split("=", 1) for a in sys.argv[2:])))' _ "$@" >"$f"
}
# pick WHAT JSON [NAME [ENV]]: one field out of a response (no eval; names are arguments).
pick() {
	python3 - "$@" <<'PY'
import json, sys
what, d = sys.argv[1], json.loads(sys.argv[2])
arg = sys.argv[3:]
edges = lambda c: [e["node"] for e in c["edges"]]
if what == "project-id":
    out = [p["id"] for p in edges(d["projects"]) if p["name"] == arg[0]]
elif what in ("environment-id", "service-id"):
    out = [n["id"] for n in edges(d["project"]["environments" if what == "environment-id" else "services"]) if n["name"] == arg[0]]
elif what == "mount":
    out = [i["mountPath"] for v in edges(d["project"]["volumes"]) for i in edges(v["volumeInstances"])
           if i["serviceId"] == arg[0] and i["environmentId"] == arg[1]]
elif what == "project-name":
    out = [d["project"]["name"]]
elif what == "image":
    out = [(d["serviceInstance"]["source"] or {}).get("image") or ""]
print(" ".join(out))
PY
}
isuuid() { [[ $1 =~ ^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$ ]]; }

# 1. Project, environment and service ids, from names or ids; checked against the prefixes.
if isuuid "$RAILWAY_PROJECT"; then
	project=$RAILWAY_PROJECT
else
	echo '{}' >"$tmp/v"
	list=$(gql 'query { projects { edges { node { id name } } } }' "$tmp/v")
	project=$(pick project-id "$list" "$RAILWAY_PROJECT")
	[[ -n $project && $project != *" "* ]] || { echo "project $RAILWAY_PROJECT not found once among the token's projects" >&2; exit 1; }
fi
vars "$tmp/v" project="$project"
info=$(gql 'query($project: String!) { project(id: $project) { name
	environments { edges { node { id name } } }
	services { edges { node { id name } } }
	volumes { edges { node { name volumeInstances { edges { node { serviceId environmentId mountPath } } } } } } } }' "$tmp/v")
if isuuid "$RAILWAY_ENVIRONMENT"; then
	env=$RAILWAY_ENVIRONMENT
else
	env=$(pick environment-id "$info" "$RAILWAY_ENVIRONMENT")
fi
service=$(pick service-id "$info" "$service_name")
[[ -n $env && $env != *" "* ]] || { echo "environment $RAILWAY_ENVIRONMENT not found once" >&2; exit 1; }
[[ -n $service && $service != *" "* ]] || { echo "service $service_name not found once" >&2; exit 1; }
for pair in "project:$project:${EXPECT_PROJECT_PREFIX:-}" "environment:$env:${EXPECT_ENVIRONMENT_PREFIX:-}" \
	"service:$service:${EXPECT_SERVICE_PREFIX:-}"; do
	IFS=: read -r what id prefix <<<"$pair"
	[[ -z $prefix || $id == "$prefix"* ]] || { echo "$what id $id does not start with $prefix" >&2; exit 1; }
done
mount=$(pick mount "$info" "$service" "$env")
[[ -n $mount ]] || { echo "service $service_name has no volume in this environment: refusing" >&2; exit 1; }

vars "$tmp/v" project="$project" env="$env" service="$service"
inst=$(gql 'query($service: String!, $env: String!) { serviceInstance(serviceId: $service, environmentId: $env) { source { image } } }' "$tmp/v")
image_now=$(pick image "$inst")
gql 'query($project: String!, $env: String!, $service: String!) { variables(projectId: $project, environmentId: $env, serviceId: $service) }' \
	"$tmp/v" >"$tmp/current" # values stay in this 0600 file
names=$(python3 -c 'import json, sys; print(" ".join(sorted(json.load(open(sys.argv[1]))["variables"])))' "$tmp/current")
value_of() { python3 -c 'import json, sys; print(json.load(open(sys.argv[1]))["variables"].get(sys.argv[2], ""))' "$tmp/current" "$1"; }

echo "project $(pick project-name "$info") ($project), environment $env"
echo "service $service_name ($service): image ${image_now:-<none>}, volume at $mount"
echo "variables: $names"
echo "DAVINCI_DKG_NETWORK now: '$(value_of DAVINCI_DKG_NETWORK)'"
saved=$STATE_DIR/$service_name-$service.json

if [[ $mode == --rollback ]]; then
	[[ -r $saved ]] || { echo "no saved state at $saved" >&2; exit 1; }
	python3 - "$saved" "$project" "$env" "$service" >"$tmp/v" <<'PY'
import json, sys
s = json.load(open(sys.argv[1]))
print(json.dumps({"input": {"projectId": sys.argv[2], "environmentId": sys.argv[3], "serviceId": sys.argv[4],
                            "variables": {k: v for k, v in s["variables"].items() if v != ""}, "skipDeploys": True}}))
PY
	[[ -n $(python3 -c 'import json, sys; print(json.load(open(sys.argv[1]))["image"])' "$saved") ]] || { echo "saved state has no image" >&2; exit 1; }
	python3 -c 'import json, sys; s = json.load(open(sys.argv[1])); print(json.dumps({"service": sys.argv[2], "env": sys.argv[3], "input": {"source": {"image": s["image"]}}}))' \
		"$saved" "$service" "$env" >"$tmp/img"
	gql 'mutation($service: String!, $env: String!, $input: ServiceInstanceUpdateInput!) { serviceInstanceUpdate(serviceId: $service, environmentId: $env, input: $input) }' "$tmp/img" >/dev/null
	gql 'mutation($input: VariableCollectionUpsertInput!) { variableCollectionUpsert(input: $input) }' "$tmp/v" >/dev/null
	for name in "${OVERRIDES[@]}" DAVINCI_DKG_NETWORK DAVINCI_DKG_WEB3_RPC; do
		if [[ -z $(python3 -c 'import json, sys; print(json.load(open(sys.argv[1]))["variables"].get(sys.argv[2], ""))' "$saved" "$name") && " $names " == *" $name "* ]]; then
			vars "$tmp/d" projectId="$project" environmentId="$env" serviceId="$service" name="$name"
			python3 -c 'import json, sys; print(json.dumps({"input": json.load(open(sys.argv[1]))}))' "$tmp/d" >"$tmp/v"
			gql 'mutation($input: VariableDeleteInput!) { variableDelete(input: $input) }' "$tmp/v" >/dev/null
		fi
	done
	vars "$tmp/v" service="$service" env="$env"
	echo "rolled back: $(gql 'mutation($service: String!, $env: String!) { serviceInstanceDeployV2(serviceId: $service, environmentId: $env) }' "$tmp/v")"
	# Set the restored state aside (kept, 0600) so a later --apply can save afresh.
	mv "$saved" "$saved.rolled-back-$(date -u +%Y%m%dT%H%M%SZ)"
	exit 0
fi

: "${IMAGE:?IMAGE, e.g. ghcr.io/vocdoni/davinci-dkg:v0.11.0 (a pinned release, not latest)}"
[[ -n $RPC_FILE && -r $RPC_FILE ]] || { echo "RPC_FILE (first line: the private RPC) is not readable" >&2; exit 1; }
private=$(head -n1 "$RPC_FILE" | tr -d '[:space:]')
[[ $private =~ ^https?:// ]] || { echo "the first line of RPC_FILE is not an http(s) URL" >&2; exit 1; }
deletes=()
for name in "${OVERRIDES[@]}"; do [[ " $names " == *" $name "* ]] && deletes+=("$name"); done
echo "plan: image ${image_now:-<none>} -> $IMAGE"
echo "plan: DAVINCI_DKG_NETWORK -> $NETWORK"
echo "plan: DAVINCI_DKG_WEB3_RPC -> <private RPC from $RPC_FILE>,$PUBLIC_RPCS"
echo "plan: delete ${deletes[*]:-nothing}"
echo "plan: keep the volume at $mount, the operator key and every other variable; deploy"
[[ -e $saved ]] && echo "note: saved state $saved exists; --apply will refuse until --rollback sets it aside"
[[ $mode == --apply ]] || { echo "(plan only: rerun with --apply)"; exit 0; }
# A second --apply (a retry after a partial failure) would save the already-moved state over the
# original one, and --rollback would then leave the service moved.
[[ -e $saved ]] && { echo "saved state exists at $saved: --rollback, or move it aside first" >&2; exit 1; }

# 2. Save what is replaced, for --rollback (0600, values never printed).
mkdir -p "$STATE_DIR"
chmod 700 "$STATE_DIR"
python3 - "$tmp/current" "$image_now" DAVINCI_DKG_NETWORK DAVINCI_DKG_WEB3_RPC "${OVERRIDES[@]}" >"$saved" <<'PY'
import json, sys
cur = json.load(open(sys.argv[1]))["variables"]
print(json.dumps({"image": sys.argv[2], "variables": {k: cur.get(k, "") for k in sys.argv[3:]}}))
PY
chmod 600 "$saved"
echo "saved the replaced values to $saved"

# 3. Image first (no deploy), variables with skipDeploys, deletes, then one deploy.
python3 -c 'import json, sys; print(json.dumps({"service": sys.argv[1], "env": sys.argv[2], "input": {"source": {"image": sys.argv[3]}}}))' \
	"$service" "$env" "$IMAGE" >"$tmp/v"
gql 'mutation($service: String!, $env: String!, $input: ServiceInstanceUpdateInput!) { serviceInstanceUpdate(serviceId: $service, environmentId: $env, input: $input) }' "$tmp/v" >/dev/null
echo "image set"
RW_PRIVATE=$private python3 - "$project" "$env" "$service" "$NETWORK" "$PUBLIC_RPCS" >"$tmp/v" <<'PY'
import json, os, sys
project, env, service, network, public = sys.argv[1:]
rpcs = [os.environ["RW_PRIVATE"]] + [u.strip() for u in public.split(",") if u.strip()]
print(json.dumps({"input": {"projectId": project, "environmentId": env, "serviceId": service, "skipDeploys": True,
                            "variables": {"DAVINCI_DKG_NETWORK": network, "DAVINCI_DKG_WEB3_RPC": ",".join(rpcs)}}}))
PY
unset private
gql 'mutation($input: VariableCollectionUpsertInput!) { variableCollectionUpsert(input: $input) }' "$tmp/v" >/dev/null
echo "variables set: DAVINCI_DKG_NETWORK, DAVINCI_DKG_WEB3_RPC"
for name in "${deletes[@]}"; do
	vars "$tmp/d" projectId="$project" environmentId="$env" serviceId="$service" name="$name"
	python3 -c 'import json, sys; print(json.dumps({"input": json.load(open(sys.argv[1]))}))' "$tmp/d" >"$tmp/v"
	gql 'mutation($input: VariableDeleteInput!) { variableDelete(input: $input) }' "$tmp/v" >/dev/null
	echo "deleted $name"
done
vars "$tmp/v" service="$service" env="$env"
echo "deployment: $(gql 'mutation($service: String!, $env: String!) { serviceInstanceDeployV2(serviceId: $service, environmentId: $env) }' "$tmp/v")"
echo "follow it with scripts/railway-node-status.sh $service"
