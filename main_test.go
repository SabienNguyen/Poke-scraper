package main

import (
	"strings"
	"testing"
)

func TestWebScraper(t *testing.T) {
	expectedLinks := []string{"https://example.com/page1", "https://example.com/page2"}

	html := `<html><body><a href="https://example.com/page1">Page 1</a><a href="https://example.com/page2">Page 2</a></body></html>`

	for _, expectedLink := range expectedLinks {
		if !containsLink(html, expectedLink) {
			t.Errorf("Expected link %s not found in HTML", expectedLink)
		}
	}
}

func containsLink(html, link string) bool {
	return strings.Contains(html, link)
}
