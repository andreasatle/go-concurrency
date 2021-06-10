package main

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"golang.org/x/net/html"
)

var fetched map[string]bool
var counter int

type result struct {
	url   string
	urls  []string
	err   error
	depth int
}

func Crawl(url string, depth int) {
	ch := make(chan *result)
	fetch := func(url string, depth int) {
		urls, err := findLinks(url)
		ch <- &result{url, urls, err, depth}
	}
	go fetch(url, depth)
	fetched[url] = true
	for fetching := 1; fetching > 0; fetching-- {
		res := <-ch
		if res.err != nil {
			continue
		}
		counter++
		fmt.Printf("Count: %d, Depth: %d, Found: %s\n", counter, res.depth, res.url)
		if res.depth > 0 {
			for _, u := range res.urls {
				if !fetched[u] {
					fetching++
					go fetch(u, res.depth-1)
					fetched[u] = true
				}
			}
		}
	}
	close(ch)
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	fetched = make(map[string]bool) // Use make since module varaiable already declared
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
