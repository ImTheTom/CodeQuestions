package main

import (
	"fmt"
)

const (
	GO_RIGHT = iota
	GO_DOWN
	GO_LEFT
	GO_UP
)

func DoQuestion(matrix [][]int) []int {
	result := []int{}

	maxHeight, maxLength := len(matrix), len(matrix[0])
	smallestHeight, smallestLength := 0, 0
	currentX, currentY := 0, 0
	currentMove := 0

	amountAdded := -1

	for amountAdded != 0 {
		switch currentMove {
		case GO_RIGHT:
			moveResult := MoveRight(matrix, currentX, maxLength, currentY)
			amountAdded = len(moveResult)
			result = append(result, moveResult...)
			currentX = maxLength - 1
			currentY++
			smallestHeight++
		case GO_DOWN:
			moveResult := MoveDown(matrix, currentY, maxHeight, currentX)
			amountAdded = len(moveResult)
			result = append(result, moveResult...)
			currentY = maxHeight - 1
			currentX--
			maxLength--
		case GO_LEFT:
			moveResult := MoveLeft(matrix, currentX, smallestLength, currentY)
			amountAdded = len(moveResult)
			result = append(result, moveResult...)
			currentX = smallestLength
			currentY--
			maxHeight--
		case GO_UP:
			moveResult := MoveUp(matrix, currentY, smallestHeight, currentX)
			amountAdded = len(moveResult)
			result = append(result, moveResult...)
			currentY = smallestHeight
			currentX++
			smallestLength++
		}

		currentMove++
		if currentMove > 3 {
			currentMove = 0
		}
	}

	return result
}

func MoveRight(matrix [][]int, startX, finishX, y int) []int {
	result := []int{}

	for i := startX; i < finishX; i++ {
		result = append(result, matrix[y][i])
	}

	return result
}

func MoveDown(matrix [][]int, startY, finishY, x int) []int {
	result := []int{}

	for i := startY; i < finishY; i++ {
		result = append(result, matrix[i][x])
	}

	return result
}

func MoveLeft(matrix [][]int, startX, finishX, y int) []int {
	result := []int{}

	for i := startX; i >= finishX; i-- {
		result = append(result, matrix[y][i])
	}

	return result
}

func MoveUp(matrix [][]int, startY, finishY, x int) []int {
	result := []int{}

	for i := startY; i >= finishY; i-- {
		result = append(result, matrix[i][x])
	}

	return result
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([][]int{
		{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12},
	})
	fmt.Printf("Got %v\n", result)
}
