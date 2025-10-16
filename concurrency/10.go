package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	Fetch(url string) (body string, urls []string, err error)
}

type SafeFetcher struct {
	fetched map[string]bool
	mu sync.Mutex
}

func (sf *SafeFetcher) Visit (url string) bool {
	sf.mu.Lock()
	defer sf.mu.Unlock()
	if sf.fetched[url] {
		return false
	}
	sf.fetched[url] = true
	return true
}

func Crawl(url string, depth int, fetcher Fetcher, sf *SafeFetcher, wg *sync.WaitGroup) {
	defer wg.Done()
	
	if depth <= 0 {
		return
	}
	if !sf.Visit(url) {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found %s %q\n", url, body)

	for _, u := range urls {
		wg.Add(1)
		Crawl(u, depth-1, fetcher, sf, wg)
	}
	
}

func func10() {

	var wg sync.WaitGroup
	sf := &SafeFetcher{fetched: make(map[string]bool)}

	wg.Add(1)
	go Crawl("https://golang.org/", 4, fetcher, sf, &wg)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}