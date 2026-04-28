# davinci-dkg UI — multi-stage build that produces the Vite static bundle
# only. The final stage carries no runtime (no nginx, no Node) — it's a
# build artifact image, intended for static hosts that build from source
# and read files out of `output_dir` (DigitalOcean App Platform, Cloudflare
# Pages with build-args, etc.) or for compose to bind-mount the dist into
# a stock nginx:alpine.
#
# Chain config is baked in at build time via the args below. Hosts that
# want a different deployment build a different image — there is no
# runtime config templating any more.

# ---------------------------------------------------------------------------
# SDK build stage — TypeScript library that the UI links to via link:../sdk
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
# UI build stage — render config.json from build args, then `vite build`.
# ---------------------------------------------------------------------------
FROM node:20-bookworm-slim AS ui-builder
WORKDIR /root
RUN corepack enable
# Mirror the repo layout so the link:../sdk dependency resolves.
COPY --from=sdk-builder /sdk ./sdk
COPY scripts/ ./scripts/

WORKDIR /root/ui
COPY ui/package.json ui/pnpm-lock.yaml* ./
RUN --mount=type=cache,target=/root/.local/share/pnpm/store \
    pnpm install --frozen-lockfile || pnpm install
COPY ui/ ./

# Chain config — every value has a sensible Sepolia default so a bare
# `docker build` produces a usable Sepolia bundle. Override on the build
# command line (or via DO BUILD_TIME envs) to target any deployment.
ARG RPC_URL=https://eth-sepolia.public.blastapi.io
ARG MANAGER_ADDRESS=0xd3ef727b695b21e108497c36f9dcec52d741298a
ARG CHAIN_ID=11155111
ARG CHAIN_NAME=sepolia
ARG REGISTRY_ADDRESS=
ARG START_BLOCK=
ARG VITE_BUILD_VERSION=dev
ARG VITE_WALLETCONNECT_PROJECT_ID=

ENV RPC_URL=${RPC_URL} \
    MANAGER_ADDRESS=${MANAGER_ADDRESS} \
    CHAIN_ID=${CHAIN_ID} \
    CHAIN_NAME=${CHAIN_NAME} \
    REGISTRY_ADDRESS=${REGISTRY_ADDRESS} \
    START_BLOCK=${START_BLOCK} \
    VITE_BUILD_VERSION=${VITE_BUILD_VERSION} \
    VITE_WALLETCONNECT_PROJECT_ID=${VITE_WALLETCONNECT_PROJECT_ID}

# Template public/config.json from the build args before vite picks it
# up (Vite copies public/ into dist/ verbatim, so the rendered file
# becomes /config.json in the served bundle).
RUN bash /root/scripts/render-ui-config.sh public/config.json

RUN pnpm run build

# ---------------------------------------------------------------------------
# Final stage — just the static files. No nginx, no shell, no runtime.
#
# Using `nginx:alpine` *only* so the output dir (/usr/share/nginx/html)
# matches the convention DO's static_sites docs assume. The actual nginx
# binary is never invoked here — DO (and compose) read files out of the
# image, they don't run it. busybox/alpine would also work; the choice of
# base layer is invisible to consumers.
# ---------------------------------------------------------------------------
FROM nginx:alpine AS static
COPY --from=ui-builder /root/ui/dist/ /usr/share/nginx/html/
