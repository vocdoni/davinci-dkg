#!/usr/bin/env bash
# Move an existing Railway node service to another image and deploy it. Railway
# runs no Watchtower: a release (new protocol, or a default preset pointing at a
# new deployment) reaches a Railway node only through a new deployment. Pin the
# release tag rather than `latest`, so the rollout is exactly that build.
#
#   RAILWAY_TOKEN_FILE=railway-api-key RAILWAY_ENVIRONMENT_ID=<uuid> \
#   scripts/railway-roll-node.sh <serviceId> ghcr.io/vocdoni/davinci-dkg:vX.Y.Z
#
# The volume, the variables and the operator key stay as they are. Follow the
# start with scripts/railway-node-status.sh <serviceId>.
set -euo pipefail

: "${RAILWAY_TOKEN_FILE:=railway-api-key}"
: "${RAILWAY_ENVIRONMENT_ID:?environment id of the service}"
service=${1:?service id}
image=${2:?image, e.g. ghcr.io/vocdoni/davinci-dkg:v0.10.0}
API=${RAILWAY_API:-https://backboard.railway.com/graphql/v2}

body=$(mktemp); chmod 600 "$body"
auth=$(mktemp); chmod 600 "$auth"
trap 'rm -f "$body" "$auth"' EXIT
# The token reaches curl through a 0600 config file, never its arguments.
printf 'header = "Authorization: Bearer %s"\n' "$(tr -d '[:space:]' < "$RAILWAY_TOKEN_FILE")" > "$auth"

# gql QUERY VARS_JSON: post one GraphQL document, fail on a GraphQL error,
# print the data object.
gql() {
	python3 -c 'import sys,json; print(json.dumps({"query": sys.argv[1], "variables": json.loads(sys.argv[2])}))' \
		"$1" "$2" > "$body"
	curl -sS --max-time 90 "$API" -K "$auth" -H 'Content-Type: application/json' \
		--data @"$body" |
		python3 -c 'import sys,json; d=json.load(sys.stdin);
sys.exit("railway error: "+json.dumps(d["errors"])) if d.get("errors") else print(json.dumps(d["data"]))'
}

ids=$(printf '{"serviceId":"%s","environmentId":"%s"' "$service" "$RAILWAY_ENVIRONMENT_ID")
gql 'mutation($serviceId: String!, $environmentId: String!, $input: ServiceInstanceUpdateInput!) {
	serviceInstanceUpdate(serviceId: $serviceId, environmentId: $environmentId, input: $input) }' \
	"$ids,\"input\":{\"source\":{\"image\":\"$image\"}}}" >/dev/null
echo "service $service: image $image"
deployment=$(gql 'mutation($serviceId: String!, $environmentId: String!) {
	serviceInstanceDeployV2(serviceId: $serviceId, environmentId: $environmentId) }' "$ids}")
echo "deployment: $deployment"
