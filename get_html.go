package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	// Create a new HTTP request
	request, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create request: %v", err)
	}

	// Set the User-Agent header
	request.Header.Set("User-Agent", "BootCrawler/1.0")

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return "", fmt.Errorf("got network error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return "", fmt.Errorf("error fetching %s: status code %d", rawURL, resp.StatusCode)
	}

	// Check Content-Type header
	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("got non-HTML response: %s", contentType)
	}

	// Read the response body
	htmlBodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("couldn't read response body: %v", err)
	}
	htmlBody := string(htmlBodyBytes)
	return htmlBody, nil
}
