package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Solution struct {
	Original []int
	Shuffled []int
}

func Constructor(nums []int) Solution {
	shuffled := make([]int, len(nums))

	for i, v := range nums {
		shuffled[i] = v
	}

	return Solution{
		Original: nums,
		Shuffled: shuffled,
	}
}

func (this *Solution) Reset() []int {
	return this.Original
}

func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	for i := len(this.Shuffled) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		this.Shuffled[i], this.Shuffled[j] = this.Shuffled[j], this.Shuffled[i]
	}

	return this.Shuffled
}

func main() {
	fmt.Println("Running main")

	obj := Constructor([]int{
		1, 2, 3,
	})
	param_1 := obj.Shuffle()
	param_2 := obj.Reset()

	fmt.Printf("Got %v\n", param_1)
	fmt.Printf("Got %v\n", param_2)
}
