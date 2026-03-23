package main

import (
	"fmt"
	"time"
)

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
	queue := make([][]int, 0)
	queue = append(queue, []int{x, y})

	for len(queue) > 0 {
		currElement := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		sourroundingLand := landAt(currElement[0], currElement[1], grid)

		for _, v := range sourroundingLand {
			queue = append(queue, []int{
				v[0], v[1],
			})
		}

		grid[currElement[0]][currElement[1]] = discoveredLand
	}

	return grid
}

func landAt(x, y int, grid [][]byte) [][]int {
	lands := make([][]int, 0, 4)

	if x > 0 && grid[x-1][y] == land {
		lands = append(lands, []int{x - 1, y})
	}

	if y > 0 && grid[x][y-1] == land {
		lands = append(lands, []int{x, y - 1})
	}

	if x < len(grid)-1 && grid[x+1][y] == land {
		lands = append(lands, []int{x + 1, y})
	}

	if y < len(grid[x])-1 && grid[x][y+1] == land {
		lands = append(lands, []int{x, y + 1})
	}

	return lands
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'},
	}
	result := DoQuestion(grid)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Completed in %v\n", elapsed)
}
