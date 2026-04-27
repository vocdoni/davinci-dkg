# davinci-dkg UI — multi-stage build producing a tiny nginx image that
# serves the Vite-built SPA. The UI talks to the chain directly via RPC;
# the davinci-dkg-node binary has no awareness of this image.

# ---------------------------------------------------------------------------
# SDK build stage — TypeScript library that the UI depends on via file:../sdk
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
# UI build stage — Vite produces /root/ui/dist
# ---------------------------------------------------------------------------
FROM node:20-bookworm-slim AS ui-builder
WORKDIR /root
RUN corepack enable
# The UI depends on the SDK via `file:../sdk`, so we have to mirror the
# repo layout: /root/sdk + /root/ui side-by-side.
COPY --from=sdk-builder /sdk ./sdk
WORKDIR /root/ui
COPY ui/package.json ui/pnpm-lock.yaml* ./
RUN --mount=type=cache,target=/root/.local/share/pnpm/store \
    pnpm install --frozen-lockfile || pnpm install
COPY ui/ ./
ARG VITE_BUILD_VERSION=dev
ENV VITE_BUILD_VERSION=${VITE_BUILD_VERSION}
RUN pnpm run build

# ---------------------------------------------------------------------------
# Runtime — nginx serves dist/. Single binary, no Node at runtime.
# ---------------------------------------------------------------------------
FROM nginx:alpine
WORKDIR /usr/share/nginx/html

# nginx default config replaced with our SPA-aware one.
COPY ui/docker/nginx.conf /etc/nginx/conf.d/default.conf
COPY ui/docker/entrypoint.sh /docker-entrypoint.d/40-render-config.sh
RUN chmod +x /docker-entrypoint.d/40-render-config.sh

# Copy the built SPA last so iterating on its content doesn't bust the
# nginx-config cache layer above.
COPY --from=ui-builder /root/ui/dist/ ./

EXPOSE 80
# nginx:alpine's default ENTRYPOINT runs every executable in
# /docker-entrypoint.d/ then calls `nginx -g 'daemon off;'` itself, so we
# don't need an ENTRYPOINT override here.
CMD ["nginx", "-g", "daemon off;"]
