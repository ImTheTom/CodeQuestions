package main

import (
	"fmt"
)

// Cheated need to get my head around bits again
func DoQuestion(num uint32) int {
	count := uint32(0)

	for ; num > uint32(0); num >>= 1 {
		count += num & 1
	}

	return int(count)
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(00000000000000000000000000001011)
	fmt.Printf("Got %v\n", result)
}
