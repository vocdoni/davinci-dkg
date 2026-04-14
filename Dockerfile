# Multi-stage build producing two binaries:
#   davinci-dkg-node  – long-running DKG participant daemon
#   dkg-runner        – testnet scenario orchestrator

# ---------------------------------------------------------------------------
# SDK build stage (TypeScript library consumed by the webapp)
# ---------------------------------------------------------------------------
FROM node:20-bookworm-slim AS sdk-builder
WORKDIR /sdk
RUN corepack enable
COPY sdk/package.json sdk/pnpm-lock.yaml* ./
RUN --mount=type=cache,target=/root/.local/share/pnpm/store \
    pnpm install --frozen-lockfile || pnpm install
COPY sdk/ ./
RUN pnpm run build

# ---------------------------------------------------------------------------
# Webapp build stage (produces webapp/dist consumed by go:embed)
# ---------------------------------------------------------------------------
FROM node:20-bookworm-slim AS webapp-builder
WORKDIR /root
RUN corepack enable
# Bring in the pre-built SDK so the file:../sdk dependency resolves correctly.
COPY --from=sdk-builder /sdk ./sdk
WORKDIR /root/webapp
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

# Copy go module files first so the download layer is cached independently
# of source changes.
COPY go.mod go.sum ./

RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download

COPY . .
# Bring in the pre-built webapp so the //go:embed directive in webapp/embed.go
# finds the dist/ directory during the Go build below.
COPY --from=webapp-builder /root/webapp/dist ./webapp/dist

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
