package config

import (
	"fmt"
	"github.com/nicholasarvelo/flareddns/internal/util"
	"log"
	"os"
)

func parseRequiredVariables() (ClientConfig, error) {
	var clientConfig ClientConfig

	required := map[string]*string{
		"CF_API_TOKEN":       &clientConfig.APIToken,
		"CF_DNS_RECORD_TYPE": &clientConfig.RecordType,
		"CF_ZONE_NAME":       &clientConfig.ZoneName,
	}

	for keyName, keyType := range required {
		keyValue, defined := os.LookupEnv(keyName)
		if !defined || keyValue == "" {
			return clientConfig, fmt.Errorf(
				"missing required environment variable: %q",
				keyName,
			)
		}
		*keyType = keyValue

		if keyName == "CF_API_TOKEN" {
			log.Printf("%q set to %q", keyName, util.ObfuscateVariable(keyValue))
		} else {
			log.Printf("%q set to %q", keyName, keyValue)
		}
	}
	return clientConfig, nil
}
