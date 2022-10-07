package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()

	fmt.Println("Running main")

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
