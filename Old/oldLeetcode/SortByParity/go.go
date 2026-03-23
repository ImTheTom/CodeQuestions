package main

import "fmt"

func DoQuestion(nums []int) []int {
	res := make([]int, 0, len(nums))
	for _, v := range nums {
		if v%2 == 0 {
			res = append([]int{v}, res...)
		} else {
			res = append(res, v)
		}
	}
	return res
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		3, 1, 2, 4,
	})
	fmt.Printf("Got %v\n", result)
}
