# Multi-stage build producing two binaries:
#   davinci-dkg-node  – long-running DKG participant daemon
#   dkg-runner        – testnet scenario orchestrator
#
# The node binary is UI-blind: the explorer ships as a separate image
# (ui/Dockerfile → ghcr.io/vocdoni/davinci-dkg-ui). This image therefore
# only needs the Go toolchain — no Node, no pnpm.

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

ENTRYPOINT ["/app/davinci-dkg-node"]
