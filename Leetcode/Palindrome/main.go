package main

import (
	"fmt"
	"strconv"
	"time"
)

func DoQuestion(first int) bool {
	if first < 0 {
		return false
	}

	firstStr := strconv.Itoa(first)
	firstStrLen := len(firstStr)

	for i := 0; i < firstStrLen/2; i++ {
		if firstStr[i] != firstStr[firstStrLen-1-i] {
			return false
		}
	}

	return true
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(121)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
