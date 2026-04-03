package main

import (
	"fmt"
	"time"
)

func DoQuestion(digits []int) []int {
	currentIdx := len(digits) - 1

	for {
		if currentIdx < 0 {
			break
		}

		current := digits[currentIdx]
		current++

		if current != 10 {
			digits[currentIdx] = current
			break
		}

		digits[currentIdx] = 0

		currentIdx--
	}

	if digits[0] == 0 {
		newNums := []int{1}

		digits = append(newNums, digits...)
	}

	return digits
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]int{9, 9, 9})
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
