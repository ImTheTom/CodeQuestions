package main

import "fmt"

var _uniquePaths = 0

func DoQuestion(n int) int {
	_uniquePaths = 0
	res := make([][]int, 0)
	recursiveSolve([]int{}, res, n, 1)
	recursiveSolve([]int{}, res, n, 2)
	return _uniquePaths
}

func recursiveSolve(steps []int, res [][]int, final, newV int) {
	steps = append(steps, newV)

	stepV := stepsValue(steps)

	if stepV == final {
		_uniquePaths++
		return
	} else if stepV > final {
		return
	}

	recursiveSolve(steps, res, final, 1)
	recursiveSolve(steps, res, final, 2)
}

func stepsValue(numbs []int) int {
	total := 0
	for _, numb := range numbs {
		total += numb
	}
	return total
}

func actualSolve(n int) int {
	if n <= 3 {
		return n
	}

	n1, n2 := 3, 2
	for i := 3; i <= n; i++ {
		tmp := n1
		n1 = n1 + n2
		n2 = tmp
	}

	return n2
}

func main() {
	fmt.Println("Running main")
	result := actualSolve(5)
	fmt.Printf("Got 5 - %v\n", result)
}
