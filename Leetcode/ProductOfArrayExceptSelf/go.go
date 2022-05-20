package main

import "fmt"

func DoQuestion(nums []int) []int {
	res := make([]int, len(nums))

	leftProduct := 1
	for i := 0; i < len(nums); i++ {
		leftProduct *= nums[i]
		res[i] = leftProduct
	}

	rightProduct := 1
	for i := len(nums) - 1; i >= 1; i-- {
		res[i] = res[i-1] * rightProduct
		rightProduct *= nums[i]
	}
	res[0] = rightProduct

	return res
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 2, 3, 4,
	})
	fmt.Printf("Got %v\n", result)
}
