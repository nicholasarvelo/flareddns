FROM golang:1.9.7

LABEL org.opencontainers.title="flareDDNS"
LABEL org.opencontainers.image.authors="nicholas@arvelo.dev"
LABEL org.opencontainers.description="Dynamic DNS Client for Cloudflare"
LABEL org.opencontainers.source="https://github.com/steptimeeditor/flareddns"

WORKDIR /usr/src/flareddns

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/flareddns ./...

CMD ["flareddns"]
