package main

import "fmt"

func DoQuestion(n int) bool {
	for i := 0; i < 50; i++ {
		nSplit := splitInt(n)
		n = squareArray(nSplit)
		if n == 1 {
			return true
		}
	}
	return false
}

func splitInt(n int) []int {
	res := []int{}

	for n != 0 {
		res = append(res, n%10)
		n = n / 10
	}

	return res
}

func squareArray(nArr []int) int {
	res := 0

	for _, v := range nArr {
		res += v * v
	}

	return res
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(19)
	fmt.Printf("Got %v\n", result)
}
