# - Builder Stage
# Compiles a statically linked binary for use in scratch
# or distroless containers.

FROM golang:alpine AS builder

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

# - Final Stage
# Builds the distroless image containing the binary.

FROM gcr.io/distroless/static:nonroot

COPY --from=builder /app/flareddns /flareddns

USER nonroot:nonroot
ENTRYPOINT ["/flareddns"]
