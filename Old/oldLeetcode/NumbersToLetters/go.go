package main

import "fmt"

var lookupTable = map[int][]string{
	0: nil,
	1: nil,
	2: {"a", "b", "c"},
	3: {"d", "e", "f"},
	4: {"g", "h", "i"},
	5: {"j", "k", "l"},
	6: {"m", "n", "o"},
	7: {"p", "q", "r", "s"},
	8: {"t", "u", "v"},
	9: {"w", "x", "y", "z"},
}

func DoQuestion(digits string) []string {
	result := make([]string, 0)
	for _, v := range digits {
		numVal := int(v - '0')
		if len(result) == 0 {
			result = append(result, lookupTable[numVal]...)
		} else {
			var tempNewResult []string
			newLetters := lookupTable[numVal]
			for _, i := range result {
				for _, k := range newLetters {
					tempNewResult = append(tempNewResult, fmt.Sprintf("%s%s", i, k))
				}
			}
			result = tempNewResult
		}
	}
	return result
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion("23")
	fmt.Printf("Got %v\n", result)
}
