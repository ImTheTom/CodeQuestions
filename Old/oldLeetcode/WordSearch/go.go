package main

import "fmt"

func DoQuestion(board [][]byte, word string) bool {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[0]); y++ {
			if checkForWordAt(x, y, board, word) {
				return true
			}
		}
	}
	return false
}

func checkForWordAt(x, y int, board [][]byte, word string) bool {
	if len(word) == 0 {
		return true
	}

	if x < 0 || x >= len(board) || y < 0 || y >= len(board[0]) {
		return false
	}

	if word[0] != board[x][y] {
		return false
	}

	tmpPos := board[x][y]

	board[x][y] = byte('!')

	if checkForWordAt(x+1, y, board, word[1:]) {
		return true
	}

	if checkForWordAt(x, y+1, board, word[1:]) {
		return true
	}

	if checkForWordAt(x-1, y, board, word[1:]) {
		return true
	}

	if checkForWordAt(x, y-1, board, word[1:]) {
		return true
	}

	board[x][y] = tmpPos

	return false
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([][]byte{
		{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'},
	}, "ABCCED")
	fmt.Printf("Got %v\n", result)
}
