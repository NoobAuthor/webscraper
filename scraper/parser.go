package scraper

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

func Parse(htmlContent string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		return nil, fmt.Errorf("failed to parse HTML: %v", err)
	}

	var links []string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key == "href" {
					links = append(links, attr.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)

	return links, nil
}

func FetchAndParse(url string) ([]string, error) {
	htmlContent, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	return Parse(htmlContent)
}
