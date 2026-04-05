package main

import (
	"fmt"
	"time"
)

func DoQuestion(first int, second int) int {
	totalDivisible := 0
	totalNonDivisible := 0

	check := 1

	for {
		if check > first {
			break
		}

		if check%second != 0 {
			totalNonDivisible += check
		} else {
			totalDivisible += check
		}

		check++
	}

	return totalNonDivisible - totalDivisible
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(10, 3)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
