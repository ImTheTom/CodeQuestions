package main

import (
	"fmt"
	"time"
)

const (
	OPENING_TYPE_ONE   = "("
	CLOSING_TYPE_ONE   = ")"
	OPENING_TYPE_TWO   = "["
	CLOSING_TYPE_TWO   = "]"
	OPENING_TYPE_THREE = "{"
	CLOSING_TYPE_THREE = "}"
)

var closingToOpening = map[string]string{
	CLOSING_TYPE_ONE:   OPENING_TYPE_ONE,
	CLOSING_TYPE_TWO:   OPENING_TYPE_TWO,
	CLOSING_TYPE_THREE: OPENING_TYPE_THREE,
}

var isOpening = map[string]interface{}{
	OPENING_TYPE_ONE:   struct{}{},
	OPENING_TYPE_TWO:   struct{}{},
	OPENING_TYPE_THREE: struct{}{},
}

func DoQuestion(str string) bool {
	openingStack := make([]string, 0, len(str))

	for _, s := range str {
		if _, exists := isOpening[string(s)]; exists {
			openingStack = append(openingStack, string(s))

			continue
		}

		if len(openingStack) == 0 {
			return false
		}

		lastElement := len(openingStack) - 1

		lastOpening := openingStack[lastElement]

		expectedOpening := closingToOpening[string(s)]

		if lastOpening != expectedOpening {
			return false
		}

		openingStack = openingStack[:lastElement]
	}

	if len(openingStack) != 0 {
		return false
	}

	return true
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("()[]{}")
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
