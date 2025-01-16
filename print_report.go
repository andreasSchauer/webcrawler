package main

import (
	"fmt"
	"slices"
	"strings"
)


type page struct {
	name	string
	visits	int
}


func printReport(pages map[string]int, baseURL string) {
	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	pagesSlice := getPagesSlice(pages)

	for _, page := range pagesSlice {
		fmt.Printf("Found %d internal links to %s\n", page.visits, page.name)
	}
}


func getPagesSlice(pages map[string]int) []page {
	var pagesSlice []page

	for url, amount := range pages {
		page := page{
			name: url,
			visits: amount,
		}
		pagesSlice = append(pagesSlice, page)
	}

	return sortPages(pagesSlice)
}


func sortPages(pagesSlice []page) []page {
	slices.SortFunc(pagesSlice, func(a, b page) int {
		if n := sortByVisitDesc(a.visits, b.visits); n != 0 {
			return n
		}

		return strings.Compare(a.name, b.name)
	})

	return pagesSlice
}


func sortByVisitDesc(a, b int) int {
	if a < b {
		return 1
	}

	if a > b {
		return -1
	}

	return 0
}

/*

- slices.sortFunc
- need to write my own sorting function for descending pageviews
- inside sortFunc: first use custom function. if n != 0 return n; else (if they're equal), return strings.Compare
*/