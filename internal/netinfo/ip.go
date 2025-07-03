package netinfo

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func QueryPublicIP(recordType string) (string, error) {
	var protocol string
	switch recordType {
	case "A":
		protocol = "ipv4"
	case "AAAA":
		protocol = "ipv6"
	default:
		return "", fmt.Errorf("unsupported record type: %s", recordType)
	}

	url := fmt.Sprintf("https://%s.icanhazip.com", protocol)
	request, err := http.Get(url)
	if err != nil {
		log.Println(err)
	}
	// It's good practice to close the response body after processing the
	// response. This ensures that any associated network resources are released
	// and also prevent resource leaks.
	defer func() {
		if err := request.Body.Close(); err != nil {
			log.Printf("Failed to close response body: %v", err)
		}
	}()

	// Read the entire response body from the HTTP request and stores it.
	response, err := io.ReadAll(request.Body)
	if err != nil {
		log.Println(err)
	}

	currentPublicIP := strings.TrimRight(string(response), "\n")

	return currentPublicIP, nil
}
