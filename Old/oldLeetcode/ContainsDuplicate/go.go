package main

import "fmt"

func DoQuestion(nums []int) bool {
	numMap := make(map[int]int, len(nums))
	for _, v := range nums {
		numMap[v]++
	}
	for _, v := range numMap {
		if v > 1 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 2, 3, 1,
	})
	fmt.Printf("Got %v\n", result)
}
