package main

import (
	"fmt"
	"math"
)

func DoQuestion(nums []int, target int) int {
	closest := 99999999999
	value := 9999999999999

	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				tmpValue := nums[i] + nums[j] + nums[k]
				diff := int(math.Abs(float64(tmpValue - target)))
				if diff < closest {
					closest = diff
					value = tmpValue
				}
			}
		}
	}

	return value
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		-1, 2, 1, -4,
	}, 1)
	fmt.Printf("Got %v\n", result)
}
