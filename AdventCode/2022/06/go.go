package main

import (
	"fmt"
	"time"
)

// const file = "text.txt"

const file = "actual.txt"

func DoQuestion(input []string) int {
	for i := 0; i < len(input[0]); i++ {
		current := make(map[byte]int, 4)
		for j := i; j < i+4; j++ {
			current[input[0][j]]++
		}

		isAllOne := true
		for _, v := range current {
			if v != 1 {
				isAllOne = false
			}
		}

		if isAllOne {
			return i + 4
		}
	}
	return -1
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
