package main

import "fmt"

// DoQuestion way fails with large arrayss
func DoQuestion(digits []int) []int {
	value := sliceToInt(digits)
	return intToSlice(value + 1)
}

func sliceToInt(digits []int) int {
	res := 0
	op := 1
	for i := len(digits) - 1; i >= 0; i-- {
		res += digits[i] * op
		op *= 10
	}
	return res
}

func intToSlice(value int) []int {
	res := []int{}
	for value > 0 {
		newDigit := value % 10
		value = value / 10
		res = append([]int{newDigit}, res...)
	}
	return res
}

func CorrectWay(digits []int) []int {
	res := []int{}
	needToIncrease := false
	addedOne := false
	for i := len(digits) - 1; i >= 0; i-- {
		if !addedOne {
			newV := digits[i] + 1
			addedOne = true
			if newV == 10 {
				newV = 0
				needToIncrease = true
			}
			res = append([]int{newV}, res...)
		} else if needToIncrease {
			newV := digits[i] + 1
			if newV == 10 {
				newV = 0
				needToIncrease = true
			} else {
				needToIncrease = false
			}
			res = append([]int{newV}, res...)
		} else {
			res = append([]int{digits[i]}, res...)
		}
	}
	if needToIncrease {
		res = append([]int{1}, res...)
	}
	return res
}

func main() {
	fmt.Println("Running main")
	result := CorrectWay([]int{
		1, 2, 3,
	})
	fmt.Printf("Got %v\n", result)
}
