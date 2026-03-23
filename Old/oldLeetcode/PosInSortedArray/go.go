package main

import "fmt"

func DoQuestion(nums []int, target int) []int {
	var start, end int

	start, end = BinarySearch(nums, target)

	if start != -1 {
		for i := start - 1; i >= 0; i-- {
			if nums[i] == target {
				start--
			}
		}

		for i := end + 1; i < len(nums); i++ {
			if nums[i] == target {
				end++
			}
		}
	}

	return []int{
		start, end,
	}
}

func BinarySearch(nums []int, target int) (int, int) {
	l, r := 0, len(nums)-1

	for l <= r {
		m := (l + r) / 2

		if nums[m] < target {
			l = m + 1
		} else if nums[m] > target {
			r = m - 1
		} else {
			return m, m
		}
	}

	return -1, -1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		5, 7, 7, 8, 8, 10,
	}, 8)
	fmt.Printf("Got %v\n", result)
}
