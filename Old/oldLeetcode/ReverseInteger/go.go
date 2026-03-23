package main

import (
	"fmt"
	"os"
)

const (
	input  = 123
	output = 321

	input2  = -123
	output2 = -321

	MAXINT32 = 2147483647
	MININT32 = -2147483647
)

func reverse(x int) int {
	var new, remainder int

	for x != 0 {
		remainder = x % 10
		new = new*10 + remainder
		x /= 10
	}

	if new >= MAXINT32 || new <= MININT32 {
		return 0
	}

	return new
}

func main() {
	fmt.Println("Running main")
	actual := reverse(input)
	fmt.Printf("For - %d I got %d wanted %d\n", input, actual, output)
	if actual != output {
		os.Exit(1)
	}

	actual2 := reverse(input2)
	fmt.Printf("For - %d I got %d wanted %d\n", input2, actual2, output2)
	if actual2 != output2 {
		os.Exit(1)
	}
}
