package main

import "fmt"

func DoQuestion(nums []int) {
	lastPos := len(nums) - 1
	for i := 0; i <= lastPos; i++ {
		if nums[i] == 0 {
			shiftDown(i, lastPos, nums)
			lastPos -= 1
			i -= 1
		}
	}
}

func shiftDown(index, lastPos int, nums []int) {
	for i := index; i < lastPos; i++ {
		nums[i] = nums[i+1]
	}
	nums[lastPos] = 0
}

func main() {
	fmt.Println("Running main")
	nums := []int{
		0, 1, 0, 3, 12,
	}
	DoQuestion(nums)
	fmt.Printf("Got %v\n", nums)
}
