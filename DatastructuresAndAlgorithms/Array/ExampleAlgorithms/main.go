package main

import (
	"fmt"
)

const (
	ArrayBeforePrefix = "Array before "
	ArrayAfterPrefix  = "Array after "
)

/*
Arrays are declared such as - var arr [5]int

But usually in Go we use slices - arr := make([]int, 0, 5) or arr := make([]int, 5)

This will allow []int to be passed to functions instead of specifying [5]int

All algorithms will use slices instead of arrays.
*/

func main() {
	fmt.Println("Running main")

	ShowcaseRotate()

	fmt.Println("Total execution over")
}

// printArray is a helper function to print out arrays
func printArray(nums []int, prefix string) {
	fmt.Printf("%s%v\n", prefix, nums)
}
