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
	if recordType == "A" {
		protocol = "ipv4"
	} else if recordType == "AAAA" {
		protocol = "ipv6"
	}
	url := fmt.Sprintf("https://%s.icanhazip.com", protocol)
	request, errorOccurred := http.Get(url)
	if errorOccurred != nil {
		log.Println(errorOccurred)
	}
	// It's good practice to close the response body after processing the
	// response. This ensures that any associated network resources are released
	// and also prevent resource leaks.
	defer func() {
		if errorOccurred := request.Body.Close(); errorOccurred != nil {
			log.Printf("Failed to close response body: %v", errorOccurred)
		}
	}()

	// Read the entire response body from the HTTP request and stores it.
	response, errorOccurred := io.ReadAll(request.Body)
	if errorOccurred != nil {
		log.Println(errorOccurred)
	}

	currentPublicIP := strings.TrimRight(string(response), "\n")

	return currentPublicIP, nil
}
