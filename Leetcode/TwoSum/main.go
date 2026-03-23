package main

import (
	"fmt"
	"time"
)

/*
Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

You can return the answer in any order.
*/

func DoQuestion(nums []int, target int) []int {
	for iIdx, iNum := range nums {
		for jIdx, jNum := range nums {
			if iIdx == jIdx {
				continue
			}

			if iNum+jNum == target {
				return []int{iIdx, jIdx}
			}
		}
	}

	return []int{}
}

func DoQuestion2(nums []int, target int) []int {
	numMap := make(map[int]int, len(nums))

	for idx, num := range nums {
		numMap[num] = idx
	}

	for idx, num := range nums {
		lookUpVal := target - num

		if lookUpIdx, exists := numMap[lookUpVal]; exists {
			if idx == lookUpIdx {
				continue
			}

			return []int{idx, lookUpIdx}
		}
	}

	return []int{}
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion2([]int{2, 7, 11, 15}, 9)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
