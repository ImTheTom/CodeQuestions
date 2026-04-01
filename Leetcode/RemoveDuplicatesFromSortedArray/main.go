package main

import (
	"fmt"
	"time"
)

func DoQuestionActual(nums []int) int {
	uniqueIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == nums[uniqueIndex] {
			continue
		}

		nums[uniqueIndex+1] = nums[i]
		uniqueIndex++
	}

	return uniqueIndex + 1
}

func DoQuestion(nums []int) int {
	total := 0

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] != nums[i+1] {
			total++
			continue
		}

		if i == 0 {
			nums = nums[i+1:]
			i--
			continue
		}

		nums = append(nums[:i], nums[i+1:]...)

		i--
	}

	return total + 1
}

func main() {
	start := time.Now()

	nums := []int{1, 1, 2}

	fmt.Println("Running main")
	result := DoQuestionActual(nums)
	fmt.Printf("Got %v\n", result)

	fmt.Println(nums)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
