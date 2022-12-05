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

	lines, err := readFile("text.txt")
	if err != nil {
		fmt.Println("reading failed " + err.Error())
	}

	fmt.Printf("loaded file with %d lines\n", len(lines))

	fmt.Println("Running main")
	result := DoQuestion(1, 1)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
