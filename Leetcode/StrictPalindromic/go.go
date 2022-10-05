package main

import (
	"fmt"
	"strconv"
	"time"
)

func DoQuestion(n int) bool {
	parsedN := int64(n)
	for i := 2; i <= n-2; i++ {
		convertedToBase := strconv.FormatInt(parsedN, i)
		if !IsAPalindrome(convertedToBase) {
			return false
		}
	}
	return true
}

func IsAPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(9)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
