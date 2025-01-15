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

	rawBaseURL := args[0]

	const maxConcurrency = 3
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("Error - configure: %v", err)
		return
	}

	fmt.Printf("starting crawl of: %s...", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	for page, count := range cfg.pages {
		fmt.Printf("amount: %d, page: %s\n",count, page)
	}
}