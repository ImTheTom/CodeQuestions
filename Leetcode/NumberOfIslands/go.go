package main

import "fmt"

var (
	land           byte = '1'
	water          byte = '0'
	discoveredLand byte = '2'
)

func DoQuestion(grid [][]byte) int {
	islands := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == land {
				grid = exploreAt(i, j, grid)
				islands++
			}
		}
	}
	return islands
}

func exploreAt(x, y int, grid [][]byte) [][]byte {
	explored := make([][]int, 0, len(grid)*len(grid[0]))
	queue := make([][]int, 0, len(grid)*len(grid[0]))
	queue = append(queue, []int{x, y})

	for len(queue) > 0 {
		currElement := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		sourroundingLand := landAt(currElement[0], currElement[1], grid)

		for _, v := range sourroundingLand {
			if !doesItExistsInArray(v[0], v[1], queue) && !doesItExistsInArray(v[0], v[1], explored) {
				queue = append(queue, []int{
					v[0], v[1],
				})
			}
		}
		explored = append(explored, currElement)
	}

	for _, v := range explored {
		grid[v[0]][v[1]] = discoveredLand
	}

	return grid
}

func landAt(x, y int, grid [][]byte) [][]int {
	possibleLands := make([][]int, 0, 4)
	lands := make([][]int, 0, 4)

	if x > 0 {
		possibleLands = append(possibleLands, []int{x - 1, y})
	}

	if y > 0 {
		possibleLands = append(possibleLands, []int{x, y - 1})
	}

	if x < len(grid)-1 {
		possibleLands = append(possibleLands, []int{x + 1, y})
	}

	if y < len(grid[x])-1 {
		possibleLands = append(possibleLands, []int{x, y + 1})
	}

	for _, v := range possibleLands {
		if grid[v[0]][v[1]] == land {
			lands = append(lands, v)
		}
	}

	return lands
}

func doesItExistsInArray(x, y int, compare [][]int) bool {
	for _, v := range compare {
		if v[0] == x && v[1] == y {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Running main")
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	result := DoQuestion(grid)
	fmt.Printf("Got %v\n", result)
}
