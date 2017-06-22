package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

type crawlResult struct {
	depth int
	urls  []string
}

func crawlOne(url string, depth int, ch chan crawlResult) {
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("found: %s %q\n", url, body)
	}
	ch <- crawlResult{depth: depth - 1, urls: urls}
	return
}

// Crawl creates map for caching. Takes channel recieves
// from crawlOne with depth, and urls. Checking for the
// unique url in urlsCache.
func Crawl(url string, depth int, fetcher Fetcher) {
	urlsCache := make(map[string]bool)

	ch := make(chan crawlResult)

	go crawlOne(url, depth, ch)
	urlsCache[url] = true

	for i := depth; i >= 0; i-- {
		cr := <-ch
		for _, u := range cr.urls {
			if _, ok := urlsCache[u]; !ok {
				go crawlOne(u, cr.depth, ch)
				urlsCache[u] = true
			}
		}
	}
	return
}

// TODO: Fetch URLs in parallel.
// TODO: Don't fetch the same URL twice.
// This implementation doesn't do either:
func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

// Fetch fetches taken url from fetcher and returns its body
// and embedding urls. In other case returns an error.
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
