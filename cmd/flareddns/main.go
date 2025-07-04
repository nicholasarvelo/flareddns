package main

import (
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/nicholasarvelo/flareddns/internal/client"
	"github.com/nicholasarvelo/flareddns/internal/config"
	"github.com/nicholasarvelo/flareddns/internal/dns"
	"github.com/nicholasarvelo/flareddns/internal/scheduler"
	"github.com/nicholasarvelo/flareddns/internal/ui"
	"log"
	"runtime"
)

func main() {
	ui.PrintBanner()
	clientConfig := loadConfig()
	cloudflareClient := client.CreateCloudflareClient(clientConfig.APIToken)
	log.Println("flareDDNS started")
	initialRecordSync(cloudflareClient, clientConfig)
	cronSchedule := fmt.Sprintf("@every %dm", clientConfig.PollingInterval)

	scheduler.StartCronJob(cronSchedule, cloudflareClient, clientConfig)
	runtime.Goexit()
}

func loadConfig() config.ClientConfig {
	cfg, err := config.ParseVariables()
	if err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}
	return cfg
}

func initialRecordSync(
	client *cloudflare.API,
	clientConfig config.ClientConfig,
) {
	log.Println("Running initial DNS record sync")
	dns.SyncDNSRecord(client, clientConfig)
	if clientConfig.PollingInterval > 1 {
		log.Printf(
			"flareDDNS will now poll every %d minutes",
			clientConfig.PollingInterval,
		)
	} else {
		log.Printf(
			"flareDDNS will now poll every %d minute",
			clientConfig.PollingInterval,
		)
	}
}
