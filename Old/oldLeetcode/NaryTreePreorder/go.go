package main

import (
	"fmt"
	"time"
)

type Node struct {
	Val      int
	Children []*Node
}

var (
	lvl31 = &Node{
		Val: 5,
	}
	lvl32 = &Node{
		Val: 6,
	}
	lvl21 = &Node{
		Val: 3,
		Children: []*Node{
			lvl31, lvl32,
		},
	}
	lvl22 = &Node{
		Val: 2,
	}
	lvl23 = &Node{
		Val: 4,
	}
	root = &Node{
		Val: 1,
		Children: []*Node{
			lvl21, lvl22, lvl23,
		},
	}
)

func DoQuestion(root *Node) []int {
	res := make([]int, 0)
	res = DoPreorder(root, res)
	return res
}

func DoPreorder(root *Node, res []int) []int {
	if root == nil {
		return res
	}

	res = append(res, root.Val)

	for _, n := range root.Children {
		res = DoPreorder(n, res)
	}

	return res
}

func main() {
	start := time.Now()

	fmt.Println("Running main")
	result := DoQuestion(root)
	fmt.Printf("Got %v\n", result)

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
