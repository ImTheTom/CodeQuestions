package main

import "fmt"

const (
	in  = "dvdf"
	out = 3
)

func lengthOfLongestSubstring(s string) int {
	longest := 0
	current := 0
	longestString := ""

	currentlySeen := make(map[rune]int)
	for i := 0; i < len(s); i++ {
		v := rune(s[i])

		if currentlySeen[v] == 1 {
			currentlySeen = make(map[rune]int)
			if current > longest {
				longest = current
			}

			i = i - len(longestString)
			longestString = ""
			current = 0
		} else {
			currentlySeen[v] = 1
			current++
			longestString += string(v)
		}
	}

	if current > longest {
		longest = current
	}

	return longest
}

func main() {
	fmt.Println("Running main")
	length := lengthOfLongestSubstring(in)
	fmt.Printf("Expected %d Got %d\n", out, length)
}
