package main

import "fmt"

func DoQuestion(nums []int) int {
	infoMap := make(map[int]int, len(nums)/2)
	for _, v := range nums {
		infoMap[v]++
	}

	for k, v := range infoMap {
		if v >= 2 {
			return k
		}
	}

	return -1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 3, 4, 2, 2,
	})
	fmt.Printf("Got %v\n", result)
}
