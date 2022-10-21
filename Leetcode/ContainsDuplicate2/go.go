package main

import (
	"fmt"
	"math"
	"time"
)

func DoQuestion(nums []int, k int) bool {
	counts := CountOccurrences(nums)

	for _, arr := range counts {
		if len(arr) < 2 {
			continue
		}
		if CheckIndexes(arr, k) {
			return true
		}
	}
	return false
}

func CountOccurrences(nums []int) map[int][]int {
	counts := make(map[int][]int, len(nums))
	for i, v := range nums {
		counts[v] = append(counts[v], i)
	}
	return counts
}

func CheckIndexes(nums []int, k int) bool {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			first, second := nums[i], nums[j]
			if math.Abs(float64(first-second)) <= float64(k) {
				return true
			}
		}
	}
	return false
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 2, 3, 1,
	}, 3)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
