package config

import (
	"fmt"
	"github.com/nicholasarvelo/flareddns/internal/util"
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
	var cfg ClientConfig

	required := map[string]*string{
		"CF_API_KEY":         &cfg.APIKey,
		"CF_DNS_RECORD_TYPE": &cfg.RecordType,
		"CF_ZONE_NAME":       &cfg.ZoneName,
	}

	for key, ref := range required {
		val, ok := os.LookupEnv(key)
		if !ok || val == "" {
			return cfg, fmt.Errorf("missing required environment variable: %s", key)
		}
		*ref = val

		if key == "CF_API_KEY" {
			log.Printf("%s set to %q", key, util.ObfuscateVariable(val))
		} else {
			log.Printf("%s set to %q", key, val)
		}
	}

	// Optional: CF_DNS_RECORD
	if val, ok := os.LookupEnv("CF_DNS_RECORD"); ok && val != "" {
		cfg.RecordValue = val
		log.Printf("CF_DNS_RECORD set to %q", val)
	} else {
		cfg.RecordValue = cfg.ZoneName
		log.Printf("CF_DNS_RECORD not set; using apex record %q", cfg.ZoneName)
	}

	// Optional: CF_POLLING_INTERVAL
	if val, ok := os.LookupEnv("CF_POLLING_INTERVAL"); ok && val != "" {
		polling, err := strconv.Atoi(val)
		if err != nil {
			return cfg, fmt.Errorf("invalid CF_POLLING_INTERVAL: %w", err)
		}
		cfg.PollingInterval = polling
		log.Printf("CF_POLLING_INTERVAL set to %d", polling)
	} else {
		cfg.PollingInterval = 60
		log.Printf("CF_POLLING_INTERVAL not set; using default %d", cfg.PollingInterval)
	}

	// Optional: CF_PROXIED
	if val, ok := os.LookupEnv("CF_PROXIED"); ok && val != "" {
		proxied, err := strconv.ParseBool(val)
		if err != nil {
			return cfg, fmt.Errorf("invalid CF_PROXIED: %w", err)
		}
		cfg.Proxied = proxied
		log.Printf("CF_PROXIED set to %t", proxied)
	} else {
		cfg.Proxied = false
		log.Printf("CF_PROXIED not set; using default %t", cfg.Proxied)
	}

	return cfg, nil
}
