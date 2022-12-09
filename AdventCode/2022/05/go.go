package main

import (
	"fmt"
	"time"
)

const file = "text.txt"

// const file = "actual.txt"

func DoQuestion(input []string) int {
	return 0
}

func main() {
	start := time.Now()

	lines, err := readFile(file)
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
