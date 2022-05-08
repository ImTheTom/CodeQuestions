package main

import "fmt"

func DoQuestion(prices []int) int {
	if len(prices) <= 1 {
		return 0
	}

	profit := 0
	bestBuyValue := 100000

	for _, v := range prices {
		if v < bestBuyValue {
			bestBuyValue = v
			continue
		}

		potentialNewProfit := v - bestBuyValue

		if potentialNewProfit > profit {
			profit = potentialNewProfit
		}
	}

	return profit
}

func main() {
	fmt.Println("Running main")
	result := DoQuestion([]int{
		7, 1, 5, 3, 6, 4,
	})
	fmt.Printf("Got %v\n", result)
}
