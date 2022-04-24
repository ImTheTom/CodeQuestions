package main

import "fmt"

func DoQuestion(first int) [][]int {
	res := [][]int{}
	for i := 1; i <= first; i++ {
		currList := make([]int, i)
		currList[0] = 1
		currList[i-1] = 1
		for j := 1; j < len(currList)-1; j++ {
			idxV := res[i-2][j-1] + res[i-2][j]
			currList[j] = idxV
		}
		res = append(res, currList)
	}
	return res
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(5)
	fmt.Printf("Got %v\n", result)
}
