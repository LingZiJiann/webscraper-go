package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages              map[string]PageData // Keep track of the pages we've crawled
	baseURL            *url.URL            // Keep track of the original URL
	mu                 *sync.Mutex         // Ensure the pages map is thread-safe
	concurrencyControl chan struct{}       // Ensure we don't spawn too many goroutines at once
	wg                 *sync.WaitGroup     // Ensure the main function waits until goroutines are done
}

func (cfg *config) AddPageVisit(normalizedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		return false
	}
	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}
	return true
}

func (cfg *config) SetPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data
}

func configure(rawBaseURL string, maxConcurrency int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	return &config{
		pages:              make(map[string]PageData),
		baseURL:            baseURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrency),
		wg:                 &sync.WaitGroup{},
	}, nil
}
