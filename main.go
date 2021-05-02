package main

import (
	"fmt"
	"os"

	"sitesearch/cmd"
)

func main() {
  reader := bufio.NewReader(os.Stdin)
  fmt.Print("\033[H\033[2J")
	for {
    fmt.Print("#>")
    text, _ := reader.ReadString('\n')
    text = strings.Replace(text, "\n", "", -1)
    command = strings.Fields(text)
		links, err := cmd.FindTag(command[0], command[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}
