package scraper

import (
	"net/http"
	"strings"
)

func CheckLinks(links []string) []string {
	var deadLinks []string

	for _, link := range links {
		if !strings.HasPrefix(link, "http") {
			continue // skip non-http links
		}

		resp, err := http.Head(link)
		if err != nil || resp.StatusCode >= 400 {
			deadLinks = append(deadLinks, link)
		}
	}

	return deadLinks
}
