# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:1.25-alpine AS golang
ARG TARGETOS
ARG TARGETARCH
WORKDIR /backend
COPY ./backend .
RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -trimpath -ldflags="-s -w" -o rustdesk-api-server-pro .

FROM --platform=$BUILDPLATFORM node:22-alpine AS node
WORKDIR /frontend
COPY ./soybean-admin .
RUN npm install --global pnpm@9.15.9 \
    && pnpm install --frozen-lockfile \
    && pnpm build

FROM alpine:3.22
RUN apk add --no-cache ca-certificates tzdata \
    && addgroup -S app \
    && adduser -S -D -H -u 10001 -G app app \
    && mkdir -p /app/data \
    && chown -R app:app /app

WORKDIR /app
COPY --chown=app:app ./docker/start.sh .
COPY --chown=app:app --from=golang /backend/rustdesk-api-server-pro .
COPY --chown=app:app --from=node /frontend/dist ./dist

USER app
EXPOSE 8080
CMD ["sh", "/app/start.sh"]