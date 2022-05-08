package main

import (
	"fmt"
	"math"
)

func NotLengthSpecified(size int) []int {
	var res []int
	for i := 0; i < size; i++ {
		res = append(res, int(math.Sqrt(float64(i))))
	}
	return res
}

func LengthSpecified(size int) []int {
	var res = make([]int, size)
	for i := 0; i < size; i++ {
		res[i] = int(math.Sqrt(float64(i)))
	}
	return res
}

func main() {
	fmt.Println("Running main")
	result := LengthSpecified(1000)
	fmt.Printf("Got %v\n", result)
}
