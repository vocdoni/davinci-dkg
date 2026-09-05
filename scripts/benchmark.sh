#!/usr/bin/env bash
# benchmark.sh — Run the DKG testnet at multiple (n, threshold) configurations
# and collect the gas of every protocol transaction plus scenario timing.
#
# Usage (from repo root):
#   bash scripts/benchmark.sh [output_file]
#
# Output: JSON lines written to output_file (default: /tmp/dkg-benchmark.jsonl)
# Requires: jq, curl, docker compose (v2), go (the battery drives the scenario)
# Optional env: SIZES="4 8" (node counts), ANVIL_PORT, DEPLOYER_PORT,
#               BATTERY_ORGANIZERS / BATTERY_CIPHERTEXTS (swarm size per run)
#
# For every size the script tears down any running testnet, starts a fresh
# fleet (Anvil + deployer + n nodes) with the battery's compose override so
# the nodes activate the whole MAX_K pool of each epoch, lets the fleet
# create, contribute to, finalize and activate an epoch on its own, then runs
# the battery's organizer swarm against it (registrations, ciphertexts, an
# organizer reveal, and the nodes' partials and combines).
#
# Gas comes from two places: the node logs for the committee-side
# transactions (submitContribution, finalizeEpoch, activatePoolKey) and the
# battery's JSON report for the application-side ones (registerApplication,
# submitCiphertext, revealOrganizerSecret, submitPartialDecryption,
# combineDecryption).
#
# Benchmark matrix:
#   n (nodes) = 4, 8, 12, 16, 20, 24, 28, 32
#   threshold  = ceil(n * 2/3)   (2/3 majority)
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
OUTPUT="${1:-/tmp/dkg-benchmark.jsonl}"
TESTNET_DIR="${REPO_ROOT}/testnet"
# The battery override sets DAVINCI_DKG_ACTIVATE_AHEAD=8 on the nodes (see
# tests/battery/compose.battery.yml for why).
COMPOSE=(docker compose -f docker-compose.yml -f ../tests/battery/compose.battery.yml)

# Participant counts to benchmark.
read -r -a SIZES <<< "${SIZES:-4 8 12 16 20 24 28 32}"
ANVIL_PORT="${ANVIL_PORT:-8545}"
DEPLOYER_PORT="${DEPLOYER_PORT:-8888}"
# Swarm size per run: small on purpose, the gas figures do not depend on it.
BATTERY_ORGANIZERS="${BATTERY_ORGANIZERS:-2}"
BATTERY_CIPHERTEXTS="${BATTERY_CIPHERTEXTS:-2}"

cd "${REPO_ROOT}"

compose() {
    (cd "${TESTNET_DIR}" && "${COMPOSE[@]}" "$@")
}

# Highest "gas=N" on the node-log lines matching the label (0 when absent).
# The nodes write zerolog console lines, ANSI-coloured even without a TTY,
# such as:  2026-09-04T10:00:00Z INF finalizeEpoch tx mined tx=0x… gas=35123
parse_node_gas() {
    local log="$1" label="$2" gas
    gas=$(sed 's/\x1b\[[0-9;]*m//g' "${log}" 2>/dev/null \
        | grep -F "${label}" | grep -o 'gas=[0-9]*' | cut -d= -f2 | sort -n | tail -1 || true)
    echo "${gas:-0}"
}

# Highest gas among the battery report rows of one transaction kind (0 when absent).
parse_report_gas() {
    local report="$1" kind="$2"
    if [ ! -f "${report}" ]; then
        echo 0
        return
    fi
    jq -r --arg kind "${kind}" '[.[] | select(.kind == $kind and .gas > 0) | .gas] | max // 0' "${report}"
}

# Average of a latency field over the combine rows of the battery report.
parse_report_latency() {
    local report="$1" field="$2"
    if [ ! -f "${report}" ]; then
        echo 0
        return
    fi
    jq -r --arg field "${field}" \
        '[.[] | select(.kind == "combineDecryption(node)" and .pass) | .[$field]] | if length == 0 then 0 else add / length end' \
        "${report}"
}

# Run one benchmark iteration.
run_once() {
    local n="$1"
    local t="$2"
    local log_file="/tmp/dkg-bench-n${n}.log"
    local node_log="/tmp/dkg-bench-n${n}-nodes.log"
    local report="/tmp/dkg-bench-n${n}-report.json"
    local addresses="/tmp/dkg-bench-n${n}-addresses.env"

    echo "=== Benchmark n=${n} t=${t} ===" | tee -a "${OUTPUT}.log"

    # Tear down any previous run.
    compose down -v --remove-orphans 2>/dev/null || true
    sleep 2

    # Start the network: every node in the committee, threshold t.
    rm -f "${addresses}" "${report}"
    (cd "${TESTNET_DIR}" && \
        DKG_NODE_COUNT="${n}" \
        DKG_THRESHOLD="${t}" \
        DKG_COMMITTEE_SIZE="${n}" \
        DKG_MIN_VALID_CONTRIBUTIONS="${t}" \
        "${COMPOSE[@]}" up -d --build --scale "dkg-node=${n}" \
        anvil deployer dkg-node 2>&1) | tee -a "${OUTPUT}.log"

    echo "Waiting for the deployer to publish the contract addresses..."
    local i
    for i in $(seq 1 100); do
        if curl -fsS "http://127.0.0.1:${DEPLOYER_PORT}/addresses.env" > "${addresses}" 2>/dev/null \
            && [ -s "${addresses}" ]; then
            break
        fi
        sleep 3
    done

    # Run the scenario and capture timing. The battery waits for the fleet's
    # first Live epoch by itself (BATTERY_CONNECT_TIMEOUT, 10 minutes).
    local start_ts
    start_ts=$(date +%s%N)

    set +e
    DAVINCI_DKG_BATTERY=1 \
    DAVINCI_DKG_TEST_RPC_URL="http://127.0.0.1:${ANVIL_PORT}" \
    DAVINCI_DKG_TEST_ADDRESSES="${addresses}" \
    DAVINCI_DKG_BATTERY_REPORT="${report}" \
    BATTERY_ORGANIZERS="${BATTERY_ORGANIZERS}" \
    BATTERY_CIPHERTEXTS="${BATTERY_CIPHERTEXTS}" \
    go test ./tests/battery -run 'TestOrganizerSwarm' -count=1 -v -timeout 40m 2>&1 | tee "${log_file}"
    local exit_code=${PIPESTATUS[0]}
    set -e

    local end_ts
    end_ts=$(date +%s%N)
    local elapsed_ms=$(( (end_ts - start_ts) / 1000000 ))

    local success="false"
    if [ "${exit_code}" -eq 0 ] && grep -q '^ok' "${log_file}"; then
        success="true"
    fi

    # Committee-side gas from the node logs.
    compose logs --no-color dkg-node > "${node_log}" 2>&1 || true
    local gas_contribution gas_finalize gas_activate
    gas_contribution=$(parse_node_gas "${node_log}" "contribution submitted")
    gas_finalize=$(parse_node_gas "${node_log}" "finalizeEpoch tx mined")
    gas_activate=$(parse_node_gas "${node_log}" "activatePoolKey tx mined")

    # Application-side gas and latency from the battery report.
    local gas_register gas_ciphertext gas_reveal gas_partial_decrypt gas_combine
    gas_register=$(parse_report_gas "${report}" "registerApplication")
    gas_ciphertext=$(parse_report_gas "${report}" "submitCiphertext")
    gas_reveal=$(parse_report_gas "${report}" "revealOrganizerSecret")
    gas_partial_decrypt=$(parse_report_gas "${report}" "submitPartialDecryption(node)")
    gas_combine=$(parse_report_gas "${report}" "combineDecryption(node)")
    local combine_blocks combine_seconds
    combine_blocks=$(parse_report_latency "${report}" "latencyBlocks")
    combine_seconds=$(parse_report_latency "${report}" "latencySeconds")

    # Write JSON result.
    jq -n \
        --argjson n "${n}" \
        --argjson t "${t}" \
        --argjson elapsed_ms "${elapsed_ms}" \
        --argjson success "${success}" \
        --argjson gas_contribution "${gas_contribution}" \
        --argjson gas_finalize "${gas_finalize}" \
        --argjson gas_activate "${gas_activate}" \
        --argjson gas_register "${gas_register}" \
        --argjson gas_ciphertext "${gas_ciphertext}" \
        --argjson gas_reveal "${gas_reveal}" \
        --argjson gas_partial_decrypt "${gas_partial_decrypt}" \
        --argjson gas_combine "${gas_combine}" \
        --argjson combine_blocks "${combine_blocks}" \
        --argjson combine_seconds "${combine_seconds}" \
        --arg timestamp "$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
        '{
            timestamp: $timestamp,
            n: $n,
            threshold: $t,
            success: $success,
            elapsed_ms: $elapsed_ms,
            gas: {
                contribution: $gas_contribution,
                finalize: $gas_finalize,
                activate: $gas_activate,
                register: $gas_register,
                ciphertext: $gas_ciphertext,
                reveal: $gas_reveal,
                partial_decrypt: $gas_partial_decrypt,
                combine: $gas_combine
            },
            latency: {
                combine_blocks: $combine_blocks,
                combine_seconds: $combine_seconds
            }
        }' >> "${OUTPUT}"

    echo "n=${n} t=${t}: success=${success} elapsed=${elapsed_ms}ms"
}

echo "DKG Benchmark Suite" | tee "${OUTPUT}.log"
echo "Started: $(date)" | tee -a "${OUTPUT}.log"
echo "Output: ${OUTPUT}" | tee -a "${OUTPUT}.log"
echo "" | tee -a "${OUTPUT}.log"

# Clear output file.
> "${OUTPUT}"

for n in "${SIZES[@]}"; do
    t=$(( (n * 2 + 2) / 3 ))   # ceil(2n/3)
    run_once "${n}" "${t}"
done

# Tear down final run.
compose down -v --remove-orphans 2>/dev/null || true

echo ""
echo "Benchmark complete. Results in ${OUTPUT}"
echo "Summary:"
jq -r '. | "\(.n) nodes (t=\(.threshold)): success=\(.success) elapsed=\(.elapsed_ms)ms contribution=\(.gas.contribution) finalize=\(.gas.finalize) activate=\(.gas.activate) register=\(.gas.register) reveal=\(.gas.reveal) partial=\(.gas.partial_decrypt) combine=\(.gas.combine)"' "${OUTPUT}"
