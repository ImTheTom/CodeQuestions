package main

import (
	"fmt"
	"time"
)

func DoQuestion(edges [][]int) int {
	counts := countValues(edges)

	return findMostCommonValue(counts)
}

func countValues(edges [][]int) map[int]int {
	counts := make(map[int]int)

	for _, edge := range edges {
		for _, value := range edge {
			counts[value]++
		}
	}

	return counts
}

func findMostCommonValue(counts map[int]int) int {
	max := 0
	maxIndex := 0
	for index, count := range counts {
		if count > max {
			max = count
			maxIndex = index
		}
	}

	return maxIndex
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([][]int{
		{1, 2}, {2, 3}, {4, 2},
	})
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
