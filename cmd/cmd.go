package cmd

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func visit(keysearch string, links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == keysearch {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(keysearch, links, c)
	}
	return links
}

func FindTag(url, keysearch string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("failed %s: with %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("failed to parse %s as HTNL: %v", url, err)
	}
	return visit(keysearch, nil, doc), nil
}
