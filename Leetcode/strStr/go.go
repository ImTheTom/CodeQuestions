package main

import "fmt"

func DoQuestion(haystack string, needle string) int {
	lengthOfHaystack := len(haystack)
	lengthOfNeedle := len(needle)

	// Special case for 0
	if lengthOfHaystack == 0 || lengthOfNeedle == 0 {
		return 0
	}

	for i := 0; i <= lengthOfHaystack-lengthOfNeedle; i++ {
		// Is this i even possible?
		if needle[0] == haystack[i] {
			matches := true
			m := 0
			for j := i; j < lengthOfNeedle+i; j++ {
				if haystack[j] != needle[m] {
					matches = false
					break
				}
				m++
			}
			if matches {
				return i
			}
		}
	}
	return -1
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion("mississippi", "issip")
	fmt.Printf("Got %v\n", result)
}
