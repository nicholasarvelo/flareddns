# Build Stage
FROM golang:1.24-alpine AS builder

LABEL org.opencontainers.image.title="flareDDNS"
LABEL org.opencontainers.image.authors="nicholas@arvelo.dev"
LABEL org.opencontainers.image.description="Dynamic DNS Client for Cloudflare"
LABEL org.opencontainers.image.source="https://github.com/steptimeeditor/flareddns"

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux

RUN go build -v -o flareddns ./cmd/flareddns

# Runtime Stage
FROM alpine:latest

RUN apk --no-cache add ca-certificates

COPY --from=builder /app/flareddns /usr/local/bin/flareddns

ENTRYPOINT ["/usr/local/bin/flareddns"]
