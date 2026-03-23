package main

import "fmt"

func DoQuestion(nums []int) int {
	majorityElementMap := make(map[int]int, len(nums))
	for _, v := range nums {
		majorityElementMap[v]++
	}
	majoritySize := len(nums) / 2
	for k, v := range majorityElementMap {
		if v > majoritySize {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		2, 2, 1, 1, 1, 2, 2,
	})
	fmt.Printf("Got %v\n", result)
}
