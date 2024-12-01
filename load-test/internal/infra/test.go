package infra

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Result struct {
	StatusCode int
}

func RunLoadTest(url string, totalRequests, concurrency int) []Result {
	var wg sync.WaitGroup
	results := make([]Result, 0, totalRequests)
	resultsChan := make(chan Result, totalRequests)

	startTime := time.Now()

	go func() {
		for result := range resultsChan {
			results = append(results, result)
		}
	}()

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < totalRequests/concurrency; j++ {
				resp, err := http.Get(url)
				if err != nil {
					resultsChan <- Result{StatusCode: 0}
					continue
				}
				resultsChan <- Result{StatusCode: resp.StatusCode}
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()
	close(resultsChan)
	totalTime := time.Since(startTime)

	fmt.Printf("\nTotal time: %v\n", totalTime)
	return results
}
