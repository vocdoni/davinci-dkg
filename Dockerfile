# Multi-stage build producing two binaries:
#   davinci-dkg-node  – long-running DKG participant daemon
#   dkg-runner        – testnet scenario orchestrator

# ---------------------------------------------------------------------------
# Webapp build stage (produces webapp/dist consumed by go:embed)
# ---------------------------------------------------------------------------
FROM node:20-bookworm-slim AS webapp-builder
WORKDIR /webapp
RUN corepack enable
COPY webapp/package.json webapp/pnpm-lock.yaml* ./
RUN --mount=type=cache,target=/root/.local/share/pnpm/store \
    pnpm install --frozen-lockfile || pnpm install
COPY webapp/ ./
RUN pnpm run build

# ---------------------------------------------------------------------------
# Go build stage
# ---------------------------------------------------------------------------
FROM golang:1.25 AS builder
WORKDIR /src

# Copy go module files for both the main module and the local replace-directive
# modules (davinci-node/, gnark-crypto-primitives/) so the download layer
# is cached independently of source changes.
COPY go.mod go.sum ./
COPY davinci-node/go.mod davinci-node/go.sum ./davinci-node/
COPY gnark-crypto-primitives/go.mod gnark-crypto-primitives/go.sum ./gnark-crypto-primitives/

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
# Bring in the pre-built webapp so the //go:embed directive in webapp/embed.go
# finds the dist/ directory during the Go build below.
COPY --from=webapp-builder /webapp/dist ./webapp/dist

ARG VERSION=dev

# Build node daemon
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    go build -trimpath \
    -ldflags="-w -s -X=github.com/vocdoni/davinci-dkg/internal/version.Version=${VERSION}" \
    -o davinci-dkg-node \
    ./cmd/davinci-dkg-node

# Build scenario runner
RUN --mount=type=cache,target=/go/pkg/mod \
    --mount=type=cache,target=/root/.cache/go-build \
    CGO_ENABLED=0 \
    go build -trimpath \
    -ldflags="-w -s" \
    -o dkg-runner \
    ./cmd/dkg-runner

# ---------------------------------------------------------------------------
# Final minimal runtime image
# ---------------------------------------------------------------------------
FROM debian:bookworm-slim
WORKDIR /app

RUN apt-get update && \
    apt-get install --no-install-recommends -y ca-certificates && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

COPY --from=builder /src/davinci-dkg-node ./
COPY --from=builder /src/dkg-runner ./

EXPOSE 8081
ENTRYPOINT ["/app/davinci-dkg-node"]
