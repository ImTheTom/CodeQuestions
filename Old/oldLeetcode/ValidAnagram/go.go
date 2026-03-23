package main

import "fmt"

func DoQuestion(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	sMap := make(map[rune]int, len(s))
	tMap := make(map[rune]int, len(s))

	for _, v := range s {
		sMap[v]++
	}

	for _, v := range t {
		tMap[v]++
	}

	for k, v := range sMap {
		if tMap[k] != v {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion("anagram", "nagaram")
	fmt.Printf("Got %v\n", result)
}
