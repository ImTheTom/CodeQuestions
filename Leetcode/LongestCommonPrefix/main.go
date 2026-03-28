package main

import (
	"fmt"
	"time"
)

func DoQuestion(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}

	firstString := strs[0]
	secondString := strs[1]

	initialPrefix := make([]byte, 0, len(firstString))

	for i := 0; i < len(firstString); i++ {
		if i >= len(secondString) {
			break
		}

		if firstString[i] == secondString[i] {
			initialPrefix = append(initialPrefix, firstString[i])
		} else {
			break
		}
	}

	returnedPrefix := initialPrefix

	for i := 2; i < len(strs); i++ {
		tempPrefix := make([]byte, 0, len(returnedPrefix))
		for j := 0; j < len(returnedPrefix); j++ {
			if len(strs[i]) == 0 {
				return ""
			}

			if j >= len(strs[i]) {
				break
			}

			if strs[i][j] == returnedPrefix[j] {
				tempPrefix = append(tempPrefix, strs[i][j])
			} else {
				break
			}
		}
		returnedPrefix = tempPrefix
	}

	return string(returnedPrefix)
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion([]string{"flower", "flow", "flight"})
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
