package main

import "fmt"

const (
	in  = "abcabcbb"
	out = 3 // abc
)

func lengthOfLongestSubstring(s string) int {
	return 0
}

func main() {
	fmt.Println("Running main")
	length := lengthOfLongestSubstring(in)
	fmt.Printf("Expected %d Got %d\n", out, length)
}
