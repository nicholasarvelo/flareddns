package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type ClientConfig struct {
	APIKey          string
	RecordType      string
	RecordValue     string
	ZoneName        string
	PollingInterval int
	Proxied         bool
}

func ParseVariables() (ClientConfig, error) {
	requiredVariables := []string{
		"CF_API_KEY",
		"CF_DNS_RECORD_TYPE",
		"CF_ZONE_NAME",
	}

	clientConfig := ClientConfig{}
	for _, variable := range requiredVariables {
		val := os.Getenv(variable)
		if val == "" {
			return clientConfig, fmt.Errorf("environment variable '%s' is required", variable)
		}
	}

	if val := os.Getenv("CF_DNS_RECORD"); val != "" {
		log.Println("'CF_DNS_RECORD' undefined, proceeding with Apex record.")
		apexRecord := os.Getenv("CF_ZONE_NAME")
		clientConfig.RecordValue = apexRecord
	}

	if val := os.Getenv("CF_POLLING_INTERVAL"); val != "" {
		polling, err := strconv.Atoi(val)
		if err != nil {
			return clientConfig, fmt.Errorf("invalid value for 'CF_POLLING_INTERVAL': %w", err)
		}
		clientConfig.PollingInterval = polling
	} else {
		clientConfig.PollingInterval = 60
	}

	if val := os.Getenv("CF_PROXIED"); val != "" {
		proxied, err := strconv.ParseBool(val)
		if err != nil {
			return clientConfig, fmt.Errorf("invalid boolean value for 'CF_PROXIED': %w", err)
		}
		clientConfig.Proxied = proxied
	} else {
		clientConfig.Proxied = false
	}

	return clientConfig, nil
}
