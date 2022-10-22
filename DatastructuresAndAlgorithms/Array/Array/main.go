package main

import (
	"fmt"
)

/*
Arrays are declared such as - var arr [5]int

But usually in Go we use slices - arr := make([]int, 0, 5) or arr := make([]int, 5)

This will allow []int to be passed to functions instead of specifying [5]int

*/

func main() {
	fmt.Println("Running main")

	ShowcaseArrays()

	ShowcaseSlices()

	fmt.Println("Total execution over")
}
