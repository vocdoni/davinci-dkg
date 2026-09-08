#!/usr/bin/env bash
# Show a Railway node service's latest deployment status and log tail.
#   RAILWAY_TOKEN_FILE=railway-api-key RAILWAY_PROJECT_ID=… RAILWAY_ENVIRONMENT_ID=… \
#   scripts/railway-node-status.sh <serviceId> [lines]
set -euo pipefail
: "${RAILWAY_TOKEN_FILE:=railway-api-key}"
: "${RAILWAY_PROJECT_ID:?}"
: "${RAILWAY_ENVIRONMENT_ID:?}"
service=${1:?service id}
lines=${2:-40}
tok=$(tr -d '[:space:]' < "$RAILWAY_TOKEN_FILE")
gql() {
	python3 -c 'import sys,json; print(json.dumps({"query": sys.argv[1], "variables": json.loads(sys.argv[2])}))' "$1" "$2" |
		curl -sS --max-time 60 https://backboard.railway.com/graphql/v2 -H "Authorization: Bearer $tok" -H 'Content-Type: application/json' --data @-
}
dep=$(gql 'query($in: DeploymentListInput!) { deployments(input: $in, first: 1) { edges { node { id status createdAt } } } }' \
	"{\"in\":{\"projectId\":\"$RAILWAY_PROJECT_ID\",\"environmentId\":\"$RAILWAY_ENVIRONMENT_ID\",\"serviceId\":\"$service\"}}")
id=$(python3 -c 'import sys,json; e=json.loads(sys.argv[1])["data"]["deployments"]["edges"]; n=e[0]["node"] if e else {}; print(n.get("id",""))' "$dep")
python3 -c 'import sys,json; e=json.loads(sys.argv[1])["data"]["deployments"]["edges"]; n=e[0]["node"] if e else {}; print("deployment", n.get("id","-"), n.get("status","-"), n.get("createdAt","-"))' "$dep"
[ -n "$id" ] || exit 0
gql 'query($id: String!, $limit: Int!) { deploymentLogs(deploymentId: $id, limit: $limit) { timestamp message } }' "{\"id\":\"$id\",\"limit\":$lines}" |
	python3 -c 'import sys,json,re; d=json.load(sys.stdin); [print(l["timestamp"][11:19], re.sub(r"0x[0-9a-fA-F]{6,}", "0x…", re.sub(r"\x1b\[[0-9;]*m", "", l["message"]))[:150]) for l in d.get("data",{}).get("deploymentLogs",[])]; print(json.dumps(d["errors"])[:300]) if d.get("errors") else None'
