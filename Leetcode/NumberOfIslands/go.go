package main

import "fmt"

func DoQuestion(grid [][]byte) int {
	return -1
}

func main() {
	fmt.Println("Running main")
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	result := DoQuestion(grid)
	fmt.Printf("Got %v\n", result)
}
