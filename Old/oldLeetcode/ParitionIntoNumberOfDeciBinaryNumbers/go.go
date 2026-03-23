package main

import (
	"fmt"
	"strconv"
	"time"
)

func DoQuestion(n string) int {
	highest := 0

	for _, v := range n {
		value, err := strconv.Atoi(string(v))

		if err != nil {
			panic(err)
		}

		if value > highest {
			highest = value
		}

		if highest == 9 {
			return 9
		}
	}

	return highest
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("32")
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
