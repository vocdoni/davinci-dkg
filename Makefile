.PHONY: help \
        circuits-compile circuits-update-hashes circuits circuits-release \
        solidity-build solidity-bind solidity-deploy \
        testnet-up testnet-run testnet-down testnet-logs \
        ui-install ui-build ui-dev ui-clean ui-test ui-config \
        build test test-integration

# Default node count and threshold for testnet
DKG_NODE_COUNT        ?= 3
DKG_THRESHOLD         ?= 2
DKG_DISCLOSURE_ALLOWED ?= false

# Circuit artifact cache directory (mirrors DAVINCI_DKG_ARTIFACTS_DIR default)
ARTIFACTS_DIR ?= $(HOME)/.davinci/artifacts

# S3 / DigitalOcean Spaces settings for circuits-release
# Override S3_ACCESS_KEY and S3_SECRET_KEY on the command line or via env.
S3_ACCESS_KEY ?=
S3_SECRET_KEY ?=
S3_SPACE      ?= circuits
S3_BUCKET     ?= dev

# Deployment + UI config parameters (override on command line or via env).
# RPC_URL and CHAIN_ID are read by both the solidity-deploy target and by
# `make ui-dev` / `make ui-build` (which template ui/public/config.json so
# the dev server / built bundle targets the chain you specify). When unset,
# the UI keeps the defaults baked into ui/public/config.json (Sepolia).
RPC_URL         ?=
CHAIN_ID        ?=
CHAIN_NAME      ?=
MANAGER_ADDRESS ?=
PRIVATE_KEY     ?=

# Temporary file used to pass compile output between targets
CIRCUIT_ARTIFACTS_JSON ?= /tmp/circuit-artifacts.json

help: ## Show this help message
	@echo "DAVINCI DKG Makefile"
	@echo ""
	@echo "Usage: make [target] [VAR=value ...]"
	@echo ""
	@echo "Circuit Update Pipeline:"
	@echo "  circuits              Full pipeline: compile → update hashes → Solidity build → Go bind"
	@echo "  circuits-compile      Compile all 6 circuits; write artifacts + Solidity verifier files"
	@echo "                        ARTIFACTS_DIR (default: ~/.davinci/artifacts)"
	@echo "  circuits-release      Compile + upload to CDN + update hashes + Solidity rebuild"
	@echo "                        Requires: S3_ACCESS_KEY=<key> S3_SECRET_KEY=<secret>"
	@echo "                        Optional: S3_BUCKET=<channel> (default: dev)"
	@echo "  circuits-update-hashes Patch config/circuit_artifacts.go with hashes from last compile"
	@echo "  solidity-build        Build Solidity contracts (forge build)"
	@echo "  solidity-bind         Regenerate Go ABI bindings (go_bind.sh)"
	@echo "  solidity-deploy       Deploy contracts  (requires RPC_URL, CHAIN_ID, PRIVATE_KEY)"
	@echo "                        Optional: ETHERSCAN_API_KEY, VERIFIER_URL, SKIP_TESTS=1"
	@echo ""
	@echo "Testnet Commands:"
	@echo "  testnet-up      Start the local DKG testnet: Anvil + deployer +"
	@echo "                  N dkg-node replicas + the standalone UI service."
	@echo "                  DKG_NODE_COUNT     (default 3, max 32 containers)"
	@echo "                  DKG_THRESHOLD      (default 2)"
	@echo "                  UI_PORT            (default 8081, host binding)"
	@echo "                  UI_PUBLIC_RPC      RPC URL advertised to browsers"
	@echo "                                     — set to http://<host-ip>:8545"
	@echo "                                     when accessing from another host."
	@echo "                  Note: committee size is capped by the circuit"
	@echo "                  bound MaxN (see circuits/common/sizes.go, currently 32)."
	@echo "  testnet-run     Run the full DKG scenario (create round → encrypt → decrypt)"
	@echo "                  DKG_NODE_COUNT         (default 3)"
	@echo "                  DKG_THRESHOLD          (default 2)"
	@echo "                  DKG_DISCLOSURE_ALLOWED (default false) enables the"
	@echo "                                         reveal-share disclosure phase"
	@echo "  testnet-logs    Tail logs of the dkg-node containers"
	@echo "  testnet-down    Stop the testnet and wipe Docker volumes"
	@echo ""
	@echo "UI Commands:"
	@echo "  ui-install       Install UI dependencies (pnpm install)"
	@echo "  ui-build         Build the DKG explorer UI to ui/dist"
	@echo "  ui-dev           Run the UI in dev mode (Vite, hot reload on :5174)"
	@echo "  ui-test          Run UI unit tests (vitest)"
	@echo "  ui-clean         Remove ui/dist and ui/node_modules"
	@echo "  ui-config        Re-render ui/public/config.json from RPC_URL et al."
	@echo "                   Optional vars (override which chain the UI targets):"
	@echo "                     RPC_URL          (default: Sepolia public RPC)"
	@echo "                     MANAGER_ADDRESS  (default: Sepolia DKGManager)"
	@echo "                     CHAIN_ID         (default: 11155111)"
	@echo "                     CHAIN_NAME       (default: sepolia)"
	@echo "                     REGISTRY_ADDRESS (optional)"
	@echo "                     START_BLOCK      (optional)"
	@echo "                   ui-build / ui-dev call ui-config automatically when"
	@echo "                   RPC_URL is set on the command line."
	@echo ""
	@echo "Development Commands:"
	@echo "  build            Build all Go binaries (node, runner, circuit compiler)"
	@echo "  test             Run fast Go unit tests (no chain, no Docker)"
	@echo "  test-integration Run heavy chain-backed integration tests via Docker"
	@echo ""

# ── Circuit update pipeline ───────────────────────────────────────────────

circuits-compile: ## Compile all circuits; write artifacts and Solidity verifier stubs
	@echo "Compiling circuits → $(ARTIFACTS_DIR) ..."
	@mkdir -p $(ARTIFACTS_DIR)
	go run ./cmd/circuit-compile \
		--destination=$(ARTIFACTS_DIR) \
		--verifiers-dir=solidity/src/verifiers \
		--output-json=$(CIRCUIT_ARTIFACTS_JSON)
	@echo ""
	@echo "Artifacts written to $(ARTIFACTS_DIR)"
	@echo "Solidity verifiers updated in solidity/src/verifiers"
	@echo "Hash JSON saved to $(CIRCUIT_ARTIFACTS_JSON)"
	@echo "Next step: make circuits-update-hashes"

circuits-update-hashes: ## Patch config/circuit_artifacts.go with hashes from the last compile
	@[ -f $(CIRCUIT_ARTIFACTS_JSON) ] || \
		{ echo "error: $(CIRCUIT_ARTIFACTS_JSON) not found — run 'make circuits-compile' first"; exit 1; }
	@echo "Patching config/circuit_artifacts.go ..."
	@bash scripts/update-circuit-hashes.sh $(CIRCUIT_ARTIFACTS_JSON) config/circuit_artifacts.go
	@echo "Next step: make solidity-build"

solidity-build: ## Build Solidity contracts with Foundry
	@echo "Building Solidity contracts ..."
	@cd solidity && forge build --force

solidity-bind: ## Regenerate Go ABI bindings from compiled Solidity artifacts
	@echo "Generating Go ABI bindings ..."
	@cd solidity && bash go_bind.sh

solidity-deploy: ## Deploy contracts (set RPC_URL, CHAIN_ID, PRIVATE_KEY)
	@# The script also loads RPC_URL / CHAIN_ID / PRIVATE_KEY / ETHERSCAN_API_KEY
	@# from a repo-root .env file when present, so the Makefile only enforces
	@# explicit overrides on the command line.
	@echo "Running solidity/deploy_all.sh ..."
	@bash solidity/deploy_all.sh

circuits: circuits-compile circuits-update-hashes solidity-build solidity-bind ## Full circuit update pipeline
	@echo ""
	@echo "Circuit update complete."

circuits-release: ## Compile circuits, upload to CDN, update hashes, rebuild Solidity
	@[ -n "$(S3_ACCESS_KEY)" ] || { echo "error: S3_ACCESS_KEY is not set"; exit 1; }
	@[ -n "$(S3_SECRET_KEY)" ] || { echo "error: S3_SECRET_KEY is not set"; exit 1; }
	@echo "Compiling circuits and uploading to CDN (bucket: $(S3_BUCKET)) ..."
	@mkdir -p $(ARTIFACTS_DIR)
	go run ./cmd/circuit-compile \
		--destination=$(ARTIFACTS_DIR) \
		--verifiers-dir=solidity/src/verifiers \
		--output-json=$(CIRCUIT_ARTIFACTS_JSON) \
		--s3.enabled \
		--s3.access-key=$(S3_ACCESS_KEY) \
		--s3.secret-key=$(S3_SECRET_KEY) \
		--s3.space=$(S3_SPACE) \
		--s3.bucket=$(S3_BUCKET)
	@echo ""
	@echo "Patching config/circuit_artifacts.go ..."
	@bash scripts/update-circuit-hashes.sh $(CIRCUIT_ARTIFACTS_JSON) config/circuit_artifacts.go
	@$(MAKE) solidity-build solidity-bind
	@echo ""
	@echo "Release complete — commit config/circuit_artifacts.go and the updated Solidity files."
	@echo "Commit config/circuit_artifacts.go, solidity/src/verifiers/*, and solidity/bindings/*"

# ── Testnet ───────────────────────────────────────────────────────────────

testnet-up: ## Start the testnet with N nodes
	@echo "Starting testnet with $(DKG_NODE_COUNT) nodes..."
	@cd testnet && \
	DKG_NODE_COUNT=$(DKG_NODE_COUNT) \
	docker compose up -d --scale dkg-node=$(DKG_NODE_COUNT) --build

testnet-run: ## Run the DKG orchestration scenario
	@echo "Running DKG scenario (nodes=$(DKG_NODE_COUNT), threshold=$(DKG_THRESHOLD), disclosure=$(DKG_DISCLOSURE_ALLOWED))..."
	@cd testnet && \
	DKG_RUNNER_NODES=$(DKG_NODE_COUNT) \
	DKG_RUNNER_THRESHOLD=$(DKG_THRESHOLD) \
	DKG_RUNNER_DISCLOSURE_ALLOWED=$(DKG_DISCLOSURE_ALLOWED) \
	docker compose run --rm dkg-runner

testnet-logs: ## Tail logs for the DKG nodes
	@cd testnet && docker compose logs -f dkg-node

testnet-down: ## Stop the testnet and wipe state
	@echo "Tearing down testnet..."
	@cd testnet && docker compose down -v

# ── UI ────────────────────────────────────────────────────────────────────
#
# The UI is fully decoupled from the Go binary — it ships as its own
# Docker image (ui/Dockerfile → ghcr.io/vocdoni/davinci-dkg-ui) and talks
# to the chain directly via RPC. None of the targets here are a build
# prerequisite for the Go binaries.

ui-install: ## Install UI dependencies with pnpm
	@echo "Installing UI dependencies ..."
	@cd ui && pnpm install

# Render ui/public/config.json from RPC_URL / MANAGER_ADDRESS / CHAIN_ID /
# CHAIN_NAME / REGISTRY_ADDRESS / START_BLOCK env vars. Idempotent and safe
# to call repeatedly. Run as a prerequisite by ui-build / ui-dev whenever
# the user passes a non-empty RPC_URL on the command line; otherwise we
# leave the committed default config alone.
ui-config: ## Re-render ui/public/config.json from RPC_URL et al.
	@RPC_URL='$(RPC_URL)' \
	 MANAGER_ADDRESS='$(MANAGER_ADDRESS)' \
	 CHAIN_ID='$(CHAIN_ID)' \
	 CHAIN_NAME='$(CHAIN_NAME)' \
	 REGISTRY_ADDRESS='$(REGISTRY_ADDRESS)' \
	 START_BLOCK='$(START_BLOCK)' \
	 bash scripts/render-ui-config.sh ui/public/config.json

ui-build: ## Build the DKG explorer UI to ui/dist
	@echo "Building UI ..."
	@cd sdk && [ -d node_modules ] || pnpm install --frozen-lockfile
	@cd sdk && [ -d dist ] || pnpm run build
	@cd ui && [ -d node_modules ] || pnpm install
	@if [ -n "$(RPC_URL)" ]; then $(MAKE) --no-print-directory ui-config; fi
	@cd ui && pnpm run build
	@echo "UI built → ui/dist (consumed by ui/Dockerfile)"

ui-dev: ## Run the UI in Vite dev mode with hot reload
	@cd sdk && [ -d node_modules ] || pnpm install --frozen-lockfile
	@cd sdk && [ -d dist ] || pnpm run build
	@cd ui && [ -d node_modules ] || pnpm install
	@if [ -n "$(RPC_URL)" ]; then $(MAKE) --no-print-directory ui-config; fi
	@cd ui && pnpm run dev

ui-test: ## Run UI unit tests (vitest)
	@cd ui && [ -d node_modules ] || pnpm install
	@cd ui && pnpm run test

ui-clean: ## Remove UI build output and node_modules
	@rm -rf ui/dist ui/node_modules

# ── Development ───────────────────────────────────────────────────────────

build: ## Build all Go binaries
	@echo "Building binaries..."
	go build ./cmd/...

test: ## Run fast unit tests
	@echo "Running unit tests..."
	go test -v $$(go list ./... | grep -v github.com/vocdoni/davinci-dkg/tests) -timeout=30m -failfast

test-integration: ## Run heavy integration tests
	@echo "Running integration tests..."
	RUN_INTEGRATION_TESTS=true go test -v ./tests/... -timeout=2h -failfast -count=1
