package main

import "fmt"

// Cheated looked at answers LOL!

func DoQuestion(first int) []string {
	var result []string
	GenerateCombos(first, first, &result, "")
	return result
}

func GenerateCombos(left, right int, result *[]string, cur string) {
	if left == 0 && right == 0 {
		*result = append(*result, cur)
		return
	}

	// Don't add invalid parenthesis
	if right < left {
		return
	}

	if left > 0 {
		GenerateCombos(left-1, right, result, cur+"(")
	}

	if right > 0 {
		GenerateCombos(left, right-1, result, cur+")")
	}
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(3)
	fmt.Printf("Got %v\n", result)
}
