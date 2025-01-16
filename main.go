package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	args := os.Args[1:]

	if len(args) < 3 {
		fmt.Println("not enough arguments provided")
		fmt.Println("usage: crawler <baseURL> <maxConcurrency> <maxPages>")
		return
	}

	if len(args) > 3 {
		fmt.Println("too many arguments provided")
		return
	}

	rawBaseURL := args[0]

	maxConcurrency, err := strconv.Atoi(args[1])
	if err != nil {
		fmt.Printf("Error - maxConcurrency: %v", err)
	}

	maxPages, err := strconv.Atoi(args[2])
	if err != nil {
		fmt.Printf("Error - maxPages: %v", err)
	}

	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
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