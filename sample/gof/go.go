package main

import (
	"fmt"
	"time"
)

func DoQuestion(input []string) int {
	return 0
}

func main() {
	start := time.Now()

	lines, err := readFile("text.txt")
	if err != nil {
		fmt.Println("reading failed " + err.Error())
	}

	fmt.Printf("loaded file with %d lines\n", len(lines))

	fmt.Println("Running main")
	result := DoQuestion(lines)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
