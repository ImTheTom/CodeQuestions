package main

import (
	"fmt"
)

func DoQuestion(x int) int {
	// Special case for 0
	if x == 0 {
		return 0
	}

	for i := 1; i <= x; i++ {
		squaredValue := i * i
		if x == squaredValue {
			return i
		}
		if x < squaredValue {
			return i - 1
		}
	}

	return -1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(8)
	fmt.Printf("Got %v\n", result)
}
