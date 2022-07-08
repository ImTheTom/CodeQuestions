package main

import (
	"fmt"
	"time"
)

/*
Arrays
[N][M]Type
[N][M]Type{ {value1, value2 }, ..., { valueX, valueN} }
[...][...]Type{ {value1, value2 }, ..., { valueX, valueN} }
*/

/*
Slices
make([][]Type, length, capacity)
make([]Type, length)
[]Type{}
[]Type{value1, value2, ..., valueN}
*/

// Arrays are fixed size.

// Working with arrays is very hard in go. Usually use slices instead.
func slcTraverse(arr [][]int) {
	fmt.Println("Printing Slice")
	for i, row := range arr {
		fmt.Printf("Printing row %d\n", i)
		for _, v := range row {
			fmt.Println(v)
		}
	}
	fmt.Println("Done")
}

func slcSearch(arr [][]int, needle int) (int, int) {
	for x, row := range arr {
		for y, v := range row {
			if v == needle {
				return x, y
			}
		}
	}
	return -1, -1
}

func slcUpdate(arr [][]int, locX, locY, newValue int) {
	arr[locX][locY] = newValue
}

func main() {
	start := time.Now()

	matExample := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
	}

	fmt.Println("Running main")

	slcTraverse(matExample)

	fmt.Println(slcSearch(matExample, 7))
	fmt.Println(slcSearch(matExample, 11))

	slcUpdate(matExample, 1, 3, 11)
	fmt.Println(slcSearch(matExample, 11))

	slcTraverse(matExample)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
