package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}
	var urls []string
	doc.Find("a[href]").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if !exists {
			return
		}
		href = strings.TrimSpace(href)
		if href == "" {
			return
		}

		parsedURL, err := url.Parse(href)
		if err != nil {
			// fmt.Printf("couldn't parse href %q: %v\n", href, err)
			return
		}
		resolvedURL := baseURL.ResolveReference(parsedURL)
		urls = append(urls, resolvedURL.String())
	})
	return urls, nil
}
