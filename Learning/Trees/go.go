package main

import (
	"Trees/tree"
	"fmt"
)

func main() {
	fmt.Println("Running main")
	tree := tree.ExampleOne()
	fmt.Printf("Got %v\n", tree)
}
