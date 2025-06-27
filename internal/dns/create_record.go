package dns

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/nicholasarvelo/flareddns/internal/config"
	"github.com/nicholasarvelo/flareddns/internal/netinfo"
	"log"
	"time"
)

func CreateRecord(
	client cloudflare.API,
	clientConfig config.ClientConfig,
	zoneID *cloudflare.ResourceContainer,
) error {
	ctx := context.Background()
	currentPublicIP, err := netinfo.QueryPublicIP(clientConfig.RecordType)
	if err != nil {
		log.Printf("failed to retrieve public IP: %s", err)
	}

	timeStamp := time.Now().Format(time.DateTime)
	comment := fmt.Sprintf("Created by flareDDNS [%s]", timeStamp)
	_, err = client.CreateDNSRecord(
		ctx, zoneID, cloudflare.CreateDNSRecordParams{
			Type:      clientConfig.RecordType,
			Name:      clientConfig.RecordValue,
			Content:   currentPublicIP,
			Comment:   comment,
			Proxiable: true,
			Proxied:   &clientConfig.Proxied,
		},
	)
	if err != nil {
		log.Printf("Failed to create record: %s", err)
	}
	log.Printf(
		"Record Created: '%s' is resolving to '%s'",
		clientConfig.RecordValue,
		currentPublicIP,
	)

	return nil
}
