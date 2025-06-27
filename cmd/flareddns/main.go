package main

import (
	"fmt"
	"github.com/nicholasarvelo/flareddns/internal/client"
	"github.com/nicholasarvelo/flareddns/internal/config"
	"github.com/nicholasarvelo/flareddns/internal/scheduler"
	"log"
	"runtime"
)

func main() {
	clientConfig := loadConfig()
	cloudflareClient := client.CreateCloudflareClient(clientConfig.APIKey)
	cronSchedule := fmt.Sprintf("@every %dm", clientConfig.PollingInterval)

	scheduler.StartCronJob(cronSchedule, cloudflareClient, clientConfig)
	log.Println("flareDDNS started")
	runtime.Goexit()
}

func loadConfig() config.ClientConfig {
	cfg, err := config.ParseVariables()
	if err != nil {
		log.Fatalf("Failed to parse environment variables: %v", err)
	}
	return cfg
}
