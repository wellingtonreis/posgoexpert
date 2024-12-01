package main

import (
	"flag"
	"fmt"
	infra "load-test/internal/infra"
)

type Result struct {
	StatusCode int
}

func main() {
	url := flag.String("url", "", "URL of the service to test")
	requests := flag.Int("requests", 10, "Total number of requests")
	concurrency := flag.Int("concurrency", 1, "Number of simultaneous calls")
	flag.Parse()

	if *url == "" || *requests <= 0 || *concurrency <= 0 {
		fmt.Println("Invalid parameters. Use --url, --requests and --concurrency correctly.	")
		return
	}
	results := infra.RunLoadTest(*url, *requests, *concurrency)
	infra.Generate(results)
}
