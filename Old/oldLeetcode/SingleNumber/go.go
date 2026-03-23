package main

import "fmt"

func DoQuestion(nums []int) int {
	intMap := make(map[int]int, len(nums))
	for _, v := range nums {
		intMap[v]++
	}
	for k, v := range intMap {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		2, 2, 1,
	})
	fmt.Printf("Got %v\n", result)
}
