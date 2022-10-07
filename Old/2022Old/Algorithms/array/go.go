package main

import (
	"fmt"
	"time"
)

func PrintArray(arr []int) {
	fmt.Println("PRINTING ARRAY")
	for _, v := range arr {
		fmt.Println(v)
	}
	fmt.Println("DONE")
}

func doRotate() {
	rotateArr := []int{
		1, 2, 3, 4, 5, 6, 7,
	}

	Rotate(rotateArr, 2)

	PrintArray(rotateArr)
}

func main() {
	start := time.Now()

	fmt.Println("Running main")

	doRotate()

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
