package main

import "fmt"

// Bubble sort
func DoQuestion(nums []int) {
	n := len(nums)
	swapped := true
	for swapped {
		swapped = false
		for i := 1; i < n; i++ {
			if nums[i-1] > nums[i] {
				tmp := nums[i-1]
				nums[i-1] = nums[i]
				nums[i] = tmp
				swapped = true
			}
		}
	}
}

func main() {
	fmt.Println("Running main")
	matrix := []int{2, 0, 2, 1, 1, 0}
	DoQuestion(matrix)
	fmt.Printf("Got %v\n", matrix)
}
