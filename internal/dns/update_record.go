package dns

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
	"github.com/nicholasarvelo/flareddns/internal/config"
	"github.com/nicholasarvelo/flareddns/internal/netinfo"
	"github.com/nicholasarvelo/flareddns/internal/util"
	"log"
	"time"
)

func UpdateRecord(
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
	comment := util.StringPointer(
		fmt.Sprintf("Updated by flareDDNS [%s]", timeStamp),
	)
	_, err = client.UpdateDNSRecord(
		ctx, zoneID, cloudflare.UpdateDNSRecordParams{
			Name:    clientConfig.RecordValue,
			Content: currentPublicIP,
			Comment: comment,
			ID:      zoneID.Identifier,
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
