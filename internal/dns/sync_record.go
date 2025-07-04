package dns

import (
	"github.com/cloudflare/cloudflare-go"
	"github.com/nicholasarvelo/flareddns/internal/config"
	"github.com/nicholasarvelo/flareddns/internal/netinfo"
	"log"
)

// SyncDNSRecord compares the current public IP with the DNS record and
// creates or updates the record on Cloudflare as needed.
func SyncDNSRecord(client *cloudflare.API, cfg config.ClientConfig) {
	currentIP, err := netinfo.QueryPublicIP(cfg.RecordType)
	if err != nil {
		log.Printf("Failed to get public IP: %v", err)
		return
	}

	record, err := RetrieveRecord(client, cfg.RecordValue, cfg.ZoneName)
	if err != nil {
		log.Printf("Failed to retrieve DNS record: %v", err)
		return
	}

	switch {
	case record.Value == "":
		log.Println("DNS record not found. Creating...")
		if err := CreateRecord(
			client,
			cfg,
			record.ZoneIdentifier,
		); err != nil {
			log.Printf("Failed to create record: %v", err)
		}
	case record.Value != currentIP:
		log.Println("DNS record outdated. Updating...")
		if err := UpdateRecord(
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
