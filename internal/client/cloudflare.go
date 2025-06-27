package client

import (
	"github.com/cloudflare/cloudflare-go"
	"log"
)

func CreateCloudflareClient(apiKey string) *cloudflare.API {
	client, err := cloudflare.NewWithAPIToken(apiKey)
	if err != nil {
		log.Fatalf("Failed to create Cloudflare client: %v", err)
	}
	return client
}
