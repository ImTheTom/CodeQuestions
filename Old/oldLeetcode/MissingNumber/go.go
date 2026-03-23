package main

import (
	"fmt"
	"sort"
)

func DoQuestion(nums []int) int {
	sort.Ints(nums)
	for i, v := range nums {
		if i != v {
			return i
		}
	}
	return len(nums)
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		3, 0, 1,
	})
	fmt.Printf("Got %v\n", result)
}
