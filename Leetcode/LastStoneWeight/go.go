package main

import (
	"fmt"
	"sort"
)

func DoQuestion(stones []int) int {
	// Special case for single index
	if len(stones) == 1 {
		return stones[0]
	}

	for {
		// Sort array so we know the two heaviest stones are at the end
		sort.Ints(stones)
		heaviest := stones[len(stones)-1]
		secondHeaviest := stones[len(stones)-2]
		// If the second heaviest is zero, the single one left is the heaviest
		if secondHeaviest == 0 {
			return heaviest
		}
		// Smash and update values
		res := heaviest - secondHeaviest
		stones[len(stones)-2] = 0
		stones[len(stones)-1] = res
	}
}

func main() {
	fmt.Println("Running main")
	stones := []int{
		2, 7, 4, 1, 8, 1,
	}
	result := DoQuestion(stones)
	fmt.Printf("Got %v\n", result)
}
