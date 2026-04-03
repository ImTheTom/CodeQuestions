package main

import (
	"fmt"
	"time"
)

func DoQuestion(nums []int, k int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total % k
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]int{
		3, 9, 7,
	}, 5)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
