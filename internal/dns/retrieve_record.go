package dns

import (
	"context"
	"fmt"
	"github.com/cloudflare/cloudflare-go"
)

type ZoneRecord struct {
	Value          string                        `json:"value"`
	ZoneIdentifier *cloudflare.ResourceContainer `json:"zone_identifier"`
}

// RetrieveRecord queries Cloudflare for a DNS record by name and zone name.
// Returns the first matching record found, or an error if none exist.
func RetrieveRecord(
	apiClient *cloudflare.API,
	recordName string,
	zoneName string,
) (ZoneRecord, error) {
	ctx := context.Background()

	zoneID, err := apiClient.ZoneIDByName(zoneName)
	if err != nil {
		return ZoneRecord{}, fmt.Errorf(
			"failed to retrieve Cloudflare zone ID: %w",
			err,
		)
	}

	records, _, err := apiClient.ListDNSRecords(
		ctx,
		cloudflare.ZoneIdentifier(zoneID),
		cloudflare.ListDNSRecordsParams{Name: recordName},
	)
	if err != nil {
		return ZoneRecord{}, fmt.Errorf("failed to list DNS records: %w", err)
	}

	if len(records) == 0 {
		return ZoneRecord{}, fmt.Errorf(
			"no DNS records found for name: %q",
			recordName,
		)
	}

	return ZoneRecord{
		Value:          records[0].Content,
		ZoneIdentifier: cloudflare.ZoneIdentifier(zoneID),
	}, nil
}
