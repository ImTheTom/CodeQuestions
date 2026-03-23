package main

import "fmt"

func DoQuestion(heights []int) int {
	maxV := 0
	for i := 0; i < len(heights)-1; i++ {
		usingRight := true
		for j := len(heights) - 1; j >= i; j-- {
			maxVHeight := heights[j]
			if maxVHeight > heights[i] {
				maxVHeight = heights[i]
				usingRight = false
			}
			l := j - i
			v := maxVHeight * l
			if v > maxV {
				maxV = v
			}
			if !usingRight {
				break
			}
		}
	}
	return maxV
}

func main() {
	fmt.Println("Running main")
}
