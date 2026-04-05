package main

import (
	"fmt"
	"time"
)

func DoQuestion(nums []int) []int {
	result := make([]int, len(nums))

	for i := range nums {
		result[i] = nums[nums[i]]
	}

	return result
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]int{
		5, 0, 1, 2, 3, 4,
	})
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
