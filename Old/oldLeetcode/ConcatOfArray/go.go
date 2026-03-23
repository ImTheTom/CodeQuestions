package main

import (
	"fmt"
	"time"
)

func DoQuestion(nums []int) []int {
	res := make([]int, 0, len(nums)*2)
	for i := 0; i < 2; i++ {
		res = append(res, nums...)
	}
	return res
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 2, 1,
	})
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
