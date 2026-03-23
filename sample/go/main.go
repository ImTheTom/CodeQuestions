package main

import (
	"fmt"
	"time"
)

func DoQuestion(first int, second int) int {
	return first + second
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(1, 1)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
