package main

import (
	"fmt"
	"os"
)

const (
	s       = "ABCD"
	numRows = 2
	output  = "ACBD"
)

func convert(s string, numRows int) string {
	matrix := make([]string, numRows*len(s))

	isZigging := true // Going down vertically
	currentColumn := 0

	currentZig := 0
	numberOfZags := numRows - 2
	currentZag := 0
	if numberOfZags < 0 {
		numberOfZags = 0
	}

	for _, v := range s {
		vString := string(v)
		// Going down
		if isZigging {
			index := currentColumn + (currentZig * len(s))
			matrix[index] = vString

			currentZig++

			if currentZig >= numRows {
				isZigging = false
				currentZig = 0
				currentColumn++

				if currentZag >= numberOfZags {
					isZigging = true
					currentZag = 0
				}
			}
		} else {
			index := currentColumn + ((numberOfZags - currentZag) * len(s))
			matrix[index] = vString

			currentZag++
			currentColumn++

			if currentZag >= numberOfZags {
				isZigging = true
				currentZag = 0
			}
		}
	}
	return flattenArray(matrix)
}

func flattenArray(arr []string) string {
	var current string
	for _, v := range arr {
		current += v
	}
	return current
}

func main() {
	fmt.Println("Running main")
	actual := convert(s, numRows)
	fmt.Printf("For %s I got %s when I wanted %s\n", s, actual, output)
	if actual != output {
		os.Exit(1)
	}
}
