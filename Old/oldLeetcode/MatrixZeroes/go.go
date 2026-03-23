package main

import "fmt"

func DoQuestion(matrix [][]int) {
	zeroRows := make([]int, 0, len(matrix))
	zeroColumns := make([]int, 0, len(matrix[0]))
	for x := 0; x < len(matrix); x++ {
		for y := 0; y < len(matrix[x]); y++ {
			if matrix[x][y] == 0 {
				zeroRows = append(zeroRows, x)
				zeroColumns = append(zeroColumns, y)
			}
		}
	}

	for _, v := range zeroRows {
		for i := 0; i < len(matrix[v]); i++ {
			matrix[v][i] = 0
		}
	}

	for _, v := range zeroColumns {
		for i := 0; i < len(matrix); i++ {
			matrix[i][v] = 0
		}
	}
}

func main() {
	fmt.Println("Running main")
	matrix := [][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}
	DoQuestion(matrix)
	fmt.Printf("Got %v\n", matrix)
}
