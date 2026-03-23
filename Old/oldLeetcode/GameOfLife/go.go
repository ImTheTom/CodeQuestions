package main

import "fmt"

const (
	DEAD = iota
	ALIVE
	WILL_DIE
	WILL_ALIVE
)

func DoQuestion(board [][]int) {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			n := countNeighbours(x, y, board)
			value := board[x][y]

			if value == ALIVE && n < 2 {
				value = WILL_DIE // Under population
			} else if value == ALIVE && n > 3 {
				value = WILL_DIE // Over population
			} else if value == DEAD && n == 3 {
				value = WILL_ALIVE
			}

			board[x][y] = value
		}
	}

	CleanBoard(board)
}

func countNeighbours(x, y int, board [][]int) int {
	maxX := len(board) - 1
	maxY := len(board[0]) - 1
	neighbours := 0
	// Left
	if y != 0 {
		if board[x][y-1] == ALIVE || board[x][y-1] == WILL_DIE {
			neighbours++
		}
	}
	// Top Left
	if x != 0 && y != 0 {
		if board[x-1][y-1] == ALIVE || board[x-1][y-1] == WILL_DIE {
			neighbours++
		}
	}
	// Top
	if x != 0 {
		if board[x-1][y] == ALIVE || board[x-1][y] == WILL_DIE {
			neighbours++
		}
	}
	// Top Right
	if x != 0 && y < maxY {
		if board[x-1][y+1] == ALIVE || board[x-1][y+1] == WILL_DIE {
			neighbours++
		}
	}
	// Right
	if y < maxY {
		if board[x][y+1] == ALIVE || board[x][y+1] == WILL_DIE {
			neighbours++
		}
	}
	// Bottom Right
	if x < maxX && y < maxY {
		if board[x+1][y+1] == ALIVE || board[x+1][y+1] == WILL_DIE {
			neighbours++
		}
	}
	// Bottom
	if x < maxX {
		if board[x+1][y] == ALIVE || board[x+1][y] == WILL_DIE {
			neighbours++
		}
	}
	// Bottom Left
	if x < maxX && y != 0 {
		if board[x+1][y-1] == ALIVE || board[x+1][y-1] == WILL_DIE {
			neighbours++
		}
	}
	return neighbours
}

func CleanBoard(board [][]int) {
	for x := 0; x < len(board); x++ {
		for y := 0; y < len(board[x]); y++ {
			if board[x][y] == WILL_DIE {
				board[x][y] = DEAD
			} else if board[x][y] == WILL_ALIVE {
				board[x][y] = ALIVE
			}
		}
	}
}

func main() {
	fmt.Println("Running main")
	board := [][]int{
		{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0},
	}
	DoQuestion(board)
	fmt.Printf("Got %v\n", board)
}
