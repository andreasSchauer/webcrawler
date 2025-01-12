package main

import (
	"os"
	"fmt"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := args[0]
	fmt.Printf("starting crawl of: %s...", baseURL)
	htmlBody, err := getHTML(baseURL) 
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(htmlBody)
}