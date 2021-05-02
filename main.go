package main

import (
	"fmt"
	"os"

	"sitesearch/cmd"
)

func main() {
	text := os.Args[1:]

	links, err := cmd.FindTag(text[0], text[1], text[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	for _, link := range links {
		fmt.Println(link)
	}
}
