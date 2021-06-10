package main

import (
	"fmt"
	"net/http"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool
var counter int

func Crawl(url string, depth int) {
	// Terminate when depth negative
	if depth < 0 {
		return
	}

	urls, err := findLinks(url)
	if err != nil {
		//log.Println("Something went wrong calling findLinks:", err)
		return
	}
	counter++
	fmt.Printf("Count: %d, Depth: %d, Found: %s\n", counter, depth, url)
	fetched[url] = true

	for _, u := range urls {
		if !fetched[u] {
			Crawl(u, depth-1)
		}
	}
}

func main() {
	fetched = make(map[string]bool) // Use make since module variable already declared
	start := time.Now()
	Crawl("http://andcloud.io", 2)
	fmt.Println("Elapsed Time:", time.Since(start))
}

func findLinks(url string) ([]string, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("getting url %s: %v", url, err)
	}
	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, res.Status)
	}
	doc, err := html.Parse(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

func visit(links []string, node *html.Node) []string {
	if node.Type == html.ElementNode && node.Data == "a" {
		for _, attr := range node.Attr {
			if attr.Key == "href" {
				links = append(links, attr.Val)
			}
		}
	}
	for c := node.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
