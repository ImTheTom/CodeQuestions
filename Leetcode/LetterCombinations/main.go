package main

import (
	"fmt"
	"time"
)

var NumberMapping = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func DoQuestion(digits string) []string {
	options := make([]string, 0)
	for i := 0; i < len(digits); i++ {
		numberMap := NumberMapping[string(digits[i])]

		if len(options) == 0 {
			for _, option := range numberMap {
				options = append(options, string(option))
			}

			continue
		}

		newMappings := make([]string, 0)
		for _, number := range numberMap {
			for _, option := range options {
				newMappings = append(newMappings, option+string(number))
			}
		}

		options = newMappings
	}

	return options
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("23")
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
