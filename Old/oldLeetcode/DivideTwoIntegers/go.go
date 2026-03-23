package main

import (
	"fmt"
	"math"
)

const (
	maxInt = 2147483647
	minInt = -2147483648
)

func DoQuestion(dividend int, divisor int) int {
	steps, total := 0, 0
	absDividend, absDivisor := int(math.Abs(float64(dividend))), int(math.Abs(float64(divisor)))
	isDividendNegative, isDivisorNegative := dividend < 0, divisor < 0

	for total <= absDividend-absDivisor {
		total += absDivisor
		steps++
	}

	if isDividendNegative && isDivisorNegative {
		if steps > maxInt {
			return maxInt
		}
		return steps
	} else if isDividendNegative || isDivisorNegative {
		if steps < minInt {
			return minInt
		}
		return -steps
	}
	return steps
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(10, -3)
	fmt.Printf("Got %v\n", result)
}
