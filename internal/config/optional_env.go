// Package config provides functions for parsing required and optional
// environment variables into the ClientConfig struct for application
// configuration.
package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseOptionalVariables(config ClientConfig) (ClientConfig, error) {
	var clientConfig ClientConfig
	dnsRecord, err := parseDNSRecord(config.ZoneName)
	if err != nil {
		return clientConfig, fmt.Errorf("%w", err)
	}

	pollingInterval, err := parsePollingInterval()
	if err != nil {
		return clientConfig, fmt.Errorf("%w", err)
	}

	proxied, err := parseCFProxied()
	if err != nil {
		return clientConfig, fmt.Errorf("%w", err)
	}

	clientConfig = ClientConfig{
		RecordValue:     dnsRecord.RecordValue,
		PollingInterval: pollingInterval.PollingInterval,
		Proxied:         proxied.Proxied,
	}
	return clientConfig, nil
}

func parseDNSRecord(zoneName string) (ClientConfig, error) {
	var clientConfig ClientConfig
	apexRecord := zoneName
	value, defined := os.LookupEnv("CF_DNS_RECORD")
	if defined && value != "" {
		clientConfig.RecordValue = value
		log.Printf("\"CF_DNS_RECORD\" set to %q", value)
		return clientConfig, nil
	}
	clientConfig.RecordValue = apexRecord
	log.Printf(
		"\"CF_DNS_RECORD\" not set; using apex record %q",
		clientConfig.RecordValue,
	)
	return clientConfig, nil
}

func parsePollingInterval() (ClientConfig, error) {
	var clientConfig ClientConfig
	value, defined := os.LookupEnv("CF_POLLING_INTERVAL")
	if defined && value != "" {
		pollingInterval, err := strconv.Atoi(value)
		if err != nil {
			return clientConfig, fmt.Errorf(
				"invalid \"CF_POLLING_INTERVAL\" value: %w",
				err,
			)
		}
		clientConfig.PollingInterval = pollingInterval
		log.Printf(
			"\"CF_POLLING_INTERVAL\" set to %q",
			strconv.Itoa(clientConfig.PollingInterval),
		)
		return clientConfig, nil
	}
	clientConfig.PollingInterval = 60
	log.Printf(
		"\"CF_DNS_RECORD\" not set; using apex record %q",
		clientConfig.RecordValue,
	)
	return clientConfig, nil
}

func parseCFProxied() (ClientConfig, error) {
	var clientConfig ClientConfig
	value, defined := os.LookupEnv("CF_PROXIED")
	if defined && value != "" {
		proxied, err := strconv.ParseBool(value)
		if err != nil {
			return clientConfig, fmt.Errorf("invalid \"CF_PROXIED\" value: %w", err)
		}
		clientConfig.Proxied = proxied
		log.Printf("\"CF_PROXIED\" set to %q", value)
		return clientConfig, nil
	}
	clientConfig.Proxied = false
	log.Printf(
		"\"CF_PROXIED\" not set; using default %q",
		strconv.FormatBool(clientConfig.Proxied),
	)
	return clientConfig, nil
}
