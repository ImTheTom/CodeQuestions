package main

import (
	"fmt"
	"time"
)

func DoQuestion(nums []int) int {
	digitCounts := make(map[int]int, len(nums))
	for _, v := range nums {
		digitCounts[v]++
	}

	total := 0

	for index, value := range digitCounts {
		if value == 1 {
			total += index
		}
	}

	return total
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 2, 3, 2,
	})
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
