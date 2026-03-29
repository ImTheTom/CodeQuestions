package main

import (
	"fmt"
	"time"
)

const (
	GOING_DOWN = iota + 1
	GOING_UP
)

func DoQuestion(str string, numRows int) string {
	if numRows == 1 {
		return str
	}

	mapping := make([][]string, numRows)
	for i := 0; i < len(mapping); i++ {
		mapping[i] = make([]string, len(str))
	}

	currentX, currentY := 0, 0

	direction := GOING_DOWN

	for _, s := range str {
		mapping[currentY][currentX] = string(s)

		switch direction {
		case GOING_DOWN:
			if currentY == len(mapping)-1 {
				currentY -= 1
				currentX++
				direction = GOING_UP
				break
			}

			currentY++

		case GOING_UP:
			if currentY == 0 {
				currentY += 1
				direction = GOING_DOWN
				break
			}

			currentX++
			currentY--
		}

	}

	result := ""

	for _, q := range mapping {
		for _, w := range q {
			if w != "" {
				result += w
			}
		}
	}

	return result
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion("PAYPALISHIRING", 3)
	fmt.Printf("Got %v - Want PAHNAPLSIIGYIR\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
