package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		return
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		return
	}

	baseURL := args[0]
	fmt.Printf("starting crawl of: %s...", baseURL)
	pages := make(map[string]int)
	crawlPage(baseURL, baseURL, pages)

	for page, count := range pages {
		fmt.Printf("amount: %d, page: %s\n",count, page)
	}
}