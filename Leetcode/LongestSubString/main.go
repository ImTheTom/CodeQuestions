package main

import (
	"fmt"
	"time"
)

func DoQuestion(s string) int {
	highest := 0

	for i := 0; i < len(s); i++ {
		seen := make(map[byte]struct{}, len(s)-i)
		for j := i; j < len(s); j++ {
			if _, exists := seen[s[j]]; exists {
				break
			}
			seen[s[j]] = struct{}{}
		}

		if len(seen) > highest {
			highest = len(seen)
		}

		if highest > len(s)-i {
			break
		}

	}

	return highest
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("abcabcbb")
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
