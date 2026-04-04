package main

import (
	"fmt"
	"time"
)

func DoQuestion(str string) int {
	distinctStarts := make(map[rune]interface{})
	total := 0

	for _, s := range str {
		if _, exists := distinctStarts[s]; !exists {
			distinctStarts[s] = struct{}{}

			total++
		}
	}

	return total
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("abab")
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
