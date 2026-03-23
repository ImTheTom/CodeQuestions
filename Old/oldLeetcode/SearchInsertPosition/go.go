package main

import "fmt"

func DoQuestion(first []int, second int) int {
	start := 0
	end := len(first) - 1

	for start <= end {
		mid := (start + end) / 2

		if first[mid] == second {
			return mid
		} else if first[mid] < second {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}

	return end + 1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		1, 3, 5, 6,
	}, 5)
	fmt.Printf("Got %v\n", result)
}
