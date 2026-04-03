package scheduler

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/nicholasarvelo/flareddns/internal/config"
	"github.com/nicholasarvelo/flareddns/internal/dns"
	"github.com/robfig/cron/v3"
	"log"
)

func StartCronJob(
	schedule string,
	client *cloudflare.API,
	cfg config.ClientConfig,
) {
	cronJob := cron.New()
	_, err := cronJob.AddFunc(
		schedule, func() {
			dns.SyncDNSRecord(client, cfg)
		},
	)
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}
	cronJob.Start()
}
