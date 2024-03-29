package main

import (
	"fmt"
	"os"
	"strings"

	"sitesearch/cmd"
)

func main() {
	text := os.Args[1:]
	url := cmd.CheckInput(text[0])
	tag := text[1]
	element := text[2]

	links, err := cmd.FindTag(url, tag, element)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
	for _, link := range links {
		if strings.HasPrefix(link, "http") {
			fmt.Println(link)
		} else {
			continue
		}
	}
}
