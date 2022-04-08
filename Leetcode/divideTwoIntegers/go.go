package main

import "fmt"

const (
	maxInt = 2147483647
	minInt = -2147483648
)

func DoQuestion(dividend int, divisor int) int {
	result := dividend / divisor
	if result > maxInt {
		return maxInt
	}
	if result < minInt {
		return minInt
	}
	return result
}

func main() {
	res := DoQuestion(10, 3)
	fmt.Println(res)
}
