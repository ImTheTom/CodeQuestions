package main

import "fmt"

var romanMap = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func DoQuestion(input string) int {
	total := 0
	for i := 0; i < len(input); i++ {
		value := romanMap[string(input[i])]
		if i == len(input)-1 {
			total += value
			continue
		}
		nextValue := romanMap[string(input[i+1])]
		if value == 1 && (nextValue == 5 || nextValue == 10) {
			total += nextValue - value
			i++
			continue
		} else if value == 10 && (nextValue == 50 || nextValue == 100) {
			total += nextValue - value
			i++
			continue
		} else if value == 100 && (nextValue == 500 || nextValue == 1000) {
			total += nextValue - value
			i++
			continue
		}
		total += value
	}
	return total
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion("III")
	fmt.Printf("Got %v\n", result)
}
