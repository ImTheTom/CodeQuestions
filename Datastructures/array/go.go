package main

import (
	"fmt"
	"time"
)

/*
Arrays
[N]Type
[N]Type{value1, value2, ..., valueN}
[...]Type{value1, value2, ..., valueN}
*/

/*
Slices
make([]Type, length, capacity)
make([]Type, length)
[]Type{}
[]Type{value1, value2, ..., valueN}
*/

// Arrays are fixed size.

// Working with arrays is very hard in go. Usually use slices instead.

// Array functions

func arrTraverse(arr [5]int) {
	fmt.Println("Printing array")
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("Done")
}

func arrSearch(arr [5]int, needle int) int {
	for i, v := range arr {
		if v == needle {
			return i
		}
	}
	return -1
}

func arrUpdate(arr [5]int, loc, newValue int) [5]int {
	arr[loc] = newValue
	return arr
}

// Slice functions

func slcTraverse(arr []int) {
	fmt.Println("Printing Slice")
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("Done")
}

func slcSearch(arr []int, needle int) int {
	for i, v := range arr {
		if v == needle {
			return i
		}
	}
	return -1
}

func slcUpdate(arr []int, loc, newValue int) {
	arr[loc] = newValue
}

func runArray() {
	fmt.Println("Running Array funcs")

	arrExample := [5]int{
		1, 2, 3, 4, 5,
	}

	arrTraverse(arrExample)

	res := arrSearch(arrExample, 3)
	fmt.Printf("Location of 3 is at %d\n", res)

	res = arrSearch(arrExample, 6)
	fmt.Printf("Location of 6 is at %d\n", res)

	arrExample = arrUpdate(arrExample, 4, 6)

	arrTraverse(arrExample)
}

func runSlice() {
	fmt.Println("Running Slice funcs")

	slcExample := []int{
		1, 2, 3, 4, 5,
	}

	slcTraverse(slcExample)

	res := slcSearch(slcExample, 3)
	fmt.Printf("Location of 3 is at %d\n", res)

	res = slcSearch(slcExample, 6)
	fmt.Printf("Location of 6 is at %d\n", res)

	slcUpdate(slcExample, 4, 6)

	slcTraverse(slcExample)
}

func main() {
	start := time.Now()

	fmt.Println("Running main")

	runArray()
	fmt.Println()
	runSlice()

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
