package main

import "fmt"

func DoQuestion(matrix [][]int) {
	maxRightPos := len(matrix) - 1
	for i := 0; i < len(matrix)/2; i++ {
		for j := i; j < maxRightPos-i; j++ {
			tmp := matrix[i][j]
			matrix[i][j] = matrix[maxRightPos-j][i]
			matrix[maxRightPos-j][i] = matrix[maxRightPos-i][maxRightPos-j]
			matrix[maxRightPos-i][maxRightPos-j] = matrix[j][maxRightPos-i]
			matrix[j][maxRightPos-i] = tmp
		}
	}
}

func main() {
	fmt.Println("Running main")
	matrix := [][]int{
		{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
	}
	DoQuestion(matrix)
	fmt.Printf("Got %v\n", matrix)
}
