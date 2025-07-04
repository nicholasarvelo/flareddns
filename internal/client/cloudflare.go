package client

import (
	"github.com/cloudflare/cloudflare-go"
	"log"
)

// CreateCloudflareClient initializes and returns a new Cloudflare API client
// using the provided API token. This function logs a fatal error and exits
// the program if the client cannot be created.
func CreateCloudflareClient(apiToken string) *cloudflare.API {
	client, err := cloudflare.NewWithAPIToken(apiToken)
	if err != nil {
		log.Fatalf("Failed to create Cloudflare client: %v", err)
	}
	return client
}
