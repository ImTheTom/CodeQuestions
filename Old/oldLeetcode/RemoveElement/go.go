package main

import "fmt"

func DoQuestion(nums []int, val int) int {
	lastShift := len(nums)
	for i := 0; i < lastShift; i++ {
		if nums[i] == val {
			shiftDownElements(nums, i, val)
			i--
			lastShift--
		}
	}

	return lastShift
}

func shiftDownElements(nums []int, start int, val int) []int {
	for i := start; i < len(nums)-1; i++ {
		nums[i] = nums[i+1]
	}
	nums[len(nums)-1] = val
	return nums
}

func main() {
	fmt.Println("Running main")
	nums := []int{
		2, 2, 3,
	}
	result := DoQuestion(nums, 2)
	fmt.Printf("Got %v\n", result)
	fmt.Printf("Array %v\n", nums)
}
