package main

import (
	"fmt"
	"math/rand"
	"time"
)

func CreateInputArray() []int {
	originals := make([]int, 100)

	for i := 0; i < 100; i++ {
		originals[i] = i
	}

	return originals
}

func Synchronous(digits []int) []int {
	squared := make([]int, 100)

	for i := 0; i < 100; i++ {
		squared[i] = i * i
	}

	return squared
}

func main() {
	fmt.Println("Running main")

	rand.Seed(time.Now().UTC().UnixNano())

	fmt.Println(rand.Intn(2))

	originals := CreateInputArray()

	result := Synchronous(originals)

	fmt.Printf("Got %v\n", result)
}
