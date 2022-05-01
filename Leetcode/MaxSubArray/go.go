package main

import "fmt"

const defaultValue = 0

func maxSubArray(nums []int) int {
	res := nums[0]
	tmpRes := defaultValue

	for _, n := range nums {
		if tmpRes+n > 0 {
			tmpRes += n
			if tmpRes > res {
				res = tmpRes
			}
		} else {
			tmpRes = defaultValue
			if n > res {
				res = n
			}
		}
	}

	return res
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		-2, 1, -3, 4, -1, 2, 1, -5, 4,
	})
	fmt.Printf("Got %v\n", result)
}
