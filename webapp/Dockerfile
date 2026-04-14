# Build from the repository root so both webapp/ and sdk/ are in context:
#
#   docker build -f webapp/Dockerfile -t davinci-dkg-webapp .
#
# ── Stage 1: build ────────────────────────────────────────────────────────
FROM node:24-alpine AS build

# Enable pnpm via corepack (ships with Node ≥ 16).
RUN corepack enable && corepack prepare pnpm@10.33.0 --activate

WORKDIR /build

# Copy the local SDK package first so pnpm can resolve the
# "file:../sdk" workspace dependency declared in webapp/package.json.
COPY sdk/package.json          sdk/pnpm-lock.yaml ./sdk/
COPY sdk/src                                       ./sdk/src/
COPY sdk/tsconfig.json                             ./sdk/

# Build the SDK so its dist/ is available before the webapp install.
RUN cd sdk && pnpm install --frozen-lockfile && pnpm build

# Copy webapp manifests and install dependencies (lockfile-frozen for
# reproducible builds; the resolved sdk path ../sdk points to /build/sdk).
COPY webapp/package.json       webapp/pnpm-lock.yaml ./webapp/
RUN cd webapp && pnpm install --frozen-lockfile

# Copy the rest of the webapp sources and build.
COPY webapp/index.html         ./webapp/
COPY webapp/vite.config.ts     ./webapp/
COPY webapp/tsconfig.json      ./webapp/
COPY webapp/src                ./webapp/src/

RUN cd webapp && pnpm build

# ── Stage 2: serve ────────────────────────────────────────────────────────
FROM nginx:1.27-alpine AS serve

# SPA routing: all paths fall through to index.html so react-router-dom
# handles client-side navigation.
RUN printf 'server {\n\
    listen       80;\n\
    root         /usr/share/nginx/html;\n\
    index        index.html;\n\
    location / {\n\
        try_files $uri $uri/ /index.html;\n\
    }\n\
    # Disable access logs for health-check noise reduction.\n\
    access_log   off;\n\
}\n' > /etc/nginx/conf.d/default.conf

COPY --from=build /build/webapp/dist /usr/share/nginx/html

EXPOSE 80
