package main

import (
	"fmt"
	"os"

	"sitesearch/cmd"
)

func main() {
	for _, url := range os.Args[1:] {
		links, err := cmd.FindTag(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
