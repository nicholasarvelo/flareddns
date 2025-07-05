package config

import (
	"fmt"
)

// ClientConfig holds configuration values loaded from environment
// variables.
type ClientConfig struct {
	APIToken        string
	RecordType      string
	RecordValue     string
	ZoneName        string
	PollingInterval int
	Proxied         bool
}

// ParseVariables loads required and optional environment variables into a
// ClientConfig struct. It returns the populated ClientConfig and an error
// if any required variable is missing or invalid.
func ParseVariables() (ClientConfig, error) {
	var clientConfig ClientConfig

	requiredVariable, err := parseRequiredVariables()
	if err != nil {
		return clientConfig, fmt.Errorf("%w", err)
	}

	optionalVariable, err := parseOptionalVariables(requiredVariable)
	if err != nil {
		return clientConfig, fmt.Errorf("%w", err)
	}

	clientConfig = ClientConfig{
		APIToken:        requiredVariable.APIToken,
		RecordType:      requiredVariable.RecordType,
		RecordValue:     optionalVariable.RecordValue,
		ZoneName:        requiredVariable.ZoneName,
		PollingInterval: optionalVariable.PollingInterval,
		Proxied:         optionalVariable.Proxied,
	}

	return clientConfig, nil
}
