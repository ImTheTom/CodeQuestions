package main

import (
	"fmt"
	"os"
	"strconv"
)

const (
	i1 = 121
	o1 = true
	i2 = -121
	o2 = false
	i3 = 10
	o3 = false
)

func isPalindrome(x int) bool {
	xString := strconv.Itoa(x)
	half := len(xString) / 2
	for i := 0; i <= half; i++ {
		if xString[i] != xString[len(xString)-1-i] {
			return false
		}
	}
	return true
}

func doFunc(input int, output bool) {
	actual := isPalindrome(input)
	fmt.Printf("For %d I got %v I want %v\n", input, actual, output)
	if actual != output {
		os.Exit(1)
	}
}

func main() {
	fmt.Println("Running main")
	doFunc(i1, o1)
	doFunc(i2, o2)
	doFunc(i3, o3)
}
