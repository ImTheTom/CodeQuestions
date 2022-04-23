package main

import "fmt"

func DoQuestion(first int, second int) int {
	return first + second
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(1, 1)
	fmt.Printf("Got %v\n", result)
}
