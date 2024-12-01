package infra

import "fmt"

func Generate(results []Result) {
	total := len(results)
	status200 := 0
	statusCodes := make(map[int]int)

	for _, result := range results {
		if result.StatusCode == 200 {
			status200++
		}
		statusCodes[result.StatusCode]++
	}

	fmt.Println("\n--- Report ---")
	fmt.Printf("Total requests: %d\n", total)
	fmt.Printf("Status 200: %d\n", status200)
	fmt.Println("Distribution of status codes:")
	for code, count := range statusCodes {
		fmt.Printf("  %d: %d\n", code, count)
	}
}
