package main

import "fmt"

func DoQuestion(s string) int {
	mapOfCounts := make(map[rune]int)
	for _, v := range s {
		mapOfCounts[v]++
	}
	for k, v := range s {
		if mapOfCounts[v] == 1 {
			return k
		}
	}
	return -1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion("loveleetcode")
	fmt.Printf("Got %v\n", result)
}
