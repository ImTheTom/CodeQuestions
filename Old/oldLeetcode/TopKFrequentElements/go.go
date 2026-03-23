package main

import (
	"fmt"
	"sort"
)

func DoQuestion(nums []int, k int) []int {
	mapOfFrequency := make(map[int]int, len(nums))

	for _, v := range nums {
		mapOfFrequency[v]++
	}

	keys := make([]int, 0, k)

	for key := range mapOfFrequency {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i int, j int) bool {
		return mapOfFrequency[keys[i]] > mapOfFrequency[keys[j]]
	})

	return keys[:k]
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 1, 1, 2, 2, 3,
	}, 2)
	fmt.Printf("Got %v\n", result)
}
