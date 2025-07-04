package scheduler

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/nicholasarvelo/flareddns/internal/config"
	"github.com/nicholasarvelo/flareddns/internal/dns"
	"github.com/nicholasarvelo/flareddns/internal/netinfo"
	"github.com/robfig/cron/v3"
	"log"
)

func syncDNSRecord(client *cloudflare.API, cfg config.ClientConfig) {
	currentIP, err := netinfo.QueryPublicIP(cfg.RecordType)
	if err != nil {
		log.Printf("Failed to get public IP: %v", err)
		return
	}

	record, err := dns.RetrieveRecord(client, cfg.RecordValue, cfg.ZoneName)
	if err != nil {
		log.Printf("Failed to retrieve DNS record: %v", err)
		return
	}

	switch {
	case record.Value == "":
		log.Println("DNS record not found. Creating...")
		if err := dns.CreateRecord(
			client,
			cfg,
			record.ZoneIdentifier,
		); err != nil {
			log.Printf("Failed to create record: %v", err)
		}
	case record.Value != currentIP:
		log.Println("DNS record outdated. Updating...")
		if err := dns.UpdateRecord(
			client,
			cfg,
			record.ZoneIdentifier,
		); err != nil {
			log.Printf("Failed to update record: %v", err)
		}
	default:
		log.Printf(
			"Record valid: %q is already resolving to %q",
			cfg.RecordValue,
			currentIP,
		)
	}
}

func StartCronJob(
	schedule string,
	client *cloudflare.API,
	cfg config.ClientConfig,
) {
	cronJob := cron.New()
	_, err := cronJob.AddFunc(
		schedule, func() {
			syncDNSRecord(client, cfg)
		},
	)
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}
	cronJob.Start()
}
