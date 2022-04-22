package main

import "fmt"

func DoQuestion(n int) []string {
	res := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		fizzBuzzified := fmt.Sprintf("%d", i)
		moduloThree := i % 3
		moduloFive := i % 5
		if moduloThree == 0 && moduloFive == 0 {
			fizzBuzzified = "FizzBuzz"
		} else if moduloThree == 0 {
			fizzBuzzified = "Fizz"
		} else if moduloFive == 0 {
			fizzBuzzified = "Buzz"
		}
		res = append(res, fizzBuzzified)
	}
	return res
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion(15)
	fmt.Printf("Got %v\n", result)
}
