package main

import (
	"net/url"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return rawURL, err
	}
	host := parsedURL.Host
	path := parsedURL.Path

	path = strings.TrimSuffix(path, "/")
	return host + path, nil
}
