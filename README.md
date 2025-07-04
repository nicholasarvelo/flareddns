![flareDDNS Logo](https://raw.githubusercontent.com/nicholasarvelo/flareddns/main/.assets/logo.png)

## A dynamic DNS client written in Go

One of my favorite features of Google Domains was their native support for Dynamic DNS at no additional cost. It's a shame that all good things eventually [come to an end](https://www.theverge.com/2023/6/16/23763340/google-domains-sunset-sell-squarespace).

I moved all my domains over to Cloudflare, and that's all well and good; however, they do not offer any support for DDNS. I wasn't happy with the solutions I came across. I wanted something simple to deploy and configure with decent logging that quickly assures if it's working and, if so, what it's doing and, if not, what went wrong.

So I decided to write my own, which resulted in this Dynamic DNS client for Cloudflare, written in Go and designed to run as a Docker container.

## What Does It Do?

* `flareddns` runs as a scheduled job polling the public IP address of the host system at a user-defined interval.
* If the DNS record does not exist, a new record is created with the current public IP address. Whether it provisions a proxied record or not is also user-defined.
* If a DNS record already exists and the IP address associated with it is not the same as the current public IP, `flareddns` will update the record with the current IP address.
* If the DNS record's IP matches the public IP, no action is taken.

---

## Quick Start

### Building and Running locally

#### Prerequisites

* A Cloudflare account
* A user-generated [Cloudflare API token](https://developers.cloudflare.com/fundamentals/api/get-started/create-token/)
* Go 1.20+ installed ([Download Go](https://golang.org/dl/))

#### Steps to Build and Run

1. Clone the repository:
    ```shell
    git clone https://github.com/nicholasarvelo/flareddns.git
   ```
2. Change into the project directory:
    ```shell
    cd flareddns
    ```
3. Build the binary:
    ```shell
    go build -o flareddns ./cmd/flareddns
   ```
4. Set the required environment variables:
    ```shell
    export CF_API_KEY=dd109b2b7b0f1a3dfab1ad8b #Your Cloudflare API token
    export CF_DNS_RECORD_TYPE=A #DNS record type (A or AAAA)
    export CF_ZONE_NAME=dobbs.dev #Your Cloudflare zone name
    ```
5. Set optional environment variables:
    ```shell
   CF_DNS_RECORD=jr.dobbs.dev #DNS record to update (default: null)
   CF_POLLING_INTERVAL=240 #Polling interval in minutes (default: 60)
   CF_PROXIED=true #The proxy status of the DNS record (default: false)
   ```
6. Run the application:
    ```shell
    ./flareddns
    ```
---

### Code Structure

- `cmd/flareddns/main.go`: Application entry point.
- `internal/client/cloudflare.go`: Cloudflare API client logic.
- `internal/config/env.go`: Environment variable parsing and configuration.
- `internal/dns/`: DNS record creation, retrieval, and update logic.
- `internal/netinfo/ip.go`: Public IP address detection.
- `internal/scheduler/cron.go`: Scheduling and polling logic.
- `internal/util/pointer.go`: Utility functions.

## License

See `LICENSE` for details.
