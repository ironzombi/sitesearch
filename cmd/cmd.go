package cmd

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func visit(keysearch, tagtype string, links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == tagtype {
		for _, a := range n.Attr {
			if a.Key == keysearch {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(keysearch, tagtype, links, c)
	}
	return links
}

func FindTag(url, keysearch, tagtype string) ([]string, error) {
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
	return visit(keysearch, tagtype, nil, doc), nil
}

func CheckInput(text string) string {
	if strings.HasPrefix(text, "http") {
		return text
	} else {
		return fmt.Sprintf("http://%s", text)
	}
}
