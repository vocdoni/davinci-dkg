.PHONY: help \
        circuits-compile circuits-update-hashes circuits circuits-release \
        solidity-build solidity-bind solidity-deploy \
        testnet-up testnet-run testnet-down testnet-logs \
        webapp-install webapp-build webapp-dev webapp-clean \
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

# Deployment parameters (override on command line or via env)
RPC_URL     ?=
CHAIN_ID    ?=
PRIVATE_KEY ?=

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
	@echo "                  N dkg-node replicas + the dkg-webapp explorer."
	@echo "                  DKG_NODE_COUNT     (default 3, max 32 containers)"
	@echo "                  DKG_THRESHOLD      (default 2)"
	@echo "                  WEBAPP_PORT        (default 8081, host binding)"
	@echo "                  WEBAPP_PUBLIC_RPC  RPC URL advertised to browsers"
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
	@echo "Webapp Commands:"
	@echo "  webapp-install   Install webapp dependencies (pnpm install)"
	@echo "  webapp-build     Build the DKG explorer webapp (embedded into dkg-node)"
	@echo "  webapp-dev       Run the webapp in dev mode (Vite, hot reload on :5173)"
	@echo "  webapp-clean     Remove webapp/dist and webapp/node_modules"
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

# ── Webapp ────────────────────────────────────────────────────────────────

webapp-install: ## Install webapp dependencies with pnpm
	@echo "Installing webapp dependencies ..."
	@cd webapp && pnpm install

webapp-build: ## Build the embedded DKG explorer webapp
	@echo "Building webapp ..."
	@cd webapp && [ -d node_modules ] || pnpm install
	@cd webapp && pnpm run build
	@echo "Webapp built → webapp/dist (embedded into dkg-node on next 'make build')"

webapp-dev: ## Run the webapp in Vite dev mode with hot reload
	@cd webapp && [ -d node_modules ] || pnpm install
	@cd webapp && pnpm run dev

webapp-clean: ## Remove webapp build output and node_modules
	@rm -rf webapp/dist webapp/node_modules

# ── Development ───────────────────────────────────────────────────────────

build: webapp-build ## Build all binaries (rebuilds webapp first so it is embedded)
	@echo "Building binaries..."
	go build ./cmd/...

test: ## Run fast unit tests
	@echo "Running unit tests..."
	go test -v $$(go list ./... | grep -v github.com/vocdoni/davinci-dkg/tests) -timeout=30m -failfast

test-integration: ## Run heavy integration tests
	@echo "Running integration tests..."
	RUN_INTEGRATION_TESTS=true go test -v ./tests/... -timeout=2h -failfast -count=1
