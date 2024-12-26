package main

import (
	"net/http"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeResult struct {
	URL   string
	Title string
	Error error
}

func scrapeURL(url string) ScrapeResult {
	resp, err := http.Get(url)
	if err != nil {
		return ScrapeResult{URL: url, Error: err}
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return ScrapeResult{URL: url, Error: err}
	}

	title := doc.Find("title").Text()
	return ScrapeResult{URL: url, Title: title}
}

func worker(urls <-chan string, results chan<- ScrapeResult, wg *sync.WaitGroup) {
	defer wg.Done()
	for url := range urls {
		results <- scrapeURL(url)
	}
}

func startScraping(urls []string, numWorkers int) []ScrapeResult {
	var wg sync.WaitGroup
	urlChan := make(chan string, len(urls))
	resultChan := make(chan ScrapeResult, len(urls))

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go worker(urlChan, resultChan, &wg)
	}

	for _, url := range urls {
		urlChan <- url
	}
	close(urlChan)

	wg.Wait()
	close(resultChan)

	var results []ScrapeResult
	for result := range resultChan {
		results = append(results, result)
	}

	return results
}
