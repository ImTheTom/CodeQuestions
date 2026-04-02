package main

import (
	"fmt"
	"time"
)

func DoQuestion(nums []int, val int) int {
	currentIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[currentIndex] = nums[i]
			currentIndex++
			continue
		}
	}

	return currentIndex
}

func main() {
	start := time.Now()

	nums := []int{
		0, 1, 2, 2, 3, 0, 4, 2,
	}

	fmt.Println("Running main")
	result := DoQuestion(nums, 2)
	fmt.Printf("Got %v\n", result)

	fmt.Println(nums)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
