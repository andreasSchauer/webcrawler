package main
import (
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"strings"
	"fmt"
	"net/url"
)


func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	doc, err := html.Parse(strings.NewReader(htmlBody))
	if err != nil {
		return nil, fmt.Errorf("couldn't read HTML body")
	}

	var URLs []string

	for node := range doc.Descendants() {
		if node.Type == html.ElementNode && node.DataAtom == atom.A {
			for _, a := range node.Attr {
				if a.Key == "href" {
					URL, err := getAbsURL(a.Val, rawBaseURL)
					if err != nil {
						return nil, err
					}

					URLs = append(URLs, URL)
				}
			}
		}
	}

	return URLs, nil
}


func getAbsURL(URL, rawBaseURL string) (string, error) {
	if strings.HasPrefix(URL, rawBaseURL) {
		return URL, nil
	}
	
	u, err := url.Parse(URL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL")
	}

	base, err := url.Parse(rawBaseURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse base URL")
	}

	AbsURL := base.ResolveReference(u).String()

	return AbsURL, nil
}