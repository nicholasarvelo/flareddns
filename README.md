![flareDDNS Logo](https://raw.githubusercontent.com/nicholasarvelo/flareddns/main/.assets/logo.png)

# flareDDNS

> A simple, Docker-friendly Dynamic DNS client for Cloudflare written in Go.

---

## Overview

One of my favorite features of Google Domains was its native support for Dynamic DNS at no additional cost. Sadly, all good things [come to an end](https://www.theverge.com/2023/6/16/23763340/google-domains-sunset-sell-squarespace). After moving all my domains to Cloudflare, I realized they didn’t offer DDNS support. Unimpressed by the alternatives, I decided to build my own.

**`flareDDNS`** is a lightweight Go application designed to run in Docker, keep your Cloudflare DNS records updated with your public IP, and provide clear, reliable logging.

---

## Features

- Polls your public IP address at a configurable interval.
- Creates a DNS record if one doesn’t exist.
- Updates existing records if the IP has changed.
- No-op if the DNS record already matches your current IP.
- Runs cleanly in Docker with environment-based configuration.
- Provides simple but informative logs.

---

## Build and Run Locally

### Prerequisites

- Go 1.20+ installed ([Download Go](https://golang.org/dl/))
- A [Cloudflare account](https://cloudflare.com)
- A [Cloudflare API token](https://developers.cloudflare.com/fundamentals/api/get-started/create-token/)

### Steps

```bash
# Clone the repository
git clone https://github.com/nicholasarvelo/flareddns.git
cd flareddns

# Build the binary
go build -o flareddns ./cmd/flareddns

# Set required environment variables
export CF_API_TOKEN=your_cloudflare_api_token
export CF_DNS_RECORD_TYPE=A
export CF_ZONE_NAME=example.com

# Optional environment variables
export CF_DNS_RECORD=sub.example.com   # Default: null
export CF_POLLING_INTERVAL=240         # In minutes; default: 60
export CF_PROXIED=true                 # Default: false

# Run the application
./flareddns
```
## Running in Docker

```shell
docker run -it --rm \
  -e CF_API_TOKEN=your_cloudflare_api_token \
  -e CF_DNS_RECORD_TYPE=A \
  -e CF_ZONE_NAME=example.com \
  steptimeeditor/flareddns:latest
```
---

## Project Structure

| Path                                | Description                          |
|-------------------------------------|--------------------------------------|
| `cmd/flareddns/main.go`             | Application entry point              |
| `internal/client/cloudflare.go`     | Cloudflare API client logic          |
| `internal/config/env.go`            | Environment variable parsing         |
| `internal/dns/`                     | DNS record creation and update logic |
| `internal/netinfo/ip.go`            | Public IP address detection          |
| `internal/scheduler/cron.go`        | Scheduling and polling logic         |
| `internal/util/pointer.go`          | Utility helper functions             |

---

## License

See LICENSE for details.
