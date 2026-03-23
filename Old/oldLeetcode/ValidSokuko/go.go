package main

import "fmt"

/*
	0 | 1 | 2
	- | - | -
	3 | 4 | 5
	- | - | -
	6 | 7 | 8
*/

const (
	boardLengthAndHeight int  = 9
	fullStopByte         byte = '.'
)

func DoQuestion(board [][]byte) bool {
	columnsPassed := checkColumns(board)
	rowsPassed := checkRows(board)
	if !rowsPassed || !columnsPassed {
		return false
	}

	for i := 0; i < 9; i++ {
		if !checkQuadrant(board, i) {
			return false
		}
	}

	return true
}

func checkColumns(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var bytesSeen []byte

		for j := 0; j < 9; j++ {
			if board[j][i] != fullStopByte {
				bytesSeen = append(bytesSeen, board[j][i])
			}
		}

		if !containsNoDuplicates(bytesSeen) {
			return false
		}

	}
	return true
}

func checkRows(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		var bytesSeen []byte

		for j := 0; j < 9; j++ {
			if board[i][j] != fullStopByte {
				bytesSeen = append(bytesSeen, board[i][j])
			}
		}

		if !containsNoDuplicates(bytesSeen) {
			return false
		}

	}
	return true
}

func checkQuadrant(board [][]byte, quadrant int) bool {
	rowQuadrant := quadrant / 3
	columnQuadrant := quadrant % 3

	var bytesSeen []byte

	for i := 0; i < 9; i++ {
		column := i%3 + (columnQuadrant * 3)
		row := i/3 + (rowQuadrant * 3)
		if board[row][column] != fullStopByte {
			bytesSeen = append(bytesSeen, board[row][column])
		}
	}

	return containsNoDuplicates(bytesSeen)
}

func containsNoDuplicates(bytesSeen []byte) bool {
	for i, v := range bytesSeen {
		for j := i + 1; j < len(bytesSeen); j++ {
			if v == bytesSeen[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	})
	fmt.Printf("Got %v\n", result)
}
