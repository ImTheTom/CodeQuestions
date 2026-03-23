package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	len := len(nums)

	for i := 0; i < len-1; i++ {
		current := nums[i]

		for j := i + 1; j < len; j++ {
			otherValue := nums[j]
			if current+otherValue == target {
				result = append(result, i)
				result = append(result, j)
				return result
			}
		}
	}

	return result
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	resultString := fmt.Sprintf("%v should be [0, 1]", result)
	fmt.Println(resultString)
}
