#!/usr/bin/env bash
# benchmark.sh — Run the DKG testnet at multiple (n, threshold) configurations
# and collect gas costs, proof times, and circuit statistics.
#
# Usage (from repo root):
#   bash scripts/benchmark.sh [output_file]
#
# Output: JSON lines written to output_file (default: /tmp/dkg-benchmark.jsonl)
# Requires: jq, docker compose (v2), make
#
# The script tears down any running testnet before each run, starts a fresh
# network, executes the scenario, and parses Anvil logs for gas and timing.
#
# Benchmark matrix:
#   n (nodes) = 4, 8, 12, 16, 20, 24, 28, 32
#   threshold  = ceil(n * 2/3)   (2/3 majority)
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
REPO_ROOT="$(cd "${SCRIPT_DIR}/.." && pwd)"
OUTPUT="${1:-/tmp/dkg-benchmark.jsonl}"
TESTNET_DIR="${REPO_ROOT}/testnet"

# Participant counts to benchmark
SIZES=(4 8 12 16 20 24 28 32)

cd "${REPO_ROOT}"

# Print circuit constraint counts from the compiled artifact JSON (if present).
circuit_constraints() {
    local json_file="${HOME}/.davinci/artifacts/circuit-artifacts.json"
    if [ ! -f "${json_file}" ]; then
        echo "n/a"
        return
    fi
    jq -r '.contribution.circuit_hash // "n/a"' "${json_file}" 2>/dev/null || echo "n/a"
}

# Parse gas from docker logs produced by the runner container.
# Logs lines like: gas=427312
parse_gas() {
    local log="$1"
    local label="$2"
    grep -o "${label}=[0-9]*" "${log}" 2>/dev/null | tail -1 | cut -d= -f2 || echo "0"
}

# Run one benchmark iteration.
run_once() {
    local n="$1"
    local t="$2"
    local log_file="/tmp/dkg-bench-n${n}.log"

    echo "=== Benchmark n=${n} t=${t} ===" | tee -a "${OUTPUT}.log"

    # Tear down any previous run.
    (cd "${TESTNET_DIR}" && docker compose down -v --remove-orphans 2>/dev/null || true)
    sleep 2

    # Start the network.
    (cd "${TESTNET_DIR}" && \
        DKG_NODE_COUNT="${n}" \
        docker compose up -d --scale "dkg-node=${n}" --build \
        anvil deployer dkg-node 2>&1) | tee -a "${OUTPUT}.log"

    echo "Waiting for nodes to register (${n} nodes)..."
    local wait_secs=$((n * 10 + 30))
    sleep "${wait_secs}"

    # Run the scenario and capture timing.
    local start_ts
    start_ts=$(date +%s%N)

    set +e
    (cd "${TESTNET_DIR}" && \
        DKG_RUNNER_NODES="${n}" \
        DKG_RUNNER_THRESHOLD="${t}" \
        docker compose run --rm \
        -e DKG_RUNNER_WAIT_READINESS=10m \
        -e DKG_RUNNER_WAIT_CONTRIB=15m \
        -e DKG_RUNNER_WAIT_DECRYPT=15m \
        dkg-runner 2>&1) | tee "${log_file}"
    local exit_code=$?
    set -e

    local end_ts
    end_ts=$(date +%s%N)
    local elapsed_ms=$(( (end_ts - start_ts) / 1000000 ))

    local success="false"
    if grep -q "DKG scenario completed successfully" "${log_file}"; then
        success="true"
    fi

    # Extract gas values from runner logs.
    local gas_contribution gas_finalize gas_partial_decrypt gas_combine
    gas_contribution=$(grep -o 'submitContribution.*gas=[0-9]*' "${log_file}" 2>/dev/null | grep -o 'gas=[0-9]*' | head -1 | cut -d= -f2 || echo "0")
    gas_finalize=$(grep -o 'finalizeRound.*gas=[0-9]*\|finalize.*gas=[0-9]*' "${log_file}" 2>/dev/null | grep -o 'gas=[0-9]*' | head -1 | cut -d= -f2 || echo "0")
    gas_partial_decrypt=$(grep -o 'submitPartialDecryption.*gas=[0-9]*\|partialDecrypt.*gas=[0-9]*' "${log_file}" 2>/dev/null | grep -o 'gas=[0-9]*' | head -1 | cut -d= -f2 || echo "0")
    gas_combine=$(grep -o 'combineDecryption.*gas=[0-9]*\|combine.*gas=[0-9]*' "${log_file}" 2>/dev/null | grep -o 'gas=[0-9]*' | head -1 | cut -d= -f2 || echo "0")

    # Write JSON result.
    jq -n \
        --argjson n "${n}" \
        --argjson t "${t}" \
        --argjson elapsed_ms "${elapsed_ms}" \
        --argjson success "${success}" \
        --arg gas_contribution "${gas_contribution}" \
        --arg gas_finalize "${gas_finalize}" \
        --arg gas_partial_decrypt "${gas_partial_decrypt}" \
        --arg gas_combine "${gas_combine}" \
        --arg timestamp "$(date -u +%Y-%m-%dT%H:%M:%SZ)" \
        '{
            timestamp: $timestamp,
            n: $n,
            threshold: $t,
            success: $success,
            elapsed_ms: $elapsed_ms,
            gas: {
                contribution: ($gas_contribution | tonumber),
                finalize: ($gas_finalize | tonumber),
                partial_decrypt: ($gas_partial_decrypt | tonumber),
                combine: ($gas_combine | tonumber)
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
(cd "${TESTNET_DIR}" && docker compose down -v --remove-orphans 2>/dev/null || true)

echo ""
echo "Benchmark complete. Results in ${OUTPUT}"
echo "Summary:"
jq -r '. | "\(.n) nodes (t=\(.threshold)): success=\(.success) elapsed=\(.elapsed_ms)ms gas_contribution=\(.gas.contribution) gas_finalize=\(.gas.finalize) gas_combine=\(.gas.combine)"' "${OUTPUT}"
