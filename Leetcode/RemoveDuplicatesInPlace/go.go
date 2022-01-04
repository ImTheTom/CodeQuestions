package main

import "fmt"

func shuffleDownArray(first []int, i int) []int {
	for j := i; j < len(first)-1; j++ {
		first[j] = first[j+1]
	}
	return first
}

func DoQuestion(first []int) int {
	totalArray := len(first)
	if totalArray == 0 {
		return 0
	}

	for i := 0; i < totalArray-1; i++ {
		if first[i] == first[i+1] {
			totalArray--
			first = shuffleDownArray(first, i)
			i--
		}
	}

	return totalArray
}

func main() {
	fmt.Println("Running main")
}
