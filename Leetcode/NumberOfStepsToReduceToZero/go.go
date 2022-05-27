package main

import "fmt"

func DoQuestion(num int) int {
	steps := 0

	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		steps++
	}

	return steps
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(14)
	fmt.Printf("Got %v\n", result)
}
