package main

import (
	"fmt"
	"time"
)

func DoQuestion(nums []int) int {
	totalRequired := 0

	for _, num := range nums {
		if num%3 == 0 {
			continue
		}

		if (num+1)%3 == 0 {
			totalRequired++

			continue
		}

		if (num-1)%3 == 0 {
			totalRequired++

			continue
		}

		if (num+2)%3 == 0 {
			totalRequired += 2

			continue
		}

		if (num-2)%3 == 0 {
			totalRequired += 2

			continue
		}

		totalRequired += 3
	}

	return totalRequired
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 2, 3, 4,
	})
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
