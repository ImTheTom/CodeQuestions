package main

import (
	"fmt"
	"time"
)

func DoQuestion(num int) int {
	for num > 9 {
		splitted := splitInt(num)

		num = sum(splitted)
	}

	return num
}

func splitInt(n int) []int {
	res := []int{}

	for n > 0 {
		res = append(res, n%10)
		n /= 10
	}

	return res
}

func sum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}

	return res
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(38)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
