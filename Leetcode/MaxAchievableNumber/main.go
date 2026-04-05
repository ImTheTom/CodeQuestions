package main

import (
	"fmt"
	"time"
)

func DoQuestion(num int, t int) int {
	additionalCount := t * 2

	total := num + additionalCount

	return total
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(4, 1)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
