package main

import "fmt"

func DoQuestion(nums []int) []int {
	res := make([]int, len(nums))
	currentSum := 0

	for i, v := range nums {
		currentSum += v
		res[i] = currentSum
	}

	return res
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 2, 3, 4,
	})
	fmt.Printf("Got %v\n", result)
}
