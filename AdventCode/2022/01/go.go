package main

import (
	"fmt"
	"strconv"
	"time"
)

func DoQuestion(input []string) int {
	max := 0
	tmp := 0
	for _, value := range input {
		// Is this an empty line to signal a new elf
		if value == "" {
			if tmp > max {
				max = tmp
			}
			tmp = 0

			continue
		}

		convertedValue, err := strconv.Atoi(value)
		if err != nil {
			panic(err)
		}

		tmp += convertedValue
	}
	return max
}

func main() {
	start := time.Now()

	lines, err := readFile("actual.txt")
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
