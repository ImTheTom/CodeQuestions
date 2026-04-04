package main

import (
	"fmt"
	"time"
)

func DoQuestion(s string) int {
	total := 0

	for i := 0; i < len(s)-1; i++ {
		difference := int(s[i]) - int(s[i+1])

		if difference < 0 {
			difference *= -1
		}

		total += difference
	}

	return total
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("hello")
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
