package main

import (
	"fmt"
	"strings"
	"time"
)

func DoQuestion(address string) string {
	return strings.Replace(address, ".", "[.]", -1)
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("1.1.1.1")
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
