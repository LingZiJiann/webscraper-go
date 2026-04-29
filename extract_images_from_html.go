package main

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("error parsing HTML: %w", err)
	}
	var image_urls []string
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if !exists || strings.TrimSpace(src) == "" {
			return
		}

		parsedURL, err := url.Parse(src)
		if err != nil {
			fmt.Printf("couldn't parse src %q: %v\n", src, err)
			return
		}
		resolvedURL := baseURL.ResolveReference(parsedURL)
		image_urls = append(image_urls, resolvedURL.String())
	})
	return image_urls, nil
}
